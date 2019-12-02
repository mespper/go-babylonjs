// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// AxisScaleGizmo represents a babylon.js AxisScaleGizmo.
// Single axis scale gizmo
type AxisScaleGizmo struct {
	*Gizmo
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (a *AxisScaleGizmo) JSObject() js.Value { return a.p }

// AxisScaleGizmo returns a AxisScaleGizmo JavaScript class.
func (ba *Babylon) AxisScaleGizmo() *AxisScaleGizmo {
	p := ba.ctx.Get("AxisScaleGizmo")
	return AxisScaleGizmoFromJSObject(p, ba.ctx)
}

// AxisScaleGizmoFromJSObject returns a wrapped AxisScaleGizmo JavaScript class.
func AxisScaleGizmoFromJSObject(p js.Value, ctx js.Value) *AxisScaleGizmo {
	return &AxisScaleGizmo{Gizmo: GizmoFromJSObject(p, ctx), ctx: ctx}
}

// AxisScaleGizmoArrayToJSArray returns a JavaScript Array for the wrapped array.
func AxisScaleGizmoArrayToJSArray(array []*AxisScaleGizmo) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewAxisScaleGizmoOpts contains optional parameters for NewAxisScaleGizmo.
type NewAxisScaleGizmoOpts struct {
	Color      *Color3
	GizmoLayer *UtilityLayerRenderer
	Parent     *ScaleGizmo
}

// NewAxisScaleGizmo returns a new AxisScaleGizmo object.
//
// https://doc.babylonjs.com/api/classes/babylon.axisscalegizmo
func (ba *Babylon) NewAxisScaleGizmo(dragAxis *Vector3, opts *NewAxisScaleGizmoOpts) *AxisScaleGizmo {
	if opts == nil {
		opts = &NewAxisScaleGizmoOpts{}
	}

	args := make([]interface{}, 0, 1+3)

	args = append(args, dragAxis.JSObject())

	if opts.Color == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Color.JSObject())
	}
	if opts.GizmoLayer == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.GizmoLayer.JSObject())
	}
	if opts.Parent == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Parent.JSObject())
	}

	p := ba.ctx.Get("AxisScaleGizmo").New(args...)
	return AxisScaleGizmoFromJSObject(p, ba.ctx)
}

// Dispose calls the Dispose method on the AxisScaleGizmo object.
//
// https://doc.babylonjs.com/api/classes/babylon.axisscalegizmo#dispose
func (a *AxisScaleGizmo) Dispose() {

	a.p.Call("dispose")
}

// AxisScaleGizmoSetCustomMeshOpts contains optional parameters for AxisScaleGizmo.SetCustomMesh.
type AxisScaleGizmoSetCustomMeshOpts struct {
	UseGizmoMaterial *bool
}

// SetCustomMesh calls the SetCustomMesh method on the AxisScaleGizmo object.
//
// https://doc.babylonjs.com/api/classes/babylon.axisscalegizmo#setcustommesh
func (a *AxisScaleGizmo) SetCustomMesh(mesh *Mesh, opts *AxisScaleGizmoSetCustomMeshOpts) {
	if opts == nil {
		opts = &AxisScaleGizmoSetCustomMeshOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, mesh.JSObject())

	if opts.UseGizmoMaterial == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.UseGizmoMaterial)
	}

	a.p.Call("setCustomMesh", args...)
}

// DragBehavior returns the DragBehavior property of class AxisScaleGizmo.
//
// https://doc.babylonjs.com/api/classes/babylon.axisscalegizmo#dragbehavior
func (a *AxisScaleGizmo) DragBehavior() *PointerDragBehavior {
	retVal := a.p.Get("dragBehavior")
	return PointerDragBehaviorFromJSObject(retVal, a.ctx)
}

// SetDragBehavior sets the DragBehavior property of class AxisScaleGizmo.
//
// https://doc.babylonjs.com/api/classes/babylon.axisscalegizmo#dragbehavior
func (a *AxisScaleGizmo) SetDragBehavior(dragBehavior *PointerDragBehavior) *AxisScaleGizmo {
	a.p.Set("dragBehavior", dragBehavior.JSObject())
	return a
}

// IsEnabled returns the IsEnabled property of class AxisScaleGizmo.
//
// https://doc.babylonjs.com/api/classes/babylon.axisscalegizmo#isenabled
func (a *AxisScaleGizmo) IsEnabled() bool {
	retVal := a.p.Get("isEnabled")
	return retVal.Bool()
}

// SetIsEnabled sets the IsEnabled property of class AxisScaleGizmo.
//
// https://doc.babylonjs.com/api/classes/babylon.axisscalegizmo#isenabled
func (a *AxisScaleGizmo) SetIsEnabled(isEnabled bool) *AxisScaleGizmo {
	a.p.Set("isEnabled", isEnabled)
	return a
}

// OnSnapObservable returns the OnSnapObservable property of class AxisScaleGizmo.
//
// https://doc.babylonjs.com/api/classes/babylon.axisscalegizmo#onsnapobservable
func (a *AxisScaleGizmo) OnSnapObservable() *Observable {
	retVal := a.p.Get("onSnapObservable")
	return ObservableFromJSObject(retVal, a.ctx)
}

// SetOnSnapObservable sets the OnSnapObservable property of class AxisScaleGizmo.
//
// https://doc.babylonjs.com/api/classes/babylon.axisscalegizmo#onsnapobservable
func (a *AxisScaleGizmo) SetOnSnapObservable(onSnapObservable *Observable) *AxisScaleGizmo {
	a.p.Set("onSnapObservable", onSnapObservable.JSObject())
	return a
}

// Sensitivity returns the Sensitivity property of class AxisScaleGizmo.
//
// https://doc.babylonjs.com/api/classes/babylon.axisscalegizmo#sensitivity
func (a *AxisScaleGizmo) Sensitivity() float64 {
	retVal := a.p.Get("sensitivity")
	return retVal.Float()
}

// SetSensitivity sets the Sensitivity property of class AxisScaleGizmo.
//
// https://doc.babylonjs.com/api/classes/babylon.axisscalegizmo#sensitivity
func (a *AxisScaleGizmo) SetSensitivity(sensitivity float64) *AxisScaleGizmo {
	a.p.Set("sensitivity", sensitivity)
	return a
}

// SnapDistance returns the SnapDistance property of class AxisScaleGizmo.
//
// https://doc.babylonjs.com/api/classes/babylon.axisscalegizmo#snapdistance
func (a *AxisScaleGizmo) SnapDistance() float64 {
	retVal := a.p.Get("snapDistance")
	return retVal.Float()
}

// SetSnapDistance sets the SnapDistance property of class AxisScaleGizmo.
//
// https://doc.babylonjs.com/api/classes/babylon.axisscalegizmo#snapdistance
func (a *AxisScaleGizmo) SetSnapDistance(snapDistance float64) *AxisScaleGizmo {
	a.p.Set("snapDistance", snapDistance)
	return a
}

// UniformScaling returns the UniformScaling property of class AxisScaleGizmo.
//
// https://doc.babylonjs.com/api/classes/babylon.axisscalegizmo#uniformscaling
func (a *AxisScaleGizmo) UniformScaling() bool {
	retVal := a.p.Get("uniformScaling")
	return retVal.Bool()
}

// SetUniformScaling sets the UniformScaling property of class AxisScaleGizmo.
//
// https://doc.babylonjs.com/api/classes/babylon.axisscalegizmo#uniformscaling
func (a *AxisScaleGizmo) SetUniformScaling(uniformScaling bool) *AxisScaleGizmo {
	a.p.Set("uniformScaling", uniformScaling)
	return a
}
