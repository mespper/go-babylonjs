// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// SphereBuilder represents a babylon.js SphereBuilder.
// Class containing static functions to help procedurally build meshes
type SphereBuilder struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (s *SphereBuilder) JSObject() js.Value { return s.p }

// SphereBuilder returns a SphereBuilder JavaScript class.
func (ba *Babylon) SphereBuilder() *SphereBuilder {
	p := ba.ctx.Get("SphereBuilder")
	return SphereBuilderFromJSObject(p, ba.ctx)
}

// SphereBuilderFromJSObject returns a wrapped SphereBuilder JavaScript class.
func SphereBuilderFromJSObject(p js.Value, ctx js.Value) *SphereBuilder {
	return &SphereBuilder{p: p, ctx: ctx}
}

// SphereBuilderArrayToJSArray returns a JavaScript Array for the wrapped array.
func SphereBuilderArrayToJSArray(array []*SphereBuilder) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// SphereBuilderCreateSphereOpts contains optional parameters for SphereBuilder.CreateSphere.
type SphereBuilderCreateSphereOpts struct {
	Scene *Scene
}

// CreateSphere calls the CreateSphere method on the SphereBuilder object.
//
// https://doc.babylonjs.com/api/classes/babylon.spherebuilder#createsphere
func (s *SphereBuilder) CreateSphere(name string, options js.Value, opts *SphereBuilderCreateSphereOpts) *Mesh {
	if opts == nil {
		opts = &SphereBuilderCreateSphereOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, name)
	args = append(args, options)

	if opts.Scene == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Scene.JSObject())
	}

	retVal := s.p.Call("CreateSphere", args...)
	return MeshFromJSObject(retVal, s.ctx)
}
