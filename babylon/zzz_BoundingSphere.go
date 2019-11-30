// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// BoundingSphere represents a babylon.js BoundingSphere.
// Class used to store bounding sphere information
type BoundingSphere struct{ p js.Value }

// JSObject returns the underlying js.Value.
func (b *BoundingSphere) JSObject() js.Value { return b.p }

// BoundingSphere returns a BoundingSphere JavaScript class.
func (ba *Babylon) BoundingSphere() *BoundingSphere {
	p := ba.ctx.Get("BoundingSphere")
	return BoundingSphereFromJSObject(p)
}

// BoundingSphereFromJSObject returns a wrapped BoundingSphere JavaScript class.
func BoundingSphereFromJSObject(p js.Value) *BoundingSphere {
	return &BoundingSphere{p: p}
}

// NewBoundingSphereOpts contains optional parameters for NewBoundingSphere.
type NewBoundingSphereOpts struct {
	WorldMatrix *Matrix
}

// NewBoundingSphere returns a new BoundingSphere object.
//
// https://doc.babylonjs.com/api/classes/babylon.boundingsphere
func (ba *Babylon) NewBoundingSphere(min *Vector3, max *Vector3, opts *NewBoundingSphereOpts) *BoundingSphere {
	if opts == nil {
		opts = &NewBoundingSphereOpts{}
	}

	p := ba.ctx.Get("BoundingSphere").New(min.JSObject(), max.JSObject(), opts.WorldMatrix.JSObject())
	return BoundingSphereFromJSObject(p)
}

// TODO: methods
