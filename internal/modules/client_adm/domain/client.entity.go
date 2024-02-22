package domain

import (
	"fmt"

	"github.com/MarcoBuarque/monolito/internal/modules/client_adm/repository"
	"github.com/MarcoBuarque/monolito/internal/modules/shared/domain/entity"
	valueobject "github.com/MarcoBuarque/monolito/internal/modules/shared/domain/value_object"
)

type ClientEntity struct {
	name    string
	address string
	email   string
	entity.BaseEntity
}

func NewClient(id, name, address, email string) (ClientEntity, error) {
	if name == "" {
		return ClientEntity{}, fmt.Errorf("clientEntity: name cannot be empty")
	}

	if address == "" {
		return ClientEntity{}, fmt.Errorf("clientEntity: description cannot be empty")
	}

	if email == "" {
		return ClientEntity{}, fmt.Errorf("clientEntity: email cannot be empty")
	}

	return ClientEntity{
		name:       name,
		address:    address,
		email:      email,
		BaseEntity: entity.CreateBaseEntity(valueobject.CreateID(id)),
	}, nil
}

func (data ClientEntity) Name() string    { return data.name }
func (data ClientEntity) Address() string { return data.address }
func (data ClientEntity) Email() string   { return data.email }
func (data ClientEntity) ToData() repository.Client {
	return repository.Client{
		ID:        string(data.ID().ToString()),
		Name:      data.name,
		Address:   data.address,
		Email:     data.email,
		CreatedAt: data.CreatedAt(),
		UpdatedAt: data.UpdatedAt(),
	}
}
