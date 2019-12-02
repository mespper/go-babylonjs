// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// FireProceduralTexture represents a babylon.js FireProceduralTexture.
//
type FireProceduralTexture struct {
	*ProceduralTexture
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (f *FireProceduralTexture) JSObject() js.Value { return f.p }

// FireProceduralTexture returns a FireProceduralTexture JavaScript class.
func (ba *Babylon) FireProceduralTexture() *FireProceduralTexture {
	p := ba.ctx.Get("FireProceduralTexture")
	return FireProceduralTextureFromJSObject(p, ba.ctx)
}

// FireProceduralTextureFromJSObject returns a wrapped FireProceduralTexture JavaScript class.
func FireProceduralTextureFromJSObject(p js.Value, ctx js.Value) *FireProceduralTexture {
	return &FireProceduralTexture{ProceduralTexture: ProceduralTextureFromJSObject(p, ctx), ctx: ctx}
}

// FireProceduralTextureArrayToJSArray returns a JavaScript Array for the wrapped array.
func FireProceduralTextureArrayToJSArray(array []*FireProceduralTexture) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewFireProceduralTextureOpts contains optional parameters for NewFireProceduralTexture.
type NewFireProceduralTextureOpts struct {
	FallbackTexture *Texture
	GenerateMipMaps *bool
}

// NewFireProceduralTexture returns a new FireProceduralTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.fireproceduraltexture
func (ba *Babylon) NewFireProceduralTexture(name string, size float64, scene *Scene, opts *NewFireProceduralTextureOpts) *FireProceduralTexture {
	if opts == nil {
		opts = &NewFireProceduralTextureOpts{}
	}

	args := make([]interface{}, 0, 3+2)

	args = append(args, name)
	args = append(args, size)
	args = append(args, scene.JSObject())

	if opts.FallbackTexture == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.FallbackTexture.JSObject())
	}
	if opts.GenerateMipMaps == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.GenerateMipMaps)
	}

	p := ba.ctx.Get("FireProceduralTexture").New(args...)
	return FireProceduralTextureFromJSObject(p, ba.ctx)
}

// Parse calls the Parse method on the FireProceduralTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.fireproceduraltexture#parse
func (f *FireProceduralTexture) Parse(parsedTexture interface{}, scene *Scene, rootUrl string) *FireProceduralTexture {

	args := make([]interface{}, 0, 3+0)

	args = append(args, parsedTexture)
	args = append(args, scene.JSObject())
	args = append(args, rootUrl)

	retVal := f.p.Call("Parse", args...)
	return FireProceduralTextureFromJSObject(retVal, f.ctx)
}

// FireProceduralTextureRenderOpts contains optional parameters for FireProceduralTexture.Render.
type FireProceduralTextureRenderOpts struct {
	UseCameraPostProcess *bool
}

// Render calls the Render method on the FireProceduralTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.fireproceduraltexture#render
func (f *FireProceduralTexture) Render(opts *FireProceduralTextureRenderOpts) {
	if opts == nil {
		opts = &FireProceduralTextureRenderOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.UseCameraPostProcess == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.UseCameraPostProcess)
	}

	f.p.Call("render", args...)
}

// Serialize calls the Serialize method on the FireProceduralTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.fireproceduraltexture#serialize
func (f *FireProceduralTexture) Serialize() interface{} {

	retVal := f.p.Call("serialize")
	return retVal
}

// UpdateShaderUniforms calls the UpdateShaderUniforms method on the FireProceduralTexture object.
//
// https://doc.babylonjs.com/api/classes/babylon.fireproceduraltexture#updateshaderuniforms
func (f *FireProceduralTexture) UpdateShaderUniforms() {

	f.p.Call("updateShaderUniforms")
}

// AlphaThreshold returns the AlphaThreshold property of class FireProceduralTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.fireproceduraltexture#alphathreshold
func (f *FireProceduralTexture) AlphaThreshold() float64 {
	retVal := f.p.Get("alphaThreshold")
	return retVal.Float()
}

// SetAlphaThreshold sets the AlphaThreshold property of class FireProceduralTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.fireproceduraltexture#alphathreshold
func (f *FireProceduralTexture) SetAlphaThreshold(alphaThreshold float64) *FireProceduralTexture {
	f.p.Set("alphaThreshold", alphaThreshold)
	return f
}

// AutoGenerateTime returns the AutoGenerateTime property of class FireProceduralTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.fireproceduraltexture#autogeneratetime
func (f *FireProceduralTexture) AutoGenerateTime() bool {
	retVal := f.p.Get("autoGenerateTime")
	return retVal.Bool()
}

// SetAutoGenerateTime sets the AutoGenerateTime property of class FireProceduralTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.fireproceduraltexture#autogeneratetime
func (f *FireProceduralTexture) SetAutoGenerateTime(autoGenerateTime bool) *FireProceduralTexture {
	f.p.Set("autoGenerateTime", autoGenerateTime)
	return f
}

// BlueFireColors returns the BlueFireColors property of class FireProceduralTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.fireproceduraltexture#bluefirecolors
func (f *FireProceduralTexture) BlueFireColors() *Color3 {
	retVal := f.p.Get("BlueFireColors")
	return Color3FromJSObject(retVal, f.ctx)
}

// SetBlueFireColors sets the BlueFireColors property of class FireProceduralTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.fireproceduraltexture#bluefirecolors
func (f *FireProceduralTexture) SetBlueFireColors(BlueFireColors *Color3) *FireProceduralTexture {
	f.p.Set("BlueFireColors", BlueFireColors.JSObject())
	return f
}

// FireColors returns the FireColors property of class FireProceduralTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.fireproceduraltexture#firecolors
func (f *FireProceduralTexture) FireColors() *Color3 {
	retVal := f.p.Get("fireColors")
	return Color3FromJSObject(retVal, f.ctx)
}

// SetFireColors sets the FireColors property of class FireProceduralTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.fireproceduraltexture#firecolors
func (f *FireProceduralTexture) SetFireColors(fireColors *Color3) *FireProceduralTexture {
	f.p.Set("fireColors", fireColors.JSObject())
	return f
}

// GreenFireColors returns the GreenFireColors property of class FireProceduralTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.fireproceduraltexture#greenfirecolors
func (f *FireProceduralTexture) GreenFireColors() *Color3 {
	retVal := f.p.Get("GreenFireColors")
	return Color3FromJSObject(retVal, f.ctx)
}

// SetGreenFireColors sets the GreenFireColors property of class FireProceduralTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.fireproceduraltexture#greenfirecolors
func (f *FireProceduralTexture) SetGreenFireColors(GreenFireColors *Color3) *FireProceduralTexture {
	f.p.Set("GreenFireColors", GreenFireColors.JSObject())
	return f
}

// PurpleFireColors returns the PurpleFireColors property of class FireProceduralTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.fireproceduraltexture#purplefirecolors
func (f *FireProceduralTexture) PurpleFireColors() *Color3 {
	retVal := f.p.Get("PurpleFireColors")
	return Color3FromJSObject(retVal, f.ctx)
}

// SetPurpleFireColors sets the PurpleFireColors property of class FireProceduralTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.fireproceduraltexture#purplefirecolors
func (f *FireProceduralTexture) SetPurpleFireColors(PurpleFireColors *Color3) *FireProceduralTexture {
	f.p.Set("PurpleFireColors", PurpleFireColors.JSObject())
	return f
}

// RedFireColors returns the RedFireColors property of class FireProceduralTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.fireproceduraltexture#redfirecolors
func (f *FireProceduralTexture) RedFireColors() *Color3 {
	retVal := f.p.Get("RedFireColors")
	return Color3FromJSObject(retVal, f.ctx)
}

// SetRedFireColors sets the RedFireColors property of class FireProceduralTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.fireproceduraltexture#redfirecolors
func (f *FireProceduralTexture) SetRedFireColors(RedFireColors *Color3) *FireProceduralTexture {
	f.p.Set("RedFireColors", RedFireColors.JSObject())
	return f
}

// Speed returns the Speed property of class FireProceduralTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.fireproceduraltexture#speed
func (f *FireProceduralTexture) Speed() *Vector2 {
	retVal := f.p.Get("speed")
	return Vector2FromJSObject(retVal, f.ctx)
}

// SetSpeed sets the Speed property of class FireProceduralTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.fireproceduraltexture#speed
func (f *FireProceduralTexture) SetSpeed(speed *Vector2) *FireProceduralTexture {
	f.p.Set("speed", speed.JSObject())
	return f
}

// Time returns the Time property of class FireProceduralTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.fireproceduraltexture#time
func (f *FireProceduralTexture) Time() float64 {
	retVal := f.p.Get("time")
	return retVal.Float()
}

// SetTime sets the Time property of class FireProceduralTexture.
//
// https://doc.babylonjs.com/api/classes/babylon.fireproceduraltexture#time
func (f *FireProceduralTexture) SetTime(time float64) *FireProceduralTexture {
	f.p.Set("time", time)
	return f
}
