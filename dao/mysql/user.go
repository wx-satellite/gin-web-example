package mysql

import "gin-web/models"

// CheckUserExist 指定用户名的用户是否存在
func CheckUserExist(username string) (bool, error) {
	var count int
	sqlStr := `select count(1) from users where username = ?`
	err := db.Get(&count, sqlStr, username)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// InsertUser 新增用户信息
func InsertUser(obj *models.User) (id int64, err error) {
	sqlStr := `insert into users(user_id,username,password) values(?,?,?)`
	result, err := db.Exec(sqlStr, obj.UserId, obj.Username, obj.Password)
	if err != nil {
		return
	}
	return result.LastInsertId()
}
