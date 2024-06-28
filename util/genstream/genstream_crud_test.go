package genstream

import (
	"coco-server/util/genstream/demo_tmpl"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestGenStream_CRUD_conf(t *testing.T) {

	b, err := os.ReadFile("../../sql/gen_stream_conf.sql")
	if err != nil {
		t.Errorf("ReadFile err = %v", err)
		return
	}

	configs := make([]ParserConfig, 0)
	configs = append(configs, ParserConfig{
		Type: "createSql",
		Text: string(b),
	})
	configs = append(configs, ParserConfig{
		Type: "temp",
		Text: `.Rows: {{ JsonIndent .Rows }}
.Obj: {{ JsonIndent .Obj }}`,
		Opts: []ParserOption{{Type: "name", Value: "debug"}},
	})

	configs = append(configs, ParserConfig{
		Type: "temp",
		Opts: []ParserOption{{Type: "name", Value: "model/{{ .Obj.tableName }}.go"}},
		Text: demo_tmpl.CrudModel,
	})

	configs = append(configs, ParserConfig{
		Type: "temp",
		Opts: []ParserOption{{Type: "name", Value: "dao/{{ .Obj.tableName }}.go"}},
		Text: demo_tmpl.CrudDao,
	})

	res, err := NewGenStream(configs).Gen(context.TODO())
	if err != nil {
		t.Errorf("Gen err = %v", err)
		return
	}
	// 生成结果
	fmt.Println(res.Show())
}

func TestGenStream_CRUD_cache(t *testing.T) {
	configs := make([]ParserConfig, 0)

	jsonMap := map[string]interface{}{
		"tableName": "demo_table",
		"comment":   "注释",
		"rows": []map[string]interface{}{
			{
				"type":    "lock",
				"name":    "demoAbc",
				"comment": "abc锁",
			},
		},
	}
	jsonBytes, _ := json.Marshal(jsonMap)
	configs = append(configs, ParserConfig{
		Type: ParserTypeJson,
		Text: string(jsonBytes),
	})

	configs = append(configs, ParserConfig{
		Type: "temp",
		Opts: []ParserOption{{Type: "name", Value: "cache/{{ .Obj.tableName }}.go"}},
		Text: demo_tmpl.CrudCache,
	})

	res, err := NewGenStream(configs).Gen(context.TODO())
	if err != nil {
		t.Errorf("Gen err = %v", err)
		return
	}
	// 生成结果
	fmt.Println(res.Show())
}
