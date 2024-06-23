package genstream

import (
	"context"
	"reflect"
	"testing"
)

func TestParserFuncJson(t *testing.T) {
	type args struct {
		ctx context.Context
		req *ParserReq
		res *ParserRes
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		wantObj map[string]interface{}
	}{
		{
			args: args{
				ctx: context.Background(),
				req: &ParserReq{
					Text: "{\"rows\":[[\"abc\",\"a\",\"b\",\"c\"]]}",
				},
			},
			wantObj: map[string]interface{}{
				"rows": []interface{}{
					[]interface{}{"abc", "a", "b", "c"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ParserFuncJson(tt.args.ctx, tt.args.req, tt.args.res); (err != nil) != tt.wantErr {
				t.Errorf("ParserFuncJson() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.args.req.Obj, tt.wantObj) {
				t.Errorf("ParserFuncJson() Obj = %v, wantObj %v", tt.args.req.Obj, tt.wantObj)
			}
		})
	}
}

func TestParserFuncTemp(t *testing.T) {
	type args struct {
		ctx context.Context
		req *ParserReq
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		wantRes *ParserRes
	}{
		{
			args: args{
				ctx: context.Background(),
				req: &ParserReq{
					Text: "Hello {{ .Obj.echo }}!",
					Obj: map[string]interface{}{
						"name": "cat",
						"echo": "World",
					},
					Opts: []ParserOption{
						{
							Type:  "name",
							Value: "{{ .Obj.name }}",
						},
					},
				},
			},
			wantRes: &ParserRes{
				List: []ParserResData{
					{
						Name:    "cat",
						Content: "Hello World!",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes := &ParserRes{}
			if err := ParserFuncTemp(tt.args.ctx, tt.args.req, gotRes); (err != nil) != tt.wantErr {
				t.Errorf("ParserFuncTemp() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("ParserFuncTemp() gotRes = %+v, wantRes %+v", gotRes, tt.wantRes)
			}
		})
	}
}
