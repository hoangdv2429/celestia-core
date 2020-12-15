// Code generated by mockery v2.4.0-beta. DO NOT EDIT.

package mocks

import (
	log "github.com/lazyledger/lazyledger-core/libs/log"
	conn "github.com/lazyledger/lazyledger-core/p2p/conn"

	mock "github.com/stretchr/testify/mock"

	net "net"

	p2p "github.com/lazyledger/lazyledger-core/p2p"
)

// Peer is an autogenerated mock type for the Peer type
type Peer struct {
	mock.Mock
}

// CloseConn provides a mock function with given fields:
func (_m *Peer) CloseConn() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FlushStop provides a mock function with given fields:
func (_m *Peer) FlushStop() {
	_m.Called()
}

// Get provides a mock function with given fields: _a0
func (_m *Peer) Get(_a0 string) interface{} {
	ret := _m.Called(_a0)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(string) interface{}); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// ID provides a mock function with given fields:
func (_m *Peer) ID() p2p.ID {
	ret := _m.Called()

	var r0 p2p.ID
	if rf, ok := ret.Get(0).(func() p2p.ID); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(p2p.ID)
	}

	return r0
}

// IsOutbound provides a mock function with given fields:
func (_m *Peer) IsOutbound() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// IsPersistent provides a mock function with given fields:
func (_m *Peer) IsPersistent() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// IsRunning provides a mock function with given fields:
func (_m *Peer) IsRunning() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// NodeInfo provides a mock function with given fields:
func (_m *Peer) NodeInfo() p2p.NodeInfo {
	ret := _m.Called()

	var r0 p2p.NodeInfo
	if rf, ok := ret.Get(0).(func() p2p.NodeInfo); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(p2p.NodeInfo)
		}
	}

	return r0
}

// OnReset provides a mock function with given fields:
func (_m *Peer) OnReset() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OnStart provides a mock function with given fields:
func (_m *Peer) OnStart() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OnStop provides a mock function with given fields:
func (_m *Peer) OnStop() {
	_m.Called()
}

// Quit provides a mock function with given fields:
func (_m *Peer) Quit() <-chan struct{} {
	ret := _m.Called()

	var r0 <-chan struct{}
	if rf, ok := ret.Get(0).(func() <-chan struct{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan struct{})
		}
	}

	return r0
}

// RemoteAddr provides a mock function with given fields:
func (_m *Peer) RemoteAddr() net.Addr {
	ret := _m.Called()

	var r0 net.Addr
	if rf, ok := ret.Get(0).(func() net.Addr); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(net.Addr)
		}
	}

	return r0
}

// RemoteIP provides a mock function with given fields:
func (_m *Peer) RemoteIP() net.IP {
	ret := _m.Called()

	var r0 net.IP
	if rf, ok := ret.Get(0).(func() net.IP); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(net.IP)
		}
	}

	return r0
}

// Reset provides a mock function with given fields:
func (_m *Peer) Reset() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Send provides a mock function with given fields: _a0, _a1
func (_m *Peer) Send(_a0 byte, _a1 []byte) bool {
	ret := _m.Called(_a0, _a1)

	var r0 bool
	if rf, ok := ret.Get(0).(func(byte, []byte) bool); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Set provides a mock function with given fields: _a0, _a1
func (_m *Peer) Set(_a0 string, _a1 interface{}) {
	_m.Called(_a0, _a1)
}

// SetLogger provides a mock function with given fields: _a0
func (_m *Peer) SetLogger(_a0 log.Logger) {
	_m.Called(_a0)
}

// SocketAddr provides a mock function with given fields:
func (_m *Peer) SocketAddr() *p2p.NetAddress {
	ret := _m.Called()

	var r0 *p2p.NetAddress
	if rf, ok := ret.Get(0).(func() *p2p.NetAddress); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*p2p.NetAddress)
		}
	}

	return r0
}

// Start provides a mock function with given fields:
func (_m *Peer) Start() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Status provides a mock function with given fields:
func (_m *Peer) Status() conn.ConnectionStatus {
	ret := _m.Called()

	var r0 conn.ConnectionStatus
	if rf, ok := ret.Get(0).(func() conn.ConnectionStatus); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(conn.ConnectionStatus)
	}

	return r0
}

// Stop provides a mock function with given fields:
func (_m *Peer) Stop() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// String provides a mock function with given fields:
func (_m *Peer) String() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// TrySend provides a mock function with given fields: _a0, _a1
func (_m *Peer) TrySend(_a0 byte, _a1 []byte) bool {
	ret := _m.Called(_a0, _a1)

	var r0 bool
	if rf, ok := ret.Get(0).(func(byte, []byte) bool); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}
