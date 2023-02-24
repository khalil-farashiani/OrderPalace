package contract

import "github.com/khalil-farashiani/OrderPalace/receiver/internal/entities"

type BrokerIntractor interface {
	ReceiveOrder(*entities.Order) error
}
