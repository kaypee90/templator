package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockUploader struct {
}

// UploadTemplate : fake uploader function for testing **/
func (c MockUploader) UploadTemplate(emailBody string) string {
	return "http://mockurl.com"
}

func (c MockUploader) shortenURL(templateURL string) string {

	return templateURL
}

func TestGetPort(t *testing.T) {
	port := GetPort()
	assert.Equal(t, port, ":9898")
}

func TestGenerateAndUploadEmailTemplate(t *testing.T) {
	var templateDetails TemplateDetails
	uploader := MockUploader{}
	isSuccessful, template := GenerateAndUploadEmailTemplate(&templateDetails, uploader)
	assert.Equal(t, isSuccessful, true)
	assert.NotEmpty(t, template)
}
