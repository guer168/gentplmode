package repo

{{$packageName := param "packageName"}}
{{$packageNameFirstUpper := CamelizeStr $packageName true}}

{{$unPreTableName := RemovePrefix .Name "erp_"}}
{{$unPreTableNameUpper := CamelizeStr $unPreTableName true}}
{{$unPreTableNameLower := CamelizeStr $unPreTableName false}}

{{$firstChar := FirstCharacter $unPreTableName}}
{{$camelizeStructName := CamelizeStr .Name false}}

{{$structName := CamelizeStr .Name true}}

import (
	"erp-collect/entity/do"
	"erp-collect/util/xcontent"
	"erp-collect/util/xlog"
	"fmt"
	"github.com/jinzhu/gorm"
)

type {{$unPreTableNameLower}} struct {
	xcontent.Xcontext
	xLog xlog.Logger
}


func New{{$unPreTableNameUpper}}Repo() *{{$unPreTableNameLower}} {
	return &{{$unPreTableNameLower}}{}
}

// Create 创建
func ({{$firstChar}} *{{$unPreTableNameLower}}) Create(data interface{}) error {
	db := {{$firstChar}}.InitMySql().Create(data)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

// DelById 通过id删除
func ({{$firstChar}} *{{$unPreTableNameLower}}) DelById(id int32) (int64, error) {
	db := {{$firstChar}}.InitMySql().Delete(&do.{{$unPreTableNameUpper}}{}, "id = ?", id)
	if db.Error != nil && db.Error != gorm.ErrRecordNotFound {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

// DelByIds 通过多id删除
func ({{$firstChar}} *{{$unPreTableNameLower}}) DelByIds(ids []int32) (int64, error) {
	db := {{$firstChar}}.InitMySql().Delete(&do.{{$unPreTableNameUpper}}{}, "id in (?)", ids)
	if db.Error != nil && db.Error != gorm.ErrRecordNotFound {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

// UpdateById 通过id更新
func ({{$firstChar}} *{{$unPreTableNameLower}}) UpdateById(id int32, updateMap map[string]interface{}) (int64, error) {
	db := {{$firstChar}}.InitMySql().Model(&do.{{$unPreTableNameUpper}}{}).Where("id = ?", id).Updates(updateMap)
	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

// GetById 通过id查找
func ({{$firstChar}} *{{$unPreTableNameLower}}) GetById(id int32) (*do.XpathGroupSys, error) {
	ety := &do.{{$unPreTableNameUpper}}{}
	err := {{$firstChar}}.InitMySql().Model(&do.LoginCookies{}).Where("id = ?", id).First(&ety).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return ety, nil
}

// SearchList 搜索列表
func ({{$firstChar}} *{{$unPreTableNameLower}}) SearchList(where *do.Where{{$unPreTableNameUpper}}) ([]do.{{$unPreTableNameUpper}}, error) {
	var items []do.{{$unPreTableNameUpper}}
	db := {{$firstChar}}.InitMySql().Model(&do.{{$unPreTableNameUpper}}{})


    {{- $first := true }}
    {{- range $index, $column := .Columns }}
        {{- if not $first }}
            if len(where.{{CamelizeStr .Name true }}) > 0 {
            		db = db.Where("{{.Name}} = ?", where.{{CamelizeStr .Name true }})
            }
        {{- end }}
        {{- $first = false }}
    {{- end }}

	res := db.Order("id DESC").Find(&items)
	if res.Error != nil {
		return nil, res.Error
	}
	return items, nil
}

// SearchPage 搜索分页
func ({{$firstChar}} *{{$unPreTableNameLower}}) SearchPage(where *do.Where{{$unPreTableNameUpper}}, page, limit int32) ([]do.{{$unPreTableNameUpper}}, int64, error) {
	offset := (page - 1) * limit

    var count int64
	var items []do.{{$unPreTableNameUpper}}
	db := {{$firstChar}}.InitMySql().Model(&do.{{$unPreTableNameUpper}}{})


	{{- $first := true }}
        {{- range $index, $column := .Columns }}
        {{- if not $first }}
            if len(where.{{CamelizeStr .Name true }}) > 0 {
                db = db.Where("{{.Name}} = ?", where.{{CamelizeStr .Name true }})
            }
        {{- end }}
        {{- $first = false }}
    {{- end }}

	dbCount := db
    dbList := db

    dbCount.Count(count)
    res := dbList.Order("id DESC").Limit(limit).Offset(offset).Find(&items)
	if res.Error != nil {
		return nil, 0, res.Error
	}
	return items, int64(len(items)), nil
}