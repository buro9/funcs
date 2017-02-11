package funcs

import (
	"bytes"
	"html/template"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	// This test involves defining a template that will test each of our funcs
	// to ensure that it exists and is correctly defined in the template.FuncMap

	type templateData struct {
		MixedCase string
	}
	data := templateData{
		MixedCase: "mixedCASE woRds",
	}

	const templateText = `
lower: {{lower .MixedCase}}
title: {{title .MixedCase}}
upper: {{upper .MixedCase}}

isNil: 
`

	// Create a template, add the function map, and parse the template
	tmpl, err := template.New("test").Funcs(Map).Parse(templateText)
	if err != nil {
		log.Fatalf("parsing: %s", err)
	}

	// Run the template to verify the output.
	var b bytes.Buffer
	err = tmpl.Execute(&b, data)
	if err != nil {
		log.Fatalf("execution: %s", err)
	}

	assert.Equal(
		t,
		`
lower: mixedcase words
title: MixedCASE WoRds
upper: MIXEDCASE WORDS

isNil: 
`,
		b.String(),
	)
}
