package domain

import (
	"github.com/MarcoBuarque/fc-monolito-golang/internal/modules/shared/domain/entity"
	valueobject "github.com/MarcoBuarque/fc-monolito-golang/internal/modules/shared/domain/value_object"
	"github.com/MarcoBuarque/fc-monolito-golang/internal/modules/shared/types"
	"github.com/shopspring/decimal"
)

type OrderEntity struct {
	status   types.Status
	client   ClientEntity
	products []ProductEntity
	entity.BaseEntity
}

func NewOrderEntity(id string, status types.Status, client ClientEntity, products []ProductEntity) OrderEntity {
	if status == "" {
		status = types.Pending
	}

	return OrderEntity{
		client:     client,
		products:   products,
		status:     status,
		BaseEntity: entity.NewBaseEntity(valueobject.NewID(id)),
	}
}

func (data OrderEntity) Status() types.Status      { return data.status }
func (data OrderEntity) Client() ClientEntity      { return data.client }
func (data OrderEntity) Products() []ProductEntity { return data.products }
func (data OrderEntity) Total() decimal.Decimal {
	sum := decimal.NewFromInt(0)

	for _, product := range data.products {
		price := product.SalesPrice()
		quantity := decimal.NewFromInt32(product.Quantity())
		sum = sum.Add(price.Mul(quantity))
	}

	return sum
}
