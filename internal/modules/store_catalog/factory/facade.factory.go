package factory

import (
	"sync"

	"github.com/MarcoBuarque/fc-monolito-golang/config"
	"github.com/MarcoBuarque/fc-monolito-golang/internal/modules/store_catalog/facade"
	"github.com/MarcoBuarque/fc-monolito-golang/internal/modules/store_catalog/repository"
	getproduct "github.com/MarcoBuarque/fc-monolito-golang/internal/modules/store_catalog/usecase/get_product"
	listproducts "github.com/MarcoBuarque/fc-monolito-golang/internal/modules/store_catalog/usecase/list_products"
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
	)
}

func NewStoreCatalogFacadeFactory() facade.IStoreCatalogFacade {
	once.Do(createSingleton)

	return singleton
}
