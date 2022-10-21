package word

import (
	"strings"
	"unicode"
)

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func ToLower(s string) string {
	return strings.ToLower(s)
}

func UnderscoreToUpperCamelCase(s string) string { // miracle_wong
	s = strings.Replace(s, "_", " ", -1)   // miracle wong
	s = strings.Title(s)                   // Miracle Wong
	return strings.Replace(s, " ", "", -1) // MiracleWong
}

func UnderscoreToLowerCamelCase(s string) string { // miracle_wong
	s = UnderscoreToUpperCamelCase(s)                  // MiracleWong
	return string(unicode.ToLower(rune(s[0]))) + s[1:] // miracleWong
}

func CamelCaseToUnderscore(s string) string { // MiracleWong
	var output []rune
	for i, r := range s {
		if i == 0 {
			output = append(output, unicode.ToLower(r))
			continue
		}
		if unicode.IsUpper(r) {
			output = append(output, '_')
		}
		output = append(output, unicode.ToLower(r))
	}
	return string(output)
}
