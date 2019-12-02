// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// FlyCameraInputsManager represents a babylon.js FlyCameraInputsManager.
// Default Inputs manager for the FlyCamera.
// It groups all the default supported inputs for ease of use.
//
// See: http://doc.babylonjs.com/how_to/customizing_camera_inputs
type FlyCameraInputsManager struct {
	*CameraInputsManager
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (f *FlyCameraInputsManager) JSObject() js.Value { return f.p }

// FlyCameraInputsManager returns a FlyCameraInputsManager JavaScript class.
func (ba *Babylon) FlyCameraInputsManager() *FlyCameraInputsManager {
	p := ba.ctx.Get("FlyCameraInputsManager")
	return FlyCameraInputsManagerFromJSObject(p, ba.ctx)
}

// FlyCameraInputsManagerFromJSObject returns a wrapped FlyCameraInputsManager JavaScript class.
func FlyCameraInputsManagerFromJSObject(p js.Value, ctx js.Value) *FlyCameraInputsManager {
	return &FlyCameraInputsManager{CameraInputsManager: CameraInputsManagerFromJSObject(p, ctx), ctx: ctx}
}

// FlyCameraInputsManagerArrayToJSArray returns a JavaScript Array for the wrapped array.
func FlyCameraInputsManagerArrayToJSArray(array []*FlyCameraInputsManager) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewFlyCameraInputsManager returns a new FlyCameraInputsManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamerainputsmanager
func (ba *Babylon) NewFlyCameraInputsManager(camera *FlyCamera) *FlyCameraInputsManager {

	args := make([]interface{}, 0, 1+0)

	args = append(args, camera.JSObject())

	p := ba.ctx.Get("FlyCameraInputsManager").New(args...)
	return FlyCameraInputsManagerFromJSObject(p, ba.ctx)
}

// AddKeyboard calls the AddKeyboard method on the FlyCameraInputsManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamerainputsmanager#addkeyboard
func (f *FlyCameraInputsManager) AddKeyboard() *FlyCameraInputsManager {

	retVal := f.p.Call("addKeyboard")
	return FlyCameraInputsManagerFromJSObject(retVal, f.ctx)
}

// FlyCameraInputsManagerAddMouseOpts contains optional parameters for FlyCameraInputsManager.AddMouse.
type FlyCameraInputsManagerAddMouseOpts struct {
	TouchEnabled *bool
}

// AddMouse calls the AddMouse method on the FlyCameraInputsManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.flycamerainputsmanager#addmouse
func (f *FlyCameraInputsManager) AddMouse(opts *FlyCameraInputsManagerAddMouseOpts) *FlyCameraInputsManager {
	if opts == nil {
		opts = &FlyCameraInputsManagerAddMouseOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.TouchEnabled == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.TouchEnabled)
	}

	retVal := f.p.Call("addMouse", args...)
	return FlyCameraInputsManagerFromJSObject(retVal, f.ctx)
}
