package getproduct

import (
	"context"

	"github.com/MarcoBuarque/monolito/internal/modules/store_catalog/repository"
)

type IGetProductUseCase interface {
	Execute(ctx context.Context, productID string) (repository.Product, error)
}
