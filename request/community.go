package request

type CommunityList struct {
	Page     int64 `json:"page" form:"page" example:"1"`
	PageSize int64 `json:"page_size" form:"page_size"  example:"20"`
}
