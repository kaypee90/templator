package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPort(t *testing.T) {
	port := GetPort()
	assert.Equal(t, port, ":9898")
}

func TestGenerateEmailTemplate(t *testing.T) {
	var templateDetails TemplateDetails
	isSuccessful := GenerateEmailTemplate(&templateDetails)
	assert.Equal(t, isSuccessful, true)
}
