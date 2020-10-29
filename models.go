package main

// ShortenerRequest : shortener request payload **/
type ShortenerRequest struct {
	TemplateURL string `json:"templateUrl"`
}

// ShortenerResponse : shortener response payload **/
type ShortenerResponse struct {
	URL string `json:"url"`
}

// ResponseMessage : format of response message to be sent to user **/
type ResponseMessage struct {
	Message  string `json:"message"`
	Template string `json:"template,omitempty"`
}

// Action : an action to be triggered by a user **/
type Action struct {
	Instructions string `json:"instructions"`
	Button       Button `json:"button"`
	InviteCode   string `json:"inviteCode"`
}

// Button : button properties to be generated in template **/
type Button struct {
	Color     string `json:"color"`
	TextColor string `json:"textColor"`
	Text      string `json:"text"`
	Link      string `json:"link"`
}

// Product : product brand details **/
type Product struct {
	Name        string `json:"name"`
	Link        string `json:"link"`
	Logo        string `json:"logo"`
	Copyright   string `json:"copyright"`
	TroubleText string `json:"troubleText"`
}

// Email : content of the email **/
type Email struct {
	Body Body `json:"body"`
}

// Body : items in the body of the email **/
type Body struct {
	Name      string   `json:"name"`
	Intros    []string `json:"intros"`
	Actions   []Action `json:"actions"`
	Outros    []string `json:"outros"`
	Greeting  string   `json:"greeting"`
	Signature string   `json:"signature"`
	Title     string   `json:"title"`
}

// TemplateDetails : contains details of the template **/
type TemplateDetails struct {
	Product Product `json:"product"`
	Email   Email   `json:"email"`
}
