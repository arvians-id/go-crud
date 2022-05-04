package web

type UserCreateRequest struct {
	Name     string `validate:"required,max=100" json:"name"`
	Age      int    `validate:"required,max=100" json:"age"`
	Email    string `validate:"required,max=100" json:"email"`
	Password string `validate:"required" json:"password"`
}