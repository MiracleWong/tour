package sql2struct

import (
	"fmt"
	"github.com/MiracleWong/tour/internal/word"
	"os"
	"text/template"
)

//const structTpl = `type {{.TableName | ToCamelCase}} struct {
//{{range .Columns}}	{{ $length := len .Comment}} {{ if gt $length 0 }}// {{.Comment}} {{else}}// {{.Name}} {{ end }}
//	{{ $typeLen := len .Type }} {{ if gt $typeLen 0 }}{{.Name | ToCamelCase}}	{{.Type}}	{{.Tag}}{{ else }}{{.Name}}{{ end }}
//{{end}}}
//func (model {{.TableName | ToCamelCase}}) TableName() string {
//	return "{{.TableName}}"
//}`
//模板对象声明
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

func NewStructTemplate() *StructTemplate {
	return &StructTemplate{structTpl: structTpl}
}

// 通过对查询 COLUMNS 表所组装得到的tbColumns，进行进一步的分解和转换
func (t *StructTemplate) AssemblyColumns(tbColumns []*TableColumn) []*StructColumn {
	tplColumns := make([]*StructColumn, 0, len(tbColumns))
	for _, column := range tbColumns {
		tag := fmt.Sprintf("`"+"json:"+"\"%s\""+"`", column.ColumnName)
		tplColumns = append(tplColumns, &StructColumn{
			Name:    column.ColumnName,
			Type:    DBTypeToStructType[column.DataType],
			Tag:     tag,
			Comment: column.ColumnComment,
		})
	}
	return tplColumns
}

// 对模块渲染的自定义函数和模板对象处理
func (t *StructTemplate) Generate(tableName string, tplColumns []*StructColumn) error {
	tpl := template.Must(template.New("sql2struct").Funcs(template.FuncMap{
		"ToCamelCase": word.UnderscoreToUpperCamelCase,
	}).Parse(t.structTpl))

	tplDB := StructTemplateDB{
		TableName: tableName,
		Columns:   tplColumns,
	}
	err := tpl.Execute(os.Stdout, tplDB)
	if err != nil {
		return err
	}
	return nil
}
