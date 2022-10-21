package cmd

import (
	"github.com/MiracleWong/tour/internal/word"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

const (
	MODE_UPPER = iota + 1
	MODE_LOWER
	MODE_UNDERSCORE_TO_UPPER_CAMELCASE
	MODE_UNDERSCORE_TO_LOWER_CAMELCASE
	MODE_CAMELCASE_TO_UNDERSCORE
)

var str string
var mode int8

var desc = strings.Join([]string{
	"1: 全部转换单词为大写",
}, "\n")

var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "单词格式转换",
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case MODE_UPPER:
			content = word.ToUpper(str)
		case MODE_LOWER:
			content = word.ToLower(str)
		case MODE_UNDERSCORE_TO_UPPER_CAMELCASE:
			content = word.UnderscoreToUpperCamelCase(str)
		case MODE_UNDERSCORE_TO_LOWER_CAMELCASE:
			content = word.UnderscoreToLowerCamelCase(str)
		case MODE_CAMELCASE_TO_UNDERSCORE:
			content = word.CamelCaseToUnderscore(str)
		default:
			log.Printf("暂不支持转换该模式，请通过 help word 查看帮助文档")
		}
		log.Printf("输出的格式是：%v", content)
	},
}

func init() {
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "请输入具体的单词")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "请输入单词转换的格式")
}
