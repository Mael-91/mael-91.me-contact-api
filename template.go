package main

import (
	"bytes"
	"html/template"
)

func initTemplate() *template.Template {
	return template.Must(template.ParseGlob("template/*"))
}

func (c *contactForm) ParseTemplate(templateName string, data interface{}, h *Handlers) error {
	buf := new(bytes.Buffer)
	if err := h.Template.ExecuteTemplate(buf, templateName, data); err != nil {
		panic(err)
	}
	c.Body = buf.String()
	return nil
}