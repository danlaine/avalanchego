// Copyright (C) 2019-2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ava-labs/avalanchego/api/server (interfaces: Server)

// Package server is a generated GoMock package.
package server

import (
	reflect "reflect"
	sync "sync"

	snow "github.com/ava-labs/avalanchego/snow"
	common "github.com/ava-labs/avalanchego/snow/engine/common"
	gomock "github.com/golang/mock/gomock"
)

// MockServer is a mock of Server interface.
type MockServer struct {
	ctrl     *gomock.Controller
	recorder *MockServerMockRecorder
}

// MockServerMockRecorder is the mock recorder for MockServer.
type MockServerMockRecorder struct {
	mock *MockServer
}

// NewMockServer creates a new mock instance.
func NewMockServer(ctrl *gomock.Controller) *MockServer {
	mock := &MockServer{ctrl: ctrl}
	mock.recorder = &MockServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServer) EXPECT() *MockServerMockRecorder {
	return m.recorder
}

// AddAliases mocks base method.
func (m *MockServer) AddAliases(arg0 string, arg1 ...string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddAliases", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddAliases indicates an expected call of AddAliases.
func (mr *MockServerMockRecorder) AddAliases(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAliases", reflect.TypeOf((*MockServer)(nil).AddAliases), varargs...)
}

// AddAliasesWithReadLock mocks base method.
func (m *MockServer) AddAliasesWithReadLock(arg0 string, arg1 ...string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddAliasesWithReadLock", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddAliasesWithReadLock indicates an expected call of AddAliasesWithReadLock.
func (mr *MockServerMockRecorder) AddAliasesWithReadLock(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAliasesWithReadLock", reflect.TypeOf((*MockServer)(nil).AddAliasesWithReadLock), varargs...)
}

// AddRoute mocks base method.
func (m *MockServer) AddRoute(arg0 *common.HTTPHandler, arg1 *sync.RWMutex, arg2, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRoute", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddRoute indicates an expected call of AddRoute.
func (mr *MockServerMockRecorder) AddRoute(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRoute", reflect.TypeOf((*MockServer)(nil).AddRoute), arg0, arg1, arg2, arg3)
}

// AddRouteWithReadLock mocks base method.
func (m *MockServer) AddRouteWithReadLock(arg0 *common.HTTPHandler, arg1 *sync.RWMutex, arg2, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRouteWithReadLock", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddRouteWithReadLock indicates an expected call of AddRouteWithReadLock.
func (mr *MockServerMockRecorder) AddRouteWithReadLock(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRouteWithReadLock", reflect.TypeOf((*MockServer)(nil).AddRouteWithReadLock), arg0, arg1, arg2, arg3)
}

// Dispatch mocks base method.
func (m *MockServer) Dispatch() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Dispatch")
	ret0, _ := ret[0].(error)
	return ret0
}

// Dispatch indicates an expected call of Dispatch.
func (mr *MockServerMockRecorder) Dispatch() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Dispatch", reflect.TypeOf((*MockServer)(nil).Dispatch))
}

// DispatchTLS mocks base method.
func (m *MockServer) DispatchTLS(arg0, arg1 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DispatchTLS", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DispatchTLS indicates an expected call of DispatchTLS.
func (mr *MockServerMockRecorder) DispatchTLS(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DispatchTLS", reflect.TypeOf((*MockServer)(nil).DispatchTLS), arg0, arg1)
}

// RegisterChain mocks base method.
func (m *MockServer) RegisterChain(arg0 string, arg1 *snow.ConsensusContext, arg2 common.VM) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterChain", arg0, arg1, arg2)
}

// RegisterChain indicates an expected call of RegisterChain.
func (mr *MockServerMockRecorder) RegisterChain(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterChain", reflect.TypeOf((*MockServer)(nil).RegisterChain), arg0, arg1, arg2)
}

// Shutdown mocks base method.
func (m *MockServer) Shutdown() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Shutdown")
	ret0, _ := ret[0].(error)
	return ret0
}

// Shutdown indicates an expected call of Shutdown.
func (mr *MockServerMockRecorder) Shutdown() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shutdown", reflect.TypeOf((*MockServer)(nil).Shutdown))
}
