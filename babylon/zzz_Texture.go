// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// Texture represents a babylon.js Texture.
// This represents a texture in babylon. It can be easily loaded from a network, base64 or html input.

//
// See: http://doc.babylonjs.com/babylon101/materials#texture
type Texture struct{ *BaseTexture }

// JSObject returns the underlying js.Value.
func (t *Texture) JSObject() js.Value { return t.p }

// Texture returns a Texture JavaScript class.
func (b *Babylon) Texture() *Texture {
	p := b.ctx.Get("Texture")
	return TextureFromJSObject(p)
}

// TextureFromJSObject returns a wrapped Texture JavaScript class.
func TextureFromJSObject(p js.Value) *Texture {
	return &Texture{BaseTextureFromJSObject(p)}
}

// NewTexture returns a new Texture object.
//
// https://doc.babylonjs.com/api/classes/babylon.texture
func (b *Babylon) NewTexture(todo parameters) *Texture {
	p := b.ctx.Get("Texture").New(todo)
	return TextureFromJSObject(p)
}

// TODO: methods