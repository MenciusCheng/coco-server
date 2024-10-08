package log

import (
	"context"
	"go.uber.org/zap"
	"testing"
	"time"
)

func TestDefaultLogger(t *testing.T) {
	for i := 0; i < 10; i++ {
		Info(context.Background(), "test", zap.Int("i", i))
		time.Sleep(500 * time.Millisecond)
	}
}
