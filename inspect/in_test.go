package inspect

import "testing"

// TestIn ...
//
// Copyright The Hugo Authors
// License Apache 2
// https://github.com/spf13/hugo/blob/master/LICENSE.md
// https://github.com/spf13/hugo/blob/master/tpl/template_funcs.go
func TestIn(t *testing.T) {
	for i, this := range []struct {
		v1     interface{}
		v2     interface{}
		expect bool
	}{
		{[]string{"a", "b", "c"}, "b", true},
		{[]interface{}{"a", "b", "c"}, "b", true},
		{[]interface{}{"a", "b", "c"}, "d", false},
		{[]string{"a", "b", "c"}, "d", false},
		{[]string{"a", "12", "c"}, 12, false},
		{[]int{1, 2, 4}, 2, true},
		{[]interface{}{1, 2, 4}, 2, true},
		{[]interface{}{1, 2, 4}, nil, false},
		{[]interface{}{nil}, nil, false},
		{[]int{1, 2, 4}, 3, false},
		{[]float64{1.23, 2.45, 4.67}, 1.23, true},
		{[]float64{1.234567, 2.45, 4.67}, 1.234568, false},
		{"this substring should be found", "substring", true},
		{"this substring should not be found", "subseastring", false},
	} {
		result := In(this.v1, this.v2)

		if result != this.expect {
			t.Errorf("[%d] got %v but expected %v", i, result, this.expect)
		}
	}
}
