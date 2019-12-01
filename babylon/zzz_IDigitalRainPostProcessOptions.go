// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// IDigitalRainPostProcessOptions represents a babylon.js IDigitalRainPostProcessOptions.
// Option available in the Digital Rain Post Process.
type IDigitalRainPostProcessOptions struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (i *IDigitalRainPostProcessOptions) JSObject() js.Value { return i.p }

// IDigitalRainPostProcessOptions returns a IDigitalRainPostProcessOptions JavaScript class.
func (ba *Babylon) IDigitalRainPostProcessOptions() *IDigitalRainPostProcessOptions {
	p := ba.ctx.Get("IDigitalRainPostProcessOptions")
	return IDigitalRainPostProcessOptionsFromJSObject(p, ba.ctx)
}

// IDigitalRainPostProcessOptionsFromJSObject returns a wrapped IDigitalRainPostProcessOptions JavaScript class.
func IDigitalRainPostProcessOptionsFromJSObject(p js.Value, ctx js.Value) *IDigitalRainPostProcessOptions {
	return &IDigitalRainPostProcessOptions{p: p, ctx: ctx}
}

// IDigitalRainPostProcessOptionsArrayToJSArray returns a JavaScript Array for the wrapped array.
func IDigitalRainPostProcessOptionsArrayToJSArray(array []*IDigitalRainPostProcessOptions) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

/*

// Font returns the Font property of class IDigitalRainPostProcessOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.idigitalrainpostprocessoptions#font
func (i *IDigitalRainPostProcessOptions) Font(font string) *IDigitalRainPostProcessOptions {
	p := ba.ctx.Get("IDigitalRainPostProcessOptions").New(font)
	return IDigitalRainPostProcessOptionsFromJSObject(p, ba.ctx)
}

// SetFont sets the Font property of class IDigitalRainPostProcessOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.idigitalrainpostprocessoptions#font
func (i *IDigitalRainPostProcessOptions) SetFont(font string) *IDigitalRainPostProcessOptions {
	p := ba.ctx.Get("IDigitalRainPostProcessOptions").New(font)
	return IDigitalRainPostProcessOptionsFromJSObject(p, ba.ctx)
}

// MixToNormal returns the MixToNormal property of class IDigitalRainPostProcessOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.idigitalrainpostprocessoptions#mixtonormal
func (i *IDigitalRainPostProcessOptions) MixToNormal(mixToNormal float64) *IDigitalRainPostProcessOptions {
	p := ba.ctx.Get("IDigitalRainPostProcessOptions").New(mixToNormal)
	return IDigitalRainPostProcessOptionsFromJSObject(p, ba.ctx)
}

// SetMixToNormal sets the MixToNormal property of class IDigitalRainPostProcessOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.idigitalrainpostprocessoptions#mixtonormal
func (i *IDigitalRainPostProcessOptions) SetMixToNormal(mixToNormal float64) *IDigitalRainPostProcessOptions {
	p := ba.ctx.Get("IDigitalRainPostProcessOptions").New(mixToNormal)
	return IDigitalRainPostProcessOptionsFromJSObject(p, ba.ctx)
}

// MixToTile returns the MixToTile property of class IDigitalRainPostProcessOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.idigitalrainpostprocessoptions#mixtotile
func (i *IDigitalRainPostProcessOptions) MixToTile(mixToTile float64) *IDigitalRainPostProcessOptions {
	p := ba.ctx.Get("IDigitalRainPostProcessOptions").New(mixToTile)
	return IDigitalRainPostProcessOptionsFromJSObject(p, ba.ctx)
}

// SetMixToTile sets the MixToTile property of class IDigitalRainPostProcessOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.idigitalrainpostprocessoptions#mixtotile
func (i *IDigitalRainPostProcessOptions) SetMixToTile(mixToTile float64) *IDigitalRainPostProcessOptions {
	p := ba.ctx.Get("IDigitalRainPostProcessOptions").New(mixToTile)
	return IDigitalRainPostProcessOptionsFromJSObject(p, ba.ctx)
}

*/
