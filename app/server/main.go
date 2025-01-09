package main

import (
	"coco-server/api"
	"coco-server/conf"
	"coco-server/middleware/db"
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
	// 解析命令行参数
	flag.Parse()

	// 初始化配置
	conf.Init(ctx, configPath)

	// init db
	db.InitMySQL(ctx)
	db.InitRedis(ctx)

	// init api
	api.InitApi(ctx)

	log.Info(ctx, "Server Start", zap.String("ServiceName", conf.Conf.ServiceName), zap.Any("port", conf.Conf.Api.Port))
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info(ctx, "Server Shutdown", zap.String("ServiceName", conf.Conf.ServiceName))
}
