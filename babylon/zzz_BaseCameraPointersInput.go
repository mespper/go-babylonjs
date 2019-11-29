// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// BaseCameraPointersInput represents a babylon.js BaseCameraPointersInput.
// Base class for Camera Pointer Inputs.
// See FollowCameraPointersInput in src/Cameras/Inputs/followCameraPointersInput.ts
// for example usage.
type BaseCameraPointersInput struct{}

// JSObject returns the underlying js.Value.
func (b *BaseCameraPointersInput) JSObject() js.Value { return b.p }

// BaseCameraPointersInput returns a BaseCameraPointersInput JavaScript class.
func (b *Babylon) BaseCameraPointersInput() *BaseCameraPointersInput {
	p := b.ctx.Get("BaseCameraPointersInput")
	return BaseCameraPointersInputFromJSObject(p)
}

// BaseCameraPointersInputFromJSObject returns a wrapped BaseCameraPointersInput JavaScript class.
func BaseCameraPointersInputFromJSObject(p js.Value) *BaseCameraPointersInput {
	return &BaseCameraPointersInput{p: p}
}

// NewBaseCameraPointersInput returns a new BaseCameraPointersInput object.
//
// https://doc.babylonjs.com/api/classes/babylon.basecamerapointersinput
func (b *Babylon) NewBaseCameraPointersInput(todo parameters) *BaseCameraPointersInput {
	p := b.ctx.Get("BaseCameraPointersInput").New(todo)
	return BaseCameraPointersInputFromJSObject(p)
}

// TODO: methods