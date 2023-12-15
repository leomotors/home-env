package services

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSensorManager(t *testing.T) {
	sensor := NewSensorManager("testid", "testname")

	// Not nil object
	assert.NotNil(t, sensor)

	// Check name and id
	assert.Equal(t, "testid", sensor.id)
	assert.Equal(t, "testname", sensor.name)

	// Check gauges not nil
	assert.NotNil(t, sensor.gauges.temperature)
	assert.NotNil(t, sensor.gauges.humidity)
	assert.NotNil(t, sensor.gauges.healthStatus)
	assert.NotNil(t, sensor.counters.alertSent)
	assert.NotNil(t, sensor.counters.pingReceived)

	sensor.Register()

	// Check initial values
	assert.True(t, math.IsNaN(sensor.values.temperature))
	assert.True(t, math.IsNaN(sensor.values.humidity))
	assert.Equal(t, true, sensor.values.healthStatus)

	assert.Equal(t, uint(0), sensor.alertLevel)

	// Set Value Test
	sensor.SetValue(69, 420)
	assert.Equal(t, float64(69), sensor.values.temperature)
	assert.Equal(t, float64(420), sensor.values.humidity)
	assert.Equal(t, true, sensor.values.healthStatus)
}
