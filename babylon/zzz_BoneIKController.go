// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// BoneIKController represents a babylon.js BoneIKController.
// Class used to apply inverse kinematics to bones
//
// See: http://doc.babylonjs.com/how_to/how_to_use_bones_and_skeletons#boneikcontroller
type BoneIKController struct{ p js.Value }

// JSObject returns the underlying js.Value.
func (b *BoneIKController) JSObject() js.Value { return b.p }

// BoneIKController returns a BoneIKController JavaScript class.
func (ba *Babylon) BoneIKController() *BoneIKController {
	p := ba.ctx.Get("BoneIKController")
	return BoneIKControllerFromJSObject(p)
}

// BoneIKControllerFromJSObject returns a wrapped BoneIKController JavaScript class.
func BoneIKControllerFromJSObject(p js.Value) *BoneIKController {
	return &BoneIKController{p: p}
}

// NewBoneIKControllerOpts contains optional parameters for NewBoneIKController.
type NewBoneIKControllerOpts struct {
	Options *JSValue
}

// NewBoneIKController returns a new BoneIKController object.
//
// https://doc.babylonjs.com/api/classes/babylon.boneikcontroller
func (ba *Babylon) NewBoneIKController(mesh *AbstractMesh, bone *Bone, opts *NewBoneIKControllerOpts) *BoneIKController {
	if opts == nil {
		opts = &NewBoneIKControllerOpts{}
	}

	p := ba.ctx.Get("BoneIKController").New(mesh.JSObject(), bone.JSObject(), opts.Options.JSObject())
	return BoneIKControllerFromJSObject(p)
}

// TODO: methods
