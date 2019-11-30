// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// FollowCameraPointersInput represents a babylon.js FollowCameraPointersInput.
// Manage the pointers inputs to control an follow camera.
//
// See: http://doc.babylonjs.com/how_to/customizing_camera_inputs
type FollowCameraPointersInput struct{ *BaseCameraPointersInput }

// JSObject returns the underlying js.Value.
func (f *FollowCameraPointersInput) JSObject() js.Value { return f.p }

// FollowCameraPointersInput returns a FollowCameraPointersInput JavaScript class.
func (ba *Babylon) FollowCameraPointersInput() *FollowCameraPointersInput {
	p := ba.ctx.Get("FollowCameraPointersInput")
	return FollowCameraPointersInputFromJSObject(p)
}

// FollowCameraPointersInputFromJSObject returns a wrapped FollowCameraPointersInput JavaScript class.
func FollowCameraPointersInputFromJSObject(p js.Value) *FollowCameraPointersInput {
	return &FollowCameraPointersInput{BaseCameraPointersInputFromJSObject(p)}
}

// TODO: methods
