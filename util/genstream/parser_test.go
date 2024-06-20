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
