package facade

import (
	"context"

	"github.com/MarcoBuarque/monolito/internal/modules/store_catalog/repository"
	getproduct "github.com/MarcoBuarque/monolito/internal/modules/store_catalog/usecase/get_product"
	listproducts "github.com/MarcoBuarque/monolito/internal/modules/store_catalog/usecase/list_products"
	updatesalesprice "github.com/MarcoBuarque/monolito/internal/modules/store_catalog/usecase/update_sales_price"
)

type StoreCatalogFacade struct {
	findAllUseCase     listproducts.IListProductsUseCase
	findUseCase        getproduct.IGetProductUseCase
	updatePriceUseCase updatesalesprice.IUpdateSalesPriceUseCase
}

func NewStoreCatalogFacade(findAllUseCase listproducts.IListProductsUseCase, findUseCase getproduct.IGetProductUseCase, updatePriceUseCase updatesalesprice.IUpdateSalesPriceUseCase) StoreCatalogFacade {
	return StoreCatalogFacade{findAllUseCase, findUseCase, updatePriceUseCase}
}

func (facade StoreCatalogFacade) ListProducts(ctx context.Context) ([]repository.Product, error) {
	return facade.findAllUseCase.Execute(ctx)
}

func (facade StoreCatalogFacade) GetProduct(ctx context.Context, productID string) (repository.Product, error) {
	return facade.findUseCase.Execute(ctx, productID)
}

func (facade StoreCatalogFacade) UpdateSalesPrice(ctx context.Context, productID string, price float64) error {
	return facade.updatePriceUseCase.Execute(ctx, productID, price)
}
