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

func (repo ProductRepository) Add(ctx context.Context, data Product) error {
	if result := repo.db.WithContext(ctx).Create(&data); result.Error != nil {
		// TODO: add log
		return result.Error
	}

	return nil
}

func (repo ProductRepository) Find(ctx context.Context, id string) (Product, error) {
	if id == "" {
		return Product{}, fmt.Errorf("id cannot be empty")
	}

	response := &Product{ID: id}
	if result := repo.db.WithContext(ctx).First(response); result.Error != nil {
		// TODO: add log
		return Product{}, result.Error
	}

	return *response, nil
}
