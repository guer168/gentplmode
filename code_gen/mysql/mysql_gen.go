package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/guer168/gentplmode/db_meta_data"
	"github.com/guer168/gentplmode/utils"
)

const (
	//获取所有表
	tableNamesSql          = `select table_name from information_schema.tables where table_schema = ? and table_type = 'base table';`
	//获取指定表
	specifiedTableNamesSql = `select table_name from information_schema.tables where table_schema = ? and table_name in ('%s') and table_type = 'base table';`
	//获取字段信息
	tableColumnsSql        = `select column_name,
is_nullable, if(column_type = 'tinyint(1)', 'boolean', data_type),
column_type like '%unsigned%', column_comment
from information_schema.columns
where table_schema = ? and  table_name = ?
order by ordinal_position;
`
)

type Gen struct {
	db     				*sql.DB
	dbName 				string
	formatDriveEngine	string
}

func (m *Gen) ConnectionDB(dsn string) error {
	fmt.Println("MySQL Connecting dsn : " + dsn)
	dbName, err := utils.GetDbNameFromDSN(dsn)
	if err != nil {
		return err
	}
	m.dbName = dbName
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	m.db = db
	m.formatDriveEngine = "db"
	return nil
}

func (m *Gen) AllTableData() (db_meta_data.TableMetaDataList, error) {
	rows, err := m.db.Query(tableNamesSql, m.dbName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	rev := db_meta_data.TableMetaDataList{}
	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			return nil, err
		}
		tableColumnsInfo, err := m.GetTableColumns(tableName)
		if err != nil {
			return nil, err
		}
		rev = append(rev, &db_meta_data.TableMetaData{Name: tableName, Columns: tableColumnsInfo})
	}

	return rev, rows.Err()
}

func (m *Gen) SpecifiedTables(tableNameList []string) (db_meta_data.TableMetaDataList, error) {
	if len(tableNameList) == 0 {
		return nil, errors.New("tableNameList is empty")
	}
	sqlStr := fmt.Sprintf(specifiedTableNamesSql, strings.Join(tableNameList, "','"))
	rows, err := m.db.Query(sqlStr, m.dbName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	rev := db_meta_data.TableMetaDataList{}
	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			return nil, err
		}
		tableColumnsInfo, err := m.GetTableColumns(tableName)
		if err != nil {
			return nil, err
		}
		rev = append(rev, &db_meta_data.TableMetaData{Name: tableName, Columns: tableColumnsInfo})
	}

	return rev, rows.Err()
}

func (m *Gen) GetTableColumns(tableName string) (db_meta_data.ColumnMetaDataList, error) {
	rows, err := m.db.Query(tableColumnsSql, m.dbName, tableName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	rev := db_meta_data.ColumnMetaDataList{}
	for rows.Next() {
		var name, isNullable, dataType, dataComment string
		var isUnsigned bool
		if err := rows.Scan(&name, &isNullable, &dataType, &isUnsigned, &dataComment); err != nil {
			return nil, err
		}
		//fmt.Println(name, dataType, dataComment)
		rev = append(rev, db_meta_data.NewColumnMetaData(name,
			strings.ToLower(isNullable) == "yes", dataType, isUnsigned, tableName, m.formatDriveEngine, dataComment))
	}
	return rev, rows.Err()
}

func (m *Gen) SetFormatDriveEngine(formatDriveEngine string) error {
	if len(formatDriveEngine)>0 {
		m.formatDriveEngine = formatDriveEngine
	}else{
		m.formatDriveEngine = "db"
	}
	return nil
}
