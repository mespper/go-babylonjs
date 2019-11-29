// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// FireProceduralTexture represents a babylon.js FireProceduralTexture.
//
type FireProceduralTexture struct{ *ProceduralTexture }

// JSObject returns the underlying js.Value.
func (f *FireProceduralTexture) JSObject() js.Value { return f.p }

// FireProceduralTexture returns a FireProceduralTexture JavaScript class.
func (b *Babylon) FireProceduralTexture() *FireProceduralTexture {
	p := b.ctx.Get("FireProceduralTexture")
	return FireProceduralTextureFromJSObject(p)
}

// FireProceduralTextureFromJSObject returns a wrapped FireProceduralTexture JavaScript class.
func FireProceduralTextureFromJSObject(p js.Value) *FireProceduralTexture {
	return &FireProceduralTexture{ProceduralTextureFromJSObject(p)}
}

// NewFireProceduralTexture returns a new FireProceduralTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.fireproceduraltexture
func (b *Babylon) NewFireProceduralTexture(todo parameters) *FireProceduralTexture {
	p := b.ctx.Get("FireProceduralTexture").New(todo)
	return FireProceduralTextureFromJSObject(p)
}

// TODO: methods