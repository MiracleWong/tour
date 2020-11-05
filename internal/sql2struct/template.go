package sql2struct

//const structTpl = `type {{.TableName | ToCamelCase}} struct {
//{{range .Columns}}	{{ $length := len .Comment}} {{ if gt $length 0 }}// {{.Comment}} {{else}}// {{.Name}} {{ end }}
//	{{ $typeLen := len .Type }} {{ if gt $typeLen 0 }}{{.Name | ToCamelCase}}	{{.Type}}	{{.Tag}}{{ else }}{{.Name}}{{ end }}
//{{end}}}
//func (model {{.TableName | ToCamelCase}}) TableName() string {
//	return "{{.TableName}}"
//}`
const structTpl = `type {{.TableName | ToCamelCase }} struct {
{{range .Columns}}
	{{ $length := len .Comment}}
	{{ if gt $length 0}}
 // {{.Comment}}
	{{ else }}
	// {{.Name}}
	{{ end }}
	{{ $typeLen := len .Type }}
	{{ if gt $typeLen 0 }}
		{{.Name | ToCamelCase}}	{{.Type}}	{{.Tag}}
	{{ else }}
		{{.Name}}
	{{end}}
{{end}}
}

func (model {{.TableName | ToCamelCase }}) TableName()  string {
	return "{{.TableName}}"
}
`

type StructTemplate struct {
	structTpl string
}

// 用于存储转换后的 Go 结构体的所有字段信息
type StructColumn struct {
	Name    string
	Type    string
	Tag     string
	Comment string
}

// 存储最终用于渲染的模板对象信息
type StructTemplateDB struct {
	TableName string
	Columns   []*StructColumn
}
