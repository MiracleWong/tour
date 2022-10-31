package cmd

import (
	"github.com/MiracleWong/tour/internal/sql2struct"
	"github.com/spf13/cobra"
	"log"
)

var dbName string
var tableName string

var sqlCmd = &cobra.Command{
	Use:   "sql",
	Short: "数据库表格式转换",
	Long:  "数据库表格式转换",
	Run: func(cmd *cobra.Command, args []string) {
		dbInfo := &sql2struct.DBInfo{
			DBType:   "mysql",
			Host:     "192.168.3.10:3306",
			UserName: "root",
			Password: "123456",
			Charset:  "utf8mb4",
		}

		dbModel := sql2struct.NewDBModel(dbInfo)
		err := dbModel.Connect()
		if err != nil {
			log.Fatalf("dbModel.Connect error: %v", err)
		}
		log.Printf("输出的格式是：%v", *dbModel.DBInfo)
		dbName = "demo7"
		tableName = "operator"
		columns, err := dbModel.GetColumns(dbName, tableName)

		log.Printf("输出的格式是：%v", columns)
		if err != nil {
			log.Fatalf("dbModel.GetColumns err: %v", err)
		}

		for _, column := range columns {
			log.Printf("输出的格式是：%v", column.ColumnName)
		}

	},
}

func init() {
	sqlCmd.Flags().StringVarP(&str, "str", "s", "", "请输入具体的名称")
}
