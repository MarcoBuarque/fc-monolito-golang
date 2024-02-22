package repository

import "context"

type IProductRepository interface {
	Add(ctx context.Context, data Product) error
	Find(ctx context.Context, id string) (Product, error)
}
