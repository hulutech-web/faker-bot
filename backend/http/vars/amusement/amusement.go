package amusement

import "time"

type Video struct {
	ItemCover string `json:"item_cover"`
	ShareURL  string `json:"share_url"`
	Title     string `json:"title"`
}

type ListItem struct {
	RankChange       string  `json:"rank_change"`
	VideoList        []Video `json:"video_list"`
	Avatar           string  `json:"avatar"`
	EffectValue      int     `json:"effect_value"`
	FollowerCount    int     `json:"follower_count"`
	Nickname         string  `json:"nickname"`
	OnbillbaordTimes int     `json:"onbillbaord_times"`
	Rank             int     `json:"rank"`
}

type Data struct {
	List        []ListItem `json:"list"`
	ErrorCode   int        `json:"error_code"`
	Description string     `json:"description"`
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
