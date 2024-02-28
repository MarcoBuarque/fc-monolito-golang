package repository

import (
	"context"
)

type ICheckoutRepository interface {
	AddOrder(ctx context.Context, order Order) error
	FindOrder(ctx context.Context, orderID string) (Order, error)
}
