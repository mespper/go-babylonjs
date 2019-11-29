// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// PostProcessesOptimization represents a babylon.js PostProcessesOptimization.
// Defines an optimization used to turn post-processes off

//
// See: http://doc.babylonjs.com/how_to/how_to_use_sceneoptimizer
type PostProcessesOptimization struct{ *SceneOptimization }

// JSObject returns the underlying js.Value.
func (p *PostProcessesOptimization) JSObject() js.Value { return p.p }

// PostProcessesOptimization returns a PostProcessesOptimization JavaScript class.
func (b *Babylon) PostProcessesOptimization() *PostProcessesOptimization {
	p := b.ctx.Get("PostProcessesOptimization")
	return PostProcessesOptimizationFromJSObject(p)
}

// PostProcessesOptimizationFromJSObject returns a wrapped PostProcessesOptimization JavaScript class.
func PostProcessesOptimizationFromJSObject(p js.Value) *PostProcessesOptimization {
	return &PostProcessesOptimization{SceneOptimizationFromJSObject(p)}
}

// NewPostProcessesOptimization returns a new PostProcessesOptimization object.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocessesoptimization
func (b *Babylon) NewPostProcessesOptimization(todo parameters) *PostProcessesOptimization {
	p := b.ctx.Get("PostProcessesOptimization").New(todo)
	return PostProcessesOptimizationFromJSObject(p)
}

// TODO: methods