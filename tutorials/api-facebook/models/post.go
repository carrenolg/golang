package models

type Post struct {
	CreatedTime Time   `json:"created_time"`
	Message     string `json:"message,omitempty"`
	Id          string `json:"id"`
}

type Paging struct {
	Previous string `json:"previous"`
	Next     string `json:"next"`
}

type Response struct {
	Data   []Post `json:"data"`
	Paging `json:"paging"`
}
