// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ScatterPanel represents a babylon.js ScatterPanel.
// Class used to create a container panel where items get randomized planar mapping
type ScatterPanel struct {
	*VolumeBasedPanel
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (s *ScatterPanel) JSObject() js.Value { return s.p }

// ScatterPanel returns a ScatterPanel JavaScript class.
func (ba *Babylon) ScatterPanel() *ScatterPanel {
	p := ba.ctx.Get("ScatterPanel")
	return ScatterPanelFromJSObject(p, ba.ctx)
}

// ScatterPanelFromJSObject returns a wrapped ScatterPanel JavaScript class.
func ScatterPanelFromJSObject(p js.Value, ctx js.Value) *ScatterPanel {
	return &ScatterPanel{VolumeBasedPanel: VolumeBasedPanelFromJSObject(p, ctx), ctx: ctx}
}

// ScatterPanelArrayToJSArray returns a JavaScript Array for the wrapped array.
func ScatterPanelArrayToJSArray(array []*ScatterPanel) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewScatterPanel returns a new ScatterPanel object.
//
// https://doc.babylonjs.com/api/classes/babylon.scatterpanel
func (ba *Babylon) NewScatterPanel() *ScatterPanel {

	args := make([]interface{}, 0, 0+0)

	p := ba.ctx.Get("ScatterPanel").New(args...)
	return ScatterPanelFromJSObject(p, ba.ctx)
}

// Iteration returns the Iteration property of class ScatterPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.scatterpanel#iteration
func (s *ScatterPanel) Iteration() float64 {
	retVal := s.p.Get("iteration")
	return retVal.Float()
}

// SetIteration sets the Iteration property of class ScatterPanel.
//
// https://doc.babylonjs.com/api/classes/babylon.scatterpanel#iteration
func (s *ScatterPanel) SetIteration(iteration float64) *ScatterPanel {
	s.p.Set("iteration", iteration)
	return s
}
