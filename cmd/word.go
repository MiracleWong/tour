package cmd

import (
	"github.com/MiracleWong/tour/internal/word"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

// 枚举类型值

const (
	MODE_UPPER                         = iota + 1 // 全部转换成大写单词
	MODE_LOWER                                    // 全部转换成小写单词
	MODE_UNDERSCORE_TO_UPPER_CAMELCASE            // 下划线单词转换成大写驼峰单词
	MODE_UNDERSCORE_TO_LOWER_CAMELCASE            // 下划线单词转换成小写驼峰单词
	MODE_CAMELCASE_TO_UNDERSCORE                  // 驼峰单词转换成下划线单词
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
		case MODE_UNDERSCORE_TO_LOWER_CAMELCASE:
			content = word.UnderScoreToLowerCamelCase(str)
		case MODE_CAMELCASE_TO_UNDERSCORE:
			content = word.CamelCaseToUnderScore(str)
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
