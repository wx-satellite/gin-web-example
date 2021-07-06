package logic

import (
	"fmt"
	"gin-web/dao/mysql"
	"gin-web/models"
	"gin-web/pkg/encryption"
	"gin-web/pkg/snowflake"
	"gin-web/request"
)

func SignUp(in *request.SignUpRequest) (err error) {
	var exist bool
	if exist, err = mysql.CheckUserExist(in.Username); nil != err {
		return
	}
	if exist {
		err = fmt.Errorf("该用户名：%s 已经存在", in.Username)
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
