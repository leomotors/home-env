package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTruncateString(t *testing.T) {
	assert.Equal(t, "test", TruncateString("test", 4))
	assert.Equal(t, "test", TruncateString("test", 69))

	assert.Equal(t, "hello worl", TruncateString("hello worl", 10))
	assert.Equal(t, "hello w...", TruncateString("hello world", 10))
	assert.Equal(t, "hello w...", TruncateString("hello world bruh", 10))
}
