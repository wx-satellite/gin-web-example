package request

type CreatePostRequest struct {
	Title       string `json:"title" binding:"required"`        // 标题
	Content     string `json:"content" binding:"required"`      // 内容
	CommunityId int64  `json:"community_id" binding:"required"` // 所属社区
	AuthorId    int64  `json:"-"`
}
