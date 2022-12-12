package templates

var templateHeader = `
// Generated at {{now}}
`

var tableModelTemplate = `
package {{param "packageName"}}
{{ if .Imports }}
import (
{{- range .Imports}}
	"{{.}}"
{{- end}}
)
{{end}}

{{$structName := CamelizeStr .Name true}}

type {{$structName}} struct {
{{- range .Columns}}
	{{CamelizeStr .Name true}} {{.GoType}} ` + "{{.Tag}}" + `
{{- end}}
}
{{$firstChar := FirstCharacter .Name}}
{{$camelizeStructName := CamelizeStr .Name false}}

// TableName
//  @Description: Getting the table name
//  @return string
func ({{param "packageName"}} *{{$structName }}) TableName() string {
	return "{{.Name}}"
}
`

//原版(本系统由原版修改)： git地址 https://gitee.com/guer168/yggdrasill.git 或 https://github.com/lpxxn/yggdrasill.git
//参考变量
//{{param "packageName"}} 包名
//{{.Name}} 数据表原名
//{{$firstChar} 数据表首字母
//{{$structName}} 结构体名
//{{$camelizeStructName}} 首字母小写结构体名
//{{.FieldName 0}} 获取字段下标0的字段名，其它字段把0换成对应下标值
//{{.FieldType 0}} 获取字段下标0的字段类型，其它字段把0换成对应下标值