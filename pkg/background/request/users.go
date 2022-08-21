package request

import (
	"encoding/json"
	"fmt"
)

type GitHubUser struct {
	Login        string `json:"login"`
	NodeId       string `json:"node_id"`
	Id           int64  `json:"id"`
	FollowersUrl string `json:"followers_url"`
	Followers    int64  `json:"followers"`
	Blog         string `json:"blog"` // this is the link in user homepage that under `Location` and behind `Achievements`
}

func Users(token string, username string) (*GitHubUser, error) {
	resp, err := GetRequester(fmt.Sprintf("https://api.github.com/users/%s", username), token)
	if err != nil {
		return nil, err
	}
	var userInfo GitHubUser
	json.Unmarshal(resp, &userInfo)
	return &userInfo, nil
}
