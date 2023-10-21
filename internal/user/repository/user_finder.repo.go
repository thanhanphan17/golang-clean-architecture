package repository

import (
	"context"
	"errors"
	cerr "go-clean-architecture/common/error"
	paging "go-clean-architecture/common/pagination"
	"go-clean-architecture/db"
	"go-clean-architecture/internal/user/business/entity"

	"gorm.io/gorm"
)

type userFinderImpl struct {
	db db.Database
}

func (repo *userFinderImpl) FindUserByCondition(ctx context.Context,
	condition map[string]interface{}, preloadKey ...string) (*entity.User, error) {
	userEntity := entity.User{}
	db := repo.db.Executor.Model(&userEntity)

	// Preload all keys
	for _, key := range preloadKey {
		db = db.Preload(key)
	}

	if err := db.Where(condition).First(&userEntity).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, cerr.ErrRecordNotFound(nil)
		}

		return nil, cerr.ErrDB(err)
	}

	return &userEntity, nil
}

// FindUsers implements UserRepoFinder.
func (repo *userFinderImpl) FindAllUsers(ctx context.Context,
	filter *entity.Filter, pagination *paging.Pagination) (*paging.Pagination, error) {
	var users []*entity.User

	db := repo.db.Executor.
		Table(entity.User{}.TableName())

	db = db.Scopes(applyFilterScope(filter))

	if err := db.
		Scopes(paging.Paginate(users, pagination, db)).
		Find(&users).Error; err != nil {
		return nil, err
	}

	pagination.Rows = users

	return pagination, nil

}

var _ userRepoFinder = (*userFinderImpl)(nil)

func NewUserFinderImpl(db db.Database) userRepoFinder {
	return &userFinderImpl{
		db: db,
	}
}
