// SPDX-License-Identifier: AGPL-3.0-or-later
package background

import (
	"bytes"
	"embed"
	"encoding/gob"
	"encoding/json"
	"gogogo/pkg/background/internal"
)

//go:embed config.json
var configFile embed.FS

type DatabaseRecord struct {
	Address  string `json:"address"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password int    `json:"password"`
}
type GithubRecord struct {
	Maintainer     string `json:"maintainer"`
	FollowingDepth int    `json:"followingDepth"`
	FollowerDepth  int    `json:"followerDepth"`
	TokenEnvName   string `json:"tokenEnvName"`
}
type ScheduleRecord struct {
	Cron string `json:"cron"`
}
type Config struct {
	Database DatabaseRecord `json:"database"`
	Github   GithubRecord   `json:"github"`
	Schedule ScheduleRecord `json:"schedule"`
}

var config Config

func init() {
	raw, err := configFile.ReadFile("config.json")
	if err != nil {
		internal.Logger.Fatal("read config fail")
		return
	}
	json.Unmarshal(raw, &config)
}

// GetConfig
// return a copy of Config, change it is useless
func GetConfig() Config {
	buf := bytes.Buffer{}
	var err error
	if err = gob.NewEncoder(&buf).Encode(config); err != nil {
		return config
	}
	dist := Config{}
	err = gob.NewDecoder(&buf).Decode(&dist)
	if err != nil {
		return config
	}
	return dist
}

func GetPureConfig() (string, error) {
	raw, err := configFile.ReadFile("config.json")
	if err != nil {
		internal.Logger.Fatal("read config fail")
		return "", err
	}
	return string(raw), nil
}
