package sql2struct

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type DBModel struct {
	DBEngine *sql.DB
	DBInfo   *DBInfo
}

type DBInfo struct {
	DBType   string
	Host     string
	UserName string
	Password string
	Charset  string
}

func NewDBModel(info *DBInfo) *DBModel {
	return &DBModel{DBInfo: info}
}

func (m *DBModel) Connect() error {
	var err error
	s := "%s:%s@tcp(%s)/information_schema?" +
		"charset=%s&parseTime=True&loc=Local"
	dsn := fmt.Sprintf(
		s,
		m.DBInfo.UserName,
		m.DBInfo.Password,
		m.DBInfo.Host,
		m.DBInfo.Charset)
	fmt.Println("dsn", dsn)
	m.DBEngine, err = sql.Open(m.DBInfo.DBType, dsn)
	//m.DBEngine, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1)/information_schema?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		return err
	} else {
		fmt.Println("数据库连接成功", m.DBEngine)
	}
	return nil
}
