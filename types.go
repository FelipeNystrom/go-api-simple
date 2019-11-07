package main

type PostBody struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Text   string `json:"text"`
}

type ExistingPost struct {
	Id int `json:"id"`
	PostBody
}
