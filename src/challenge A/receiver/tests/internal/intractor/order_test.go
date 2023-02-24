package intractor

import (
	"context"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/khalil-farashiani/OrderPalace/receiver/internal/contract"
	"github.com/khalil-farashiani/OrderPalace/receiver/internal/entities"
	"github.com/khalil-farashiani/OrderPalace/receiver/internal/interactor/order"
	mock_contract "github.com/khalil-farashiani/OrderPalace/receiver/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func setup(t *testing.T) (contract.OrderInteractor, *mock_contract.MockOrderStore, *mock_contract.MockBrokerIntractor, func()) {
	mockCtrl := gomock.NewController(t)
	mockOrderStore := mock_contract.NewMockOrderStore(mockCtrl)
	mockBroker := mock_contract.NewMockBrokerIntractor(mockCtrl)

	service := order.New(mockBroker, mockOrderStore)
	return service, mockOrderStore, mockBroker, func() {
		mockCtrl.Finish()
	}
}

func TestInteractor_CreateOrder(t *testing.T) {

	service, mockOrderStore, mockBroker, tearDown := setup(t)
	defer tearDown()
	ctx := context.Background()

	order := entities.Order{
		ID:    1,
		Title: "Alice",
		Price: 10.0,
	}

	orderBytes, err := json.Marshal(order)
	assert.Nil(t, err)

	orderChan := make(chan []byte)
	go func() {
		orderChan <- orderBytes
		close(orderChan)
	}()

	mockBroker.EXPECT().ReceiveOrder(gomock.Any()).Do(func(ch chan []byte) {
		for data := range orderChan {
			ch <- data
		}
	})

	mockOrderStore.EXPECT().CreateOrder(ctx, order).Return(order, nil)
	service.CreateOrder(ctx)
}
