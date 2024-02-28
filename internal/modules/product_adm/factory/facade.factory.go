package factory

import (
	"sync"

	"github.com/MarcoBuarque/monolito/config"
	"github.com/MarcoBuarque/monolito/internal/modules/product_adm/facade"
	"github.com/MarcoBuarque/monolito/internal/modules/product_adm/repository"
	addproduct "github.com/MarcoBuarque/monolito/internal/modules/product_adm/usecase/add_product"
	checkstock "github.com/MarcoBuarque/monolito/internal/modules/product_adm/usecase/check_stock"
)

var (
	once      sync.Once
	singleton facade.ProductAdmFacade
)

func createSingleton() {
	repo := repository.NewProductRepository(config.GetDB())

	singleton = facade.NewProductAdmFacade(
		addproduct.NewAddProductUseCase(repo),
		checkstock.NewCheckStockUseCase(repo),
	)
}
func NewProductAdmFacadeFactory() facade.ProductAdmFacade {
	once.Do(createSingleton)

	return singleton
}
