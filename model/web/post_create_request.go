package web

type PostCreateRequest struct {
	Title       string `validate:"required,max=100" json:"title"`
	Description string `validate:"required,max=256" json:"description"`
}
