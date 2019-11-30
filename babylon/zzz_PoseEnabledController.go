// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// PoseEnabledController represents a babylon.js PoseEnabledController.
// Defines the PoseEnabledController object that contains state of a vr capable controller
type PoseEnabledController struct{ *Gamepad }

// JSObject returns the underlying js.Value.
func (p *PoseEnabledController) JSObject() js.Value { return p.p }

// PoseEnabledController returns a PoseEnabledController JavaScript class.
func (ba *Babylon) PoseEnabledController() *PoseEnabledController {
	p := ba.ctx.Get("PoseEnabledController")
	return PoseEnabledControllerFromJSObject(p)
}

// PoseEnabledControllerFromJSObject returns a wrapped PoseEnabledController JavaScript class.
func PoseEnabledControllerFromJSObject(p js.Value) *PoseEnabledController {
	return &PoseEnabledController{GamepadFromJSObject(p)}
}

// NewPoseEnabledController returns a new PoseEnabledController object.
//
// https://doc.babylonjs.com/api/classes/babylon.poseenabledcontroller
func (ba *Babylon) NewPoseEnabledController(browserGamepad interface{}) *PoseEnabledController {
	p := ba.ctx.Get("PoseEnabledController").New(browserGamepad)
	return PoseEnabledControllerFromJSObject(p)
}

// TODO: methods
