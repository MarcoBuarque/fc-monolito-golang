package factory

import (
	"github.com/MarcoBuarque/monolito/internal/modules/product_adm/facade"
	"github.com/MarcoBuarque/monolito/internal/modules/product_adm/repository"
	addproduct "github.com/MarcoBuarque/monolito/internal/modules/product_adm/use_case/add_product"
	"github.com/MarcoBuarque/monolito/pkg/database"
)

type ProductAdmFacadeFactory struct {
}

func NewProductAdmFacadeFactory() facade.ProductAdmFacade {
	repo := repository.NewProductRepository(database.GetDB())
	addUseCase := addproduct.NewAddProductUseCase(repo)
	return facade.NewProductAdmFacade(addUseCase)
}
