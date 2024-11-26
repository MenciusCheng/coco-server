package event_rule

import (
	"context"
	"fmt"
	"github.com/MenciusCheng/go-util/log"
	"github.com/expr-lang/expr"
	"go.uber.org/zap"
)

// 处理事件
func HandleEvent(ctx context.Context, event *Event) {
	// 1. 获取事件定义
	//def, ok := eventDefinitions[event.Type]
	//if !ok {
	//	//log.Printf("Unknown event type: %s", event.Type)
	//	return
	//}

	// 2. 加载规则
	rules := LoadRulesForEvent(event.Type)

	for _, rule := range rules {
		// 3. 校验并匹配条件
		if match, err := evaluateCondition(rule.Conditions, event.Attributes); match {
			// 4. 执行动作
			for _, action := range rule.Actions {
				executeAction(action, event.Attributes)
			}
		} else if err != nil {
			log.Error(ctx, "evaluateCondition error", zap.Error(err))
		}
	}
}

var allRules = []Rule{
	{
		ID:        1,
		Name:      "送礼增加榜单积分,次数",
		EventType: EventTypeSendGift,
		Conditions: []RuleCondition{
			{
				Expression: "gift_id == 1001 && num > 0",
			},
		},
		Actions: []RuleAction{
			{
				Type: ActionTypeLog,
				Parameters: map[string]interface{}{
					"msg": "通过规则：送礼增加榜单积分",
				},
			},
			{
				Type: ActionTypeAddRankData,
				Parameters: map[string]interface{}{
					"pointField": "num",
				},
			},
		},
	},
}

func LoadRulesForEvent(eventType string) []Rule {
	rules := make([]Rule, 0)
	for _, rule := range allRules {
		if rule.EventType == eventType {
			rules = append(rules, rule)
		}
	}
	return rules
}

func evaluateCondition(conditions []RuleCondition, data map[string]interface{}) (bool, error) {
	if len(conditions) == 0 {
		return false, nil
	}
	for _, condition := range conditions {
		program, err := expr.Compile(condition.Expression, expr.Env(data))
		if err != nil {
			return false, err
		}

		output, err := expr.Run(program, data)
		if err != nil {
			return false, err
		}

		outputBool, ok := output.(bool)
		if !ok {
			return false, fmt.Errorf("evaluate expression returned non-boolean result")
		}
		if !outputBool {
			return false, nil
		}
	}
	return true, nil
}

// 示例：执行 API 调用动作
func executeAction(action RuleAction, data map[string]interface{}) error {
	switch action.Type {
	case ActionTypeLog:
		msg, ok := action.Parameters["msg"].(string)
		if !ok {
			msg = ""
		}
		log.Info(context.Background(), msg)
	case ActionTypeAddRankData:
		pointField, ok := action.Parameters["pointField"].(string)
		if !ok {
			pointField = "gift_id"
		}

		point, ok := data[pointField].(int)
		if !ok {
			pointField = "gift_id"
		}
		log.Info(context.Background(), fmt.Sprintf("add rank data, pointField: %s, point: %d", pointField, point))
	}
	return nil
}
