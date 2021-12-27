package models

type Content struct {
	Parse struct {
		Title  string `json:"title"`
		PageId int    `json:"pageid"`
		Text   struct {
			Html string `json:"*"`
		} `json:"text"`
	} `json:"parse"`
}
