package domain

import (
	"fmt"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	assert := assert.New(t)

	type args struct {
		id            string
		name          string
		description   string
		purchasePrice decimal.Decimal
		stock         int32
	}

	type expect struct {
		data ProductEntity
		err  error
	}
	tests := []struct {
		title  string
		args   args
		expect expect
	}{
		{
			title: "Should return an error when name is empty",
			args:  args{},
			expect: expect{
				data: ProductEntity{}, err: fmt.Errorf("productEntity: name cannot be empty"),
			},
		},
		{
			title: "Should return an error when description is empty",
			args: args{
				name: "xpto",
			},
			expect: expect{
				data: ProductEntity{}, err: fmt.Errorf("productEntity: description cannot be empty"),
			},
		},
		{
			title: "Should return an error when purchase price is less than or equal zero",
			args: args{
				name:        "xpto",
				description: "xpto_description",
			},
			expect: expect{
				data: ProductEntity{}, err: fmt.Errorf("productEntity: purchase price cannot be less than or equal zero"),
			},
		},
		{
			title: "Should return an error when stock is less than or equal zero",
			args: args{
				name:          "xpto",
				description:   "xpto_description",
				purchasePrice: decimal.NewFromInt(10),
			},
			expect: expect{
				data: ProductEntity{}, err: fmt.Errorf("productEntity: stock cannot be less than or equal zero"),
			},
		},
		{
			title: "Success",
			args: args{
				id:            "xpto_id",
				name:          "xpto",
				description:   "xpto_description",
				purchasePrice: decimal.NewFromInt(10),
				stock:         10,
			},
			expect: expect{
				data: ProductEntity{
					name:          "xpto",
					description:   "xpto_description",
					purchasePrice: decimal.NewFromInt(10),
					stock:         10,
				}, err: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			response, err := NewProduct(tt.args.id, tt.args.name, tt.args.description, tt.args.purchasePrice, tt.args.stock)

			assert.Equal(tt.expect.err, err)
			assert.Equal(tt.expect.data.Name(), response.Name())
			assert.Equal(tt.expect.data.PurchasePrice(), response.PurchasePrice())
			assert.Equal(tt.expect.data.Stock(), response.Stock())
		})
	}
}

func TestToData(t *testing.T) {
	assert := assert.New(t)

	entity := ProductEntity{
		name:          "xpto",
		description:   "xpto_description",
		purchasePrice: decimal.NewFromInt(10),
		stock:         10,
	}

	dbData := entity.ToData()

	assert.Equal(entity.ID().ToString(), dbData.ID)
	assert.Equal(entity.Name(), dbData.Name)
	assert.Equal(entity.PurchasePrice(), dbData.PurchasePrice)
	assert.Equal(entity.Stock(), dbData.Stock)
	assert.Equal(entity.CreatedAt(), dbData.CreatedAt)
	assert.Equal(entity.UpdatedAt(), dbData.UpdatedAt)
}
