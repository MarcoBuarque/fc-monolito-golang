package facade

import (
	"context"

	"github.com/MarcoBuarque/fc-monolito-golang/internal/modules/store_catalog/repository"
	getproduct "github.com/MarcoBuarque/fc-monolito-golang/internal/modules/store_catalog/usecase/get_product"
	listproducts "github.com/MarcoBuarque/fc-monolito-golang/internal/modules/store_catalog/usecase/list_products"
)

type StoreCatalogFacade struct {
	listProductsUseCase listproducts.IListProductsUseCase
	getProductUseCase   getproduct.IGetProductUseCase
}

func NewStoreCatalogFacade(findAllUseCase listproducts.IListProductsUseCase, findUseCase getproduct.IGetProductUseCase) StoreCatalogFacade {
	return StoreCatalogFacade{findAllUseCase, findUseCase}
}

func (facade StoreCatalogFacade) ListProducts(ctx context.Context) ([]repository.ProductCatalog, error) {
	return facade.listProductsUseCase.Execute(ctx)
}

func (facade StoreCatalogFacade) GetProduct(ctx context.Context, productID string) (repository.ProductCatalog, error) {
	return facade.getProductUseCase.Execute(ctx, productID)
}
