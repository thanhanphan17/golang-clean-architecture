package repository

import (
	"context"
	paging "go-clean-architecture/common/pagination"
	"go-clean-architecture/internal/user/business/entity"
)

type userRepoFinder interface {
	FindUserByCondition(ctx context.Context,
		condition map[string]interface{}, preloadKey ...string) (*entity.User, error)
	FindAllUsers(ctx context.Context, filter *entity.Filter,
		pagination *paging.Pagination) (*paging.Pagination, error)
}

type userRepoWriter interface {
	CreateUser(ctx context.Context, userEntity entity.User) (*entity.User, error)
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
