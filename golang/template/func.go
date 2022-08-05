package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	//var test = func(s string) string {
	//	return s[4:]
	//}
	type SqlResult struct {
		FinalOut []string
	}
	input := SqlResult{
		FinalOut: []string{"1", "2", "3"},
	}
	// First we create a FuncMap with which to register the function.
	funcMap := template.FuncMap{
		// The name "title" is what the function will be called in the template text.
		"value": func(s SqlResult, index int) string {
			return "===" + s.FinalOut[index] + "----"
		},
	}

	// A simple template definition to test our function.
	// We print the input text several ways:
	// - the original
	// - title-cased
	// - title-cased and then printed with %q
	// - printed with %q and then title-cased.
	const templateText = `
Input: {{printf "%q" .}}
Output 0: {{value . final_result:value:}}

`
	// Create a template, add the function map, and parse the text.
	tmpl, err := template.New("titleTest").Funcs(funcMap).Parse(templateText)
	if err != nil {
		log.Fatalf("parsing: %s", err)
	}

	// Run the template to verify the output.
	err = tmpl.Execute(os.Stdout, input)
	if err != nil {
		log.Fatalf("execution: %s", err)
	}

}
