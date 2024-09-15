package main

import (
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
	"faker-bot/backend/orm/models"
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

func (p *App) HotVideo() hotvideo.Root {
	ctrl := http.NewController()
	resp, err := ctrl.HotVideo()
	if err != nil {
		fmt.Println(err)
	}
	var result hotvideo.Root
	decoder := json.NewDecoder(&resp)
	decoder.Decode(&result)
	return result
}

func (p *App) Topic() topic.Root {
	ctrl := http.NewController()
	resp, err := ctrl.Topic()
	if err != nil {
		fmt.Println(err)
	}
	var result topic.Root

	decoder := json.NewDecoder(&resp)
	decoder.Decode(&result)
	return result
}

func (p *App) Amusement() amusement.Root {
	ctrl := http.NewController()
	resp, err := ctrl.Amusement()
	if err != nil {
		fmt.Println(err)
	}
	var result amusement.Root

	decoder := json.NewDecoder(&resp)
	decoder.Decode(&result)
	for _, li := range result.Data.List {
		for _, i2 := range li.VideoList {
			var amusement models.Amusement
			amusement.Title = i2.Title
			amusement.ShareURL = i2.ShareURL
			amusement.ItemCover = i2.ItemCover
			orm.GetInstance().Model(&models.Amusement{}).
				Where("title = ?", i2.Title).FirstOrCreate(&amusement)
		}
	}
	return result
}

func (p *App) Drama() drama.Root {
	ctrl := http.NewController()
	resp, err := ctrl.Drama()
	if err != nil {
		fmt.Println(err)
	}

	var result drama.Root

	decoder := json.NewDecoder(&resp)
	decoder.Decode(&result)
	return result
}

func (p *App) HotSearch() hotsearch.Root {
	ctrl := http.NewController()
	resp, err := ctrl.HotSearch()
	if err != nil {
		fmt.Println(err)
	}

	var result hotsearch.Root

	decoder := json.NewDecoder(&resp)
	decoder.Decode(&result)
	return result
}

// 趋势上升词
func (p *App) Trending() trending.Root {
	ctrl := http.NewController()
	resp, err := ctrl.Trending()
	if err != nil {
		fmt.Println(err)
	}
	var result trending.Root

	decoder := json.NewDecoder(&resp)
	decoder.Decode(&result)
	for _, li := range result.Data.List {
		var trend models.Trend
		trend.Sentence = li.Sentence
		trend.HotLevel = li.HotLevel
		orm.GetInstance().Model(&models.Trend{}).
			Where("sentence = ?", li.Sentence).FirstOrCreate(&trend)
	}
	return result
}

// 热搜
func (p *App) HotSentence(keyword string) hotsentence.Root {
	ctrl := http.NewController()
	resp, err := ctrl.HotSentence(keyword)
	if err != nil {
		fmt.Println(err)
	}
	var result hotsentence.Root
	decoder := json.NewDecoder(&resp)
	decoder.Decode(&result)
	return result
}
