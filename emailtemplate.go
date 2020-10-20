package main

import (
	"io/ioutil"
	"os"

	cloudinary "github.com/gotsunami/go-cloudinary"
	"github.com/matcornic/hermes/v2"
	log "github.com/sirupsen/logrus"
)

// GenerateEmailTemplate : Method for generating email tempates **/
func GenerateEmailTemplate(templateDetails *TemplateDetails) (bool, string) {
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

	temeplateURL := UploadTemplate(emailBody)

	return true, temeplateURL
}

// UploadTemplate : uploads templates to cloudinary **/
func UploadTemplate(emailBody string) string {
	cloudinaryURI := os.Getenv("CLOUDINARY_URL")
	cloudinaryFileBase := os.Getenv("CLOUDINARY_FILE_BASE")
	service, err := cloudinary.Dial(cloudinaryURI)
	if err != nil {
		panic(err)
	}

	var filename string
	filename, err = service.Upload("Preview.html", nil, "tmptr", true, cloudinary.RawType)
	if err != nil {
		panic(err)
	}

	path := cloudinaryFileBase + filename
	log.Info(path)
	return path
}
