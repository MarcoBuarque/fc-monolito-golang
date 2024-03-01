package facade

import (
	"context"

	"github.com/MarcoBuarque/monolito/internal/modules/store_catalog/repository"
)

type IStoreCatalogFacade interface {
	GetProduct(ctx context.Context, productID string) (repository.ProductCatalog, error)
	ListProducts(ctx context.Context) ([]repository.ProductCatalog, error)
}
