// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// IMaterialPbrMetallicRoughness represents a babylon.js IMaterialPbrMetallicRoughness.
// Loader interface with additional members.
type IMaterialPbrMetallicRoughness struct {
	*IMaterialPbrMetallicRoughness
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (i *IMaterialPbrMetallicRoughness) JSObject() js.Value { return i.p }

// IMaterialPbrMetallicRoughness returns a IMaterialPbrMetallicRoughness JavaScript class.
func (ba *Babylon) IMaterialPbrMetallicRoughness() *IMaterialPbrMetallicRoughness {
	p := ba.ctx.Get("IMaterialPbrMetallicRoughness")
	return IMaterialPbrMetallicRoughnessFromJSObject(p, ba.ctx)
}

// IMaterialPbrMetallicRoughnessFromJSObject returns a wrapped IMaterialPbrMetallicRoughness JavaScript class.
func IMaterialPbrMetallicRoughnessFromJSObject(p js.Value, ctx js.Value) *IMaterialPbrMetallicRoughness {
	return &IMaterialPbrMetallicRoughness{IMaterialPbrMetallicRoughness: IMaterialPbrMetallicRoughnessFromJSObject(p, ctx), ctx: ctx}
}

// IMaterialPbrMetallicRoughnessArrayToJSArray returns a JavaScript Array for the wrapped array.
func IMaterialPbrMetallicRoughnessArrayToJSArray(array []*IMaterialPbrMetallicRoughness) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

/*

// BaseColorFactor returns the BaseColorFactor property of class IMaterialPbrMetallicRoughness.
//
// https://doc.babylonjs.com/api/classes/babylon.imaterialpbrmetallicroughness#basecolorfactor
func (i *IMaterialPbrMetallicRoughness) BaseColorFactor(baseColorFactor float64) *IMaterialPbrMetallicRoughness {
	p := ba.ctx.Get("IMaterialPbrMetallicRoughness").New(baseColorFactor)
	return IMaterialPbrMetallicRoughnessFromJSObject(p, ba.ctx)
}

// SetBaseColorFactor sets the BaseColorFactor property of class IMaterialPbrMetallicRoughness.
//
// https://doc.babylonjs.com/api/classes/babylon.imaterialpbrmetallicroughness#basecolorfactor
func (i *IMaterialPbrMetallicRoughness) SetBaseColorFactor(baseColorFactor float64) *IMaterialPbrMetallicRoughness {
	p := ba.ctx.Get("IMaterialPbrMetallicRoughness").New(baseColorFactor)
	return IMaterialPbrMetallicRoughnessFromJSObject(p, ba.ctx)
}

// BaseColorTexture returns the BaseColorTexture property of class IMaterialPbrMetallicRoughness.
//
// https://doc.babylonjs.com/api/classes/babylon.imaterialpbrmetallicroughness#basecolortexture
func (i *IMaterialPbrMetallicRoughness) BaseColorTexture(baseColorTexture *ITextureInfo) *IMaterialPbrMetallicRoughness {
	p := ba.ctx.Get("IMaterialPbrMetallicRoughness").New(baseColorTexture.JSObject())
	return IMaterialPbrMetallicRoughnessFromJSObject(p, ba.ctx)
}

// SetBaseColorTexture sets the BaseColorTexture property of class IMaterialPbrMetallicRoughness.
//
// https://doc.babylonjs.com/api/classes/babylon.imaterialpbrmetallicroughness#basecolortexture
func (i *IMaterialPbrMetallicRoughness) SetBaseColorTexture(baseColorTexture *ITextureInfo) *IMaterialPbrMetallicRoughness {
	p := ba.ctx.Get("IMaterialPbrMetallicRoughness").New(baseColorTexture.JSObject())
	return IMaterialPbrMetallicRoughnessFromJSObject(p, ba.ctx)
}

// MetallicFactor returns the MetallicFactor property of class IMaterialPbrMetallicRoughness.
//
// https://doc.babylonjs.com/api/classes/babylon.imaterialpbrmetallicroughness#metallicfactor
func (i *IMaterialPbrMetallicRoughness) MetallicFactor(metallicFactor float64) *IMaterialPbrMetallicRoughness {
	p := ba.ctx.Get("IMaterialPbrMetallicRoughness").New(metallicFactor)
	return IMaterialPbrMetallicRoughnessFromJSObject(p, ba.ctx)
}

// SetMetallicFactor sets the MetallicFactor property of class IMaterialPbrMetallicRoughness.
//
// https://doc.babylonjs.com/api/classes/babylon.imaterialpbrmetallicroughness#metallicfactor
func (i *IMaterialPbrMetallicRoughness) SetMetallicFactor(metallicFactor float64) *IMaterialPbrMetallicRoughness {
	p := ba.ctx.Get("IMaterialPbrMetallicRoughness").New(metallicFactor)
	return IMaterialPbrMetallicRoughnessFromJSObject(p, ba.ctx)
}

// MetallicRoughnessTexture returns the MetallicRoughnessTexture property of class IMaterialPbrMetallicRoughness.
//
// https://doc.babylonjs.com/api/classes/babylon.imaterialpbrmetallicroughness#metallicroughnesstexture
func (i *IMaterialPbrMetallicRoughness) MetallicRoughnessTexture(metallicRoughnessTexture *ITextureInfo) *IMaterialPbrMetallicRoughness {
	p := ba.ctx.Get("IMaterialPbrMetallicRoughness").New(metallicRoughnessTexture.JSObject())
	return IMaterialPbrMetallicRoughnessFromJSObject(p, ba.ctx)
}

// SetMetallicRoughnessTexture sets the MetallicRoughnessTexture property of class IMaterialPbrMetallicRoughness.
//
// https://doc.babylonjs.com/api/classes/babylon.imaterialpbrmetallicroughness#metallicroughnesstexture
func (i *IMaterialPbrMetallicRoughness) SetMetallicRoughnessTexture(metallicRoughnessTexture *ITextureInfo) *IMaterialPbrMetallicRoughness {
	p := ba.ctx.Get("IMaterialPbrMetallicRoughness").New(metallicRoughnessTexture.JSObject())
	return IMaterialPbrMetallicRoughnessFromJSObject(p, ba.ctx)
}

// RoughnessFactor returns the RoughnessFactor property of class IMaterialPbrMetallicRoughness.
//
// https://doc.babylonjs.com/api/classes/babylon.imaterialpbrmetallicroughness#roughnessfactor
func (i *IMaterialPbrMetallicRoughness) RoughnessFactor(roughnessFactor float64) *IMaterialPbrMetallicRoughness {
	p := ba.ctx.Get("IMaterialPbrMetallicRoughness").New(roughnessFactor)
	return IMaterialPbrMetallicRoughnessFromJSObject(p, ba.ctx)
}

// SetRoughnessFactor sets the RoughnessFactor property of class IMaterialPbrMetallicRoughness.
//
// https://doc.babylonjs.com/api/classes/babylon.imaterialpbrmetallicroughness#roughnessfactor
func (i *IMaterialPbrMetallicRoughness) SetRoughnessFactor(roughnessFactor float64) *IMaterialPbrMetallicRoughness {
	p := ba.ctx.Get("IMaterialPbrMetallicRoughness").New(roughnessFactor)
	return IMaterialPbrMetallicRoughnessFromJSObject(p, ba.ctx)
}

*/
