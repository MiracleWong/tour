package word

import (
	"fmt"
	"strings"
	"unicode"
)

// 全部转换成大写单词

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// 全部转换成小写单词

func ToLower(s string) string {
	return strings.ToLower(s)
}

// 下划线单词转换成大写驼峰单词

func UnderScoreToUpperCamelCase(s string) string {
	s = strings.Replace(s, "_", " ", -1)
	//fmt.Println(s)
	s = strings.Title(s)
	//fmt.Println(s)
	return strings.Replace(s, " ", "", -1)
}

// 下划线单词转换成小写驼峰单词

func UnderScoreToLowerCamelCase(s string) string {
	s = UnderScoreToUpperCamelCase(s)
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

// 驼峰单词转换成下划线单词

func CamelCaseToUnderScore(s string) string {
	var output []rune
	for i, r := range s {
		fmt.Println(i, r)
		if i == 0 {
			output = append(output, unicode.ToLower(r))
		} else {
			if unicode.IsUpper(r) {
				output = append(output, '_')
			}

			output = append(output, unicode.ToLower(r))
		}
	}
	return string(output)
}
