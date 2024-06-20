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
	temp, err := template.New("").Funcs(parse.GetFuncMap()).Parse(req.Text)
	if err != nil {
		log.Error(ctx, "ParserFuncTemp parse template error", zap.Error(err))
		return err
	}

	var b bytes.Buffer
	err = temp.Execute(&b, req)
	if err != nil {
		log.Error(ctx, "ParserFuncTemp execute template error", zap.Error(err))
		return err
	}

	res.List = append(res.List, ParserResData{
		Name:    GetOptionName(req.Opts),
		Content: b.String(),
	})
	return nil
}
