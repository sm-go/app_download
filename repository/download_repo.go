package repository

import (
	"app-download/ds"
	"app-download/dto"
	"app-download/model"
	"context"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type DownloadLogRepo struct {
	DB  *gorm.DB
	RDB *redis.Client
}

func NewDownloadLogRepository(ds *ds.DataSource) *DownloadLogRepo {
	return &DownloadLogRepo{
		DB:  ds.DB,
		RDB: ds.RDB,
	}
}

func (r *DownloadLogRepo) Create(ctx context.Context, data *model.DownloadLog) (*model.DownloadLog, error) {
	if err := r.DB.Model(&model.DownloadLog{}).Create(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (r *DownloadLogRepo) All(ctx context.Context) []model.DownloadLog {
	var downloadlists []model.DownloadLog
	r.DB.Find(&downloadlists)
	return downloadlists
}

func (r *DownloadLogRepo) FindAll(data *model.DownloadLog, pagination *dto.Pagination) (*[]model.DownloadLog, error) {
	var logs []model.DownloadLog
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuider := r.DB.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)
	result := queryBuider.Model(&model.DownloadLog{}).Where(data).Find(&logs)
	if result.Error != nil {
		msg := result.Error
		return nil, msg
	}
	return &logs, nil
}
