package handler

import (
	"go-clean-architecture/internal/user/business/usecase"
	"go-clean-architecture/internal/user/repository"

	"github.com/go-playground/validator/v10"
)

type UserHandler struct {
	createUserController
	loginUserController
	confirmVerifyUserController
	sendOTPController
	verifyUserController
}

// NewUserHandler creates a new UserHandler instance.
//
// Parameters:
// - repo: the UserRepository implementation.
// - validatorRequest: the request validator.
// - tokenProvider: the token provider.
// - verifyTime: the verification time.
// - accessTime: the access time.
// - hasher: the hasher implementation.
//
// Return:
// - *UserHandler: a pointer to the UserHandler instance.
func NewUserHandler(
	repo repository.UserRepository,
	validatorRequest *validator.Validate,
	tokenProvider interface{},
	verifyTime uint,
	accessTime uint,
	hasher interface{},
) *UserHandler {
	return &UserHandler{
		createUserController: createUserController{
			validatorRequest: validatorRequest,
			createUserUseCase: usecase.NewUserCreator(
				repo,
				tokenProvider,
				verifyTime,
				hasher,
			),
		},
		loginUserController: loginUserController{
			validatorRequest: validatorRequest,
			loginUserUseCase: usecase.NewUserLoginer(
				repo,
				tokenProvider,
				accessTime,
				hasher,
			),
		},
		verifyUserController: verifyUserController{
			verifyUserUseCase: usecase.NewUserVerifier(
				repo,
				tokenProvider,
				verifyTime,
			),
		},
		confirmVerifyUserController: confirmVerifyUserController{
			confirmVerifyUserUseCase: usecase.NewUserVerifyConfirmer(
				repo,
				tokenProvider,
				accessTime,
			),
		},
		sendOTPController: sendOTPController{
			sendOTPUseCase: usecase.NewOTPResender(
				repo,
			),
		},
	}
}
