package main

type PostBody struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Body   string `json:"body"`
}
