package repository

import "context"

type IProductRepository interface {
	FindAll(ctx context.Context) ([]ProductData, error)
	Find(ctx context.Context, id string) (ProductData, error)
}
