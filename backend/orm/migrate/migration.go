package migrate

import (
	"faker-bot/backend/entities"
	"faker-bot/backend/orm"
)

func Migrate() error {
	has_cate := orm.GetInstance().Migrator().HasTable(&entities.Category{})
	has_cmt := orm.GetInstance().Migrator().HasTable(&entities.Comment{})
	has_vd := orm.GetInstance().Migrator().HasTable(&entities.Video{})
	has_pm := orm.GetInstance().Migrator().HasTable(&entities.Param{})
	if !has_cate || !has_cmt || !has_vd || !has_pm {
		err := orm.GetInstance().AutoMigrate(&entities.Category{}, &entities.Comment{}, &entities.Video{}, &entities.Param{})
		if err != nil {
			return err
		}
		Seeder()
	}

	return nil
}

func Seeder() {
	//分类数据
	categories := []entities.Category{
		{Name: "热门视频榜"},
		{Name: "话题榜"},
		{Name: "搞笑榜"},
		{Name: "美食榜"},
		{Name: "剧情榜"},
		{Name: "直播榜"},
		{Name: "热歌榜"},
		{Name: "热点视频"},
	}
	orm.GetInstance().Create(categories)
	//
	params := []entities.Param{
		{
			Url:   "https://open.douyin.com/data/extern/billboard/hot_video/",
			Label: "热门视频榜",
		},
		{
			Url:   "https://open.douyin.com/data/extern/billboard/topic/",
			Label: "话题榜",
		},
		{
			Url:   "https://open.douyin.com/data/extern/billboard/amusement/overall/",
			Label: "搞笑榜",
		},
		{
			Url:   "https://open.douyin.com/data/extern/billboard/drama/overall/",
			Label: "短剧",
		},
		{
			Url:   "https://open.douyin.com/hotsearch/sentences/",
			Label: "实时热词",
		},
		{
			Url:   "https://open.douyin.com/hotsearch/trending/sentences/?count=10&cursor=0",
			Label: "上升词",
		},
		{
			Url:   "https://open.douyin.com/hotsearch/videos/?",
			Label: "热词查询",
			Body:  "hot_sentence=教师节",
		},
	}
	orm.GetInstance().Create(params)
}
