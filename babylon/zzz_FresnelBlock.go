// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// FresnelBlock represents a babylon.js FresnelBlock.
// Block used to compute fresnel value
type FresnelBlock struct{ *NodeMaterialBlock }

// JSObject returns the underlying js.Value.
func (f *FresnelBlock) JSObject() js.Value { return f.p }

// FresnelBlock returns a FresnelBlock JavaScript class.
func (ba *Babylon) FresnelBlock() *FresnelBlock {
	p := ba.ctx.Get("FresnelBlock")
	return FresnelBlockFromJSObject(p)
}

// FresnelBlockFromJSObject returns a wrapped FresnelBlock JavaScript class.
func FresnelBlockFromJSObject(p js.Value) *FresnelBlock {
	return &FresnelBlock{NodeMaterialBlockFromJSObject(p)}
}

// NewFresnelBlock returns a new FresnelBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.fresnelblock
func (ba *Babylon) NewFresnelBlock(name string) *FresnelBlock {
	p := ba.ctx.Get("FresnelBlock").New(name)
	return FresnelBlockFromJSObject(p)
}

// TODO: methods
