// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// AsciiArtFontTexture represents a babylon.js AsciiArtFontTexture.
// AsciiArtFontTexture is the helper class used to easily create your ascii art font texture.
//
// It basically takes care rendering the font front the given font size to a texture.
// This is used later on in the postprocess.
type AsciiArtFontTexture struct {
	*BaseTexture
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (a *AsciiArtFontTexture) JSObject() js.Value { return a.p }

// AsciiArtFontTexture returns a AsciiArtFontTexture JavaScript class.
func (ba *Babylon) AsciiArtFontTexture() *AsciiArtFontTexture {
	p := ba.ctx.Get("AsciiArtFontTexture")
	return AsciiArtFontTextureFromJSObject(p, ba.ctx)
}

// AsciiArtFontTextureFromJSObject returns a wrapped AsciiArtFontTexture JavaScript class.
func AsciiArtFontTextureFromJSObject(p js.Value, ctx js.Value) *AsciiArtFontTexture {
	return &AsciiArtFontTexture{BaseTexture: BaseTextureFromJSObject(p, ctx), ctx: ctx}
}

// AsciiArtFontTextureArrayToJSArray returns a JavaScript Array for the wrapped array.
func AsciiArtFontTextureArrayToJSArray(array []*AsciiArtFontTexture) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewAsciiArtFontTextureOpts contains optional parameters for NewAsciiArtFontTexture.
type NewAsciiArtFontTextureOpts struct {
	Scene *Scene
}

// NewAsciiArtFontTexture returns a new AsciiArtFontTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture
func (ba *Babylon) NewAsciiArtFontTexture(name string, font string, text string, opts *NewAsciiArtFontTextureOpts) *AsciiArtFontTexture {
	if opts == nil {
		opts = &NewAsciiArtFontTextureOpts{}
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

	p := ba.ctx.Get("AsciiArtFontTexture").New(args...)
	return AsciiArtFontTextureFromJSObject(p, ba.ctx)
}

// Clone calls the Clone method on the AsciiArtFontTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#clone
func (a *AsciiArtFontTexture) Clone() *AsciiArtFontTexture {

	retVal := a.p.Call("clone")
	return AsciiArtFontTextureFromJSObject(retVal, a.ctx)
}

// Parse calls the Parse method on the AsciiArtFontTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#parse
func (a *AsciiArtFontTexture) Parse(source JSObject, scene *Scene) *AsciiArtFontTexture {

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

	retVal := a.p.Call("Parse", args...)
	return AsciiArtFontTextureFromJSObject(retVal, a.ctx)
}

// CharSize returns the CharSize property of class AsciiArtFontTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#charsize
func (a *AsciiArtFontTexture) CharSize() float64 {
	retVal := a.p.Get("charSize")
	return retVal.Float()
}

// SetCharSize sets the CharSize property of class AsciiArtFontTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.asciiartfonttexture#charsize
func (a *AsciiArtFontTexture) SetCharSize(charSize float64) *AsciiArtFontTexture {
	a.p.Set("charSize", charSize)
	return a
}
