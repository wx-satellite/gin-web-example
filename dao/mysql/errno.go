package mysql

import "errors"

// ----- 数据库错误 -----

var (
	ErrUserNotExist = errors.New("用户不存在！")

	ErrInvalidId = errors.New("无效的ID")
)
