package background

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfig(t *testing.T) {
	assert.NotEmpty(t, GetConfig(), "embed to binary")
}
func TestPureConfig(t *testing.T) {
	pureConfig, err := GetPureConfig()
	assert.Nil(t, err, "read config should not throw err")
	assert.NotEmpty(t, pureConfig, "pure str should not be empty")
}
