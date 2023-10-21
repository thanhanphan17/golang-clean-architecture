package usecase

import (
	"context"
	"go-clean-architecture/common"
	cerr "go-clean-architecture/common/error"
	biz "go-clean-architecture/internal/user/business"
	"go-clean-architecture/internal/user/business/entity"
	tokentype "go-clean-architecture/provider/tokenprovider/type"
	mail "go-clean-architecture/utils/mail"
	rootdir "go-clean-architecture/utils/rootdir"
	utils "go-clean-architecture/utils/salt"
	"log/slog"
	"time"

	"math/rand"
)

type userCreator interface {
	Execute(ctx context.Context, userEntity entity.User) (map[string]interface{}, error)
}

var _ userCreator = (*createUserUseCase)(nil)

type createUserRepo interface {
	FindUserByCondition(ctx context.Context, condition map[string]interface{},
		preloadKey ...string) (*entity.User, error)
	CreateUser(ctx context.Context, userEntity entity.User) (*entity.User, error)
	UpdateUser(ctx context.Context, userEntity entity.User) (*entity.User, error)
}

type createUserUseCase struct {
	repo          createUserRepo
	tokenProvider TokenProvider
	tokenExpiry   uint
	hashProvider  HashProvider
}

// Execute executes the createUserUseCase.
//
// It takes a context.Context and an entity.User as parameters.
// It returns a map[string]interface{} and an error.
func (useCase *createUserUseCase) Execute(ctx context.Context,
	userEntity entity.User) (map[string]interface{}, error) {
	// Check if the email already exists
	oldUser, _ := useCase.repo.FindUserByCondition(ctx, map[string]interface{}{
		"email": userEntity.Email,
	})
	if oldUser != nil {
		// Email has not been verified
		if oldUser.Status == entity.UNVERIFIED.Value() {
			return nil, biz.ErrEmailNotVerified(nil)
		}

		// Email has already been verified
		return nil, biz.ErrEmailHasExisted(nil)
	}

	salt := utils.GenSalt()
	userEntity.Salt = salt
	userEntity.Password = useCase.hashProvider.Hash(userEntity.Password + userEntity.Salt)

	// Create the new user
	user, err := useCase.repo.CreateUser(ctx, userEntity)
	if err != nil {
		return nil, cerr.ErrCannotCreateEntity(entity.EntityName, err)
	}

	// Generate ramdom OTP and save to db
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	user.OTP = r.Intn(common.OTPMax-common.OTPMin+1) + common.OTPMin

	user, err = useCase.repo.UpdateUser(ctx, *user)
	if err != nil {
		return nil, cerr.ErrInternal(err)
	}

	// Send OTP to user's email
	go func() {
		err = mail.Send(
			user.Email, // to
			"OTP Code", // subject
			rootdir.FindGoModDir()+"/utils/mail/template/otp.html", // template path
			map[string]interface{}{ // data
				"CustomerName": user.Name,
				"OTP":          user.OTP,
			},
		)
		if err != nil {
			slog.Info(err.Error())
			return
		}
	}()

	// Generate the verification token
	verifyToken, err := useCase.tokenProvider.Generate(
		map[string]interface{}{ // payload
			"user_id": user.ID,
			"role":    user.Role,
			"type":    tokentype.VERIFY_TOKEN.Value(),
		},
		useCase.tokenExpiry,
	)
	if err != nil {
		return nil, cerr.ErrInternal(err)
	}
	return verifyToken, nil
}

func NewUserCreator(
	repo createUserRepo,
	tokenProvider interface{},
	tokenExpiry uint,
	hashProvider interface{},
) userCreator {
	return &createUserUseCase{
		repo:          repo,
		tokenProvider: tokenProvider.(TokenProvider),
		tokenExpiry:   tokenExpiry,
		hashProvider:  hashProvider.(HashProvider),
	}
}
