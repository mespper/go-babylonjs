// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// VertexBuffer represents a babylon.js VertexBuffer.
// Specialized buffer used to store vertex data
type VertexBuffer struct{ p js.Value }

// JSObject returns the underlying js.Value.
func (v *VertexBuffer) JSObject() js.Value { return v.p }

// VertexBuffer returns a VertexBuffer JavaScript class.
func (ba *Babylon) VertexBuffer() *VertexBuffer {
	p := ba.ctx.Get("VertexBuffer")
	return VertexBufferFromJSObject(p)
}

// VertexBufferFromJSObject returns a wrapped VertexBuffer JavaScript class.
func VertexBufferFromJSObject(p js.Value) *VertexBuffer {
	return &VertexBuffer{p: p}
}

// NewVertexBufferOpts contains optional parameters for NewVertexBuffer.
type NewVertexBufferOpts struct {
	PostponeInternalCreation *JSBool

	Stride *JSFloat64

	Instanced *JSBool

	Offset *JSFloat64

	Size *JSFloat64

	Type *JSFloat64

	Normalized *JSBool

	UseBytes *JSBool

	Divisor *JSFloat64
}

// NewVertexBuffer returns a new VertexBuffer object.
//
// https://doc.babylonjs.com/api/classes/babylon.vertexbuffer
func (ba *Babylon) NewVertexBuffer(engine interface{}, data []float64, kind string, updatable bool, opts *NewVertexBufferOpts) *VertexBuffer {
	if opts == nil {
		opts = &NewVertexBufferOpts{}
	}

	p := ba.ctx.Get("VertexBuffer").New(engine, data, kind, updatable, opts.PostponeInternalCreation, opts.Stride, opts.Instanced, opts.Offset, opts.Size, opts.Type, opts.Normalized, opts.UseBytes, opts.Divisor)
	return VertexBufferFromJSObject(p)
}

// TODO: methods
