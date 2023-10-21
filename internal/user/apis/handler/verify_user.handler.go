package handler

import (
	"context"
	"go-clean-architecture/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userVerifer interface {
	Execute(ctx context.Context, email string) (verifyToken map[string]interface{}, err error)
}

type verifyUserController struct {
	verifyUserUseCase userVerifer
}

// processVerifyUser is a function that processes the verification of a user.
//
// It takes the following parameters:
// - ctx: a context.Context object representing the context of the function call.
// - otp: an integer representing the One-Time Password (OTP) for user verification.
//
// It returns a map[string]interface{} object representing the access token and an error object.
func (c verifyUserController) processVerifyUser(ctx context.Context,
	email string) (verifyToken map[string]interface{}, err error) {
	return c.verifyUserUseCase.Execute(ctx, email)
}

// @Summary Verify user
// @Description Verify user email
// @Tags 		user-service
// @Accept 		json
// @Produce 	json
// @Param 		email query string true "user's email"
// @Success 	200 {object} common.SuccessRes
// @failure 	400 {object} cerr.AppError
// @failure 	500 {object} cerr.AppError
// @Router 		/user/verify [get]
func (h UserHandler) HandleVerifyUser(c *gin.Context) {
	email := c.Query("email")

	verifyToken, err := h.verifyUserController.processVerifyUser(c.Request.Context(), email)

	if err != nil {
		panic(err)
	}

	response := common.SuccessRes{
		StatusCode: http.StatusOK,
		Message:    "OK",
		Data:       verifyToken,
	}

	common.ResponseSuccess(c, response)
}
