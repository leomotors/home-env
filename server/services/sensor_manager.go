package services

import (
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
	name        string
	id          string
	gauges      SensorGauge
	values      SensorValue
	lastUpdated time.Time
	alertLevel  uint
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
		sensorManager.alertLevel++
	}
}

func (sensorManager *SensorManager) LastUpdated() float64 {
	return time.Since(sensorManager.lastUpdated).Seconds()
}
