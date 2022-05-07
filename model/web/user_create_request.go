package web

type UserCreateRequest struct {
	Name     string `validate:"required,max=100" json:"name"`
	Age      int    `validate:"required,max=100" json:"age"`
	Email    string `validate:"required,max=100,email" json:"email"`
	Image    string `json:"image"`
	Password string `validate:"required,min=6,max=256" json:"password"`
}
