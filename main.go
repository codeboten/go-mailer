package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	mailgun "gopkg.in/mailgun/mailgun-go.v1"
)

type email struct {
	Name         string `json:"name"`
	EmailAddress string `json:"email"`
	Body         string `json:"message"`
}

func send(body string) {
	yourdomain := os.Getenv("MG_DOMAIN")
	key := os.Getenv("MG_KEY")
	publicAPIKey := os.Getenv("MG_PUBLIC_KEY")
	recipient := os.Getenv("MG_RECIPIENT")
	sender := os.Getenv("MG_SENDER")
	title := os.Getenv("MG_TITLE")
	mg := mailgun.NewMailgun(yourdomain, key, publicAPIKey)
	message := mailgun.NewMessage(
		sender,
		title,
		body,
		recipient)
	resp, id, err := mg.Send(message)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID: %s Resp: %s\n", id, resp)
}

func handler(response http.ResponseWriter, request *http.Request) {
	name := request.FormValue("name")
	email := request.FormValue("email")
	message := request.FormValue("message")
	body := fmt.Sprintf("New requests\nName: %s\nEmail: %s\nMessage: %s\n", name, email, message)
	send(body)
	fmt.Fprintf(response, "{}")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9000", nil)
}
