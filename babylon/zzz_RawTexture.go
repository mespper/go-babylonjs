// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// RawTexture represents a babylon.js RawTexture.
// Raw texture can help creating a texture directly from an array of data.
// This can be super useful if you either get the data from an uncompressed source or
// if you wish to create your texture pixel by pixel.
type RawTexture struct{ *Texture }

// JSObject returns the underlying js.Value.
func (r *RawTexture) JSObject() js.Value { return r.p }

// RawTexture returns a RawTexture JavaScript class.
func (b *Babylon) RawTexture() *RawTexture {
	p := b.ctx.Get("RawTexture")
	return RawTextureFromJSObject(p)
}

// RawTextureFromJSObject returns a wrapped RawTexture JavaScript class.
func RawTextureFromJSObject(p js.Value) *RawTexture {
	return &RawTexture{TextureFromJSObject(p)}
}

// NewRawTexture returns a new RawTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.rawtexture
func (b *Babylon) NewRawTexture(todo parameters) *RawTexture {
	p := b.ctx.Get("RawTexture").New(todo)
	return RawTextureFromJSObject(p)
}

// TODO: methods