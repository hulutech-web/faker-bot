package orm

import (
	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"sync"
)

var (
	once sync.Once
)

// 申明一个MYSQL连接GormIns
var gormIns *gorm.DB

type Gorm struct {
	gormIns *gorm.DB
}

func NewGorm() *Gorm {
	return &Gorm{
		gormIns: Boot(),
	}
}

func GetInstance() *gorm.DB {
	return gormIns
}

func Boot() *gorm.DB {
	once.Do(func() {
		viper.SetConfigFile(".env")
		viper.AutomaticEnv()
		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("读取配置文件错误: %v", err)
		}

		// 获取环境变量的值
		app_db_name := viper.GetString("APP_DB_NAME")
		ins, err := gorm.Open(sqlite.Open(app_db_name), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})
		if err != nil {
			panic(err)
		}
		gormIns = ins
	})
	return gormIns
}
