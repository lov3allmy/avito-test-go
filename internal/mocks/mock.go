// Code generated by MockGen. DO NOT EDIT.
// Source: domain.go

// Package mock_domain is a generated GoMock package.
package mock_domain

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/lov3allmy/avito-test-go/internal/domain"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockRepository) CreateUser(user *domain.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockRepositoryMockRecorder) CreateUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockRepository)(nil).CreateUser), user)
}

// GetUser mocks base method.
func (m *MockRepository) GetUser(userID int) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", userID)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockRepositoryMockRecorder) GetUser(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockRepository)(nil).GetUser), userID)
}

// MakeP2PTransfer mocks base method.
func (m *MockRepository) MakeP2PTransfer(p2pInput domain.P2PInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeP2PTransfer", p2pInput)
	ret0, _ := ret[0].(error)
	return ret0
}

// MakeP2PTransfer indicates an expected call of MakeP2PTransfer.
func (mr *MockRepositoryMockRecorder) MakeP2PTransfer(p2pInput interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeP2PTransfer", reflect.TypeOf((*MockRepository)(nil).MakeP2PTransfer), p2pInput)
}

// UpdateUser mocks base method.
func (m *MockRepository) UpdateUser(userID int, user *domain.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", userID, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockRepositoryMockRecorder) UpdateUser(userID, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockRepository)(nil).UpdateUser), userID, user)
}

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockService) CreateUser(user *domain.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockServiceMockRecorder) CreateUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockService)(nil).CreateUser), user)
}

// GetUser mocks base method.
func (m *MockService) GetUser(userID int) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", userID)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockServiceMockRecorder) GetUser(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockService)(nil).GetUser), userID)
}

// MakeP2PTransfer mocks base method.
func (m *MockService) MakeP2PTransfer(p2pInput domain.P2PInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeP2PTransfer", p2pInput)
	ret0, _ := ret[0].(error)
	return ret0
}

// MakeP2PTransfer indicates an expected call of MakeP2PTransfer.
func (mr *MockServiceMockRecorder) MakeP2PTransfer(p2pInput interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeP2PTransfer", reflect.TypeOf((*MockService)(nil).MakeP2PTransfer), p2pInput)
}

// UpdateUser mocks base method.
func (m *MockService) UpdateUser(userID int, user *domain.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", userID, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockServiceMockRecorder) UpdateUser(userID, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockService)(nil).UpdateUser), userID, user)
}