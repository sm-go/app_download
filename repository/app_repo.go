package repository

import (
	"app-download/ds"
	"app-download/dto"
	"app-download/model"
	"context"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type AppRepository struct {
	DB  *gorm.DB
	RDB *redis.Client
}

func NewAppRepository(ds *ds.DataSource) *AppRepository {
	return &AppRepository{
		DB:  ds.DB,
		RDB: ds.RDB,
	}
}

func (r *AppRepository) All(ctx context.Context) []model.App {
	var applist []model.App
	r.DB.Model(&model.App{}).Find(&applist)
	return applist
}

func (r *AppRepository) FindAll(data *model.App, pagination *dto.Pagination) (*[]model.App, error) {
	var apps []model.App
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuilder := r.DB.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)
	result := queryBuilder.Model(&model.App{}).Where(data).Find(&apps)
	if result.Error != nil {
		msg := result.Error
		return nil, msg
	}
	return &apps, nil
}

func (r *AppRepository) Create(ctx context.Context, data *model.App) (*model.App, error) {
	if err := r.DB.Model(&model.App{}).Create(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (r *AppRepository) UpdateApp(ctx context.Context, data *model.App) (*model.App, error) {
	if err := r.DB.Model(&model.App{}).Where("id = ?", data.Id).Updates(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (r *AppRepository) DeleteApp(ctx context.Context, app *model.App) error {
	if err := r.DB.Where("id = ?", app.Id).Delete(&app).Error; err != nil {
		return err
	}
	return nil
}
