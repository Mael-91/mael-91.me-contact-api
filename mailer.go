package main

import (
	"github.com/sendgrid/rest"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"os"
)

func mailer(cf contactForm) (resp *rest.Response, err error) {
	key := os.Getenv("SENDGRIP_API_KEY")

	receiver := os.Getenv("FIRSTNAME") + " " + os.Getenv("LASTNAME")
	receiverEmail := os.Getenv("RECEIVER")
	from := mail.NewEmail(cf.Firstname+" "+cf.Lastname, cf.Email)
	subject := cf.Subject
	to := mail.NewEmail(receiver, receiverEmail)

	plainTextContent := cf.Message
	html := "<strong>" + cf.Message + "<strong>" // Temporaire le temps de créer la template HTML pour les mails

	message := mail.NewSingleEmail(from, subject, to, plainTextContent, html)
	client := sendgrid.NewSendClient(key)
	response, err := client.Send(message)
	displayError(err)

	return response, nil
}
