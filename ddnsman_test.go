package ddnsman

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	upd, err := New(&Configuration{})
	assert.NotNil(t, upd)
	assert.NoError(t, err)
}
