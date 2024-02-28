package repository

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type CheckoutRepository struct {
	db *gorm.DB
}

func NewCheckoutRepository(db *gorm.DB) CheckoutRepository {
	return CheckoutRepository{db}
}

func (repo CheckoutRepository) AddOrder(ctx context.Context, data Order) error {
	if result := repo.db.WithContext(ctx).Create(&data); result.Error != nil {
		// TODO: add log
		return result.Error
	}

	return nil
}

func (repo CheckoutRepository) FindOrder(ctx context.Context, orderID string) (Order, error) {
	if orderID == "" {
		return Order{}, fmt.Errorf("id cannot be zero")
	}

	response := &Order{ID: orderID}
	if result := repo.db.WithContext(ctx).Preload("clients").Preload("order_products").First(response); result.Error != nil {
		// TODO: add log
		return Order{}, result.Error
	}

	return *response, nil
}
