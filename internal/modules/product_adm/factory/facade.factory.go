package factory

import (
	"sync"

	"github.com/MarcoBuarque/monolito/config"
	"github.com/MarcoBuarque/monolito/internal/modules/product_adm/facade"
	"github.com/MarcoBuarque/monolito/internal/modules/product_adm/repository"
	checkstock "github.com/MarcoBuarque/monolito/internal/modules/product_adm/usecase/check_stock"
	addproduct "github.com/MarcoBuarque/monolito/internal/modules/product_adm/usecase/create_product"
)

var (
	once      sync.Once
	singleton facade.ProductAdmFacade
)

func createSingleton() {
	repo := repository.NewProductRepository(config.GetDB())

	singleton = facade.NewProductAdmFacade(
		addproduct.NewCreateProductUseCase(repo),
		checkstock.NewCheckStockUseCase(repo),
	)
}
func NewProductAdmFacadeFactory() facade.IProductAdmFacade {
	once.Do(createSingleton)

	return singleton
}
