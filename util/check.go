package util

import (
	"coco-server/util/log"
	"context"
	"go.uber.org/zap"
)

func CheckGoPanic(ctx context.Context) {
	if r := recover(); r != nil {
		log.DPanic(ctx, "panic recovered ", zap.Any("panic", r))
	}
}
