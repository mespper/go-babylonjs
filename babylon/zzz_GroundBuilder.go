// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// GroundBuilder represents a babylon.js GroundBuilder.
// Class containing static functions to help procedurally build meshes
type GroundBuilder struct{}

// JSObject returns the underlying js.Value.
func (g *GroundBuilder) JSObject() js.Value { return g.p }

// GroundBuilder returns a GroundBuilder JavaScript class.
func (b *Babylon) GroundBuilder() *GroundBuilder {
	p := b.ctx.Get("GroundBuilder")
	return GroundBuilderFromJSObject(p)
}

// GroundBuilderFromJSObject returns a wrapped GroundBuilder JavaScript class.
func GroundBuilderFromJSObject(p js.Value) *GroundBuilder {
	return &GroundBuilder{p: p}
}

// NewGroundBuilder returns a new GroundBuilder object.
//
// https://doc.babylonjs.com/api/classes/babylon.groundbuilder
func (b *Babylon) NewGroundBuilder(todo parameters) *GroundBuilder {
	p := b.ctx.Get("GroundBuilder").New(todo)
	return GroundBuilderFromJSObject(p)
}

// TODO: methods