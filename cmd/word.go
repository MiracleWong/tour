package cmd

import (
	"github.com/MiracleWong/tour/internal/word"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

// 枚举类型值

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
	"该子命令支持各种单词的格式转换，模式如下：",
	"1：全部单词转为大写",
}, "\n")

// 用于放置 单词格式转换 的子命令 word

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
			content = word.UnderScoreToUpperCamelCase(str)
		default:
			log.Fatal("暂不支持现在的单词转换模式")
		}
		log.Printf("输出结果为：%s", content)
	},
}

func init() {
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "请输入单词内容")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "请输入单词内容")
}
