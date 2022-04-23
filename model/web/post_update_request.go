package web

type PostUpdateRequest struct {
	Id          int    `validate:"required" json:"id"`
	Title       string `validate:"required,max=100" json:"title"`
	Description string `validate:"required,max=256" json:"description"`
}
