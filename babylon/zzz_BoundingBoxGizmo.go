// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// BoundingBoxGizmo represents a babylon.js BoundingBoxGizmo.
// Bounding box gizmo
type BoundingBoxGizmo struct{ *Gizmo }

// JSObject returns the underlying js.Value.
func (b *BoundingBoxGizmo) JSObject() js.Value { return b.p }

// BoundingBoxGizmo returns a BoundingBoxGizmo JavaScript class.
func (ba *Babylon) BoundingBoxGizmo() *BoundingBoxGizmo {
	p := ba.ctx.Get("BoundingBoxGizmo")
	return BoundingBoxGizmoFromJSObject(p)
}

// BoundingBoxGizmoFromJSObject returns a wrapped BoundingBoxGizmo JavaScript class.
func BoundingBoxGizmoFromJSObject(p js.Value) *BoundingBoxGizmo {
	return &BoundingBoxGizmo{GizmoFromJSObject(p)}
}

// NewBoundingBoxGizmoOpts contains optional parameters for NewBoundingBoxGizmo.
type NewBoundingBoxGizmoOpts struct {
	Color *Color3

	GizmoLayer *UtilityLayerRenderer
}

// NewBoundingBoxGizmo returns a new BoundingBoxGizmo object.
//
// https://doc.babylonjs.com/api/classes/babylon.boundingboxgizmo
func (ba *Babylon) NewBoundingBoxGizmo(opts *NewBoundingBoxGizmoOpts) *BoundingBoxGizmo {
	if opts == nil {
		opts = &NewBoundingBoxGizmoOpts{}
	}

	p := ba.ctx.Get("BoundingBoxGizmo").New(opts.Color.JSObject(), opts.GizmoLayer.JSObject())
	return BoundingBoxGizmoFromJSObject(p)
}

// TODO: methods
