package services

import (
	"time"
)

var sensors = make(map[string]*SensorManager)

func RegisterSensor(sensorId string, sensorName string) {
	newSensor := NewSensorManager(sensorId, sensorName)
	newSensor.Register()

	sensors[sensorId] = newSensor
}

func GetSensorManager(sensorId string) *SensorManager {
	return sensors[sensorId]
}

type PublicSensorValue struct {
	Temperature float64
	Humidity    float64
	LastUpdated float64
}

func GetSensorValue(sensorId string) PublicSensorValue {
	manager := sensors[sensorId]

	return PublicSensorValue{
		manager.values.temperature,
		manager.values.humidity,
		time.Since(manager.lastUpdated).Seconds(),
	}
}

// Should be called every 5 seconds
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
