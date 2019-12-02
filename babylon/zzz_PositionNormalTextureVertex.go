// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// PositionNormalTextureVertex represents a babylon.js PositionNormalTextureVertex.
// Contains position, normal and uv vectors for a vertex
type PositionNormalTextureVertex struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (p *PositionNormalTextureVertex) JSObject() js.Value { return p.p }

// PositionNormalTextureVertex returns a PositionNormalTextureVertex JavaScript class.
func (ba *Babylon) PositionNormalTextureVertex() *PositionNormalTextureVertex {
	p := ba.ctx.Get("PositionNormalTextureVertex")
	return PositionNormalTextureVertexFromJSObject(p, ba.ctx)
}

// PositionNormalTextureVertexFromJSObject returns a wrapped PositionNormalTextureVertex JavaScript class.
func PositionNormalTextureVertexFromJSObject(p js.Value, ctx js.Value) *PositionNormalTextureVertex {
	return &PositionNormalTextureVertex{p: p, ctx: ctx}
}

// PositionNormalTextureVertexArrayToJSArray returns a JavaScript Array for the wrapped array.
func PositionNormalTextureVertexArrayToJSArray(array []*PositionNormalTextureVertex) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewPositionNormalTextureVertexOpts contains optional parameters for NewPositionNormalTextureVertex.
type NewPositionNormalTextureVertexOpts struct {
	Position *Vector3
	Normal   *Vector3
	Uv       *Vector2
}

// NewPositionNormalTextureVertex returns a new PositionNormalTextureVertex object.
//
// https://doc.babylonjs.com/api/classes/babylon.positionnormaltexturevertex
func (ba *Babylon) NewPositionNormalTextureVertex(opts *NewPositionNormalTextureVertexOpts) *PositionNormalTextureVertex {
	if opts == nil {
		opts = &NewPositionNormalTextureVertexOpts{}
	}

	args := make([]interface{}, 0, 0+3)

	if opts.Position == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Position.JSObject())
	}
	if opts.Normal == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Normal.JSObject())
	}
	if opts.Uv == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Uv.JSObject())
	}

	p := ba.ctx.Get("PositionNormalTextureVertex").New(args...)
	return PositionNormalTextureVertexFromJSObject(p, ba.ctx)
}

// Clone calls the Clone method on the PositionNormalTextureVertex object.
//
// https://doc.babylonjs.com/api/classes/babylon.positionnormaltexturevertex#clone
func (p *PositionNormalTextureVertex) Clone() *PositionNormalTextureVertex {

	retVal := p.p.Call("clone")
	return PositionNormalTextureVertexFromJSObject(retVal, p.ctx)
}

// Normal returns the Normal property of class PositionNormalTextureVertex.
//
// https://doc.babylonjs.com/api/classes/babylon.positionnormaltexturevertex#normal
func (p *PositionNormalTextureVertex) Normal() *Vector3 {
	retVal := p.p.Get("normal")
	return Vector3FromJSObject(retVal, p.ctx)
}

// SetNormal sets the Normal property of class PositionNormalTextureVertex.
//
// https://doc.babylonjs.com/api/classes/babylon.positionnormaltexturevertex#normal
func (p *PositionNormalTextureVertex) SetNormal(normal *Vector3) *PositionNormalTextureVertex {
	p.p.Set("normal", normal.JSObject())
	return p
}

// Position returns the Position property of class PositionNormalTextureVertex.
//
// https://doc.babylonjs.com/api/classes/babylon.positionnormaltexturevertex#position
func (p *PositionNormalTextureVertex) Position() *Vector3 {
	retVal := p.p.Get("position")
	return Vector3FromJSObject(retVal, p.ctx)
}

// SetPosition sets the Position property of class PositionNormalTextureVertex.
//
// https://doc.babylonjs.com/api/classes/babylon.positionnormaltexturevertex#position
func (p *PositionNormalTextureVertex) SetPosition(position *Vector3) *PositionNormalTextureVertex {
	p.p.Set("position", position.JSObject())
	return p
}

// Uv returns the Uv property of class PositionNormalTextureVertex.
//
// https://doc.babylonjs.com/api/classes/babylon.positionnormaltexturevertex#uv
func (p *PositionNormalTextureVertex) Uv() *Vector2 {
	retVal := p.p.Get("uv")
	return Vector2FromJSObject(retVal, p.ctx)
}

// SetUv sets the Uv property of class PositionNormalTextureVertex.
//
// https://doc.babylonjs.com/api/classes/babylon.positionnormaltexturevertex#uv
func (p *PositionNormalTextureVertex) SetUv(uv *Vector2) *PositionNormalTextureVertex {
	p.p.Set("uv", uv.JSObject())
	return p
}
