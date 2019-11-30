// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// FollowCameraKeyboardMoveInput represents a babylon.js FollowCameraKeyboardMoveInput.
// Manage the keyboard inputs to control the movement of a follow camera.
//
// See: http://doc.babylonjs.com/how_to/customizing_camera_inputs
type FollowCameraKeyboardMoveInput struct{ p js.Value }

// JSObject returns the underlying js.Value.
func (f *FollowCameraKeyboardMoveInput) JSObject() js.Value { return f.p }

// FollowCameraKeyboardMoveInput returns a FollowCameraKeyboardMoveInput JavaScript class.
func (ba *Babylon) FollowCameraKeyboardMoveInput() *FollowCameraKeyboardMoveInput {
	p := ba.ctx.Get("FollowCameraKeyboardMoveInput")
	return FollowCameraKeyboardMoveInputFromJSObject(p)
}

// FollowCameraKeyboardMoveInputFromJSObject returns a wrapped FollowCameraKeyboardMoveInput JavaScript class.
func FollowCameraKeyboardMoveInputFromJSObject(p js.Value) *FollowCameraKeyboardMoveInput {
	return &FollowCameraKeyboardMoveInput{p: p}
}

// TODO: methods
