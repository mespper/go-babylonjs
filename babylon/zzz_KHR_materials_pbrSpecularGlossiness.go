// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// KHR_materials_pbrSpecularGlossiness represents a babylon.js KHR_materials_pbrSpecularGlossiness.
// <a href="https://github.com/KhronosGroup/glTF/tree/master/extensions/2.0/Khronos/KHR_materials_pbrSpecularGlossiness">Specification</a>
type KHR_materials_pbrSpecularGlossiness struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (k *KHR_materials_pbrSpecularGlossiness) JSObject() js.Value { return k.p }

// KHR_materials_pbrSpecularGlossiness returns a KHR_materials_pbrSpecularGlossiness JavaScript class.
func (ba *Babylon) KHR_materials_pbrSpecularGlossiness() *KHR_materials_pbrSpecularGlossiness {
	p := ba.ctx.Get("KHR_materials_pbrSpecularGlossiness")
	return KHR_materials_pbrSpecularGlossinessFromJSObject(p, ba.ctx)
}

// KHR_materials_pbrSpecularGlossinessFromJSObject returns a wrapped KHR_materials_pbrSpecularGlossiness JavaScript class.
func KHR_materials_pbrSpecularGlossinessFromJSObject(p js.Value, ctx js.Value) *KHR_materials_pbrSpecularGlossiness {
	return &KHR_materials_pbrSpecularGlossiness{p: p, ctx: ctx}
}

// KHR_materials_pbrSpecularGlossinessArrayToJSArray returns a JavaScript Array for the wrapped array.
func KHR_materials_pbrSpecularGlossinessArrayToJSArray(array []*KHR_materials_pbrSpecularGlossiness) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// Enabled returns the Enabled property of class KHR_materials_pbrSpecularGlossiness.
//
// https://doc.babylonjs.com/api/classes/babylon.khr_materials_pbrspecularglossiness#enabled
func (k *KHR_materials_pbrSpecularGlossiness) Enabled() bool {
	retVal := k.p.Get("enabled")
	return retVal.Bool()
}

// SetEnabled sets the Enabled property of class KHR_materials_pbrSpecularGlossiness.
//
// https://doc.babylonjs.com/api/classes/babylon.khr_materials_pbrspecularglossiness#enabled
func (k *KHR_materials_pbrSpecularGlossiness) SetEnabled(enabled bool) *KHR_materials_pbrSpecularGlossiness {
	k.p.Set("enabled", enabled)
	return k
}

// Name returns the Name property of class KHR_materials_pbrSpecularGlossiness.
//
// https://doc.babylonjs.com/api/classes/babylon.khr_materials_pbrspecularglossiness#name
func (k *KHR_materials_pbrSpecularGlossiness) Name() string {
	retVal := k.p.Get("name")
	return retVal.String()
}

// SetName sets the Name property of class KHR_materials_pbrSpecularGlossiness.
//
// https://doc.babylonjs.com/api/classes/babylon.khr_materials_pbrspecularglossiness#name
func (k *KHR_materials_pbrSpecularGlossiness) SetName(name string) *KHR_materials_pbrSpecularGlossiness {
	k.p.Set("name", name)
	return k
}

// Order returns the Order property of class KHR_materials_pbrSpecularGlossiness.
//
// https://doc.babylonjs.com/api/classes/babylon.khr_materials_pbrspecularglossiness#order
func (k *KHR_materials_pbrSpecularGlossiness) Order() float64 {
	retVal := k.p.Get("order")
	return retVal.Float()
}

// SetOrder sets the Order property of class KHR_materials_pbrSpecularGlossiness.
//
// https://doc.babylonjs.com/api/classes/babylon.khr_materials_pbrspecularglossiness#order
func (k *KHR_materials_pbrSpecularGlossiness) SetOrder(order float64) *KHR_materials_pbrSpecularGlossiness {
	k.p.Set("order", order)
	return k
}
