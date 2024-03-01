package factory

import (
	"sync"

	"github.com/MarcoBuarque/fc-monolito-golang/config"
	"github.com/MarcoBuarque/fc-monolito-golang/internal/modules/product_adm/facade"
	"github.com/MarcoBuarque/fc-monolito-golang/internal/modules/product_adm/repository"
	checkstock "github.com/MarcoBuarque/fc-monolito-golang/internal/modules/product_adm/usecase/check_stock"
	addproduct "github.com/MarcoBuarque/fc-monolito-golang/internal/modules/product_adm/usecase/create_product"
	updatesalesprice "github.com/MarcoBuarque/fc-monolito-golang/internal/modules/product_adm/usecase/update_sales_price"
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
		updatesalesprice.NewUpdateSalesPriceUseCase(repo),
	)
}
func NewProductAdmFacadeFactory() facade.IProductAdmFacade {
	once.Do(createSingleton)

	return singleton
}
