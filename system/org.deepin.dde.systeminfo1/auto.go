// Code generated by "./generator ./system/org.deepin.dde.systeminfo1"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2023 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package systeminfo1

import "errors"

import "github.com/godbus/dbus/v5"

import "github.com/linuxdeepin/go-lib/dbusutil/proxy"
import "unsafe"

type SystemInfo interface {
	systemInfo // interface org.deepin.dde.SystemInfo1
	proxy.Object
}

type objectSystemInfo struct {
	interfaceSystemInfo // interface org.deepin.dde.SystemInfo1
	proxy.ImplObject
}

func NewSystemInfo(conn *dbus.Conn) SystemInfo {
	obj := new(objectSystemInfo)
	obj.ImplObject.Init_(conn, "org.deepin.dde.SystemInfo1", "/org/deepin/dde/SystemInfo1")
	return obj
}

type systemInfo interface {
	DMIInfo() PropDMIInfo
	MemorySize() proxy.PropUint64
	MemorySizeHuman() proxy.PropString
	CurrentSpeed() proxy.PropUint64
}

type interfaceSystemInfo struct{}

func (v *interfaceSystemInfo) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceSystemInfo) GetInterfaceName_() string {
	return "org.deepin.dde.SystemInfo1"
}

// property DMIInfo (ssssssssssss)

func (v *interfaceSystemInfo) DMIInfo() PropDMIInfo {
	return &implPropDMIInfo{
		Impl: v,
		Name: "DMIInfo",
	}
}

// property MemorySize t

func (v *interfaceSystemInfo) MemorySize() proxy.PropUint64 {
	return &proxy.ImplPropUint64{
		Impl: v,
		Name: "MemorySize",
	}
}

// property MemorySizeHuman s

func (v *interfaceSystemInfo) MemorySizeHuman() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "MemorySizeHuman",
	}
}

// property CurrentSpeed t

func (v *interfaceSystemInfo) CurrentSpeed() proxy.PropUint64 {
	return &proxy.ImplPropUint64{
		Impl: v,
		Name: "CurrentSpeed",
	}
}

type PropDMIInfo interface {
	Get(flags dbus.Flags) (value DMIInfo, err error)
	Set(flags dbus.Flags, value DMIInfo) error
	ConnectChanged(cb func(hasValue bool, value DMIInfo)) error
}

type implPropDMIInfo struct {
	Impl proxy.Implementer
	Name string
}

func (p implPropDMIInfo) Get(flags dbus.Flags) (value DMIInfo, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		p.Name, &value)
	return
}

func (p implPropDMIInfo) Set(flags dbus.Flags, value DMIInfo) error {
	return p.Impl.GetObject_().SetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, value)
}

func (p implPropDMIInfo) ConnectChanged(cb func(hasValue bool, value DMIInfo)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v DMIInfo
			err := dbus.Store([]interface{}{value}, &v)
			if err != nil {
				return
			}
			cb(true, v)
		} else {
			cb(false, DMIInfo{})
		}
	}
	return p.Impl.GetObject_().ConnectPropertyChanged_(p.Impl.GetInterfaceName_(),
		p.Name, cb0)
}