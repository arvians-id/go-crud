package web

type PostCreateRequest struct {
	Title       string `validate:"required,max=100"`
	Description string `validate:"required,max=256"`
}
