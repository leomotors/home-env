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

// PublicSensorValue is the public-facing sensor data
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

// PublicSensorDetail includes identity and current readings for a sensor
type PublicSensorDetail struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Temperature *float64 `json:"temperature"`
	Humidity    *float64 `json:"humidity"`
	LastUpdated float64  `json:"lastUpdated"`
	Online      bool     `json:"online"`
}

func GetAllSensors() []PublicSensorDetail {
	result := make([]PublicSensorDetail, 0, len(sensors))

	for _, sensor := range sensors {
		detail := PublicSensorDetail{
			Id:          sensor.id,
			Name:        sensor.name,
			LastUpdated: time.Since(sensor.lastUpdated).Seconds(),
			Online:      sensor.values.healthStatus,
		}
		if sensor.values.healthStatus {
			detail.Temperature = &sensor.values.temperature
			detail.Humidity = &sensor.values.humidity
		}
		result = append(result, detail)
	}

	return result
}
