// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// TubeBuilder represents a babylon.js TubeBuilder.
// Class containing static functions to help procedurally build meshes
type TubeBuilder struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (t *TubeBuilder) JSObject() js.Value { return t.p }

// TubeBuilder returns a TubeBuilder JavaScript class.
func (ba *Babylon) TubeBuilder() *TubeBuilder {
	p := ba.ctx.Get("TubeBuilder")
	return TubeBuilderFromJSObject(p, ba.ctx)
}

// TubeBuilderFromJSObject returns a wrapped TubeBuilder JavaScript class.
func TubeBuilderFromJSObject(p js.Value, ctx js.Value) *TubeBuilder {
	return &TubeBuilder{p: p, ctx: ctx}
}

// TubeBuilderArrayToJSArray returns a JavaScript Array for the wrapped array.
func TubeBuilderArrayToJSArray(array []*TubeBuilder) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// TubeBuilderCreateTubeOpts contains optional parameters for TubeBuilder.CreateTube.
type TubeBuilderCreateTubeOpts struct {
	Scene *Scene
}

// CreateTube calls the CreateTube method on the TubeBuilder object.
//
// https://doc.babylonjs.com/api/classes/babylon.tubebuilder#createtube
func (t *TubeBuilder) CreateTube(name string, options js.Value, opts *TubeBuilderCreateTubeOpts) *Mesh {
	if opts == nil {
		opts = &TubeBuilderCreateTubeOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, name)
	args = append(args, options)

	if opts.Scene == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Scene.JSObject())
	}

	retVal := t.p.Call("CreateTube", args...)
	return MeshFromJSObject(retVal, t.ctx)
}
