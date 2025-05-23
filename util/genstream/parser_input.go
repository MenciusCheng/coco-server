package genstream

import (
	"coco-server/util/generator"
	"coco-server/util/log"
	"context"
	"encoding/json"
	"go.uber.org/zap"
	"regexp"
	"strconv"
	"strings"
)

// 输入类型解析器

func ParserFuncNone(ctx context.Context, req *ParserReq, res *ParserRes) error {
	log.Warn(ctx, "ParserFuncNone called", zap.Any("req", req))
	return nil
}

func ParserFuncText(ctx context.Context, req *ParserReq, res *ParserRes) error {
	rows := [][]string{
		{req.Text},
	}

	req.Rows = rows
	return nil
}

func ParserFuncLine(ctx context.Context, req *ParserReq, res *ParserRes) error {
	lines := strings.Split(req.Text, "\n")
	rows := make([][]string, 0)
	for _, line := range lines {
		rows = append(rows, []string{line})
	}

	req.Rows = rows
	return nil
}

func ParserFuncRow(ctx context.Context, req *ParserReq, res *ParserRes) error {
	lines := strings.Split(req.Text, "\n")
	rows := make([][]string, 0)
	sep := GetOptionSep(req.Opts, "\t")
	for _, line := range lines {
		// 清洗
		lineData := strings.TrimSpace(line)
		if len(lineData) == 0 {
			continue
		}
		cols := []string{line}
		cols = append(cols, strings.Split(lineData, sep)...)
		rows = append(rows, cols)
	}

	req.Rows = rows
	return nil
}

func ParserFuncJson(ctx context.Context, req *ParserReq, res *ParserRes) error {
	obj := make(map[string]interface{})
	err := json.Unmarshal([]byte(req.Text), &obj)
	if err != nil {
		return err
	}

	req.Obj = obj
	return nil
}

// 建表语句
func ParserFuncCreateSql(ctx context.Context, req *ParserReq, res *ParserRes) error {
	obj := generator.ParserSQL2(req.Text, req.OptMap)

	req.Obj = obj
	return nil
}

// proto服务
func ParserFuncProtoService(ctx context.Context, req *ParserReq, res *ParserRes) error {
	services := parseProtoService(req.Text)
	obj := make(map[string]interface{})
	if len(services) > 0 {
		serviceBytes, err := json.Marshal(services[0])
		if err != nil {
			return err
		}
		err = json.Unmarshal(serviceBytes, &obj)
		if err != nil {
			return err
		}
	}
	req.Obj = obj
	return nil
}

// 定义结构体来存储解析结果
type ProtoService struct {
	Name    string       `json:"name"`
	Comment string       `json:"comment"`
	Methods []*RPCMethod `json:"methods"`
}

type RPCMethod struct {
	Name      string      `json:"name"`
	Req       string      `json:"req"`
	ReqFields []*ReqField `json:"req_fields"`
	Res       string      `json:"res"`
	Comment   string      `json:"comment"`
}

type ReqField struct {
	IsRepeated bool   `json:"is_repeated"`
	Type       string `json:"type"`
	Name       string `json:"name"`
	Seq        int    `json:"seq"`
	Comment    string `json:"comment"`
}

// 使用正则表达式解析proto服务接口定义的函数
func parseProtoService(protoContent string) []*ProtoService {
	lines := strings.Split(protoContent, "\n")
	var services []*ProtoService
	var currentService *ProtoService
	var currentComment string
	var currentMessage string
	reqFieldsMapByMessage := make(map[string][]*ReqField)

	servicePattern := regexp.MustCompile(`^\s*service\s+(\w+)\s*{`)
	rpcPattern := regexp.MustCompile(`^\s*rpc\s+(\w+)\s*\((\w+)\)\s*returns\s*\((\w+)\);`)
	messagePattern := regexp.MustCompile(`^\s*message\s+(\w+)\s*{`)
	reqFieldPattern := regexp.MustCompile(`^\s*(repeated\s+)?(\w+)\s+(\w+)\s*=\s*(\d+)\s*;`)
	reqFieldCommentPattern := regexp.MustCompile(`\s*;\s*//\s*(.+)$`)

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "//") {
			// 处理注释
			currentComment = strings.TrimSpace(strings.TrimPrefix(line, "//"))
		} else if servicePattern.MatchString(line) {
			// 解析服务名
			matches := servicePattern.FindStringSubmatch(line)
			if len(matches) == 2 {
				if currentService != nil {
					services = append(services, currentService)
				}
				currentService = &ProtoService{
					Name:    matches[1],
					Comment: currentComment,
					Methods: []*RPCMethod{},
				}
				currentComment = "" // 清空注释
				currentMessage = "" // 请求结构体定义
			}
		} else if rpcPattern.MatchString(line) {
			// 解析rpc方法
			matches := rpcPattern.FindStringSubmatch(line)
			if len(matches) == 4 && currentService != nil {
				rpcMethod := &RPCMethod{
					Name:      matches[1],
					Req:       matches[2],
					ReqFields: make([]*ReqField, 0),
					Res:       matches[3],
					Comment:   currentComment,
				}
				if reqFields, ok := reqFieldsMapByMessage[rpcMethod.Req]; ok {
					rpcMethod.ReqFields = reqFields
				}

				currentService.Methods = append(currentService.Methods, rpcMethod)
				currentComment = "" // 清空注释
			}
		} else if messagePattern.MatchString(line) {
			// 结构体定义
			matches := messagePattern.FindStringSubmatch(line)
			if len(matches) == 2 {
				currentMessage = matches[1]
				if _, ok := reqFieldsMapByMessage[currentMessage]; !ok {
					reqFieldsMapByMessage[currentMessage] = make([]*ReqField, 0)
				}
				currentComment = "" // 清空注释
			}
		} else if reqFieldPattern.MatchString(line) && currentMessage != "" {
			// 结构体字段定义
			matches := reqFieldPattern.FindStringSubmatch(line)
			if len(matches) == 5 {
				// 匹配行末注释
				if reqFieldCommentPattern.MatchString(line) {
					commentMatches := reqFieldCommentPattern.FindStringSubmatch(line)
					if len(commentMatches) == 2 {
						currentComment = strings.TrimSpace(commentMatches[1])
					}
				}

				reqField := &ReqField{
					IsRepeated: false,
					Type:       matches[2],
					Name:       matches[3],
					Comment:    currentComment,
				}
				if strings.TrimSpace(matches[1]) == "repeated" {
					reqField.IsRepeated = true
				}
				if matches[4] != "" {
					reqField.Seq, _ = strconv.Atoi(matches[4])
				}
				reqFieldsMapByMessage[currentMessage] = append(reqFieldsMapByMessage[currentMessage], reqField)
				currentComment = "" // 清空注释
			}
		}
	}

	if currentService != nil {
		services = append(services, currentService)
	}

	return services
}
