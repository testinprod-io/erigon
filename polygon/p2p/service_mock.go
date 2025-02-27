// Code generated by MockGen. DO NOT EDIT.
// Source: ./service.go
//
// Generated by this command:
//
//	mockgen -source=./service.go -destination=./service_mock.go -package=p2p . Service
//

// Package p2p is a generated GoMock package.
package p2p

import (
	context "context"
	reflect "reflect"

	sentry "github.com/ledgerwatch/erigon-lib/gointerfaces/sentry"
	types "github.com/ledgerwatch/erigon/core/types"
	eth "github.com/ledgerwatch/erigon/eth/protocols/eth"
	gomock "go.uber.org/mock/gomock"
)

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

// BlockNumMissing mocks base method.
func (m *MockService) BlockNumMissing(peerId *PeerId, blockNum uint64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "BlockNumMissing", peerId, blockNum)
}

// BlockNumMissing indicates an expected call of BlockNumMissing.
func (mr *MockServiceMockRecorder) BlockNumMissing(peerId, blockNum any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockNumMissing", reflect.TypeOf((*MockService)(nil).BlockNumMissing), peerId, blockNum)
}

// BlockNumPresent mocks base method.
func (m *MockService) BlockNumPresent(peerId *PeerId, blockNum uint64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "BlockNumPresent", peerId, blockNum)
}

// BlockNumPresent indicates an expected call of BlockNumPresent.
func (mr *MockServiceMockRecorder) BlockNumPresent(peerId, blockNum any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockNumPresent", reflect.TypeOf((*MockService)(nil).BlockNumPresent), peerId, blockNum)
}

// FetchBlocks mocks base method.
func (m *MockService) FetchBlocks(ctx context.Context, start, end uint64, peerId *PeerId) ([]*types.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchBlocks", ctx, start, end, peerId)
	ret0, _ := ret[0].([]*types.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchBlocks indicates an expected call of FetchBlocks.
func (mr *MockServiceMockRecorder) FetchBlocks(ctx, start, end, peerId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchBlocks", reflect.TypeOf((*MockService)(nil).FetchBlocks), ctx, start, end, peerId)
}

// FetchBodies mocks base method.
func (m *MockService) FetchBodies(ctx context.Context, headers []*types.Header, peerId *PeerId) ([]*types.Body, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchBodies", ctx, headers, peerId)
	ret0, _ := ret[0].([]*types.Body)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchBodies indicates an expected call of FetchBodies.
func (mr *MockServiceMockRecorder) FetchBodies(ctx, headers, peerId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchBodies", reflect.TypeOf((*MockService)(nil).FetchBodies), ctx, headers, peerId)
}

// FetchHeaders mocks base method.
func (m *MockService) FetchHeaders(ctx context.Context, start, end uint64, peerId *PeerId) ([]*types.Header, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchHeaders", ctx, start, end, peerId)
	ret0, _ := ret[0].([]*types.Header)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchHeaders indicates an expected call of FetchHeaders.
func (mr *MockServiceMockRecorder) FetchHeaders(ctx, start, end, peerId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchHeaders", reflect.TypeOf((*MockService)(nil).FetchHeaders), ctx, start, end, peerId)
}

// ListPeersMayHaveBlockNum mocks base method.
func (m *MockService) ListPeersMayHaveBlockNum(blockNum uint64) []*PeerId {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPeersMayHaveBlockNum", blockNum)
	ret0, _ := ret[0].([]*PeerId)
	return ret0
}

// ListPeersMayHaveBlockNum indicates an expected call of ListPeersMayHaveBlockNum.
func (mr *MockServiceMockRecorder) ListPeersMayHaveBlockNum(blockNum any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPeersMayHaveBlockNum", reflect.TypeOf((*MockService)(nil).ListPeersMayHaveBlockNum), blockNum)
}

// MaxPeers mocks base method.
func (m *MockService) MaxPeers() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MaxPeers")
	ret0, _ := ret[0].(int)
	return ret0
}

// MaxPeers indicates an expected call of MaxPeers.
func (mr *MockServiceMockRecorder) MaxPeers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaxPeers", reflect.TypeOf((*MockService)(nil).MaxPeers))
}

// PeerConnected mocks base method.
func (m *MockService) PeerConnected(peerId *PeerId) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PeerConnected", peerId)
}

// PeerConnected indicates an expected call of PeerConnected.
func (mr *MockServiceMockRecorder) PeerConnected(peerId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PeerConnected", reflect.TypeOf((*MockService)(nil).PeerConnected), peerId)
}

// PeerDisconnected mocks base method.
func (m *MockService) PeerDisconnected(peerId *PeerId) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PeerDisconnected", peerId)
}

// PeerDisconnected indicates an expected call of PeerDisconnected.
func (mr *MockServiceMockRecorder) PeerDisconnected(peerId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PeerDisconnected", reflect.TypeOf((*MockService)(nil).PeerDisconnected), peerId)
}

// Penalize mocks base method.
func (m *MockService) Penalize(ctx context.Context, peerId *PeerId) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Penalize", ctx, peerId)
	ret0, _ := ret[0].(error)
	return ret0
}

// Penalize indicates an expected call of Penalize.
func (mr *MockServiceMockRecorder) Penalize(ctx, peerId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Penalize", reflect.TypeOf((*MockService)(nil).Penalize), ctx, peerId)
}

// RegisterBlockBodiesObserver mocks base method.
func (m *MockService) RegisterBlockBodiesObserver(observer MessageObserver[*DecodedInboundMessage[*eth.BlockBodiesPacket66]]) UnregisterFunc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterBlockBodiesObserver", observer)
	ret0, _ := ret[0].(UnregisterFunc)
	return ret0
}

// RegisterBlockBodiesObserver indicates an expected call of RegisterBlockBodiesObserver.
func (mr *MockServiceMockRecorder) RegisterBlockBodiesObserver(observer any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterBlockBodiesObserver", reflect.TypeOf((*MockService)(nil).RegisterBlockBodiesObserver), observer)
}

// RegisterBlockHeadersObserver mocks base method.
func (m *MockService) RegisterBlockHeadersObserver(observer MessageObserver[*DecodedInboundMessage[*eth.BlockHeadersPacket66]]) UnregisterFunc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterBlockHeadersObserver", observer)
	ret0, _ := ret[0].(UnregisterFunc)
	return ret0
}

// RegisterBlockHeadersObserver indicates an expected call of RegisterBlockHeadersObserver.
func (mr *MockServiceMockRecorder) RegisterBlockHeadersObserver(observer any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterBlockHeadersObserver", reflect.TypeOf((*MockService)(nil).RegisterBlockHeadersObserver), observer)
}

// RegisterNewBlockHashesObserver mocks base method.
func (m *MockService) RegisterNewBlockHashesObserver(observer MessageObserver[*DecodedInboundMessage[*eth.NewBlockHashesPacket]]) UnregisterFunc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterNewBlockHashesObserver", observer)
	ret0, _ := ret[0].(UnregisterFunc)
	return ret0
}

// RegisterNewBlockHashesObserver indicates an expected call of RegisterNewBlockHashesObserver.
func (mr *MockServiceMockRecorder) RegisterNewBlockHashesObserver(observer any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterNewBlockHashesObserver", reflect.TypeOf((*MockService)(nil).RegisterNewBlockHashesObserver), observer)
}

// RegisterNewBlockObserver mocks base method.
func (m *MockService) RegisterNewBlockObserver(observer MessageObserver[*DecodedInboundMessage[*eth.NewBlockPacket]]) UnregisterFunc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterNewBlockObserver", observer)
	ret0, _ := ret[0].(UnregisterFunc)
	return ret0
}

// RegisterNewBlockObserver indicates an expected call of RegisterNewBlockObserver.
func (mr *MockServiceMockRecorder) RegisterNewBlockObserver(observer any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterNewBlockObserver", reflect.TypeOf((*MockService)(nil).RegisterNewBlockObserver), observer)
}

// RegisterPeerEventObserver mocks base method.
func (m *MockService) RegisterPeerEventObserver(observer MessageObserver[*sentry.PeerEvent]) UnregisterFunc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterPeerEventObserver", observer)
	ret0, _ := ret[0].(UnregisterFunc)
	return ret0
}

// RegisterPeerEventObserver indicates an expected call of RegisterPeerEventObserver.
func (mr *MockServiceMockRecorder) RegisterPeerEventObserver(observer any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterPeerEventObserver", reflect.TypeOf((*MockService)(nil).RegisterPeerEventObserver), observer)
}

// Run mocks base method.
func (m *MockService) Run(ctx context.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Run", ctx)
}

// Run indicates an expected call of Run.
func (mr *MockServiceMockRecorder) Run(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockService)(nil).Run), ctx)
}
