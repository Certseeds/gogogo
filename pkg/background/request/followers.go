package request

import (
	"encoding/json"
	"fmt"
)

type FollowUser struct {
	Login        string `json:"login"`
	Id           int64  `json:"id"`
	NodeId       string `json:"node_id"`
	FollowersUrl string `json:"followers_url"`
	FollowingUrl string `json:"following_url"`
}

func Followers(token string, username string, page int64) (*[]GitHubUser, error) {
	resp, err := GetRequester(fmt.Sprintf("https://api.github.com/users/%s/followers?per_page=30&page=%d", username, page), token)
	if err != nil {
		return nil, err
	}
	var userInfo []GitHubUser
	json.Unmarshal(resp, &userInfo)
	return &userInfo, nil
}
