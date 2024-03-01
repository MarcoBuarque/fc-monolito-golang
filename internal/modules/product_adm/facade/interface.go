package facade

import (
	"context"

	"github.com/MarcoBuarque/fc-monolito-golang/internal/modules/product_adm/repository"
)

type IProductAdmFacade interface {
	CreateProduct(ctx context.Context, data repository.Product) (repository.Product, error)
	CheckStock(ctx context.Context, productID string) (int32, error)
	UpdateSalesPrice(ctx context.Context, productID string, newprice float32) (repository.Product, error)
}
