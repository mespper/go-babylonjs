// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// VirtualJoysticksCamera represents a babylon.js VirtualJoysticksCamera.
// This represents a free type of camera. It can be useful in First Person Shooter game for instance.
// It is identical to the Free Camera and simply adds by default a virtual joystick.
// Virtual Joysticks are on-screen 2D graphics that are used to control the camera or other scene items.
//
// See: http://doc.babylonjs.com/features/cameras#virtual-joysticks-camera
type VirtualJoysticksCamera struct{ *FreeCamera }

// JSObject returns the underlying js.Value.
func (v *VirtualJoysticksCamera) JSObject() js.Value { return v.p }

// VirtualJoysticksCamera returns a VirtualJoysticksCamera JavaScript class.
func (ba *Babylon) VirtualJoysticksCamera() *VirtualJoysticksCamera {
	p := ba.ctx.Get("VirtualJoysticksCamera")
	return VirtualJoysticksCameraFromJSObject(p)
}

// VirtualJoysticksCameraFromJSObject returns a wrapped VirtualJoysticksCamera JavaScript class.
func VirtualJoysticksCameraFromJSObject(p js.Value) *VirtualJoysticksCamera {
	return &VirtualJoysticksCamera{FreeCameraFromJSObject(p)}
}

// NewVirtualJoysticksCamera returns a new VirtualJoysticksCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.virtualjoystickscamera
func (ba *Babylon) NewVirtualJoysticksCamera(name string, position *Vector3, scene *Scene) *VirtualJoysticksCamera {
	p := ba.ctx.Get("VirtualJoysticksCamera").New(name, position.JSObject(), scene.JSObject())
	return VirtualJoysticksCameraFromJSObject(p)
}

// TODO: methods
