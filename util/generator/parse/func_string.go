package parse

import (
	"bytes"
	"strings"
	"unicode"
)

func init() {
	funcMap["sToCamel"] = FuncStringObj.ToCamel
	funcMap["sToLCamel"] = FuncStringObj.ToLCamel
	funcMap["sToUCamel"] = FuncStringObj.ToUCamel
	funcMap["sToSnake"] = FuncStringObj.ToSnake
	funcMap["sToUp0"] = FuncStringObj.ToUp0
	funcMap["sToLow0"] = FuncStringObj.ToLow0
}

type FuncString struct{}

var FuncStringObj = &FuncString{}

// ToCamel 将字符串转换为驼峰命名
func (f *FuncString) ToCamel(s string) string {
	// 分割字符串
	words := strings.FieldsFunc(s, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})

	b := bytes.Buffer{}
	for i, word := range words {
		if i == 0 {
			// 第一个单词不变
			b.WriteString(word)
		} else {
			// 后续单词首字母大写
			b.WriteString(f.ToUp0(word))
		}
	}
	return b.String()
}

// ToLCamel 将字符串转换为驼峰命名，首字母小写
func (f *FuncString) ToLCamel(s string) string {
	// 分割字符串
	words := strings.FieldsFunc(s, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})

	b := bytes.Buffer{}
	for i, word := range words {
		if i == 0 {
			// 第一个单词小写
			b.WriteString(f.ToLow0(word))
		} else {
			// 后续单词首字母大写
			b.WriteString(f.ToUp0(word))
		}
	}
	return b.String()
}

// ToUCamel 将字符串转换为驼峰命名，首字母大写
func (f *FuncString) ToUCamel(s string) string {
	// 分割字符串
	words := strings.FieldsFunc(s, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})

	b := bytes.Buffer{}
	for _, word := range words {
		// 单词首字母大写
		b.WriteString(f.ToUp0(word))
	}
	return b.String()
}

// ToSnake 将字符串转换为蛇形命名
func (f *FuncString) ToSnake(s string) string {
	var result []rune
	for i, r := range s {
		if unicode.IsUpper(r) {
			if i > 0 {
				result = append(result, '_')
			}
			result = append(result, unicode.ToLower(r))
		} else if unicode.IsSpace(r) || r == '-' || r == '_' {
			result = append(result, '_')
		} else {
			result = append(result, r)
		}
	}
	return string(result)
}

func (f *FuncString) ToUp0(s string) string {
	if len(s) > 0 {
		b := bytes.Buffer{}
		b.WriteString(strings.ToUpper(s[:1]))
		b.WriteString(s[1:])
		return b.String()
	}
	return s
}

func (f *FuncString) ToLow0(s string) string {
	if len(s) > 0 {
		b := bytes.Buffer{}
		b.WriteString(strings.ToLower(s[:1]))
		b.WriteString(s[1:])
		return b.String()
	}
	return s
}
