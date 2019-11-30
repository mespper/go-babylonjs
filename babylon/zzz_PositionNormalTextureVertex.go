// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// PositionNormalTextureVertex represents a babylon.js PositionNormalTextureVertex.
// Contains position, normal and uv vectors for a vertex
type PositionNormalTextureVertex struct{ p js.Value }

// JSObject returns the underlying js.Value.
func (p *PositionNormalTextureVertex) JSObject() js.Value { return p.p }

// PositionNormalTextureVertex returns a PositionNormalTextureVertex JavaScript class.
func (ba *Babylon) PositionNormalTextureVertex() *PositionNormalTextureVertex {
	p := ba.ctx.Get("PositionNormalTextureVertex")
	return PositionNormalTextureVertexFromJSObject(p)
}

// PositionNormalTextureVertexFromJSObject returns a wrapped PositionNormalTextureVertex JavaScript class.
func PositionNormalTextureVertexFromJSObject(p js.Value) *PositionNormalTextureVertex {
	return &PositionNormalTextureVertex{p: p}
}

// NewPositionNormalTextureVertexOpts contains optional parameters for NewPositionNormalTextureVertex.
type NewPositionNormalTextureVertexOpts struct {
	Position *Vector3

	Normal *Vector3

	Uv *Vector2
}

// NewPositionNormalTextureVertex returns a new PositionNormalTextureVertex object.
//
// https://doc.babylonjs.com/api/classes/babylon.positionnormaltexturevertex
func (ba *Babylon) NewPositionNormalTextureVertex(opts *NewPositionNormalTextureVertexOpts) *PositionNormalTextureVertex {
	if opts == nil {
		opts = &NewPositionNormalTextureVertexOpts{}
	}

	p := ba.ctx.Get("PositionNormalTextureVertex").New(opts.Position.JSObject(), opts.Normal.JSObject(), opts.Uv.JSObject())
	return PositionNormalTextureVertexFromJSObject(p)
}

// TODO: methods
