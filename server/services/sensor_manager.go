package services

import (
	"math"
	"time"
)

type SensorValue struct {
	temperature  float64
	humidity     float64
	healthStatus bool
}

type SensorManager struct {
	name string
	id   string

	values SensorValue

	lastUpdated time.Time
	alertLevel  uint
}

func NewSensorManager(sensorId string, sensorName string) *SensorManager {
	return &SensorManager{
		name: sensorName,
		id:   sensorId,
	}
}

func (sensorManager *SensorManager) Register() {
	sensorManager.values.temperature = math.NaN()
	sensorManager.values.humidity = math.NaN()
	sensorManager.values.healthStatus = true
	sensorManager.lastUpdated = time.Now()
}

// SetValue updates sensor readings, pushes to TimescaleDB, and resets health status
func (sensorManager *SensorManager) SetValue(temperature float64, humidity float64) {
	sensorManager.values.temperature = temperature
	sensorManager.values.humidity = humidity
	sensorManager.values.healthStatus = true
	sensorManager.lastUpdated = time.Now()

	if !math.IsNaN(temperature) && !math.IsNaN(humidity) {
		go InsertReading(sensorManager.id, temperature, humidity)
	}

	if sensorManager.alertLevel > 0 {
		go InsertDowntimeEvent(sensorManager.id, "resolved")
		sensorManager.alertLevel = 0
		SendBackNotice(sensorManager.id)
	}
}

// HealthCheck checks sensor idle time, marks offline after 15s,
// and sends Discord alerts with escalating thresholds
func (sensorManager *SensorManager) HealthCheck() {
	idleTime := time.Since(sensorManager.lastUpdated).Seconds()

	if idleTime >= 15 {
		sensorManager.values.healthStatus = false
		sensorManager.values.temperature = math.NaN()
		sensorManager.values.humidity = math.NaN()
	}

	if MeetsThreshold(sensorManager.alertLevel, idleTime) {
		if sensorManager.alertLevel == 0 {
			go InsertDowntimeEvent(sensorManager.id, "down")
		}
		SendDownAlert(sensorManager.id, sensorManager.alertLevel)
		sensorManager.alertLevel++
	}
}

func (sensorManager *SensorManager) LastUpdated() float64 {
	return time.Since(sensorManager.lastUpdated).Seconds()
}
