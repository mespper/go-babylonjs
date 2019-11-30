// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// HardwareScalingOptimization represents a babylon.js HardwareScalingOptimization.
// Defines an optimization used to increase or decrease the rendering resolution
//
// See: http://doc.babylonjs.com/how_to/how_to_use_sceneoptimizer
type HardwareScalingOptimization struct{ *SceneOptimization }

// JSObject returns the underlying js.Value.
func (h *HardwareScalingOptimization) JSObject() js.Value { return h.p }

// HardwareScalingOptimization returns a HardwareScalingOptimization JavaScript class.
func (ba *Babylon) HardwareScalingOptimization() *HardwareScalingOptimization {
	p := ba.ctx.Get("HardwareScalingOptimization")
	return HardwareScalingOptimizationFromJSObject(p)
}

// HardwareScalingOptimizationFromJSObject returns a wrapped HardwareScalingOptimization JavaScript class.
func HardwareScalingOptimizationFromJSObject(p js.Value) *HardwareScalingOptimization {
	return &HardwareScalingOptimization{SceneOptimizationFromJSObject(p)}
}

// NewHardwareScalingOptimizationOpts contains optional parameters for NewHardwareScalingOptimization.
type NewHardwareScalingOptimizationOpts struct {
	Priority *JSFloat64

	MaximumScale *JSFloat64

	Step *JSFloat64
}

// NewHardwareScalingOptimization returns a new HardwareScalingOptimization object.
//
// https://doc.babylonjs.com/api/classes/babylon.hardwarescalingoptimization
func (ba *Babylon) NewHardwareScalingOptimization(opts *NewHardwareScalingOptimizationOpts) *HardwareScalingOptimization {
	if opts == nil {
		opts = &NewHardwareScalingOptimizationOpts{}
	}

	p := ba.ctx.Get("HardwareScalingOptimization").New(opts.Priority.JSObject(), opts.MaximumScale.JSObject(), opts.Step.JSObject())
	return HardwareScalingOptimizationFromJSObject(p)
}

// TODO: methods
