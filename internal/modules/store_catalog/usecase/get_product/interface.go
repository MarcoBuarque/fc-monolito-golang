package getproduct

import (
	"context"

	"github.com/MarcoBuarque/fc-monolito-golang/internal/modules/store_catalog/repository"
)

type IGetProductUseCase interface {
	Execute(ctx context.Context, productID string) (repository.ProductCatalog, error)
}
