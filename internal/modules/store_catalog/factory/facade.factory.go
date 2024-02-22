package factory

import (
	"github.com/MarcoBuarque/monolito/internal/modules/store_catalog/facade"
	"github.com/MarcoBuarque/monolito/internal/modules/store_catalog/repository"
	findallproducts "github.com/MarcoBuarque/monolito/internal/modules/store_catalog/usecase/find_all_products"
	findproduct "github.com/MarcoBuarque/monolito/internal/modules/store_catalog/usecase/find_product"
	"github.com/MarcoBuarque/monolito/pkg/database"
)

type StoreCatalogFacadeFactory struct {
}

func NewStoreCatalogFacadeFactory() facade.IStoreCatalogFacade {
	repo := repository.NewProductRepository(database.GetDB())

	return facade.NewStoreCatalogFacade(findallproducts.NewFindAllProductsUseCase(repo), findproduct.NewFindProductUseCase(repo))
}
