package database

import (
	"gogogo/pkg/background/request"
)

func GetUserByUserName(login string, connection string) (*request.GitHubUser, error) {
	readConfig(connection)
	row := mysql.QueryRow("select login,githubNodeId,githubId,githubBlog from github_user where login = ?", login)
	target := request.GitHubUser{}
	err := row.Scan(&target.Login, &target.NodeId, &target.Id, &target.Blog)
	if err != nil {
		return nil, err
	}
	return &target, nil
}
func GetIdByUserName(login string, connection string) (*int64, error) {
	readConfig(connection)
	row := mysql.QueryRow("select id from github_user where login = ?", login)
	var target int64
	err := row.Scan(&target)
	if err != nil {
		return nil, err
	}
	return &target, nil
}
func getIdByGithubId(githubId int64, connection string) (*int64, error) {
	readConfig(connection)
	row := mysql.QueryRow("select id from github_user where githubId = ?", githubId)
	var target int64
	err := row.Scan(&target)
	if err != nil {
		return nil, err
	}
	return &target, nil
}
func SyncGitHubUser(user request.GitHubUser, connection string) bool {
	readConfig(connection)
	id, _ := getIdByGithubId(user.Id, connection)
	if id != nil {
		return updateGitHubUser(user, connection)
	} else {
		return insertGitHubUser(user, connection)
	}
}
func insertGitHubUser(user request.GitHubUser, connection string) bool {
	readConfig(connection)
	_, err := mysql.Exec("insert into github_user(login, githubNodeId, githubId, githubBlog,twitterUsername)values (?,?,?,?,?)",
		user.Login, user.NodeId, user.Id, user.Blog, user.TwitterUserName)
	if err != nil {
		return false
	}
	return true
}
func updateGitHubUser(user request.GitHubUser, connection string) bool {
	readConfig(connection)
	_, err := mysql.Exec("update github_user set login = ?,githubNodeId = ?, githubBlog = ?,twitterUsername=? where githubId = ?",
		user.Login, user.NodeId, user.Blog, user.TwitterUserName, user.Id)
	if err != nil {
		return false
	}
	return true
}
