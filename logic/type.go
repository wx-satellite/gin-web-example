package logic

import "gin-web/models"

type ApiPostDetail struct {
	AuthorName string `json:"author_name"`
	*models.Community
	*models.Post
}
