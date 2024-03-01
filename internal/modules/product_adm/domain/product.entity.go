package domain

import (
	"fmt"

	"github.com/MarcoBuarque/fc-monolito-golang/internal/modules/shared/domain/entity"
	valueobject "github.com/MarcoBuarque/fc-monolito-golang/internal/modules/shared/domain/value_object"
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
		BaseEntity:    entity.NewBaseEntity(valueobject.NewID(id)),
	}, nil
}

func (data ProductEntity) Name() string                   { return data.name }
func (data ProductEntity) Description() string            { return data.description }
func (data ProductEntity) PurchasePrice() decimal.Decimal { return data.purchasePrice }
func (data ProductEntity) Stock() int32                   { return data.stock }
func (data *ProductEntity) SetStock(stock int32)          { data.stock = stock }
