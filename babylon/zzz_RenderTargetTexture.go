// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// RenderTargetTexture represents a babylon.js RenderTargetTexture.
// This Helps creating a texture that will be created from a camera in your scene.
// It is basically a dynamic texture that could be used to create special effects for instance.
// Actually, It is the base of lot of effects in the framework like post process, shadows, effect layers and rendering pipelines...
type RenderTargetTexture struct{ *Texture }

// JSObject returns the underlying js.Value.
func (r *RenderTargetTexture) JSObject() js.Value { return r.p }

// RenderTargetTexture returns a RenderTargetTexture JavaScript class.
func (ba *Babylon) RenderTargetTexture() *RenderTargetTexture {
	p := ba.ctx.Get("RenderTargetTexture")
	return RenderTargetTextureFromJSObject(p)
}

// RenderTargetTextureFromJSObject returns a wrapped RenderTargetTexture JavaScript class.
func RenderTargetTextureFromJSObject(p js.Value) *RenderTargetTexture {
	return &RenderTargetTexture{TextureFromJSObject(p)}
}

// NewRenderTargetTextureOpts contains optional parameters for NewRenderTargetTexture.
type NewRenderTargetTextureOpts struct {
	GenerateMipMaps *JSBool

	DoNotChangeAspectRatio *JSBool

	Type *JSFloat64

	IsCube *JSBool

	SamplingMode *JSFloat64

	GenerateDepthBuffer *JSBool

	GenerateStencilBuffer *JSBool

	IsMulti *JSBool

	Format *JSFloat64

	DelayAllocation *JSBool
}

// NewRenderTargetTexture returns a new RenderTargetTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.rendertargettexture
func (ba *Babylon) NewRenderTargetTexture(name string, size float64, scene *Scene, opts *NewRenderTargetTextureOpts) *RenderTargetTexture {
	if opts == nil {
		opts = &NewRenderTargetTextureOpts{}
	}

	p := ba.ctx.Get("RenderTargetTexture").New(name, size, scene.JSObject(), opts.GenerateMipMaps.JSObject(), opts.DoNotChangeAspectRatio.JSObject(), opts.Type.JSObject(), opts.IsCube.JSObject(), opts.SamplingMode.JSObject(), opts.GenerateDepthBuffer.JSObject(), opts.GenerateStencilBuffer.JSObject(), opts.IsMulti.JSObject(), opts.Format.JSObject(), opts.DelayAllocation.JSObject())
	return RenderTargetTextureFromJSObject(p)
}

// TODO: methods
