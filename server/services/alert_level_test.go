package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMeetsThreshold(t *testing.T) {
	for i := 0; i < 30; i++ {
		for j := uint(0); j < 10; j++ {
			assert.False(t, MeetsThreshold(j, float64(i)))
		}
	}

	for i := 30; i < 60; i++ {
		assert.True(t, MeetsThreshold(0, float64(i)))

		for j := 1; j < 10; j++ {
			assert.False(t, MeetsThreshold(1, float64(i)))
		}
	}

	assert.True(t, MeetsThreshold(1, 60))
	assert.True(t, MeetsThreshold(1, 69))
	assert.False(t, MeetsThreshold(2, 69))

	assert.True(t, MeetsThreshold(2, 600))
	assert.True(t, MeetsThreshold(2, 699))
	assert.False(t, MeetsThreshold(3, 699))
	assert.False(t, MeetsThreshold(3, 3599))

	assert.True(t, MeetsThreshold(3, 3600))

	assert.False(t, MeetsThreshold(4, 7199))
	assert.True(t, MeetsThreshold(4, 7200))

	assert.False(t, MeetsThreshold(^uint(0), 6942069420))
}

func TestGetAlertText(t *testing.T) {
	assert.Equal(t, "30秒", GetAlertText(0))
	assert.Equal(t, "1分", GetAlertText(1))
	assert.Equal(t, "10分", GetAlertText(2))
	assert.Equal(t, "1時間", GetAlertText(3))
	assert.Equal(t, "2時間", GetAlertText(4))
	assert.Equal(t, "3時間", GetAlertText(5))
	assert.Equal(t, "69時間", GetAlertText(71))
}
