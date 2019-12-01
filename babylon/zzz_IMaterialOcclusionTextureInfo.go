// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// IMaterialOcclusionTextureInfo represents a babylon.js IMaterialOcclusionTextureInfo.
// Loader interface with additional members.
type IMaterialOcclusionTextureInfo struct {
	*IMaterialOcclusionTextureInfo
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (i *IMaterialOcclusionTextureInfo) JSObject() js.Value { return i.p }

// IMaterialOcclusionTextureInfo returns a IMaterialOcclusionTextureInfo JavaScript class.
func (ba *Babylon) IMaterialOcclusionTextureInfo() *IMaterialOcclusionTextureInfo {
	p := ba.ctx.Get("IMaterialOcclusionTextureInfo")
	return IMaterialOcclusionTextureInfoFromJSObject(p, ba.ctx)
}

// IMaterialOcclusionTextureInfoFromJSObject returns a wrapped IMaterialOcclusionTextureInfo JavaScript class.
func IMaterialOcclusionTextureInfoFromJSObject(p js.Value, ctx js.Value) *IMaterialOcclusionTextureInfo {
	return &IMaterialOcclusionTextureInfo{IMaterialOcclusionTextureInfo: IMaterialOcclusionTextureInfoFromJSObject(p, ctx), ctx: ctx}
}

// IMaterialOcclusionTextureInfoArrayToJSArray returns a JavaScript Array for the wrapped array.
func IMaterialOcclusionTextureInfoArrayToJSArray(array []*IMaterialOcclusionTextureInfo) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

/*

// Extensions returns the Extensions property of class IMaterialOcclusionTextureInfo.
//
// https://doc.babylonjs.com/api/classes/babylon.imaterialocclusiontextureinfo#extensions
func (i *IMaterialOcclusionTextureInfo) Extensions(extensions js.Value) *IMaterialOcclusionTextureInfo {
	p := ba.ctx.Get("IMaterialOcclusionTextureInfo").New(extensions)
	return IMaterialOcclusionTextureInfoFromJSObject(p, ba.ctx)
}

// SetExtensions sets the Extensions property of class IMaterialOcclusionTextureInfo.
//
// https://doc.babylonjs.com/api/classes/babylon.imaterialocclusiontextureinfo#extensions
func (i *IMaterialOcclusionTextureInfo) SetExtensions(extensions js.Value) *IMaterialOcclusionTextureInfo {
	p := ba.ctx.Get("IMaterialOcclusionTextureInfo").New(extensions)
	return IMaterialOcclusionTextureInfoFromJSObject(p, ba.ctx)
}

// Extras returns the Extras property of class IMaterialOcclusionTextureInfo.
//
// https://doc.babylonjs.com/api/classes/babylon.imaterialocclusiontextureinfo#extras
func (i *IMaterialOcclusionTextureInfo) Extras(extras interface{}) *IMaterialOcclusionTextureInfo {
	p := ba.ctx.Get("IMaterialOcclusionTextureInfo").New(extras)
	return IMaterialOcclusionTextureInfoFromJSObject(p, ba.ctx)
}

// SetExtras sets the Extras property of class IMaterialOcclusionTextureInfo.
//
// https://doc.babylonjs.com/api/classes/babylon.imaterialocclusiontextureinfo#extras
func (i *IMaterialOcclusionTextureInfo) SetExtras(extras interface{}) *IMaterialOcclusionTextureInfo {
	p := ba.ctx.Get("IMaterialOcclusionTextureInfo").New(extras)
	return IMaterialOcclusionTextureInfoFromJSObject(p, ba.ctx)
}

// Index returns the Index property of class IMaterialOcclusionTextureInfo.
//
// https://doc.babylonjs.com/api/classes/babylon.imaterialocclusiontextureinfo#index
func (i *IMaterialOcclusionTextureInfo) Index(index float64) *IMaterialOcclusionTextureInfo {
	p := ba.ctx.Get("IMaterialOcclusionTextureInfo").New(index)
	return IMaterialOcclusionTextureInfoFromJSObject(p, ba.ctx)
}

// SetIndex sets the Index property of class IMaterialOcclusionTextureInfo.
//
// https://doc.babylonjs.com/api/classes/babylon.imaterialocclusiontextureinfo#index
func (i *IMaterialOcclusionTextureInfo) SetIndex(index float64) *IMaterialOcclusionTextureInfo {
	p := ba.ctx.Get("IMaterialOcclusionTextureInfo").New(index)
	return IMaterialOcclusionTextureInfoFromJSObject(p, ba.ctx)
}

// Strength returns the Strength property of class IMaterialOcclusionTextureInfo.
//
// https://doc.babylonjs.com/api/classes/babylon.imaterialocclusiontextureinfo#strength
func (i *IMaterialOcclusionTextureInfo) Strength(strength float64) *IMaterialOcclusionTextureInfo {
	p := ba.ctx.Get("IMaterialOcclusionTextureInfo").New(strength)
	return IMaterialOcclusionTextureInfoFromJSObject(p, ba.ctx)
}

// SetStrength sets the Strength property of class IMaterialOcclusionTextureInfo.
//
// https://doc.babylonjs.com/api/classes/babylon.imaterialocclusiontextureinfo#strength
func (i *IMaterialOcclusionTextureInfo) SetStrength(strength float64) *IMaterialOcclusionTextureInfo {
	p := ba.ctx.Get("IMaterialOcclusionTextureInfo").New(strength)
	return IMaterialOcclusionTextureInfoFromJSObject(p, ba.ctx)
}

// TexCoord returns the TexCoord property of class IMaterialOcclusionTextureInfo.
//
// https://doc.babylonjs.com/api/classes/babylon.imaterialocclusiontextureinfo#texcoord
func (i *IMaterialOcclusionTextureInfo) TexCoord(texCoord float64) *IMaterialOcclusionTextureInfo {
	p := ba.ctx.Get("IMaterialOcclusionTextureInfo").New(texCoord)
	return IMaterialOcclusionTextureInfoFromJSObject(p, ba.ctx)
}

// SetTexCoord sets the TexCoord property of class IMaterialOcclusionTextureInfo.
//
// https://doc.babylonjs.com/api/classes/babylon.imaterialocclusiontextureinfo#texcoord
func (i *IMaterialOcclusionTextureInfo) SetTexCoord(texCoord float64) *IMaterialOcclusionTextureInfo {
	p := ba.ctx.Get("IMaterialOcclusionTextureInfo").New(texCoord)
	return IMaterialOcclusionTextureInfoFromJSObject(p, ba.ctx)
}

*/