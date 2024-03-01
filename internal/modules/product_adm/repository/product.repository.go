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

func (repo ProductRepository) CreateProduct(ctx context.Context, data Product) error {
	if result := repo.db.WithContext(ctx).Create(&data); result.Error != nil {
		// TODO: add log
		return fmt.Errorf("ProductRepository - CreateProduct: %w", result.Error)
	}

	return nil
}

func (repo ProductRepository) GetProduct(ctx context.Context, id string) (Product, error) {
	if id == "" {
		return Product{}, fmt.Errorf("ProductRepository - GetProduct: id cannot be empty")
	}

	response := &Product{ID: id}
	if result := repo.db.WithContext(ctx).First(response); result.Error != nil {
		// TODO: add log
		return Product{}, fmt.Errorf("ProductRepository - GetProduct: %w", result.Error)
	}

	return *response, nil
}

func (repo ProductRepository) UpdateSalesPrice(ctx context.Context, id string, price float32) error {
	if id == "" {
		return fmt.Errorf("id cannot be empty")
	}

	if result := repo.db.WithContext(ctx).Where("id = ?", id).Update("sales_price", price); result.Error != nil {
		// TODO: add log
		return result.Error
	}

	return nil
}
