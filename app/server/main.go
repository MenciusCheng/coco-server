package main

import (
	"coco-server/api"
	"coco-server/conf"
	"coco-server/middleware/db"
	"coco-server/util/bookmark"
	"context"
	"flag"
	"github.com/MenciusCheng/go-util/log"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx := context.TODO()
	// init config
	var configPath string
	flag.StringVar(&configPath, "c", "", "Configuration file")
	conf.Init(ctx, configPath)

	// init db
	db.InitMySQL(ctx)
	db.InitRedis(ctx)

	// init api
	api.InitApi(ctx)

	// 插件
	bookmark.Setup(api.GetRouterGroup(), db.MySQLCon)

	log.Info(ctx, "Server Start", zap.String("ServiceName", conf.Conf.ServiceName), zap.Any("port", conf.Conf.Api.Port))
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info(ctx, "Server Shutdown", zap.String("ServiceName", conf.Conf.ServiceName))
}
