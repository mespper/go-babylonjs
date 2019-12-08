// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// Octree represents a babylon.js Octree.
// Octrees are a really powerful data structure that can quickly select entities based on space coordinates.
//
// See: https://doc.babylonjs.com/how_to/optimizing_your_scene_with_octrees
type Octree struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (o *Octree) JSObject() js.Value { return o.p }

// Octree returns a Octree JavaScript class.
func (ba *Babylon) Octree() *Octree {
	p := ba.ctx.Get("Octree")
	return OctreeFromJSObject(p, ba.ctx)
}

// OctreeFromJSObject returns a wrapped Octree JavaScript class.
func OctreeFromJSObject(p js.Value, ctx js.Value) *Octree {
	return &Octree{p: p, ctx: ctx}
}

// OctreeArrayToJSArray returns a JavaScript Array for the wrapped array.
func OctreeArrayToJSArray(array []*Octree) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewOctreeOpts contains optional parameters for NewOctree.
type NewOctreeOpts struct {
	MaxBlockCapacity *float64
	MaxDepth         *float64
}

// NewOctree returns a new Octree object.
//
// https://doc.babylonjs.com/api/classes/babylon.octree
func (ba *Babylon) NewOctree(creationFunc JSFunc, opts *NewOctreeOpts) *Octree {
	if opts == nil {
		opts = &NewOctreeOpts{}
	}

	args := make([]interface{}, 0, 1+2)

	args = append(args, js.FuncOf(creationFunc))

	if opts.MaxBlockCapacity == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.MaxBlockCapacity)
	}
	if opts.MaxDepth == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.MaxDepth)
	}

	p := ba.ctx.Get("Octree").New(args...)
	return OctreeFromJSObject(p, ba.ctx)
}

// AddMesh calls the AddMesh method on the Octree object.
//
// https://doc.babylonjs.com/api/classes/babylon.octree#addmesh
func (o *Octree) AddMesh(entry *T) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, entry.JSObject())

	o.p.Call("addMesh", args...)
}

// OctreeIntersectsOpts contains optional parameters for Octree.Intersects.
type OctreeIntersectsOpts struct {
	AllowDuplicate *bool
}

// Intersects calls the Intersects method on the Octree object.
//
// https://doc.babylonjs.com/api/classes/babylon.octree#intersects
func (o *Octree) Intersects(sphereCenter *Vector3, sphereRadius float64, opts *OctreeIntersectsOpts) *SmartArray {
	if opts == nil {
		opts = &OctreeIntersectsOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, sphereCenter.JSObject())
	args = append(args, sphereRadius)

	if opts.AllowDuplicate == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.AllowDuplicate)
	}

	retVal := o.p.Call("intersects", args...)
	return SmartArrayFromJSObject(retVal, o.ctx)
}

// IntersectsRay calls the IntersectsRay method on the Octree object.
//
// https://doc.babylonjs.com/api/classes/babylon.octree#intersectsray
func (o *Octree) IntersectsRay(ray *Ray) *SmartArray {

	args := make([]interface{}, 0, 1+0)

	args = append(args, ray.JSObject())

	retVal := o.p.Call("intersectsRay", args...)
	return SmartArrayFromJSObject(retVal, o.ctx)
}

// RemoveMesh calls the RemoveMesh method on the Octree object.
//
// https://doc.babylonjs.com/api/classes/babylon.octree#removemesh
func (o *Octree) RemoveMesh(entry *T) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, entry.JSObject())

	o.p.Call("removeMesh", args...)
}

// OctreeSelectOpts contains optional parameters for Octree.Select.
type OctreeSelectOpts struct {
	AllowDuplicate *bool
}

// Select calls the Select method on the Octree object.
//
// https://doc.babylonjs.com/api/classes/babylon.octree#select
func (o *Octree) Select(frustumPlanes []*Plane, opts *OctreeSelectOpts) *SmartArray {
	if opts == nil {
		opts = &OctreeSelectOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, PlaneArrayToJSArray(frustumPlanes))

	if opts.AllowDuplicate == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.AllowDuplicate)
	}

	retVal := o.p.Call("select", args...)
	return SmartArrayFromJSObject(retVal, o.ctx)
}

// Update calls the Update method on the Octree object.
//
// https://doc.babylonjs.com/api/classes/babylon.octree#update
func (o *Octree) Update(worldMin *Vector3, worldMax *Vector3, entries []*T) {

	args := make([]interface{}, 0, 3+0)

	args = append(args, worldMin.JSObject())
	args = append(args, worldMax.JSObject())
	args = append(args, TArrayToJSArray(entries))

	o.p.Call("update", args...)
}

// Blocks returns the Blocks property of class Octree.
//
// https://doc.babylonjs.com/api/classes/babylon.octree#blocks
func (o *Octree) Blocks() []*OctreeBlock {
	retVal := o.p.Get("blocks")
	result := []*OctreeBlock{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, OctreeBlockFromJSObject(retVal.Index(ri), o.ctx))
	}
	return result
}

// SetBlocks sets the Blocks property of class Octree.
//
// https://doc.babylonjs.com/api/classes/babylon.octree#blocks
func (o *Octree) SetBlocks(blocks []*OctreeBlock) *Octree {
	o.p.Set("blocks", blocks)
	return o
}

// CreationFuncForMeshes returns the CreationFuncForMeshes property of class Octree.
//
// https://doc.babylonjs.com/api/classes/babylon.octree#creationfuncformeshes
func (o *Octree) CreationFuncForMeshes() js.Value {
	retVal := o.p.Get("CreationFuncForMeshes")
	return retVal
}

// SetCreationFuncForMeshes sets the CreationFuncForMeshes property of class Octree.
//
// https://doc.babylonjs.com/api/classes/babylon.octree#creationfuncformeshes
func (o *Octree) SetCreationFuncForMeshes(CreationFuncForMeshes JSFunc) *Octree {
	o.p.Set("CreationFuncForMeshes", js.FuncOf(CreationFuncForMeshes))
	return o
}

// CreationFuncForSubMeshes returns the CreationFuncForSubMeshes property of class Octree.
//
// https://doc.babylonjs.com/api/classes/babylon.octree#creationfuncforsubmeshes
func (o *Octree) CreationFuncForSubMeshes() js.Value {
	retVal := o.p.Get("CreationFuncForSubMeshes")
	return retVal
}

// SetCreationFuncForSubMeshes sets the CreationFuncForSubMeshes property of class Octree.
//
// https://doc.babylonjs.com/api/classes/babylon.octree#creationfuncforsubmeshes
func (o *Octree) SetCreationFuncForSubMeshes(CreationFuncForSubMeshes JSFunc) *Octree {
	o.p.Set("CreationFuncForSubMeshes", js.FuncOf(CreationFuncForSubMeshes))
	return o
}

// DynamicContent returns the DynamicContent property of class Octree.
//
// https://doc.babylonjs.com/api/classes/babylon.octree#dynamiccontent
func (o *Octree) DynamicContent() []*T {
	retVal := o.p.Get("dynamicContent")
	result := []*T{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, TFromJSObject(retVal.Index(ri), o.ctx))
	}
	return result
}

// SetDynamicContent sets the DynamicContent property of class Octree.
//
// https://doc.babylonjs.com/api/classes/babylon.octree#dynamiccontent
func (o *Octree) SetDynamicContent(dynamicContent []*T) *Octree {
	o.p.Set("dynamicContent", dynamicContent)
	return o
}

// MaxDepth returns the MaxDepth property of class Octree.
//
// https://doc.babylonjs.com/api/classes/babylon.octree#maxdepth
func (o *Octree) MaxDepth() float64 {
	retVal := o.p.Get("maxDepth")
	return retVal.Float()
}

// SetMaxDepth sets the MaxDepth property of class Octree.
//
// https://doc.babylonjs.com/api/classes/babylon.octree#maxdepth
func (o *Octree) SetMaxDepth(maxDepth float64) *Octree {
	o.p.Set("maxDepth", maxDepth)
	return o
}
