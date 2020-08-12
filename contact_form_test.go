package main

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var handlers = Handlers{Template: initTemplate()}

func TestContactFormHandlerWithSuccess(t *testing.T) {
	form := &contactForm{
		Firstname: "Mael",
		Lastname:  "Constantin",
		Email:     "mael@test.com",
		Subject:   "Test",
		Message:   "Hello all",
		ReCaptcha: "aze",
	}
	jsonForm, _ := json.Marshal(form)
	req, err := http.NewRequest("POST", "/", bytes.NewReader(jsonForm))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.contactFormHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code :got %v want %v", status, http.StatusOK)
	}

}

func TestContactFormWithEmptyContent(t *testing.T) {
	req, err := http.NewRequest("POST", "/contact", bytes.NewBuffer(nil))
	if err != nil {
		t.Fatal(err)
	}
	resp := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.contactFormHandler)

	handler.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusBadRequest, resp.Code)
}

func TestContactFormWithBadMethod(t *testing.T) {
	req, err := http.NewRequest("GET", "/contact", nil)
	if err != nil {
		t.Fatal(err)
	}

	resp := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.contactFormHandler)

	handler.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusMethodNotAllowed, resp.Code)
}

func TestContactFormWithBadUrl(t *testing.T) {
	req, err := http.NewRequest("POST", "/contat", bytes.NewBuffer(nil))
	if err != nil {
		t.Fatal(err)
	}

	resp := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.contactFormHandler)

	handler.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusNotFound, resp.Code)
}
