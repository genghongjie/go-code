package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

func main() {
	txt := `{{printf "%q" "output"}}`

	tpl, err := template.New("test").Parse(txt)
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.Execute(os.Stdout, "")
	fmt.Println()
}
