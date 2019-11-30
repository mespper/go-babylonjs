// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// KeepAssets represents a babylon.js KeepAssets.
// Set of assets to keep when moving a scene into an asset container.
type KeepAssets struct{ *AbstractScene }

// JSObject returns the underlying js.Value.
func (k *KeepAssets) JSObject() js.Value { return k.p }

// KeepAssets returns a KeepAssets JavaScript class.
func (ba *Babylon) KeepAssets() *KeepAssets {
	p := ba.ctx.Get("KeepAssets")
	return KeepAssetsFromJSObject(p)
}

// KeepAssetsFromJSObject returns a wrapped KeepAssets JavaScript class.
func KeepAssetsFromJSObject(p js.Value) *KeepAssets {
	return &KeepAssets{AbstractSceneFromJSObject(p)}
}

// TODO: methods
