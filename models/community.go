package models

import "time"

type Community struct {
	// `db:"community_id"` 指定字段名
	Id            int64  `json:"community_id,omitempty" db:"community_id"`
	CommunityName string `json:"community_name,omitempty"`
	Introduction  string `json:"introduction,omitempty"`

	// 也可以返回给前端时间戳，前端根据用户所在的时区进行转换，这样子如果是一个国际化的应用就比较兼容
	CreatedAt *time.Time `json:"created_at"`
}
