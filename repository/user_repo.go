package repository

import (
	"app-download/ds"
	"app-download/model"
	"context"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB  *gorm.DB
	RDB *redis.Client
}

func NewUserRepository(ds *ds.DataSource) *UserRepository {
	return &UserRepository{
		DB:  ds.DB,
		RDB: ds.RDB,
	}
}

func (r *UserRepository) FindByUsername(ctx context.Context, username string) (*model.User, error) {
	user := &model.User{}
	if err := r.DB.Model(&model.User{}).Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) FindByUserId(ctx context.Context, userId uint64) (*model.User, error) {
	user := &model.User{}
	if err := r.DB.Model(&model.User{}).Where("id = ?", userId).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
