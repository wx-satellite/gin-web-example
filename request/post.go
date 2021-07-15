package request

type CreatePostRequest struct {
	Title       string `json:"title" binding:"required"`        // 标题
	Content     string `json:"content" binding:"required"`      // 内容
	CommunityId int64  `json:"community_id" binding:"required"` // 所属社区
	AuthorId    int64  `json:"-"`
}

type VotePostRequest struct {
	PostId    int64 `json:"post_id,string" binding:"required"`       // 防止ID传递到前端失真，使用字符串传递
	Direction int8  `json:"direction" binding:"required,oneof=1 -1"` // 1 表示赞成 -1 表示反对
	UserId    int64 `json:"-"`                                       // 当前登陆用户
}
