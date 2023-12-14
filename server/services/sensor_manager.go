package services

import (
	"math"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type SensorGauge struct {
	temperature  prometheus.Gauge
	humidity     prometheus.Gauge
	healthStatus prometheus.Gauge
}

type SensorValue struct {
	temperature  float64
	humidity     float64
	healthStatus bool
}

type SensorManager struct {
	name   string
	id     string
	gauges SensorGauge
	values SensorValue

	lastUpdated  time.Time
	alertLevel   uint
	alertCounter prometheus.Counter
}

func NewSensorManager(sensorId string, sensorName string) *SensorManager {
	label := prometheus.Labels{
		"sensorId": sensorId,
	}

	newSensor := &SensorManager{
		name: sensorName,
		id:   sensorId,
		gauges: SensorGauge{
			temperature: prometheus.NewGauge(prometheus.GaugeOpts{
				Name:        "home_temperature",
				Help:        "Current temperature in degrees Celsius.",
				ConstLabels: label,
			}),
			humidity: prometheus.NewGauge(prometheus.GaugeOpts{
				Name:        "home_humidity",
				Help:        "Current humidity level as a percentage.",
				ConstLabels: label,
			}),
			healthStatus: prometheus.NewGauge(prometheus.GaugeOpts{
				Name:        "home_health_status",
				Help:        "ESP32 is maintaining connection with server.",
				ConstLabels: label,
			}),
		},
		alertCounter: prometheus.NewCounter(prometheus.CounterOpts{
			Name:        "home_health_alert_sent",
			Help:        "Number of alerts sent to Discord",
			ConstLabels: label,
		}),
	}

	return newSensor
}

func (sensorManager *SensorManager) Register() {
	prometheus.MustRegister(sensorManager.gauges.temperature)
	prometheus.MustRegister(sensorManager.gauges.humidity)
	prometheus.MustRegister(sensorManager.gauges.healthStatus)

	prometheus.MustRegister(sensorManager.alertCounter)

	sensorManager.SetValue(math.NaN(), math.NaN())
}

func (sensorManager *SensorManager) SetValue(temperature float64, humidity float64) {
	sensorManager.gauges.temperature.Set(temperature)
	sensorManager.values.temperature = temperature

	sensorManager.gauges.humidity.Set(humidity)
	sensorManager.values.humidity = humidity

	sensorManager.gauges.healthStatus.Set(1)
	sensorManager.values.healthStatus = true

	sensorManager.lastUpdated = time.Now()

	if sensorManager.alertLevel > 0 {
		sensorManager.alertLevel = 0
		SendBackNotice(sensorManager.id)
	}
}

func (sensorManager *SensorManager) HealthCheck() {
	idleTime := time.Since(sensorManager.lastUpdated).Seconds()

	if idleTime >= 15 {
		sensorManager.gauges.healthStatus.Set(0)
		sensorManager.values.healthStatus = false
	}

	if MeetsThreshold(sensorManager.alertLevel, idleTime) {
		SendDownAlert(sensorManager.id, sensorManager.alertLevel)
		sensorManager.alertCounter.Inc()
		sensorManager.alertLevel++
	}
}

func (sensorManager *SensorManager) LastUpdated() float64 {
	return time.Since(sensorManager.lastUpdated).Seconds()
}
