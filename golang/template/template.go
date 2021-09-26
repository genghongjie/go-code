package main

import (
	"fmt"
	"os"
	"text/template"
)

type Latlng struct {
	Lat float32
	Lng float32
}

func (latlng Latlng) String() string {
	return fmt.Sprintf("%g/%g", latlng.Lat, latlng.Lng)
}

func main() {
	data := []template.FuncMap{}
	data = append(data, template.FuncMap{"name": "dotcoo1", "url": "http://www.dotcoo.com/", "latlng": Latlng{24.1, 135.1}})
	data = append(data, template.FuncMap{"name": "dotcoo2", "url": "http://www.dotcoo.com/", "latlng": Latlng{24.2, 135.2}})
	data = append(data, template.FuncMap{"name": "dotcoo2", "url": "http://www.dotcoo.com/", "latlng": Latlng{24.3, 135.3}})

	datatpl := `{{range .}}{{template "user" .}}{{end}}`
	usertpl := `{{define "user"}}name:{{.name}}, url:{{.url}}, latlng:{{.latlng}} lat:{{.latlng.Lat}} lng:{{.latlng.Lng}}
{{end}}`

	tpl, err := template.New("data").Parse(datatpl)
	if err != nil {
		panic(err)
	}
	_, err = tpl.Parse(usertpl)
	if err != nil {
		panic(err)
	}

	err = tpl.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
	println()

	tpl.ExecuteTemplate(os.Stdout, "data", data)
	newtmplate := `{{define "T1"}}ONE{{end}}
{{define "T2"}}TWO{{end}}
{{define "T3"}}{{template "T1"}} {{template "T2"}}{{end}}
{{template "T3"}}`
	_, err = tpl.Parse(newtmplate)
	if err != nil {
		panic(err)
	}
	err = tpl.Execute(os.Stdout, "")
	if err != nil {
		panic(err)
	}

	println()
}
