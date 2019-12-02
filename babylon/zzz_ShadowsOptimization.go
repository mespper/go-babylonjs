// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ShadowsOptimization represents a babylon.js ShadowsOptimization.
// Defines an optimization used to remove shadows
//
// See: http://doc.babylonjs.com/how_to/how_to_use_sceneoptimizer
type ShadowsOptimization struct {
	*SceneOptimization
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (s *ShadowsOptimization) JSObject() js.Value { return s.p }

// ShadowsOptimization returns a ShadowsOptimization JavaScript class.
func (ba *Babylon) ShadowsOptimization() *ShadowsOptimization {
	p := ba.ctx.Get("ShadowsOptimization")
	return ShadowsOptimizationFromJSObject(p, ba.ctx)
}

// ShadowsOptimizationFromJSObject returns a wrapped ShadowsOptimization JavaScript class.
func ShadowsOptimizationFromJSObject(p js.Value, ctx js.Value) *ShadowsOptimization {
	return &ShadowsOptimization{SceneOptimization: SceneOptimizationFromJSObject(p, ctx), ctx: ctx}
}

// ShadowsOptimizationArrayToJSArray returns a JavaScript Array for the wrapped array.
func ShadowsOptimizationArrayToJSArray(array []*ShadowsOptimization) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewShadowsOptimizationOpts contains optional parameters for NewShadowsOptimization.
type NewShadowsOptimizationOpts struct {
	Priority *float64
}

// NewShadowsOptimization returns a new ShadowsOptimization object.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowsoptimization
func (ba *Babylon) NewShadowsOptimization(opts *NewShadowsOptimizationOpts) *ShadowsOptimization {
	if opts == nil {
		opts = &NewShadowsOptimizationOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.Priority == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Priority)
	}

	p := ba.ctx.Get("ShadowsOptimization").New(args...)
	return ShadowsOptimizationFromJSObject(p, ba.ctx)
}

// Apply calls the Apply method on the ShadowsOptimization object.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowsoptimization#apply
func (s *ShadowsOptimization) Apply(scene *Scene, optimizer *SceneOptimizer) bool {

	args := make([]interface{}, 0, 2+0)

	args = append(args, scene.JSObject())
	args = append(args, optimizer.JSObject())

	retVal := s.p.Call("apply", args...)
	return retVal.Bool()
}

// GetDescription calls the GetDescription method on the ShadowsOptimization object.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowsoptimization#getdescription
func (s *ShadowsOptimization) GetDescription() string {

	retVal := s.p.Call("getDescription")
	return retVal.String()
}
