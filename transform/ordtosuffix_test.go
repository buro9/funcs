package transform

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrdToSuffix(t *testing.T) {
	assert.Equal(t, `st`, OrdToSuffix(1))
	assert.Equal(t, `nd`, OrdToSuffix(2))
	assert.Equal(t, `rd`, OrdToSuffix(3))
	assert.Equal(t, `th`, OrdToSuffix(4))
	assert.Equal(t, `th`, OrdToSuffix(20))
	assert.Equal(t, `st`, OrdToSuffix(21))
	assert.Equal(t, `nd`, OrdToSuffix(22))
	assert.Equal(t, `rd`, OrdToSuffix(23))
	assert.Equal(t, `th`, OrdToSuffix(24))
	assert.Equal(t, `th`, OrdToSuffix(30))
	assert.Equal(t, `st`, OrdToSuffix(31))
	assert.Equal(t, `nd`, OrdToSuffix(32))
	assert.Equal(t, `rd`, OrdToSuffix(33))
	assert.Equal(t, `th`, OrdToSuffix(34))

	// Nothing for a float or string
	assert.Equal(t, ``, OrdToSuffix(1.1))
	assert.Equal(t, ``, OrdToSuffix("hello"))

}
