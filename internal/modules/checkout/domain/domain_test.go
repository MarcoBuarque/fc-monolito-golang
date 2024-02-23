package domain

import (
	"fmt"
	"testing"

	valueobject "github.com/MarcoBuarque/monolito/internal/modules/shared/domain/value_object"
	"github.com/MarcoBuarque/monolito/internal/modules/shared/types"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	assert := assert.New(t)

	type args struct {
		id      string
		name    string
		email   string
		address valueobject.Address
	}

	type expect struct {
		data ClientEntity
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
				data: ClientEntity{}, err: fmt.Errorf("clientEntity: name cannot be empty"),
			},
		},
		{
			title: "Should return an error when purchase email is empty",
			args: args{
				name: "xpto",
			},
			expect: expect{
				data: ClientEntity{}, err: fmt.Errorf("clientEntity: email cannot be empty"),
			},
		},

		{
			title: "Success",
			args: args{
				id:    "xpto_id",
				name:  "xpto",
				email: "email_test",
			},
			expect: expect{
				data: ClientEntity{
					name:    "xpto",
					address: valueobject.Address{},
					email:   "email_test",
				}, err: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			response, err := NewClient(tt.args.id, tt.args.name, tt.args.email, tt.args.address)

			assert.Equal(tt.expect.err, err)
			assert.Equal(tt.expect.data.Name(), response.Name())
			assert.Equal(tt.expect.data.Email(), response.Email())
		})
	}
}

func TestNewProduct(t *testing.T) {
	assert := assert.New(t)

	type args struct {
		id          string
		name        string
		description string
		salesPrice  decimal.Decimal
		quantity    int32
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
			title: "Should return an error when sales price is less than or equal zero",
			args: args{
				name:        "xpto",
				description: "xpto_description",
			},
			expect: expect{
				data: ProductEntity{}, err: fmt.Errorf("productEntity: sales price cannot be less than or equal zero"),
			},
		},
		{
			title: "Should return an error quantity is less than or equal zero",
			args: args{
				id:          "xpto_id",
				name:        "xpto",
				description: "xpto_description",
				salesPrice:  decimal.NewFromInt(10),
				quantity:    0,
			},
			expect: expect{
				data: ProductEntity{}, err: fmt.Errorf("productEntity: quantity cannot be less than or equal zero"),
			},
		},
		{
			title: "Success",
			args: args{
				id:          "xpto_id",
				name:        "xpto",
				description: "xpto_description",
				salesPrice:  decimal.NewFromInt(10),
				quantity:    10,
			},
			expect: expect{
				data: ProductEntity{
					name:        "xpto",
					description: "xpto_description",
					salesPrice:  decimal.NewFromInt(10),
				}, err: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			response, err := NewProduct(tt.args.id, tt.args.name, tt.args.description, tt.args.salesPrice, tt.args.quantity)

			assert.Equal(tt.expect.err, err)
			assert.Equal(tt.expect.data.Name(), response.Name())
			assert.Equal(tt.expect.data.Description(), response.Description())
			assert.Equal(tt.expect.data.SalesPrice(), response.SalesPrice())
		})
	}
}

func TestNewOrder(t *testing.T) {
	assert := assert.New(t)

	client := ClientEntity{}
	products := []ProductEntity{}

	type args struct {
		id       string
		status   types.Status
		client   ClientEntity
		products []ProductEntity
	}

	tests := []struct {
		title  string
		args   args
		expect OrderEntity
	}{
		{
			title: "Success without status",
			args: args{
				id: "xpto_id",
				// status:   "xpto",
				client:   client,
				products: products,
			},
			expect: OrderEntity{
				client:   client,
				products: products,
				status:   types.Pending,
			},
		},
		{
			title: "Success without status",
			args: args{
				id:       "xpto_id",
				status:   "xpto",
				client:   client,
				products: products,
			},
			expect: OrderEntity{
				client:   client,
				products: products,
				status:   "xpto",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			response := NewOrderEntity(tt.args.id, tt.args.status, tt.args.client, tt.args.products)

			assert.Equal(tt.expect.Status(), response.Status())
			assert.Equal(tt.expect.Client(), response.Client())
			assert.Equal(tt.expect.Products(), response.Products())
		})
	}
}

func TestTotal(t *testing.T) {
	order := OrderEntity{
		products: []ProductEntity{
			{
				quantity:   1,
				salesPrice: decimal.NewFromInt(10),
			},
			{
				quantity:   3,
				salesPrice: decimal.NewFromInt(5),
			},
			{
				quantity:   1,
				salesPrice: decimal.NewFromInt(7),
			},
		},
	}

	assert.Equal(t, "32", order.Total().String())
}
