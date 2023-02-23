package order

import (
	"context"
	"github.com/khalil-farashiani/OrderPalace/sender/internal/contract"
	"github.com/khalil-farashiani/OrderPalace/sender/internal/dto"
	"github.com/khalil-farashiani/OrderPalace/sender/internal/entities"
)

type Interactor struct {
	broker contract.BrokerIntractor
}

func New(broker contract.BrokerIntractor) Interactor {
	return Interactor{broker: broker}
}

func (i Interactor) SendOrder(ctx context.Context, req dto.SendOrderRequest) (dto.SendOrderResponse, error) {
	order := entities.Order{
		ID:    req.ID,
		Price: req.Price,
		Title: req.Title,
	}
	if err := i.broker.AddOrder(ctx, &order); err != nil {
		return dto.SendOrderResponse{}, err
	}
	return dto.SendOrderResponse{Message: "Order Successfully Push"}, nil
}
