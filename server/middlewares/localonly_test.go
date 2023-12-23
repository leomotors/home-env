package middlewares

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsLocalIP(t *testing.T) {
	assert.True(t, isLocalIP("192.168.1.112"))
	assert.True(t, isLocalIP("192.168.10.123"))

	assert.False(t, isLocalIP("48.123.45.67"))
	assert.False(t, isLocalIP("192.166.69.420"))

	assert.True(t, isLocalIP("172.24.0.1"))
	assert.True(t, isLocalIP("172.18.0.1"))

	assert.False(t, isLocalIP("172.32.0.0"))
	assert.False(t, isLocalIP("172.15.0.0"))

	assert.False(t, isLocalIP("invalid string"))
	assert.False(t, isLocalIP("1234:5678:abcd:efgh:ijkl:mnop:qrst:uvwx"))
	assert.False(t, isLocalIP("172.24.0"))
}
