// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// BoneLookController represents a babylon.js BoneLookController.
// Class used to make a bone look toward a point in space
//
// See: http://doc.babylonjs.com/how_to/how_to_use_bones_and_skeletons#bonelookcontroller
type BoneLookController struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (b *BoneLookController) JSObject() js.Value { return b.p }

// BoneLookController returns a BoneLookController JavaScript class.
func (ba *Babylon) BoneLookController() *BoneLookController {
	p := ba.ctx.Get("BoneLookController")
	return BoneLookControllerFromJSObject(p, ba.ctx)
}

// BoneLookControllerFromJSObject returns a wrapped BoneLookController JavaScript class.
func BoneLookControllerFromJSObject(p js.Value, ctx js.Value) *BoneLookController {
	return &BoneLookController{p: p, ctx: ctx}
}

// BoneLookControllerArrayToJSArray returns a JavaScript Array for the wrapped array.
func BoneLookControllerArrayToJSArray(array []*BoneLookController) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewBoneLookControllerOpts contains optional parameters for NewBoneLookController.
type NewBoneLookControllerOpts struct {
	Options map[string]interface{}
}

// NewBoneLookController returns a new BoneLookController object.
//
// https://doc.babylonjs.com/api/classes/babylon.bonelookcontroller
func (ba *Babylon) NewBoneLookController(mesh *AbstractMesh, bone *Bone, target *Vector3, opts *NewBoneLookControllerOpts) *BoneLookController {
	if opts == nil {
		opts = &NewBoneLookControllerOpts{}
	}

	args := make([]interface{}, 0, 3+1)

	args = append(args, mesh.JSObject())
	args = append(args, bone.JSObject())
	args = append(args, target.JSObject())

	if opts.Options == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Options)
	}

	p := ba.ctx.Get("BoneLookController").New(args...)
	return BoneLookControllerFromJSObject(p, ba.ctx)
}

// Update calls the Update method on the BoneLookController object.
//
// https://doc.babylonjs.com/api/classes/babylon.bonelookcontroller#update
func (b *BoneLookController) Update() {

	b.p.Call("update")
}

// AdjustPitch returns the AdjustPitch property of class BoneLookController.
//
// https://doc.babylonjs.com/api/classes/babylon.bonelookcontroller#adjustpitch
func (b *BoneLookController) AdjustPitch() float64 {
	retVal := b.p.Get("adjustPitch")
	return retVal.Float()
}

// SetAdjustPitch sets the AdjustPitch property of class BoneLookController.
//
// https://doc.babylonjs.com/api/classes/babylon.bonelookcontroller#adjustpitch
func (b *BoneLookController) SetAdjustPitch(adjustPitch float64) *BoneLookController {
	b.p.Set("adjustPitch", adjustPitch)
	return b
}

// AdjustRoll returns the AdjustRoll property of class BoneLookController.
//
// https://doc.babylonjs.com/api/classes/babylon.bonelookcontroller#adjustroll
func (b *BoneLookController) AdjustRoll() float64 {
	retVal := b.p.Get("adjustRoll")
	return retVal.Float()
}

// SetAdjustRoll sets the AdjustRoll property of class BoneLookController.
//
// https://doc.babylonjs.com/api/classes/babylon.bonelookcontroller#adjustroll
func (b *BoneLookController) SetAdjustRoll(adjustRoll float64) *BoneLookController {
	b.p.Set("adjustRoll", adjustRoll)
	return b
}

// AdjustYaw returns the AdjustYaw property of class BoneLookController.
//
// https://doc.babylonjs.com/api/classes/babylon.bonelookcontroller#adjustyaw
func (b *BoneLookController) AdjustYaw() float64 {
	retVal := b.p.Get("adjustYaw")
	return retVal.Float()
}

// SetAdjustYaw sets the AdjustYaw property of class BoneLookController.
//
// https://doc.babylonjs.com/api/classes/babylon.bonelookcontroller#adjustyaw
func (b *BoneLookController) SetAdjustYaw(adjustYaw float64) *BoneLookController {
	b.p.Set("adjustYaw", adjustYaw)
	return b
}

// Bone returns the Bone property of class BoneLookController.
//
// https://doc.babylonjs.com/api/classes/babylon.bonelookcontroller#bone
func (b *BoneLookController) Bone() *Bone {
	retVal := b.p.Get("bone")
	return BoneFromJSObject(retVal, b.ctx)
}

// SetBone sets the Bone property of class BoneLookController.
//
// https://doc.babylonjs.com/api/classes/babylon.bonelookcontroller#bone
func (b *BoneLookController) SetBone(bone *Bone) *BoneLookController {
	b.p.Set("bone", bone.JSObject())
	return b
}

// MaxPitch returns the MaxPitch property of class BoneLookController.
//
// https://doc.babylonjs.com/api/classes/babylon.bonelookcontroller#maxpitch
func (b *BoneLookController) MaxPitch() float64 {
	retVal := b.p.Get("maxPitch")
	return retVal.Float()
}

// SetMaxPitch sets the MaxPitch property of class BoneLookController.
//
// https://doc.babylonjs.com/api/classes/babylon.bonelookcontroller#maxpitch
func (b *BoneLookController) SetMaxPitch(maxPitch float64) *BoneLookController {
	b.p.Set("maxPitch", maxPitch)
	return b
}

// MaxYaw returns the MaxYaw property of class BoneLookController.
//
// https://doc.babylonjs.com/api/classes/babylon.bonelookcontroller#maxyaw
func (b *BoneLookController) MaxYaw() float64 {
	retVal := b.p.Get("maxYaw")
	return retVal.Float()
}

// SetMaxYaw sets the MaxYaw property of class BoneLookController.
//
// https://doc.babylonjs.com/api/classes/babylon.bonelookcontroller#maxyaw
func (b *BoneLookController) SetMaxYaw(maxYaw float64) *BoneLookController {
	b.p.Set("maxYaw", maxYaw)
	return b
}

// Mesh returns the Mesh property of class BoneLookController.
//
// https://doc.babylonjs.com/api/classes/babylon.bonelookcontroller#mesh
func (b *BoneLookController) Mesh() *AbstractMesh {
	retVal := b.p.Get("mesh")
	return AbstractMeshFromJSObject(retVal, b.ctx)
}

// SetMesh sets the Mesh property of class BoneLookController.
//
// https://doc.babylonjs.com/api/classes/babylon.bonelookcontroller#mesh
func (b *BoneLookController) SetMesh(mesh *AbstractMesh) *BoneLookController {
	b.p.Set("mesh", mesh.JSObject())
	return b
}

// MinPitch returns the MinPitch property of class BoneLookController.
//
// https://doc.babylonjs.com/api/classes/babylon.bonelookcontroller#minpitch
func (b *BoneLookController) MinPitch() float64 {
	retVal := b.p.Get("minPitch")
	return retVal.Float()
}

// SetMinPitch sets the MinPitch property of class BoneLookController.
//
// https://doc.babylonjs.com/api/classes/babylon.bonelookcontroller#minpitch
func (b *BoneLookController) SetMinPitch(minPitch float64) *BoneLookController {
	b.p.Set("minPitch", minPitch)
	return b
}

// MinYaw returns the MinYaw property of class BoneLookController.
//
// https://doc.babylonjs.com/api/classes/babylon.bonelookcontroller#minyaw
func (b *BoneLookController) MinYaw() float64 {
	retVal := b.p.Get("minYaw")
	return retVal.Float()
}

// SetMinYaw sets the MinYaw property of class BoneLookController.
//
// https://doc.babylonjs.com/api/classes/babylon.bonelookcontroller#minyaw
func (b *BoneLookController) SetMinYaw(minYaw float64) *BoneLookController {
	b.p.Set("minYaw", minYaw)
	return b
}

// SlerpAmount returns the SlerpAmount property of class BoneLookController.
//
// https://doc.babylonjs.com/api/classes/babylon.bonelookcontroller#slerpamount
func (b *BoneLookController) SlerpAmount() float64 {
	retVal := b.p.Get("slerpAmount")
	return retVal.Float()
}

// SetSlerpAmount sets the SlerpAmount property of class BoneLookController.
//
// https://doc.babylonjs.com/api/classes/babylon.bonelookcontroller#slerpamount
func (b *BoneLookController) SetSlerpAmount(slerpAmount float64) *BoneLookController {
	b.p.Set("slerpAmount", slerpAmount)
	return b
}

// Target returns the Target property of class BoneLookController.
//
// https://doc.babylonjs.com/api/classes/babylon.bonelookcontroller#target
func (b *BoneLookController) Target() *Vector3 {
	retVal := b.p.Get("target")
	return Vector3FromJSObject(retVal, b.ctx)
}

// SetTarget sets the Target property of class BoneLookController.
//
// https://doc.babylonjs.com/api/classes/babylon.bonelookcontroller#target
func (b *BoneLookController) SetTarget(target *Vector3) *BoneLookController {
	b.p.Set("target", target.JSObject())
	return b
}

// UpAxis returns the UpAxis property of class BoneLookController.
//
// https://doc.babylonjs.com/api/classes/babylon.bonelookcontroller#upaxis
func (b *BoneLookController) UpAxis() *Vector3 {
	retVal := b.p.Get("upAxis")
	return Vector3FromJSObject(retVal, b.ctx)
}

// SetUpAxis sets the UpAxis property of class BoneLookController.
//
// https://doc.babylonjs.com/api/classes/babylon.bonelookcontroller#upaxis
func (b *BoneLookController) SetUpAxis(upAxis *Vector3) *BoneLookController {
	b.p.Set("upAxis", upAxis.JSObject())
	return b
}

// UpAxisSpace returns the UpAxisSpace property of class BoneLookController.
//
// https://doc.babylonjs.com/api/classes/babylon.bonelookcontroller#upaxisspace
func (b *BoneLookController) UpAxisSpace() js.Value {
	retVal := b.p.Get("upAxisSpace")
	return retVal
}

// SetUpAxisSpace sets the UpAxisSpace property of class BoneLookController.
//
// https://doc.babylonjs.com/api/classes/babylon.bonelookcontroller#upaxisspace
func (b *BoneLookController) SetUpAxisSpace(upAxisSpace js.Value) *BoneLookController {
	b.p.Set("upAxisSpace", upAxisSpace)
	return b
}
