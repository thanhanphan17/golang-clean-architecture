package req

type CreateUserReq struct {
	Email    string `json:"email" validate:"email,required"`
	Name     string `json:"name" validate:"min=10,max=30,required"`
	Role     string `json:"role" validate:"required"`
	Password string `json:"password" validate:"password,required"`
}
