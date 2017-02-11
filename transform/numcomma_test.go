package transform

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumComma(t *testing.T) {
	assert.Equal(t, `1`, NumComma(1))
	assert.Equal(t, `100`, NumComma(100))
	assert.Equal(t, `1,000`, NumComma(1000))
	assert.Equal(t, `1.1`, NumComma(1.1))
	assert.Equal(t, `1,000.1`, NumComma(1000.1))
}
