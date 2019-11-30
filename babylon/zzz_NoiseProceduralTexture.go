// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// NoiseProceduralTexture represents a babylon.js NoiseProceduralTexture.
// Class used to generate noise procedural textures
type NoiseProceduralTexture struct{ *ProceduralTexture }

// JSObject returns the underlying js.Value.
func (n *NoiseProceduralTexture) JSObject() js.Value { return n.p }

// NoiseProceduralTexture returns a NoiseProceduralTexture JavaScript class.
func (ba *Babylon) NoiseProceduralTexture() *NoiseProceduralTexture {
	p := ba.ctx.Get("NoiseProceduralTexture")
	return NoiseProceduralTextureFromJSObject(p)
}

// NoiseProceduralTextureFromJSObject returns a wrapped NoiseProceduralTexture JavaScript class.
func NoiseProceduralTextureFromJSObject(p js.Value) *NoiseProceduralTexture {
	return &NoiseProceduralTexture{ProceduralTextureFromJSObject(p)}
}

// NewNoiseProceduralTextureOpts contains optional parameters for NewNoiseProceduralTexture.
type NewNoiseProceduralTextureOpts struct {
	Size *JSFloat64

	Scene *Scene

	FallbackTexture *Texture

	GenerateMipMaps *JSBool
}

// NewNoiseProceduralTexture returns a new NoiseProceduralTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.noiseproceduraltexture
func (ba *Babylon) NewNoiseProceduralTexture(name string, opts *NewNoiseProceduralTextureOpts) *NoiseProceduralTexture {
	if opts == nil {
		opts = &NewNoiseProceduralTextureOpts{}
	}

	p := ba.ctx.Get("NoiseProceduralTexture").New(name, opts.Size, opts.Scene.JSObject(), opts.FallbackTexture.JSObject(), opts.GenerateMipMaps)
	return NoiseProceduralTextureFromJSObject(p)
}

// TODO: methods
