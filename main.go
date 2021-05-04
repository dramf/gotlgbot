package main

import (
	"github.com/dramf/gotlgbot/tlgbot"
	"log"
)

func main() {
	port := ":8082"
	log.Print("Starting Go Telegram Bot")
	tlgbot.Instance().Listen(port)
}
