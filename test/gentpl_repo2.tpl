package repo

{{$packageName := param "packageName"}}
{{$packageNameFirstUpper := CamelizeStr $packageName true}}

{{$unPreTableName := RemovePrefix .Name "mt_"}}
{{$unPreTableNameUpper := CamelizeStr $unPreTableName true}}
{{$unPreTableNameLower := CamelizeStr $unPreTableName false}}

{{$firstChar := FirstCharacter $unPreTableName}}
{{$camelizeStructName := CamelizeStr .Name false}}

{{$structName := CamelizeStr .Name true}}

import (
	"appgo/cmd/gin/entity/do"
	"appgo/pkg/xcontent"
	"github.com/jinzhu/gorm"
)

type {{$unPreTableNameUpper}}Repo struct {
	xcontent.XContext
}

func New{{$unPreTableNameUpper}}Repo() *{{$unPreTableNameUpper}}Repo {
	return &{{$unPreTableNameUpper}}Repo{}
}

// Create 创建
func ({{$firstChar}} *{{$unPreTableNameUpper}}Repo) Create(data interface{}) error {
	db := {{$firstChar}}.NewDb().Create(data)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

// DelById 通过id删除
func ({{$firstChar}} *{{$unPreTableNameUpper}}Repo) DelById(id int32) (int64, error) {
	db := {{$firstChar}}.NewDb().Delete(&do.{{$unPreTableNameUpper}}{}, "id = ?", id)
	if db.Error != nil && db.Error != gorm.ErrRecordNotFound {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

// DelByIds 通过多id删除
func ({{$firstChar}} *{{$unPreTableNameUpper}}Repo) DelByIds(ids []int32) (int64, error) {
	db := {{$firstChar}}.NewDb().Delete(&do.{{$unPreTableNameUpper}}{}, "id in (?)", ids)
	if db.Error != nil && db.Error != gorm.ErrRecordNotFound {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

// Del 通过多条件删除
func ({{$firstChar}} *{{$unPreTableNameUpper}}Repo) Del(where *do.Where{{$unPreTableNameUpper}}) (int64, error) {
	db := {{$firstChar}}.NewDb().Model(&do.{{$unPreTableNameUpper}}{})

	{{- $first := true }}
    {{- range $index, $column := .Columns }}
        {{- if not $first }}
            if len(where.{{CamelizeStr .Name true }}) > 0 {
                    db = db.Where("{{.Name}} = ?", where.{{CamelizeStr .Name true }})
            }
        {{- end }}
        {{- $first = false }}
    {{- end }}

	result := db.Delete(&do.{{$unPreTableNameUpper}}{})
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

// UpdateById 通过id更新
func ({{$firstChar}} *{{$unPreTableNameUpper}}Repo) UpdateById(id int32, updateMap map[string]interface{}) (int64, error) {
	db := {{$firstChar}}.NewDb().Model(&do.{{$unPreTableNameUpper}}{}).Where("id = ?", id).Updates(updateMap)
	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

// Update 通过多条件更新
func ({{$firstChar}} *{{$unPreTableNameUpper}}Repo) Update(where *do.Where{{$unPreTableNameUpper}}, updateMap map[string]interface{}) (int64, error) {
	db := {{$firstChar}}.NewDb().Model(&do.{{$unPreTableNameUpper}}{})

	{{- $first := true }}
    {{- range $index, $column := .Columns }}
        {{- if not $first }}
            if len(where.{{CamelizeStr .Name true }}) > 0 {
                    db = db.Where("{{.Name}} = ?", where.{{CamelizeStr .Name true }})
            }
        {{- end }}
        {{- $first = false }}
    {{- end }}

	db = db.Updates(updateMap)
	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

// GetById 通过id查找
func ({{$firstChar}} *{{$unPreTableNameUpper}}Repo) GetById(id int32) (*do.{{$unPreTableNameUpper}}, error) {
	ety := &do.{{$unPreTableNameUpper}}{}
	err := {{$firstChar}}.NewDb().Model(&do.{{$unPreTableNameUpper}}{}).Where("id = ?", id).First(&ety).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return ety, nil
}

// Find 通过条件查找一条数据
func ({{$firstChar}} *{{$unPreTableNameUpper}}Repo) Find(where *do.Where{{$unPreTableNameUpper}}) (*do.{{$unPreTableNameUpper}}, error) {
	ety := &do.{{$unPreTableNameUpper}}{}
	db := {{$firstChar}}.NewDb().Model(&do.{{$unPreTableNameUpper}}{})

	{{- $first := true }}
    {{- range $index, $column := .Columns }}
        {{- if not $first }}
            if len(where.{{CamelizeStr .Name true }}) > 0 {
                    db = db.Where("{{.Name}} = ?", where.{{CamelizeStr .Name true }})
            }
        {{- end }}
        {{- $first = false }}
    {{- end }}

	err := db.First(&ety).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return ety, nil
}

// SearchList 搜索列表
func ({{$firstChar}} *{{$unPreTableNameUpper}}Repo) SearchList(where *do.Where{{$unPreTableNameUpper}}) ([]do.{{$unPreTableNameUpper}}, error) {
	var items []do.{{$unPreTableNameUpper}}
	db := {{$firstChar}}.NewDb().Model(&do.{{$unPreTableNameUpper}}{})


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
func ({{$firstChar}} *{{$unPreTableNameUpper}}Repo) SearchPage(where *do.Where{{$unPreTableNameUpper}}, page, limit int32) ([]do.{{$unPreTableNameUpper}}, int64, error) {
	offset := (page - 1) * limit

    var count int64
	var items []do.{{$unPreTableNameUpper}}
	db := {{$firstChar}}.NewDb().Model(&do.{{$unPreTableNameUpper}}{})


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

// QueryOne 通过sql查询单条数据
func ({{$firstChar}} *{{$unPreTableNameUpper}}Repo) QueryOne(sql string) (*do.{{$unPreTableNameUpper}}, error) {
	ety := &do.{{$unPreTableNameUpper}}{}
	err := {{$firstChar}}.NewDb().Raw(sql).Scan(&ety).Error // 等价于 SQL 中添加了 LIMIT 1
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return ety, nil
}

// QueryMore 通过sql查询多条数据
func ({{$firstChar}} *{{$unPreTableNameUpper}}Repo) QueryMore(sql string) ([]do.{{$unPreTableNameUpper}}, error) {
	var items []do.{{$unPreTableNameUpper}}
	err := {{$firstChar}}.NewDb().Raw(sql).Scan(&items).Error // 会返回所有符合条件的记录
	if err != nil {
		return nil, err
	}
	return items, nil
}

// Exec 执行不返回结果的 SQL
func ({{$firstChar}} *{{$unPreTableNameUpper}}Repo) Exec(sql string) error {
	err := {{$firstChar}}.NewDb().Exec(sql).Error
	return err
}