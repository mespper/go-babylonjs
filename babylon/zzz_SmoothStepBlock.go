// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// SmoothStepBlock represents a babylon.js SmoothStepBlock.
// Block used to smooth step a value
type SmoothStepBlock struct{ *NodeMaterialBlock }

// JSObject returns the underlying js.Value.
func (s *SmoothStepBlock) JSObject() js.Value { return s.p }

// SmoothStepBlock returns a SmoothStepBlock JavaScript class.
func (b *Babylon) SmoothStepBlock() *SmoothStepBlock {
	p := b.ctx.Get("SmoothStepBlock")
	return SmoothStepBlockFromJSObject(p)
}

// SmoothStepBlockFromJSObject returns a wrapped SmoothStepBlock JavaScript class.
func SmoothStepBlockFromJSObject(p js.Value) *SmoothStepBlock {
	return &SmoothStepBlock{NodeMaterialBlockFromJSObject(p)}
}

// NewSmoothStepBlock returns a new SmoothStepBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.smoothstepblock
func (b *Babylon) NewSmoothStepBlock(todo parameters) *SmoothStepBlock {
	p := b.ctx.Get("SmoothStepBlock").New(todo)
	return SmoothStepBlockFromJSObject(p)
}

// TODO: methods