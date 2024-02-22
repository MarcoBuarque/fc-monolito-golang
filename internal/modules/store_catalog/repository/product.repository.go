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

func (repo ProductRepository) FindAll(ctx context.Context) ([]Product, error) {
	var response []Product

	if result := repo.db.WithContext(ctx).Find(&response); result.Error != nil {
		// TODO: add log
		return []Product{}, result.Error
	}

	return response, nil
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
