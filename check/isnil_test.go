package check

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIsNil(t *testing.T) {
	var sp *string
	assert.True(t, IsNil(sp))
	assert.True(t, IsNil(nil))

	s := "testing"
	sp = &s
	assert.False(t, IsNil(sp))
	assert.False(t, IsNil(s))

	m := map[string]interface{}{"this": "that"}
	assert.False(t, IsNil(m))
	assert.True(t, IsNil(m["other"]))

	var tm *time.Time
	assert.True(t, IsNil(&tm))
	now := time.Now()
	tm = &now
	assert.False(t, IsNil(&tm))
}
