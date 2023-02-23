package dto

import (
	"errors"
	"github.com/khalil-farashiani/OrderPalace/sender/internal/entities"
	"strings"
)

type SendOrderRequest struct {
	entities.Order
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
