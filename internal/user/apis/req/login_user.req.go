package req

type LoginUserReq struct {
	Email    string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"required"`
}
