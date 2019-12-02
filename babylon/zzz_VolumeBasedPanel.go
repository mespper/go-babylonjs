// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// VolumeBasedPanel represents a babylon.js VolumeBasedPanel.
// Abstract class used to create a container panel deployed on the surface of a volume
type VolumeBasedPanel struct {
	*Container3D
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (v *VolumeBasedPanel) JSObject() js.Value { return v.p }

// VolumeBasedPanel returns a VolumeBasedPanel JavaScript class.
func (ba *Babylon) VolumeBasedPanel() *VolumeBasedPanel {
	p := ba.ctx.Get("VolumeBasedPanel")
	return VolumeBasedPanelFromJSObject(p, ba.ctx)
}

// VolumeBasedPanelFromJSObject returns a wrapped VolumeBasedPanel JavaScript class.
func VolumeBasedPanelFromJSObject(p js.Value, ctx js.Value) *VolumeBasedPanel {
	return &VolumeBasedPanel{Container3D: Container3DFromJSObject(p, ctx), ctx: ctx}
}

// VolumeBasedPanelArrayToJSArray returns a JavaScript Array for the wrapped array.
func VolumeBasedPanelArrayToJSArray(array []*VolumeBasedPanel) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewVolumeBasedPanel returns a new VolumeBasedPanel object.
//
// https://doc.babylonjs.com/api/classes/babylon.volumebasedpanel
func (ba *Babylon) NewVolumeBasedPanel() *VolumeBasedPanel {

	args := make([]interface{}, 0, 0+0)

	p := ba.ctx.Get("VolumeBasedPanel").New(args...)
	return VolumeBasedPanelFromJSObject(p, ba.ctx)
}

// Columns returns the Columns property of class VolumeBasedPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.volumebasedpanel#columns
func (v *VolumeBasedPanel) Columns() int {
	retVal := v.p.Get("columns")
	return retVal.Int()
}

// SetColumns sets the Columns property of class VolumeBasedPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.volumebasedpanel#columns
func (v *VolumeBasedPanel) SetColumns(columns int) *VolumeBasedPanel {
	v.p.Set("columns", columns)
	return v
}

// Margin returns the Margin property of class VolumeBasedPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.volumebasedpanel#margin
func (v *VolumeBasedPanel) Margin() float64 {
	retVal := v.p.Get("margin")
	return retVal.Float()
}

// SetMargin sets the Margin property of class VolumeBasedPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.volumebasedpanel#margin
func (v *VolumeBasedPanel) SetMargin(margin float64) *VolumeBasedPanel {
	v.p.Set("margin", margin)
	return v
}

// Orientation returns the Orientation property of class VolumeBasedPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.volumebasedpanel#orientation
func (v *VolumeBasedPanel) Orientation() float64 {
	retVal := v.p.Get("orientation")
	return retVal.Float()
}

// SetOrientation sets the Orientation property of class VolumeBasedPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.volumebasedpanel#orientation
func (v *VolumeBasedPanel) SetOrientation(orientation float64) *VolumeBasedPanel {
	v.p.Set("orientation", orientation)
	return v
}

// Rows returns the Rows property of class VolumeBasedPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.volumebasedpanel#rows
func (v *VolumeBasedPanel) Rows() int {
	retVal := v.p.Get("rows")
	return retVal.Int()
}

// SetRows sets the Rows property of class VolumeBasedPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.volumebasedpanel#rows
func (v *VolumeBasedPanel) SetRows(rows int) *VolumeBasedPanel {
	v.p.Set("rows", rows)
	return v
}
