package model

func GetVotesByName(name string) (votes int64) {
	sql := "SELECT votes FROM `t_user_votes` where name = ?"
	DB.Raw(sql, name).Scan(&votes)
	return
}
