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

func (repo CheckoutRepository) CreateOrder(ctx context.Context, data Order) error {
	order, products := data.convertToCreate()

	tx := repo.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	if result := tx.Create(order); result.Error != nil {
		// TODO: add log
		return result.Error
	}

	if result := tx.Create(products); result.Error != nil {
		// TODO: add log
		return result.Error
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

func (repo CheckoutRepository) GetOrder(ctx context.Context, orderID string) (Order, error) {
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
