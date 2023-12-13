package services_test

import (
	"testing"

	"github.com/leomotors/home-env/services"
	"github.com/stretchr/testify/assert"
)

func TestMeetsThreshold(t *testing.T) {
	for i := 0; i < 30; i++ {
		for j := uint(0); j < 10; j++ {
			assert.False(t, services.MeetsThreshold(j, float64(i)))
		}
	}

	for i := 30; i < 60; i++ {
		assert.True(t, services.MeetsThreshold(0, float64(i)))

		for j := 1; j < 10; j++ {
			assert.False(t, services.MeetsThreshold(1, float64(i)))
		}
	}

	assert.True(t, services.MeetsThreshold(1, 60))
	assert.True(t, services.MeetsThreshold(1, 69))
	assert.False(t, services.MeetsThreshold(2, 69))

	assert.True(t, services.MeetsThreshold(2, 600))
	assert.True(t, services.MeetsThreshold(2, 699))
	assert.False(t, services.MeetsThreshold(3, 699))
	assert.False(t, services.MeetsThreshold(3, 3599))

	assert.True(t, services.MeetsThreshold(3, 3600))

	assert.False(t, services.MeetsThreshold(4, 7199))
	assert.True(t, services.MeetsThreshold(4, 7200))

	assert.False(t, services.MeetsThreshold(^uint(0), 6942069420))
}

func TestGetAlertText(t *testing.T) {
	assert.Equal(t, "30秒", services.GetAlertText(0))
	assert.Equal(t, "1分", services.GetAlertText(1))
	assert.Equal(t, "10分", services.GetAlertText(2))
	assert.Equal(t, "1時間", services.GetAlertText(3))
	assert.Equal(t, "2時間", services.GetAlertText(4))
	assert.Equal(t, "3時間", services.GetAlertText(5))
	assert.Equal(t, "69時間", services.GetAlertText(71))
}
