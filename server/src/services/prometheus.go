package services

import (
	"math"

	"github.com/prometheus/client_golang/prometheus"
)

var sensors = make(map[string]*SensorManager)

func RegisterSensor(sensorId string, sensorName string) {
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
	}

	sensors[sensorId] = newSensor

	prometheus.MustRegister(newSensor.gauges.temperature)
	prometheus.MustRegister(newSensor.gauges.humidity)
	prometheus.MustRegister(newSensor.gauges.healthStatus)

	newSensor.SetValue(math.NaN(), math.NaN())
}

func GetSensorManager(sensorId string) *SensorManager {
	return sensors[sensorId]
}

func GetSensorValue(sensorId string) SensorValue {
	return sensors[sensorId].values
}

func HealthCheck() {
	for _, sensor := range sensors {
		sensor.HealthCheck()
	}
}

func GetAllSensorHealth() *map[string]bool {
	status := make(map[string]bool)

	for _, sensor := range sensors {
		status[sensor.name] = sensor.values.healthStatus
	}

	return &status
}
