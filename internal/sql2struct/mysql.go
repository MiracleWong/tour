package sql2struct

import (
	"database/sql"
	"errors"
	"fmt"
	//导入驱动
	_ "github.com/go-sql-driver/mysql"
)

// 数据库连接的核心对象
type DBModel struct {
	DBEngine *sql.DB
	DBInfo   *DBInfo
}

// 结构体DBInfo：用于存储连接 MySQL 的基本信息
type DBInfo struct {
	DBType   string
	Host     string
	UserName string
	Password string
	Charset  string
}

// TableColumn 是用于存储 Columns 表中需要的字段
type TableColumn struct {
	ColumnName    string
	DataType      string
	IsNullable    string
	ColumnKey     string
	ColumnType    string
	ColumnComment string
}

// 表字段的类型映射，使用枚举，用 ap 作映射获取
var DBTypeToStructType = map[string]string{
	"int":        "int32",
	"tinyint":    "int8",
	"smallint":   "int",
	"mediumint":  "int64",
	"bigint":     "int64",
	"bit":        "int",
	"bool":       "bool",
	"enum":       "string",
	"set":        "string",
	"varchar":    "string",
	"char":       "string",
	"tinytext":   "string",
	"mediumtext": "string",
	"text":       "string",
	"longtext":   "string",
	"blob":       "string",
	"tinyblob":   "string",
	"mediumblob": "string",
	"longblob":   "string",
	"date":       "time.Time",
	"datetime":   "time.Time",
	"timestamp":  "time.Time",
	"time":       "time.Time",
	"float":      "float64",
	"double":     "float64",
}

func NewDBModel(dbInfo *DBInfo) *DBModel {
	return &DBModel{DBInfo: dbInfo}
}

// Go 中，连接数据库的方法
func (m *DBModel) Connect() error {
	var err error
	s := "%s:%s@tcp(%s)/information_schema?" +
		"charset=%s&parseTime=True&loc=Local"
	dsn := fmt.Sprintf(
		s,
		m.DBInfo.UserName,
		m.DBInfo.Password,
		m.DBInfo.Host,
		m.DBInfo.Charset,
	)
	// 使用的是标准库的 Open 方法
	m.DBEngine, err = sql.Open(m.DBInfo.DBType, dsn)
	if err != nil {
		return err
	}
	return nil
}

// 我自己的数据库信息
//SELECT COLUMN_NAME, DATA_TYPE, COLUMN_KEY, IS_NULLABLE, COLUMN_TYPE, COLUMN_COMMENT FROM COLUMNS WHERE TABLE_SCHEMA = 'nba' AND TABLE_NAME = 'player';
//+-------------+-----------+------------+-------------+--------------+----------------+
//| COLUMN_NAME | DATA_TYPE | COLUMN_KEY | IS_NULLABLE | COLUMN_TYPE  | COLUMN_COMMENT |
//+-------------+-----------+------------+-------------+--------------+----------------+
//| player_id   | int       | PRI        | NO          | int(11)      | 球员ID         |
//| team_id     | int       |            | NO          | int(11)      | 球队ID         |
//| player_name | varchar   | UNI        | NO          | varchar(255) | 球员姓名       |
//| height      | float     |            | YES         | float(3,2)   | 球员身高       |
//+-------------+-----------+------------+-------------+--------------+----------------+
//4 rows in set (0.00 sec)

//SELECT COLUMN_NAME, DATA_TYPE, COLUMN_KEY, IS_NULLABLE, COLUMN_TYPE, COLUMN_COMMENT FROM COLUMNS WHERE TABLE_SCHEMA = 'nba' AND TABLE_NAME = 'player_score';
//+------------------+-----------+------------+-------------+-------------+-----------------------------------+
//| COLUMN_NAME      | DATA_TYPE | COLUMN_KEY | IS_NULLABLE | COLUMN_TYPE | COLUMN_COMMENT                    |
//+------------------+-----------+------------+-------------+-------------+-----------------------------------+
//| game_id          | int       |            | NO          | int(11)     | 比赛ID                            |
//| player_id        | int       |            | NO          | int(11)     | 球员ID                            |
//| is_first         | tinyint   |            | NO          | tinyint(1)  | 是否首发                          |
//| playing_time     | int       |            | NO          | int(11)     | 该球员本次比赛出场时间            |
//| rebound          | int       |            | NO          | int(11)     | 篮板球                            |
//| rebound_o        | int       |            | NO          | int(11)     | 前场篮板                          |
//| rebound_d        | int       |            | NO          | int(11)     | 后场篮板                          |
//| assist           | int       |            | NO          | int(11)     | 助攻                              |
//| score            | int       |            | NO          | int(11)     | 比分                              |
//| steal            | int       |            | NO          | int(11)     | 抢断                              |
//| blockshot        | int       |            | NO          | int(11)     | 盖帽                              |
//| fault            | int       |            | NO          | int(11)     | 失误                              |
//| foul             | int       |            | NO          | int(11)     | 犯规                              |
//| shoot_attempts   | int       |            | NO          | int(11)     | 总出手                            |
//| shoot_hits       | int       |            | NO          | int(11)     | 命中                              |
//| shoot_3_attempts | int       |            | NO          | int(11)     | 3分出手                           |
//| shoot_3_hits     | int       |            | NO          | int(11)     | 3分命中                           |
//| shoot_p_attempts | int       |            | NO          | int(11)     | 罚球出手                          |
//| shoot_p_hits     | int       |            | NO          | int(11)     | 罚球命中                          |
//+------------------+-----------+------------+-------------+-------------+-----------------------------------+
//19 rows in set (0.00 sec)

// 针对columns 表进行查询和数据组装
func (m *DBModel) GetColumns(dbName, tableName string) ([]*TableColumn, error) {
	query := "SELECT COLUMN_NAME, DATA_TYPE, COLUMN_KEY, " +
		"IS_NULLABLE, COLUMN_TYPE, COLUMN_COMMENT " +
		"FROM COLUMNS WHERE TABLE_SCHEMA = ? AND TABLE_NAME = ? "
	rows, err := m.DBEngine.Query(query, dbName, tableName)
	if err != nil {
		return nil, err
	}
	if rows == nil {
		return nil, errors.New("没有数据")
	}
	defer rows.Close()
	var columns []*TableColumn

	for rows.Next() {
		var column TableColumn
		err := rows.Scan(&column.ColumnName, &column.DataType, &column.ColumnKey,
			&column.IsNullable, &column.ColumnType, &column.ColumnComment)
		if err != nil {
			return nil, err
		}
		columns = append(columns, &column)
	}
	return columns, nil
}
