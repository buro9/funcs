package inspect

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSet(t *testing.T) {
	aSlice := []interface{}{1, 2, 3, 5}
	aMap := map[string]interface{}{"a": 1, "b": 2}

	assert.True(t, IsSet(aSlice, 2))
	assert.True(t, IsSet(aMap, "b"))
	assert.False(t, IsSet(aSlice, 22))
	assert.False(t, IsSet(aMap, "bc"))
}
