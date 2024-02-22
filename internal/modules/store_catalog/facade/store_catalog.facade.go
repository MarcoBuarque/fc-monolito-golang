package facade

import (
	"context"

	"github.com/MarcoBuarque/monolito/internal/modules/store_catalog/repository"
	findallproducts "github.com/MarcoBuarque/monolito/internal/modules/store_catalog/usecase/find_all_products"
	findproduct "github.com/MarcoBuarque/monolito/internal/modules/store_catalog/usecase/find_product"
)

type StoreCatalogFacade struct {
	findAllUseCase findallproducts.IFindAllProductsUseCase
	findUseCase    findproduct.IFindProductUseCase
}

func NewStoreCatalogFacade(findAllUseCase findallproducts.IFindAllProductsUseCase, findUseCase findproduct.IFindProductUseCase) StoreCatalogFacade {
	return StoreCatalogFacade{findAllUseCase, findUseCase}
}

func (facade StoreCatalogFacade) FindAll(ctx context.Context) ([]repository.ProductData, error) {
	return facade.findAllUseCase.Execute(ctx)
}

func (facade StoreCatalogFacade) Find(ctx context.Context, productID string) (repository.ProductData, error) {
	return facade.findUseCase.Execute(ctx, productID)
}
