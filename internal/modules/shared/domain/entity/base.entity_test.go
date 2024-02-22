package entity

import (
	"testing"

	valueobject "github.com/MarcoBuarque/monolito/internal/modules/shared/domain/value_object"
	"github.com/stretchr/testify/assert"
)

func TestNewBaseEntity(t *testing.T) {
	assert := assert.New(t)

	baseEntity := NewBaseEntity(valueobject.NewID("abc"))

	assert.Equal(baseEntity.ID().ToString(), "abc")
}
