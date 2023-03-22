package main

import (
	"bytes"
	"fmt"
	"text/template"
)

func main() {
	//txt := `{{printf "%q" "output"}}`
	txt := `你好 {{.IP}}`
	str := ReplaceTemplate(txt, nil)
	fmt.Printf(str)
}

//使用template替换变量的方法
func ReplaceTemplate(str string, vars ...interface{}) string {
	var tmplBytes bytes.Buffer
	tpl, err := template.New("test").Parse(str)
	if err == nil {
		tpl.Execute(&tmplBytes, vars)
	}
	return tmplBytes.String()
}
