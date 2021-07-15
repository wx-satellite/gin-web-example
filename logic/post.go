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

// PostDetail 根据帖子ID获取帖子详情
func PostDetail(id int64) (res *ApiPostDetail, err error) {
	// 查询并组合接口想要的数据
	return
}
