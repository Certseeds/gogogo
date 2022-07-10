package background

import (
	"bytes"
	"embed"
	_ "embed"
	"encoding/gob"
	"encoding/json"
	_ "encoding/json"
	_ "fmt"
	"log"
)

//go:embed config.json
var configFile embed.FS

type Database struct {
	Address  string `json:"address"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password int    `json:"password"`
}
type Github struct {
	FollowingDepth int    `json:"followingDepth"`
	FollowerDepth  int    `json:"followerDepth"`
	TokenEnvName   string `json:"tokenEnvName"`
}
type Config struct {
	Database Database `json:"database"`
	Github   Github   `json:"github"`
}

var config Config

func init() {
	raw, err := configFile.ReadFile("config.json")
	if err != nil {
		log.Fatalln("read config fail")
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
		log.Fatalln("read config fail")
		return "", err
	}
	return string(raw), nil
}
