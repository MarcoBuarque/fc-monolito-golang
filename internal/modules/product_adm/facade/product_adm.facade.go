package facade

import (
	"context"

	"github.com/MarcoBuarque/monolito/internal/modules/product_adm/repository"
	addproduct "github.com/MarcoBuarque/monolito/internal/modules/product_adm/use_case/add_product"
)

type ProductAdmFacade struct {
	addUseCase addproduct.IAddProductUseCase
	// checkStockUseCase  checkStock.checkStockUseCase
}

func NewProductAdmFacade(addUseCase addproduct.IAddProductUseCase) ProductAdmFacade {
	return ProductAdmFacade{addUseCase: addUseCase}
}

func (facade ProductAdmFacade) AddProduct(ctx context.Context, data repository.ProductData) (repository.ProductData, error) {
	return facade.addUseCase.Execute(ctx, repository.ProductData{})
}

func (facade ProductAdmFacade) CheckStock(ctx context.Context, productID string) {
	facade.addUseCase.Execute(ctx, repository.ProductData{})
}
