package dto

import (
	"errors"
	"strings"
)

type SendOrderRequest struct {
	ID    uint64 `json:"order_id"`
	Price uint64 `json:"price"`
	Title string `json:"title"`
}

type SendOrderResponse struct {
	Message string `json:"message"`
}

const (
	requiredIdErrorMsg    = "id is required field"
	requiredPriceErrorMsg = "price is required field and should be positive integer"
	requiredTitleErrorMsg = "title cannot be empty"
)

func (s SendOrderRequest) Validate() error {
	switch {
	case s.ID == 0:
		return errors.New(requiredIdErrorMsg)
	case s.Price <= 0:
		return errors.New(requiredPriceErrorMsg)
	case strings.TrimSpace(s.Title) == "":
		return errors.New(requiredTitleErrorMsg)
	}
	return nil
}
