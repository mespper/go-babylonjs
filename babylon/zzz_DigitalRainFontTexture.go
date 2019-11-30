// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// DigitalRainFontTexture represents a babylon.js DigitalRainFontTexture.
// DigitalRainFontTexture is the helper class used to easily create your digital rain font texture.
//
// It basically takes care rendering the font front the given font size to a texture.
// This is used later on in the postprocess.
type DigitalRainFontTexture struct{ *BaseTexture }

// JSObject returns the underlying js.Value.
func (d *DigitalRainFontTexture) JSObject() js.Value { return d.p }

// DigitalRainFontTexture returns a DigitalRainFontTexture JavaScript class.
func (ba *Babylon) DigitalRainFontTexture() *DigitalRainFontTexture {
	p := ba.ctx.Get("DigitalRainFontTexture")
	return DigitalRainFontTextureFromJSObject(p)
}

// DigitalRainFontTextureFromJSObject returns a wrapped DigitalRainFontTexture JavaScript class.
func DigitalRainFontTextureFromJSObject(p js.Value) *DigitalRainFontTexture {
	return &DigitalRainFontTexture{BaseTextureFromJSObject(p)}
}

// NewDigitalRainFontTextureOpts contains optional parameters for NewDigitalRainFontTexture.
type NewDigitalRainFontTextureOpts struct {
	Scene *Scene
}

// NewDigitalRainFontTexture returns a new DigitalRainFontTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainfonttexture
func (ba *Babylon) NewDigitalRainFontTexture(name string, font string, text string, opts *NewDigitalRainFontTextureOpts) *DigitalRainFontTexture {
	if opts == nil {
		opts = &NewDigitalRainFontTextureOpts{}
	}

	p := ba.ctx.Get("DigitalRainFontTexture").New(name, font, text, opts.Scene.JSObject())
	return DigitalRainFontTextureFromJSObject(p)
}

// TODO: methods
