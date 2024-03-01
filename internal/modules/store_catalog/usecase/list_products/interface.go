package listproducts

import (
	"context"

	"github.com/MarcoBuarque/fc-monolito-golang/internal/modules/store_catalog/repository"
)

type IListProductsUseCase interface {
	Execute(ctx context.Context) ([]repository.ProductCatalog, error)
}
