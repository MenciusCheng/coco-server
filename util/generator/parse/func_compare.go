package parse

import "fmt"

func init() {
	funcMap["eqStr"] = FuncCompareObj.EqStr
}

type FuncCompare struct{}

var FuncCompareObj = &FuncCompare{}

func (c *FuncCompare) EqStr(s1, s2 interface{}) bool {
	return fmt.Sprintf("%v", s1) == fmt.Sprintf("%v", s2)
}
