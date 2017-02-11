package transform

import (
	"reflect"
	"testing"
)

// TestDictionary ...
//
// Copyright The Hugo Authors
// License Apache 2
// https://github.com/spf13/hugo/blob/master/LICENSE.md
// https://github.com/spf13/hugo/blob/master/tpl/template_funcs.go
func TestDictionary(t *testing.T) {
	for i, this := range []struct {
		v1            []interface{}
		expecterr     bool
		expectedValue map[string]interface{}
	}{
		{[]interface{}{"a", "b"}, false, map[string]interface{}{"a": "b"}},
		{[]interface{}{5, "b"}, true, nil},
		{[]interface{}{"a", 12, "b", []int{4}}, false, map[string]interface{}{"a": 12, "b": []int{4}}},
		{[]interface{}{"a", "b", "c"}, true, nil},
	} {
		r, e := Dictionary(this.v1...)

		if (this.expecterr && e == nil) || (!this.expecterr && e != nil) {
			t.Errorf("[%d] got an unexpected error: %s", i, e)
		} else if !this.expecterr {
			if !reflect.DeepEqual(r, this.expectedValue) {
				t.Errorf("[%d] got %v but expected %v", i, r, this.expectedValue)
			}
		}
	}
}
