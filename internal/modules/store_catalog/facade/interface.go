package facade

import (
	"context"

	"github.com/MarcoBuarque/monolito/internal/modules/store_catalog/repository"
)

type IStoreCatalogFacade interface {
	Find(ctx context.Context, productID string) (repository.ProductData, error)
	FindAll(ctx context.Context) ([]repository.ProductData, error)
}
