package main

import (
	"log"
	"os"

	"github.com/nurudianto/polantas/handler"
	"github.com/subosito/gotenv"
	"github.com/yanzay/tbot"
)

func main() {
	gotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	token := "781127446:AAHnr_BRiwHIRfMRTJ7mP1x1A0Wa93JbzHI"
	opt := tbot.WithWebhook("", ":"+port)
	bot, err := tbot.NewServer(token, opt)
	if err != nil {
		log.Fatal(err)
	}
	bot.Handle("/healthz", "ok")
	bot.HandleFunc("/start", handler.HandleStart)
	bot.HandleFunc("/next", handler.HandleNext)
	bot.HandleFunc("/reset", handler.HandleReset)
	bot.HandleFunc("/help", handler.HandleHelp)
	bot.HandleFunc("/bot_bodoh", handler.HandleBodoh)
	bot.HandleFunc("/cuti", handler.HandleCuti)
	bot.HandleFunc("/skip", handler.HandleSkip)
	bot.ListenAndServe()
}
