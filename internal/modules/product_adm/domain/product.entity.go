package domain

import (
	"fmt"

	"github.com/MarcoBuarque/monolito/internal/modules/product_adm/repository"
	"github.com/MarcoBuarque/monolito/internal/modules/shared/domain/entity"
	valueobject "github.com/MarcoBuarque/monolito/internal/modules/shared/domain/value_object"
	"github.com/shopspring/decimal"
)

type ProductEntity struct {
	name          string
	description   string
	purchasePrice decimal.Decimal
	stock         int32
	entity.BaseEntity
}

func NewProduct(id, name, description string, purchasePrice decimal.Decimal, stock int32) (ProductEntity, error) {
	if name == "" {
		return ProductEntity{}, fmt.Errorf("productEntity: name cannot be empty")
	}

	if description == "" {
		return ProductEntity{}, fmt.Errorf("productEntity: description cannot be empty")
	}

	if purchasePrice.LessThanOrEqual(decimal.NewFromInt(0)) {
		return ProductEntity{}, fmt.Errorf("productEntity: purchase price cannot be less than or equal zero")
	}

	if stock <= 0 {
		return ProductEntity{}, fmt.Errorf("productEntity: stock cannot be less than or equal zero")
	}

	return ProductEntity{
		name:          name,
		description:   description,
		purchasePrice: purchasePrice,
		stock:         stock,
		BaseEntity:    entity.CreateBaseEntity(valueobject.CreateID(id)),
	}, nil
}

func (data ProductEntity) Name() string                   { return data.name }
func (data ProductEntity) Description() string            { return data.description }
func (data ProductEntity) PurchasePrice() decimal.Decimal { return data.purchasePrice }
func (data ProductEntity) Stock() int32                   { return data.stock }
func (data *ProductEntity) SetStock(stock int32)          { data.stock = stock }
func (data ProductEntity) ToData() repository.ProductData {
	return repository.ProductData{
		ID:            string(data.ID().ToString()),
		Name:          data.name,
		Description:   data.description,
		PurchasePrice: data.purchasePrice,
		Stock:         data.stock,
		CreatedAt:     data.CreatedAt(),
		UpdatedAt:     data.UpdatedAt(),
	}
}
