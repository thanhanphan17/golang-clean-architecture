package usecase

import (
	"context"
	cerr "go-clean-architecture/common/error"
	biz "go-clean-architecture/internal/user/business"
	"go-clean-architecture/internal/user/business/entity"
	"go-clean-architecture/provider/tokenprovider"
)

type UserCreator interface {
	Execute(
		ctx context.Context,
		userEntity entity.User,
	) (*tokenprovider.Token, error)
}

type createUserRepo interface {
	FindUserByPhone(ctx context.Context, phone string) (*entity.User, error)
	CreateUser(ctx context.Context, userEntity entity.User) (*entity.User, error)
}

type createUserUseCase struct {
	repo          createUserRepo
	tokenProvider tokenprovider.Provider
	expiry        int
}

// Execute implements UserCreator.
func (u *createUserUseCase) Execute(
	ctx context.Context,
	userEntity entity.User,
) (*tokenprovider.Token, error) {
	oldUser, _ := u.repo.FindUserByPhone(ctx, userEntity.Phone)

	// phone is existed
	if oldUser != nil {
		// phone has not been verified
		if oldUser.Status == entity.NOT_VERIFIED.Value() {
			return nil, biz.ErrPhoneHasNotVerified(nil)
		}

		return nil, biz.ErrPhoneHasExisted(nil)
	}

	user, err := u.repo.CreateUser(ctx, userEntity)

	if err != nil {
		return nil, err
	}

	// create payload
	payload := tokenprovider.TokenPayload{
		UserID: *user.ID,
		Role:   user.Role,
		Type:   "verify",
	}

	verifyToken, err := u.tokenProvider.Generate(payload, u.expiry)

	if err != nil {
		return nil, cerr.ErrInternal(err)
	}

	return verifyToken, nil
}

var _ UserCreator = (*createUserUseCase)(nil)

func NewUserCreator(
	repo createUserRepo,
	tokenProvider tokenprovider.Provider,
	expiry int,
) UserCreator {
	return &createUserUseCase{
		repo:          repo,
		tokenProvider: tokenProvider,
		expiry:        expiry,
	}
}
