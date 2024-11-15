package generator

import (
	"coco-server/util/log"
	"context"
	"encoding/json"
	"go.uber.org/zap"
	"regexp"
	"strings"
)

type SqlTable struct {
	SqlHead
	SqlFoot
	Rows []SqlField `json:"rows"`
}

type SqlHead struct {
	TableName string `json:"tableName"`
}

type SqlField struct {
	ColName string `json:"colName"`
	ColType string `json:"colType"`
	Comment string `json:"comment"`
	GoType  string `json:"goType"`
}

type SqlFoot struct {
	Comment string `json:"comment"`
}

func CalSqlHead(line string) *SqlHead {
	headReg := regexp.MustCompile("^\\s*CREATE\\s+TABLE\\s+([a-zA-Z0-9_`]+\\.)?([a-zA-Z0-9_`]+)\\s*\\(?\\s*$") // CREATE TABLE tbl_name
	if !headReg.MatchString(line) {
		return nil
	}
	submatch := headReg.FindStringSubmatch(line)
	if len(submatch) < 3 {
		return nil
	}

	tableName := strings.Trim(submatch[2], "`")
	return &SqlHead{
		TableName: tableName,
	}
}

func CalSqlField(line string, colNameGoTypeMap map[string]string) *SqlField {
	fieldReg := regexp.MustCompile("^\\s*([a-zA-Z0-9_`]+)\\s+([a-zA-Z0-9]+)\\(?[0-9]*\\)?\\s+") // col_name column_definition
	if !fieldReg.MatchString(line) {
		return nil
	}
	submatch := fieldReg.FindStringSubmatch(line)
	if len(submatch) < 3 {
		return nil
	}
	colName := strings.Trim(submatch[1], "`")
	colType := strings.ToLower(submatch[2])

	colTypeToGoTypeMap := map[string]string{
		"int":      "int64",
		"bigint":   "int64",
		"varchar":  "string",
		"json":     "string",
		"text":     "string",
		"datetime": "time.Time",
	}
	goType := ""
	if goTypeVal, ok := colNameGoTypeMap[colName]; ok {
		goType = goTypeVal
	} else if goTypeVal, ok := colTypeToGoTypeMap[colType]; ok {
		goType = goTypeVal
	}

	res := &SqlField{
		ColName: colName,
		ColType: colType,
		GoType:  goType,
	}

	fieldCommentReg := regexp.MustCompile("\\s+COMMENT\\s+'(.+)',?\\s*$") // COMMENT
	commentSubmatch := fieldCommentReg.FindStringSubmatch(line)
	if len(commentSubmatch) >= 2 {
		res.Comment = commentSubmatch[1]
	}

	return res
}

func CalSqlFoot(line string) *SqlFoot {
	footReg := regexp.MustCompile("^\\s*\\)")
	if !footReg.MatchString(line) {
		return nil
	}
	res := &SqlFoot{}

	footCommentReg := regexp.MustCompile("COMMENT\\s*=\\s*'(.+)'\\S*;?$")
	commentSubmatch := footCommentReg.FindStringSubmatch(line)
	if len(commentSubmatch) >= 2 {
		res.Comment = commentSubmatch[1]
	}
	return res
}

func ParserSQL2NoMap(text string) map[string]interface{} {
	// 兼容旧函数签名
	return ParserSQL2(text, nil)
}

func ParserSQL2(text string, colNameGoTypeMap map[string]string) map[string]interface{} {
	res := make(map[string]interface{})

	sqlTable := SqlTable{}
	lines := strings.Split(text, "\n")
	for _, line := range lines {
		// 清洗
		lineData := strings.TrimSpace(line)
		if len(lineData) == 0 {
			continue
		}

		if strings.HasPrefix(lineData, "PRIMARY KEY") {
			// 预留
		} else if strings.HasPrefix(lineData, "KEY") {
			// 预留
		} else if CalSqlHead(lineData) != nil {
			sqlHead := CalSqlHead(lineData)
			if sqlHead != nil {
				sqlTable.TableName = sqlHead.TableName
			}
		} else if sqlField := CalSqlField(lineData, colNameGoTypeMap); sqlField != nil {
			sqlTable.Rows = append(sqlTable.Rows, *sqlField)
		} else if CalSqlFoot(lineData) != nil {
			sqlFoot := CalSqlFoot(lineData)
			if sqlFoot != nil {
				sqlTable.Comment = sqlFoot.Comment
			}
		}
	}
	bytes, err := json.Marshal(sqlTable)
	if err != nil {
		log.Error(context.TODO(), "Marshal failed", zap.Error(err))
		return res
	}

	err = json.Unmarshal(bytes, &res)
	if err != nil {
		log.Error(context.TODO(), "Unmarshal failed", zap.Error(err))
		return res
	}

	return res
}
