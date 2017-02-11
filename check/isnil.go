package check

import "reflect"

// IsNil returns true if the provided variable is nil
func IsNil(v interface{}) (isNil bool) {
	if isNil = (v == nil); isNil {
		return
	}

	vo := reflect.ValueOf(v)
	switch vo.Kind() {
	case reflect.String:
		// always ok
	default:
		isNil = vo.IsNil()
	}

	return
}
