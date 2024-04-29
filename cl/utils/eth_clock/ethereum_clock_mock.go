// Code generated by MockGen. DO NOT EDIT.
// Source: ./ethereum_clock.go
//
// Generated by this command:
//
//	mockgen -typed=true -source=./ethereum_clock.go -destination=./ethereum_clock_mock.go -package=eth_clock . EthereumClock
//

// Package eth_clock is a generated GoMock package.
package eth_clock

import (
	reflect "reflect"
	time "time"

	common "github.com/ledgerwatch/erigon-lib/common"
	clparams "github.com/ledgerwatch/erigon/cl/clparams"
	gomock "go.uber.org/mock/gomock"
)

// MockEthereumClock is a mock of EthereumClock interface.
type MockEthereumClock struct {
	ctrl     *gomock.Controller
	recorder *MockEthereumClockMockRecorder
}

// MockEthereumClockMockRecorder is the mock recorder for MockEthereumClock.
type MockEthereumClockMockRecorder struct {
	mock *MockEthereumClock
}

// NewMockEthereumClock creates a new mock instance.
func NewMockEthereumClock(ctrl *gomock.Controller) *MockEthereumClock {
	mock := &MockEthereumClock{ctrl: ctrl}
	mock.recorder = &MockEthereumClockMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEthereumClock) EXPECT() *MockEthereumClockMockRecorder {
	return m.recorder
}

// ComputeForkDigestForVersion mocks base method.
func (m *MockEthereumClock) ComputeForkDigestForVersion(currentVersion common.Bytes4) (common.Bytes4, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ComputeForkDigestForVersion", currentVersion)
	ret0, _ := ret[0].(common.Bytes4)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ComputeForkDigestForVersion indicates an expected call of ComputeForkDigestForVersion.
func (mr *MockEthereumClockMockRecorder) ComputeForkDigestForVersion(currentVersion any) *MockEthereumClockComputeForkDigestForVersionCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ComputeForkDigestForVersion", reflect.TypeOf((*MockEthereumClock)(nil).ComputeForkDigestForVersion), currentVersion)
	return &MockEthereumClockComputeForkDigestForVersionCall{Call: call}
}

// MockEthereumClockComputeForkDigestForVersionCall wrap *gomock.Call
type MockEthereumClockComputeForkDigestForVersionCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockEthereumClockComputeForkDigestForVersionCall) Return(digest common.Bytes4, err error) *MockEthereumClockComputeForkDigestForVersionCall {
	c.Call = c.Call.Return(digest, err)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockEthereumClockComputeForkDigestForVersionCall) Do(f func(common.Bytes4) (common.Bytes4, error)) *MockEthereumClockComputeForkDigestForVersionCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockEthereumClockComputeForkDigestForVersionCall) DoAndReturn(f func(common.Bytes4) (common.Bytes4, error)) *MockEthereumClockComputeForkDigestForVersionCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// CurrentForkDigest mocks base method.
func (m *MockEthereumClock) CurrentForkDigest() (common.Bytes4, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CurrentForkDigest")
	ret0, _ := ret[0].(common.Bytes4)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CurrentForkDigest indicates an expected call of CurrentForkDigest.
func (mr *MockEthereumClockMockRecorder) CurrentForkDigest() *MockEthereumClockCurrentForkDigestCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CurrentForkDigest", reflect.TypeOf((*MockEthereumClock)(nil).CurrentForkDigest))
	return &MockEthereumClockCurrentForkDigestCall{Call: call}
}

// MockEthereumClockCurrentForkDigestCall wrap *gomock.Call
type MockEthereumClockCurrentForkDigestCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockEthereumClockCurrentForkDigestCall) Return(arg0 common.Bytes4, arg1 error) *MockEthereumClockCurrentForkDigestCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockEthereumClockCurrentForkDigestCall) Do(f func() (common.Bytes4, error)) *MockEthereumClockCurrentForkDigestCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockEthereumClockCurrentForkDigestCall) DoAndReturn(f func() (common.Bytes4, error)) *MockEthereumClockCurrentForkDigestCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ForkId mocks base method.
func (m *MockEthereumClock) ForkId() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForkId")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ForkId indicates an expected call of ForkId.
func (mr *MockEthereumClockMockRecorder) ForkId() *MockEthereumClockForkIdCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForkId", reflect.TypeOf((*MockEthereumClock)(nil).ForkId))
	return &MockEthereumClockForkIdCall{Call: call}
}

// MockEthereumClockForkIdCall wrap *gomock.Call
type MockEthereumClockForkIdCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockEthereumClockForkIdCall) Return(arg0 []byte, arg1 error) *MockEthereumClockForkIdCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockEthereumClockForkIdCall) Do(f func() ([]byte, error)) *MockEthereumClockForkIdCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockEthereumClockForkIdCall) DoAndReturn(f func() ([]byte, error)) *MockEthereumClockForkIdCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GenesisTime mocks base method.
func (m *MockEthereumClock) GenesisTime() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenesisTime")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GenesisTime indicates an expected call of GenesisTime.
func (mr *MockEthereumClockMockRecorder) GenesisTime() *MockEthereumClockGenesisTimeCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenesisTime", reflect.TypeOf((*MockEthereumClock)(nil).GenesisTime))
	return &MockEthereumClockGenesisTimeCall{Call: call}
}

// MockEthereumClockGenesisTimeCall wrap *gomock.Call
type MockEthereumClockGenesisTimeCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockEthereumClockGenesisTimeCall) Return(arg0 uint64) *MockEthereumClockGenesisTimeCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockEthereumClockGenesisTimeCall) Do(f func() uint64) *MockEthereumClockGenesisTimeCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockEthereumClockGenesisTimeCall) DoAndReturn(f func() uint64) *MockEthereumClockGenesisTimeCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GenesisValidatorsRoot mocks base method.
func (m *MockEthereumClock) GenesisValidatorsRoot() common.Hash {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenesisValidatorsRoot")
	ret0, _ := ret[0].(common.Hash)
	return ret0
}

// GenesisValidatorsRoot indicates an expected call of GenesisValidatorsRoot.
func (mr *MockEthereumClockMockRecorder) GenesisValidatorsRoot() *MockEthereumClockGenesisValidatorsRootCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenesisValidatorsRoot", reflect.TypeOf((*MockEthereumClock)(nil).GenesisValidatorsRoot))
	return &MockEthereumClockGenesisValidatorsRootCall{Call: call}
}

// MockEthereumClockGenesisValidatorsRootCall wrap *gomock.Call
type MockEthereumClockGenesisValidatorsRootCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockEthereumClockGenesisValidatorsRootCall) Return(arg0 common.Hash) *MockEthereumClockGenesisValidatorsRootCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockEthereumClockGenesisValidatorsRootCall) Do(f func() common.Hash) *MockEthereumClockGenesisValidatorsRootCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockEthereumClockGenesisValidatorsRootCall) DoAndReturn(f func() common.Hash) *MockEthereumClockGenesisValidatorsRootCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetCurrentEpoch mocks base method.
func (m *MockEthereumClock) GetCurrentEpoch() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentEpoch")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetCurrentEpoch indicates an expected call of GetCurrentEpoch.
func (mr *MockEthereumClockMockRecorder) GetCurrentEpoch() *MockEthereumClockGetCurrentEpochCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentEpoch", reflect.TypeOf((*MockEthereumClock)(nil).GetCurrentEpoch))
	return &MockEthereumClockGetCurrentEpochCall{Call: call}
}

// MockEthereumClockGetCurrentEpochCall wrap *gomock.Call
type MockEthereumClockGetCurrentEpochCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockEthereumClockGetCurrentEpochCall) Return(arg0 uint64) *MockEthereumClockGetCurrentEpochCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockEthereumClockGetCurrentEpochCall) Do(f func() uint64) *MockEthereumClockGetCurrentEpochCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockEthereumClockGetCurrentEpochCall) DoAndReturn(f func() uint64) *MockEthereumClockGetCurrentEpochCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetCurrentSlot mocks base method.
func (m *MockEthereumClock) GetCurrentSlot() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentSlot")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetCurrentSlot indicates an expected call of GetCurrentSlot.
func (mr *MockEthereumClockMockRecorder) GetCurrentSlot() *MockEthereumClockGetCurrentSlotCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentSlot", reflect.TypeOf((*MockEthereumClock)(nil).GetCurrentSlot))
	return &MockEthereumClockGetCurrentSlotCall{Call: call}
}

// MockEthereumClockGetCurrentSlotCall wrap *gomock.Call
type MockEthereumClockGetCurrentSlotCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockEthereumClockGetCurrentSlotCall) Return(arg0 uint64) *MockEthereumClockGetCurrentSlotCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockEthereumClockGetCurrentSlotCall) Do(f func() uint64) *MockEthereumClockGetCurrentSlotCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockEthereumClockGetCurrentSlotCall) DoAndReturn(f func() uint64) *MockEthereumClockGetCurrentSlotCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetSlotByTime mocks base method.
func (m *MockEthereumClock) GetSlotByTime(time time.Time) uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSlotByTime", time)
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetSlotByTime indicates an expected call of GetSlotByTime.
func (mr *MockEthereumClockMockRecorder) GetSlotByTime(time any) *MockEthereumClockGetSlotByTimeCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSlotByTime", reflect.TypeOf((*MockEthereumClock)(nil).GetSlotByTime), time)
	return &MockEthereumClockGetSlotByTimeCall{Call: call}
}

// MockEthereumClockGetSlotByTimeCall wrap *gomock.Call
type MockEthereumClockGetSlotByTimeCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockEthereumClockGetSlotByTimeCall) Return(arg0 uint64) *MockEthereumClockGetSlotByTimeCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockEthereumClockGetSlotByTimeCall) Do(f func(time.Time) uint64) *MockEthereumClockGetSlotByTimeCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockEthereumClockGetSlotByTimeCall) DoAndReturn(f func(time.Time) uint64) *MockEthereumClockGetSlotByTimeCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetSlotTime mocks base method.
func (m *MockEthereumClock) GetSlotTime(slot uint64) time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSlotTime", slot)
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// GetSlotTime indicates an expected call of GetSlotTime.
func (mr *MockEthereumClockMockRecorder) GetSlotTime(slot any) *MockEthereumClockGetSlotTimeCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSlotTime", reflect.TypeOf((*MockEthereumClock)(nil).GetSlotTime), slot)
	return &MockEthereumClockGetSlotTimeCall{Call: call}
}

// MockEthereumClockGetSlotTimeCall wrap *gomock.Call
type MockEthereumClockGetSlotTimeCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockEthereumClockGetSlotTimeCall) Return(arg0 time.Time) *MockEthereumClockGetSlotTimeCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockEthereumClockGetSlotTimeCall) Do(f func(uint64) time.Time) *MockEthereumClockGetSlotTimeCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockEthereumClockGetSlotTimeCall) DoAndReturn(f func(uint64) time.Time) *MockEthereumClockGetSlotTimeCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// IsSlotCurrentSlotWithMaximumClockDisparity mocks base method.
func (m *MockEthereumClock) IsSlotCurrentSlotWithMaximumClockDisparity(slot uint64) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsSlotCurrentSlotWithMaximumClockDisparity", slot)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsSlotCurrentSlotWithMaximumClockDisparity indicates an expected call of IsSlotCurrentSlotWithMaximumClockDisparity.
func (mr *MockEthereumClockMockRecorder) IsSlotCurrentSlotWithMaximumClockDisparity(slot any) *MockEthereumClockIsSlotCurrentSlotWithMaximumClockDisparityCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsSlotCurrentSlotWithMaximumClockDisparity", reflect.TypeOf((*MockEthereumClock)(nil).IsSlotCurrentSlotWithMaximumClockDisparity), slot)
	return &MockEthereumClockIsSlotCurrentSlotWithMaximumClockDisparityCall{Call: call}
}

// MockEthereumClockIsSlotCurrentSlotWithMaximumClockDisparityCall wrap *gomock.Call
type MockEthereumClockIsSlotCurrentSlotWithMaximumClockDisparityCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockEthereumClockIsSlotCurrentSlotWithMaximumClockDisparityCall) Return(arg0 bool) *MockEthereumClockIsSlotCurrentSlotWithMaximumClockDisparityCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockEthereumClockIsSlotCurrentSlotWithMaximumClockDisparityCall) Do(f func(uint64) bool) *MockEthereumClockIsSlotCurrentSlotWithMaximumClockDisparityCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockEthereumClockIsSlotCurrentSlotWithMaximumClockDisparityCall) DoAndReturn(f func(uint64) bool) *MockEthereumClockIsSlotCurrentSlotWithMaximumClockDisparityCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// LastFork mocks base method.
func (m *MockEthereumClock) LastFork() (common.Bytes4, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LastFork")
	ret0, _ := ret[0].(common.Bytes4)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LastFork indicates an expected call of LastFork.
func (mr *MockEthereumClockMockRecorder) LastFork() *MockEthereumClockLastForkCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LastFork", reflect.TypeOf((*MockEthereumClock)(nil).LastFork))
	return &MockEthereumClockLastForkCall{Call: call}
}

// MockEthereumClockLastForkCall wrap *gomock.Call
type MockEthereumClockLastForkCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockEthereumClockLastForkCall) Return(arg0 common.Bytes4, arg1 error) *MockEthereumClockLastForkCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockEthereumClockLastForkCall) Do(f func() (common.Bytes4, error)) *MockEthereumClockLastForkCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockEthereumClockLastForkCall) DoAndReturn(f func() (common.Bytes4, error)) *MockEthereumClockLastForkCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// NextForkDigest mocks base method.
func (m *MockEthereumClock) NextForkDigest() (common.Bytes4, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NextForkDigest")
	ret0, _ := ret[0].(common.Bytes4)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NextForkDigest indicates an expected call of NextForkDigest.
func (mr *MockEthereumClockMockRecorder) NextForkDigest() *MockEthereumClockNextForkDigestCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NextForkDigest", reflect.TypeOf((*MockEthereumClock)(nil).NextForkDigest))
	return &MockEthereumClockNextForkDigestCall{Call: call}
}

// MockEthereumClockNextForkDigestCall wrap *gomock.Call
type MockEthereumClockNextForkDigestCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockEthereumClockNextForkDigestCall) Return(arg0 common.Bytes4, arg1 error) *MockEthereumClockNextForkDigestCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockEthereumClockNextForkDigestCall) Do(f func() (common.Bytes4, error)) *MockEthereumClockNextForkDigestCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockEthereumClockNextForkDigestCall) DoAndReturn(f func() (common.Bytes4, error)) *MockEthereumClockNextForkDigestCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// StateVersionByForkDigest mocks base method.
func (m *MockEthereumClock) StateVersionByForkDigest(arg0 common.Bytes4) (clparams.StateVersion, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateVersionByForkDigest", arg0)
	ret0, _ := ret[0].(clparams.StateVersion)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateVersionByForkDigest indicates an expected call of StateVersionByForkDigest.
func (mr *MockEthereumClockMockRecorder) StateVersionByForkDigest(arg0 any) *MockEthereumClockStateVersionByForkDigestCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateVersionByForkDigest", reflect.TypeOf((*MockEthereumClock)(nil).StateVersionByForkDigest), arg0)
	return &MockEthereumClockStateVersionByForkDigestCall{Call: call}
}

// MockEthereumClockStateVersionByForkDigestCall wrap *gomock.Call
type MockEthereumClockStateVersionByForkDigestCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockEthereumClockStateVersionByForkDigestCall) Return(arg0 clparams.StateVersion, arg1 error) *MockEthereumClockStateVersionByForkDigestCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockEthereumClockStateVersionByForkDigestCall) Do(f func(common.Bytes4) (clparams.StateVersion, error)) *MockEthereumClockStateVersionByForkDigestCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockEthereumClockStateVersionByForkDigestCall) DoAndReturn(f func(common.Bytes4) (clparams.StateVersion, error)) *MockEthereumClockStateVersionByForkDigestCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
