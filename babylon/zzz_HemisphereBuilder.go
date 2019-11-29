// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// HemisphereBuilder represents a babylon.js HemisphereBuilder.
// Class containing static functions to help procedurally build meshes
type HemisphereBuilder struct{}

// JSObject returns the underlying js.Value.
func (h *HemisphereBuilder) JSObject() js.Value { return h.p }

// HemisphereBuilder returns a HemisphereBuilder JavaScript class.
func (b *Babylon) HemisphereBuilder() *HemisphereBuilder {
	p := b.ctx.Get("HemisphereBuilder")
	return HemisphereBuilderFromJSObject(p)
}

// HemisphereBuilderFromJSObject returns a wrapped HemisphereBuilder JavaScript class.
func HemisphereBuilderFromJSObject(p js.Value) *HemisphereBuilder {
	return &HemisphereBuilder{p: p}
}

// NewHemisphereBuilder returns a new HemisphereBuilder object.
//
// https://doc.babylonjs.com/api/classes/babylon.hemispherebuilder
func (b *Babylon) NewHemisphereBuilder(todo parameters) *HemisphereBuilder {
	p := b.ctx.Get("HemisphereBuilder").New(todo)
	return HemisphereBuilderFromJSObject(p)
}

// TODO: methods