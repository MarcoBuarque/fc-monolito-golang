package domain

import (
	"fmt"
	"testing"

	valueobject "github.com/MarcoBuarque/monolito/internal/modules/shared/domain/value_object"
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
