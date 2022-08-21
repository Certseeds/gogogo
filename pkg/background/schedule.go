// Package background
// SPDX-License-Identifier: AGPL-3.0-or-later
package background

import (
	"gogogo/pkg/background/internal"
	"gogogo/pkg/background/request"
	"math"
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
	pages := int64(math.Ceil(float64(MaintainerInfo.Followers) / 30))
	for i := int64(1); i <= pages; i++ {
		internal.Logger.Info(request.Followers(token, MaintainerInfo.Login, i))
	}
}
