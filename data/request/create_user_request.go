package request

type CreateUserRequest struct {
	Name string `validate:"required,max = 10, min =1" json:"name"`
}
