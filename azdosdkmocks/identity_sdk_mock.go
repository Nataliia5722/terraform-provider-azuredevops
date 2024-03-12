// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Nataliia5722/azure-devops-go-api/azuredevops/v7/identity (interfaces: Client)

// Package azdosdkmocks is a generated GoMock package.
package azdosdkmocks

import (
	context "context"
	delegatedauthorization "github.com/Nataliia5722/azure-devops-go-api/azuredevops/v7/delegatedauthorization"
	identity "github.com/Nataliia5722/azure-devops-go-api/azuredevops/v7/identity"
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
	reflect "reflect"
)

// MockIdentityClient is a mock of Client interface
type MockIdentityClient struct {
	ctrl     *gomock.Controller
	recorder *MockIdentityClientMockRecorder
}

// MockIdentityClientMockRecorder is the mock recorder for MockIdentityClient
type MockIdentityClientMockRecorder struct {
	mock *MockIdentityClient
}

// NewMockIdentityClient creates a new mock instance
func NewMockIdentityClient(ctrl *gomock.Controller) *MockIdentityClient {
	mock := &MockIdentityClient{ctrl: ctrl}
	mock.recorder = &MockIdentityClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIdentityClient) EXPECT() *MockIdentityClientMockRecorder {
	return m.recorder
}

// AddMember mocks base method
func (m *MockIdentityClient) AddMember(arg0 context.Context, arg1 identity.AddMemberArgs) (*bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddMember", arg0, arg1)
	ret0, _ := ret[0].(*bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddMember indicates an expected call of AddMember
func (mr *MockIdentityClientMockRecorder) AddMember(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMember", reflect.TypeOf((*MockIdentityClient)(nil).AddMember), arg0, arg1)
}

// CreateGroups mocks base method
func (m *MockIdentityClient) CreateGroups(arg0 context.Context, arg1 identity.CreateGroupsArgs) (*[]identity.Identity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateGroups", arg0, arg1)
	ret0, _ := ret[0].(*[]identity.Identity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateGroups indicates an expected call of CreateGroups
func (mr *MockIdentityClientMockRecorder) CreateGroups(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateGroups", reflect.TypeOf((*MockIdentityClient)(nil).CreateGroups), arg0, arg1)
}

// CreateIdentity mocks base method
func (m *MockIdentityClient) CreateIdentity(arg0 context.Context, arg1 identity.CreateIdentityArgs) (*identity.Identity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateIdentity", arg0, arg1)
	ret0, _ := ret[0].(*identity.Identity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateIdentity indicates an expected call of CreateIdentity
func (mr *MockIdentityClientMockRecorder) CreateIdentity(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateIdentity", reflect.TypeOf((*MockIdentityClient)(nil).CreateIdentity), arg0, arg1)
}

// CreateOrBindWithClaims mocks base method
func (m *MockIdentityClient) CreateOrBindWithClaims(arg0 context.Context, arg1 identity.CreateOrBindWithClaimsArgs) (*identity.Identity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrBindWithClaims", arg0, arg1)
	ret0, _ := ret[0].(*identity.Identity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrBindWithClaims indicates an expected call of CreateOrBindWithClaims
func (mr *MockIdentityClientMockRecorder) CreateOrBindWithClaims(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrBindWithClaims", reflect.TypeOf((*MockIdentityClient)(nil).CreateOrBindWithClaims), arg0, arg1)
}

// CreateScope mocks base method
func (m *MockIdentityClient) CreateScope(arg0 context.Context, arg1 identity.CreateScopeArgs) (*identity.IdentityScope, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateScope", arg0, arg1)
	ret0, _ := ret[0].(*identity.IdentityScope)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateScope indicates an expected call of CreateScope
func (mr *MockIdentityClientMockRecorder) CreateScope(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateScope", reflect.TypeOf((*MockIdentityClient)(nil).CreateScope), arg0, arg1)
}

// DeleteGroup mocks base method
func (m *MockIdentityClient) DeleteGroup(arg0 context.Context, arg1 identity.DeleteGroupArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteGroup", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteGroup indicates an expected call of DeleteGroup
func (mr *MockIdentityClientMockRecorder) DeleteGroup(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteGroup", reflect.TypeOf((*MockIdentityClient)(nil).DeleteGroup), arg0, arg1)
}

// DeleteScope mocks base method
func (m *MockIdentityClient) DeleteScope(arg0 context.Context, arg1 identity.DeleteScopeArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteScope", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteScope indicates an expected call of DeleteScope
func (mr *MockIdentityClientMockRecorder) DeleteScope(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteScope", reflect.TypeOf((*MockIdentityClient)(nil).DeleteScope), arg0, arg1)
}

// ForceRemoveMember mocks base method
func (m *MockIdentityClient) ForceRemoveMember(arg0 context.Context, arg1 identity.ForceRemoveMemberArgs) (*bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForceRemoveMember", arg0, arg1)
	ret0, _ := ret[0].(*bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ForceRemoveMember indicates an expected call of ForceRemoveMember
func (mr *MockIdentityClientMockRecorder) ForceRemoveMember(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForceRemoveMember", reflect.TypeOf((*MockIdentityClient)(nil).ForceRemoveMember), arg0, arg1)
}

// GetDescriptorById mocks base method
func (m *MockIdentityClient) GetDescriptorById(arg0 context.Context, arg1 identity.GetDescriptorByIdArgs) (*string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDescriptorById", arg0, arg1)
	ret0, _ := ret[0].(*string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDescriptorById indicates an expected call of GetDescriptorById
func (mr *MockIdentityClientMockRecorder) GetDescriptorById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDescriptorById", reflect.TypeOf((*MockIdentityClient)(nil).GetDescriptorById), arg0, arg1)
}

// GetIdentityChanges mocks base method
func (m *MockIdentityClient) GetIdentityChanges(arg0 context.Context, arg1 identity.GetIdentityChangesArgs) (*identity.ChangedIdentities, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIdentityChanges", arg0, arg1)
	ret0, _ := ret[0].(*identity.ChangedIdentities)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIdentityChanges indicates an expected call of GetIdentityChanges
func (mr *MockIdentityClientMockRecorder) GetIdentityChanges(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIdentityChanges", reflect.TypeOf((*MockIdentityClient)(nil).GetIdentityChanges), arg0, arg1)
}

// GetIdentitySnapshot mocks base method
func (m *MockIdentityClient) GetIdentitySnapshot(arg0 context.Context, arg1 identity.GetIdentitySnapshotArgs) (*identity.IdentitySnapshot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIdentitySnapshot", arg0, arg1)
	ret0, _ := ret[0].(*identity.IdentitySnapshot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIdentitySnapshot indicates an expected call of GetIdentitySnapshot
func (mr *MockIdentityClientMockRecorder) GetIdentitySnapshot(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIdentitySnapshot", reflect.TypeOf((*MockIdentityClient)(nil).GetIdentitySnapshot), arg0, arg1)
}

// GetMaxSequenceId mocks base method
func (m *MockIdentityClient) GetMaxSequenceId(arg0 context.Context, arg1 identity.GetMaxSequenceIdArgs) (*uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMaxSequenceId", arg0, arg1)
	ret0, _ := ret[0].(*uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMaxSequenceId indicates an expected call of GetMaxSequenceId
func (mr *MockIdentityClientMockRecorder) GetMaxSequenceId(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMaxSequenceId", reflect.TypeOf((*MockIdentityClient)(nil).GetMaxSequenceId), arg0, arg1)
}

// GetScopeById mocks base method
func (m *MockIdentityClient) GetScopeById(arg0 context.Context, arg1 identity.GetScopeByIdArgs) (*identity.IdentityScope, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetScopeById", arg0, arg1)
	ret0, _ := ret[0].(*identity.IdentityScope)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetScopeById indicates an expected call of GetScopeById
func (mr *MockIdentityClientMockRecorder) GetScopeById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetScopeById", reflect.TypeOf((*MockIdentityClient)(nil).GetScopeById), arg0, arg1)
}

// GetScopeByName mocks base method
func (m *MockIdentityClient) GetScopeByName(arg0 context.Context, arg1 identity.GetScopeByNameArgs) (*identity.IdentityScope, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetScopeByName", arg0, arg1)
	ret0, _ := ret[0].(*identity.IdentityScope)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetScopeByName indicates an expected call of GetScopeByName
func (mr *MockIdentityClientMockRecorder) GetScopeByName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetScopeByName", reflect.TypeOf((*MockIdentityClient)(nil).GetScopeByName), arg0, arg1)
}

// GetSelf mocks base method
func (m *MockIdentityClient) GetSelf(arg0 context.Context, arg1 identity.GetSelfArgs) (*identity.IdentitySelf, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSelf", arg0, arg1)
	ret0, _ := ret[0].(*identity.IdentitySelf)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSelf indicates an expected call of GetSelf
func (mr *MockIdentityClientMockRecorder) GetSelf(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSelf", reflect.TypeOf((*MockIdentityClient)(nil).GetSelf), arg0, arg1)
}

// GetSignedInToken mocks base method
func (m *MockIdentityClient) GetSignedInToken(arg0 context.Context, arg1 identity.GetSignedInTokenArgs) (*delegatedauthorization.AccessTokenResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSignedInToken", arg0, arg1)
	ret0, _ := ret[0].(*delegatedauthorization.AccessTokenResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSignedInToken indicates an expected call of GetSignedInToken
func (mr *MockIdentityClientMockRecorder) GetSignedInToken(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSignedInToken", reflect.TypeOf((*MockIdentityClient)(nil).GetSignedInToken), arg0, arg1)
}

// GetSignoutToken mocks base method
func (m *MockIdentityClient) GetSignoutToken(arg0 context.Context, arg1 identity.GetSignoutTokenArgs) (*delegatedauthorization.AccessTokenResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSignoutToken", arg0, arg1)
	ret0, _ := ret[0].(*delegatedauthorization.AccessTokenResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSignoutToken indicates an expected call of GetSignoutToken
func (mr *MockIdentityClientMockRecorder) GetSignoutToken(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSignoutToken", reflect.TypeOf((*MockIdentityClient)(nil).GetSignoutToken), arg0, arg1)
}

// GetTenant mocks base method
func (m *MockIdentityClient) GetTenant(arg0 context.Context, arg1 identity.GetTenantArgs) (*identity.TenantInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTenant", arg0, arg1)
	ret0, _ := ret[0].(*identity.TenantInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTenant indicates an expected call of GetTenant
func (mr *MockIdentityClientMockRecorder) GetTenant(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTenant", reflect.TypeOf((*MockIdentityClient)(nil).GetTenant), arg0, arg1)
}

// GetUserIdentityIdsByDomainId mocks base method
func (m *MockIdentityClient) GetUserIdentityIdsByDomainId(arg0 context.Context, arg1 identity.GetUserIdentityIdsByDomainIdArgs) (*[]uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserIdentityIdsByDomainId", arg0, arg1)
	ret0, _ := ret[0].(*[]uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserIdentityIdsByDomainId indicates an expected call of GetUserIdentityIdsByDomainId
func (mr *MockIdentityClientMockRecorder) GetUserIdentityIdsByDomainId(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserIdentityIdsByDomainId", reflect.TypeOf((*MockIdentityClient)(nil).GetUserIdentityIdsByDomainId), arg0, arg1)
}

// ListGroups mocks base method
func (m *MockIdentityClient) ListGroups(arg0 context.Context, arg1 identity.ListGroupsArgs) (*[]identity.Identity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListGroups", arg0, arg1)
	ret0, _ := ret[0].(*[]identity.Identity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListGroups indicates an expected call of ListGroups
func (mr *MockIdentityClientMockRecorder) ListGroups(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListGroups", reflect.TypeOf((*MockIdentityClient)(nil).ListGroups), arg0, arg1)
}

// ReadIdentities mocks base method
func (m *MockIdentityClient) ReadIdentities(arg0 context.Context, arg1 identity.ReadIdentitiesArgs) (*[]identity.Identity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadIdentities", arg0, arg1)
	ret0, _ := ret[0].(*[]identity.Identity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadIdentities indicates an expected call of ReadIdentities
func (mr *MockIdentityClientMockRecorder) ReadIdentities(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadIdentities", reflect.TypeOf((*MockIdentityClient)(nil).ReadIdentities), arg0, arg1)
}

// ReadIdentitiesByScope mocks base method
func (m *MockIdentityClient) ReadIdentitiesByScope(arg0 context.Context, arg1 identity.ReadIdentitiesByScopeArgs) (*[]identity.Identity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadIdentitiesByScope", arg0, arg1)
	ret0, _ := ret[0].(*[]identity.Identity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadIdentitiesByScope indicates an expected call of ReadIdentitiesByScope
func (mr *MockIdentityClientMockRecorder) ReadIdentitiesByScope(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadIdentitiesByScope", reflect.TypeOf((*MockIdentityClient)(nil).ReadIdentitiesByScope), arg0, arg1)
}

// ReadIdentity mocks base method
func (m *MockIdentityClient) ReadIdentity(arg0 context.Context, arg1 identity.ReadIdentityArgs) (*identity.Identity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadIdentity", arg0, arg1)
	ret0, _ := ret[0].(*identity.Identity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadIdentity indicates an expected call of ReadIdentity
func (mr *MockIdentityClientMockRecorder) ReadIdentity(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadIdentity", reflect.TypeOf((*MockIdentityClient)(nil).ReadIdentity), arg0, arg1)
}

// ReadIdentityBatch mocks base method
func (m *MockIdentityClient) ReadIdentityBatch(arg0 context.Context, arg1 identity.ReadIdentityBatchArgs) (*[]identity.Identity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadIdentityBatch", arg0, arg1)
	ret0, _ := ret[0].(*[]identity.Identity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadIdentityBatch indicates an expected call of ReadIdentityBatch
func (mr *MockIdentityClientMockRecorder) ReadIdentityBatch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadIdentityBatch", reflect.TypeOf((*MockIdentityClient)(nil).ReadIdentityBatch), arg0, arg1)
}

// ReadMember mocks base method
func (m *MockIdentityClient) ReadMember(arg0 context.Context, arg1 identity.ReadMemberArgs) (*string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadMember", arg0, arg1)
	ret0, _ := ret[0].(*string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadMember indicates an expected call of ReadMember
func (mr *MockIdentityClientMockRecorder) ReadMember(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadMember", reflect.TypeOf((*MockIdentityClient)(nil).ReadMember), arg0, arg1)
}

// ReadMemberOf mocks base method
func (m *MockIdentityClient) ReadMemberOf(arg0 context.Context, arg1 identity.ReadMemberOfArgs) (*string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadMemberOf", arg0, arg1)
	ret0, _ := ret[0].(*string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadMemberOf indicates an expected call of ReadMemberOf
func (mr *MockIdentityClientMockRecorder) ReadMemberOf(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadMemberOf", reflect.TypeOf((*MockIdentityClient)(nil).ReadMemberOf), arg0, arg1)
}

// ReadMembers mocks base method
func (m *MockIdentityClient) ReadMembers(arg0 context.Context, arg1 identity.ReadMembersArgs) (*[]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadMembers", arg0, arg1)
	ret0, _ := ret[0].(*[]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadMembers indicates an expected call of ReadMembers
func (mr *MockIdentityClientMockRecorder) ReadMembers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadMembers", reflect.TypeOf((*MockIdentityClient)(nil).ReadMembers), arg0, arg1)
}

// ReadMembersOf mocks base method
func (m *MockIdentityClient) ReadMembersOf(arg0 context.Context, arg1 identity.ReadMembersOfArgs) (*[]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadMembersOf", arg0, arg1)
	ret0, _ := ret[0].(*[]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadMembersOf indicates an expected call of ReadMembersOf
func (mr *MockIdentityClientMockRecorder) ReadMembersOf(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadMembersOf", reflect.TypeOf((*MockIdentityClient)(nil).ReadMembersOf), arg0, arg1)
}

// RefreshMembersOf mocks base method
func (m *MockIdentityClient) RefreshMembersOf(arg0 context.Context, arg1 identity.RefreshMembersOfArgs) (*[]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshMembersOf", arg0, arg1)
	ret0, _ := ret[0].(*[]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RefreshMembersOf indicates an expected call of RefreshMembersOf
func (mr *MockIdentityClientMockRecorder) RefreshMembersOf(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshMembersOf", reflect.TypeOf((*MockIdentityClient)(nil).RefreshMembersOf), arg0, arg1)
}

// RemoveMember mocks base method
func (m *MockIdentityClient) RemoveMember(arg0 context.Context, arg1 identity.RemoveMemberArgs) (*bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveMember", arg0, arg1)
	ret0, _ := ret[0].(*bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveMember indicates an expected call of RemoveMember
func (mr *MockIdentityClientMockRecorder) RemoveMember(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveMember", reflect.TypeOf((*MockIdentityClient)(nil).RemoveMember), arg0, arg1)
}

// UpdateIdentities mocks base method
func (m *MockIdentityClient) UpdateIdentities(arg0 context.Context, arg1 identity.UpdateIdentitiesArgs) (*[]identity.IdentityUpdateData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateIdentities", arg0, arg1)
	ret0, _ := ret[0].(*[]identity.IdentityUpdateData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateIdentities indicates an expected call of UpdateIdentities
func (mr *MockIdentityClientMockRecorder) UpdateIdentities(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateIdentities", reflect.TypeOf((*MockIdentityClient)(nil).UpdateIdentities), arg0, arg1)
}

// UpdateIdentity mocks base method
func (m *MockIdentityClient) UpdateIdentity(arg0 context.Context, arg1 identity.UpdateIdentityArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateIdentity", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateIdentity indicates an expected call of UpdateIdentity
func (mr *MockIdentityClientMockRecorder) UpdateIdentity(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateIdentity", reflect.TypeOf((*MockIdentityClient)(nil).UpdateIdentity), arg0, arg1)
}

// UpdateScope mocks base method
func (m *MockIdentityClient) UpdateScope(arg0 context.Context, arg1 identity.UpdateScopeArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateScope", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateScope indicates an expected call of UpdateScope
func (mr *MockIdentityClientMockRecorder) UpdateScope(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateScope", reflect.TypeOf((*MockIdentityClient)(nil).UpdateScope), arg0, arg1)
}
