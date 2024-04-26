package conf

import (
	"coco-server/util/log"
	"context"
	_ "embed"
	"encoding/json"
	"go.uber.org/zap"
	"os"
)

func Init(ctx context.Context, configPath string) {
	config := make([]byte, 0)
	if configPath != "" {
		loadConfig, err := os.ReadFile(configPath)
		if err != nil {
			log.Error(ctx, "ReadFile", zap.Error(err))
			panic(err)
		}
		config = loadConfig
	} else {
		config = devConfig
	}
	err := json.Unmarshal(config, &Conf)
	if err != nil {
		log.Error(ctx, "Unmarshal", zap.Error(err))
		panic(err)
	}
}

//go:embed config/dev/config.json
var devConfig []byte

var Conf struct {
	ServiceName  string         `json:"serviceName"`
	Api          ApiConf        `json:"api"`
	Databases    []DatabaseConf `json:"databases"`
	RedisConfigs []RedisConf    `json:"redisConfigs"`
}

type ApiConf struct {
	Port    int    `json:"port"`
	GinMode string `json:"ginMode"`
}

type DatabaseConf struct {
	Name   string `json:"name"`
	Master string `json:"master"`
}

type RedisConf struct {
	Addr     string `json:"addr"`
	Password string `json:"password"`
	Database int    `json:"database"`
}
