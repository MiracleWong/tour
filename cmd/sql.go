package cmd

import (
	"fmt"
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
	Short: "数据库sql 转换与处理",
	Long:  "数据库sql 转换与处理",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello Parent Cmd")

	},
}

var sql2structCmd = &cobra.Command{
	Use:   "struct",
	Short: "sql 转换",
	Long:  "sql 转换",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello Parent Cmd")
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
			log.Fatalf("dbModel.Connect err: %v", err)
		} else {
			fmt.Println("数据库连接成功")
		}

	},
}

func init() {
	fmt.Println("Hello Init sql.go")
	sqlCmd.AddCommand(sql2structCmd)
	sql2structCmd.Flags().StringVarP(&username, "username", "", "root", "请输入数据库的账号")
	sql2structCmd.Flags().StringVarP(&password, "password", "", "123456", "请输入数据库的密码")
	sql2structCmd.Flags().StringVarP(&host, "host", "", "127.0.0.1", "请输入数据库的Host")
	sql2structCmd.Flags().StringVarP(&charset, "charset", "", "utf8mb4", "请输入数据库的编码")
	sql2structCmd.Flags().StringVarP(&dbType, "dbType", "", "mysql", "请输入数据库实例类型")
	sql2structCmd.Flags().StringVarP(&dbName, "dbName", "", "test", "请输入数据库名称")
	sql2structCmd.Flags().StringVarP(&tableName, "tableName", "", "test", "请输入表名称")
}
