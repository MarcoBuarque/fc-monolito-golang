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
