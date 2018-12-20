package main

import (
	"log"
	"os"

	"github.com/heroku/polantas/handler"
	"github.com/subosito/gotenv"
	"github.com/yanzay/tbot"
)

func main() {
	gotenv.Load()
	token := os.Getenv("TELEGRAM_TOKEN")
	opt := tbot.WithWebhook("", ":8080")
	bot, err := tbot.NewServer(token, opt)
	if err != nil {
		log.Fatal(err)
	}
	bot.Handle("/healthz", "ok")
	bot.HandleFunc("/mulai", handler.HandleStart)
	bot.HandleFunc("/selesai", handler.HandleFinish)
	bot.HandleFunc("/reset", handler.HandleReset)
	bot.ListenAndServe()
}
