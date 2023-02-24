package contract

import (
	"context"
	"github.com/khalil-farashiani/OrderPalace/receiver/internal/entities"
)

type UserStore interface {
	CreateUser(ctx context.Context, order entities.Order) (entities.Order, error)
}
