package hotvideo

type VideoData struct {
	List []struct {
		DiggCount    int    `json:"digg_count"`
		HotValue     int    `json:"hot_value"`
		HotWords     string `json:"hot_words"`
		ItemCover    string `json:"item_cover"`
		PlayCount    int    `json:"play_count"`
		Rank         int    `json:"rank"`
		ShareURL     string `json:"share_url"`
		Title        string `json:"title"`
		Author       string `json:"author"`
		CommentCount int    `json:"comment_count"`
	} `json:"list"`
	ErrorCode   int    `json:"error_code"`
	Description string `json:"description"`
}

type Extra struct {
	ErrorCode      int    `json:"error_code"`
	Description    string `json:"description"`
	SubErrorCode   int    `json:"sub_error_code"`
	SubDescription string `json:"sub_description"`
	Logid          string `json:"logid"`
	Now            int64  `json:"now"`
}

type Root struct {
	Data  VideoData `json:"data"`
	Extra Extra     `json:"extra"`
}
