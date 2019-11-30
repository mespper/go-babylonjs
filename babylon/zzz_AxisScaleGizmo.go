// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// AxisScaleGizmo represents a babylon.js AxisScaleGizmo.
// Single axis scale gizmo
type AxisScaleGizmo struct{ *Gizmo }

// JSObject returns the underlying js.Value.
func (a *AxisScaleGizmo) JSObject() js.Value { return a.p }

// AxisScaleGizmo returns a AxisScaleGizmo JavaScript class.
func (ba *Babylon) AxisScaleGizmo() *AxisScaleGizmo {
	p := ba.ctx.Get("AxisScaleGizmo")
	return AxisScaleGizmoFromJSObject(p)
}

// AxisScaleGizmoFromJSObject returns a wrapped AxisScaleGizmo JavaScript class.
func AxisScaleGizmoFromJSObject(p js.Value) *AxisScaleGizmo {
	return &AxisScaleGizmo{GizmoFromJSObject(p)}
}

// NewAxisScaleGizmoOpts contains optional parameters for NewAxisScaleGizmo.
type NewAxisScaleGizmoOpts struct {
	Color *Color3

	GizmoLayer *UtilityLayerRenderer

	Parent *ScaleGizmo
}

// NewAxisScaleGizmo returns a new AxisScaleGizmo object.
//
// https://doc.babylonjs.com/api/classes/babylon.axisscalegizmo
func (ba *Babylon) NewAxisScaleGizmo(dragAxis *Vector3, opts *NewAxisScaleGizmoOpts) *AxisScaleGizmo {
	if opts == nil {
		opts = &NewAxisScaleGizmoOpts{}
	}

	p := ba.ctx.Get("AxisScaleGizmo").New(dragAxis.JSObject(), opts.Color.JSObject(), opts.GizmoLayer.JSObject(), opts.Parent.JSObject())
	return AxisScaleGizmoFromJSObject(p)
}

// TODO: methods
