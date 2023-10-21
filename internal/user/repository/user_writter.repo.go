package repository

import (
	"context"
	cerr "go-clean-architecture/common/error"
	"go-clean-architecture/db"
	"go-clean-architecture/internal/user/business/entity"
)

type userWriterImpl struct {
	db db.Database
}

// CreateUser implements UserRepoWriter.
func (repo *userWriterImpl) CreateUser(ctx context.Context,
	userEntity entity.User) (*entity.User, error) {
	if err := repo.db.Executor.Create(&userEntity).Error; err != nil {
		return nil, err
	}

	return &userEntity, nil
}

// UpdateUser implements UserRepoWriter.
func (repo *userWriterImpl) UpdateUser(ctx context.Context,
	userEntity entity.User) (*entity.User, error) {
	if err := repo.db.Executor.Updates(&userEntity).Error; err != nil {
		return nil, err
	}

	return &userEntity, nil
}

// VerifyUser implements UserRepoWriter.
func (repo *userWriterImpl) VerifyUser(ctx context.Context, userID string) error {
	if err := repo.db.Executor.
		Table(entity.User{}.TableName()).
		Where("id = ?", userID).
		Updates(map[string]interface{}{
			"status": entity.ACTIVE.Value(),
		}).Error; err != nil {
		return cerr.ErrDB(err)
	}

	return nil
}

var _ userRepoWriter = (*userWriterImpl)(nil)

func NewUserWriterImpl(db db.Database) userRepoWriter {
	return &userWriterImpl{
		db: db,
	}
}
