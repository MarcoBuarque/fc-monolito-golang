package valueobject

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewID(t *testing.T) {
	require := require.New(t)
	assert := assert.New(t)

	id := NewID("")

	require.NotEmpty(id)

	id = NewID("xpto-10")
	assert.Equal("xpto-10", id.id)

	assert.Equal("xpto-10", id.ToString())
}

func TestNewClient(t *testing.T) {
	assert := assert.New(t)

	type args struct{ street, number, complement, city, state, zipCode string }

	type expect struct {
		data Address
		err  error
	}
	tests := []struct {
		title  string
		args   args
		expect expect
	}{
		{
			title: "Should return an error when street is empty",
			args:  args{},
			expect: expect{
				data: Address{}, err: fmt.Errorf("address: street cannot be empty"),
			},
		},
		{
			title: "Should return an error when number is empty",
			args: args{
				street: "test_street",
			},
			expect: expect{
				data: Address{}, err: fmt.Errorf("address: number cannot be empty"),
			},
		},
		{
			title: "Should return an error when complement is empty",
			args: args{
				street: "test_street",
				number: "test_number",
			},
			expect: expect{
				data: Address{}, err: fmt.Errorf("address: complement cannot be empty"),
			},
		},
		{
			title: "Should return an error when city is empty",
			args: args{
				street:     "test_street",
				number:     "test_number",
				complement: "complement",
			},
			expect: expect{
				data: Address{}, err: fmt.Errorf("address: city cannot be empty"),
			},
		},
		{
			title: "Should return an error when state is empty",
			args: args{
				street:     "test_street",
				number:     "test_number",
				complement: "complement",
				city:       "Aracaju",
			},
			expect: expect{
				data: Address{}, err: fmt.Errorf("address: state cannot be empty"),
			},
		},
		{
			title: "Should return an error when zipCode is empty",
			args: args{
				street:     "test_street",
				number:     "test_number",
				complement: "complement",
				city:       "Aracaju",
				state:      "Sergipe",
			},
			expect: expect{
				data: Address{}, err: fmt.Errorf("address: zipCode cannot be empty"),
			},
		},
		{
			title: "Sucess",
			args: args{
				street:     "test_street",
				number:     "test_number",
				complement: "complement",
				city:       "Aracaju",
				state:      "Sergipe",
				zipCode:    "4900",
			},
			expect: expect{
				data: Address{
					street:     "test_street",
					number:     "test_number",
					complement: "complement",
					city:       "Aracaju",
					state:      "Sergipe",
					zipCode:    "4900",
				}, err: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			response, err := NewAddress(tt.args.street, tt.args.number, tt.args.complement, tt.args.city, tt.args.state, tt.args.zipCode)

			assert.Equal(tt.expect.err, err)
			assert.Equal(tt.expect.data.Street(), response.Street())
			assert.Equal(tt.expect.data.Number(), response.Number())
			assert.Equal(tt.expect.data.Complement(), response.Complement())
			assert.Equal(tt.expect.data.City(), response.City())
			assert.Equal(tt.expect.data.State(), response.State())
			assert.Equal(tt.expect.data.ZipCode(), response.ZipCode())
		})
	}
}
