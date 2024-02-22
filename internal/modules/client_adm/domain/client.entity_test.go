package domain

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	assert := assert.New(t)

	type args struct {
		id          string
		name        string
		description string
		email       string
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
			title: "Should return an error when description is empty",
			args: args{
				name: "xpto",
			},
			expect: expect{
				data: ClientEntity{}, err: fmt.Errorf("clientEntity: description cannot be empty"),
			},
		},
		{
			title: "Should return an error when purchase email is empty",
			args: args{
				name:        "xpto",
				description: "xpto_description",
			},
			expect: expect{
				data: ClientEntity{}, err: fmt.Errorf("clientEntity: email cannot be empty"),
			},
		},

		{
			title: "Success",
			args: args{
				id:          "xpto_id",
				name:        "xpto",
				description: "xpto_description",
				email:       "email_test",
			},
			expect: expect{
				data: ClientEntity{
					name:    "xpto",
					address: "xpto_description",
					email:   "email_test",
				}, err: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			response, err := NewClient(tt.args.id, tt.args.name, tt.args.description, tt.args.email)

			assert.Equal(tt.expect.err, err)
			assert.Equal(tt.expect.data.Name(), response.Name())
			assert.Equal(tt.expect.data.Email(), response.Email())
		})
	}
}

func TestToData(t *testing.T) {
	assert := assert.New(t)

	entity := ClientEntity{
		name:    "xpto",
		address: "xpto_description",
		email:   "email",
	}

	dbData := entity.ToData()

	assert.Equal(entity.ID().ToString(), dbData.ID)
	assert.Equal(entity.Name(), dbData.Name)
	assert.Equal(entity.Email(), dbData.Email)
	assert.Equal(entity.CreatedAt(), dbData.CreatedAt)
	assert.Equal(entity.UpdatedAt(), dbData.UpdatedAt)
}
