package trending

import "time"

type ListItem struct {
	HotLevel int    `json:"hot_level"`
	Label    int    `json:"label"`
	Sentence string `json:"sentence"`
}

type Data struct {
	HasMore     bool       `json:"has_more"`
	List        []ListItem `json:"list"`
	Total       int        `json:"total"`
	Cursor      int        `json:"cursor"`
	Description string     `json:"description"`
	ErrorCode   int        `json:"error_code"`
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
