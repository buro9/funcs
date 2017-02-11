package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubtract(t *testing.T) {
	assert.Equal(t, 2, Subtract(3, 1))
}
