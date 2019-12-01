// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// GradientHelper represents a babylon.js GradientHelper.
// Helper used to simplify some generic gradient tasks
type GradientHelper struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (g *GradientHelper) JSObject() js.Value { return g.p }

// GradientHelper returns a GradientHelper JavaScript class.
func (ba *Babylon) GradientHelper() *GradientHelper {
	p := ba.ctx.Get("GradientHelper")
	return GradientHelperFromJSObject(p, ba.ctx)
}

// GradientHelperFromJSObject returns a wrapped GradientHelper JavaScript class.
func GradientHelperFromJSObject(p js.Value, ctx js.Value) *GradientHelper {
	return &GradientHelper{p: p, ctx: ctx}
}

// GradientHelperArrayToJSArray returns a JavaScript Array for the wrapped array.
func GradientHelperArrayToJSArray(array []*GradientHelper) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// GetCurrentGradient calls the GetCurrentGradient method on the GradientHelper object.
//
// https://doc.babylonjs.com/api/classes/babylon.gradienthelper#getcurrentgradient
func (g *GradientHelper) GetCurrentGradient(ratio float64, gradients *IValueGradient, updateFunc func()) {

	args := make([]interface{}, 0, 3+0)

	args = append(args, ratio)
	args = append(args, gradients.JSObject())
	args = append(args, js.FuncOf(func(this js.Value, args []js.Value) interface{} { updateFunc(); return nil }))

	g.p.Call("GetCurrentGradient", args...)
}

/*

 */
