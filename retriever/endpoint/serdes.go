package endpoint

type Record struct {
	UserID  int    `json:"user_id"`
	DataID  int    `json:"data_id"`
	Version int    `json:"version"`
	Content string `json:"content"`
}
