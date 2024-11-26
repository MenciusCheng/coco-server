package event_rule

const (
	EventTypeSendGift = "send_gift"
)

var eventDefinitions = map[string]EventDefinition{
	EventTypeSendGift: {
		Attributes: map[string]string{},
		AllowedActions: []string{
			ActionTypeLog,
		},
	},
}

const (
	ActionTypeLog         = "log"
	ActionTypeAddRankData = "add_rank_data"
)
