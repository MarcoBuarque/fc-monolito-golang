package factory

import (
	"github.com/MarcoBuarque/monolito/internal/modules/product_adm/facade"
	"github.com/MarcoBuarque/monolito/internal/modules/product_adm/repository"
	addproduct "github.com/MarcoBuarque/monolito/internal/modules/product_adm/usecase/add_product"
	checkstock "github.com/MarcoBuarque/monolito/internal/modules/product_adm/usecase/check_stock"
	"github.com/MarcoBuarque/monolito/pkg/database"
)

type ProductAdmFacadeFactory struct {
}

func NewProductAdmFacadeFactory() facade.ProductAdmFacade {
	repo := repository.NewProductRepository(database.GetDB())
	addUseCase := addproduct.NewAddProductUseCase(repo)
	checkStockUseCase := checkstock.NewCheckStockUseCase(repo)
	return facade.NewProductAdmFacade(addUseCase, checkStockUseCase)
}
