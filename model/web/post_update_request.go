package web

type PostUpdateRequest struct {
	Id          int    `validate:"required"`
	Title       string `validate:"required,max=100"`
	Description string `validate:"required,max=256"`
}
