package models

type Sections struct {
	Parse struct {
		Title    string    `json:"title"`
		PageId   int       `json:"pageid"`
		Sections []Section `json:"sections"`
	} `json:"parse"`
}

type Section struct {
	Index string `json:"index"`
	Line  string `json:"line"`
}
