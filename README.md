[中文简介](README.md)    
[English](README_US.md)

gentplmode 把数据库的表转换成`go`语言的`struct`，支持 `PostgreSQL`, `MySQL`    

## 安装 
安装到`GOPATH`的 `bin`目录.
```
GO111MODULE=on go get -u github.com/guer168/gentplmode/cmd/gentplmode
```
### 帮助
```
gentplmode -help 
```
```
Usage of gentplmode:
  -dir string
        Destination dir for files generated. (default "./tmp")
  -dsn string
        dsn (default "postgresql")
  -package_name string
        package name default model. (default "model")
  -table_names value
        if it is empty, will generate all tables in database
  -target string
        mysql postgresql[pg] (default "postgresql")
  -template_path string
        custom template file path
  -formatDriveEngine string
        format the data structure to the corresponding database engine

```


## 命令

### MySql
`-target`为 `mysql`
默认生成数据库内的所有表
```
gentplmode -target=mysql -dsn="root:123456@tcp(127.0.0.1:3306)/test" 
```
也可以生成指定表，使用 `-table_names` 指定想生成的表    
```
gentplmode -target=mysql -dsn="root:123456@tcp(127.0.0.1:3306)/test" -table_names=employee -table_names=user
```

### PostgreSql
`-target` 为 `postgresql`或者`pg`
默认生成数据库内的所有表
```
gentplmode -target=pg -dsn="postgres://:@127.0.0.1:5432/test?sslmode=disable"
```
使用 `-table_names` 指定想生成的表    
```
gentplmode -target=pg -dsn="root:123456@tcp(127.0.0.1:3306)/test" -table_names=employee -table_names=user
```

自定义 template 使用 `-template_path` 自定义模板 
```
gentplmode  -target=pg -dsn="postgres://:@127.0.0.1:5432/test?sslmode=disable" -package_name=db_model -template_path=../../test/test_template.tml 
```

自定义 template 列子参考：
```
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
//  @receiver {{$firstChar}}
//  @return string
func ({{param "packageName"}} *{{$structName}}) TableName() string {
	return "{{.Name}}"
}
`
```

参考变量：
```
{{param "packageName"}} 包名

{{.Name}} 数据表原名

{{$firstChar} 数据表首字母

{{$structName}} 结构体名

{{$camelizeStructName}} 首字母小写结构体名
```
