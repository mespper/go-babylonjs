// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// DiscBuilder represents a babylon.js DiscBuilder.
// Class containing static functions to help procedurally build meshes
type DiscBuilder struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (d *DiscBuilder) JSObject() js.Value { return d.p }

// DiscBuilder returns a DiscBuilder JavaScript class.
func (ba *Babylon) DiscBuilder() *DiscBuilder {
	p := ba.ctx.Get("DiscBuilder")
	return DiscBuilderFromJSObject(p, ba.ctx)
}

// DiscBuilderFromJSObject returns a wrapped DiscBuilder JavaScript class.
func DiscBuilderFromJSObject(p js.Value, ctx js.Value) *DiscBuilder {
	return &DiscBuilder{p: p, ctx: ctx}
}

// DiscBuilderArrayToJSArray returns a JavaScript Array for the wrapped array.
func DiscBuilderArrayToJSArray(array []*DiscBuilder) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// DiscBuilderCreateDiscOpts contains optional parameters for DiscBuilder.CreateDisc.
type DiscBuilderCreateDiscOpts struct {
	Scene *Scene
}

// CreateDisc calls the CreateDisc method on the DiscBuilder object.
//
// https://doc.babylonjs.com/api/classes/babylon.discbuilder#createdisc
func (d *DiscBuilder) CreateDisc(name string, options js.Value, opts *DiscBuilderCreateDiscOpts) *Mesh {
	if opts == nil {
		opts = &DiscBuilderCreateDiscOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, name)

	args = append(args, options)

	if opts.Scene == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Scene.JSObject())
	}

	retVal := d.p.Call("CreateDisc", args...)
	return MeshFromJSObject(retVal, d.ctx)
}
