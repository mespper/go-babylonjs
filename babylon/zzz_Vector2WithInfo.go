// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// Vector2WithInfo represents a babylon.js Vector2WithInfo.
// Class used to transport BABYLON.Vector2 information for pointer events
type Vector2WithInfo struct{ *Vector2 }

// JSObject returns the underlying js.Value.
func (v *Vector2WithInfo) JSObject() js.Value { return v.p }

// Vector2WithInfo returns a Vector2WithInfo JavaScript class.
func (ba *Babylon) Vector2WithInfo() *Vector2WithInfo {
	p := ba.ctx.Get("Vector2WithInfo")
	return Vector2WithInfoFromJSObject(p)
}

// Vector2WithInfoFromJSObject returns a wrapped Vector2WithInfo JavaScript class.
func Vector2WithInfoFromJSObject(p js.Value) *Vector2WithInfo {
	return &Vector2WithInfo{Vector2FromJSObject(p)}
}

// NewVector2WithInfoOpts contains optional parameters for NewVector2WithInfo.
type NewVector2WithInfoOpts struct {
	ButtonIndex *JSFloat64
}

// NewVector2WithInfo returns a new Vector2WithInfo object.
//
// https://doc.babylonjs.com/api/classes/babylon.vector2withinfo
func (ba *Babylon) NewVector2WithInfo(source *Vector2, opts *NewVector2WithInfoOpts) *Vector2WithInfo {
	if opts == nil {
		opts = &NewVector2WithInfoOpts{}
	}

	p := ba.ctx.Get("Vector2WithInfo").New(source.JSObject(), opts.ButtonIndex.JSObject())
	return Vector2WithInfoFromJSObject(p)
}

// TODO: methods
