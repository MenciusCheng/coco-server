package parse

import "fmt"

func init() {
	funcMap["eqStr"] = FuncCompareObj.EqStr
	funcMap["inStr"] = FuncCompareObj.InStr
}

type FuncCompare struct{}

var FuncCompareObj = &FuncCompare{}

func (c *FuncCompare) EqStr(s1, s2 interface{}) bool {
	return fmt.Sprintf("%v", s1) == fmt.Sprintf("%v", s2)
}

func (c *FuncCompare) InStr(s1 interface{}, arr ...interface{}) bool {
	for _, s2 := range arr {
		if fmt.Sprintf("%v", s1) == fmt.Sprintf("%v", s2) {
			return true
		}
	}
	return false
}
