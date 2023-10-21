package usecase

import (
	"context"
	cerr "go-clean-architecture/common/error"
	biz "go-clean-architecture/internal/user/business"
	"go-clean-architecture/internal/user/business/entity"
	tokentype "go-clean-architecture/provider/tokenprovider/type"
)

type userLoginer interface {
	Execute(ctx context.Context,
		userEntity entity.User) (token map[string]interface{}, err error)
}

var _ userLoginer = (*loginUserUseCase)(nil)

type loginUserRepo interface {
	FindUserByCondition(ctx context.Context, condition map[string]interface{},
		preloadKey ...string) (*entity.User, error)
	UpdateUser(ctx context.Context, userEntity entity.User) (*entity.User, error)
}

type loginUserUseCase struct {
	repo              loginUserRepo
	tokenProvider     TokenProvider
	accessTokenExpiry uint
	hashProvider      HashProvider
}

// Execute is a function that executes the login use case.
//
// It takes a context.Context and a userEntity of type entity.User as parameters.
// It returns an accessToken of type map[string]interface{} and an error.
func (useCase *loginUserUseCase) Execute(ctx context.Context,
	userEntity entity.User) (token map[string]interface{}, err error) {
	// Check if the email already exists
	user, err := useCase.repo.FindUserByCondition(ctx, map[string]interface{}{
		"email": userEntity.Email,
	})
	if err != nil {
		return nil, biz.ErrEmailIsNotExisted(err)
	}

	if user != nil {
		// Email has not been verified
		if user.Status == entity.UNVERIFIED.Value() {
			return nil, biz.ErrEmailNotVerified(nil)
		}
	}

	// Check password
	password := useCase.hashProvider.Hash(userEntity.Password + user.Salt)
	if password != user.Password {
		return nil, biz.ErrInvalidInfo(nil)
	}

	// Generate the access token
	accessToken, err := useCase.tokenProvider.Generate(
		map[string]interface{}{ // payload
			"user_id": user.ID,
			"role":    user.Role,
			"type":    tokentype.ACCESS_TOKEN.Value(),
		},
		useCase.accessTokenExpiry,
	)
	if err != nil {
		return nil, cerr.ErrInternal(err)
	}

	return accessToken, nil
}

func NewUserLoginer(
	repo loginUserRepo,
	tokenProvider interface{},
	accessTokenExpiry uint,
	hashProvider interface{},
) userLoginer {
	return &loginUserUseCase{
		repo:              repo,
		tokenProvider:     tokenProvider.(TokenProvider),
		accessTokenExpiry: accessTokenExpiry,
		hashProvider:      hashProvider.(HashProvider),
	}
}
