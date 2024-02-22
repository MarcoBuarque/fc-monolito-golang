package facade

import (
	"context"

	"github.com/MarcoBuarque/monolito/internal/modules/product_adm/repository"
	addproduct "github.com/MarcoBuarque/monolito/internal/modules/product_adm/usecase/add_product"
	checkstock "github.com/MarcoBuarque/monolito/internal/modules/product_adm/usecase/check_stock"
)

type ProductAdmFacade struct {
	addUseCase        addproduct.IAddProductUseCase
	checkStockUseCase checkstock.ICheckStockUseCase
}

func NewProductAdmFacade(addUseCase addproduct.IAddProductUseCase, checkStockUseCase checkstock.ICheckStockUseCase) ProductAdmFacade {
	return ProductAdmFacade{addUseCase, checkStockUseCase}
}

func (facade ProductAdmFacade) AddProduct(ctx context.Context, data repository.ProductData) (repository.ProductData, error) {
	return facade.addUseCase.Execute(ctx, data)
}

func (facade ProductAdmFacade) CheckStock(ctx context.Context, id string) (int32, error) {
	return facade.checkStockUseCase.Execute(ctx, id)
}
