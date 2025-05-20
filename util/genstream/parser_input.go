package genstream

import (
	"coco-server/util/generator"
	"coco-server/util/log"
	"context"
	"encoding/json"
	"go.uber.org/zap"
	"regexp"
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
	Name    string `json:"name"`
	Req     string `json:"req"`
	Res     string `json:"res"`
	Comment string `json:"comment"`
}

// 使用正则表达式解析proto服务接口定义的函数
func parseProtoService(protoContent string) []*ProtoService {
	lines := strings.Split(protoContent, "\n")
	var services []*ProtoService
	var currentService *ProtoService
	var currentComment string

	servicePattern := regexp.MustCompile(`service\s+(\w+)\s*{`)
	rpcPattern := regexp.MustCompile(`rpc\s+(\w+)\s*\((\w+)\)\s*returns\s*\((\w+)\);`)

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
			}
		} else if rpcPattern.MatchString(line) {
			// 解析rpc方法
			matches := rpcPattern.FindStringSubmatch(line)
			if len(matches) == 4 && currentService != nil {
				currentService.Methods = append(currentService.Methods, &RPCMethod{
					Name:    matches[1],
					Req:     matches[2],
					Res:     matches[3],
					Comment: currentComment,
				})
				currentComment = "" // 清空注释
			}
		}
	}

	if currentService != nil {
		services = append(services, currentService)
	}

	return services
}
