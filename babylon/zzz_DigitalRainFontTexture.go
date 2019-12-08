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
type DigitalRainFontTexture struct {
	*BaseTexture
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (d *DigitalRainFontTexture) JSObject() js.Value { return d.p }

// DigitalRainFontTexture returns a DigitalRainFontTexture JavaScript class.
func (ba *Babylon) DigitalRainFontTexture() *DigitalRainFontTexture {
	p := ba.ctx.Get("DigitalRainFontTexture")
	return DigitalRainFontTextureFromJSObject(p, ba.ctx)
}

// DigitalRainFontTextureFromJSObject returns a wrapped DigitalRainFontTexture JavaScript class.
func DigitalRainFontTextureFromJSObject(p js.Value, ctx js.Value) *DigitalRainFontTexture {
	return &DigitalRainFontTexture{BaseTexture: BaseTextureFromJSObject(p, ctx), ctx: ctx}
}

// DigitalRainFontTextureArrayToJSArray returns a JavaScript Array for the wrapped array.
func DigitalRainFontTextureArrayToJSArray(array []*DigitalRainFontTexture) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
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

	args := make([]interface{}, 0, 3+1)

	args = append(args, name)
	args = append(args, font)
	args = append(args, text)

	if opts.Scene == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Scene.JSObject())
	}

	p := ba.ctx.Get("DigitalRainFontTexture").New(args...)
	return DigitalRainFontTextureFromJSObject(p, ba.ctx)
}

// Clone calls the Clone method on the DigitalRainFontTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainfonttexture#clone
func (d *DigitalRainFontTexture) Clone() *DigitalRainFontTexture {

	retVal := d.p.Call("clone")
	return DigitalRainFontTextureFromJSObject(retVal, d.ctx)
}

// Parse calls the Parse method on the DigitalRainFontTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainfonttexture#parse
func (d *DigitalRainFontTexture) Parse(source JSObject, scene *Scene) *DigitalRainFontTexture {

	args := make([]interface{}, 0, 2+0)

	if source == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, source.JSObject())
	}

	if scene == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, scene.JSObject())
	}

	retVal := d.p.Call("Parse", args...)
	return DigitalRainFontTextureFromJSObject(retVal, d.ctx)
}

// CharSize returns the CharSize property of class DigitalRainFontTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainfonttexture#charsize
func (d *DigitalRainFontTexture) CharSize() float64 {
	retVal := d.p.Get("charSize")
	return retVal.Float()
}

// SetCharSize sets the CharSize property of class DigitalRainFontTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.digitalrainfonttexture#charsize
func (d *DigitalRainFontTexture) SetCharSize(charSize float64) *DigitalRainFontTexture {
	d.p.Set("charSize", charSize)
	return d
}
