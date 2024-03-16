package db

func GetVotesByName(name string) (votes int64, err error) {
	sql := "SELECT votes FROM `t_user_votes` where name = ?"
	result := db.Raw(sql, name).Scan(&votes)
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}
