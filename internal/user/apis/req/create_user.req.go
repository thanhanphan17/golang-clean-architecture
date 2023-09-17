package req

type CreateUserReq struct {
	Phone    string `json:"phone" validate:"vnphone,required"`
	Name     string `json:"name" validate:"min=10,max=30,required"`
	Password string `json:"password" validate:"password,required"`
}
