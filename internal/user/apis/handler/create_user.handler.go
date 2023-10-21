package handler

import (
	"context"
	"go-clean-architecture/common"
	cerr "go-clean-architecture/common/error"
	"go-clean-architecture/internal/user/apis/mapper"
	"go-clean-architecture/internal/user/apis/req"
	"go-clean-architecture/internal/user/business/entity"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type userCreator interface {
	Execute(ctx context.Context, entity entity.User) (map[string]interface{}, error)
}

type createUserController struct {
	validatorRequest  *validator.Validate
	createUserUseCase userCreator
}

// processCreateUser processes the creation of a user.
//
// ctx: the context.Context to use for the operation.
// req: the CreateUserReq containing the user information.
// Returns: a map[string]interface{} containing the user data,
// and an error if there was a problem executing the operation.
func (c createUserController) processCreateUser(ctx context.Context,
	req req.CreateUserReq) (map[string]interface{}, error) {
	return c.createUserUseCase.Execute(ctx, mapper.TransformCreateUserReq(req))
}

// @Summary 	Create user account
// @Description Create user account with "email", "name" and "password"
// @Tags 		user-service
// @Accept  	json
// @Produce  	json
// @Param 		data body req.CreateUserReq true "user"
// @Success 	200 {object} common.SuccessRes
// @failure 	400 {object} cerr.AppError
// @failure 	500 {object} cerr.AppError
// @Router 		/user/register [post]
func (h UserHandler) HandleCreateUser(c *gin.Context) {
	var createUserReq req.CreateUserReq

	if err := c.BindJSON(&createUserReq); err != nil {
		panic(cerr.ErrInvalidRequest(err))
	}

	if err := h.createUserController.validatorRequest.Struct(createUserReq); err != nil {
		panic(cerr.ErrInvalidRequest(err))
	}

	verifyToken, err := h.createUserController.processCreateUser(c.Request.Context(), createUserReq)
	if err != nil {
		panic(err)
	}

	response := common.SuccessRes{
		StatusCode: http.StatusCreated,
		Message:    "OK",
		Data:       verifyToken,
	}

	common.ResponseSuccess(c, response)
}
