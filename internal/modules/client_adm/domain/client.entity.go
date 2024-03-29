package domain

import (
	"fmt"

	"github.com/MarcoBuarque/fc-monolito-golang/internal/modules/shared/domain/entity"
	valueobject "github.com/MarcoBuarque/fc-monolito-golang/internal/modules/shared/domain/value_object"
)

type ClientEntity struct {
	name     string
	email    string
	document valueobject.Document
	address  valueobject.Address
	entity.BaseEntity
}

func NewClient(id, name, email string, document valueobject.Document, address valueobject.Address) (ClientEntity, error) {
	if name == "" {
		return ClientEntity{}, fmt.Errorf("clientEntity: name cannot be empty")
	}

	if email == "" {
		return ClientEntity{}, fmt.Errorf("clientEntity: email cannot be empty")
	}

	return ClientEntity{
		name:       name,
		email:      email,
		document:   document,
		address:    address,
		BaseEntity: entity.NewBaseEntity(valueobject.NewID(id)),
	}, nil
}

func (data ClientEntity) Name() string                   { return data.name }
func (data ClientEntity) Address() valueobject.Address   { return data.address }
func (data ClientEntity) Document() valueobject.Document { return data.document }
func (data ClientEntity) Email() string                  { return data.email }
