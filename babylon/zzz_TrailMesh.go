// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// TrailMesh represents a babylon.js TrailMesh.
// Class used to create a trail following a mesh
type TrailMesh struct {
	*Mesh
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (t *TrailMesh) JSObject() js.Value { return t.p }

// TrailMesh returns a TrailMesh JavaScript class.
func (ba *Babylon) TrailMesh() *TrailMesh {
	p := ba.ctx.Get("TrailMesh")
	return TrailMeshFromJSObject(p, ba.ctx)
}

// TrailMeshFromJSObject returns a wrapped TrailMesh JavaScript class.
func TrailMeshFromJSObject(p js.Value, ctx js.Value) *TrailMesh {
	return &TrailMesh{Mesh: MeshFromJSObject(p, ctx), ctx: ctx}
}

// TrailMeshArrayToJSArray returns a JavaScript Array for the wrapped array.
func TrailMeshArrayToJSArray(array []*TrailMesh) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewTrailMeshOpts contains optional parameters for NewTrailMesh.
type NewTrailMeshOpts struct {
	Diameter  *float64
	Length    *float64
	AutoStart *bool
}

// NewTrailMesh returns a new TrailMesh object.
//
// https://doc.babylonjs.com/api/classes/babylon.trailmesh
func (ba *Babylon) NewTrailMesh(name string, generator *AbstractMesh, scene *Scene, opts *NewTrailMeshOpts) *TrailMesh {
	if opts == nil {
		opts = &NewTrailMeshOpts{}
	}

	args := make([]interface{}, 0, 3+3)

	args = append(args, name)
	args = append(args, generator.JSObject())
	args = append(args, scene.JSObject())

	if opts.Diameter == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Diameter)
	}
	if opts.Length == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Length)
	}
	if opts.AutoStart == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.AutoStart)
	}

	p := ba.ctx.Get("TrailMesh").New(args...)
	return TrailMeshFromJSObject(p, ba.ctx)
}

// Clone calls the Clone method on the TrailMesh object.
//
// https://doc.babylonjs.com/api/classes/babylon.trailmesh#clone
func (t *TrailMesh) Clone(name string, newGenerator *AbstractMesh) *TrailMesh {

	args := make([]interface{}, 0, 2+0)

	args = append(args, name)

	if newGenerator == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, newGenerator.JSObject())
	}

	retVal := t.p.Call("clone", args...)
	return TrailMeshFromJSObject(retVal, t.ctx)
}

// GetClassName calls the GetClassName method on the TrailMesh object.
//
// https://doc.babylonjs.com/api/classes/babylon.trailmesh#getclassname
func (t *TrailMesh) GetClassName() string {

	retVal := t.p.Call("getClassName")
	return retVal.String()
}

// Parse calls the Parse method on the TrailMesh object.
//
// https://doc.babylonjs.com/api/classes/babylon.trailmesh#parse
func (t *TrailMesh) Parse(parsedMesh JSObject, scene *Scene) *TrailMesh {

	args := make([]interface{}, 0, 2+0)

	if parsedMesh == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, parsedMesh.JSObject())
	}

	if scene == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, scene.JSObject())
	}

	retVal := t.p.Call("Parse", args...)
	return TrailMeshFromJSObject(retVal, t.ctx)
}

// Serialize calls the Serialize method on the TrailMesh object.
//
// https://doc.babylonjs.com/api/classes/babylon.trailmesh#serialize
func (t *TrailMesh) Serialize(serializationObject JSObject) {

	args := make([]interface{}, 0, 1+0)

	if serializationObject == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, serializationObject.JSObject())
	}

	t.p.Call("serialize", args...)
}

// Start calls the Start method on the TrailMesh object.
//
// https://doc.babylonjs.com/api/classes/babylon.trailmesh#start
func (t *TrailMesh) Start() {

	t.p.Call("start")
}

// Stop calls the Stop method on the TrailMesh object.
//
// https://doc.babylonjs.com/api/classes/babylon.trailmesh#stop
func (t *TrailMesh) Stop() {

	t.p.Call("stop")
}

// Update calls the Update method on the TrailMesh object.
//
// https://doc.babylonjs.com/api/classes/babylon.trailmesh#update
func (t *TrailMesh) Update() {

	t.p.Call("update")
}
