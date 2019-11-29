// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// Bone represents a babylon.js Bone.
// Class used to store bone information

//
// See: http://doc.babylonjs.com/how_to/how_to_use_bones_and_skeletons
type Bone struct{ *Node }

// JSObject returns the underlying js.Value.
func (b *Bone) JSObject() js.Value { return b.p }

// Bone returns a Bone JavaScript class.
func (b *Babylon) Bone() *Bone {
	p := b.ctx.Get("Bone")
	return BoneFromJSObject(p)
}

// BoneFromJSObject returns a wrapped Bone JavaScript class.
func BoneFromJSObject(p js.Value) *Bone {
	return &Bone{NodeFromJSObject(p)}
}

// NewBone returns a new Bone object.
//
// https://doc.babylonjs.com/api/classes/babylon.bone
func (b *Babylon) NewBone(todo parameters) *Bone {
	p := b.ctx.Get("Bone").New(todo)
	return BoneFromJSObject(p)
}

// TODO: methods