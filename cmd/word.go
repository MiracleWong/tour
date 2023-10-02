package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// 用于放置 单词格式转换 的子命令 word

var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "单词格式转换",
	Long:  "支持多种格式的单词格式转换",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Word Cmd Body")
	},
}

func init() {

}
