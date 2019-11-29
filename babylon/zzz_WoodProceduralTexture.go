// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// WoodProceduralTexture represents a babylon.js WoodProceduralTexture.
//
type WoodProceduralTexture struct{ *ProceduralTexture }

// JSObject returns the underlying js.Value.
func (w *WoodProceduralTexture) JSObject() js.Value { return w.p }

// WoodProceduralTexture returns a WoodProceduralTexture JavaScript class.
func (b *Babylon) WoodProceduralTexture() *WoodProceduralTexture {
	p := b.ctx.Get("WoodProceduralTexture")
	return WoodProceduralTextureFromJSObject(p)
}

// WoodProceduralTextureFromJSObject returns a wrapped WoodProceduralTexture JavaScript class.
func WoodProceduralTextureFromJSObject(p js.Value) *WoodProceduralTexture {
	return &WoodProceduralTexture{ProceduralTextureFromJSObject(p)}
}

// NewWoodProceduralTexture returns a new WoodProceduralTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.woodproceduraltexture
func (b *Babylon) NewWoodProceduralTexture(todo parameters) *WoodProceduralTexture {
	p := b.ctx.Get("WoodProceduralTexture").New(todo)
	return WoodProceduralTextureFromJSObject(p)
}

// TODO: methods