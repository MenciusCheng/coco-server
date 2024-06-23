package genstream

import (
	"bytes"
	"coco-server/util/generator/parse"
	"coco-server/util/log"
	"context"
	"go.uber.org/zap"
	"text/template"
)

// 输出类型解析器

func ParserFuncTemp(ctx context.Context, req *ParserReq, res *ParserRes) error {
	content, err := TemplateExecute(ctx, req.Text, req)
	if err != nil {
		log.Error(ctx, "TemplateExecute text error", zap.Error(err))
		return err
	}

	optName := GetOptionName(req.Opts)
	name, err := TemplateExecute(ctx, optName, req)
	if err != nil {
		log.Error(ctx, "TemplateExecute optName error", zap.Error(err))
		return err
	}

	res.List = append(res.List, ParserResData{
		Name:    name,
		Content: content,
	})
	return nil
}

func TemplateExecute(ctx context.Context, tempText string, data any) (string, error) {
	temp, err := template.New("").Funcs(parse.GetFuncMap()).Parse(tempText)
	if err != nil {
		log.Error(ctx, "parse error", zap.Error(err), zap.String("text", tempText))
		return "", err
	}

	var b bytes.Buffer
	err = temp.Execute(&b, data)
	if err != nil {
		log.Error(ctx, "execute error", zap.Error(err), zap.String("text", tempText), zap.Any("data", data))
		return "", err
	}
	return b.String(), nil
}
