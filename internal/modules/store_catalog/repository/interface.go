package repository

import "context"

type IProductRepository interface {
	FindAll(ctx context.Context) ([]Product, error)
	Find(ctx context.Context, id string) (Product, error)
}
