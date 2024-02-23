package facade

import (
	"context"

	"github.com/MarcoBuarque/monolito/internal/modules/store_catalog/repository"
)

type IStoreCatalogFacade interface {
	Find(ctx context.Context, productID string) (repository.Product, error)
	FindAll(ctx context.Context) ([]repository.Product, error)
	UpdateSalesPrice(ctx context.Context, id string, price float64) error
}
