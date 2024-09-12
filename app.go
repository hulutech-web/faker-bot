package main

import (
	"context"
	"faker-bot/backend/entities"
	"faker-bot/backend/orm"
	"faker-bot/backend/orm/migrate"
	"faker-bot/backend/utils"
	"fmt"
	"log"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	//初始化数据库
	orm.NewGorm()
	//执行数据库迁移
	err := migrate.Migrate()
	//初始化缓存
	cache := utils.Boot()
	err_auth := cache.Authorize()
	if err_auth != nil {
		log.Fatal("接口访问失败", err_auth)
	}
	if err != nil {
		fmt.Println(err)
	}
	a.ctx = ctx
}
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s!", name)
}
func (p *App) List() []entities.Param {
	var params []entities.Param
	orm.GetInstance().Model(&entities.Param{}).Find(&params)
	return params
}
