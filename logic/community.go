package logic

import (
	"gin-web/dao/mysql"
	"gin-web/models"
)

func GetCommunityList() (data []*models.Community, err error) {
	return mysql.Communities()
}

func CommunityDetail(id int64) (obj *models.Community, err error) {
	return
}
