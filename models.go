package main

type Request struct {
	Title string  `json:"title"`
}

type ResponseMessage struct {
	Message string `json:"message"`
}