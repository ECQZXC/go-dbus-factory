// Code generated by "./generator ./org.deepin.system.systeminfo1"; DO NOT EDIT.

package systeminfo1

import "github.com/godbus/dbus"

import "github.com/linuxdeepin/go-lib/dbusutil/proxy"
import "unsafe"

type SystemInfo interface {
	systemInfo // interface org.deepin.system.SystemInfo1
	proxy.Object
}

type objectSystemInfo struct {
	interfaceSystemInfo // interface org.deepin.system.SystemInfo1
	proxy.ImplObject
}

func NewSystemInfo(conn *dbus.Conn) SystemInfo {
	obj := new(objectSystemInfo)
	obj.ImplObject.Init_(conn, "org.deepin.system.SystemInfo1", "/org/deepin/system/SystemInfo1")
	return obj
}

type systemInfo interface {
	MemorySizeHuman() proxy.PropString
	MemorySize() proxy.PropUint64
	CurrentSpeed() proxy.PropUint64
}

type interfaceSystemInfo struct{}

func (v *interfaceSystemInfo) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceSystemInfo) GetInterfaceName_() string {
	return "org.deepin.system.SystemInfo1"
}

// property MemorySizeHuman s

func (v *interfaceSystemInfo) MemorySizeHuman() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "MemorySizeHuman",
	}
}

// property MemorySize t

func (v *interfaceSystemInfo) MemorySize() proxy.PropUint64 {
	return &proxy.ImplPropUint64{
		Impl: v,
		Name: "MemorySize",
	}
}

// property CurrentSpeed t

func (v *interfaceSystemInfo) CurrentSpeed() proxy.PropUint64 {
	return &proxy.ImplPropUint64{
		Impl: v,
		Name: "CurrentSpeed",
	}
}
