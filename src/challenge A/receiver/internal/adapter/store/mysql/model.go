package mysql

import "github.com/khalil-farashiani/OrderPalace/receiver/internal/entities"

type Order struct {
	ID    uint64 `json:"order_id" gorm:"primary_key"`
	Price uint64 `json:"price"`
	Title string `json:"title"`
}

func mapToOrderModel(o entities.Order) Order {
	return Order{
		Price: o.Price,
		Title: o.Title,
	}
}

func mapFromOrderModel(o Order) entities.Order {
	return entities.Order{
		ID:    o.ID,
		Price: o.Price,
		Title: o.Title,
	}
}
