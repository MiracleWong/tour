package word

import (
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
