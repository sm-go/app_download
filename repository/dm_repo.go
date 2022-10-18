package repository

import (
	"app-download/ds"
	"app-download/model"
	"context"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type DmRepository struct {
	DB  *gorm.DB
	RDB *redis.Client
}

func NewDmRepository(ds *ds.DataSource) *DmRepository {
	return &DmRepository{
		DB:  ds.DB,
		RDB: ds.RDB,
	}
}

func (r *DmRepository) All(ctx context.Context) []model.Domain {
	var dmlist []model.Domain
	r.DB.Model(&model.Domain{}).Find(&dmlist)
	return dmlist
}

func (r *DmRepository) CreateDm(ctx context.Context, data *model.Domain) (*model.Domain, error) {
	if err := r.DB.Model(&model.Domain{}).Create(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (r *DmRepository) UpdateDomain(ctx context.Context, data *model.Domain) (*model.Domain, error) {
	if err := r.DB.Model(&model.Domain{}).Where("id = ?", data.Id).Updates(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (r *DmRepository) DeleteDomain(ctx context.Context, data *model.Domain) error {
	if err := r.DB.Model(&model.Domain{}).Where("id = ?", data.Id).Delete(&data).Error; err != nil {
		return err
	}
	return nil
}
