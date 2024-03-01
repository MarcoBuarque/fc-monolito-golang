package repository

import "context"

type IProductRepository interface {
	CreateProduct(ctx context.Context, data Product) error
	GetProduct(ctx context.Context, productID string) (Product, error)
	UpdateSalesPrice(ctx context.Context, productID string, newprice float32) error
}
