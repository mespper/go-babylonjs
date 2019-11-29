// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// FollowCameraInputsManager represents a babylon.js FollowCameraInputsManager.
// Default Inputs manager for the FollowCamera.
// It groups all the default supported inputs for ease of use.

//
// See: http://doc.babylonjs.com/how_to/customizing_camera_inputs
type FollowCameraInputsManager struct {
	*CameraInputsManager
	*FollowCamera
}

// JSObject returns the underlying js.Value.
func (f *FollowCameraInputsManager) JSObject() js.Value { return f.p }

// FollowCameraInputsManager returns a FollowCameraInputsManager JavaScript class.
func (b *Babylon) FollowCameraInputsManager() *FollowCameraInputsManager {
	p := b.ctx.Get("FollowCameraInputsManager")
	return FollowCameraInputsManagerFromJSObject(p)
}

// FollowCameraInputsManagerFromJSObject returns a wrapped FollowCameraInputsManager JavaScript class.
func FollowCameraInputsManagerFromJSObject(p js.Value) *FollowCameraInputsManager {
	return &FollowCameraInputsManager{CameraInputsManagerFromJSObject(p)}
}

// NewFollowCameraInputsManager returns a new FollowCameraInputsManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.followcamerainputsmanager
func (b *Babylon) NewFollowCameraInputsManager(todo parameters) *FollowCameraInputsManager {
	p := b.ctx.Get("FollowCameraInputsManager").New(todo)
	return FollowCameraInputsManagerFromJSObject(p)
}

// TODO: methods