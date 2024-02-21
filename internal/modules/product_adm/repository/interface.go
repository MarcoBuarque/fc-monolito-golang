package repository

import "context"

type IProductRepository interface {
	Add(ctx context.Context, data ProductData) error
	Find(ctx context.Context, id string) (ProductData, error)
}
