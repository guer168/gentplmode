package do

import (
{{ if .Imports }}
    {{- range .Imports}}
        "{{.}}"
    {{- end}}
{{end}}
    "github.com/jinzhu/gorm"
)


{{$packageName := param "packageName"}}
{{$packageNameFirstUpper := CamelizeStr $packageName true}}

{{$unPreTableName := RemovePrefix .Name "erp_"}}
{{$unPreTableNameUpper := CamelizeStr $unPreTableName true}}

{{$firstChar := FirstCharacter $unPreTableName}}
{{$camelizeStructName := CamelizeStr .Name false}}

{{$structName := CamelizeStr .Name true}}

type {{$unPreTableNameUpper}} struct {
{{- range .Columns}}
	{{CamelizeStr .Name true}} {{.GoType}} {{.Tag}} {{.Comment}}
{{- end}}
}


type Where{{$unPreTableNameUpper}} struct {
    {{- $first := true }}
    {{- range $index, $column := .Columns }}
        {{- if not $first }}
            {{CamelizeStr .Name true }} string `json:"{{StrToLower .Name}},omitempty"`
        {{- end }}
        {{- $first = false }}
    {{- end }}
}

// TableName
//  @Description: 获取表名
//  @return string
func ({{$firstChar}} *{{$unPreTableNameUpper}}) TableName() string {
	return "{{.Name}}"
}

// BeforeCreate
//  @Description: 创建钩子函数
//  @param scope
//  @return error
func ({{$firstChar}} *{{$unPreTableNameUpper}}) BeforeCreate(scope *gorm.Scope) error {
	//scope.SetColumn("created_at", time.Now())
	//scope.SetColumn("updated_at", time.Now())
	return nil
}

// BeforeUpdate
//  @Description: 更新钩子函数
//  @param scope
//  @return error
func ({{$firstChar}} *{{$unPreTableNameUpper}}) BeforeUpdate(scope *gorm.Scope) error {
	//scope.SetColumn("updated_at", time.Now())
	return nil
}