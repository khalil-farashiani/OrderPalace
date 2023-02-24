package contract

import (
	"context"
)

type OrderInteractor interface {
	CreateOrder(context.Context)
}
