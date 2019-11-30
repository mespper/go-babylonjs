// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// FollowCamera represents a babylon.js FollowCamera.
// A follow camera takes a mesh as a target and follows it as it moves. Both a free camera version followCamera and
// an arc rotate version arcFollowCamera are available.
//
// See: http://doc.babylonjs.com/features/cameras#follow-camera
type FollowCamera struct{ *TargetCamera }

// JSObject returns the underlying js.Value.
func (f *FollowCamera) JSObject() js.Value { return f.p }

// FollowCamera returns a FollowCamera JavaScript class.
func (ba *Babylon) FollowCamera() *FollowCamera {
	p := ba.ctx.Get("FollowCamera")
	return FollowCameraFromJSObject(p)
}

// FollowCameraFromJSObject returns a wrapped FollowCamera JavaScript class.
func FollowCameraFromJSObject(p js.Value) *FollowCamera {
	return &FollowCamera{TargetCameraFromJSObject(p)}
}

// NewFollowCameraOpts contains optional parameters for NewFollowCamera.
type NewFollowCameraOpts struct {
	LockedTarget *AbstractMesh
}

// NewFollowCamera returns a new FollowCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.followcamera
func (ba *Babylon) NewFollowCamera(name string, position *Vector3, scene *Scene, opts *NewFollowCameraOpts) *FollowCamera {
	if opts == nil {
		opts = &NewFollowCameraOpts{}
	}

	p := ba.ctx.Get("FollowCamera").New(name, position.JSObject(), scene.JSObject(), opts.LockedTarget.JSObject())
	return FollowCameraFromJSObject(p)
}

// TODO: methods
