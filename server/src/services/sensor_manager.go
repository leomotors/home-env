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
	temperatureGauge float64
	humidity         float64
	healthStatus     bool
}

type SensorManager struct {
	name        string
	gauges      SensorGauge
	values      SensorValue
	lastUpdated time.Time
	alertLevel  int
}

func (sensorManager *SensorManager) SetValue(temperature float64, humidity float64) {
	sensorManager.gauges.temperature.Set(temperature)
	sensorManager.values.temperatureGauge = temperature

	sensorManager.gauges.humidity.Set(humidity)
	sensorManager.values.humidity = humidity

	sensorManager.gauges.healthStatus.Set(1)
	sensorManager.values.healthStatus = true

	sensorManager.lastUpdated = time.Now()

	// TODO Manage alert level
	sensorManager.alertLevel = 0
}
