package request

import (
	"encoding/json"
	"fmt"
	"gogogo/pkg/background/internal"
	"math"
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
func GetUserFollower(user *GitHubUser, token string) (*[]GitHubUser, error) {
	pages := int64(math.Ceil(float64(user.Followers) / 30))
	doneChallen := make(chan bool, pages)
	userChan := make(chan *[]GitHubUser, pages)
	for i := int64(1); i <= pages; i++ {
		go func(token string, userLogin string, pageId int64, done chan bool, userChan chan *[]GitHubUser) {
			followers, err := Followers(token, userLogin, pageId)
			if err != nil {
				done <- false
				userChan <- nil
			}
			done <- true
			userChan <- followers
		}(token, user.Login, i, doneChallen, userChan)
	}
	userLists := make([]*[]GitHubUser, pages)
	for i := int64(0); i < pages; i++ {
		<-doneChallen
		userLists[i] = <-userChan
	}
	willReturn := make([]GitHubUser, user.Followers)
	count := 0
	for _, userList := range userLists {
		for _, user := range *userList {
			willReturn[count] = user
			internal.Logger.Info(user)
			count += 1
		}
	}
	return &willReturn, nil
}
