package main

import (
	"encoding/json"
	"github.com/gorilla/schema"
	"net/http"
)

var decoder = schema.NewDecoder()

type contactForm struct {
	Firstname string
	Lastname  string
	Email     string
	Subject   string
	Message   string
	ReCaptcha string
	Body string
}

type response struct {
	Code    int64
	Message string
}

func (h *Handlers) contactFormHandler(w http.ResponseWriter, r *http.Request) {
	response := response{
		Code:    200,
		Message: "The email has been sent",
	}
	w.Header().Set("Content-Type", "application/json")

	if r.URL.Path != "/contact" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "An error occurred. Please contact administrator (webmaster@mael-91.me)", http.StatusInternalServerError)
		return
	}

	if r.ContentLength == 0 {
		http.Error(w, "Information cannot be empty", http.StatusBadRequest)
		return
	}

	var contactForm contactForm

	err = decoder.Decode(&contactForm, r.Form)
	if err != nil {
		response.Code = 400
		response.Message = "Invalid schema"
		json.NewEncoder(w).Encode(response)
		return
	}

	err = contactForm.ParseTemplate("mail.html", contactForm, h)
	checkErr(err)

	_ = mailer(contactForm)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}