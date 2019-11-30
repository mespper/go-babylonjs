// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// GeometryBufferRenderer represents a babylon.js GeometryBufferRenderer.
// This renderer is helpfull to fill one of the render target with a geometry buffer.
type GeometryBufferRenderer struct{ p js.Value }

// JSObject returns the underlying js.Value.
func (g *GeometryBufferRenderer) JSObject() js.Value { return g.p }

// GeometryBufferRenderer returns a GeometryBufferRenderer JavaScript class.
func (ba *Babylon) GeometryBufferRenderer() *GeometryBufferRenderer {
	p := ba.ctx.Get("GeometryBufferRenderer")
	return GeometryBufferRendererFromJSObject(p)
}

// GeometryBufferRendererFromJSObject returns a wrapped GeometryBufferRenderer JavaScript class.
func GeometryBufferRendererFromJSObject(p js.Value) *GeometryBufferRenderer {
	return &GeometryBufferRenderer{p: p}
}

// NewGeometryBufferRendererOpts contains optional parameters for NewGeometryBufferRenderer.
type NewGeometryBufferRendererOpts struct {
	Ratio *JSFloat64
}

// NewGeometryBufferRenderer returns a new GeometryBufferRenderer object.
//
// https://doc.babylonjs.com/api/classes/babylon.geometrybufferrenderer
func (ba *Babylon) NewGeometryBufferRenderer(scene *Scene, opts *NewGeometryBufferRendererOpts) *GeometryBufferRenderer {
	if opts == nil {
		opts = &NewGeometryBufferRendererOpts{}
	}

	p := ba.ctx.Get("GeometryBufferRenderer").New(scene.JSObject(), opts.Ratio.JSObject())
	return GeometryBufferRendererFromJSObject(p)
}

// TODO: methods
