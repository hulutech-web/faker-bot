package orm

import (
	"faker-bot/backend/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"sync"
)

var (
	once sync.Once
)

// 申明一个MYSQL连接GormIns
var gormIns *gorm.DB

var dest interface{}
var result interface{}

type Gorm struct {
	gormIns *gorm.DB
}

func NewGorm(db string) *Gorm {
	return &Gorm{
		gormIns: BootMS(db),
	}
}
func BootMS(database string) *gorm.DB {
	once.Do(func() {
		ins, err := gorm.Open(sqlite.Open(database), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})
		if err != nil {
			panic(err)
		}
		gormIns = ins
	})
	return gormIns
}

func (o *Gorm) Migrate() error {
	return o.gormIns.AutoMigrate(&models.Category{}, &models.Comment{}, &models.Video{}, &models.Param{})
}
