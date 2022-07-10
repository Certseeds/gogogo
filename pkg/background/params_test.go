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
func TestConfigImmutable(t *testing.T) {
	assert.NotEmpty(t, GetConfig(), "embed to binary")
	assert.NotEqual(t, len(GetConfig().Github.TokenEnvName), 0)
	config := GetConfig()
	config.Github = Github{
		0,
		0,
		"",
	}
	assert.NotEqual(t, len(GetConfig().Github.TokenEnvName), 0)
}
