package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

func main() {
	funcMap := template.FuncMap{
		"title": strings.Title,
	}

	const templateText = `
Input: {{printf "%q" .}}
InputLen: {{len .}}
Output : {{title .}}
`
	tmpl, err := template.New("titleTest").Funcs(funcMap).Parse(templateText)
	if err != nil {
		log.Fatalf("parsing: %s", err)
	}
	err = tmpl.Execute(os.Stdout, "FuncMap test")
	if err != nil {
		log.Fatalf("execution: %s", err)
	}
}


