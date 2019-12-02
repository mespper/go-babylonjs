// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// EngineOptions represents a babylon.js EngineOptions.
// Interface defining initialization parameters for Engine class
type EngineOptions struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (e *EngineOptions) JSObject() js.Value { return e.p }

// EngineOptions returns a EngineOptions JavaScript class.
func (ba *Babylon) EngineOptions() *EngineOptions {
	p := ba.ctx.Get("EngineOptions")
	return EngineOptionsFromJSObject(p, ba.ctx)
}

// EngineOptionsFromJSObject returns a wrapped EngineOptions JavaScript class.
func EngineOptionsFromJSObject(p js.Value, ctx js.Value) *EngineOptions {
	return &EngineOptions{p: p, ctx: ctx}
}

// EngineOptionsArrayToJSArray returns a JavaScript Array for the wrapped array.
func EngineOptionsArrayToJSArray(array []*EngineOptions) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// AudioEngine returns the AudioEngine property of class EngineOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.engineoptions#audioengine
func (e *EngineOptions) AudioEngine() bool {
	retVal := e.p.Get("audioEngine")
	return retVal.Bool()
}

// SetAudioEngine sets the AudioEngine property of class EngineOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.engineoptions#audioengine
func (e *EngineOptions) SetAudioEngine(audioEngine bool) *EngineOptions {
	e.p.Set("audioEngine", audioEngine)
	return e
}

// AutoEnableWebVR returns the AutoEnableWebVR property of class EngineOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.engineoptions#autoenablewebvr
func (e *EngineOptions) AutoEnableWebVR() bool {
	retVal := e.p.Get("autoEnableWebVR")
	return retVal.Bool()
}

// SetAutoEnableWebVR sets the AutoEnableWebVR property of class EngineOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.engineoptions#autoenablewebvr
func (e *EngineOptions) SetAutoEnableWebVR(autoEnableWebVR bool) *EngineOptions {
	e.p.Set("autoEnableWebVR", autoEnableWebVR)
	return e
}

// DeterministicLockstep returns the DeterministicLockstep property of class EngineOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.engineoptions#deterministiclockstep
func (e *EngineOptions) DeterministicLockstep() bool {
	retVal := e.p.Get("deterministicLockstep")
	return retVal.Bool()
}

// SetDeterministicLockstep sets the DeterministicLockstep property of class EngineOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.engineoptions#deterministiclockstep
func (e *EngineOptions) SetDeterministicLockstep(deterministicLockstep bool) *EngineOptions {
	e.p.Set("deterministicLockstep", deterministicLockstep)
	return e
}

// DisableWebGL2Support returns the DisableWebGL2Support property of class EngineOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.engineoptions#disablewebgl2support
func (e *EngineOptions) DisableWebGL2Support() bool {
	retVal := e.p.Get("disableWebGL2Support")
	return retVal.Bool()
}

// SetDisableWebGL2Support sets the DisableWebGL2Support property of class EngineOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.engineoptions#disablewebgl2support
func (e *EngineOptions) SetDisableWebGL2Support(disableWebGL2Support bool) *EngineOptions {
	e.p.Set("disableWebGL2Support", disableWebGL2Support)
	return e
}

// DoNotHandleContextLost returns the DoNotHandleContextLost property of class EngineOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.engineoptions#donothandlecontextlost
func (e *EngineOptions) DoNotHandleContextLost() bool {
	retVal := e.p.Get("doNotHandleContextLost")
	return retVal.Bool()
}

// SetDoNotHandleContextLost sets the DoNotHandleContextLost property of class EngineOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.engineoptions#donothandlecontextlost
func (e *EngineOptions) SetDoNotHandleContextLost(doNotHandleContextLost bool) *EngineOptions {
	e.p.Set("doNotHandleContextLost", doNotHandleContextLost)
	return e
}

// DoNotHandleTouchAction returns the DoNotHandleTouchAction property of class EngineOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.engineoptions#donothandletouchaction
func (e *EngineOptions) DoNotHandleTouchAction() bool {
	retVal := e.p.Get("doNotHandleTouchAction")
	return retVal.Bool()
}

// SetDoNotHandleTouchAction sets the DoNotHandleTouchAction property of class EngineOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.engineoptions#donothandletouchaction
func (e *EngineOptions) SetDoNotHandleTouchAction(doNotHandleTouchAction bool) *EngineOptions {
	e.p.Set("doNotHandleTouchAction", doNotHandleTouchAction)
	return e
}

// LimitDeviceRatio returns the LimitDeviceRatio property of class EngineOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.engineoptions#limitdeviceratio
func (e *EngineOptions) LimitDeviceRatio() float64 {
	retVal := e.p.Get("limitDeviceRatio")
	return retVal.Float()
}

// SetLimitDeviceRatio sets the LimitDeviceRatio property of class EngineOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.engineoptions#limitdeviceratio
func (e *EngineOptions) SetLimitDeviceRatio(limitDeviceRatio float64) *EngineOptions {
	e.p.Set("limitDeviceRatio", limitDeviceRatio)
	return e
}

// LockstepMaxSteps returns the LockstepMaxSteps property of class EngineOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.engineoptions#lockstepmaxsteps
func (e *EngineOptions) LockstepMaxSteps() float64 {
	retVal := e.p.Get("lockstepMaxSteps")
	return retVal.Float()
}

// SetLockstepMaxSteps sets the LockstepMaxSteps property of class EngineOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.engineoptions#lockstepmaxsteps
func (e *EngineOptions) SetLockstepMaxSteps(lockstepMaxSteps float64) *EngineOptions {
	e.p.Set("lockstepMaxSteps", lockstepMaxSteps)
	return e
}

// TimeStep returns the TimeStep property of class EngineOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.engineoptions#timestep
func (e *EngineOptions) TimeStep() float64 {
	retVal := e.p.Get("timeStep")
	return retVal.Float()
}

// SetTimeStep sets the TimeStep property of class EngineOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.engineoptions#timestep
func (e *EngineOptions) SetTimeStep(timeStep float64) *EngineOptions {
	e.p.Set("timeStep", timeStep)
	return e
}

// UseHighPrecisionFloats returns the UseHighPrecisionFloats property of class EngineOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.engineoptions#usehighprecisionfloats
func (e *EngineOptions) UseHighPrecisionFloats() bool {
	retVal := e.p.Get("useHighPrecisionFloats")
	return retVal.Bool()
}

// SetUseHighPrecisionFloats sets the UseHighPrecisionFloats property of class EngineOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.engineoptions#usehighprecisionfloats
func (e *EngineOptions) SetUseHighPrecisionFloats(useHighPrecisionFloats bool) *EngineOptions {
	e.p.Set("useHighPrecisionFloats", useHighPrecisionFloats)
	return e
}
