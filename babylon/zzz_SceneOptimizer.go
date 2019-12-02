// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// SceneOptimizer represents a babylon.js SceneOptimizer.
// Class used to run optimizations in order to reach a target frame rate
//
// See: http://doc.babylonjs.com/how_to/how_to_use_sceneoptimizer
type SceneOptimizer struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (s *SceneOptimizer) JSObject() js.Value { return s.p }

// SceneOptimizer returns a SceneOptimizer JavaScript class.
func (ba *Babylon) SceneOptimizer() *SceneOptimizer {
	p := ba.ctx.Get("SceneOptimizer")
	return SceneOptimizerFromJSObject(p, ba.ctx)
}

// SceneOptimizerFromJSObject returns a wrapped SceneOptimizer JavaScript class.
func SceneOptimizerFromJSObject(p js.Value, ctx js.Value) *SceneOptimizer {
	return &SceneOptimizer{p: p, ctx: ctx}
}

// SceneOptimizerArrayToJSArray returns a JavaScript Array for the wrapped array.
func SceneOptimizerArrayToJSArray(array []*SceneOptimizer) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewSceneOptimizerOpts contains optional parameters for NewSceneOptimizer.
type NewSceneOptimizerOpts struct {
	Options                *SceneOptimizerOptions
	AutoGeneratePriorities *bool
	ImprovementMode        *bool
}

// NewSceneOptimizer returns a new SceneOptimizer object.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneoptimizer
func (ba *Babylon) NewSceneOptimizer(scene *Scene, opts *NewSceneOptimizerOpts) *SceneOptimizer {
	if opts == nil {
		opts = &NewSceneOptimizerOpts{}
	}

	args := make([]interface{}, 0, 1+3)

	args = append(args, scene.JSObject())

	if opts.Options == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Options.JSObject())
	}
	if opts.AutoGeneratePriorities == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.AutoGeneratePriorities)
	}
	if opts.ImprovementMode == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.ImprovementMode)
	}

	p := ba.ctx.Get("SceneOptimizer").New(args...)
	return SceneOptimizerFromJSObject(p, ba.ctx)
}

// Dispose calls the Dispose method on the SceneOptimizer object.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneoptimizer#dispose
func (s *SceneOptimizer) Dispose() {

	s.p.Call("dispose")
}

// SceneOptimizerOptimizeAsyncOpts contains optional parameters for SceneOptimizer.OptimizeAsync.
type SceneOptimizerOptimizeAsyncOpts struct {
	Options   *SceneOptimizerOptions
	OnSuccess func()
	OnFailure func()
}

// OptimizeAsync calls the OptimizeAsync method on the SceneOptimizer object.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneoptimizer#optimizeasync
func (s *SceneOptimizer) OptimizeAsync(scene *Scene, opts *SceneOptimizerOptimizeAsyncOpts) *SceneOptimizer {
	if opts == nil {
		opts = &SceneOptimizerOptimizeAsyncOpts{}
	}

	args := make([]interface{}, 0, 1+3)

	args = append(args, scene.JSObject())

	if opts.Options == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Options.JSObject())
	}
	if opts.OnSuccess == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.OnSuccess)
	}
	if opts.OnFailure == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.OnFailure)
	}

	retVal := s.p.Call("OptimizeAsync", args...)
	return SceneOptimizerFromJSObject(retVal, s.ctx)
}

// Reset calls the Reset method on the SceneOptimizer object.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneoptimizer#reset
func (s *SceneOptimizer) Reset() {

	s.p.Call("reset")
}

// Start calls the Start method on the SceneOptimizer object.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneoptimizer#start
func (s *SceneOptimizer) Start() {

	s.p.Call("start")
}

// Stop calls the Stop method on the SceneOptimizer object.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneoptimizer#stop
func (s *SceneOptimizer) Stop() {

	s.p.Call("stop")
}

// CurrentFrameRate returns the CurrentFrameRate property of class SceneOptimizer.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneoptimizer#currentframerate
func (s *SceneOptimizer) CurrentFrameRate() float64 {
	retVal := s.p.Get("currentFrameRate")
	return retVal.Float()
}

// SetCurrentFrameRate sets the CurrentFrameRate property of class SceneOptimizer.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneoptimizer#currentframerate
func (s *SceneOptimizer) SetCurrentFrameRate(currentFrameRate float64) *SceneOptimizer {
	s.p.Set("currentFrameRate", currentFrameRate)
	return s
}

// CurrentPriorityLevel returns the CurrentPriorityLevel property of class SceneOptimizer.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneoptimizer#currentprioritylevel
func (s *SceneOptimizer) CurrentPriorityLevel() float64 {
	retVal := s.p.Get("currentPriorityLevel")
	return retVal.Float()
}

// SetCurrentPriorityLevel sets the CurrentPriorityLevel property of class SceneOptimizer.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneoptimizer#currentprioritylevel
func (s *SceneOptimizer) SetCurrentPriorityLevel(currentPriorityLevel float64) *SceneOptimizer {
	s.p.Set("currentPriorityLevel", currentPriorityLevel)
	return s
}

// IsInImprovementMode returns the IsInImprovementMode property of class SceneOptimizer.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneoptimizer#isinimprovementmode
func (s *SceneOptimizer) IsInImprovementMode() bool {
	retVal := s.p.Get("isInImprovementMode")
	return retVal.Bool()
}

// SetIsInImprovementMode sets the IsInImprovementMode property of class SceneOptimizer.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneoptimizer#isinimprovementmode
func (s *SceneOptimizer) SetIsInImprovementMode(isInImprovementMode bool) *SceneOptimizer {
	s.p.Set("isInImprovementMode", isInImprovementMode)
	return s
}

// OnFailureObservable returns the OnFailureObservable property of class SceneOptimizer.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneoptimizer#onfailureobservable
func (s *SceneOptimizer) OnFailureObservable() *Observable {
	retVal := s.p.Get("onFailureObservable")
	return ObservableFromJSObject(retVal, s.ctx)
}

// SetOnFailureObservable sets the OnFailureObservable property of class SceneOptimizer.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneoptimizer#onfailureobservable
func (s *SceneOptimizer) SetOnFailureObservable(onFailureObservable *Observable) *SceneOptimizer {
	s.p.Set("onFailureObservable", onFailureObservable.JSObject())
	return s
}

// OnNewOptimizationAppliedObservable returns the OnNewOptimizationAppliedObservable property of class SceneOptimizer.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneoptimizer#onnewoptimizationappliedobservable
func (s *SceneOptimizer) OnNewOptimizationAppliedObservable() *Observable {
	retVal := s.p.Get("onNewOptimizationAppliedObservable")
	return ObservableFromJSObject(retVal, s.ctx)
}

// SetOnNewOptimizationAppliedObservable sets the OnNewOptimizationAppliedObservable property of class SceneOptimizer.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneoptimizer#onnewoptimizationappliedobservable
func (s *SceneOptimizer) SetOnNewOptimizationAppliedObservable(onNewOptimizationAppliedObservable *Observable) *SceneOptimizer {
	s.p.Set("onNewOptimizationAppliedObservable", onNewOptimizationAppliedObservable.JSObject())
	return s
}

// OnSuccessObservable returns the OnSuccessObservable property of class SceneOptimizer.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneoptimizer#onsuccessobservable
func (s *SceneOptimizer) OnSuccessObservable() *Observable {
	retVal := s.p.Get("onSuccessObservable")
	return ObservableFromJSObject(retVal, s.ctx)
}

// SetOnSuccessObservable sets the OnSuccessObservable property of class SceneOptimizer.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneoptimizer#onsuccessobservable
func (s *SceneOptimizer) SetOnSuccessObservable(onSuccessObservable *Observable) *SceneOptimizer {
	s.p.Set("onSuccessObservable", onSuccessObservable.JSObject())
	return s
}

// Optimizations returns the Optimizations property of class SceneOptimizer.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneoptimizer#optimizations
func (s *SceneOptimizer) Optimizations() *SceneOptimization {
	retVal := s.p.Get("optimizations")
	return SceneOptimizationFromJSObject(retVal, s.ctx)
}

// SetOptimizations sets the Optimizations property of class SceneOptimizer.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneoptimizer#optimizations
func (s *SceneOptimizer) SetOptimizations(optimizations *SceneOptimization) *SceneOptimizer {
	s.p.Set("optimizations", optimizations.JSObject())
	return s
}

// TargetFrameRate returns the TargetFrameRate property of class SceneOptimizer.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneoptimizer#targetframerate
func (s *SceneOptimizer) TargetFrameRate() float64 {
	retVal := s.p.Get("targetFrameRate")
	return retVal.Float()
}

// SetTargetFrameRate sets the TargetFrameRate property of class SceneOptimizer.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneoptimizer#targetframerate
func (s *SceneOptimizer) SetTargetFrameRate(targetFrameRate float64) *SceneOptimizer {
	s.p.Set("targetFrameRate", targetFrameRate)
	return s
}

// TrackerDuration returns the TrackerDuration property of class SceneOptimizer.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneoptimizer#trackerduration
func (s *SceneOptimizer) TrackerDuration() float64 {
	retVal := s.p.Get("trackerDuration")
	return retVal.Float()
}

// SetTrackerDuration sets the TrackerDuration property of class SceneOptimizer.
//
// https://doc.babylonjs.com/api/classes/babylon.sceneoptimizer#trackerduration
func (s *SceneOptimizer) SetTrackerDuration(trackerDuration float64) *SceneOptimizer {
	s.p.Set("trackerDuration", trackerDuration)
	return s
}
