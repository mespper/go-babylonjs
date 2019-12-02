// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// IValueGradient represents a babylon.js IValueGradient.
// Interface used by value gradients (color, factor, ...)
type IValueGradient struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (i *IValueGradient) JSObject() js.Value { return i.p }

// IValueGradient returns a IValueGradient JavaScript class.
func (ba *Babylon) IValueGradient() *IValueGradient {
	p := ba.ctx.Get("IValueGradient")
	return IValueGradientFromJSObject(p, ba.ctx)
}

// IValueGradientFromJSObject returns a wrapped IValueGradient JavaScript class.
func IValueGradientFromJSObject(p js.Value, ctx js.Value) *IValueGradient {
	return &IValueGradient{p: p, ctx: ctx}
}

// IValueGradientArrayToJSArray returns a JavaScript Array for the wrapped array.
func IValueGradientArrayToJSArray(array []*IValueGradient) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// Gradient returns the Gradient property of class IValueGradient.
//
// https://doc.babylonjs.com/api/classes/babylon.ivaluegradient#gradient
func (i *IValueGradient) Gradient() float64 {
	retVal := i.p.Get("gradient")
	return retVal.Float()
}

// SetGradient sets the Gradient property of class IValueGradient.
//
// https://doc.babylonjs.com/api/classes/babylon.ivaluegradient#gradient
func (i *IValueGradient) SetGradient(gradient float64) *IValueGradient {
	i.p.Set("gradient", gradient)
	return i
}
