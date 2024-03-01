package repository

import (
	"context"
)

type IProductRepository interface {
	ListProducts(ctx context.Context) ([]ProductCatalog, error)
	GetProduct(ctx context.Context, id string) (ProductCatalog, error)
}
