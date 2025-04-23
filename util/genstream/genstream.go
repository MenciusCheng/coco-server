package genstream

import (
	"context"
)

type GenStream struct {
	Configs []ParserConfig `json:"configs"`
}

func NewGenStream(configs []ParserConfig) *GenStream {
	return &GenStream{
		Configs: configs,
	}
}

func (g *GenStream) Gen(ctx context.Context) (*ParserRes, error) {
	req := &ParserReq{
		OptMap: make(map[string]string),
	}
	res := &ParserRes{}
	for _, config := range g.Configs {
		req.Text = config.Text
		req.Opts = config.Opts
		configOptMap := GetOptionMap(config.Opts)
		for k, v := range configOptMap {
			// 字典选项可以继承上一个 OptMap
			req.OptMap[k] = v
		}
		parserFunc := NewParserFunc(config.Type)

		if err := parserFunc(ctx, req, res); err != nil {
			return nil, err
		}
	}
	return res, nil
}
