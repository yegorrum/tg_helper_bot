package main

import (
	"flag"
	"log"
	tgClient "tg_bot/clients/telegram"
	event_consumer "tg_bot/consumer/event-consumer"
	tgEvent "tg_bot/events/telegram"
	"tg_bot/storage/files"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "file_storage"
	batchSize   = 100
)

// 8565782509:AAF5gy1a4bsFvS0_AV5KTH1Y6smTEheJ-xY
func main() {
	token := mustToken()

	eventsProcessor := tgEvent.New(
		tgClient.New(tgBotHost, token),
		files.New(storagePath),
	)

	log.Print("service started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stoped", err)
	}
}

func mustToken() string {
	token := flag.String("token", "", "token for access to tg bot")

	flag.Parse()
	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}
