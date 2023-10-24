package request

type CreateBlogRequest struct {
	Title       string `json:"title" form:"title" validate:"required"`
	Description string `json:"description" form:"description" validate:"required"`
	Image       string `json:"image" form:"image" validate:"required"`
}

type UpdateBlogRequest struct {
	Title       string `json:"title" form:"title" validate:"required"`
	Description string `json:"description" form:"description" validate:"required"`
	Image       string `json:"image" form:"image" validate:"required"`
}
