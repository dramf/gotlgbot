package tlgbot

import (
	"io/ioutil"
	"log"
	"net/http"
)

func (bot *tlgBot) webhook(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("read of a request body error: %v", err)
		return
	}
	log.Printf("Webhook: %s", string(b))
	w.WriteHeader(http.StatusOK)
}
