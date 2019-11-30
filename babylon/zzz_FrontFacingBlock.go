// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// FrontFacingBlock represents a babylon.js FrontFacingBlock.
// Block used to test if the fragment shader is front facing
type FrontFacingBlock struct{ *NodeMaterialBlock }

// JSObject returns the underlying js.Value.
func (f *FrontFacingBlock) JSObject() js.Value { return f.p }

// FrontFacingBlock returns a FrontFacingBlock JavaScript class.
func (ba *Babylon) FrontFacingBlock() *FrontFacingBlock {
	p := ba.ctx.Get("FrontFacingBlock")
	return FrontFacingBlockFromJSObject(p)
}

// FrontFacingBlockFromJSObject returns a wrapped FrontFacingBlock JavaScript class.
func FrontFacingBlockFromJSObject(p js.Value) *FrontFacingBlock {
	return &FrontFacingBlock{NodeMaterialBlockFromJSObject(p)}
}

// NewFrontFacingBlock returns a new FrontFacingBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.frontfacingblock
func (ba *Babylon) NewFrontFacingBlock(name string) *FrontFacingBlock {
	p := ba.ctx.Get("FrontFacingBlock").New(name)
	return FrontFacingBlockFromJSObject(p)
}

// TODO: methods
