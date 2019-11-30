// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// RefractBlock represents a babylon.js RefractBlock.
// Block used to get the refracted vector from a direction and a normal
type RefractBlock struct{ *NodeMaterialBlock }

// JSObject returns the underlying js.Value.
func (r *RefractBlock) JSObject() js.Value { return r.p }

// RefractBlock returns a RefractBlock JavaScript class.
func (ba *Babylon) RefractBlock() *RefractBlock {
	p := ba.ctx.Get("RefractBlock")
	return RefractBlockFromJSObject(p)
}

// RefractBlockFromJSObject returns a wrapped RefractBlock JavaScript class.
func RefractBlockFromJSObject(p js.Value) *RefractBlock {
	return &RefractBlock{NodeMaterialBlockFromJSObject(p)}
}

// NewRefractBlock returns a new RefractBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.refractblock
func (ba *Babylon) NewRefractBlock(name string) *RefractBlock {
	p := ba.ctx.Get("RefractBlock").New(name)
	return RefractBlockFromJSObject(p)
}

// TODO: methods
