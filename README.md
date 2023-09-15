[中文简介](README.md)    

gentplmode 把数据库的表转换成`go`语言的`struct`，支持 `PostgreSQL`, `MySQL`    

## 安装 
安装到`GOPATH`的 `bin`目录.
```
go clean -modcache
go get -u github.com/guer168/gentplmode/cmd/gentplmode		#go 1.16
go install github.com/guer168/gentplmode/cmd/gentplmode@v1.0.1   	#go 1.17及以上
```
## 卸载 
卸载 `gentplmode`.
```
go clean -i github.com/guer168/gentplmode/cmd/gentplmode
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
  -drive_engine string
        format the data structure to the corresponding database engine

```


## 命令

### MySql
`-target`为 `mysql`
默认生成数据库内的所有表
```
gentplmode -target=mysql -dsn="root:123456@tcp(127.0.0.1:3306)/test" 
```
生成指定表，使用 `-table_names`
```
gentplmode -target=mysql -dsn="root:123456@tcp(127.0.0.1:3306)/test" -table_names=employee -table_names=user
```
格式指定数据格式，使用 `-drive_engine`    
```
gentplmode -target=mysql -dsn="root:123456@tcp(127.0.0.1:3306)/test" -table_names=employee -table_names=user -drive_engine=gorm
```
指定生成目录，使用 `-dir`    
```
gentplmode -target=mysql -dsn="root:123456@tcp(127.0.0.1:3306)/test" -table_names=employee -table_names=user -drive_engine=gorm -dir="./model"
```
按指定模板生成，使用 `-template_path`    
```
gentplmode -target=mysql -dsn="root:123456@tcp(127.0.0.1:3306)/test" -table_names=employee -table_names=user -drive_engine=gorm -dir="./model" -template_path="D:/test/test_template.tpl"
```

### PostgreSql
`-target` 为 `postgresql`或者`pg`
默认生成数据库内的所有表
```
gentplmode -target=pg -dsn="postgres://:@127.0.0.1:5432/test?sslmode=disable"
```
生成指定表，使用 `-table_names`   
```
gentplmode -target=pg -dsn="postgres://:@127.0.0.1:5432/test?sslmode=disable" -table_names=employee -table_names=user
```
格式指定数据格式，使用 `-drive_engine`  
```
gentplmode -target=pg -dsn="postgres://:@127.0.0.1:5432/test?sslmode=disable" -table_names=employee -table_names=user -drive_engine=db
```
指定生成目录或文件相对路径，使用 `-dir`   
```
gentplmode -target=pg -dsn="postgres://:@127.0.0.1:5432/test?sslmode=disable" -table_names=employee -table_names=user -drive_engine=db -dir="./model"
或
gentplmode -target=pg -dsn="postgres://:@127.0.0.1:5432/test?sslmode=disable" -table_names=employee -table_names=user -drive_engine=db -dir="./model/test.go"
```
自定义 template 使用 `-template_path` 自定义模板 
```
gentplmode  -target=pg -dsn="postgres://:@127.0.0.1:5432/test?sslmode=disable" -package_name=db_model -drive_engine=db -dir="./model" -template_path="D:/test/test_template.tpl" 
```

自定义模板 template 列子1：
```
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
```

自定义模板 template 列子2：
```
package {{param "packageName"}}

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

{{$firstChar := FirstCharacter .Name}}
{{$camelizeStructName := CamelizeStr .Name false}}

{{$structName := CamelizeStr .Name true}}

type {{$unPreTableNameUpper}} struct {
{{- range .Columns}}
	{{CamelizeStr .Name true}} {{.GoType}} {{.Tag}} {{.Comment}}
{{- end}}
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
```

参考变量：
```
{{param "packageName"}}、{{$packageName}} 包名

{{$packageNameFirstUpper}} 首字符大写包名

{{.Name}} 数据表原名

{{$firstChar}} 数据表首字母

{{$structName}} 结构体名

{{$unPreTableName}} 无前缀表名

{{$unPreTableNameFirstUpper}} 无前缀表名首字母大写

{{$camelizeStructName}} 首字母小写结构体名

{{.FieldName 0}} 获取字段下标0的字段名，其它字段把0换成对应下标值

{{.FieldType 0}} 获取字段下标0的字段类型，其它字段把0换成对应下标值
```

参考方法：
```
{{CamelizeStr string bool}} 转换驼峰 参数1：处理字符串 参数2：true=首字符大写 false=首字符非大写

{{FirstCharacter string}} 获取首字母 参数1：处理字符串

{{FirstLowerWord string}} 首字母小写 参数1：处理字符串

{{RemovePrefix string string}} 移除表前缀 参数1：表名 参数2：表前缀

{{StrToLower string string}} 转换小写字母 参数1：处理字符串	

{{Replace string string string}} 替换字符串 参数1：处理字符串 参数2：要替换字符 参数3：替换成为字符串

{{Add int int}} 加法 参数1：数字 参数2: 数字	

{{now}} 获取当前时间

{{.FieldName int}} 获取字段下标n的字段名 参数1：其它字段把n换成对应下标值

{{.FieldType int}} 获取字段下标n的字段类型 参数1：其它字段把n换成对应下标值
```
