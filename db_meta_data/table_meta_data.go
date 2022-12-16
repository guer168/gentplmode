package db_meta_data

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/guer168/gentplmode/utils"
)

type TableMetaData struct {
	Name    string
	Columns ColumnMetaDataList
}
var tableMetaData *TableMetaData

type TableMetaDataList []*TableMetaData


func (t TableMetaData) Imports() []string {
	imports := map[string]string{}

	for _, column := range t.Columns {
		columnType := column.GoType
		if v, ok := customerColumnDataTypeImport[columnType]; ok {
			imports[columnType] = v
			continue
		}
		switch columnType {
		case "time.Time":
			imports["time.Time"] = "time"
		}
	}
	rev := []string{}
	for _, packageImport := range imports {
		rev = append(rev, packageImport)
	}
	return rev
}

// FieldName
//  @Description: 获取字段名称
//  @receiver t
//  @param index	字段下标
//  @return string
func (t TableMetaData) FieldName(index int) string {
	rev := ""
	for idx, item := range t.Columns {
		if idx == index{
			rev = item.Name
			break
		}
	}
	return rev
}

// FieldType
//  @Description: 获取字段类型
//  @receiver t
//  @param index	字段下标
//  @return string
func (t TableMetaData) FieldType(index int) string {
	rev := ""
	for idx, item := range t.Columns {
		if idx == index{
			rev = item.GoType
			break
		}
	}
	return rev
}

// Replace
//  @Description: 替换字符串
//  @receiver t
//  @param s
//  @param oldStr
//  @param newStr
//  @param n
//  @return string
func (t TableMetaData) Replace(s string, oldStr string, newStr string, n int) string {
	return strings.Replace(s, oldStr, newStr, n)
}

func (t TableMetaData) ColumnsNameWithPrefixAndIgnoreColumn(col string, prefix string) string {
	rev := ""
	for _, item := range t.Columns {
		if strings.ToLower(item.Name) == col {
			continue
		}
		if len(rev) > 0 {
			rev += ", "
		}
		rev += prefix + "." + utils.CamelizeStr(item.Name, true)
	}
	return rev
}

type ColumnMetaData struct {
	Name       			string
	DBType     			string
	GoType     			string
	IsUnsigned 			bool
	IsNullable 			bool
	TableName  			string
	FormatDriveEngine  	string
}

type ColumnMetaDataList []*ColumnMetaData

var customerColumnDataType map[string]string
var customerColumnDataTypeImport map[string]string

func NewColumnMetaData(name string, isNullable bool, dataType string, isUnsigned bool, tableName string, formatDriveEngine string) *ColumnMetaData {
	columnMetaData := &ColumnMetaData{
		Name:       name,
		IsNullable: isNullable,
		DBType:     dataType,
		IsUnsigned: isUnsigned,
		TableName:  tableName,
		FormatDriveEngine: formatDriveEngine,
	}
	columnMetaData.GoType = columnMetaData.getGoType()
	return columnMetaData
}

func CustomerColumnDataType(dbColumnType string, customerType string, importStr string) {
	customerColumnDataType[dbColumnType] = customerType
	customerColumnDataTypeImport[customerType] = importStr
}

func (c ColumnMetaData) getGoType() string {
	if value, ok := customerColumnDataType[c.DBType]; ok {
		return value
	}
	switch c.DBType {
	case "boolean":
		return "bool"
	case "tinyint":
		return "int8"
	case "smallint", "year":
		return "int16"
	case "integer", "mediumint", "int":
		return "int32"
	case "bigint":
		return "int64"
	case "date", "timestamp without time zone", "timestamp with time zone", "time with time zone", "time without time zone",
		"timestamp", "datetime", "time":
		return "time.Time"
	case "bytea",
		"binary", "varbinary", "tinyblob", "blob", "mediumblob", "longblob":
		return "[]byte"
	case "text", "character", "character varying", "tsvector", "bit", "bit varying", "money", "json", "jsonb", "xml", "point", "interval", "line", "ARRAY",
		"char", "varchar", "tinytext", "mediumtext", "longtext":
		return "string"
	case "real":
		return "float32"
	case "numeric", "decimal", "double precision", "float", "double":
		return "float64"
	default:
		return "string"
	}
}

func (c ColumnMetaData) Tag() string {
	return fmt.Sprintf("`%s:\"%s\" json:\"%s,omitempty\"`", c.FormatDriveEngine, c.Name, utils.CamelizeStr(c.Name, false))
}

// TagSetDbStr
//  @Description: 设置db扩展参数
//  @receiver c
//  @param indexNums	下标字符串 如 0,1,2
//  @param dbStr		db扩展字符串
//  @return string
func (c ColumnMetaData) TagSetDbStr(indexNums string, dbStr string) string {
	indexArr := strings.Split(indexNums,",")
	currFieldName := ""
	isSetStr := false
	for _,v := range indexArr {
		indexInt,_ := strconv.Atoi(v)
		currFieldName = tableMetaData.FieldName(indexInt)
		if(currFieldName == c.Name){
			isSetStr = true
			break
		}
	}

	if isSetStr == true{
		return fmt.Sprintf("`%s:\"%s%s\" json:\"%s,omitempty\"`", c.FormatDriveEngine, c.Name, dbStr, utils.CamelizeStr(c.Name, false))
	}

	return fmt.Sprintf("`%s:\"%s\" json:\"%s,omitempty\"`", c.FormatDriveEngine, c.Name, utils.CamelizeStr(c.Name, false))
}

// TagSetJsonStr
//  @Description: 设置json扩展参数
//  @receiver c
//  @param indexNums	下标字符串 如 0,1,2
//  @param jsonStr		json扩展字符串
//  @return string
func (c ColumnMetaData) TagSetJsonStr(indexNums string, jsonStr string) string {
	indexArr := strings.Split(indexNums,",")
	currFieldName := ""
	isSetStr := false
	for _,v := range indexArr {
		indexInt,_ := strconv.Atoi(v)
		currFieldName = tableMetaData.FieldName(indexInt)
		if(currFieldName == c.Name){
			isSetStr = true
			break
		}
	}

	if isSetStr == true{
		return fmt.Sprintf("`%s:\"%s\" json:\"%s%s\"`", c.FormatDriveEngine, c.Name, utils.CamelizeStr(c.Name, false), jsonStr)
	}
	return fmt.Sprintf("`%s:\"%s\" json:\"%s,omitempty\"`", c.FormatDriveEngine, c.Name, utils.CamelizeStr(c.Name, false))
}

// TagSetDbJsonStr
//  @Description: 设置db和json扩展参数
//  @receiver c
//  @param indexNums	下标字符串 如 0,1,2
//  @param dbStr		db扩展字符串
//  @param jsonStr		json扩展字符串
//  @return string
func (c ColumnMetaData) TagSetDbJsonStr(indexNums string, dbStr string, jsonStr string) string {
	indexArr := strings.Split(indexNums,",")
	currFieldName := ""
	isSetStr := false
	for _,v := range indexArr {
		indexInt,_ := strconv.Atoi(v)
		currFieldName = tableMetaData.FieldName(indexInt)
		if(currFieldName == c.Name){
			isSetStr = true
			break
		}
	}

	if isSetStr == true{
		return fmt.Sprintf("`%s:\"%s%s\" json:\"%s%s\"`", c.FormatDriveEngine, c.Name, dbStr, utils.CamelizeStr(c.Name, false), jsonStr)
	}
	return fmt.Sprintf("`%s:\"%s\" json:\"%s,omitempty\"`", c.FormatDriveEngine, c.Name, utils.CamelizeStr(c.Name, false))
}