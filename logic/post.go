package logic

import (
	"gin-web/models"
	"gin-web/pkg/snowflake"
	"gin-web/request"
)

func CreatePost(in *request.CreatePostRequest) (err error) {
	obj := new(models.Post)
	obj.Id = snowflake.GetID()
	obj.Status = models.PostStatusUp
	obj.Title = in.Title
	obj.CommunityId = in.CommunityId
	obj.Content = in.Content
	obj.AuthorId = in.AuthorId
	// 插入数据库

	return
}
