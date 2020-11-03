package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var sqlCmd = &cobra.Command{
	Use:   "sql",
	Short: "sql 转换和处理",
	Long:  "sql 转换和处理",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Hello World")
	},
}

func init() {}
