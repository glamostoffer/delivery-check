package natsutil

import (
	"github.com/nats-io/nats.go"
	"log"
)

func MessageHandler(m *nats.Msg) {
	log.Printf("Received a message: %s", string(m.Data))
}
