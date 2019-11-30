// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ScaleBlock represents a babylon.js ScaleBlock.
// Block used to scale a vector by a float
type ScaleBlock struct{ *NodeMaterialBlock }

// JSObject returns the underlying js.Value.
func (s *ScaleBlock) JSObject() js.Value { return s.p }

// ScaleBlock returns a ScaleBlock JavaScript class.
func (ba *Babylon) ScaleBlock() *ScaleBlock {
	p := ba.ctx.Get("ScaleBlock")
	return ScaleBlockFromJSObject(p)
}

// ScaleBlockFromJSObject returns a wrapped ScaleBlock JavaScript class.
func ScaleBlockFromJSObject(p js.Value) *ScaleBlock {
	return &ScaleBlock{NodeMaterialBlockFromJSObject(p)}
}

// NewScaleBlock returns a new ScaleBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.scaleblock
func (ba *Babylon) NewScaleBlock(name string) *ScaleBlock {
	p := ba.ctx.Get("ScaleBlock").New(name)
	return ScaleBlockFromJSObject(p)
}

// TODO: methods
