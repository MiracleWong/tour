package main

import (
	"html/template"
	"os"
	"strings"
)

// Go 模板解析
//

const templateText = `
Output 0: {{ title .Name1 }}
Output 1: {{ title .Name2 }} 
Output 2: {{ .Name3 | title }} 
`

// 如果想要执行时，反注释掉 main，注释 ll()即可
// func main() {
func ll() {
	funcMap := template.FuncMap{"title": strings.Title}
	tpl, _ := template.New("go-programming-tour").Funcs(funcMap).Parse(templateText)
	data := map[string]string{
		"Name1": "go",
		"Name2": "programming",
		"Name3": "tour",
	}
	_ = tpl.Execute(os.Stdout, data)
}
