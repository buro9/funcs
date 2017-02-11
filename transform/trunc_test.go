package transform

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrunc(t *testing.T) {
	assert.Equal(t, `...`, Trunc(`Hello World`, 0))
	assert.Equal(t, `H...`, Trunc(`Hello World`, 1))
	assert.Equal(t, `Hello...`, Trunc(`Hello World`, 5))
	assert.Equal(t, `Hello Wo...`, Trunc(`Hello World`, 8))
	assert.Equal(t, `Hello World`, Trunc(`Hello World`, 50))

	assert.Equal(t, `...`, Trunc(`Привет Мир`, 0))
	assert.Equal(t, `П...`, Trunc(`Привет Мир`, 1))
	assert.Equal(t, `Приве...`, Trunc(`Привет Мир`, 5))
	assert.Equal(t, `Привет М...`, Trunc(`Привет Мир`, 8))
	assert.Equal(t, `Привет Мир`, Trunc(`Привет Мир`, 50))
}
