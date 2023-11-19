package broker

import (
	"deliveryCheck/pkg/natsconnector"
	"github.com/nats-io/nats.go"
)

type Nats interface {
	//Start(ctx context.Context)
	Subscribe(channelName string) error
	Publish(message []byte) error
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

func (nb *natsBroker) Subscribe(channelName string) error {
	_, err := nb.connector.Client.Subscribe(channelName, nb.uc.HandleMessage)
	if err != nil {
		return err
	}

	return nil
}

func (nb *natsBroker) Publish(message []byte) error {
	return nil
}
