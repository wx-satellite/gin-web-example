package models

import "time"

type PostStatus int8

const (
	PostStatusUnknown PostStatus = iota
	PostStatusDown
	PostStatusUp
)

type Post struct {
	// 模型的ID对应到数据的post_id，关于数据库的id在这个项目中不需要，因为分布式的应用表ID会有重复的
	Id          int64      `json:"id" db:"post_id"`
	Title       string     `json:"title"`        // 标题
	Content     string     `json:"content"`      // 内容
	AuthorId    int64      `json:"author_id"`    // 作者
	CommunityId int64      `json:"community_id"` // 所属社区
	Status      PostStatus `json:"status"`       // 状态
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}
