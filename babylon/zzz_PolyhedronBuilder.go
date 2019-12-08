// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// PolyhedronBuilder represents a babylon.js PolyhedronBuilder.
// Class containing static functions to help procedurally build meshes
type PolyhedronBuilder struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (p *PolyhedronBuilder) JSObject() js.Value { return p.p }

// PolyhedronBuilder returns a PolyhedronBuilder JavaScript class.
func (ba *Babylon) PolyhedronBuilder() *PolyhedronBuilder {
	p := ba.ctx.Get("PolyhedronBuilder")
	return PolyhedronBuilderFromJSObject(p, ba.ctx)
}

// PolyhedronBuilderFromJSObject returns a wrapped PolyhedronBuilder JavaScript class.
func PolyhedronBuilderFromJSObject(p js.Value, ctx js.Value) *PolyhedronBuilder {
	return &PolyhedronBuilder{p: p, ctx: ctx}
}

// PolyhedronBuilderArrayToJSArray returns a JavaScript Array for the wrapped array.
func PolyhedronBuilderArrayToJSArray(array []*PolyhedronBuilder) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// PolyhedronBuilderCreatePolyhedronOpts contains optional parameters for PolyhedronBuilder.CreatePolyhedron.
type PolyhedronBuilderCreatePolyhedronOpts struct {
	Scene *Scene
}

// CreatePolyhedron calls the CreatePolyhedron method on the PolyhedronBuilder object.
//
// https://doc.babylonjs.com/api/classes/babylon.polyhedronbuilder#createpolyhedron
func (p *PolyhedronBuilder) CreatePolyhedron(name string, options js.Value, opts *PolyhedronBuilderCreatePolyhedronOpts) *Mesh {
	if opts == nil {
		opts = &PolyhedronBuilderCreatePolyhedronOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, name)

	args = append(args, options)

	if opts.Scene == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Scene.JSObject())
	}

	retVal := p.p.Call("CreatePolyhedron", args...)
	return MeshFromJSObject(retVal, p.ctx)
}
