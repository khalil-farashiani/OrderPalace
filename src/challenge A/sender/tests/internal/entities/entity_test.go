package entities

import (
	"github.com/khalil-farashiani/OrderPalace/sender/internal/entities"
	"testing"
)

func TestOrder(t *testing.T) {
	// Create an example order
	order := entities.Order{ID: 1, Price: 1000, Title: "Example Order"}

	// Test the order ID
	if order.ID != 1 {
		t.Errorf("Order ID is incorrect, expected 1 but got %d", order.ID)
	}

	// Test the order price
	if order.Price != 1000 {
		t.Errorf("Order price is incorrect, expected 1000 but got %d", order.Price)
	}

	// Test the order title
	if order.Title != "Example Order" {
		t.Errorf("Order title is incorrect, expected 'Example Order' but got '%s'", order.Title)
	}
}
