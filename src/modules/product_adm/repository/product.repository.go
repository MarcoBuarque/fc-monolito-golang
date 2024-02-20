package repository

import (
	"context"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func (repo ProductRepository) Add(ctx context.Context, data ProductData) (ProductData, error) {
	if result := repo.db.WithContext(ctx).Create(&data); result.Error != nil {
		return ProductData{}, result.Error
	}
	return data, nil
}

func (repo ProductRepository) Find(ctx context.Context, id string) (ProductData, error) {
	response := &ProductData{ID: id}
	if result := repo.db.WithContext(ctx).First(response); result.Error != nil {
		return ProductData{}, result.Error
	}
	return *response, nil
}
