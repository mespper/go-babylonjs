// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// Ray represents a babylon.js Ray.
// Class representing a ray with position and direction
type Ray struct{ p js.Value }

// JSObject returns the underlying js.Value.
func (r *Ray) JSObject() js.Value { return r.p }

// Ray returns a Ray JavaScript class.
func (ba *Babylon) Ray() *Ray {
	p := ba.ctx.Get("Ray")
	return RayFromJSObject(p)
}

// RayFromJSObject returns a wrapped Ray JavaScript class.
func RayFromJSObject(p js.Value) *Ray {
	return &Ray{p: p}
}

// NewRayOpts contains optional parameters for NewRay.
type NewRayOpts struct {
	Length *JSFloat64
}

// NewRay returns a new Ray object.
//
// https://doc.babylonjs.com/api/classes/babylon.ray
func (ba *Babylon) NewRay(origin *Vector3, direction *Vector3, opts *NewRayOpts) *Ray {
	if opts == nil {
		opts = &NewRayOpts{}
	}

	p := ba.ctx.Get("Ray").New(origin.JSObject(), direction.JSObject(), opts.Length)
	return RayFromJSObject(p)
}

// TODO: methods
