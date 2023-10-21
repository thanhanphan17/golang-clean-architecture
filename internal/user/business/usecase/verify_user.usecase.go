package usecase

import (
	"context"
	"go-clean-architecture/common"
	cerr "go-clean-architecture/common/error"
	biz "go-clean-architecture/internal/user/business"
	"go-clean-architecture/internal/user/business/entity"
	tokentype "go-clean-architecture/provider/tokenprovider/type"
	utils "go-clean-architecture/utils/mail"
	rootdir "go-clean-architecture/utils/rootdir"
	"log/slog"
	"math/rand"
	"time"
)

type userVerifier interface {
	Execute(ctx context.Context, email string) (accessToken map[string]interface{}, err error)
}

var _ userVerifier = (*verifyUserUsecase)(nil)

type verifyUserRepo interface {
	FindUserByCondition(ctx context.Context, condition map[string]interface{},
		preloadKey ...string) (*entity.User, error)
	VerifyUser(ctx context.Context, userID string) error
	UpdateUser(ctx context.Context, userEntity entity.User) (*entity.User, error)
}

type verifyUserUsecase struct {
	repo          verifyUserRepo
	tokenProvider TokenProvider
	tokenExpiry   uint
}

// Execute executes the verifyUser use case.
//
// It takes a context and an email as parameters.
// It returns a map[string]interface{} containing the verify token,
// and an error.
func (usecase *verifyUserUsecase) Execute(ctx context.Context,
	email string) (verifyToken map[string]interface{}, err error) {

	user, err := usecase.repo.FindUserByCondition(ctx, map[string]interface{}{
		"email": email,
	})
	if err != nil {
		return nil, cerr.ErrEntityNotFound(entity.EntityName, err)
	}

	if user.Status == entity.ACTIVE.Value() {
		return nil, biz.ErrEmailAlreadyVerified(nil)
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	user.OTP = r.Intn(common.OTPMax-common.OTPMin+1) + common.OTPMin
	user, err = usecase.repo.UpdateUser(ctx, *user)
	if err != nil {
		return nil, cerr.ErrInternal(err)
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
		slog.Error(err.Error())
	}

	// Generate the access token
	verifyToken, err = usecase.tokenProvider.Generate(
		map[string]interface{}{ // payload
			"user_id": user.ID,
			"role":    user.Role,
			"type":    tokentype.VERIFY_TOKEN.Value(),
		},
		usecase.tokenExpiry,
	)
	if err != nil {
		return nil, cerr.ErrInternal(err)
	}

	return verifyToken, nil
}

func NewUserVerifier(
	repo verifyUserRepo,
	tokenProvider interface{},
	tokenExpiry uint,
) userVerifier {
	return &verifyUserUsecase{
		repo:          repo,
		tokenProvider: tokenProvider.(TokenProvider),
		tokenExpiry:   tokenExpiry,
	}
}
