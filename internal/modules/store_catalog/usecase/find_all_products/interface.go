package findallproducts

import (
	"context"

	"github.com/MarcoBuarque/monolito/internal/modules/store_catalog/repository"
)

type IFindAllProductsUseCase interface {
	Execute(ctx context.Context) ([]repository.Product, error)
}
