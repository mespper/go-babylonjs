// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// NormalBlendBlock represents a babylon.js NormalBlendBlock.
// Block used to blend normals
type NormalBlendBlock struct{ *NodeMaterialBlock }

// JSObject returns the underlying js.Value.
func (n *NormalBlendBlock) JSObject() js.Value { return n.p }

// NormalBlendBlock returns a NormalBlendBlock JavaScript class.
func (b *Babylon) NormalBlendBlock() *NormalBlendBlock {
	p := b.ctx.Get("NormalBlendBlock")
	return NormalBlendBlockFromJSObject(p)
}

// NormalBlendBlockFromJSObject returns a wrapped NormalBlendBlock JavaScript class.
func NormalBlendBlockFromJSObject(p js.Value) *NormalBlendBlock {
	return &NormalBlendBlock{NodeMaterialBlockFromJSObject(p)}
}

// NewNormalBlendBlock returns a new NormalBlendBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.normalblendblock
func (b *Babylon) NewNormalBlendBlock(todo parameters) *NormalBlendBlock {
	p := b.ctx.Get("NormalBlendBlock").New(todo)
	return NormalBlendBlockFromJSObject(p)
}

// TODO: methods