package main

import (
	"io/ioutil"

	"github.com/matcornic/hermes/v2"
)

// GenerateEmailTemplate : Method for generating email tempates **/
func GenerateEmailTemplate(templateDetails *TemplateDetails) bool {
	productDetails := templateDetails.Product
	emailBodyDetails := templateDetails.Email.Body

	h := hermes.Hermes{
		Product: hermes.Product{
			Name:        productDetails.Name,
			Link:        productDetails.Link,
			Logo:        productDetails.Logo,
			Copyright:   productDetails.Copyright,
			TroubleText: productDetails.TroubleText,
		},
	}

	email := hermes.Email{
		Body: hermes.Body{
			Name:     emailBodyDetails.Name,
			Title:    emailBodyDetails.Title,
			Intros:   emailBodyDetails.Intros,
			Greeting: emailBodyDetails.Greeting,
			Actions:  []hermes.Action{},
			Outros:   emailBodyDetails.Outros,
		},
	}

	if emailBodyDetails.Actions != nil {
		for _, action := range emailBodyDetails.Actions {
			hermesAction := hermes.Action{
				Instructions: action.Instructions,
				Button: hermes.Button{
					Color: action.Button.Color, // Optional action button color
					Text:  action.Button.Text,
					Link:  action.Button.Link,
				},
			}

			email.Body.Actions = append(email.Body.Actions, hermesAction)
		}
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
