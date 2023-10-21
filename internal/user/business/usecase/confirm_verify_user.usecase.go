package usecase

import (
	"context"
	"go-clean-architecture/common"
	cerr "go-clean-architecture/common/error"
	"go-clean-architecture/common/requester"
	biz "go-clean-architecture/internal/user/business"
	"go-clean-architecture/internal/user/business/entity"
	tokentype "go-clean-architecture/provider/tokenprovider/type"
)

type userVerifyConfirmer interface {
	Execute(ctx context.Context, otp int) (accessToken map[string]interface{}, err error)
}

var _ userVerifyConfirmer = (*confirmVerifyUserUsecase)(nil)

type comfirmVerifyUserRepo interface {
	FindUserByCondition(ctx context.Context, condition map[string]interface{},
		preloadKey ...string) (*entity.User, error)
	VerifyUser(ctx context.Context, userID string) error
}

type confirmVerifyUserUsecase struct {
	repo          comfirmVerifyUserRepo
	tokenProvider TokenProvider
	tokenExpiry   uint
}

// Execute executes the confirmVerifyUserUsecase.
//
// It takes the following parameters:
// - ctx (context.Context): The context.
// - otp (int): The OTP.
//
// It returns the following:
// - accessToken (map[string]interface{}): The access token.
// - err (error): The error, if any.
func (useCase *confirmVerifyUserUsecase) Execute(ctx context.Context,
	otp int) (accessToken map[string]interface{}, err error) {
	requester := ctx.Value(common.VerifyTokenKey{}).(requester.Requester)

	user, err := useCase.repo.FindUserByCondition(ctx, map[string]interface{}{
		"id": requester.GetUserId(),
	})
	if err != nil {
		return nil, cerr.ErrEntityNotFound(entity.EntityName, err)
	}

	if user.OTP != otp {
		return nil, biz.ErrInvalidOTP(nil)
	}

	if user.Status == entity.ACTIVE.Value() {
		return nil, biz.ErrEmailAlreadyVerified(nil)
	}

	if err := useCase.repo.VerifyUser(ctx, user.ID); err != nil {
		return nil, cerr.ErrInternal(err)
	}

	// Generate the access token
	accessToken, err = useCase.tokenProvider.Generate(
		map[string]interface{}{ // payload
			"user_id": user.ID,
			"role":    user.Role,
			"type":    tokentype.ACCESS_TOKEN.Value(),
		},
		useCase.tokenExpiry,
	)
	if err != nil {
		return nil, cerr.ErrInternal(err)
	}

	return accessToken, nil
}

func NewUserVerifyConfirmer(
	repo verifyUserRepo,
	tokenProvider interface{},
	tokenExpiry uint,
) userVerifyConfirmer {
	return &confirmVerifyUserUsecase{
		repo:          repo,
		tokenProvider: tokenProvider.(TokenProvider),
		tokenExpiry:   tokenExpiry,
	}
}
