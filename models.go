package main

type ResponseMessage struct {
	Message string `json:"message"`
}

type Action struct {
	Instructions string `json:"instructions"`
	Button       Button `json:"button"`
	InviteCode   string `json:"inviteCode"`
}

type Button struct {
	Color     string `json:"color"`
	TextColor string `json:"textColor"`
	Text      string `json:"text"`
	Link      string `json:"link"`
}

type Product struct {
	Name        string `json:"name"`
	Link        string `json:"link"`
	Logo        string `json:"logo"`
	Copyright   string `json:"copyright"`
	TroubleText string `json:"troubleText"`
}

type Email struct {
	Body Body `json:"body"`
}

type Body struct {
	Name      string   `json:"name"`
	Intros    []string `json:"intros"`
	Actions   []Action `json:"actions"`
	Outros    []string `json:"outros"`
	Greeting  string   `json:"greeting"`
	Signature string   `json:"signature"`
	Title     string   `json:"title"`
}

type TemplateDetails struct {
	Product Product `json:"product"`
	Email   Email   `json:"email"`
}
