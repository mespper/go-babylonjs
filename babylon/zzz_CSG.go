// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// CSG represents a babylon.js CSG.
// Class for building Constructive Solid Geometry
type CSG struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (c *CSG) JSObject() js.Value { return c.p }

// CSG returns a CSG JavaScript class.
func (ba *Babylon) CSG() *CSG {
	p := ba.ctx.Get("CSG")
	return CSGFromJSObject(p, ba.ctx)
}

// CSGFromJSObject returns a wrapped CSG JavaScript class.
func CSGFromJSObject(p js.Value, ctx js.Value) *CSG {
	return &CSG{p: p, ctx: ctx}
}

// CSGArrayToJSArray returns a JavaScript Array for the wrapped array.
func CSGArrayToJSArray(array []*CSG) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// CSGBuildMeshGeometryOpts contains optional parameters for CSG.BuildMeshGeometry.
type CSGBuildMeshGeometryOpts struct {
	Scene         *Scene
	KeepSubMeshes *bool
}

// BuildMeshGeometry calls the BuildMeshGeometry method on the CSG object.
//
// https://doc.babylonjs.com/api/classes/babylon.csg#buildmeshgeometry
func (c *CSG) BuildMeshGeometry(name string, opts *CSGBuildMeshGeometryOpts) *Mesh {
	if opts == nil {
		opts = &CSGBuildMeshGeometryOpts{}
	}

	args := make([]interface{}, 0, 1+2)

	args = append(args, name)

	if opts.Scene == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Scene.JSObject())
	}
	if opts.KeepSubMeshes == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.KeepSubMeshes)
	}

	retVal := c.p.Call("buildMeshGeometry", args...)
	return MeshFromJSObject(retVal, c.ctx)
}

// Clone calls the Clone method on the CSG object.
//
// https://doc.babylonjs.com/api/classes/babylon.csg#clone
func (c *CSG) Clone() *CSG {

	retVal := c.p.Call("clone")
	return CSGFromJSObject(retVal, c.ctx)
}

// CopyTransformAttributes calls the CopyTransformAttributes method on the CSG object.
//
// https://doc.babylonjs.com/api/classes/babylon.csg#copytransformattributes
func (c *CSG) CopyTransformAttributes(csg *CSG) *CSG {

	args := make([]interface{}, 0, 1+0)

	if csg == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, csg.JSObject())
	}

	retVal := c.p.Call("copyTransformAttributes", args...)
	return CSGFromJSObject(retVal, c.ctx)
}

// FromMesh calls the FromMesh method on the CSG object.
//
// https://doc.babylonjs.com/api/classes/babylon.csg#frommesh
func (c *CSG) FromMesh(mesh *Mesh) *CSG {

	args := make([]interface{}, 0, 1+0)

	if mesh == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, mesh.JSObject())
	}

	retVal := c.p.Call("FromMesh", args...)
	return CSGFromJSObject(retVal, c.ctx)
}

// Intersect calls the Intersect method on the CSG object.
//
// https://doc.babylonjs.com/api/classes/babylon.csg#intersect
func (c *CSG) Intersect(csg *CSG) *CSG {

	args := make([]interface{}, 0, 1+0)

	if csg == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, csg.JSObject())
	}

	retVal := c.p.Call("intersect", args...)
	return CSGFromJSObject(retVal, c.ctx)
}

// IntersectInPlace calls the IntersectInPlace method on the CSG object.
//
// https://doc.babylonjs.com/api/classes/babylon.csg#intersectinplace
func (c *CSG) IntersectInPlace(csg *CSG) {

	args := make([]interface{}, 0, 1+0)

	if csg == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, csg.JSObject())
	}

	c.p.Call("intersectInPlace", args...)
}

// Inverse calls the Inverse method on the CSG object.
//
// https://doc.babylonjs.com/api/classes/babylon.csg#inverse
func (c *CSG) Inverse() *CSG {

	retVal := c.p.Call("inverse")
	return CSGFromJSObject(retVal, c.ctx)
}

// InverseInPlace calls the InverseInPlace method on the CSG object.
//
// https://doc.babylonjs.com/api/classes/babylon.csg#inverseinplace
func (c *CSG) InverseInPlace() {

	c.p.Call("inverseInPlace")
}

// Subtract calls the Subtract method on the CSG object.
//
// https://doc.babylonjs.com/api/classes/babylon.csg#subtract
func (c *CSG) Subtract(csg *CSG) *CSG {

	args := make([]interface{}, 0, 1+0)

	if csg == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, csg.JSObject())
	}

	retVal := c.p.Call("subtract", args...)
	return CSGFromJSObject(retVal, c.ctx)
}

// SubtractInPlace calls the SubtractInPlace method on the CSG object.
//
// https://doc.babylonjs.com/api/classes/babylon.csg#subtractinplace
func (c *CSG) SubtractInPlace(csg *CSG) {

	args := make([]interface{}, 0, 1+0)

	if csg == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, csg.JSObject())
	}

	c.p.Call("subtractInPlace", args...)
}

// CSGToMeshOpts contains optional parameters for CSG.ToMesh.
type CSGToMeshOpts struct {
	Material      *Material
	Scene         *Scene
	KeepSubMeshes *bool
}

// ToMesh calls the ToMesh method on the CSG object.
//
// https://doc.babylonjs.com/api/classes/babylon.csg#tomesh
func (c *CSG) ToMesh(name string, opts *CSGToMeshOpts) *Mesh {
	if opts == nil {
		opts = &CSGToMeshOpts{}
	}

	args := make([]interface{}, 0, 1+3)

	args = append(args, name)

	if opts.Material == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Material.JSObject())
	}
	if opts.Scene == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Scene.JSObject())
	}
	if opts.KeepSubMeshes == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.KeepSubMeshes)
	}

	retVal := c.p.Call("toMesh", args...)
	return MeshFromJSObject(retVal, c.ctx)
}

// Union calls the Union method on the CSG object.
//
// https://doc.babylonjs.com/api/classes/babylon.csg#union
func (c *CSG) Union(csg *CSG) *CSG {

	args := make([]interface{}, 0, 1+0)

	if csg == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, csg.JSObject())
	}

	retVal := c.p.Call("union", args...)
	return CSGFromJSObject(retVal, c.ctx)
}

// UnionInPlace calls the UnionInPlace method on the CSG object.
//
// https://doc.babylonjs.com/api/classes/babylon.csg#unioninplace
func (c *CSG) UnionInPlace(csg *CSG) {

	args := make([]interface{}, 0, 1+0)

	if csg == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, csg.JSObject())
	}

	c.p.Call("unionInPlace", args...)
}

// Matrix returns the Matrix property of class CSG.
//
// https://doc.babylonjs.com/api/classes/babylon.csg#matrix
func (c *CSG) Matrix() *Matrix {
	retVal := c.p.Get("matrix")
	return MatrixFromJSObject(retVal, c.ctx)
}

// SetMatrix sets the Matrix property of class CSG.
//
// https://doc.babylonjs.com/api/classes/babylon.csg#matrix
func (c *CSG) SetMatrix(matrix *Matrix) *CSG {
	c.p.Set("matrix", matrix.JSObject())
	return c
}

// Position returns the Position property of class CSG.
//
// https://doc.babylonjs.com/api/classes/babylon.csg#position
func (c *CSG) Position() *Vector3 {
	retVal := c.p.Get("position")
	return Vector3FromJSObject(retVal, c.ctx)
}

// SetPosition sets the Position property of class CSG.
//
// https://doc.babylonjs.com/api/classes/babylon.csg#position
func (c *CSG) SetPosition(position *Vector3) *CSG {
	c.p.Set("position", position.JSObject())
	return c
}

// Rotation returns the Rotation property of class CSG.
//
// https://doc.babylonjs.com/api/classes/babylon.csg#rotation
func (c *CSG) Rotation() *Vector3 {
	retVal := c.p.Get("rotation")
	return Vector3FromJSObject(retVal, c.ctx)
}

// SetRotation sets the Rotation property of class CSG.
//
// https://doc.babylonjs.com/api/classes/babylon.csg#rotation
func (c *CSG) SetRotation(rotation *Vector3) *CSG {
	c.p.Set("rotation", rotation.JSObject())
	return c
}

// RotationQuaternion returns the RotationQuaternion property of class CSG.
//
// https://doc.babylonjs.com/api/classes/babylon.csg#rotationquaternion
func (c *CSG) RotationQuaternion() *Quaternion {
	retVal := c.p.Get("rotationQuaternion")
	return QuaternionFromJSObject(retVal, c.ctx)
}

// SetRotationQuaternion sets the RotationQuaternion property of class CSG.
//
// https://doc.babylonjs.com/api/classes/babylon.csg#rotationquaternion
func (c *CSG) SetRotationQuaternion(rotationQuaternion *Quaternion) *CSG {
	c.p.Set("rotationQuaternion", rotationQuaternion.JSObject())
	return c
}

// Scaling returns the Scaling property of class CSG.
//
// https://doc.babylonjs.com/api/classes/babylon.csg#scaling
func (c *CSG) Scaling() *Vector3 {
	retVal := c.p.Get("scaling")
	return Vector3FromJSObject(retVal, c.ctx)
}

// SetScaling sets the Scaling property of class CSG.
//
// https://doc.babylonjs.com/api/classes/babylon.csg#scaling
func (c *CSG) SetScaling(scaling *Vector3) *CSG {
	c.p.Set("scaling", scaling.JSObject())
	return c
}
