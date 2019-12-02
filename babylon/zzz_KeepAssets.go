// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// KeepAssets represents a babylon.js KeepAssets.
// Set of assets to keep when moving a scene into an asset container.
type KeepAssets struct {
	*AbstractScene
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (k *KeepAssets) JSObject() js.Value { return k.p }

// KeepAssets returns a KeepAssets JavaScript class.
func (ba *Babylon) KeepAssets() *KeepAssets {
	p := ba.ctx.Get("KeepAssets")
	return KeepAssetsFromJSObject(p, ba.ctx)
}

// KeepAssetsFromJSObject returns a wrapped KeepAssets JavaScript class.
func KeepAssetsFromJSObject(p js.Value, ctx js.Value) *KeepAssets {
	return &KeepAssets{AbstractScene: AbstractSceneFromJSObject(p, ctx), ctx: ctx}
}

// KeepAssetsArrayToJSArray returns a JavaScript Array for the wrapped array.
func KeepAssetsArrayToJSArray(array []*KeepAssets) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}
