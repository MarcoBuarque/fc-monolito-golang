package repository

import (
	"context"
)

type ICheckoutRepository interface {
	CreateOrder(ctx context.Context, order Order) error
	GetOrder(ctx context.Context, orderID string) (Order, error)
}
