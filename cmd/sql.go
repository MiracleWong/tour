package cmd

import (
	"github.com/MiracleWong/tour/internal/sql2struct"
	"github.com/spf13/cobra"
	"log"
)

var username string
var password string
var host string
var charset string
var dbType string
var dbName string
var tableName string

var sqlCmd = &cobra.Command{
	Use:   "sql",
	Short: "sql 转换和处理",
	Long:  "sql 转换和处理",
	Run:   func(cmd *cobra.Command, args []string) {},
}
var sql2structCmd = &cobra.Command{
	Use:   "struct",
	Short: "sql 转换",
	Long:  "sql 转换",
	Run: func(cmd *cobra.Command, args []string) {
		//dbInfo := &sql2struct.DBInfo{
		//	DBType:   "mysql",
		//	Host:     "192.168.3.10:3306",
		//	UserName: "root",
		//	Password: "123456",
		//	Charset:  "utf8mb4",
		//}
		dbInfo := &sql2struct.DBInfo{
			DBType:   dbType,
			Host:     host,
			UserName: username,
			Password: password,
			Charset:  charset,
		}
		dbModel := sql2struct.NewDBModel(dbInfo)
		err := dbModel.Connect()
		if err != nil {
			log.Fatalf("dbModel.Connect error: %v", err)
		}
		log.Printf("输出的格式是：%v", *dbModel.DBInfo)
		//dbName = "demo7"
		//tableName = "operator"
		columns, err := dbModel.GetColumns(dbName, tableName)

		log.Printf("columns 输出的格式是：%v", columns)
		if err != nil {
			log.Fatalf("dbModel.GetColumns err: %v", err)
		}

		for _, column := range columns {
			log.Printf("column 输出的格式是：%v", column.ColumnName)
		}

	},
}

// go run .\main.go sql struct --username "root" --password "123456" --db "demo7" --table "operator"
// go run .\main.go sql struct --username "root" --password "123456" --db "demo" --table "blog_tag"

func init() {
	sqlCmd.AddCommand(sql2structCmd)
	sql2structCmd.Flags().StringVarP(&username, "username", "", "", "请输入数据库的账号")
	sql2structCmd.Flags().StringVarP(&password, "password", "", "", "请输入数据库的密码")
	sql2structCmd.Flags().StringVarP(&host, "host", "", "192.168.3.10:3306", "请输入数据库的HOST")
	sql2structCmd.Flags().StringVarP(&charset, "charset", "", "utf8mb4", "请输入数据库的编码")
	sql2structCmd.Flags().StringVarP(&dbType, "dbType", "", "mysql", "请输入数据库实例类型")
	sql2structCmd.Flags().StringVarP(&dbName, "db", "", "", "请输入数据库名称")
	sql2structCmd.Flags().StringVarP(&tableName, "table", "", "", "请输入表名称")
}
