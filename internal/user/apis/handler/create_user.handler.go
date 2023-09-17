package handler

import (
	"context"
	cerr "go-clean-architecture/common/error"
	"go-clean-architecture/common/res"
	"go-clean-architecture/internal/user/apis/mapper"
	"go-clean-architecture/internal/user/apis/req"
	"go-clean-architecture/internal/user/business/usecase"
	"go-clean-architecture/middleware"
	"go-clean-architecture/provider/tokenprovider"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type createUserController struct {
	validatorRequest  *validator.Validate
	createUserUseCase usecase.UserCreator
}

func (c createUserController) processCreateUser(
	ctx context.Context,
	req req.CreateUserReq,
) (*tokenprovider.Token, error) {
	return c.createUserUseCase.Execute(ctx, mapper.TranformCreateUserReq(req))
}

func (h UserHandler) HandleCreateUser(c *gin.Context) {
	var createUserReq req.CreateUserReq

	if err := c.BindJSON(&createUserReq); err != nil {
		panic(cerr.ErrInvalidRequest(err))
	}

	if err := h.createUserController.
		validatorRequest.Struct(createUserReq); err != nil {
		panic(err)
	}

	verifyToken, err := h.createUserController.processCreateUser(
		c.Request.Context(),
		createUserReq,
	)

	if err != nil {
		panic(err)
	}

	respone := res.OK{
		StatusCode: http.StatusCreated,
		Message:    "OK",
		RequestID:  middleware.GetRequestIDFromContext(c),
		Data:       verifyToken,
	}

	res.ResponseOK(c, respone)
}
