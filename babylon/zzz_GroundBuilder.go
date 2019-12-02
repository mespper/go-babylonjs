// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// GroundBuilder represents a babylon.js GroundBuilder.
// Class containing static functions to help procedurally build meshes
type GroundBuilder struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (g *GroundBuilder) JSObject() js.Value { return g.p }

// GroundBuilder returns a GroundBuilder JavaScript class.
func (ba *Babylon) GroundBuilder() *GroundBuilder {
	p := ba.ctx.Get("GroundBuilder")
	return GroundBuilderFromJSObject(p, ba.ctx)
}

// GroundBuilderFromJSObject returns a wrapped GroundBuilder JavaScript class.
func GroundBuilderFromJSObject(p js.Value, ctx js.Value) *GroundBuilder {
	return &GroundBuilder{p: p, ctx: ctx}
}

// GroundBuilderArrayToJSArray returns a JavaScript Array for the wrapped array.
func GroundBuilderArrayToJSArray(array []*GroundBuilder) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// CreateGround calls the CreateGround method on the GroundBuilder object.
//
// https://doc.babylonjs.com/api/classes/babylon.groundbuilder#createground
func (g *GroundBuilder) CreateGround(name string, options js.Value, scene interface{}) *Mesh {

	args := make([]interface{}, 0, 3+0)

	args = append(args, name)
	args = append(args, options)
	args = append(args, scene)

	retVal := g.p.Call("CreateGround", args...)
	return MeshFromJSObject(retVal, g.ctx)
}

// GroundBuilderCreateGroundFromHeightMapOpts contains optional parameters for GroundBuilder.CreateGroundFromHeightMap.
type GroundBuilderCreateGroundFromHeightMapOpts struct {
	Scene *Scene
}

// CreateGroundFromHeightMap calls the CreateGroundFromHeightMap method on the GroundBuilder object.
//
// https://doc.babylonjs.com/api/classes/babylon.groundbuilder#creategroundfromheightmap
func (g *GroundBuilder) CreateGroundFromHeightMap(name string, url string, options js.Value, opts *GroundBuilderCreateGroundFromHeightMapOpts) *GroundMesh {
	if opts == nil {
		opts = &GroundBuilderCreateGroundFromHeightMapOpts{}
	}

	args := make([]interface{}, 0, 3+1)

	args = append(args, name)
	args = append(args, url)
	args = append(args, options)

	if opts.Scene == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Scene.JSObject())
	}

	retVal := g.p.Call("CreateGroundFromHeightMap", args...)
	return GroundMeshFromJSObject(retVal, g.ctx)
}

// GroundBuilderCreateTiledGroundOpts contains optional parameters for GroundBuilder.CreateTiledGround.
type GroundBuilderCreateTiledGroundOpts struct {
	Scene *Scene
}

// CreateTiledGround calls the CreateTiledGround method on the GroundBuilder object.
//
// https://doc.babylonjs.com/api/classes/babylon.groundbuilder#createtiledground
func (g *GroundBuilder) CreateTiledGround(name string, options js.Value, opts *GroundBuilderCreateTiledGroundOpts) *Mesh {
	if opts == nil {
		opts = &GroundBuilderCreateTiledGroundOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, name)
	args = append(args, options)

	if opts.Scene == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Scene.JSObject())
	}

	retVal := g.p.Call("CreateTiledGround", args...)
	return MeshFromJSObject(retVal, g.ctx)
}
