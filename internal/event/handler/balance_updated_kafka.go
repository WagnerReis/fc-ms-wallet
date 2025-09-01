package handler

import (
	"fmt"
	"sync"

	"github.com/WagnerReis/fc-ms-wallet/pkg/events"
	"github.com/WagnerReis/fc-ms-wallet/pkg/kafka"
)

type UpdateBalanceKafkaHandler struct {
	Kafka *kafka.Producer
}

func NewBalanceUpdatedKafkaHandler(kafka *kafka.Producer) *UpdateBalanceKafkaHandler {
	return &UpdateBalanceKafkaHandler{
		Kafka: kafka,
	}
}

func (h *UpdateBalanceKafkaHandler) Handle(message events.EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()
	h.Kafka.Publish(message, nil, "balances")
	fmt.Println("Message sent to topic", message.GetName())
}
