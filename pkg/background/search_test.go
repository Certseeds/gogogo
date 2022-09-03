package background

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"gogogo/pkg/background/internal"
	"gogogo/pkg/background/request"
	"os"
	"testing"
)

func TestSyncGithubUser(t *testing.T) {
	tokenEnv := GetConfig().Github.TokenEnvName
	token := os.Getenv(tokenEnv)
	maintainer := GetConfig().Github.Maintainer
	internal.Logger.Info(request.Zen(token))

	SyncGithubUser(maintainer, token, 0)
}

func TestFollowers(t *testing.T) {
	config := GetConfig()
	assert.NotEmpty(t, config.Github.Maintainer, "maintainer should not be empty")
	tokenEnv := config.Github.TokenEnvName
	assert.NotEmpty(t, tokenEnv, "tokenEnv should not be empty")
	token := os.Getenv(tokenEnv)
	user, err := request.Users(token, config.Github.Maintainer)
	assert.Nilf(t, err, "err should be nil")
	assert.NotEmpty(t, user.Login, "login name should not be empty")
	followers, err := request.GetUserFollower(user, token)
	assert.Nilf(t, err, "err should be nil")
	assert.NotEmpty(t, followers, fmt.Sprintf("followers should not be empty %s", user))
}

func TestZen(t *testing.T) {
	config := GetConfig()
	tokenEnv := config.Github.TokenEnvName
	token := os.Getenv(tokenEnv)
	zen, err := request.Zen(token)
	assert.Nilf(t, err, "err should be nil")
	assert.NotNilf(t, zen, "zen should not be nil")
	assert.NotEmpty(t, zen, "zen should not be empty")

}
