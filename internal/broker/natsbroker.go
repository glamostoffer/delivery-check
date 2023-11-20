package broker

import (
	"deliveryCheck/pkg/natsconnector"
	"github.com/nats-io/nats.go"
)

type Nats interface {
	//Start(ctx context.Context)
	Subscribe(channel string) error
	Publish(channel, message string) error
}

type NatsUsecase interface {
	HandleMessage(m *nats.Msg)
}

type natsBroker struct {
	connector natsconnector.NatsConnector
	uc        NatsUsecase
}

func NewNatsBroker(conn natsconnector.NatsConnector, uc NatsUsecase) Nats {
	return &natsBroker{
		connector: conn,
		uc:        uc,
	}
}

func (nb *natsBroker) Subscribe(channel string) error {
	_, err := nb.connector.Client.Subscribe(channel, nb.uc.HandleMessage)
	if err != nil {
		return err
	}

	return nil
}

func (nb *natsBroker) Publish(channel, message string) error {
	err := nb.connector.Client.Publish(channel, []byte(message))

	return err
}
