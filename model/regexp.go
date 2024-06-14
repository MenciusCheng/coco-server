package model

type RegexpFindReq struct {
	Expr string `json:"expr"` // 正则表达式
	S    string `json:"s"`    // 源文本
}

type RegexpFindRes struct {
	MatchString                bool       `json:"matchString"` // 是否匹配成功
	FindString                 string     `json:"findString"`
	FindStringIndex            []int      `json:"findStringIndex"`
	FindStringSubmatch         []string   `json:"findStringSubmatch"`
	FindStringSubmatchIndex    []int      `json:"findStringSubmatchIndex"`
	FindAllString              []string   `json:"findAllString"`
	FindAllStringIndex         [][]int    `json:"findAllStringIndex"`
	FindAllStringSubmatch      [][]string `json:"findAllStringSubmatch"`
	FindAllStringSubmatchIndex [][]int    `json:"findAllStringSubmatchIndex"`
}

type RegexpReplaceReq struct {
	Expr string `json:"expr"` // 正则表达式
	S    string `json:"s"`    // 源文本
	Repl string `json:"repl"` // 替换文本
}

type RegexpReplaceRes struct {
	MatchString bool   `json:"matchString"` // 是否匹配成功
	Des         string `json:"des"`         // 替换结果文本
}
