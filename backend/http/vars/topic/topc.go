package topic

type ListItem struct {
	RankChange  string `json:"rank_change"`
	Title       string `json:"title"`
	EffectValue int    `json:"effect_value"`
	Rank        int    `json:"rank"`
}

type VideoData struct {
	List        []ListItem `json:"list"`
	ErrorCode   int        `json:"error_code"`
	Description string     `json:"description"`
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
