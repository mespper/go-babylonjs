// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// IPhysicsEnabledObject represents a babylon.js IPhysicsEnabledObject.
// Interface for a physics-enabled object
//
// See: https://doc.babylonjs.com/how_to/using_the_physics_engine
type IPhysicsEnabledObject struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (i *IPhysicsEnabledObject) JSObject() js.Value { return i.p }

// IPhysicsEnabledObject returns a IPhysicsEnabledObject JavaScript class.
func (ba *Babylon) IPhysicsEnabledObject() *IPhysicsEnabledObject {
	p := ba.ctx.Get("IPhysicsEnabledObject")
	return IPhysicsEnabledObjectFromJSObject(p, ba.ctx)
}

// IPhysicsEnabledObjectFromJSObject returns a wrapped IPhysicsEnabledObject JavaScript class.
func IPhysicsEnabledObjectFromJSObject(p js.Value, ctx js.Value) *IPhysicsEnabledObject {
	return &IPhysicsEnabledObject{p: p, ctx: ctx}
}

// IPhysicsEnabledObjectArrayToJSArray returns a JavaScript Array for the wrapped array.
func IPhysicsEnabledObjectArrayToJSArray(array []*IPhysicsEnabledObject) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// ComputeWorldMatrix calls the ComputeWorldMatrix method on the IPhysicsEnabledObject object.
//
// https://doc.babylonjs.com/api/classes/babylon.iphysicsenabledobject#computeworldmatrix
func (i *IPhysicsEnabledObject) ComputeWorldMatrix(force bool) *Matrix {

	args := make([]interface{}, 0, 1+0)

	args = append(args, force)

	retVal := i.p.Call("computeWorldMatrix", args...)
	return MatrixFromJSObject(retVal, i.ctx)
}

// GetAbsolutePivotPoint calls the GetAbsolutePivotPoint method on the IPhysicsEnabledObject object.
//
// https://doc.babylonjs.com/api/classes/babylon.iphysicsenabledobject#getabsolutepivotpoint
func (i *IPhysicsEnabledObject) GetAbsolutePivotPoint() *Vector3 {

	retVal := i.p.Call("getAbsolutePivotPoint")
	return Vector3FromJSObject(retVal, i.ctx)
}

// GetAbsolutePosition calls the GetAbsolutePosition method on the IPhysicsEnabledObject object.
//
// https://doc.babylonjs.com/api/classes/babylon.iphysicsenabledobject#getabsoluteposition
func (i *IPhysicsEnabledObject) GetAbsolutePosition() *Vector3 {

	retVal := i.p.Call("getAbsolutePosition")
	return Vector3FromJSObject(retVal, i.ctx)
}

// GetBoundingInfo calls the GetBoundingInfo method on the IPhysicsEnabledObject object.
//
// https://doc.babylonjs.com/api/classes/babylon.iphysicsenabledobject#getboundinginfo
func (i *IPhysicsEnabledObject) GetBoundingInfo() *BoundingInfo {

	retVal := i.p.Call("getBoundingInfo")
	return BoundingInfoFromJSObject(retVal, i.ctx)
}

// IPhysicsEnabledObjectGetChildMeshesOpts contains optional parameters for IPhysicsEnabledObject.GetChildMeshes.
type IPhysicsEnabledObjectGetChildMeshesOpts struct {
	DirectDescendantsOnly *bool
}

// GetChildMeshes calls the GetChildMeshes method on the IPhysicsEnabledObject object.
//
// https://doc.babylonjs.com/api/classes/babylon.iphysicsenabledobject#getchildmeshes
func (i *IPhysicsEnabledObject) GetChildMeshes(opts *IPhysicsEnabledObjectGetChildMeshesOpts) []*AbstractMesh {
	if opts == nil {
		opts = &IPhysicsEnabledObjectGetChildMeshesOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.DirectDescendantsOnly == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.DirectDescendantsOnly)
	}

	retVal := i.p.Call("getChildMeshes", args...)
	return retVal
}

// GetClassName calls the GetClassName method on the IPhysicsEnabledObject object.
//
// https://doc.babylonjs.com/api/classes/babylon.iphysicsenabledobject#getclassname
func (i *IPhysicsEnabledObject) GetClassName() string {

	retVal := i.p.Call("getClassName")
	return retVal.String()
}

// GetIndices calls the GetIndices method on the IPhysicsEnabledObject object.
//
// https://doc.babylonjs.com/api/classes/babylon.iphysicsenabledobject#getindices
func (i *IPhysicsEnabledObject) GetIndices() js.Value {

	retVal := i.p.Call("getIndices")
	return retVal
}

// GetScene calls the GetScene method on the IPhysicsEnabledObject object.
//
// https://doc.babylonjs.com/api/classes/babylon.iphysicsenabledobject#getscene
func (i *IPhysicsEnabledObject) GetScene() *Scene {

	retVal := i.p.Call("getScene")
	return SceneFromJSObject(retVal, i.ctx)
}

// GetVerticesData calls the GetVerticesData method on the IPhysicsEnabledObject object.
//
// https://doc.babylonjs.com/api/classes/babylon.iphysicsenabledobject#getverticesdata
func (i *IPhysicsEnabledObject) GetVerticesData(kind string) []*float64 {

	args := make([]interface{}, 0, 1+0)

	args = append(args, kind)

	retVal := i.p.Call("getVerticesData", args...)
	return retVal
}

// GetWorldMatrix calls the GetWorldMatrix method on the IPhysicsEnabledObject object.
//
// https://doc.babylonjs.com/api/classes/babylon.iphysicsenabledobject#getworldmatrix
func (i *IPhysicsEnabledObject) GetWorldMatrix() *Matrix {

	retVal := i.p.Call("getWorldMatrix")
	return MatrixFromJSObject(retVal, i.ctx)
}

// IPhysicsEnabledObjectRotateOpts contains optional parameters for IPhysicsEnabledObject.Rotate.
type IPhysicsEnabledObjectRotateOpts struct {
	Space js.Value
}

// Rotate calls the Rotate method on the IPhysicsEnabledObject object.
//
// https://doc.babylonjs.com/api/classes/babylon.iphysicsenabledobject#rotate
func (i *IPhysicsEnabledObject) Rotate(axis *Vector3, amount float64, opts *IPhysicsEnabledObjectRotateOpts) *TransformNode {
	if opts == nil {
		opts = &IPhysicsEnabledObjectRotateOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, axis.JSObject())
	args = append(args, amount)

	if opts.Space == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Space)
	}

	retVal := i.p.Call("rotate", args...)
	return TransformNodeFromJSObject(retVal, i.ctx)
}

// SetAbsolutePosition calls the SetAbsolutePosition method on the IPhysicsEnabledObject object.
//
// https://doc.babylonjs.com/api/classes/babylon.iphysicsenabledobject#setabsoluteposition
func (i *IPhysicsEnabledObject) SetAbsolutePosition(absolutePosition *Vector3) *TransformNode {

	args := make([]interface{}, 0, 1+0)

	args = append(args, absolutePosition.JSObject())

	retVal := i.p.Call("setAbsolutePosition", args...)
	return TransformNodeFromJSObject(retVal, i.ctx)
}

// IPhysicsEnabledObjectTranslateOpts contains optional parameters for IPhysicsEnabledObject.Translate.
type IPhysicsEnabledObjectTranslateOpts struct {
	Space js.Value
}

// Translate calls the Translate method on the IPhysicsEnabledObject object.
//
// https://doc.babylonjs.com/api/classes/babylon.iphysicsenabledobject#translate
func (i *IPhysicsEnabledObject) Translate(axis *Vector3, distance float64, opts *IPhysicsEnabledObjectTranslateOpts) *TransformNode {
	if opts == nil {
		opts = &IPhysicsEnabledObjectTranslateOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, axis.JSObject())
	args = append(args, distance)

	if opts.Space == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Space)
	}

	retVal := i.p.Call("translate", args...)
	return TransformNodeFromJSObject(retVal, i.ctx)
}

/*

// Parent returns the Parent property of class IPhysicsEnabledObject.
//
// https://doc.babylonjs.com/api/classes/babylon.iphysicsenabledobject#parent
func (i *IPhysicsEnabledObject) Parent(parent interface{}) *IPhysicsEnabledObject {
	p := ba.ctx.Get("IPhysicsEnabledObject").New(parent)
	return IPhysicsEnabledObjectFromJSObject(p, ba.ctx)
}

// SetParent sets the Parent property of class IPhysicsEnabledObject.
//
// https://doc.babylonjs.com/api/classes/babylon.iphysicsenabledobject#parent
func (i *IPhysicsEnabledObject) SetParent(parent interface{}) *IPhysicsEnabledObject {
	p := ba.ctx.Get("IPhysicsEnabledObject").New(parent)
	return IPhysicsEnabledObjectFromJSObject(p, ba.ctx)
}

// Position returns the Position property of class IPhysicsEnabledObject.
//
// https://doc.babylonjs.com/api/classes/babylon.iphysicsenabledobject#position
func (i *IPhysicsEnabledObject) Position(position *Vector3) *IPhysicsEnabledObject {
	p := ba.ctx.Get("IPhysicsEnabledObject").New(position.JSObject())
	return IPhysicsEnabledObjectFromJSObject(p, ba.ctx)
}

// SetPosition sets the Position property of class IPhysicsEnabledObject.
//
// https://doc.babylonjs.com/api/classes/babylon.iphysicsenabledobject#position
func (i *IPhysicsEnabledObject) SetPosition(position *Vector3) *IPhysicsEnabledObject {
	p := ba.ctx.Get("IPhysicsEnabledObject").New(position.JSObject())
	return IPhysicsEnabledObjectFromJSObject(p, ba.ctx)
}

// Rotation returns the Rotation property of class IPhysicsEnabledObject.
//
// https://doc.babylonjs.com/api/classes/babylon.iphysicsenabledobject#rotation
func (i *IPhysicsEnabledObject) Rotation(rotation *Vector3) *IPhysicsEnabledObject {
	p := ba.ctx.Get("IPhysicsEnabledObject").New(rotation.JSObject())
	return IPhysicsEnabledObjectFromJSObject(p, ba.ctx)
}

// SetRotation sets the Rotation property of class IPhysicsEnabledObject.
//
// https://doc.babylonjs.com/api/classes/babylon.iphysicsenabledobject#rotation
func (i *IPhysicsEnabledObject) SetRotation(rotation *Vector3) *IPhysicsEnabledObject {
	p := ba.ctx.Get("IPhysicsEnabledObject").New(rotation.JSObject())
	return IPhysicsEnabledObjectFromJSObject(p, ba.ctx)
}

// RotationQuaternion returns the RotationQuaternion property of class IPhysicsEnabledObject.
//
// https://doc.babylonjs.com/api/classes/babylon.iphysicsenabledobject#rotationquaternion
func (i *IPhysicsEnabledObject) RotationQuaternion(rotationQuaternion *Quaternion) *IPhysicsEnabledObject {
	p := ba.ctx.Get("IPhysicsEnabledObject").New(rotationQuaternion.JSObject())
	return IPhysicsEnabledObjectFromJSObject(p, ba.ctx)
}

// SetRotationQuaternion sets the RotationQuaternion property of class IPhysicsEnabledObject.
//
// https://doc.babylonjs.com/api/classes/babylon.iphysicsenabledobject#rotationquaternion
func (i *IPhysicsEnabledObject) SetRotationQuaternion(rotationQuaternion *Quaternion) *IPhysicsEnabledObject {
	p := ba.ctx.Get("IPhysicsEnabledObject").New(rotationQuaternion.JSObject())
	return IPhysicsEnabledObjectFromJSObject(p, ba.ctx)
}

// Scaling returns the Scaling property of class IPhysicsEnabledObject.
//
// https://doc.babylonjs.com/api/classes/babylon.iphysicsenabledobject#scaling
func (i *IPhysicsEnabledObject) Scaling(scaling *Vector3) *IPhysicsEnabledObject {
	p := ba.ctx.Get("IPhysicsEnabledObject").New(scaling.JSObject())
	return IPhysicsEnabledObjectFromJSObject(p, ba.ctx)
}

// SetScaling sets the Scaling property of class IPhysicsEnabledObject.
//
// https://doc.babylonjs.com/api/classes/babylon.iphysicsenabledobject#scaling
func (i *IPhysicsEnabledObject) SetScaling(scaling *Vector3) *IPhysicsEnabledObject {
	p := ba.ctx.Get("IPhysicsEnabledObject").New(scaling.JSObject())
	return IPhysicsEnabledObjectFromJSObject(p, ba.ctx)
}

*/
