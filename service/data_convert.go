package service

import (
	"coco-server/model"
	"coco-server/util/generator"
	"coco-server/util/log"
	"context"
	"go.uber.org/zap"
	"regexp"
	"strings"
)

type dataConvertService struct{}

var DataConvertService = new(dataConvertService)

func (s *dataConvertService) Gen(ctx context.Context, req *model.DataConvertGenReq) (*model.DataConvertGenRes, error) {
	res := &model.DataConvertGenRes{}

	if len(req.DataSource) == 0 || len(req.Template) == 0 {
		return res, nil
	}

	// 读取请求参数
	dataSourceFieldMap := make(map[string]string)
	dataSourceDatas := make([]string, 0)
	dataSourceLines := strings.Split(req.DataSource, "\n")
	for _, line := range dataSourceLines {
		if IsCustomField(line) {
			field, value := GetCustomField(line)
			if field != "" {
				dataSourceFieldMap[field] = value
				continue
			}
		}
		dataSourceDatas = append(dataSourceDatas, line)
	}
	dataSource := strings.Join(dataSourceDatas, "\n")

	var g *generator.Generator
	switch req.DataSourceType {
	case "tabRow":
		sep, ok := dataSourceFieldMap["sep"]
		if ok {
			switch sep {
			case "\\n":
				sep = "\n"
			case "\\t":
				sep = "\t"
			case "\\s":
				sep = " "
			}
			g = generator.NewGenerator(generator.ConfigParser(generator.WithParserTabRowBySep(sep)))
		} else {
			g = generator.NewGenerator(generator.ConfigParser(generator.ParserTabRow))
		}
	case "json":
		g = generator.NewGenerator(generator.ConfigParser(generator.ParserJson))
	case "sql":
		g = generator.NewGenerator(generator.ConfigParser(generator.ParserSQL))
	default:
		g = generator.NewGenerator(generator.ConfigParser(generator.ParserTabRow))
	}
	g.Source(dataSource)

	// 拆分配置模板
	templates := make([]TemplateParam, 0)
	curName := ""
	curDatas := make([]string, 0)
	templateLines := strings.Split(req.Template, "\n")
	for _, line := range templateLines {
		if IsCustomField(line) {
			field, value := GetCustomField(line)
			if field == "name" {
				if curName != "" || len(curDatas) > 0 {
					templates = append(templates, TemplateParam{
						Name: curName,
						Data: strings.Join(curDatas, "\n"),
					})
				}
				curName = value
				curDatas = []string{}
				continue
			}
		}
		curDatas = append(curDatas, line)
	}
	if curName != "" || len(curDatas) > 0 {
		templates = append(templates, TemplateParam{
			Name: curName,
			Data: strings.Join(curDatas, "\n"),
		})
	}

	for _, item := range templates {
		err := g.Temp(item.Data)
		if err != nil {
			log.Error(ctx, "Gen Temp err", zap.Error(err))
			return nil, err
		}
		content := g.Exec()

		res.List = append(res.List, model.DataConvertGenData{
			Name:    item.Name,
			Content: content,
		})
	}

	return res, nil
}

type TemplateParam struct {
	Name string
	Data string
}

func IsCustomField(line string) bool {
	ns := strings.TrimSpace(line)
	return (strings.HasPrefix(ns, "{{/*") || strings.HasPrefix(ns, "{{- /*")) &&
		(strings.HasSuffix(ns, "*/}}") || strings.HasSuffix(ns, "*/ -}}")) &&
		fieldValue.Match([]byte(ns))
}

func GetCustomField(line string) (field string, value string) {
	if !IsCustomField(line) {
		return
	}
	ns := strings.TrimSpace(line)
	if strings.HasPrefix(ns, "{{/*") {
		ns = strings.TrimLeft(ns, "{{/*")
	}
	if strings.HasPrefix(ns, "{{- /*") {
		ns = strings.TrimLeft(ns, "{{- /*")
	}
	if strings.HasSuffix(ns, "*/}}") {
		ns = strings.TrimRight(ns, "*/}}")
	}
	if strings.HasSuffix(ns, "*/ -}}") {
		ns = strings.TrimRight(ns, "*/ -}}")
	}
	ns = strings.TrimSpace(ns)

	split := strings.Split(ns, "=")
	if len(split) != 2 {
		return
	}
	return split[0], split[1]
}

var fieldValue = regexp.MustCompile(`([a-zA-Z0-9_]+)=`)
