// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// TargetCamera represents a babylon.js TargetCamera.
// A target camera takes a mesh or position as a target and continues to look at it while it moves.
// This is the base of the follow, arc rotate cameras and Free camera
//
// See: http://doc.babylonjs.com/features/cameras
type TargetCamera struct{ *Camera }

// JSObject returns the underlying js.Value.
func (t *TargetCamera) JSObject() js.Value { return t.p }

// TargetCamera returns a TargetCamera JavaScript class.
func (ba *Babylon) TargetCamera() *TargetCamera {
	p := ba.ctx.Get("TargetCamera")
	return TargetCameraFromJSObject(p)
}

// TargetCameraFromJSObject returns a wrapped TargetCamera JavaScript class.
func TargetCameraFromJSObject(p js.Value) *TargetCamera {
	return &TargetCamera{CameraFromJSObject(p)}
}

// NewTargetCameraOpts contains optional parameters for NewTargetCamera.
type NewTargetCameraOpts struct {
	SetActiveOnSceneIfNoneActive *JSBool
}

// NewTargetCamera returns a new TargetCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.targetcamera
func (ba *Babylon) NewTargetCamera(name string, position *Vector3, scene *Scene, opts *NewTargetCameraOpts) *TargetCamera {
	if opts == nil {
		opts = &NewTargetCameraOpts{}
	}

	p := ba.ctx.Get("TargetCamera").New(name, position.JSObject(), scene.JSObject(), opts.SetActiveOnSceneIfNoneActive.JSObject())
	return TargetCameraFromJSObject(p)
}

// TODO: methods
