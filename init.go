package funcs

import (
	"html/template"
	"strings"

	"github.com/buro9/funcs/check"
)

// Map is a `template.FuncMap` containing all of the funcs within the child
// packages as well as some from the core Go library
var Map template.FuncMap

func init() {
	Map = template.FuncMap{

		// String manipulation
		"lower": strings.ToLower,
		"title": strings.Title,
		"upper": strings.ToUpper,

		// Check vars and returns boolean answer to questions about their state
		"isNil": check.IsNil,
	}
}
