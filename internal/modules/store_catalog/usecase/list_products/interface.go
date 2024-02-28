package listproducts

import (
	"context"

	"github.com/MarcoBuarque/monolito/internal/modules/store_catalog/repository"
)

type IListProductsUseCase interface {
	Execute(ctx context.Context) ([]repository.Product, error)
}
