package mysql

import (
	"database/sql"
	"gin-web/models"
)

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

// FindUserByUsername 根据用户名获取用户
func FindUserByUsername(username string) (obj *models.User, err error) {
	obj = new(models.User)
	sqlStr := `select * from users where username = ?`
	err = db.Get(obj, sqlStr, username)
	//  err 为 sql.ErrNoRows 表示用户不存在
	if err == sql.ErrNoRows {
		err = ErrUserNotExist
	}
	return
}
