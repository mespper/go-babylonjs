// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// CylinderBuilder represents a babylon.js CylinderBuilder.
// Class containing static functions to help procedurally build meshes
type CylinderBuilder struct{}

// JSObject returns the underlying js.Value.
func (c *CylinderBuilder) JSObject() js.Value { return c.p }

// CylinderBuilder returns a CylinderBuilder JavaScript class.
func (b *Babylon) CylinderBuilder() *CylinderBuilder {
	p := b.ctx.Get("CylinderBuilder")
	return CylinderBuilderFromJSObject(p)
}

// CylinderBuilderFromJSObject returns a wrapped CylinderBuilder JavaScript class.
func CylinderBuilderFromJSObject(p js.Value) *CylinderBuilder {
	return &CylinderBuilder{p: p}
}

// NewCylinderBuilder returns a new CylinderBuilder object.
//
// https://doc.babylonjs.com/api/classes/babylon.cylinderbuilder
func (b *Babylon) NewCylinderBuilder(todo parameters) *CylinderBuilder {
	p := b.ctx.Get("CylinderBuilder").New(todo)
	return CylinderBuilderFromJSObject(p)
}

// TODO: methods