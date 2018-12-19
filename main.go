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
	bot, err := tbot.NewServer(os.Getenv("TELEGRAM_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}
	bot.Handle("/healthz", "ok")
	bot.HandleFunc("/mulai", handler.HandleStart)
	bot.HandleFunc("/selesai", handler.HandleFinish)
	bot.HandleFunc("/reset", handler.HandleReset)
	bot.ListenAndServe()
}
