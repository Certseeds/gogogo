package request

import (
	"encoding/json"
	"fmt"
	"math"
)

type FollowUser struct {
	Login        string `json:"login"`
	Id           int64  `json:"id"`
	NodeId       string `json:"node_id"`
	FollowersUrl string `json:"followers_url"`
	FollowingUrl string `json:"following_url"`
}

func followers(token string, username string, page int64) (*[]GitHubUser, error) {
	resp, err := GetRequester(fmt.Sprintf("https://api.github.com/users/%s/followers?per_page=30&page=%d", username, page), token)
	if err != nil {
		return nil, err
	}
	var userInfo []GitHubUser
	// the follower will be the last level of rec, so just get for pages will make get less cost
	_ = json.Unmarshal(resp, &userInfo)
	return &userInfo, nil
}
func GetUserFollower(user *GitHubUser, token string) (*[]GitHubUser, error) {
	pages_ := int64(math.Ceil(float64(user.Followers) / 30))
	pages := int64(math.Min(float64(pages_), 5))
	doneChallen := make(chan bool, pages)
	userChan := make(chan *[]GitHubUser, pages)
	for i := int64(1); i <= pages; i++ {
		go func(token string, userLogin string, pageId int64, done chan bool, userChan chan *[]GitHubUser) {
			followers, err := followers(token, userLogin, pageId)
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
			// internal.Logger.Info(user)
			count += 1
		}
	}
	return &willReturn, nil
}
