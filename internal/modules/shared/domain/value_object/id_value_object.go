package valueobject

import "github.com/google/uuid"

type ID struct {
	id string
}

func CreateID(newID string) ID {
	if newID == "" {
		return ID{id: uuid.New().String()}
	}

	return ID{id: newID}
}

func (data ID) ToString() string { return data.id }
