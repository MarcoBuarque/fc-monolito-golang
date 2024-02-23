package valueobject

import "fmt"

type Address struct {
	street     string
	number     string
	complement string
	city       string
	state      string
	zipCode    string
}

func NewAddress(street, number, complement, city, state, zipCode string) (Address, error) {
	if street == "" {
		return Address{}, fmt.Errorf("address: street cannot be empty")
	}

	if number == "" {
		return Address{}, fmt.Errorf("address: number cannot be empty")
	}

	if complement == "" {
		return Address{}, fmt.Errorf("address: complement cannot be empty")
	}

	if city == "" {
		return Address{}, fmt.Errorf("address: city cannot be empty")
	}

	if state == "" {
		return Address{}, fmt.Errorf("address: state cannot be empty")
	}

	if zipCode == "" {
		return Address{}, fmt.Errorf("address: zipCode cannot be empty")
	}

	return Address{street, number, complement, city, state, zipCode}, nil
}

func (data Address) Street() string     { return data.street }
func (data Address) Number() string     { return data.number }
func (data Address) Complement() string { return data.complement }
func (data Address) City() string       { return data.city }
func (data Address) State() string      { return data.state }
func (data Address) ZipCode() string    { return data.zipCode }
