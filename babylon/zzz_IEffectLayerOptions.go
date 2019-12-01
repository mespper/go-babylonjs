// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// IEffectLayerOptions represents a babylon.js IEffectLayerOptions.
// Effect layer options. This helps customizing the behaviour
// of the effect layer.
type IEffectLayerOptions struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (i *IEffectLayerOptions) JSObject() js.Value { return i.p }

// IEffectLayerOptions returns a IEffectLayerOptions JavaScript class.
func (ba *Babylon) IEffectLayerOptions() *IEffectLayerOptions {
	p := ba.ctx.Get("IEffectLayerOptions")
	return IEffectLayerOptionsFromJSObject(p, ba.ctx)
}

// IEffectLayerOptionsFromJSObject returns a wrapped IEffectLayerOptions JavaScript class.
func IEffectLayerOptionsFromJSObject(p js.Value, ctx js.Value) *IEffectLayerOptions {
	return &IEffectLayerOptions{p: p, ctx: ctx}
}

// IEffectLayerOptionsArrayToJSArray returns a JavaScript Array for the wrapped array.
func IEffectLayerOptionsArrayToJSArray(array []*IEffectLayerOptions) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

/*

// AlphaBlendingMode returns the AlphaBlendingMode property of class IEffectLayerOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.ieffectlayeroptions#alphablendingmode
func (i *IEffectLayerOptions) AlphaBlendingMode(alphaBlendingMode float64) *IEffectLayerOptions {
	p := ba.ctx.Get("IEffectLayerOptions").New(alphaBlendingMode)
	return IEffectLayerOptionsFromJSObject(p, ba.ctx)
}

// SetAlphaBlendingMode sets the AlphaBlendingMode property of class IEffectLayerOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.ieffectlayeroptions#alphablendingmode
func (i *IEffectLayerOptions) SetAlphaBlendingMode(alphaBlendingMode float64) *IEffectLayerOptions {
	p := ba.ctx.Get("IEffectLayerOptions").New(alphaBlendingMode)
	return IEffectLayerOptionsFromJSObject(p, ba.ctx)
}

// Camera returns the Camera property of class IEffectLayerOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.ieffectlayeroptions#camera
func (i *IEffectLayerOptions) Camera(camera *Camera) *IEffectLayerOptions {
	p := ba.ctx.Get("IEffectLayerOptions").New(camera.JSObject())
	return IEffectLayerOptionsFromJSObject(p, ba.ctx)
}

// SetCamera sets the Camera property of class IEffectLayerOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.ieffectlayeroptions#camera
func (i *IEffectLayerOptions) SetCamera(camera *Camera) *IEffectLayerOptions {
	p := ba.ctx.Get("IEffectLayerOptions").New(camera.JSObject())
	return IEffectLayerOptionsFromJSObject(p, ba.ctx)
}

// MainTextureFixedSize returns the MainTextureFixedSize property of class IEffectLayerOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.ieffectlayeroptions#maintexturefixedsize
func (i *IEffectLayerOptions) MainTextureFixedSize(mainTextureFixedSize float64) *IEffectLayerOptions {
	p := ba.ctx.Get("IEffectLayerOptions").New(mainTextureFixedSize)
	return IEffectLayerOptionsFromJSObject(p, ba.ctx)
}

// SetMainTextureFixedSize sets the MainTextureFixedSize property of class IEffectLayerOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.ieffectlayeroptions#maintexturefixedsize
func (i *IEffectLayerOptions) SetMainTextureFixedSize(mainTextureFixedSize float64) *IEffectLayerOptions {
	p := ba.ctx.Get("IEffectLayerOptions").New(mainTextureFixedSize)
	return IEffectLayerOptionsFromJSObject(p, ba.ctx)
}

// MainTextureRatio returns the MainTextureRatio property of class IEffectLayerOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.ieffectlayeroptions#maintextureratio
func (i *IEffectLayerOptions) MainTextureRatio(mainTextureRatio float64) *IEffectLayerOptions {
	p := ba.ctx.Get("IEffectLayerOptions").New(mainTextureRatio)
	return IEffectLayerOptionsFromJSObject(p, ba.ctx)
}

// SetMainTextureRatio sets the MainTextureRatio property of class IEffectLayerOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.ieffectlayeroptions#maintextureratio
func (i *IEffectLayerOptions) SetMainTextureRatio(mainTextureRatio float64) *IEffectLayerOptions {
	p := ba.ctx.Get("IEffectLayerOptions").New(mainTextureRatio)
	return IEffectLayerOptionsFromJSObject(p, ba.ctx)
}

// RenderingGroupId returns the RenderingGroupId property of class IEffectLayerOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.ieffectlayeroptions#renderinggroupid
func (i *IEffectLayerOptions) RenderingGroupId(renderingGroupId float64) *IEffectLayerOptions {
	p := ba.ctx.Get("IEffectLayerOptions").New(renderingGroupId)
	return IEffectLayerOptionsFromJSObject(p, ba.ctx)
}

// SetRenderingGroupId sets the RenderingGroupId property of class IEffectLayerOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.ieffectlayeroptions#renderinggroupid
func (i *IEffectLayerOptions) SetRenderingGroupId(renderingGroupId float64) *IEffectLayerOptions {
	p := ba.ctx.Get("IEffectLayerOptions").New(renderingGroupId)
	return IEffectLayerOptionsFromJSObject(p, ba.ctx)
}

*/