package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sskmy1024/PartnerAssistant/api"
	"github.com/sskmy1024/PartnerAssistant/infrastructures"
	"github.com/stretchr/testify/assert"
)

func TestHelloSuccess(t *testing.T) {
	infrastructures.InitEnvironment()

	s := infrastructures.NewServer()
	api.Router(s)

	testData := `
	{
		"message": "焼肉食べたい"
	}
	`

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/sample", bytes.NewBuffer([]byte(testData)))
	req.Header.Set("Content-Type", "application/json")
	s.ServeHTTP(w, req)

	json := `{"message":"そうですか","image_path":""}
`
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, json, w.Body.String())
}
