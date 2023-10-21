package handler

import (
	"context"
	"go-clean-architecture/common"
	cerr "go-clean-architecture/common/error"
	"go-clean-architecture/common/requester"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type userVerifyConfirmer interface {
	Execute(ctx context.Context, otp int) (accessToken map[string]interface{}, err error)
}

type confirmVerifyUserController struct {
	confirmVerifyUserUseCase userVerifyConfirmer
}

// processVerifyUser is a function that processes the verification of a user.
//
// It takes the following parameters:
// - ctx: a context.Context object representing the context of the function call.
// - otp: an integer representing the One-Time Password (OTP) for user verification.
//
// It returns a map[string]interface{} object representing the access token and an error object.
func (c confirmVerifyUserController) processConfirmVerifyUser(ctx context.Context,
	otp int) (accessToken map[string]interface{}, err error) {
	return c.confirmVerifyUserUseCase.Execute(ctx, otp)
}

// @Summary Verify user with OTP
// @Description Verify a user with the by OTP.
// @Tags 		user-service
// @Accept 		json
// @Produce 	json
// @Param 		otp query string true "OTP code"
// @Success 	200 {object} common.SuccessRes
// @failure 	400 {object} cerr.AppError
// @failure 	500 {object} cerr.AppError
// @Router 		/user/confirm-verify [post]
// @Security    JWT
func (h UserHandler) HandleConfirmVerifyUser(c *gin.Context) {
	otp, err := strconv.Atoi(c.Query("otp"))

	if err != nil {
		panic(cerr.ErrInvalidRequest(err))
	}

	requester := c.MustGet(requester.CurrentRequester).(requester.Requester)
	accessToken, err := h.confirmVerifyUserController.processConfirmVerifyUser(
		context.WithValue(c.Request.Context(), common.VerifyTokenKey{}, requester), otp)

	if err != nil {
		panic(err)
	}

	response := common.SuccessRes{
		StatusCode: http.StatusOK,
		Message:    "OK",
		Data:       accessToken,
	}

	common.ResponseSuccess(c, response)
}
