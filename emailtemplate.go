package main

import (
	"io/ioutil"
	"os"

	cloudinary "github.com/kaypee90/go-cloudinary"
	"github.com/matcornic/hermes/v2"
	log "github.com/sirupsen/logrus"
)

const templateName = "Preview.html"

// GenerateAndUploadEmailTemplate : Method for generating email tempates **/
func GenerateAndUploadEmailTemplate(templateDetails *TemplateDetails, uploader Uploader) (bool, string) {
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

	err = saveEmailTemplateToDiskForUpload(emailBody)
	if err != nil {
		panic(err)
	}

	temeplateURL := uploader.UploadTemplate(emailBody)

	return true, temeplateURL
}

func saveEmailTemplateToDiskForUpload(emailBody string) error {
	return ioutil.WriteFile(templateName, []byte(emailBody), 0644)
}

// Uploader : Interface for uploading html files **/
type Uploader interface {
	UploadTemplate(emailBody string) string
}

// CloudinaryUploader : Implementation for uploading files to cloudinary **/
type CloudinaryUploader struct {
}

// UploadTemplate : uploads templates to cloudinary **/
func (c CloudinaryUploader) UploadTemplate(emailBody string) string {
	cloudinaryURI := os.Getenv("CLOUDINARY_URL")
	cloudinaryFileBase := os.Getenv("CLOUDINARY_FILE_BASE")
	service, err := cloudinary.Dial(cloudinaryURI)
	if err != nil {
		panic(err)
	}

	var filename string
	// Read generated template from filesytem and upload
	filename, err = service.Upload(templateName, nil, "tmptr", true, cloudinary.RawType)
	if err != nil {
		panic(err)
	}

	path := cloudinaryFileBase + filename
	log.Info(path)
	return path
}
