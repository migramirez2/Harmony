// Code generated by MockGen. DO NOT EDIT.
// Source: host.go

// Package mock_p2p is a generated GoMock package.
package mock_p2p

import (
	gomock "github.com/golang/mock/gomock"
	p2p "github.com/harmony-one/harmony/p2p"
	go_libp2p_host "github.com/libp2p/go-libp2p-host"
	go_libp2p_peer "github.com/libp2p/go-libp2p-peer"
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

// AddPeer mocks base method
func (m *MockHost) AddPeer(arg0 *p2p.Peer) error {
	ret := m.ctrl.Call(m, "AddPeer", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddPeer indicates an expected call of AddPeer
func (mr *MockHostMockRecorder) AddPeer(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPeer", reflect.TypeOf((*MockHost)(nil).AddPeer), arg0)
}

// GetID mocks base method
func (m *MockHost) GetID() go_libp2p_peer.ID {
	ret := m.ctrl.Call(m, "GetID")
	ret0, _ := ret[0].(go_libp2p_peer.ID)
	return ret0
}

// GetID indicates an expected call of GetID
func (mr *MockHostMockRecorder) GetID() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetID", reflect.TypeOf((*MockHost)(nil).GetID))
}

// GetP2PHost mocks base method
func (m *MockHost) GetP2PHost() go_libp2p_host.Host {
	ret := m.ctrl.Call(m, "GetP2PHost")
	ret0, _ := ret[0].(go_libp2p_host.Host)
	return ret0
}

// GetP2PHost indicates an expected call of GetP2PHost
func (mr *MockHostMockRecorder) GetP2PHost() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetP2PHost", reflect.TypeOf((*MockHost)(nil).GetP2PHost))
}

// SendMessageToGroups mocks base method
func (m *MockHost) SendMessageToGroups(groups []p2p.GroupID, msg []byte) error {
	ret := m.ctrl.Call(m, "SendMessageToGroups", groups, msg)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMessageToGroups indicates an expected call of SendMessageToGroups
func (mr *MockHostMockRecorder) SendMessageToGroups(groups, msg interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMessageToGroups", reflect.TypeOf((*MockHost)(nil).SendMessageToGroups), groups, msg)
}

// GroupReceiver mocks base method
func (m *MockHost) GroupReceiver(arg0 p2p.GroupID) (p2p.GroupReceiver, error) {
	ret := m.ctrl.Call(m, "GroupReceiver", arg0)
	ret0, _ := ret[0].(p2p.GroupReceiver)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GroupReceiver indicates an expected call of GroupReceiver
func (mr *MockHostMockRecorder) GroupReceiver(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GroupReceiver", reflect.TypeOf((*MockHost)(nil).GroupReceiver), arg0)
}
