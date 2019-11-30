// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ViveController represents a babylon.js ViveController.
// Vive Controller
type ViveController struct{ *WebVRController }

// JSObject returns the underlying js.Value.
func (v *ViveController) JSObject() js.Value { return v.p }

// ViveController returns a ViveController JavaScript class.
func (ba *Babylon) ViveController() *ViveController {
	p := ba.ctx.Get("ViveController")
	return ViveControllerFromJSObject(p)
}

// ViveControllerFromJSObject returns a wrapped ViveController JavaScript class.
func ViveControllerFromJSObject(p js.Value) *ViveController {
	return &ViveController{WebVRControllerFromJSObject(p)}
}

// NewViveController returns a new ViveController object.
//
// https://doc.babylonjs.com/api/classes/babylon.vivecontroller
func (ba *Babylon) NewViveController(vrGamepad interface{}) *ViveController {
	p := ba.ctx.Get("ViveController").New(vrGamepad)
	return ViveControllerFromJSObject(p)
}

// TODO: methods
