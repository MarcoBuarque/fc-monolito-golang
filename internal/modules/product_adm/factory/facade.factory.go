package factory

import (
	"github.com/MarcoBuarque/monolito/config"
	"github.com/MarcoBuarque/monolito/internal/modules/product_adm/facade"
	"github.com/MarcoBuarque/monolito/internal/modules/product_adm/repository"
	addproduct "github.com/MarcoBuarque/monolito/internal/modules/product_adm/usecase/add_product"
	checkstock "github.com/MarcoBuarque/monolito/internal/modules/product_adm/usecase/check_stock"
)

type ProductAdmFacadeFactory struct {
}

func NewProductAdmFacadeFactory() facade.ProductAdmFacade {
	repo := repository.NewProductRepository(config.GetDB())

	return facade.NewProductAdmFacade(
		addproduct.NewAddProductUseCase(repo),
		checkstock.NewCheckStockUseCase(repo),
	)
}
