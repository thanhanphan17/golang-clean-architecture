package handler

import (
	"context"
	"go-clean-architecture/common"
	"go-clean-architecture/common/requester"
	"net/http"

	"github.com/gin-gonic/gin"
)

type otpSender interface {
	Execute(ctx context.Context) error
}

type sendOTPController struct {
	sendOTPUseCase otpSender
}

// processSendOTP is a function that processes the sending of an OTP (One-Time Password).
//
// It takes in the following parameters:
//   - ctx: the context of the request
//
// It return errors that occurred.
func (c sendOTPController) processSendOTP(ctx context.Context) error {
	return c.sendOTPUseCase.Execute(ctx)
}

// @Summary 	Resend OTP
// @Description Resend One-Time-Password (use verify token responsed from verify API)
// @Tags 		user-service
// @Accept  	json
// @Produce  	json
// @Success 	200 {object} common.SuccessRes
// @failure 	400 {object} cerr.AppError
// @failure 	500 {object} cerr.AppError
// @Router 		/user/otp-resend [get]
// @Security    JWT
func (h UserHandler) HandleSendOTP(c *gin.Context) {

	requester := c.MustGet(requester.CurrentRequester).(requester.Requester)
	ctx := context.WithValue(c.Request.Context(), common.VerifyTokenKey{}, requester)
	err := h.sendOTPController.processSendOTP(ctx)
	if err != nil {
		panic(err)
	}

	response := common.SuccessRes{
		StatusCode: http.StatusOK,
		Message:    "OK",
		Data:       "OTP was sent",
	}

	common.ResponseSuccess(c, response)
}
