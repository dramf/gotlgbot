package tlgbot

import "net/http"

func (bot *tlgBot) echo(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Echoooo"))
}
