package {{param "packageName"}}
{{ if .Imports }}
import (
{{- range .Imports}}
	"{{.}}"
{{- end}}
    "github.com/jinzhu/gorm"
)
{{end}}

{{$packageName := param "packageName"}}
{{$packageNameFirstUpper := CamelizeStr $packageName true}}

{{$unPreTableName := .RemovePrefix .Name "jy_"}}
{{$unPreTableNameUpper := CamelizeStr $unPreTableName true}}

{{$firstChar := FirstCharacter .Name}}
{{$camelizeStructName := CamelizeStr .Name false}}

{{$structName := CamelizeStr .Name true}}

type {{$structName}} struct {
{{- range .Columns}}
	{{CamelizeStr .Name true}} {{.GoType}} {{.Tag}} {{.Comment}}
{{- end}}
}
var {{$unPreTableNameUpper}}{{$packageNameFirstUpper}} *{{$structName}}

// TableName
//  @Description: 获取表名
//  @return string
func ({{param "packageName"}} *{{$structName }}) TableName() string {
	return "{{.Name}}"
}

// BeforeCreate
//  @Description: 创建钩子函数
//  @param scope
//  @return error
func ({{param "packageName"}} *{{$structName }}) BeforeCreate(scope *gorm.Scope) error {
	//scope.SetColumn("created_at", time.Now())
	//scope.SetColumn("updated_at", time.Now())
	return nil
}

// BeforeUpdate
//  @Description: 更新钩子函数
//  @param scope
//  @return error
func ({{param "packageName"}} *{{$structName }}) BeforeUpdate(scope *gorm.Scope) error {
	//scope.SetColumn("updated_at", time.Now())
	return nil
}