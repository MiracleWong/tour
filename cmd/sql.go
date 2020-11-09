package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// 声明7各 cmd 全局变量。用于接收外部的命令行参数
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
		fmt.Println("struct ")
	},
}

func init() {
	sqlCmd.AddCommand(sql2structCmd)
}
