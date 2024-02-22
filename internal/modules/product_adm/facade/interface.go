package facade

import (
	"context"

	"github.com/MarcoBuarque/monolito/internal/modules/product_adm/repository"
)

type IProductAdmFacade interface {
	AddProduct(ctx context.Context, data repository.ProductData) (repository.ProductData, error)
	CheckStock(ctx context.Context, productID string) (int32, error)
}
