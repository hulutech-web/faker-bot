package http

import "net/http"

type Api interface {
	//热门视频榜https://open.douyin.com/data/extern/billboard/hot_video/
	HotVideo() (http.Response, error)
	//	话题榜https://open.douyin.com/data/extern/billboard/topic/
	Topic() (http.Response, error)
	//	搞笑榜https://open.douyin.com/data/extern/billboard/amusement/overall/
	Amusement() (http.Response, error)
	//	短剧 https://open.douyin.com/data/extern/billboard/drama/overall/
	Drama() (http.Response, error)
	//	关键词搜索 https://open.douyin.com/video/search/?count=20&keyword=毒舌扒姨太&open_id={{access_token}}&cursor=1
	Search(keyword string) (http.Response, error)
	//	热词 https://open.douyin.com/hotsearch/sentences/
	HotSearch() (http.Response, error)
	//趋势上升 https://open.douyin.com/hotsearch/trending/sentences/?count=10&cursor=0
	Trending() (http.Response, error)
	//	热词聚合视频 https://open.douyin.com/hotsearch/videos/?hot_sentence=开学
	HotSentence(hot_sentence string) (http.Response, error)
}
