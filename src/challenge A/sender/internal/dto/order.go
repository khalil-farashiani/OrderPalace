package dto

import "github.com/khalil-farashiani/OrderPalace/sender/internal/entities"

type SendOrderRequest struct {
	entities.Order
}

type SendOrderResponse struct {
	Message string `json:"message"`
}
