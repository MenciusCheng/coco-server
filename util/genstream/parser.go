package genstream

import (
	"bytes"
	"context"
	"fmt"
)

type ParserConfig struct {
	Type ParserType     `json:"type"` // 解析器类型
	Text string         `json:"text"` // 文本
	Opts []ParserOption `json:"opts"` // 选项
}

type ParserOption struct {
	Type  ParserOptionType `json:"type"`  // 选项类型
	Value string           `json:"value"` // 选项值
}

type ParserReq struct {
	Text string                 `json:"text"` // 文本
	Opts []ParserOption         `json:"opts"` // 选项
	Rows [][]string             `json:"rows"` // 行列文本
	Obj  map[string]interface{} `json:"obj"`  // obj数据
}

type ParserRes struct {
	List []ParserResData `json:"list"`
}

func (r *ParserRes) Show() string {
	buf := new(bytes.Buffer)
	for _, data := range r.List {
		buf.WriteString(fmt.Sprintf("name: %s\n", data.Name))
		buf.WriteString(fmt.Sprintf("content:\n%s\n", data.Content))
	}
	return buf.String()
}

type ParserResData struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

func NewParserFunc(parserType ParserType) ParserFunc {
	switch parserType {
	case ParserTypeText:
		return ParserFuncText
	case ParserTypeLine:
		return ParserFuncLine
	case ParserTypeRow:
		return ParserFuncRow
	case ParserTypeJson:
		return ParserFuncJson
	case ParserTypeReg:
		return ParserFuncReg
	case ParserTypeSplit:
		return ParserFuncSplit
	case ParserTypeJoin:
		return ParserFuncJoin
	case ParserTypeTemp:
		return ParserFuncTemp
	}
	return ParserFuncNone
}

type ParserFunc func(ctx context.Context, req *ParserReq, res *ParserRes) error
