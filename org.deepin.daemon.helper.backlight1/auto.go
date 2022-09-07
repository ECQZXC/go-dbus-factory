// Code generated by "./generator ./org.deepin.daemon.helper.backlight1"; DO NOT EDIT.

package backlight1

import "github.com/godbus/dbus"

import "github.com/linuxdeepin/go-lib/dbusutil/proxy"
import "unsafe"

type Backlight interface {
	backlight // interface org.deepin.daemon.helper.Backlight1
	proxy.Object
}

type objectBacklight struct {
	interfaceBacklight // interface org.deepin.daemon.helper.Backlight1
	proxy.ImplObject
}

func NewBacklight(conn *dbus.Conn) Backlight {
	obj := new(objectBacklight)
	obj.ImplObject.Init_(conn, "org.deepin.daemon.helper.Backlight1", "/org/deepin/daemon/helper/Backlight1")
	return obj
}

type backlight interface {
	GoSetBrightness(flags dbus.Flags, ch chan *dbus.Call, type0 uint8, name string, value int32) *dbus.Call
	SetBrightness(flags dbus.Flags, type0 uint8, name string, value int32) error
}

type interfaceBacklight struct{}

func (v *interfaceBacklight) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceBacklight) GetInterfaceName_() string {
	return "org.deepin.daemon.helper.Backlight1"
}

// method SetBrightness

func (v *interfaceBacklight) GoSetBrightness(flags dbus.Flags, ch chan *dbus.Call, type0 uint8, name string, value int32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetBrightness", flags, ch, type0, name, value)
}

func (v *interfaceBacklight) SetBrightness(flags dbus.Flags, type0 uint8, name string, value int32) error {
	return (<-v.GoSetBrightness(flags, make(chan *dbus.Call, 1), type0, name, value).Done).Err
}

type DDCCI interface {
	ddcci // interface org.deepin.daemon.helper.Backlight1.DDCCI
	proxy.Object
}

type objectDDCCI struct {
	interfaceDdcci // interface org.deepin.daemon.helper.Backlight1.DDCCI
	proxy.ImplObject
}

func NewDDCCI(conn *dbus.Conn) DDCCI {
	obj := new(objectDDCCI)
	obj.ImplObject.Init_(conn, "org.deepin.daemon.helper.Backlight1", "/org/deepin/daemon/helper/Backlight1/DDCCI")
	return obj
}

type ddcci interface {
	GoCheckSupport(flags dbus.Flags, ch chan *dbus.Call, edidChecksum string) *dbus.Call
	CheckSupport(flags dbus.Flags, edidChecksum string) (bool, error)
	GoGetBrightness(flags dbus.Flags, ch chan *dbus.Call, edidChecksum string) *dbus.Call
	GetBrightness(flags dbus.Flags, edidChecksum string) (int32, error)
	GoSetBrightness(flags dbus.Flags, ch chan *dbus.Call, edidChecksum string, value int32) *dbus.Call
	SetBrightness(flags dbus.Flags, edidChecksum string, value int32) error
	GoRefreshDisplays(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	RefreshDisplays(flags dbus.Flags) error
}

type interfaceDdcci struct{}

func (v *interfaceDdcci) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceDdcci) GetInterfaceName_() string {
	return "org.deepin.daemon.helper.Backlight1.DDCCI"
}

// method CheckSupport

func (v *interfaceDdcci) GoCheckSupport(flags dbus.Flags, ch chan *dbus.Call, edidChecksum string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CheckSupport", flags, ch, edidChecksum)
}

func (*interfaceDdcci) StoreCheckSupport(call *dbus.Call) (support bool, err error) {
	err = call.Store(&support)
	return
}

func (v *interfaceDdcci) CheckSupport(flags dbus.Flags, edidChecksum string) (bool, error) {
	return v.StoreCheckSupport(
		<-v.GoCheckSupport(flags, make(chan *dbus.Call, 1), edidChecksum).Done)
}

// method GetBrightness

func (v *interfaceDdcci) GoGetBrightness(flags dbus.Flags, ch chan *dbus.Call, edidChecksum string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetBrightness", flags, ch, edidChecksum)
}

func (*interfaceDdcci) StoreGetBrightness(call *dbus.Call) (value int32, err error) {
	err = call.Store(&value)
	return
}

func (v *interfaceDdcci) GetBrightness(flags dbus.Flags, edidChecksum string) (int32, error) {
	return v.StoreGetBrightness(
		<-v.GoGetBrightness(flags, make(chan *dbus.Call, 1), edidChecksum).Done)
}

// method SetBrightness

func (v *interfaceDdcci) GoSetBrightness(flags dbus.Flags, ch chan *dbus.Call, edidChecksum string, value int32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetBrightness", flags, ch, edidChecksum, value)
}

func (v *interfaceDdcci) SetBrightness(flags dbus.Flags, edidChecksum string, value int32) error {
	return (<-v.GoSetBrightness(flags, make(chan *dbus.Call, 1), edidChecksum, value).Done).Err
}

// method RefreshDisplays

func (v *interfaceDdcci) GoRefreshDisplays(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RefreshDisplays", flags, ch)
}

func (v *interfaceDdcci) RefreshDisplays(flags dbus.Flags) error {
	return (<-v.GoRefreshDisplays(flags, make(chan *dbus.Call, 1)).Done).Err
}
