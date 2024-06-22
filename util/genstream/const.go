package genstream

type ParserType string

const (
	// 输入类型解析器
	ParserTypeText      ParserType = "text"
	ParserTypeLine      ParserType = "line"
	ParserTypeRow       ParserType = "row"
	ParserTypeJson      ParserType = "json"
	ParserTypeCreateSql ParserType = "createSql"

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
