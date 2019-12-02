// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// IMaterialCompilationOptions represents a babylon.js IMaterialCompilationOptions.
// Options for compiling materials.
type IMaterialCompilationOptions struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (i *IMaterialCompilationOptions) JSObject() js.Value { return i.p }

// IMaterialCompilationOptions returns a IMaterialCompilationOptions JavaScript class.
func (ba *Babylon) IMaterialCompilationOptions() *IMaterialCompilationOptions {
	p := ba.ctx.Get("IMaterialCompilationOptions")
	return IMaterialCompilationOptionsFromJSObject(p, ba.ctx)
}

// IMaterialCompilationOptionsFromJSObject returns a wrapped IMaterialCompilationOptions JavaScript class.
func IMaterialCompilationOptionsFromJSObject(p js.Value, ctx js.Value) *IMaterialCompilationOptions {
	return &IMaterialCompilationOptions{p: p, ctx: ctx}
}

// IMaterialCompilationOptionsArrayToJSArray returns a JavaScript Array for the wrapped array.
func IMaterialCompilationOptionsArrayToJSArray(array []*IMaterialCompilationOptions) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// ClipPlane returns the ClipPlane property of class IMaterialCompilationOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.imaterialcompilationoptions#clipplane
func (i *IMaterialCompilationOptions) ClipPlane() bool {
	retVal := i.p.Get("clipPlane")
	return retVal.Bool()
}

// SetClipPlane sets the ClipPlane property of class IMaterialCompilationOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.imaterialcompilationoptions#clipplane
func (i *IMaterialCompilationOptions) SetClipPlane(clipPlane bool) *IMaterialCompilationOptions {
	i.p.Set("clipPlane", clipPlane)
	return i
}

// UseInstances returns the UseInstances property of class IMaterialCompilationOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.imaterialcompilationoptions#useinstances
func (i *IMaterialCompilationOptions) UseInstances() bool {
	retVal := i.p.Get("useInstances")
	return retVal.Bool()
}

// SetUseInstances sets the UseInstances property of class IMaterialCompilationOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.imaterialcompilationoptions#useinstances
func (i *IMaterialCompilationOptions) SetUseInstances(useInstances bool) *IMaterialCompilationOptions {
	i.p.Set("useInstances", useInstances)
	return i
}
