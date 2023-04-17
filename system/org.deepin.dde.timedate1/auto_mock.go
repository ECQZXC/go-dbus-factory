// Code generated by "./generator ./system/org.deepin.dde.timedate1"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2023 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package timedate1

import "fmt"
import "github.com/godbus/dbus/v5"

import "github.com/linuxdeepin/go-lib/dbusutil/proxy"
import "github.com/stretchr/testify/mock"

type MockTimedate struct {
	MockInterfaceTimedate // interface org.deepin.dde.Timedate1
	proxy.MockObject
}

type MockInterfaceTimedate struct {
	mock.Mock
}

// method SetLocalRTC

func (v *MockInterfaceTimedate) GoSetLocalRTC(flags dbus.Flags, ch chan *dbus.Call, enabled bool, fixSystem bool, message string) *dbus.Call {
	mockArgs := v.Called(flags, ch, enabled, fixSystem, message)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceTimedate) SetLocalRTC(flags dbus.Flags, enabled bool, fixSystem bool, message string) error {
	mockArgs := v.Called(flags, enabled, fixSystem, message)

	return mockArgs.Error(0)
}

// method SetNTP

func (v *MockInterfaceTimedate) GoSetNTP(flags dbus.Flags, ch chan *dbus.Call, enabled bool, message string) *dbus.Call {
	mockArgs := v.Called(flags, ch, enabled, message)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceTimedate) SetNTP(flags dbus.Flags, enabled bool, message string) error {
	mockArgs := v.Called(flags, enabled, message)

	return mockArgs.Error(0)
}

// method SetNTPServer

func (v *MockInterfaceTimedate) GoSetNTPServer(flags dbus.Flags, ch chan *dbus.Call, server string, message string) *dbus.Call {
	mockArgs := v.Called(flags, ch, server, message)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceTimedate) SetNTPServer(flags dbus.Flags, server string, message string) error {
	mockArgs := v.Called(flags, server, message)

	return mockArgs.Error(0)
}

// method SetTime

func (v *MockInterfaceTimedate) GoSetTime(flags dbus.Flags, ch chan *dbus.Call, usec int64, relative bool, message string) *dbus.Call {
	mockArgs := v.Called(flags, ch, usec, relative, message)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceTimedate) SetTime(flags dbus.Flags, usec int64, relative bool, message string) error {
	mockArgs := v.Called(flags, usec, relative, message)

	return mockArgs.Error(0)
}

// method SetTimezone

func (v *MockInterfaceTimedate) GoSetTimezone(flags dbus.Flags, ch chan *dbus.Call, timezone string, message string) *dbus.Call {
	mockArgs := v.Called(flags, ch, timezone, message)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceTimedate) SetTimezone(flags dbus.Flags, timezone string, message string) error {
	mockArgs := v.Called(flags, timezone, message)

	return mockArgs.Error(0)
}

// property NTPServer s

func (v *MockInterfaceTimedate) NTPServer() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Timezone s

func (v *MockInterfaceTimedate) Timezone() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}