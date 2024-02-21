package facade

import (
	"context"
)

type IProductAdmFacade interface {
	AddProduct(ctx context.Context, productID string)
	CheckStock(ctx context.Context, productID string) (id string, stock int32)
}
