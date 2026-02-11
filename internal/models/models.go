package models

type User struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

type Post struct {
	ID       int    `json:"id"`
	IdAuthor int    `json:"author_id"`
	Text     string `json:"text"`
	Like     []int  `json:"like"`
}
