// Package background
// SPDX-License-Identifier: AGPL-3.0-or-later
package background

import (
	"gogogo/pkg/background/internal"
	"gogogo/pkg/background/request"
	"os"
	"time"
)

func Schedule() func() {
	return schedule
}

func schedule() {
	internal.Logger.Info(time.Now())
	tokenEnv := GetConfig().Github.TokenEnvName
	token := os.Getenv(tokenEnv)
	maintainer := GetConfig().Github.Maintainer
	internal.Logger.Info(request.Zen(token))
	MaintainerInfo, err := request.Users(token, maintainer)
	if err != nil {
		return
	}
	internal.Logger.Info(MaintainerInfo)
	followers, err := request.GetUserFollower(MaintainerInfo, token)
	if err != nil {
		internal.Logger.Info(followers)
	}
}
