package main

import (
	"flag"
	"log"

	tgClient "github.com/yegorrum/tg_helper_bot/clients/telegram"
	event_consumer "github.com/yegorrum/tg_helper_bot/consumer/event-consumer"
	tgEvent "github.com/yegorrum/tg_helper_bot/events/telegram"
	"github.com/yegorrum/tg_helper_bot/storage/files"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "file_storage"
	batchSize   = 100
)

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
