package database

// login is in the 'src' column
func GetFollowingByUserName(from int, connection string) (*[]int, error) {
	readConfig(connection)
	row, err := mysql.Query("select dst from github_follow where src = ?", from)
	if err != nil {
		return nil, err
	}
	target := make([]int, 0)
	for row.Next() {
		var to int
		_ = row.Scan(&to)
		target = append(target, to)
	}
	return &target, nil
}

// SyncFollowConnection
// from and to are both `id` in table 'github_user'
func SyncFollowConnection(from int64, to int64, connection string) (*struct{}, error) {
	readConfig(connection)
	row := mysql.QueryRow("select id from github_follow where src = ? and dst = ?", from, to)
	var id int64
	err := row.Scan(&id)
	if err == nil {
		return nil, nil
	}
	_, err = mysql.Exec("insert into github_follow (src, dst) values (?,?)", from, to)
	if err != nil {
		return nil, err
	}
	return &struct{}{}, nil
}
