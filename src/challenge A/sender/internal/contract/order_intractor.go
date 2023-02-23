package contract

import (
	"context"
	"github.com/khalil-farashiani/OrderPalace/sender/internal/dto"
)

type OrderIntractor interface {
	SendOrder(context.Context, dto.SendOrderRequest) (dto.SendOrderResponse, error)
}
