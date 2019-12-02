// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// EffectFallbacks represents a babylon.js EffectFallbacks.
// EffectFallbacks can be used to add fallbacks (properties to disable) to certain properties when desired to improve performance.
// (Eg. Start at high quality with reflection and fog, if fps is low, remove reflection, if still low remove fog)
type EffectFallbacks struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (e *EffectFallbacks) JSObject() js.Value { return e.p }

// EffectFallbacks returns a EffectFallbacks JavaScript class.
func (ba *Babylon) EffectFallbacks() *EffectFallbacks {
	p := ba.ctx.Get("EffectFallbacks")
	return EffectFallbacksFromJSObject(p, ba.ctx)
}

// EffectFallbacksFromJSObject returns a wrapped EffectFallbacks JavaScript class.
func EffectFallbacksFromJSObject(p js.Value, ctx js.Value) *EffectFallbacks {
	return &EffectFallbacks{p: p, ctx: ctx}
}

// EffectFallbacksArrayToJSArray returns a JavaScript Array for the wrapped array.
func EffectFallbacksArrayToJSArray(array []*EffectFallbacks) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// AddCPUSkinningFallback calls the AddCPUSkinningFallback method on the EffectFallbacks object.
//
// https://doc.babylonjs.com/api/classes/babylon.effectfallbacks#addcpuskinningfallback
func (e *EffectFallbacks) AddCPUSkinningFallback(rank float64, mesh *AbstractMesh) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, rank)
	args = append(args, mesh.JSObject())

	e.p.Call("addCPUSkinningFallback", args...)
}

// AddFallback calls the AddFallback method on the EffectFallbacks object.
//
// https://doc.babylonjs.com/api/classes/babylon.effectfallbacks#addfallback
func (e *EffectFallbacks) AddFallback(rank float64, define string) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, rank)
	args = append(args, define)

	e.p.Call("addFallback", args...)
}

// Reduce calls the Reduce method on the EffectFallbacks object.
//
// https://doc.babylonjs.com/api/classes/babylon.effectfallbacks#reduce
func (e *EffectFallbacks) Reduce(currentDefines string, effect *Effect) string {

	args := make([]interface{}, 0, 2+0)

	args = append(args, currentDefines)
	args = append(args, effect.JSObject())

	retVal := e.p.Call("reduce", args...)
	return retVal.String()
}

// UnBindMesh calls the UnBindMesh method on the EffectFallbacks object.
//
// https://doc.babylonjs.com/api/classes/babylon.effectfallbacks#unbindmesh
func (e *EffectFallbacks) UnBindMesh() {

	e.p.Call("unBindMesh")
}

// HasMoreFallbacks returns the HasMoreFallbacks property of class EffectFallbacks.
//
// https://doc.babylonjs.com/api/classes/babylon.effectfallbacks#hasmorefallbacks
func (e *EffectFallbacks) HasMoreFallbacks() bool {
	retVal := e.p.Get("hasMoreFallbacks")
	return retVal.Bool()
}

// SetHasMoreFallbacks sets the HasMoreFallbacks property of class EffectFallbacks.
//
// https://doc.babylonjs.com/api/classes/babylon.effectfallbacks#hasmorefallbacks
func (e *EffectFallbacks) SetHasMoreFallbacks(hasMoreFallbacks bool) *EffectFallbacks {
	e.p.Set("hasMoreFallbacks", hasMoreFallbacks)
	return e
}
