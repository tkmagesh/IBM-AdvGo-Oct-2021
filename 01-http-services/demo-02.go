package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type ContactDetails struct {
	Email   string
	Subject string
	Message string
}

func main() {
	templ := template.Must(template.ParseFiles("contact.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			if err := templ.Execute(w, nil); err != nil {
				panic(err)
			}
			return
		}
		// Process form submission
		details := ContactDetails{
			r.FormValue("email"),
			r.FormValue("subject"),
			r.FormValue("message"),
		}
		fmt.Println(details)
		templ.Execute(w, struct {
			Email   string
			Message string
			Subject string
			Success bool
		}{details.Email, details.Message, details.Subject, true})
	})
	http.ListenAndServe(":8080", nil)
}
