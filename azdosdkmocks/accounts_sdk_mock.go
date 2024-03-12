// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Nataliia5722/azure-devops-go-api/azuredevops/v7/accounts (interfaces: Client)

// Package azdosdkmocks is a generated GoMock package.
package azdosdkmocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	accounts "github.com/Nataliia5722/azure-devops-go-api/azuredevops/v7/accounts"
	reflect "reflect"
)

// MockAccountsClient is a mock of Client interface
type MockAccountsClient struct {
	ctrl     *gomock.Controller
	recorder *MockAccountsClientMockRecorder
}

// MockAccountsClientMockRecorder is the mock recorder for MockAccountsClient
type MockAccountsClientMockRecorder struct {
	mock *MockAccountsClient
}

// NewMockAccountsClient creates a new mock instance
func NewMockAccountsClient(ctrl *gomock.Controller) *MockAccountsClient {
	mock := &MockAccountsClient{ctrl: ctrl}
	mock.recorder = &MockAccountsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAccountsClient) EXPECT() *MockAccountsClientMockRecorder {
	return m.recorder
}

// GetAccounts mocks base method
func (m *MockAccountsClient) GetAccounts(arg0 context.Context, arg1 accounts.GetAccountsArgs) (*[]accounts.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccounts", arg0, arg1)
	ret0, _ := ret[0].(*[]accounts.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccounts indicates an expected call of GetAccounts
func (mr *MockAccountsClientMockRecorder) GetAccounts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccounts", reflect.TypeOf((*MockAccountsClient)(nil).GetAccounts), arg0, arg1)
}
