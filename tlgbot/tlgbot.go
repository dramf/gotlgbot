package tlgbot

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type tlgBot struct {
	router *mux.Router
}

func Instance() *tlgBot {
	r := &tlgBot{
		router: mux.NewRouter().StrictSlash(true),
	}
	r.router.HandleFunc("/echo", r.echo).Methods("GET")
	r.router.HandleFunc("/webhook", r.webhook).Methods("POST")
	return r
}

func (bot *tlgBot) Listen(url string) {
	log.Println("Listening at", url)
	log.Fatal(http.ListenAndServe(url, bot.router))
}
