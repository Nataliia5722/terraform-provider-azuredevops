// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Nataliia5722/azure-devops-go-api/azuredevops/v7/memberentitlementmanagement (interfaces: Client)

// Package azdosdkmocks is a generated GoMock package.
package azdosdkmocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	memberentitlementmanagement "github.com/Nataliia5722/azure-devops-go-api/azuredevops/v7/memberentitlementmanagement"
	reflect "reflect"
)

// MockMemberentitlementmanagementClient is a mock of Client interface
type MockMemberentitlementmanagementClient struct {
	ctrl     *gomock.Controller
	recorder *MockMemberentitlementmanagementClientMockRecorder
}

// MockMemberentitlementmanagementClientMockRecorder is the mock recorder for MockMemberentitlementmanagementClient
type MockMemberentitlementmanagementClientMockRecorder struct {
	mock *MockMemberentitlementmanagementClient
}

// NewMockMemberentitlementmanagementClient creates a new mock instance
func NewMockMemberentitlementmanagementClient(ctrl *gomock.Controller) *MockMemberentitlementmanagementClient {
	mock := &MockMemberentitlementmanagementClient{ctrl: ctrl}
	mock.recorder = &MockMemberentitlementmanagementClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMemberentitlementmanagementClient) EXPECT() *MockMemberentitlementmanagementClientMockRecorder {
	return m.recorder
}

// AddGroupEntitlement mocks base method
func (m *MockMemberentitlementmanagementClient) AddGroupEntitlement(arg0 context.Context, arg1 memberentitlementmanagement.AddGroupEntitlementArgs) (*memberentitlementmanagement.GroupEntitlementOperationReference, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddGroupEntitlement", arg0, arg1)
	ret0, _ := ret[0].(*memberentitlementmanagement.GroupEntitlementOperationReference)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddGroupEntitlement indicates an expected call of AddGroupEntitlement
func (mr *MockMemberentitlementmanagementClientMockRecorder) AddGroupEntitlement(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddGroupEntitlement", reflect.TypeOf((*MockMemberentitlementmanagementClient)(nil).AddGroupEntitlement), arg0, arg1)
}

// AddMemberToGroup mocks base method
func (m *MockMemberentitlementmanagementClient) AddMemberToGroup(arg0 context.Context, arg1 memberentitlementmanagement.AddMemberToGroupArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddMemberToGroup", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddMemberToGroup indicates an expected call of AddMemberToGroup
func (mr *MockMemberentitlementmanagementClientMockRecorder) AddMemberToGroup(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMemberToGroup", reflect.TypeOf((*MockMemberentitlementmanagementClient)(nil).AddMemberToGroup), arg0, arg1)
}

// AddServicePrincipalEntitlement mocks base method
func (m *MockMemberentitlementmanagementClient) AddServicePrincipalEntitlement(arg0 context.Context, arg1 memberentitlementmanagement.AddServicePrincipalEntitlementArgs) (*memberentitlementmanagement.ServicePrincipalEntitlementsPostResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddServicePrincipalEntitlement", arg0, arg1)
	ret0, _ := ret[0].(*memberentitlementmanagement.ServicePrincipalEntitlementsPostResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddServicePrincipalEntitlement indicates an expected call of AddServicePrincipalEntitlement
func (mr *MockMemberentitlementmanagementClientMockRecorder) AddServicePrincipalEntitlement(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddServicePrincipalEntitlement", reflect.TypeOf((*MockMemberentitlementmanagementClient)(nil).AddServicePrincipalEntitlement), arg0, arg1)
}

// AddUserEntitlement mocks base method
func (m *MockMemberentitlementmanagementClient) AddUserEntitlement(arg0 context.Context, arg1 memberentitlementmanagement.AddUserEntitlementArgs) (*memberentitlementmanagement.UserEntitlementsPostResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUserEntitlement", arg0, arg1)
	ret0, _ := ret[0].(*memberentitlementmanagement.UserEntitlementsPostResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddUserEntitlement indicates an expected call of AddUserEntitlement
func (mr *MockMemberentitlementmanagementClientMockRecorder) AddUserEntitlement(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUserEntitlement", reflect.TypeOf((*MockMemberentitlementmanagementClient)(nil).AddUserEntitlement), arg0, arg1)
}

// DeleteGroupEntitlement mocks base method
func (m *MockMemberentitlementmanagementClient) DeleteGroupEntitlement(arg0 context.Context, arg1 memberentitlementmanagement.DeleteGroupEntitlementArgs) (*memberentitlementmanagement.GroupEntitlementOperationReference, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteGroupEntitlement", arg0, arg1)
	ret0, _ := ret[0].(*memberentitlementmanagement.GroupEntitlementOperationReference)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteGroupEntitlement indicates an expected call of DeleteGroupEntitlement
func (mr *MockMemberentitlementmanagementClientMockRecorder) DeleteGroupEntitlement(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteGroupEntitlement", reflect.TypeOf((*MockMemberentitlementmanagementClient)(nil).DeleteGroupEntitlement), arg0, arg1)
}

// DeleteServicePrincipalEntitlement mocks base method
func (m *MockMemberentitlementmanagementClient) DeleteServicePrincipalEntitlement(arg0 context.Context, arg1 memberentitlementmanagement.DeleteServicePrincipalEntitlementArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteServicePrincipalEntitlement", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteServicePrincipalEntitlement indicates an expected call of DeleteServicePrincipalEntitlement
func (mr *MockMemberentitlementmanagementClientMockRecorder) DeleteServicePrincipalEntitlement(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteServicePrincipalEntitlement", reflect.TypeOf((*MockMemberentitlementmanagementClient)(nil).DeleteServicePrincipalEntitlement), arg0, arg1)
}

// DeleteUserEntitlement mocks base method
func (m *MockMemberentitlementmanagementClient) DeleteUserEntitlement(arg0 context.Context, arg1 memberentitlementmanagement.DeleteUserEntitlementArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUserEntitlement", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUserEntitlement indicates an expected call of DeleteUserEntitlement
func (mr *MockMemberentitlementmanagementClientMockRecorder) DeleteUserEntitlement(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUserEntitlement", reflect.TypeOf((*MockMemberentitlementmanagementClient)(nil).DeleteUserEntitlement), arg0, arg1)
}

// GetGroupEntitlement mocks base method
func (m *MockMemberentitlementmanagementClient) GetGroupEntitlement(arg0 context.Context, arg1 memberentitlementmanagement.GetGroupEntitlementArgs) (*memberentitlementmanagement.GroupEntitlement, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGroupEntitlement", arg0, arg1)
	ret0, _ := ret[0].(*memberentitlementmanagement.GroupEntitlement)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGroupEntitlement indicates an expected call of GetGroupEntitlement
func (mr *MockMemberentitlementmanagementClientMockRecorder) GetGroupEntitlement(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroupEntitlement", reflect.TypeOf((*MockMemberentitlementmanagementClient)(nil).GetGroupEntitlement), arg0, arg1)
}

// GetGroupEntitlements mocks base method
func (m *MockMemberentitlementmanagementClient) GetGroupEntitlements(arg0 context.Context, arg1 memberentitlementmanagement.GetGroupEntitlementsArgs) (*[]memberentitlementmanagement.GroupEntitlement, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGroupEntitlements", arg0, arg1)
	ret0, _ := ret[0].(*[]memberentitlementmanagement.GroupEntitlement)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGroupEntitlements indicates an expected call of GetGroupEntitlements
func (mr *MockMemberentitlementmanagementClientMockRecorder) GetGroupEntitlements(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroupEntitlements", reflect.TypeOf((*MockMemberentitlementmanagementClient)(nil).GetGroupEntitlements), arg0, arg1)
}

// GetGroupMembers mocks base method
func (m *MockMemberentitlementmanagementClient) GetGroupMembers(arg0 context.Context, arg1 memberentitlementmanagement.GetGroupMembersArgs) (*memberentitlementmanagement.PagedGraphMemberList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGroupMembers", arg0, arg1)
	ret0, _ := ret[0].(*memberentitlementmanagement.PagedGraphMemberList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGroupMembers indicates an expected call of GetGroupMembers
func (mr *MockMemberentitlementmanagementClientMockRecorder) GetGroupMembers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroupMembers", reflect.TypeOf((*MockMemberentitlementmanagementClient)(nil).GetGroupMembers), arg0, arg1)
}

// GetServicePrincipalEntitlement mocks base method
func (m *MockMemberentitlementmanagementClient) GetServicePrincipalEntitlement(arg0 context.Context, arg1 memberentitlementmanagement.GetServicePrincipalEntitlementArgs) (*memberentitlementmanagement.ServicePrincipalEntitlement, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServicePrincipalEntitlement", arg0, arg1)
	ret0, _ := ret[0].(*memberentitlementmanagement.ServicePrincipalEntitlement)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetServicePrincipalEntitlement indicates an expected call of GetServicePrincipalEntitlement
func (mr *MockMemberentitlementmanagementClientMockRecorder) GetServicePrincipalEntitlement(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServicePrincipalEntitlement", reflect.TypeOf((*MockMemberentitlementmanagementClient)(nil).GetServicePrincipalEntitlement), arg0, arg1)
}

// GetUserEntitlement mocks base method
func (m *MockMemberentitlementmanagementClient) GetUserEntitlement(arg0 context.Context, arg1 memberentitlementmanagement.GetUserEntitlementArgs) (*memberentitlementmanagement.UserEntitlement, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserEntitlement", arg0, arg1)
	ret0, _ := ret[0].(*memberentitlementmanagement.UserEntitlement)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserEntitlement indicates an expected call of GetUserEntitlement
func (mr *MockMemberentitlementmanagementClientMockRecorder) GetUserEntitlement(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserEntitlement", reflect.TypeOf((*MockMemberentitlementmanagementClient)(nil).GetUserEntitlement), arg0, arg1)
}

// GetUsersSummary mocks base method
func (m *MockMemberentitlementmanagementClient) GetUsersSummary(arg0 context.Context, arg1 memberentitlementmanagement.GetUsersSummaryArgs) (*memberentitlementmanagement.UsersSummary, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsersSummary", arg0, arg1)
	ret0, _ := ret[0].(*memberentitlementmanagement.UsersSummary)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsersSummary indicates an expected call of GetUsersSummary
func (mr *MockMemberentitlementmanagementClientMockRecorder) GetUsersSummary(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsersSummary", reflect.TypeOf((*MockMemberentitlementmanagementClient)(nil).GetUsersSummary), arg0, arg1)
}

// RemoveMemberFromGroup mocks base method
func (m *MockMemberentitlementmanagementClient) RemoveMemberFromGroup(arg0 context.Context, arg1 memberentitlementmanagement.RemoveMemberFromGroupArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveMemberFromGroup", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveMemberFromGroup indicates an expected call of RemoveMemberFromGroup
func (mr *MockMemberentitlementmanagementClientMockRecorder) RemoveMemberFromGroup(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveMemberFromGroup", reflect.TypeOf((*MockMemberentitlementmanagementClient)(nil).RemoveMemberFromGroup), arg0, arg1)
}

// SearchMemberEntitlements mocks base method
func (m *MockMemberentitlementmanagementClient) SearchMemberEntitlements(arg0 context.Context, arg1 memberentitlementmanagement.SearchMemberEntitlementsArgs) (*[]memberentitlementmanagement.MemberEntitlement2, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchMemberEntitlements", arg0, arg1)
	ret0, _ := ret[0].(*[]memberentitlementmanagement.MemberEntitlement2)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchMemberEntitlements indicates an expected call of SearchMemberEntitlements
func (mr *MockMemberentitlementmanagementClientMockRecorder) SearchMemberEntitlements(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchMemberEntitlements", reflect.TypeOf((*MockMemberentitlementmanagementClient)(nil).SearchMemberEntitlements), arg0, arg1)
}

// SearchUserEntitlements mocks base method
func (m *MockMemberentitlementmanagementClient) SearchUserEntitlements(arg0 context.Context, arg1 memberentitlementmanagement.SearchUserEntitlementsArgs) (*memberentitlementmanagement.PagedGraphMemberList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchUserEntitlements", arg0, arg1)
	ret0, _ := ret[0].(*memberentitlementmanagement.PagedGraphMemberList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchUserEntitlements indicates an expected call of SearchUserEntitlements
func (mr *MockMemberentitlementmanagementClientMockRecorder) SearchUserEntitlements(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchUserEntitlements", reflect.TypeOf((*MockMemberentitlementmanagementClient)(nil).SearchUserEntitlements), arg0, arg1)
}

// UpdateGroupEntitlement mocks base method
func (m *MockMemberentitlementmanagementClient) UpdateGroupEntitlement(arg0 context.Context, arg1 memberentitlementmanagement.UpdateGroupEntitlementArgs) (*memberentitlementmanagement.GroupEntitlementOperationReference, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateGroupEntitlement", arg0, arg1)
	ret0, _ := ret[0].(*memberentitlementmanagement.GroupEntitlementOperationReference)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateGroupEntitlement indicates an expected call of UpdateGroupEntitlement
func (mr *MockMemberentitlementmanagementClientMockRecorder) UpdateGroupEntitlement(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateGroupEntitlement", reflect.TypeOf((*MockMemberentitlementmanagementClient)(nil).UpdateGroupEntitlement), arg0, arg1)
}

// UpdateServicePrincipalEntitlement mocks base method
func (m *MockMemberentitlementmanagementClient) UpdateServicePrincipalEntitlement(arg0 context.Context, arg1 memberentitlementmanagement.UpdateServicePrincipalEntitlementArgs) (*memberentitlementmanagement.ServicePrincipalEntitlementsPatchResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateServicePrincipalEntitlement", arg0, arg1)
	ret0, _ := ret[0].(*memberentitlementmanagement.ServicePrincipalEntitlementsPatchResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateServicePrincipalEntitlement indicates an expected call of UpdateServicePrincipalEntitlement
func (mr *MockMemberentitlementmanagementClientMockRecorder) UpdateServicePrincipalEntitlement(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateServicePrincipalEntitlement", reflect.TypeOf((*MockMemberentitlementmanagementClient)(nil).UpdateServicePrincipalEntitlement), arg0, arg1)
}

// UpdateServicePrincipalEntitlements mocks base method
func (m *MockMemberentitlementmanagementClient) UpdateServicePrincipalEntitlements(arg0 context.Context, arg1 memberentitlementmanagement.UpdateServicePrincipalEntitlementsArgs) (*memberentitlementmanagement.ServicePrincipalEntitlementOperationReference, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateServicePrincipalEntitlements", arg0, arg1)
	ret0, _ := ret[0].(*memberentitlementmanagement.ServicePrincipalEntitlementOperationReference)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateServicePrincipalEntitlements indicates an expected call of UpdateServicePrincipalEntitlements
func (mr *MockMemberentitlementmanagementClientMockRecorder) UpdateServicePrincipalEntitlements(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateServicePrincipalEntitlements", reflect.TypeOf((*MockMemberentitlementmanagementClient)(nil).UpdateServicePrincipalEntitlements), arg0, arg1)
}

// UpdateUserEntitlement mocks base method
func (m *MockMemberentitlementmanagementClient) UpdateUserEntitlement(arg0 context.Context, arg1 memberentitlementmanagement.UpdateUserEntitlementArgs) (*memberentitlementmanagement.UserEntitlementsPatchResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserEntitlement", arg0, arg1)
	ret0, _ := ret[0].(*memberentitlementmanagement.UserEntitlementsPatchResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUserEntitlement indicates an expected call of UpdateUserEntitlement
func (mr *MockMemberentitlementmanagementClientMockRecorder) UpdateUserEntitlement(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserEntitlement", reflect.TypeOf((*MockMemberentitlementmanagementClient)(nil).UpdateUserEntitlement), arg0, arg1)
}

// UpdateUserEntitlements mocks base method
func (m *MockMemberentitlementmanagementClient) UpdateUserEntitlements(arg0 context.Context, arg1 memberentitlementmanagement.UpdateUserEntitlementsArgs) (*memberentitlementmanagement.UserEntitlementOperationReference, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserEntitlements", arg0, arg1)
	ret0, _ := ret[0].(*memberentitlementmanagement.UserEntitlementOperationReference)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUserEntitlements indicates an expected call of UpdateUserEntitlements
func (mr *MockMemberentitlementmanagementClientMockRecorder) UpdateUserEntitlements(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserEntitlements", reflect.TypeOf((*MockMemberentitlementmanagementClient)(nil).UpdateUserEntitlements), arg0, arg1)
}
