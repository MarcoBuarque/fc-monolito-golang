package repository

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return ProductRepository{db}
}

func (repo ProductRepository) FindAll(ctx context.Context) ([]ProductData, error) {
	var response []ProductData

	if result := repo.db.WithContext(ctx).Find(&response); result.Error != nil {
		// TODO: add log
		return []ProductData{}, result.Error
	}

	return response, nil
}

func (repo ProductRepository) Find(ctx context.Context, id string) (ProductData, error) {
	if id == "" {
		return ProductData{}, fmt.Errorf("id cannot be empty")
	}

	response := &ProductData{ID: id}
	if result := repo.db.WithContext(ctx).First(response); result.Error != nil {
		// TODO: add log
		return ProductData{}, result.Error
	}

	return *response, nil
}
