package findproduct

import (
	"context"

	"github.com/MarcoBuarque/monolito/internal/modules/store_catalog/repository"
)

type IFindProductUseCase interface {
	Execute(ctx context.Context, productID string) (repository.ProductData, error)
}
