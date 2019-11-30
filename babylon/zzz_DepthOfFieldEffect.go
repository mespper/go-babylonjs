// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// DepthOfFieldEffect represents a babylon.js DepthOfFieldEffect.
// The depth of field effect applies a blur to objects that are closer or further from where the camera is focusing.
type DepthOfFieldEffect struct{ *PostProcessRenderEffect }

// JSObject returns the underlying js.Value.
func (d *DepthOfFieldEffect) JSObject() js.Value { return d.p }

// DepthOfFieldEffect returns a DepthOfFieldEffect JavaScript class.
func (ba *Babylon) DepthOfFieldEffect() *DepthOfFieldEffect {
	p := ba.ctx.Get("DepthOfFieldEffect")
	return DepthOfFieldEffectFromJSObject(p)
}

// DepthOfFieldEffectFromJSObject returns a wrapped DepthOfFieldEffect JavaScript class.
func DepthOfFieldEffectFromJSObject(p js.Value) *DepthOfFieldEffect {
	return &DepthOfFieldEffect{PostProcessRenderEffectFromJSObject(p)}
}

// NewDepthOfFieldEffectOpts contains optional parameters for NewDepthOfFieldEffect.
type NewDepthOfFieldEffectOpts struct {
	BlurLevel *DepthOfFieldEffectBlurLevel

	PipelineTextureType *JSFloat64

	BlockCompilation *JSBool
}

// NewDepthOfFieldEffect returns a new DepthOfFieldEffect object.
//
// https://doc.babylonjs.com/api/classes/babylon.depthoffieldeffect
func (ba *Babylon) NewDepthOfFieldEffect(scene *Scene, depthTexture *RenderTargetTexture, opts *NewDepthOfFieldEffectOpts) *DepthOfFieldEffect {
	if opts == nil {
		opts = &NewDepthOfFieldEffectOpts{}
	}

	p := ba.ctx.Get("DepthOfFieldEffect").New(scene.JSObject(), depthTexture.JSObject(), opts.BlurLevel.JSObject(), opts.PipelineTextureType.JSObject(), opts.BlockCompilation.JSObject())
	return DepthOfFieldEffectFromJSObject(p)
}

// TODO: methods
