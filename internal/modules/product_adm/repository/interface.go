package repository

import "context"

type IProductRepository interface {
	CreateProduct(ctx context.Context, data Product) error
	GetProduct(ctx context.Context, id string) (Product, error)
}
