package logic

import (
	"errors"
	"gin-web/dao/mysql"
	"gin-web/models"
	"gin-web/pkg/encryption"
	"gin-web/pkg/jwt"
	"gin-web/pkg/snowflake"
	"gin-web/request"
)

var (
	ErrUserNameExist = errors.New("用户名已经存在")
)

func SignUp(in *request.SignUpRequest) (err error) {
	var exist bool
	if exist, err = mysql.CheckUserExist(in.Username); nil != err {
		return
	}
	if exist {
		err = ErrUserNameExist
		return
	}
	obj := new(models.User)
	// 生成分布式ID
	obj.UserId = snowflake.GetID()
	obj.Username = in.Username
	// 密码加密
	obj.Password = encryption.Md5(in.Password)
	_, err = mysql.InsertUser(obj)
	return
}

func SignIn(in *request.SignInRequest) (token string, err error) {
	obj := new(models.User)
	return jwt.GenToken(obj.UserId)
}
