// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/simonswine/cluster-api-provider-hetzner/pkg/cloud/scope (interfaces: HetznerClient)

// Package mock_scope is a generated GoMock package.
package mock_scope

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	hcloud "github.com/hetznercloud/hcloud-go/hcloud"
	reflect "reflect"
)

// MockHetznerClient is a mock of HetznerClient interface
type MockHetznerClient struct {
	ctrl     *gomock.Controller
	recorder *MockHetznerClientMockRecorder
}

// MockHetznerClientMockRecorder is the mock recorder for MockHetznerClient
type MockHetznerClientMockRecorder struct {
	mock *MockHetznerClient
}

// NewMockHetznerClient creates a new mock instance
func NewMockHetznerClient(ctrl *gomock.Controller) *MockHetznerClient {
	mock := &MockHetznerClient{ctrl: ctrl}
	mock.recorder = &MockHetznerClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHetznerClient) EXPECT() *MockHetznerClientMockRecorder {
	return m.recorder
}

// CreateFloatingIP mocks base method
func (m *MockHetznerClient) CreateFloatingIP(arg0 context.Context, arg1 hcloud.FloatingIPCreateOpts) (hcloud.FloatingIPCreateResult, *hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFloatingIP", arg0, arg1)
	ret0, _ := ret[0].(hcloud.FloatingIPCreateResult)
	ret1, _ := ret[1].(*hcloud.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateFloatingIP indicates an expected call of CreateFloatingIP
func (mr *MockHetznerClientMockRecorder) CreateFloatingIP(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFloatingIP", reflect.TypeOf((*MockHetznerClient)(nil).CreateFloatingIP), arg0, arg1)
}

// CreateNetwork mocks base method
func (m *MockHetznerClient) CreateNetwork(arg0 context.Context, arg1 hcloud.NetworkCreateOpts) (*hcloud.Network, *hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNetwork", arg0, arg1)
	ret0, _ := ret[0].(*hcloud.Network)
	ret1, _ := ret[1].(*hcloud.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateNetwork indicates an expected call of CreateNetwork
func (mr *MockHetznerClientMockRecorder) CreateNetwork(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNetwork", reflect.TypeOf((*MockHetznerClient)(nil).CreateNetwork), arg0, arg1)
}

// CreateServer mocks base method
func (m *MockHetznerClient) CreateServer(arg0 context.Context, arg1 hcloud.ServerCreateOpts) (hcloud.ServerCreateResult, *hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateServer", arg0, arg1)
	ret0, _ := ret[0].(hcloud.ServerCreateResult)
	ret1, _ := ret[1].(*hcloud.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateServer indicates an expected call of CreateServer
func (mr *MockHetznerClientMockRecorder) CreateServer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateServer", reflect.TypeOf((*MockHetznerClient)(nil).CreateServer), arg0, arg1)
}

// CreateVolume mocks base method
func (m *MockHetznerClient) CreateVolume(arg0 context.Context, arg1 hcloud.VolumeCreateOpts) (hcloud.VolumeCreateResult, *hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateVolume", arg0, arg1)
	ret0, _ := ret[0].(hcloud.VolumeCreateResult)
	ret1, _ := ret[1].(*hcloud.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateVolume indicates an expected call of CreateVolume
func (mr *MockHetznerClientMockRecorder) CreateVolume(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateVolume", reflect.TypeOf((*MockHetznerClient)(nil).CreateVolume), arg0, arg1)
}

// DeleteFloatingIP mocks base method
func (m *MockHetznerClient) DeleteFloatingIP(arg0 context.Context, arg1 *hcloud.FloatingIP) (*hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFloatingIP", arg0, arg1)
	ret0, _ := ret[0].(*hcloud.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteFloatingIP indicates an expected call of DeleteFloatingIP
func (mr *MockHetznerClientMockRecorder) DeleteFloatingIP(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFloatingIP", reflect.TypeOf((*MockHetznerClient)(nil).DeleteFloatingIP), arg0, arg1)
}

// DeleteNetwork mocks base method
func (m *MockHetznerClient) DeleteNetwork(arg0 context.Context, arg1 *hcloud.Network) (*hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteNetwork", arg0, arg1)
	ret0, _ := ret[0].(*hcloud.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteNetwork indicates an expected call of DeleteNetwork
func (mr *MockHetznerClientMockRecorder) DeleteNetwork(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNetwork", reflect.TypeOf((*MockHetznerClient)(nil).DeleteNetwork), arg0, arg1)
}

// DeleteServer mocks base method
func (m *MockHetznerClient) DeleteServer(arg0 context.Context, arg1 *hcloud.Server) (*hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteServer", arg0, arg1)
	ret0, _ := ret[0].(*hcloud.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteServer indicates an expected call of DeleteServer
func (mr *MockHetznerClientMockRecorder) DeleteServer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteServer", reflect.TypeOf((*MockHetznerClient)(nil).DeleteServer), arg0, arg1)
}

// DeleteVolume mocks base method
func (m *MockHetznerClient) DeleteVolume(arg0 context.Context, arg1 *hcloud.Volume) (*hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteVolume", arg0, arg1)
	ret0, _ := ret[0].(*hcloud.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteVolume indicates an expected call of DeleteVolume
func (mr *MockHetznerClientMockRecorder) DeleteVolume(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteVolume", reflect.TypeOf((*MockHetznerClient)(nil).DeleteVolume), arg0, arg1)
}

// ListFloatingIPs mocks base method
func (m *MockHetznerClient) ListFloatingIPs(arg0 context.Context, arg1 hcloud.FloatingIPListOpts) ([]*hcloud.FloatingIP, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFloatingIPs", arg0, arg1)
	ret0, _ := ret[0].([]*hcloud.FloatingIP)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFloatingIPs indicates an expected call of ListFloatingIPs
func (mr *MockHetznerClientMockRecorder) ListFloatingIPs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFloatingIPs", reflect.TypeOf((*MockHetznerClient)(nil).ListFloatingIPs), arg0, arg1)
}

// ListImages mocks base method
func (m *MockHetznerClient) ListImages(arg0 context.Context, arg1 hcloud.ImageListOpts) ([]*hcloud.Image, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListImages", arg0, arg1)
	ret0, _ := ret[0].([]*hcloud.Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListImages indicates an expected call of ListImages
func (mr *MockHetznerClientMockRecorder) ListImages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListImages", reflect.TypeOf((*MockHetznerClient)(nil).ListImages), arg0, arg1)
}

// ListLocation mocks base method
func (m *MockHetznerClient) ListLocation(arg0 context.Context) ([]*hcloud.Location, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListLocation", arg0)
	ret0, _ := ret[0].([]*hcloud.Location)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListLocation indicates an expected call of ListLocation
func (mr *MockHetznerClientMockRecorder) ListLocation(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListLocation", reflect.TypeOf((*MockHetznerClient)(nil).ListLocation), arg0)
}

// ListNetworks mocks base method
func (m *MockHetznerClient) ListNetworks(arg0 context.Context, arg1 hcloud.NetworkListOpts) ([]*hcloud.Network, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListNetworks", arg0, arg1)
	ret0, _ := ret[0].([]*hcloud.Network)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListNetworks indicates an expected call of ListNetworks
func (mr *MockHetznerClientMockRecorder) ListNetworks(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListNetworks", reflect.TypeOf((*MockHetznerClient)(nil).ListNetworks), arg0, arg1)
}

// ListServers mocks base method
func (m *MockHetznerClient) ListServers(arg0 context.Context, arg1 hcloud.ServerListOpts) ([]*hcloud.Server, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListServers", arg0, arg1)
	ret0, _ := ret[0].([]*hcloud.Server)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListServers indicates an expected call of ListServers
func (mr *MockHetznerClientMockRecorder) ListServers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListServers", reflect.TypeOf((*MockHetznerClient)(nil).ListServers), arg0, arg1)
}

// ListVolumes mocks base method
func (m *MockHetznerClient) ListVolumes(arg0 context.Context, arg1 hcloud.VolumeListOpts) ([]*hcloud.Volume, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListVolumes", arg0, arg1)
	ret0, _ := ret[0].([]*hcloud.Volume)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListVolumes indicates an expected call of ListVolumes
func (mr *MockHetznerClientMockRecorder) ListVolumes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListVolumes", reflect.TypeOf((*MockHetznerClient)(nil).ListVolumes), arg0, arg1)
}

// ShutdownServer mocks base method
func (m *MockHetznerClient) ShutdownServer(arg0 context.Context, arg1 *hcloud.Server) (*hcloud.Action, *hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShutdownServer", arg0, arg1)
	ret0, _ := ret[0].(*hcloud.Action)
	ret1, _ := ret[1].(*hcloud.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ShutdownServer indicates an expected call of ShutdownServer
func (mr *MockHetznerClientMockRecorder) ShutdownServer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShutdownServer", reflect.TypeOf((*MockHetznerClient)(nil).ShutdownServer), arg0, arg1)
}
