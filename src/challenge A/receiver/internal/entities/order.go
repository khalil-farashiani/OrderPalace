package entities

type Order struct {
	ID    uint64 `json:"order_id"`
	Price uint64 `json:"price"`
	Title string `json:"title"`
}
