// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// XRWindowsMotionController represents a babylon.js XRWindowsMotionController.
// This class represents a new windows motion controller in XR.
type XRWindowsMotionController struct{ *WindowsMotionController }

// JSObject returns the underlying js.Value.
func (x *XRWindowsMotionController) JSObject() js.Value { return x.p }

// XRWindowsMotionController returns a XRWindowsMotionController JavaScript class.
func (ba *Babylon) XRWindowsMotionController() *XRWindowsMotionController {
	p := ba.ctx.Get("XRWindowsMotionController")
	return XRWindowsMotionControllerFromJSObject(p)
}

// XRWindowsMotionControllerFromJSObject returns a wrapped XRWindowsMotionController JavaScript class.
func XRWindowsMotionControllerFromJSObject(p js.Value) *XRWindowsMotionController {
	return &XRWindowsMotionController{WindowsMotionControllerFromJSObject(p)}
}

// NewXRWindowsMotionController returns a new XRWindowsMotionController object.
//
// https://doc.babylonjs.com/api/classes/babylon.xrwindowsmotioncontroller
func (ba *Babylon) NewXRWindowsMotionController(gamepadInfo interface{}) *XRWindowsMotionController {
	p := ba.ctx.Get("XRWindowsMotionController").New(gamepadInfo)
	return XRWindowsMotionControllerFromJSObject(p)
}

// TODO: methods
