package model

func UpdateVotesByNames(names []string) (bool, error) {
	sql := "UPDATE `t_user_votes` SET votes=votes+1 where name in ?"
	res := DB.Exec(sql, names)
	if res.Error != nil {
		return false, err
	}
	return true, nil
}
