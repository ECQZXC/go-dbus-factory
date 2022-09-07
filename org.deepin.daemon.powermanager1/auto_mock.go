// Code generated by "./generator ./org.deepin.daemon.powermanager1"; DO NOT EDIT.

package powermanager1

import "fmt"
import "github.com/godbus/dbus"

import "github.com/linuxdeepin/go-lib/dbusutil/proxy"
import "github.com/stretchr/testify/mock"

type MockPowerManager struct {
	MockInterfacePowerManager // interface org.deepin.daemon.PowerManager1
	proxy.MockObject
}

type MockInterfacePowerManager struct {
	mock.Mock
}

// method CanShutdown

func (v *MockInterfacePowerManager) GoCanShutdown(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfacePowerManager) CanShutdown(flags dbus.Flags) (bool, error) {
	mockArgs := v.Called(flags)

	return mockArgs.Bool(0), mockArgs.Error(1)
}

// method CanReboot

func (v *MockInterfacePowerManager) GoCanReboot(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfacePowerManager) CanReboot(flags dbus.Flags) (bool, error) {
	mockArgs := v.Called(flags)

	return mockArgs.Bool(0), mockArgs.Error(1)
}

// method CanSuspend

func (v *MockInterfacePowerManager) GoCanSuspend(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfacePowerManager) CanSuspend(flags dbus.Flags) (bool, error) {
	mockArgs := v.Called(flags)

	return mockArgs.Bool(0), mockArgs.Error(1)
}

// method CanHibernate

func (v *MockInterfacePowerManager) GoCanHibernate(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfacePowerManager) CanHibernate(flags dbus.Flags) (bool, error) {
	mockArgs := v.Called(flags)

	return mockArgs.Bool(0), mockArgs.Error(1)
}
