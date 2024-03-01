package domain

import (
	"fmt"

	"github.com/MarcoBuarque/fc-monolito-golang/internal/modules/shared/domain/entity"
	valueobject "github.com/MarcoBuarque/fc-monolito-golang/internal/modules/shared/domain/value_object"
	"github.com/shopspring/decimal"
)

type ProductEntity struct {
	name        string
	description string
	salesPrice  decimal.Decimal
	quantity    int32
	entity.BaseEntity
}

func NewProduct(id, name, description string, salesPrice decimal.Decimal, quantity int32) (ProductEntity, error) {
	if name == "" {
		return ProductEntity{}, fmt.Errorf("productEntity: name cannot be empty")
	}

	if description == "" {
		return ProductEntity{}, fmt.Errorf("productEntity: description cannot be empty")
	}

	if salesPrice.LessThanOrEqual(decimal.NewFromInt(0)) {
		return ProductEntity{}, fmt.Errorf("productEntity: sales price cannot be less than or equal zero")
	}

	if quantity <= 0 {
		return ProductEntity{}, fmt.Errorf("productEntity: quantity cannot be less than or equal zero")
	}

	return ProductEntity{
		name:        name,
		description: description,
		salesPrice:  salesPrice,
		quantity:    quantity,
		BaseEntity:  entity.NewBaseEntity(valueobject.NewID(id)),
	}, nil
}

func (data ProductEntity) Name() string                { return data.name }
func (data ProductEntity) Description() string         { return data.description }
func (data ProductEntity) SalesPrice() decimal.Decimal { return data.salesPrice }
func (data ProductEntity) Quantity() int32             { return data.quantity }
