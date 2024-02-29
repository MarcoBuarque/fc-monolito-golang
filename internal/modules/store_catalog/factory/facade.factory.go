package factory

import (
	"sync"

	"github.com/MarcoBuarque/monolito/config"
	"github.com/MarcoBuarque/monolito/internal/modules/store_catalog/facade"
	"github.com/MarcoBuarque/monolito/internal/modules/store_catalog/repository"
	getproduct "github.com/MarcoBuarque/monolito/internal/modules/store_catalog/usecase/get_product"
	listproducts "github.com/MarcoBuarque/monolito/internal/modules/store_catalog/usecase/list_products"
	updatesalesprice "github.com/MarcoBuarque/monolito/internal/modules/store_catalog/usecase/update_sales_price"
)

var (
	once      sync.Once
	singleton facade.IStoreCatalogFacade
)

func createSingleton() {
	repo := repository.NewProductRepository(config.GetDB())

	singleton = facade.NewStoreCatalogFacade(
		listproducts.NewListProductsUseCase(repo),
		getproduct.NewGetProductUseCase(repo),
		updatesalesprice.NewUpdateSalesPriceUseCase(repo),
	)
}

func NewStoreCatalogFacadeFactory() facade.IStoreCatalogFacade {
	once.Do(createSingleton)

	return singleton
}
