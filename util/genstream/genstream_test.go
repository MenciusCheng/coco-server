package genstream

import (
	"context"
	"fmt"
	"testing"
)

// 生成文本示例
func TestGenStream_Row(t *testing.T) {
	configs := make([]ParserConfig, 0)
	configs = append(configs, ParserConfig{
		Type: "row",
		Text: `index	name	age	created_at
1	cat	13	1998-09-03 00:02:18
1	dog	15	1996-05-03 18:03:22`,
		Opts: nil,
	})
	configs = append(configs, ParserConfig{
		Type: "temp",
		Text: `===start===
{{- range $index, $row := .Rows }}
{{ $index }}({{ len $row }}) --- {{ slice $row 1 }}
{{- end }}
===end===`,
		Opts: []ParserOption{
			{Type: "name", Value: "test-temp-name"},
		},
	})

	res, err := NewGenStream(configs).Gen(context.TODO())
	if err != nil {
		t.Errorf("Gen err = %v", err)
		return
	}
	// 生成结果
	fmt.Println(res.Show())
}

func TestGenStream_Reg(t *testing.T) {
	configs := make([]ParserConfig, 0)
	configs = append(configs, ParserConfig{
		Type: "text",
		Text: `
  // 请求接口1
  rpc ListGod(ListGodReq) returns (ListGodRes);
  // 请求接口2
  rpc ListGodTemple(ListGodTempleReq) returns (ListGodTempleRes);
`,
	})
	configs = append(configs, ParserConfig{
		Type: "reg",
		Text: `\/\/\s*(.+)\s*\n\s*rpc\s+([a-zA-Z0-9]+)\s*\(\s*([a-zA-Z0-9]+)\s*\)\s*returns\s*\(\s*([a-zA-Z0-9]+)\s*\);`,
	})
	configs = append(configs, ParserConfig{
		Type: "json",
		Text: `{"name":"act_god"}`,
	})
	configs = append(configs, ParserConfig{
		Type: "temp",
		Text: `.Rows: {{ JsonIndent .Rows }}
.Obj: {{ JsonIndent .Obj }}`,
		Opts: []ParserOption{{Type: "name", Value: "debug"}},
	})
	configs = append(configs, ParserConfig{
		Type: "temp",
		Text: `
{{- $obj := .Obj }}
{{- range $index, $row := .Rows }}
// {{ index $row 1 }}
func (extObj *ActivityExtObj) {{ index $row 2 }}(ctx context.Context, req *pb.{{ index $row 3 }}) (*pb.{{ index $row 4 }}, error) {
	return {{ $obj.name }}.{{ SnakeToCamel $obj.name }}HandlerObj.{{ index $row 2 }}(ctx, req)
}
{{ end }}`,
		Opts: []ParserOption{{Type: "name", Value: "res"}},
	})

	res, err := NewGenStream(configs).Gen(context.TODO())
	if err != nil {
		t.Errorf("Gen err = %v", err)
		return
	}
	// 生成结果
	fmt.Println(res.Show())
}
