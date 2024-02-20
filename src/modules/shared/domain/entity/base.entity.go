package entity

import (
	"time"

	valueobject "github.com/MarcoBuarque/monolito/src/modules/shared/domain/value_object"
)

type BaseEntity struct {
	id        valueobject.ID
	createdAt time.Time
	updatedAt time.Time
}

func CreateBaseEntity(id valueobject.ID) BaseEntity {
	return BaseEntity{
		id:        id,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}
}

func (data BaseEntity) ID() valueobject.ID { return data.id }

func (data BaseEntity) CreatedAt() time.Time { return data.createdAt }

func (data BaseEntity) UpdatedAt() time.Time { return data.updatedAt }

func (data *BaseEntity) SetUpdatedAt(newTime time.Time) { data.updatedAt = newTime }
