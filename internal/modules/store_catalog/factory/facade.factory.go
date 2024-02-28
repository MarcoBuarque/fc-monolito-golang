package factory

import (
	"github.com/MarcoBuarque/monolito/config"
	"github.com/MarcoBuarque/monolito/internal/modules/store_catalog/facade"
	"github.com/MarcoBuarque/monolito/internal/modules/store_catalog/repository"
	findallproducts "github.com/MarcoBuarque/monolito/internal/modules/store_catalog/usecase/find_all_products"
	findproduct "github.com/MarcoBuarque/monolito/internal/modules/store_catalog/usecase/find_product"
	updatesalesprice "github.com/MarcoBuarque/monolito/internal/modules/store_catalog/usecase/update_sales_price"
)

type StoreCatalogFacadeFactory struct {
}

func NewStoreCatalogFacadeFactory() facade.IStoreCatalogFacade {
	repo := repository.NewProductRepository(config.GetDB())

	return facade.NewStoreCatalogFacade(
		findallproducts.NewFindAllProductsUseCase(repo),
		findproduct.NewFindProductUseCase(repo),
		updatesalesprice.NewUpdateSalesPriceUseCase(repo),
	)
}
