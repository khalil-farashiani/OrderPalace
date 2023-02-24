package interactor

import (
	"context"
	"fmt"
	"github.com/go-faker/faker/v4"
	"github.com/golang/mock/gomock"
	"github.com/khalil-farashiani/OrderPalace/sender/internal/contract"
	"github.com/khalil-farashiani/OrderPalace/sender/internal/dto"
	"github.com/khalil-farashiani/OrderPalace/sender/internal/entities"
	orderInteractor "github.com/khalil-farashiani/OrderPalace/sender/internal/interactor/order"
	mockContract "github.com/khalil-farashiani/OrderPalace/sender/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	successMessageOnAddOrder = "Order Successfully Push"
)

func setup(t *testing.T) (contract.OrderIntractor, *mockContract.MockBrokerIntractor, func()) {

	mockCtrl := gomock.NewController(t)
	mockUserStore := mockContract.NewMockBrokerIntractor(mockCtrl)

	service := orderInteractor.New(mockUserStore)
	return service, mockUserStore, func() {
		mockCtrl.Finish()
	}
}
func TestSendOrder(t *testing.T) {

	t.Run("fail on add order error", func(t *testing.T) {
		interactor, mockBroker, teardown := setup(t)
		defer teardown()
		randomNum, _ := faker.RandomInt(0, 2<<6, 2)
		req := dto.SendOrderRequest{
			ID:    uint64(randomNum[0]),
			Price: uint64(randomNum[1]),
			Title: faker.Name(),
		}

		ctx := context.Background()
		mockBroker.EXPECT().AddOrder(ctx, gomock.Any()).Return(fmt.Errorf(""))

		_, err := interactor.SendOrder(ctx, req)
		assert.NotNil(t, err)
	})

	t.Run("successful add order", func(t *testing.T) {
		interactor, mockBroker, teardown := setup(t)
		defer teardown()

		randomNum, _ := faker.RandomInt(0, 2<<6, 2)
		req := dto.SendOrderRequest{
			ID:    uint64(randomNum[0]),
			Price: uint64(randomNum[1]),
			Title: faker.Name(),
		}

		ctx := context.Background()

		passedOrder := entities.Order{
			ID:    req.ID,
			Price: req.Price,
			Title: req.Title,
		}

		mockBroker.EXPECT().AddOrder(ctx, &passedOrder).Return(nil)
		resp, err := interactor.SendOrder(ctx, req)
		assert.Nil(t, err)
		assert.EqualValues(t, successMessageOnAddOrder, resp.Message)
	})
}
