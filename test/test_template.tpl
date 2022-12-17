package {{param "packageName"}}
{{ if .Imports }}
import (
{{- range .Imports}}
	"{{.}}"
{{- end}}
)
{{end}}

{{$packageName := param "packageName"}}
{{$packageNameFirstUpper := CamelizeStr $packageName true}}

{{$unPreTableName := .RemovePrefix .Name "jy_"}}
{{$unPreTableNameFirstUpper := CamelizeStr $unPreTableName true}

{{$structName := CamelizeStr .Name true}}

{{$firstChar := FirstCharacter .Name}}
{{$camelizeStructName := CamelizeStr .Name false}}

type {{$structName}} struct {
{{- range .Columns}}
	{{CamelizeStr .Name true}} {{.GoType}} {{.Tag}} {{.Comment}}
{{- end}}
}
var {{$unPreTableNameFirstUpper}}{{$packageNameFirstUpper}} *{{$structName}}

// TableName
//  @Description: 获取表名
//  @return string
func ({{$firstChar}} *{{$structName}}) TableName() string {
	return "{{.Name}}"
}

//原版(本系统由原版修改)： git地址 https://gitee.com/guer168/yggdrasill.git 或 https://github.com/lpxxn/yggdrasill.git
//参考变量
//{{param "packageName"}}、{{$packageName}} 包名
//{{$packageNameFirstUpper}}   首字符大写包名
//{{.Name}} 数据表原名
//{{$firstChar}} 数据表首字母
//{{$structName}} 结构体名
//{{$unPreTableName}} 无前缀表名
//{{$unPreTableNameFirstUpper}} 无前缀表名首字母大写
//{{$camelizeStructName}} 首字母小写结构体名
//{{.FieldName 0}} 获取字段下标0的字段名，其它字段把0换成对应下标值
//{{.FieldType 0}} 获取字段下标0的字段类型，其它字段把0换成对应下标值