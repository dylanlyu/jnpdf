package schema

// Reply is public return structure
type Reply struct {
	UpdateDate string `json:"updateDate" description:"更新日期"`
	IsError    bool   `json:"is_error" description:"是否錯誤 true為錯誤 flase為沒有錯誤"`
	ID         string `json:"id" description:"回傳ID"`
	Message    string `json:"message" description:"回傳訊息"`
}
