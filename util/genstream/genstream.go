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
	req := &ParserReq{}
	res := &ParserRes{}
	for _, config := range g.Configs {
		req.Text = config.Text
		req.Opts = config.Opts
		req.OptMap = GetOptionMap(req.Opts)
		parserFunc := NewParserFunc(config.Type)

		if err := parserFunc(ctx, req, res); err != nil {
			return nil, err
		}
	}
	return res, nil
}
