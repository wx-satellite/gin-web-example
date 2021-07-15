package mysql

import (
	"database/sql"
	"gin-web/models"
)

func Communities() (objs []*models.Community, err error) {
	objs = make([]*models.Community, 0, 20)
	sqlStr := `select community_id, community_name from communities`
	// sql.ErrNoRows  查询不到数据
	if err = db.Select(&objs, sqlStr); err != nil && err != sql.ErrNoRows {
		return
	}
	err = nil
	return
}

// FindCommunity 根据社区ID获取社区信息
func FindCommunity(id int64) (obj *models.Community, err error) {
	obj = new(models.Community)
	sqlStr := `select * from communities where id = ?`
	if err = db.Get(obj, sqlStr, id); err != nil {
		if err == sql.ErrNoRows {
			err = ErrInvalidId
		}
		return
	}
	return
}
