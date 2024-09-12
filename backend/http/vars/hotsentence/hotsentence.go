package hotsentence

import "time"

type Statistics struct {
	CommentCount  int `json:"comment_count"`
	DiggCount     int `json:"digg_count"`
	DownloadCount int `json:"download_count"`
	ForwardCount  int `json:"forward_count"`
	PlayCount     int `json:"play_count"`
	ShareCount    int `json:"share_count"`
}

type ListItem struct {
	Statistics  Statistics `json:"statistics"`
	Title       string     `json:"title"`
	Cover       string     `json:"cover"`
	CreateTime  int64      `json:"create_time"`
	IsReviewed  bool       `json:"is_reviewed"`
	IsTop       bool       `json:"is_top"`
	ItemID      string     `json:"item_id"`
	ShareURL    string     `json:"share_url"`
	VideoStatus int        `json:"video_status"`
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
