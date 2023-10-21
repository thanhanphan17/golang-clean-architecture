package usecase

import (
	"context"
	"go-clean-architecture/common"
	cerr "go-clean-architecture/common/error"
	"go-clean-architecture/common/requester"
	biz "go-clean-architecture/internal/user/business"
	"go-clean-architecture/internal/user/business/entity"
	utils "go-clean-architecture/utils/mail"
	rootdir "go-clean-architecture/utils/rootdir"
	"time"

	"math/rand"
)

type otpResender interface {
	Execute(ctx context.Context) error
}

var _ otpResender = (*resendOTPUseCase)(nil)

type resendOTPRepo interface {
	FindUserByCondition(ctx context.Context, condition map[string]interface{},
		preloadKey ...string) (*entity.User, error)
	UpdateUser(ctx context.Context, userEntity entity.User) (*entity.User, error)
}

type resendOTPUseCase struct {
	repo resendOTPRepo
}

// Execute executes the sendOTP use case.
//
// Parameters:
// - ctx: the context.Context object for the execution.
// - email: the email to which OTP will be sent.
//
// Returns:
// - err: an error, if any occurred during the execution.
func (useCase *resendOTPUseCase) Execute(ctx context.Context) error {
	requester := ctx.Value(common.VerifyTokenKey{}).(requester.Requester)

	// Check if the email already exists
	user, err := useCase.repo.FindUserByCondition(ctx, map[string]interface{}{
		"id": requester.GetUserId(),
	})
	if err != nil {
		return biz.ErrEmailIsNotExisted(err)
	}

	// Generate ramdom OTP and save to db
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	user.OTP = r.Intn(common.OTPMax-common.OTPMin+1) + common.OTPMin

	user, err = useCase.repo.UpdateUser(ctx, *user)
	if err != nil {
		return cerr.ErrInternal(err)
	}

	// Send OTP to user's email
	err = utils.Send(
		user.Email, // to
		"OTP Code", // subject
		rootdir.FindGoModDir()+"/utils/mail/template/otp.html", // template path
		map[string]interface{}{ // data
			"CustomerName": user.Name,
			"OTP":          user.OTP,
		},
	)
	if err != nil {
		return cerr.ErrInternal(err)
	}

	return nil
}

func NewOTPResender(
	repo resendOTPRepo,
) otpResender {
	return &resendOTPUseCase{
		repo: repo,
	}
}
