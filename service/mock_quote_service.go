// Code generated by MockGen. DO NOT EDIT.
// Source: service/quote_service_interface.go

// Package service is a generated GoMock package.
package service

import (
	reflect "reflect"

	model "github.com/dilaragorum/api-quote-reverser/model"
	gomock "github.com/golang/mock/gomock"
)

// MockIQuoteService is a mock of IQuoteService interface.
type MockIQuoteService struct {
	ctrl     *gomock.Controller
	recorder *MockIQuoteServiceMockRecorder
}

// MockIQuoteServiceMockRecorder is the mock recorder for MockIQuoteService.
type MockIQuoteServiceMockRecorder struct {
	mock *MockIQuoteService
}

// NewMockIQuoteService creates a new mock instance.
func NewMockIQuoteService(ctrl *gomock.Controller) *MockIQuoteService {
	mock := &MockIQuoteService{ctrl: ctrl}
	mock.recorder = &MockIQuoteServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIQuoteService) EXPECT() *MockIQuoteServiceMockRecorder {
	return m.recorder
}

// GetQuotesReversed mocks base method.
func (m *MockIQuoteService) GetQuotesReversed() ([]model.GroupQuote, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetQuotesReversed")
	ret0, _ := ret[0].([]model.GroupQuote)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetQuotesReversed indicates an expected call of GetQuotesReversed.
func (mr *MockIQuoteServiceMockRecorder) GetQuotesReversed() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQuotesReversed", reflect.TypeOf((*MockIQuoteService)(nil).GetQuotesReversed))
}
