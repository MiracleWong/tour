package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var demoCmd = &cobra.Command{
	Use:   "demo",
	Short: "测试命令",
	Long:  "描述：测试命令",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Hello World")
	},
}

func init() {}
