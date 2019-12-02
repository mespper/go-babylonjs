// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// KHR_materials_sheen represents a babylon.js KHR_materials_sheen.
// <a href="https://github.com/KhronosGroup/glTF/pull/1688">Proposed Specification</a>
// <a href="https://www.babylonjs-playground.com/frame.html#BNIZX6#4">Playground Sample</a>
// !!! Experimental Extension Subject to Changes !!!
type KHR_materials_sheen struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (k *KHR_materials_sheen) JSObject() js.Value { return k.p }

// KHR_materials_sheen returns a KHR_materials_sheen JavaScript class.
func (ba *Babylon) KHR_materials_sheen() *KHR_materials_sheen {
	p := ba.ctx.Get("KHR_materials_sheen")
	return KHR_materials_sheenFromJSObject(p, ba.ctx)
}

// KHR_materials_sheenFromJSObject returns a wrapped KHR_materials_sheen JavaScript class.
func KHR_materials_sheenFromJSObject(p js.Value, ctx js.Value) *KHR_materials_sheen {
	return &KHR_materials_sheen{p: p, ctx: ctx}
}

// KHR_materials_sheenArrayToJSArray returns a JavaScript Array for the wrapped array.
func KHR_materials_sheenArrayToJSArray(array []*KHR_materials_sheen) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// Enabled returns the Enabled property of class KHR_materials_sheen.
//
// https://doc.babylonjs.com/api/classes/babylon.khr_materials_sheen#enabled
func (k *KHR_materials_sheen) Enabled() bool {
	retVal := k.p.Get("enabled")
	return retVal.Bool()
}

// SetEnabled sets the Enabled property of class KHR_materials_sheen.
//
// https://doc.babylonjs.com/api/classes/babylon.khr_materials_sheen#enabled
func (k *KHR_materials_sheen) SetEnabled(enabled bool) *KHR_materials_sheen {
	k.p.Set("enabled", enabled)
	return k
}

// Name returns the Name property of class KHR_materials_sheen.
//
// https://doc.babylonjs.com/api/classes/babylon.khr_materials_sheen#name
func (k *KHR_materials_sheen) Name() string {
	retVal := k.p.Get("name")
	return retVal.String()
}

// SetName sets the Name property of class KHR_materials_sheen.
//
// https://doc.babylonjs.com/api/classes/babylon.khr_materials_sheen#name
func (k *KHR_materials_sheen) SetName(name string) *KHR_materials_sheen {
	k.p.Set("name", name)
	return k
}

// Order returns the Order property of class KHR_materials_sheen.
//
// https://doc.babylonjs.com/api/classes/babylon.khr_materials_sheen#order
func (k *KHR_materials_sheen) Order() float64 {
	retVal := k.p.Get("order")
	return retVal.Float()
}

// SetOrder sets the Order property of class KHR_materials_sheen.
//
// https://doc.babylonjs.com/api/classes/babylon.khr_materials_sheen#order
func (k *KHR_materials_sheen) SetOrder(order float64) *KHR_materials_sheen {
	k.p.Set("order", order)
	return k
}
