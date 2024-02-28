package facade

import (
	"context"

	"github.com/MarcoBuarque/monolito/internal/modules/store_catalog/repository"
)

type IStoreCatalogFacade interface {
	GetProduct(ctx context.Context, productID string) (repository.Product, error)
	ListProducts(ctx context.Context) ([]repository.Product, error)
	UpdateSalesPrice(ctx context.Context, id string, price float64) error
}
