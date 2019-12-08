// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// TorusKnotBuilder represents a babylon.js TorusKnotBuilder.
// Class containing static functions to help procedurally build meshes
type TorusKnotBuilder struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (t *TorusKnotBuilder) JSObject() js.Value { return t.p }

// TorusKnotBuilder returns a TorusKnotBuilder JavaScript class.
func (ba *Babylon) TorusKnotBuilder() *TorusKnotBuilder {
	p := ba.ctx.Get("TorusKnotBuilder")
	return TorusKnotBuilderFromJSObject(p, ba.ctx)
}

// TorusKnotBuilderFromJSObject returns a wrapped TorusKnotBuilder JavaScript class.
func TorusKnotBuilderFromJSObject(p js.Value, ctx js.Value) *TorusKnotBuilder {
	return &TorusKnotBuilder{p: p, ctx: ctx}
}

// TorusKnotBuilderArrayToJSArray returns a JavaScript Array for the wrapped array.
func TorusKnotBuilderArrayToJSArray(array []*TorusKnotBuilder) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// CreateTorusKnot calls the CreateTorusKnot method on the TorusKnotBuilder object.
//
// https://doc.babylonjs.com/api/classes/babylon.torusknotbuilder#createtorusknot
func (t *TorusKnotBuilder) CreateTorusKnot(name string, options js.Value, scene JSObject) *Mesh {

	args := make([]interface{}, 0, 3+0)

	args = append(args, name)

	args = append(args, options)

	if scene == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, scene.JSObject())
	}

	retVal := t.p.Call("CreateTorusKnot", args...)
	return MeshFromJSObject(retVal, t.ctx)
}
