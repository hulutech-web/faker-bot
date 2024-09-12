package main

import (
	"bytes"
	"context"
	"encoding/json"
	"faker-bot/backend/entities"
	"faker-bot/backend/http"
	"faker-bot/backend/http/vars/amusement"
	"faker-bot/backend/http/vars/drama"
	"faker-bot/backend/http/vars/hotsearch"
	"faker-bot/backend/http/vars/hotsentence"
	"faker-bot/backend/http/vars/hotvideo"
	"faker-bot/backend/http/vars/topic"
	"faker-bot/backend/http/vars/trending"
	"faker-bot/backend/orm"
	"faker-bot/backend/orm/migrate"
	"faker-bot/backend/utils"
	"fmt"
	"io"
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

func (p *App) HotVideo() {
	ctrl := http.NewController()
	resp, err := ctrl.HotVideo()
	if err != nil {
		fmt.Println(err)
	}
	resultBuffer := bytes.NewBuffer([]byte{})
	_, _ = io.Copy(resultBuffer, resp.Body)
	fmt.Println(resultBuffer.String())
	var result hotvideo.Root
	_ = json.Unmarshal(resultBuffer.Bytes(), &result)
}

func (p *App) Topic() {
	ctrl := http.NewController()
	resp, err := ctrl.Topic()
	if err != nil {
		fmt.Println(err)
	}
	resultBuffer := bytes.NewBuffer([]byte{})
	_, _ = io.Copy(resultBuffer, resp.Body)
	fmt.Println(resultBuffer.String())
	var result topic.Root
	_ = json.Unmarshal(resultBuffer.Bytes(), &result)
}

func (p *App) Amusement() {
	ctrl := http.NewController()
	resp, err := ctrl.Amusement()
	if err != nil {
		fmt.Println(err)
	}
	resultBuffer := bytes.NewBuffer([]byte{})
	_, _ = io.Copy(resultBuffer, resp.Body)
	fmt.Println(resultBuffer.String())
	var result amusement.Root
	_ = json.Unmarshal(resultBuffer.Bytes(), &result)
}

func (p *App) Drama() {
	ctrl := http.NewController()
	resp, err := ctrl.Drama()
	if err != nil {
		fmt.Println(err)
	}
	resultBuffer := bytes.NewBuffer([]byte{})
	_, _ = io.Copy(resultBuffer, resp.Body)
	fmt.Println(resultBuffer.String())
	var result drama.Root
	_ = json.Unmarshal(resultBuffer.Bytes(), &result)
}

func (p *App) Search(keyword string) {

}

func (p *App) HotSearch(keyword string) {
	ctrl := http.NewController()
	resp, err := ctrl.HotSearch()
	if err != nil {
		fmt.Println(err)
	}
	resultBuffer := bytes.NewBuffer([]byte{})
	_, _ = io.Copy(resultBuffer, resp.Body)
	fmt.Println(resultBuffer.String())
	var result hotsearch.Root
	_ = json.Unmarshal(resultBuffer.Bytes(), &result)
}

func (p *App) Trending() {
	ctrl := http.NewController()
	resp, err := ctrl.Trending()
	if err != nil {
		fmt.Println(err)
	}
	resultBuffer := bytes.NewBuffer([]byte{})
	_, _ = io.Copy(resultBuffer, resp.Body)
	fmt.Println(resultBuffer.String())
	var result trending.Root
	_ = json.Unmarshal(resultBuffer.Bytes(), &result)
}

func (p *App) HotSentence(keyword string) {
	ctrl := http.NewController()
	resp, err := ctrl.HotSentence(keyword)
	if err != nil {
		fmt.Println(err)
	}
	resultBuffer := bytes.NewBuffer([]byte{})
	_, _ = io.Copy(resultBuffer, resp.Body)
	fmt.Println(resultBuffer.String())
	var result hotsentence.Root
	_ = json.Unmarshal(resultBuffer.Bytes(), &result)
}
