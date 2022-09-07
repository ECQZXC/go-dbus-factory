// Code generated by "./generator ./org.deepin.daemon.greeter1"; DO NOT EDIT.

package greeter1

import "github.com/godbus/dbus"

import "github.com/linuxdeepin/go-lib/dbusutil/proxy"
import "unsafe"

type Greeter interface {
	greeter // interface org.deepin.daemon.Greeter1
	proxy.Object
}

type objectGreeter struct {
	interfaceGreeter // interface org.deepin.daemon.Greeter1
	proxy.ImplObject
}

func NewGreeter(conn *dbus.Conn) Greeter {
	obj := new(objectGreeter)
	obj.ImplObject.Init_(conn, "org.deepin.daemon.Greeter1", "/org/deepin/daemon/Greeter1")
	return obj
}

type greeter interface {
	GoUpdateGreeterQtTheme(flags dbus.Flags, ch chan *dbus.Call, fd dbus.UnixFD) *dbus.Call
	UpdateGreeterQtTheme(flags dbus.Flags, fd dbus.UnixFD) error
}

type interfaceGreeter struct{}

func (v *interfaceGreeter) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceGreeter) GetInterfaceName_() string {
	return "org.deepin.daemon.Greeter1"
}

// method UpdateGreeterQtTheme

func (v *interfaceGreeter) GoUpdateGreeterQtTheme(flags dbus.Flags, ch chan *dbus.Call, fd dbus.UnixFD) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".UpdateGreeterQtTheme", flags, ch, fd)
}

func (v *interfaceGreeter) UpdateGreeterQtTheme(flags dbus.Flags, fd dbus.UnixFD) error {
	return (<-v.GoUpdateGreeterQtTheme(flags, make(chan *dbus.Call, 1), fd).Done).Err
}
