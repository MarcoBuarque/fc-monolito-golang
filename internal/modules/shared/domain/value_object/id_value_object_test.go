package valueobject

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateID(t *testing.T) {
	require := require.New(t)
	assert := assert.New(t)

	id := CreateID("")

	require.NotEmpty(id)

	id = CreateID("xpto-10")
	assert.Equal("xpto-10", id.id)

	assert.Equal("xpto-10", id.ToString())

}
