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

type userLoginer interface {
	Execute(ctx context.Context, entity entity.User) (map[string]interface{}, error)
}

type loginUserController struct {
	validatorRequest *validator.Validate
	loginUserUseCase userLoginer
}

// processLoginUser processes the login user request.
//
// It takes a context and a LoginUserReq as parameters.
// It returns a map[string]interface{} and an error.
func (c loginUserController) processLoginUser(ctx context.Context,
	req req.LoginUserReq) (map[string]interface{}, error) {
	return c.loginUserUseCase.Execute(ctx, mapper.TransformLogineUserReq(req))
}

// @Summary 	Login user account
// @Description Login user account with "email" and "password"
// @Tags 		user-service
// @Accept  	json
// @Produce  	json
// @Param 		data body req.LoginUserReq true "user"
// @Success 	200 {object} common.SuccessRes
// @failure 	400 {object} cerr.AppError
// @failure 	500 {object} cerr.AppError
// @Router 		/user/login [post]
func (h UserHandler) HandleLoginUser(c *gin.Context) {
	var loginUserReq req.LoginUserReq

	if err := c.BindJSON(&loginUserReq); err != nil {
		panic(cerr.ErrInvalidRequest(err))
	}

	if err := h.loginUserController.validatorRequest.Struct(loginUserReq); err != nil {
		panic(cerr.ErrInvalidRequest(err))
	}

	token, err := h.loginUserController.processLoginUser(c.Request.Context(), loginUserReq)
	if err != nil {
		panic(err)
	}

	response := common.SuccessRes{
		StatusCode: http.StatusOK,
		Message:    "OK",
		Data:       token,
	}
	common.ResponseSuccess(c, response)
}
