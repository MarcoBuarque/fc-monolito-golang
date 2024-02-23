package facade

import (
	"context"

	"github.com/MarcoBuarque/monolito/internal/modules/store_catalog/repository"
	findallproducts "github.com/MarcoBuarque/monolito/internal/modules/store_catalog/usecase/find_all_products"
	findproduct "github.com/MarcoBuarque/monolito/internal/modules/store_catalog/usecase/find_product"
	updatesalesprice "github.com/MarcoBuarque/monolito/internal/modules/store_catalog/usecase/update_sales_price"
)

type StoreCatalogFacade struct {
	findAllUseCase     findallproducts.IFindAllProductsUseCase
	findUseCase        findproduct.IFindProductUseCase
	updatePriceUseCase updatesalesprice.IUpdateSalesPriceUseCase
}

func NewStoreCatalogFacade(findAllUseCase findallproducts.IFindAllProductsUseCase, findUseCase findproduct.IFindProductUseCase, updatePriceUseCase updatesalesprice.IUpdateSalesPriceUseCase) StoreCatalogFacade {
	return StoreCatalogFacade{findAllUseCase, findUseCase, updatePriceUseCase}
}

func (facade StoreCatalogFacade) FindAll(ctx context.Context) ([]repository.Product, error) {
	return facade.findAllUseCase.Execute(ctx)
}

func (facade StoreCatalogFacade) Find(ctx context.Context, productID string) (repository.Product, error) {
	return facade.findUseCase.Execute(ctx, productID)
}

func (facade StoreCatalogFacade) UpdateSalesPrice(ctx context.Context, productID string, price float64) error {
	return facade.updatePriceUseCase.Execute(ctx, productID, price)
}
