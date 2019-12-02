// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// BoneIKController represents a babylon.js BoneIKController.
// Class used to apply inverse kinematics to bones
//
// See: http://doc.babylonjs.com/how_to/how_to_use_bones_and_skeletons#boneikcontroller
type BoneIKController struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (b *BoneIKController) JSObject() js.Value { return b.p }

// BoneIKController returns a BoneIKController JavaScript class.
func (ba *Babylon) BoneIKController() *BoneIKController {
	p := ba.ctx.Get("BoneIKController")
	return BoneIKControllerFromJSObject(p, ba.ctx)
}

// BoneIKControllerFromJSObject returns a wrapped BoneIKController JavaScript class.
func BoneIKControllerFromJSObject(p js.Value, ctx js.Value) *BoneIKController {
	return &BoneIKController{p: p, ctx: ctx}
}

// BoneIKControllerArrayToJSArray returns a JavaScript Array for the wrapped array.
func BoneIKControllerArrayToJSArray(array []*BoneIKController) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewBoneIKControllerOpts contains optional parameters for NewBoneIKController.
type NewBoneIKControllerOpts struct {
	Options map[string]interface{}
}

// NewBoneIKController returns a new BoneIKController object.
//
// https://doc.babylonjs.com/api/classes/babylon.boneikcontroller
func (ba *Babylon) NewBoneIKController(mesh *AbstractMesh, bone *Bone, opts *NewBoneIKControllerOpts) *BoneIKController {
	if opts == nil {
		opts = &NewBoneIKControllerOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, mesh.JSObject())
	args = append(args, bone.JSObject())

	if opts.Options == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Options)
	}

	p := ba.ctx.Get("BoneIKController").New(args...)
	return BoneIKControllerFromJSObject(p, ba.ctx)
}

// Update calls the Update method on the BoneIKController object.
//
// https://doc.babylonjs.com/api/classes/babylon.boneikcontroller#update
func (b *BoneIKController) Update() {

	b.p.Call("update")
}

// MaxAngle returns the MaxAngle property of class BoneIKController.
//
// https://doc.babylonjs.com/api/classes/babylon.boneikcontroller#maxangle
func (b *BoneIKController) MaxAngle() float64 {
	retVal := b.p.Get("maxAngle")
	return retVal.Float()
}

// SetMaxAngle sets the MaxAngle property of class BoneIKController.
//
// https://doc.babylonjs.com/api/classes/babylon.boneikcontroller#maxangle
func (b *BoneIKController) SetMaxAngle(maxAngle float64) *BoneIKController {
	b.p.Set("maxAngle", maxAngle)
	return b
}

// Mesh returns the Mesh property of class BoneIKController.
//
// https://doc.babylonjs.com/api/classes/babylon.boneikcontroller#mesh
func (b *BoneIKController) Mesh() *AbstractMesh {
	retVal := b.p.Get("mesh")
	return AbstractMeshFromJSObject(retVal, b.ctx)
}

// SetMesh sets the Mesh property of class BoneIKController.
//
// https://doc.babylonjs.com/api/classes/babylon.boneikcontroller#mesh
func (b *BoneIKController) SetMesh(mesh *AbstractMesh) *BoneIKController {
	b.p.Set("mesh", mesh.JSObject())
	return b
}

// PoleAngle returns the PoleAngle property of class BoneIKController.
//
// https://doc.babylonjs.com/api/classes/babylon.boneikcontroller#poleangle
func (b *BoneIKController) PoleAngle() float64 {
	retVal := b.p.Get("poleAngle")
	return retVal.Float()
}

// SetPoleAngle sets the PoleAngle property of class BoneIKController.
//
// https://doc.babylonjs.com/api/classes/babylon.boneikcontroller#poleangle
func (b *BoneIKController) SetPoleAngle(poleAngle float64) *BoneIKController {
	b.p.Set("poleAngle", poleAngle)
	return b
}

// PoleTargetBone returns the PoleTargetBone property of class BoneIKController.
//
// https://doc.babylonjs.com/api/classes/babylon.boneikcontroller#poletargetbone
func (b *BoneIKController) PoleTargetBone() *Bone {
	retVal := b.p.Get("poleTargetBone")
	return BoneFromJSObject(retVal, b.ctx)
}

// SetPoleTargetBone sets the PoleTargetBone property of class BoneIKController.
//
// https://doc.babylonjs.com/api/classes/babylon.boneikcontroller#poletargetbone
func (b *BoneIKController) SetPoleTargetBone(poleTargetBone *Bone) *BoneIKController {
	b.p.Set("poleTargetBone", poleTargetBone.JSObject())
	return b
}

// PoleTargetLocalOffset returns the PoleTargetLocalOffset property of class BoneIKController.
//
// https://doc.babylonjs.com/api/classes/babylon.boneikcontroller#poletargetlocaloffset
func (b *BoneIKController) PoleTargetLocalOffset() *Vector3 {
	retVal := b.p.Get("poleTargetLocalOffset")
	return Vector3FromJSObject(retVal, b.ctx)
}

// SetPoleTargetLocalOffset sets the PoleTargetLocalOffset property of class BoneIKController.
//
// https://doc.babylonjs.com/api/classes/babylon.boneikcontroller#poletargetlocaloffset
func (b *BoneIKController) SetPoleTargetLocalOffset(poleTargetLocalOffset *Vector3) *BoneIKController {
	b.p.Set("poleTargetLocalOffset", poleTargetLocalOffset.JSObject())
	return b
}

// PoleTargetMesh returns the PoleTargetMesh property of class BoneIKController.
//
// https://doc.babylonjs.com/api/classes/babylon.boneikcontroller#poletargetmesh
func (b *BoneIKController) PoleTargetMesh() *AbstractMesh {
	retVal := b.p.Get("poleTargetMesh")
	return AbstractMeshFromJSObject(retVal, b.ctx)
}

// SetPoleTargetMesh sets the PoleTargetMesh property of class BoneIKController.
//
// https://doc.babylonjs.com/api/classes/babylon.boneikcontroller#poletargetmesh
func (b *BoneIKController) SetPoleTargetMesh(poleTargetMesh *AbstractMesh) *BoneIKController {
	b.p.Set("poleTargetMesh", poleTargetMesh.JSObject())
	return b
}

// PoleTargetPosition returns the PoleTargetPosition property of class BoneIKController.
//
// https://doc.babylonjs.com/api/classes/babylon.boneikcontroller#poletargetposition
func (b *BoneIKController) PoleTargetPosition() *Vector3 {
	retVal := b.p.Get("poleTargetPosition")
	return Vector3FromJSObject(retVal, b.ctx)
}

// SetPoleTargetPosition sets the PoleTargetPosition property of class BoneIKController.
//
// https://doc.babylonjs.com/api/classes/babylon.boneikcontroller#poletargetposition
func (b *BoneIKController) SetPoleTargetPosition(poleTargetPosition *Vector3) *BoneIKController {
	b.p.Set("poleTargetPosition", poleTargetPosition.JSObject())
	return b
}

// SlerpAmount returns the SlerpAmount property of class BoneIKController.
//
// https://doc.babylonjs.com/api/classes/babylon.boneikcontroller#slerpamount
func (b *BoneIKController) SlerpAmount() float64 {
	retVal := b.p.Get("slerpAmount")
	return retVal.Float()
}

// SetSlerpAmount sets the SlerpAmount property of class BoneIKController.
//
// https://doc.babylonjs.com/api/classes/babylon.boneikcontroller#slerpamount
func (b *BoneIKController) SetSlerpAmount(slerpAmount float64) *BoneIKController {
	b.p.Set("slerpAmount", slerpAmount)
	return b
}

// TargetMesh returns the TargetMesh property of class BoneIKController.
//
// https://doc.babylonjs.com/api/classes/babylon.boneikcontroller#targetmesh
func (b *BoneIKController) TargetMesh() *AbstractMesh {
	retVal := b.p.Get("targetMesh")
	return AbstractMeshFromJSObject(retVal, b.ctx)
}

// SetTargetMesh sets the TargetMesh property of class BoneIKController.
//
// https://doc.babylonjs.com/api/classes/babylon.boneikcontroller#targetmesh
func (b *BoneIKController) SetTargetMesh(targetMesh *AbstractMesh) *BoneIKController {
	b.p.Set("targetMesh", targetMesh.JSObject())
	return b
}

// TargetPosition returns the TargetPosition property of class BoneIKController.
//
// https://doc.babylonjs.com/api/classes/babylon.boneikcontroller#targetposition
func (b *BoneIKController) TargetPosition() *Vector3 {
	retVal := b.p.Get("targetPosition")
	return Vector3FromJSObject(retVal, b.ctx)
}

// SetTargetPosition sets the TargetPosition property of class BoneIKController.
//
// https://doc.babylonjs.com/api/classes/babylon.boneikcontroller#targetposition
func (b *BoneIKController) SetTargetPosition(targetPosition *Vector3) *BoneIKController {
	b.p.Set("targetPosition", targetPosition.JSObject())
	return b
}
