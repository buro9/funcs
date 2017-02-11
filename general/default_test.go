package general

import (
	"reflect"
	"testing"
	"time"
)

// TestDefaultFunc ...
//
// Copyright The Hugo Authors
// License Apache 2
// https://github.com/spf13/hugo/blob/master/LICENSE.md
// https://github.com/spf13/hugo/blob/master/tpl/template_funcs.go
func TestDefaultFunc(t *testing.T) {
	then := time.Now()
	now := time.Now()

	for i, this := range []struct {
		dflt     interface{}
		given    interface{}
		expected interface{}
	}{
		{true, false, false},
		{"5", 0, "5"},

		{"test1", "set", "set"},
		{"test2", "", "test2"},
		{"test3", nil, "test3"},

		{[2]int{10, 20}, [2]int{1, 2}, [2]int{1, 2}},
		{[2]int{10, 20}, [0]int{}, [2]int{10, 20}},
		{[2]int{100, 200}, nil, [2]int{100, 200}},

		{[]string{"one"}, []string{"uno"}, []string{"uno"}},
		{[]string{"two"}, []string{}, []string{"two"}},
		{[]string{"three"}, nil, []string{"three"}},

		{map[string]int{"one": 1}, map[string]int{"uno": 1}, map[string]int{"uno": 1}},
		{map[string]int{"one": 1}, map[string]int{}, map[string]int{"one": 1}},
		{map[string]int{"two": 2}, nil, map[string]int{"two": 2}},

		{10, 1, 1},
		{10, 0, 10},
		{20, nil, 20},

		{float32(10), float32(1), float32(1)},
		{float32(10), 0, float32(10)},
		{float32(20), nil, float32(20)},

		{complex(2, -2), complex(1, -1), complex(1, -1)},
		{complex(2, -2), complex(0, 0), complex(2, -2)},
		{complex(3, -3), nil, complex(3, -3)},

		{struct{ f string }{f: "one"}, struct{ f string }{}, struct{ f string }{}},
		{struct{ f string }{f: "two"}, nil, struct{ f string }{f: "two"}},

		{then, now, now},
		{then, time.Time{}, then},
	} {
		res, err := Default(this.dflt, this.given)
		if err != nil {
			t.Errorf("[%d] default returned an error: %s", i, err)
			continue
		}
		if !reflect.DeepEqual(this.expected, res) {
			t.Errorf("[%d] default returned %v, but expected %v", i, res, this.expected)
		}
	}
}
