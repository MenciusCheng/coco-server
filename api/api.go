package api

import (
	"coco-server/conf"
	"coco-server/util/log"
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

var router = func() *gin.Engine {
	return gin.Default()
}()

func GetRouterGroup() *gin.RouterGroup {
	return router.Group("/api")
}

// 初始化路由
func InitApi(ctx context.Context) {
	apiConf := conf.Conf.Api
	gin.SetMode(apiConf.GinMode)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", apiConf.Port),
		Handler: router,
	}
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error(ctx, "ListenAndServe", zap.Error(err))
		}
	}()
}

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data,omitempty"`
}
