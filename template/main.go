package main

import (
	"html/template"
	"os"
	"strings"
)

// Go 模板解析
// template 是 Go 的文本模板引擎
// html/template: 基于模板输出文本内容
// text/template: 基于模板输出HTML 格式内容
// 双层大括号{{和}} 为标识符。在 template 中，所有的动作、数据评估、控制流转，都需要{{}} 包裹
// . 为点（DOT）。在 template 中，会根据点标识符进行模板变量的渲染。
// 函数调用：| 管道符。在 template 中，会把管道符前面的运算结果作为参数传递给管道符后面的函数
// templateText 预定义的解析模板
const templateText = `
Output 0: {{ title .Name1 }}
Output 1: {{ title .Name2 }} 
Output 2: {{ .Name3 | title }} 
Real Name: {{ .UserName | title }} 
`

// 如果想要执行时，反注释掉 main，注释 ll()即可，还需要将 main.go 中的 main 改名字
func main() {
	// 名称标识：FuncMap 注册了自定义函数
	funcMap := template.FuncMap{"title": strings.Title}
	// New 根据给的的名称标识，创建新的模板对象
	// Parse 将常量templateText，解析为当前文本的主体内容
	tpl, _ := template.New("go-programming-tour").Funcs(funcMap).Parse(templateText)
	data := map[string]string{
		"Name1":    "go",
		"Name2":    "programming",
		"Name3":    "tour",
		"UserName": "MiracleWong",
	}
	// Execute 进行模板渲染
	_ = tpl.Execute(os.Stdout, data)
}
