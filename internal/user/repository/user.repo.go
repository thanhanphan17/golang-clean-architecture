package repository

import (
	"context"
	paging "go-clean-architecture/common/pagination"
	"go-clean-architecture/internal/user/business/entity"
)

type userRepoFinder interface {
	FindUserByID(ctx context.Context, userID string) (*entity.User, error)
	FindUserByPhone(ctx context.Context, phone string) (*entity.User, error)
	FindUsers(ctx context.Context, filter entity.User, page, limit int) (paging.Pagination, error)
}

type userRepoWriter interface {
	CreateUser(ctx context.Context, usrentity entity.User) (*entity.User, error)
	VerifyUser(ctx context.Context, userID string) error
	UpdateUser(ctx context.Context, user entity.User) (*entity.User, error)
}

type UserRepository struct {
	userRepoFinder
	userRepoWriter
}

func NewUserRepository(
	userFinder userRepoFinder,
	userWriter userRepoWriter,
) *UserRepository {
	return &UserRepository{
		userRepoFinder: userFinder,
		userRepoWriter: userWriter,
	}
}
