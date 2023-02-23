package v1

import (
	"context"
	"encoding/json"
	"github.com/khalil-farashiani/OrderPalace/sender/internal/contract"
	"github.com/khalil-farashiani/OrderPalace/sender/internal/dto"
	"github.com/khalil-farashiani/OrderPalace/sender/internal/interactor/order"
	"log"
	"net/http"
)

func SendOrder(broker contract.BrokerIntractor) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req dto.SendOrderRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "invalid request payload", http.StatusBadRequest)
			return
		}
		if err := req.Validate(); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		ctx := context.Background()
		resp, err := order.New(broker).SendOrder(ctx, req)
		if err != nil {
			log.Printf("failed to add order to Redis queue: %v", err)
			http.Error(w, "failed to add order to queue", http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode(resp); err != nil {
			log.Printf("failed to encode response: %v", err)
			http.Error(w, "failed to encode response", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}
