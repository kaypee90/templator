package main

import (
	"io/ioutil"

	"github.com/matcornic/hermes/v2"
)

// GenerateEmailTemplate : Method for generating email tempates **/
func GenerateEmailTemplate(templateDetails *TemplateDetails) bool {
	h := hermes.Hermes{
		Product: hermes.Product{
			Name:      "Hermes",
			Link:      "https://example-hermes.com/",
			Logo:      "http://www.duchess-france.org/wp-content/uploads/2016/01/gopher.png",
			Copyright: "Copyright © 2020. All rights reserved.",
		},
	}

	email := hermes.Email{
		Body: hermes.Body{
			Title: "",
			Intros: []string{
				"Welcome to Hermes! We're very execited to have you on board",
			},
			Greeting: "Hi",
			Actions: []hermes.Action{
				hermes.Action{
					Instructions: "To get started with Hermes, please click here:",
					Button: hermes.Button{
						Color: "#22BC66", // Optional action button color
						Text:  "Confirm your account",
						Link:  "https://hermes-example.com/confirm?token=d9729feb74992cc3482b350163a1a010",
					},
				},
			},
			Outros: []string{
				"Need help, or have questions? Just reply to this email, we'd love to help.",
			},
		},
	}

	emailBody, err := h.GenerateHTML(email)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("Preview.html", []byte(emailBody), 0644)
	if err != nil {
		panic(err)
	}

	return true
}
