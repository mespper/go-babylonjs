// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// EdgesRenderer represents a babylon.js EdgesRenderer.
// This class is used to generate edges of the mesh that could then easily be rendered in a scene.
type EdgesRenderer struct{ p js.Value }

// JSObject returns the underlying js.Value.
func (e *EdgesRenderer) JSObject() js.Value { return e.p }

// EdgesRenderer returns a EdgesRenderer JavaScript class.
func (ba *Babylon) EdgesRenderer() *EdgesRenderer {
	p := ba.ctx.Get("EdgesRenderer")
	return EdgesRendererFromJSObject(p)
}

// EdgesRendererFromJSObject returns a wrapped EdgesRenderer JavaScript class.
func EdgesRendererFromJSObject(p js.Value) *EdgesRenderer {
	return &EdgesRenderer{p: p}
}

// NewEdgesRendererOpts contains optional parameters for NewEdgesRenderer.
type NewEdgesRendererOpts struct {
	Epsilon *JSFloat64

	CheckVerticesInsteadOfIndices *JSBool

	GenerateEdgesLines *JSBool
}

// NewEdgesRenderer returns a new EdgesRenderer object.
//
// https://doc.babylonjs.com/api/classes/babylon.edgesrenderer
func (ba *Babylon) NewEdgesRenderer(source *AbstractMesh, opts *NewEdgesRendererOpts) *EdgesRenderer {
	if opts == nil {
		opts = &NewEdgesRendererOpts{}
	}

	p := ba.ctx.Get("EdgesRenderer").New(source.JSObject(), opts.Epsilon, opts.CheckVerticesInsteadOfIndices, opts.GenerateEdgesLines)
	return EdgesRendererFromJSObject(p)
}

// TODO: methods
