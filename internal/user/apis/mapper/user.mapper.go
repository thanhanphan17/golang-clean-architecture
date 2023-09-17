package mapper

import (
	"go-clean-architecture/internal/user/apis/req"
	"go-clean-architecture/internal/user/business/entity"
)

func TranformCreateUserReq(req req.CreateUserReq) entity.User {
	return entity.User{
		Phone:    req.Phone,
		Name:     req.Name,
		Password: req.Password,
		Status:   entity.BLOCK.Value(),
	}
}
