package handler

import (
	"go-clean-architecture/internal/user/business/usecase"
	"go-clean-architecture/internal/user/repository"
	"go-clean-architecture/provider/tokenprovider"

	"github.com/go-playground/validator/v10"
)

type UserHandler struct {
	createUserController
}

func NewUserHandler(
	repo repository.UserRepository,
	validatorRequest *validator.Validate,
	tokenProvider tokenprovider.Provider,
	verifyTime int,
	accessTime int,
) *UserHandler {
	return &UserHandler{
		createUserController: createUserController{
			validatorRequest: validatorRequest,
			createUserUseCase: usecase.NewUserCreator(
				repo,
				tokenProvider,
				verifyTime,
			),
		},
	}
}
