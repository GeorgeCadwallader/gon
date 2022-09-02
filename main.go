package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/joho/godotenv"
	"github.com/nlopes/slack"
)

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	http.HandleFunc("/recieve", slashCommandHandler)

	fmt.Println("[INFO] Server listening")
	http.ListenAndServe(":8080", nil)

}

func slashCommandHandler(w http.ResponseWriter, r *http.Request) {

	s, err := slack.SlashCommandParse(r)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	switch s.Command {
	case "/gonline":

		var uri string = s.Text

		u, err := url.ParseRequestURI(uri)

		if err != nil {
			w.Write([]byte("This is an invalid url dawg"))
			return
		}

		var response string = uri

		resp, err := http.Get(u.String())

		if err != nil {
			response = response + " could not be reached!"
		} else if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
			response = response + " is up!"
		} else {
			response = response + " is down!"
		}

		w.Write([]byte(response))

	default:
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
