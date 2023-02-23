package entities

type Order struct {
	OrderID uint64 `json:"order_id"`
	Price   uint64 `json:"price"`
	Title   string `json:"title"`
}
