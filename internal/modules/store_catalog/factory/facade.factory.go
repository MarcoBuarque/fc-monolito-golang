package factory

import (
	"github.com/MarcoBuarque/monolito/config"
	"github.com/MarcoBuarque/monolito/internal/modules/store_catalog/facade"
	"github.com/MarcoBuarque/monolito/internal/modules/store_catalog/repository"
	getproduct "github.com/MarcoBuarque/monolito/internal/modules/store_catalog/usecase/get_product"
	listproducts "github.com/MarcoBuarque/monolito/internal/modules/store_catalog/usecase/list_products"
	updatesalesprice "github.com/MarcoBuarque/monolito/internal/modules/store_catalog/usecase/update_sales_price"
)

type StoreCatalogFacadeFactory struct {
}

func NewStoreCatalogFacadeFactory() facade.IStoreCatalogFacade {
	repo := repository.NewProductRepository(config.GetDB())

	return facade.NewStoreCatalogFacade(
		listproducts.NewListProductsUseCase(repo),
		getproduct.NewGetProductUseCase(repo),
		updatesalesprice.NewUpdateSalesPriceUseCase(repo),
	)
}
