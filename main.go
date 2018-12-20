package main

import (
	"fmt"
	"log"
	"os"

	"github.com/heroku/polantas/handler"
	"github.com/subosito/gotenv"
	"github.com/yanzay/tbot"
)

func main() {
	gotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	token := os.Getenv("TELEGRAM_TOKEN")
	opt := tbot.WithWebhook("", ":"+port)
	fmt.Println("port=" + port)
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
