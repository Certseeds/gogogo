package background

import (
	"gogogo/pkg/background/database"
	"gogogo/pkg/background/internal"
	"gogogo/pkg/background/request"
	"sync"
)

func SyncGithubUser(userName string, token string, depth int) {
	MaintainerInfo, err := request.Users(token, userName)
	if err != nil {
		return
	}
	connection := GetDataBaseConnection()
	database.SyncGitHubUser(*MaintainerInfo, connection)
	internal.Logger.Info(MaintainerInfo)
	followers, err := request.GetUserFollower(MaintainerInfo, token)
	if err != nil {
		return
	}
	toId, _ := database.GetIdByUserName(userName, connection)
	wg1 := sync.WaitGroup{}
	for index, follower := range *followers {
		go func(toId int64, follower request.GitHubUser, wg2 *sync.WaitGroup) {
			wg2.Add(1)
			defer wg2.Done()
			database.SyncGitHubUser(follower, connection)
			fromId, _ := database.GetIdByUserName(follower.Login, connection)
			internal.Logger.Info("from: ", *fromId, ", to: ", toId)
			database.SyncFollowConnection(*fromId, toId, connection)
		}(*toId, follower, &wg1)
		internal.Logger.Info(index, " follower update finish")
		//internal.Logger.Info(follower)
	}
	wg1.Wait()
	if depth == 0 {
		return
	}
	for _, follower := range *followers {
		go SyncGithubUser(follower.Login, token, depth-1)
	}
}
