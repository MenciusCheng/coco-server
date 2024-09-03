package parse

func init() {
	funcMap["add"] = FuncIntObj.Add
}

type FuncInt struct{}

var FuncIntObj = &FuncInt{}

func (c *FuncInt) Add(num1, num2 interface{}) int {
	return num1.(int) + num2.(int)
}
