package event_rule

import (
	"context"
	"testing"
	"time"
)

func TestHandleEvent_SendGift(t *testing.T) {
	event := &Event{
		Type: EventTypeSendGift,
		Attributes: map[string]interface{}{
			"rel_id":    1,
			"player_id": 123,
			"num":       1,
			"gift_id":   1001,
			"room_id":   88,
		},
		Timestamp: time.Now().Unix(),
	}
	HandleEvent(context.Background(), event)
}
