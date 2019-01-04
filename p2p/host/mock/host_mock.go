// Code generated by MockGen. DO NOT EDIT.
// Source: ./p2p/host/host.go

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	p2p "github.com/harmony-one/harmony/p2p"
	reflect "reflect"
)

// MockHost is a mock of Host interface
type MockHost struct {
	ctrl     *gomock.Controller
	recorder *MockHostMockRecorder
}

// MockHostMockRecorder is the mock recorder for MockHost
type MockHostMockRecorder struct {
	mock *MockHost
}

// NewMockHost creates a new mock instance
func NewMockHost(ctrl *gomock.Controller) *MockHost {
	mock := &MockHost{ctrl: ctrl}
	mock.recorder = &MockHostMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHost) EXPECT() *MockHostMockRecorder {
	return m.recorder
}

// GetSelfPeer mocks base method
func (m *MockHost) GetSelfPeer() p2p.Peer {
	ret := m.ctrl.Call(m, "GetSelfPeer")
	ret0, _ := ret[0].(p2p.Peer)
	return ret0
}

// GetSelfPeer indicates an expected call of GetSelfPeer
func (mr *MockHostMockRecorder) GetSelfPeer() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSelfPeer", reflect.TypeOf((*MockHost)(nil).GetSelfPeer))
}

// SendMessage mocks base method
func (m *MockHost) SendMessage(arg0 p2p.Peer, arg1 []byte) error {
	ret := m.ctrl.Call(m, "SendMessage", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMessage indicates an expected call of SendMessage
func (mr *MockHostMockRecorder) SendMessage(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMessage", reflect.TypeOf((*MockHost)(nil).SendMessage), arg0, arg1)
}

// BindHandlerAndServe mocks base method
func (m *MockHost) BindHandlerAndServe(handler p2p.StreamHandler) {
	m.ctrl.Call(m, "BindHandlerAndServe", handler)
}

// BindHandlerAndServe indicates an expected call of BindHandlerAndServe
func (mr *MockHostMockRecorder) BindHandlerAndServe(handler interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BindHandlerAndServe", reflect.TypeOf((*MockHost)(nil).BindHandlerAndServe), handler)
}

// Close mocks base method
func (m *MockHost) Close() error {
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockHostMockRecorder) Close() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockHost)(nil).Close))
}
