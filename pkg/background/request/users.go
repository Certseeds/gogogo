package request

import (
	"encoding/json"
	"errors"
	"fmt"
	"gogogo/pkg/background/internal"
)

type GitHubUser struct {
	Login           string `json:"login"`
	NodeId          string `json:"node_id"`
	Id              int64  `json:"id"`
	FollowersUrl    string `json:"followers_url"`
	Followers       int64  `json:"followers"`
	Following       int64  `json:"following"`
	Blog            string `json:"blog"` // this is the link in user homepage that under `Location` and behind `Achievements`
	TwitterUserName string `json:"twitter_username"`
}

func (user GitHubUser) String() string {
	return fmt.Sprintf("User: %s, NodeId: %s, Id: %d, Followers: %d,Following: %d,\nBlog %s,TwitterUserName: %s",
		user.Login, user.NodeId, user.Id, user.Followers, user.Following, user.Blog, user.TwitterUserName,
	)
}
func Users(token string, username string) (*GitHubUser, error) {
	resp, err := GetRequester(fmt.Sprintf("https://api.github.com/users/%s", username), token)
	if err != nil {
		return nil, err
	}
	var userInfo GitHubUser
	json.Unmarshal(resp, &userInfo)
	if userInfo.Following > 450 && userInfo.Followers > 450 {
		internal.Logger.Info("this is a bugger", userInfo)
		return &userInfo, errors.New("this is a bugger user")
	}
	return &userInfo, nil
}
