package order

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/khalil-farashiani/OrderPalace/receiver/internal/contract"
	"github.com/khalil-farashiani/OrderPalace/receiver/internal/entities"
	"log"
)

type Interactor struct {
	broker contract.BrokerIntractor
	store  contract.OrderStore
}

func New(broker contract.BrokerIntractor, store contract.OrderStore) Interactor {
	return Interactor{
		broker: broker,
		store:  store,
	}
}

func (i Interactor) CreateOrder(ctx context.Context) {
	orderChan := make(chan []byte)
	log.Println("start receiving orders to store")
	go func() {
		defer close(orderChan)
		i.broker.ReceiveOrder(orderChan)
	}()

	for data := range orderChan {
		var order entities.Order
		if err := json.Unmarshal(data, &order); err != nil {
			fmt.Println("Failed to unmarshal order:", err)
		} else {
			if _, err := i.store.CreateOrder(ctx, order); err != nil {
				fmt.Println(err.Error())
			}
		}
	}
}
