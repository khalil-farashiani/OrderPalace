// Code generated by MockGen. DO NOT EDIT.
// Source: ../sender/internal/contract/broker_intractor.go

// Package mock_contract is a generated GoMock package.
package mock_contract

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entities "github.com/khalil-farashiani/OrderPalace/sender/internal/entities"
)

// MockBrokerIntractor is a mock of BrokerIntractor interface.
type MockBrokerIntractor struct {
	ctrl     *gomock.Controller
	recorder *MockBrokerIntractorMockRecorder
}

// MockBrokerIntractorMockRecorder is the mock recorder for MockBrokerIntractor.
type MockBrokerIntractorMockRecorder struct {
	mock *MockBrokerIntractor
}

// NewMockBrokerIntractor creates a new mock instance.
func NewMockBrokerIntractor(ctrl *gomock.Controller) *MockBrokerIntractor {
	mock := &MockBrokerIntractor{ctrl: ctrl}
	mock.recorder = &MockBrokerIntractorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBrokerIntractor) EXPECT() *MockBrokerIntractorMockRecorder {
	return m.recorder
}

// AddOrder mocks base method.
func (m *MockBrokerIntractor) AddOrder(ctx context.Context, order *entities.Order) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddOrder", ctx, order)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddOrder indicates an expected call of AddOrder.
func (mr *MockBrokerIntractorMockRecorder) AddOrder(ctx, order interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddOrder", reflect.TypeOf((*MockBrokerIntractor)(nil).AddOrder), ctx, order)
}
