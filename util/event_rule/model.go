package event_rule

type Event struct {
	Type       string                 // 事件类型，例如 "OrderCreated"
	Attributes map[string]interface{} // 对象属性，例如 {"order_id": 123, "amount": 100.0}
	Timestamp  int64                  // 事件发生时间
}

type EventDefinition struct {
	Attributes     map[string]string // 属性及其类型，如 {"order_id": "int", "amount": "float"}
	AllowedActions []string          // 允许的操作，如 ["send_notification", "apply_discount"]
}

type Rule struct {
	ID         int64           // 规则 ID
	Name       string          // 规则名称
	EventType  string          // 绑定的事件类型
	Conditions []RuleCondition // 条件表达式，例如 "amount > 100 && user_id != null"
	Actions    []RuleAction    // 动作列表
}

type RuleCondition struct {
	Expression string // 条件表达式，例如 "amount > 100 && user_id != null"
}

type RuleAction struct {
	Type       string                 // 动作类型，例如 "send_notification"
	Parameters map[string]interface{} // 动作参数
}
