// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// PlaneRotationGizmo represents a babylon.js PlaneRotationGizmo.
// Single plane rotation gizmo
type PlaneRotationGizmo struct{ *Gizmo }

// JSObject returns the underlying js.Value.
func (p *PlaneRotationGizmo) JSObject() js.Value { return p.p }

// PlaneRotationGizmo returns a PlaneRotationGizmo JavaScript class.
func (ba *Babylon) PlaneRotationGizmo() *PlaneRotationGizmo {
	p := ba.ctx.Get("PlaneRotationGizmo")
	return PlaneRotationGizmoFromJSObject(p)
}

// PlaneRotationGizmoFromJSObject returns a wrapped PlaneRotationGizmo JavaScript class.
func PlaneRotationGizmoFromJSObject(p js.Value) *PlaneRotationGizmo {
	return &PlaneRotationGizmo{GizmoFromJSObject(p)}
}

// NewPlaneRotationGizmoOpts contains optional parameters for NewPlaneRotationGizmo.
type NewPlaneRotationGizmoOpts struct {
	Color *Color3

	GizmoLayer *UtilityLayerRenderer

	Tessellation *JSFloat64

	Parent *RotationGizmo

	UseEulerRotation *JSBool
}

// NewPlaneRotationGizmo returns a new PlaneRotationGizmo object.
//
// https://doc.babylonjs.com/api/classes/babylon.planerotationgizmo
func (ba *Babylon) NewPlaneRotationGizmo(planeNormal *Vector3, opts *NewPlaneRotationGizmoOpts) *PlaneRotationGizmo {
	if opts == nil {
		opts = &NewPlaneRotationGizmoOpts{}
	}

	p := ba.ctx.Get("PlaneRotationGizmo").New(planeNormal.JSObject(), opts.Color.JSObject(), opts.GizmoLayer.JSObject(), opts.Tessellation, opts.Parent.JSObject(), opts.UseEulerRotation)
	return PlaneRotationGizmoFromJSObject(p)
}

// TODO: methods
