package contract

import (
	"context"
	"github.com/khalil-farashiani/OrderPalace/receiver/internal/entities"
)

type OrderStore interface {
	CreateOrder(ctx context.Context, order entities.Order) (entities.Order, error)
}
