package repository

import (
	"app-download/ds"
	"app-download/dto"
	"app-download/model"
	"context"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type InstallLogRepository struct {
	DB  *gorm.DB
	RDB *redis.Client
}

func NewInstallLogRepository(ds *ds.DataSource) *InstallLogRepository {
	return &InstallLogRepository{
		DB:  ds.DB,
		RDB: ds.RDB,
	}
}

func (r *InstallLogRepository) Create(ctx context.Context, data *model.InstallLog) (*model.InstallLog, error) {
	if err := r.DB.Model(&model.InstallLog{}).Create(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (r *InstallLogRepository) FindAll(data *model.InstallLog, pagination *dto.Pagination) (*[]model.InstallLog, error) {
	var logs []model.InstallLog
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuider := r.DB.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)
	result := queryBuider.Model(&model.InstallLog{}).Where(data).Find(&logs)
	if result.Error != nil {
		msg := result.Error
		return nil, msg
	}
	return &logs, nil
}
