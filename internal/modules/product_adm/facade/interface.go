package facade

import (
	"context"

	"github.com/MarcoBuarque/monolito/internal/modules/product_adm/repository"
)

type IProductAdmFacade interface {
	AddProduct(ctx context.Context, data repository.Product) (repository.Product, error)
	CheckStock(ctx context.Context, productID string) (int32, error)
}
