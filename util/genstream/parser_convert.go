package genstream

import (
	"coco-server/util/log"
	"context"
	"go.uber.org/zap"
	"regexp"
	"strings"
)

// 转换类型解析器

func ParserFuncReg(ctx context.Context, req *ParserReq, res *ParserRes) error {
	expr, err := regexp.Compile(req.Text)
	if err != nil {
		log.Error(ctx, "ParserFuncReg Compile error", zap.Error(err), zap.String("reqText", req.Text))
		return err
	}
	rows := make([][]string, 0)
	for _, row := range req.Rows {
		if len(row) == 0 {
			continue
		}
		if !expr.MatchString(row[0]) {
			continue
		}

		if replace := GetOptionReplace(req.Opts); len(replace) > 0 {
			// 替换模式
			replaceRes := expr.ReplaceAllString(row[0], replace)
			rows = append(rows, []string{replaceRes})
		} else {
			// 匹配模式
			matchRes := expr.FindAllStringSubmatch(row[0], -1)
			rows = append(rows, matchRes...)
		}
	}

	req.Rows = rows
	return nil
}

// ParserFuncSplit ["a,b,c"] => ["a,b,c", "a", "b", "c"]
func ParserFuncSplit(ctx context.Context, req *ParserReq, res *ParserRes) error {
	sep := GetOptionSep(req.Opts, "\t")
	rows := make([][]string, 0)
	for _, row := range req.Rows {
		if len(row) == 0 {
			continue
		}
		cols := []string{row[0]}
		cols = append(cols, strings.Split(row[0], sep)...)
		rows = append(rows, cols)
	}

	req.Rows = rows
	return nil
}

// ParserFuncJoin ["a,b,c", "a", "b", "c"] => ["a,b,c"]
func ParserFuncJoin(ctx context.Context, req *ParserReq, res *ParserRes) error {
	sep := GetOptionSep(req.Opts, "\t")
	rows := make([][]string, 0)
	for _, row := range req.Rows {
		if len(row) <= 1 {
			continue
		}
		cols := []string{strings.Join(row[1:], sep)}
		rows = append(rows, cols)
	}

	req.Rows = rows
	return nil
}
