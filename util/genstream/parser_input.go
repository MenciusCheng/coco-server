package genstream

import (
	"coco-server/util/generator"
	"coco-server/util/log"
	"context"
	"encoding/json"
	"go.uber.org/zap"
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
	obj := generator.ParserSQL2(req.Text)

	req.Obj = obj
	return nil
}
