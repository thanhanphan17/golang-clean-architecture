package repository

import (
	"context"
	cerr "go-clean-architecture/common/error"
	paging "go-clean-architecture/common/pagination"
	"go-clean-architecture/db"
	"go-clean-architecture/internal/user/business/entity"

	"gorm.io/gorm"
)

type userFinderImpl struct {
	db db.Database
}

// FindUserByID implements UserRepoFinder.
func (u *userFinderImpl) FindUserByID(
	ctx context.Context,
	userID string,
) (*entity.User, error) {
	userEntity := entity.User{}

	if err := u.db.Executor.
		Where("id = ?", userID).
		First(&userEntity).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, cerr.ErrRecordNotFound(nil)
		}

		return nil, cerr.ErrDB(err)
	}

	return &userEntity, nil
}

// FindUserByPhone implements UserRepoFinder.
func (u *userFinderImpl) FindUserByPhone(
	ctx context.Context,
	phone string,
) (*entity.User, error) {
	userEntity := entity.User{}

	if err := u.db.Executor.
		Where("phone = ?", phone).
		First(&userEntity).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, cerr.ErrRecordNotFound(nil)
		}

		return nil, cerr.ErrDB(err)
	}

	return &userEntity, nil
}

// FindUsers implements UserRepoFinder.
func (u *userFinderImpl) FindUsers(
	ctx context.Context,
	filter entity.User,
	page int,
	limit int,
) (paging.Pagination, error) {
	panic("unimplemented")
}

var _ userRepoFinder = (*userFinderImpl)(nil)

func NewUserFinderImpl(db db.Database) userRepoFinder {
	return &userFinderImpl{
		db: db,
	}
}
