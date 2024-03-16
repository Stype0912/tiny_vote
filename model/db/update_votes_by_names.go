package db

import "gorm.io/gorm"

func UpdateVotesByNames(names []string) (bool, error) {
	var result *gorm.DB
	for _, name := range names {
		var sql string
		sql = "UPDATE `t_user_votes` SET votes=votes+1 where name = ?"
		result = db.Exec(sql, name)
		if result.Error != nil {
			return false, result.Error
		}
		if result.RowsAffected == 0 {
			sql = "INSERT INTO `t_user_votes` (name, votes) VALUES (?, ?)"
			result = db.Exec(sql, name, 1)
			if result.Error != nil {
				return false, result.Error
			}
		}
	}
	return true, nil
}
