package db

import (
	"coco-server/conf"
	ulog "coco-server/util/log"
	"context"
	"errors"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var MySQLCon *gorm.DB

func InitMySQL(ctx context.Context) {
	databases := conf.Conf.Databases
	if len(databases) == 0 {
		err := errors.New("databases is empty")
		ulog.Error(ctx, "InitMySQL err", zap.Error(err))
		panic(err)
	}

	// 初始化GORM日志配置
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level(这里记得根据需求改一下)
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)

	connString := databases[0].Master
	db, err := gorm.Open(mysql.Open(connString), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		ulog.Error(ctx, "db gorm.Open", zap.Error(err))
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		ulog.Error(ctx, "db lost", zap.Error(err))
		panic(err)
	}

	//设置连接池
	//空闲
	sqlDB.SetMaxIdleConns(10)
	//打开
	sqlDB.SetMaxOpenConns(20)

	MySQLCon = db
}
