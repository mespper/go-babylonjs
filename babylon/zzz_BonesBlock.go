// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// BonesBlock represents a babylon.js BonesBlock.
// Block used to add support for vertex skinning (bones)
type BonesBlock struct{ *NodeMaterialBlock }

// JSObject returns the underlying js.Value.
func (b *BonesBlock) JSObject() js.Value { return b.p }

// BonesBlock returns a BonesBlock JavaScript class.
func (ba *Babylon) BonesBlock() *BonesBlock {
	p := ba.ctx.Get("BonesBlock")
	return BonesBlockFromJSObject(p)
}

// BonesBlockFromJSObject returns a wrapped BonesBlock JavaScript class.
func BonesBlockFromJSObject(p js.Value) *BonesBlock {
	return &BonesBlock{NodeMaterialBlockFromJSObject(p)}
}

// NewBonesBlock returns a new BonesBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.bonesblock
func (ba *Babylon) NewBonesBlock(name string) *BonesBlock {
	p := ba.ctx.Get("BonesBlock").New(name)
	return BonesBlockFromJSObject(p)
}

// TODO: methods
