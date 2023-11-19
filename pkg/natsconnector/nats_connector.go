package natsconnector

import (
	"deliveryCheck/app"
	"github.com/nats-io/nats.go"
)

type NatsConnector struct {
	Client *nats.Conn
}

func NewNatsConnector(env *app.Env) (NatsConnector, error) {
	nc, err := nats.Connect(env.NatsServer)
	if err != nil {
		return NatsConnector{}, nil
	}

	return NatsConnector{Client: nc}, err
}

func (nc *NatsConnector) CloseConnection() {
	nc.Client.Close()
}
