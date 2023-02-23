package contract

import (
	"context"
	"github.com/khalil-farashiani/OrderPalace/sender/internal/entities"
)

type BrokerIntractor interface {
	AddOrder(ctx context.Context, order *entities.Order) error
}
