package dto

import "github.com/khalil-farashiani/OrderPalace/receiver/internal/entities"

type OrderConsume struct {
	ID    uint64 `json:"order_id"`
	Price uint64 `json:"price"`
	Title string `json:"title"`
}

type OrderSetLog struct {
	entities.Order
}
