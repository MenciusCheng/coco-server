package genstream

import "strings"

type ParserType string

const (
	// 输入类型解析器
	ParserTypeText         ParserType = "text"
	ParserTypeLine         ParserType = "line"
	ParserTypeRow          ParserType = "row"
	ParserTypeJson         ParserType = "json"
	ParserTypeCreateSql    ParserType = "createSql"
	ParserTypeProtoService ParserType = "protoService"

	// 转换类型解析器
	ParserTypeReg   ParserType = "reg"
	ParserTypeSplit ParserType = "split"
	ParserTypeJoin  ParserType = "join"

	// 输出类型解析器
	ParserTypeTemp ParserType = "temp"
)

type ParserOptionType string

const (
	ParserOptionTypeName    ParserOptionType = "name"    // 名称
	ParserOptionTypeSep     ParserOptionType = "sep"     // 分隔符
	ParserOptionTypeReplace ParserOptionType = "replace" // 替换文本
	ParserOptionTypeMap     ParserOptionType = "map"     // 字典
)

func GetOptionName(opts []ParserOption) string {
	for _, opt := range opts {
		if opt.Type == ParserOptionTypeName {
			return opt.Value
		}
	}
	return ""
}

func GetOptionSep(opts []ParserOption, defaultSep string) string {
	for _, opt := range opts {
		if opt.Type == ParserOptionTypeSep {
			return opt.Value
		}
	}
	return defaultSep
}

func GetOptionReplace(opts []ParserOption) string {
	for _, opt := range opts {
		if opt.Type == ParserOptionTypeReplace {
			return opt.Value
		}
	}
	return ""
}

func GetOptionMap(opts []ParserOption) map[string]string {
	m := make(map[string]string)
	for _, opt := range opts {
		if opt.Type == ParserOptionTypeMap {
			for _, kv := range strings.Split(opt.Value, "\n") {
				splitKV := strings.Split(kv, ":")
				if len(splitKV) != 2 {
					continue
				}
				m[strings.TrimSpace(splitKV[0])] = strings.TrimSpace(splitKV[1])
			}
		}
	}
	return m
}
