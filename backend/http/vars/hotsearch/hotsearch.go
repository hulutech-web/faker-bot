package hotsearch

import "time"

type Statistics struct {
	PlayCount     int `json:"play_count"`
	ShareCount    int `json:"share_count"`
	CommentCount  int `json:"comment_count"`
	DiggCount     int `json:"digg_count"`
	DownloadCount int `json:"download_count"`
	ForwardCount  int `json:"forward_count"`
}

type ListItem struct {
	CreateTime  int64      `json:"create_time"`
	IsReviewed  bool       `json:"is_reviewed"`
	ShareURL    string     `json:"share_url"`
	Title       string     `json:"title"`
	VideoStatus int        `json:"video_status"`
	Cover       string     `json:"cover"`
	IsTop       bool       `json:"is_top"`
	ItemID      string     `json:"item_id"`
	Statistics  Statistics `json:"statistics"`
}

type Data struct {
	Description string     `json:"description"`
	ErrorCode   int        `json:"error_code"`
	List        []ListItem `json:"list"`
}

type Extra struct {
	ErrorCode       int    `json:"error_code"`
	Description     string `json:"description"`
	SubErrorCode    int    `json:"sub_error_code"`
	SubDescription  string `json:"sub_description"`
	Logid           string `json:"logid"`
	Now             int64  `json:"now"`
	CurrentDateTime time.Time
}

type Root struct {
	Data  Data  `json:"data"`
	Extra Extra `json:"extra"`
}
