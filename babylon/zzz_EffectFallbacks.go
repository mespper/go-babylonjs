// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// EffectFallbacks represents a babylon.js EffectFallbacks.
// EffectFallbacks can be used to add fallbacks (properties to disable) to certain properties when desired to improve performance.
// (Eg. Start at high quality with reflection and fog, if fps is low, remove reflection, if still low remove fog)
type EffectFallbacks struct{ p js.Value }

// JSObject returns the underlying js.Value.
func (e *EffectFallbacks) JSObject() js.Value { return e.p }

// EffectFallbacks returns a EffectFallbacks JavaScript class.
func (ba *Babylon) EffectFallbacks() *EffectFallbacks {
	p := ba.ctx.Get("EffectFallbacks")
	return EffectFallbacksFromJSObject(p)
}

// EffectFallbacksFromJSObject returns a wrapped EffectFallbacks JavaScript class.
func EffectFallbacksFromJSObject(p js.Value) *EffectFallbacks {
	return &EffectFallbacks{p: p}
}

// TODO: methods
