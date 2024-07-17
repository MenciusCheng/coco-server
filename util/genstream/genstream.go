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
		if len(config.Text) > 0 {
			req.Text = config.Text
		}
		req.Opts = config.Opts
		parserFunc := NewParserFunc(config.Type)

		if err := parserFunc(ctx, req, res); err != nil {
			return nil, err
		}
	}
	return res, nil
}
