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

func (repo ProductRepository) ListProducts(ctx context.Context) ([]ProductCatalog, error) {
	var response []ProductCatalog

	if result := repo.db.WithContext(ctx).Find(&response); result.Error != nil {
		// TODO: add log
		return []ProductCatalog{}, result.Error
	}

	return response, nil
}

func (repo ProductRepository) GetProduct(ctx context.Context, id string) (ProductCatalog, error) {
	if id == "" {
		return ProductCatalog{}, fmt.Errorf("id cannot be empty")
	}

	response := &ProductCatalog{ID: id}
	if result := repo.db.WithContext(ctx).First(response); result.Error != nil {
		// TODO: add log
		return ProductCatalog{}, result.Error
	}

	return *response, nil
}
