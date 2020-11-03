package sql2struct

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
