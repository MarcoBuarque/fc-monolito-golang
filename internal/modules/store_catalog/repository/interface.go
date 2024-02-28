package repository

import (
	"context"
)

type IProductRepository interface {
	ListProducts(ctx context.Context) ([]Product, error)
	GetProduct(ctx context.Context, id string) (Product, error)
	UpdateSalesPrice(ctx context.Context, id string, price float64) error
}
