package authenticate

import "errors"
import "fmt"
import "github.com/godbus/dbus"
import "pkg.deepin.io/lib/dbusutil"
import "pkg.deepin.io/lib/dbusutil/proxy"
import "unsafe"

/* prevent compile error */
var _ = errors.New
var _ dbusutil.SignalHandlerId
var _ = fmt.Sprintf
var _ unsafe.Pointer

type Authenticate struct {
	authenticate // interface com.deepin.daemon.Authenticate
	proxy.Object
}

func NewAuthenticate(conn *dbus.Conn) *Authenticate {
	obj := new(Authenticate)
	obj.Object.Init_(conn, "com.deepin.daemon.Authenticate", "/com/deepin/daemon/Authenticate")
	return obj
}

type authenticate struct{}

func (v *authenticate) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*authenticate) GetInterfaceName_() string {
	return "com.deepin.daemon.Authenticate"
}

// method Authenticate

func (v *authenticate) GoAuthenticate(flags dbus.Flags, ch chan *dbus.Call, username string, authFlags int32, appType int32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Authenticate", flags, ch, username, authFlags, appType)
}

func (*authenticate) StoreAuthenticate(call *dbus.Call) (path string, err error) {
	err = call.Store(&path)
	return
}

func (v *authenticate) Authenticate(flags dbus.Flags, username string, authFlags int32, appType int32) (path string, err error) {
	return v.StoreAuthenticate(
		<-v.GoAuthenticate(flags, make(chan *dbus.Call, 1), username, authFlags, appType).Done)
}

// method GetLimits

func (v *authenticate) GoGetLimits(flags dbus.Flags, ch chan *dbus.Call, username string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetLimits", flags, ch, username)
}

func (*authenticate) StoreGetLimits(call *dbus.Call) (limits string, err error) {
	err = call.Store(&limits)
	return
}

func (v *authenticate) GetLimits(flags dbus.Flags, username string) (limits string, err error) {
	return v.StoreGetLimits(
		<-v.GoGetLimits(flags, make(chan *dbus.Call, 1), username).Done)
}

// method PreOneKeyLogin

func (v *authenticate) GoPreOneKeyLogin(flags dbus.Flags, ch chan *dbus.Call, flag int32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".PreOneKeyLogin", flags, ch, flag)
}

func (*authenticate) StorePreOneKeyLogin(call *dbus.Call) (result string, err error) {
	err = call.Store(&result)
	return
}

func (v *authenticate) PreOneKeyLogin(flags dbus.Flags, flag int32) (result string, err error) {
	return v.StorePreOneKeyLogin(
		<-v.GoPreOneKeyLogin(flags, make(chan *dbus.Call, 1), flag).Done)
}

// signal LimitUpdated

func (v *authenticate) ConnectLimitUpdated(cb func(username string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "LimitUpdated", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".LimitUpdated",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var username string
		err := dbus.Store(sig.Body, &username)
		if err == nil {
			cb(username)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property SupportEncrypts s

func (v *authenticate) SupportEncrypts() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "SupportEncrypts",
	}
}

// property FrameworkState i

func (v *authenticate) FrameworkState() proxy.PropInt32 {
	return proxy.PropInt32{
		Impl: v,
		Name: "FrameworkState",
	}
}

// property SupportedFlags i

func (v *authenticate) SupportedFlags() proxy.PropInt32 {
	return proxy.PropInt32{
		Impl: v,
		Name: "SupportedFlags",
	}
}

type Fingerprint struct {
	fingerprint // interface com.deepin.daemon.Authenticate.Fingerprint
	proxy.Object
}

func NewFingerprint(conn *dbus.Conn) *Fingerprint {
	obj := new(Fingerprint)
	obj.Object.Init_(conn, "com.deepin.daemon.Authenticate", "/com/deepin/daemon/Authenticate/Fingerprint")
	return obj
}

type fingerprint struct{}

func (v *fingerprint) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*fingerprint) GetInterfaceName_() string {
	return "com.deepin.daemon.Authenticate.Fingerprint"
}

// method Claim

func (v *fingerprint) GoClaim(flags dbus.Flags, ch chan *dbus.Call, username string, claimed bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Claim", flags, ch, username, claimed)
}

func (v *fingerprint) Claim(flags dbus.Flags, username string, claimed bool) error {
	return (<-v.GoClaim(flags, make(chan *dbus.Call, 1), username, claimed).Done).Err
}

// method DeleteAllFingers

func (v *fingerprint) GoDeleteAllFingers(flags dbus.Flags, ch chan *dbus.Call, username string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".DeleteAllFingers", flags, ch, username)
}

func (v *fingerprint) DeleteAllFingers(flags dbus.Flags, username string) error {
	return (<-v.GoDeleteAllFingers(flags, make(chan *dbus.Call, 1), username).Done).Err
}

// method DeleteFinger

func (v *fingerprint) GoDeleteFinger(flags dbus.Flags, ch chan *dbus.Call, username string, finger string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".DeleteFinger", flags, ch, username, finger)
}

func (v *fingerprint) DeleteFinger(flags dbus.Flags, username string, finger string) error {
	return (<-v.GoDeleteFinger(flags, make(chan *dbus.Call, 1), username, finger).Done).Err
}

// method Enroll

func (v *fingerprint) GoEnroll(flags dbus.Flags, ch chan *dbus.Call, finger string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Enroll", flags, ch, finger)
}

func (v *fingerprint) Enroll(flags dbus.Flags, finger string) error {
	return (<-v.GoEnroll(flags, make(chan *dbus.Call, 1), finger).Done).Err
}

// method ListFingers

func (v *fingerprint) GoListFingers(flags dbus.Flags, ch chan *dbus.Call, username string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ListFingers", flags, ch, username)
}

func (*fingerprint) StoreListFingers(call *dbus.Call) (fingers []string, err error) {
	err = call.Store(&fingers)
	return
}

func (v *fingerprint) ListFingers(flags dbus.Flags, username string) (fingers []string, err error) {
	return v.StoreListFingers(
		<-v.GoListFingers(flags, make(chan *dbus.Call, 1), username).Done)
}

// method SetDefaultDevice

func (v *fingerprint) GoSetDefaultDevice(flags dbus.Flags, ch chan *dbus.Call, device string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetDefaultDevice", flags, ch, device)
}

func (v *fingerprint) SetDefaultDevice(flags dbus.Flags, device string) error {
	return (<-v.GoSetDefaultDevice(flags, make(chan *dbus.Call, 1), device).Done).Err
}

// method StopEnroll

func (v *fingerprint) GoStopEnroll(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".StopEnroll", flags, ch)
}

func (v *fingerprint) StopEnroll(flags dbus.Flags) error {
	return (<-v.GoStopEnroll(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method StopVerify

func (v *fingerprint) GoStopVerify(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".StopVerify", flags, ch)
}

func (v *fingerprint) StopVerify(flags dbus.Flags) error {
	return (<-v.GoStopVerify(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method Verify

func (v *fingerprint) GoVerify(flags dbus.Flags, ch chan *dbus.Call, finger string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Verify", flags, ch, finger)
}

func (v *fingerprint) Verify(flags dbus.Flags, finger string) error {
	return (<-v.GoVerify(flags, make(chan *dbus.Call, 1), finger).Done).Err
}

// signal EnrollStatus

func (v *fingerprint) ConnectEnrollStatus(cb func(id string, code int32, msg string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "EnrollStatus", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".EnrollStatus",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var id string
		var code int32
		var msg string
		err := dbus.Store(sig.Body, &id, &code, &msg)
		if err == nil {
			cb(id, code, msg)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal VerifyStatus

func (v *fingerprint) ConnectVerifyStatus(cb func(id string, code int32, msg string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "VerifyStatus", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".VerifyStatus",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var id string
		var code int32
		var msg string
		err := dbus.Store(sig.Body, &id, &code, &msg)
		if err == nil {
			cb(id, code, msg)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal Touch

func (v *fingerprint) ConnectTouch(cb func(id string, pressed bool)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "Touch", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".Touch",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var id string
		var pressed bool
		err := dbus.Store(sig.Body, &id, &pressed)
		if err == nil {
			cb(id, pressed)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property DefaultDevice s

func (v *fingerprint) DefaultDevice() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "DefaultDevice",
	}
}

// property Devices s

func (v *fingerprint) Devices() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Devices",
	}
}

type UKey struct {
	ukey // interface com.deepin.daemon.Authenticate.UKey
	proxy.Object
}

func NewUKey(conn *dbus.Conn) *UKey {
	obj := new(UKey)
	obj.Object.Init_(conn, "com.deepin.daemon.Authenticate", "/com/deepin/daemon/Authenticate/UKey")
	return obj
}

type ukey struct{}

func (v *ukey) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*ukey) GetInterfaceName_() string {
	return "com.deepin.daemon.Authenticate.UKey"
}

// method ConstructVerification

func (v *ukey) GoConstructVerification(flags dbus.Flags, ch chan *dbus.Call, serviceName string, username string, useDefaultService bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ConstructVerification", flags, ch, serviceName, username, useDefaultService)
}

func (*ukey) StoreConstructVerification(call *dbus.Call) (id string, err error) {
	err = call.Store(&id)
	return
}

func (v *ukey) ConstructVerification(flags dbus.Flags, serviceName string, username string, useDefaultService bool) (id string, err error) {
	return v.StoreConstructVerification(
		<-v.GoConstructVerification(flags, make(chan *dbus.Call, 1), serviceName, username, useDefaultService).Done)
}

// method SetDefaultDevice

func (v *ukey) GoSetDefaultDevice(flags dbus.Flags, ch chan *dbus.Call, device string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetDefaultDevice", flags, ch, device)
}

func (v *ukey) SetDefaultDevice(flags dbus.Flags, device string) error {
	return (<-v.GoSetDefaultDevice(flags, make(chan *dbus.Call, 1), device).Done).Err
}

// method SetPin

func (v *ukey) GoSetPin(flags dbus.Flags, ch chan *dbus.Call, id string, pin string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetPin", flags, ch, id, pin)
}

func (v *ukey) SetPin(flags dbus.Flags, id string, pin string) error {
	return (<-v.GoSetPin(flags, make(chan *dbus.Call, 1), id, pin).Done).Err
}

// method SetSessionPath

func (v *ukey) GoSetSessionPath(flags dbus.Flags, ch chan *dbus.Call, id string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetSessionPath", flags, ch, id)
}

func (v *ukey) SetSessionPath(flags dbus.Flags, id string) error {
	return (<-v.GoSetSessionPath(flags, make(chan *dbus.Call, 1), id).Done).Err
}

// method StartVerify

func (v *ukey) GoStartVerify(flags dbus.Flags, ch chan *dbus.Call, id string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".StartVerify", flags, ch, id)
}

func (v *ukey) StartVerify(flags dbus.Flags, id string) error {
	return (<-v.GoStartVerify(flags, make(chan *dbus.Call, 1), id).Done).Err
}

// method StopVerify

func (v *ukey) GoStopVerify(flags dbus.Flags, ch chan *dbus.Call, id string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".StopVerify", flags, ch, id)
}

func (v *ukey) StopVerify(flags dbus.Flags, id string) error {
	return (<-v.GoStopVerify(flags, make(chan *dbus.Call, 1), id).Done).Err
}

// signal VerifyResult

func (v *ukey) ConnectVerifyResult(cb func(id string, msg string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "VerifyResult", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".VerifyResult",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var id string
		var msg string
		err := dbus.Store(sig.Body, &id, &msg)
		if err == nil {
			cb(id, msg)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal State

func (v *ukey) ConnectState(cb func(id string, state int32)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "State", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".State",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var id string
		var state int32
		err := dbus.Store(sig.Body, &id, &state)
		if err == nil {
			cb(id, state)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property ValidDevices s

func (v *ukey) ValidDevices() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "ValidDevices",
	}
}

// property DefaultDevice s

func (v *ukey) DefaultDevice() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "DefaultDevice",
	}
}
