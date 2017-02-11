package funcs

import (
	"html/template"
	"strings"

	"github.com/buro9/funcs/inspect"
	"github.com/buro9/funcs/safe"
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

		// Trusted content
		"safeCSS":      safe.CSS,
		"safeHTML":     safe.HTML,
		"safeHTMLAttr": safe.HTMLAttr,
		"safeJS":       safe.JS,
		"safeURL":      safe.URL,

		// Check vars and returns boolean answer to questions about their state
		"isNil": inspect.IsNil,
	}
}
