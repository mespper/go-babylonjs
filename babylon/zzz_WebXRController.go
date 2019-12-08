// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// WebXRController represents a babylon.js WebXRController.
// Represents an XR input
type WebXRController struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (w *WebXRController) JSObject() js.Value { return w.p }

// WebXRController returns a WebXRController JavaScript class.
func (ba *Babylon) WebXRController() *WebXRController {
	p := ba.ctx.Get("WebXRController")
	return WebXRControllerFromJSObject(p, ba.ctx)
}

// WebXRControllerFromJSObject returns a wrapped WebXRController JavaScript class.
func WebXRControllerFromJSObject(p js.Value, ctx js.Value) *WebXRController {
	return &WebXRController{p: p, ctx: ctx}
}

// WebXRControllerArrayToJSArray returns a JavaScript Array for the wrapped array.
func WebXRControllerArrayToJSArray(array []*WebXRController) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewWebXRControllerOpts contains optional parameters for NewWebXRController.
type NewWebXRControllerOpts struct {
	ParentContainer *AbstractMesh
}

// NewWebXRController returns a new WebXRController object.
//
// https://doc.babylonjs.com/api/classes/babylon.webxrcontroller
func (ba *Babylon) NewWebXRController(scene *Scene, inputSource js.Value, opts *NewWebXRControllerOpts) *WebXRController {
	if opts == nil {
		opts = &NewWebXRControllerOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, scene.JSObject())
	args = append(args, inputSource)

	if opts.ParentContainer == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.ParentContainer.JSObject())
	}

	p := ba.ctx.Get("WebXRController").New(args...)
	return WebXRControllerFromJSObject(p, ba.ctx)
}

// Dispose calls the Dispose method on the WebXRController object.
//
// https://doc.babylonjs.com/api/classes/babylon.webxrcontroller#dispose
func (w *WebXRController) Dispose() {

	w.p.Call("dispose")
}

// GetScene calls the GetScene method on the WebXRController object.
//
// https://doc.babylonjs.com/api/classes/babylon.webxrcontroller#getscene
func (w *WebXRController) GetScene() *Scene {

	retVal := w.p.Call("getScene")
	return SceneFromJSObject(retVal, w.ctx)
}

// GetWorldPointerRayToRef calls the GetWorldPointerRayToRef method on the WebXRController object.
//
// https://doc.babylonjs.com/api/classes/babylon.webxrcontroller#getworldpointerraytoref
func (w *WebXRController) GetWorldPointerRayToRef(result *Ray) {

	args := make([]interface{}, 0, 1+0)

	if result == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, result.JSObject())
	}

	w.p.Call("getWorldPointerRayToRef", args...)
}

// UpdateFromXRFrame calls the UpdateFromXRFrame method on the WebXRController object.
//
// https://doc.babylonjs.com/api/classes/babylon.webxrcontroller#updatefromxrframe
func (w *WebXRController) UpdateFromXRFrame(xrFrame js.Value, referenceSpace js.Value) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, xrFrame)

	args = append(args, referenceSpace)

	w.p.Call("updateFromXRFrame", args...)
}

// GamepadController returns the GamepadController property of class WebXRController.
//
// https://doc.babylonjs.com/api/classes/babylon.webxrcontroller#gamepadcontroller
func (w *WebXRController) GamepadController() *WebVRController {
	retVal := w.p.Get("gamepadController")
	return WebVRControllerFromJSObject(retVal, w.ctx)
}

// SetGamepadController sets the GamepadController property of class WebXRController.
//
// https://doc.babylonjs.com/api/classes/babylon.webxrcontroller#gamepadcontroller
func (w *WebXRController) SetGamepadController(gamepadController *WebVRController) *WebXRController {
	w.p.Set("gamepadController", gamepadController.JSObject())
	return w
}

// Grip returns the Grip property of class WebXRController.
//
// https://doc.babylonjs.com/api/classes/babylon.webxrcontroller#grip
func (w *WebXRController) Grip() *AbstractMesh {
	retVal := w.p.Get("grip")
	return AbstractMeshFromJSObject(retVal, w.ctx)
}

// SetGrip sets the Grip property of class WebXRController.
//
// https://doc.babylonjs.com/api/classes/babylon.webxrcontroller#grip
func (w *WebXRController) SetGrip(grip *AbstractMesh) *WebXRController {
	w.p.Set("grip", grip.JSObject())
	return w
}

// InputSource returns the InputSource property of class WebXRController.
//
// https://doc.babylonjs.com/api/classes/babylon.webxrcontroller#inputsource
func (w *WebXRController) InputSource() js.Value {
	retVal := w.p.Get("inputSource")
	return retVal
}

// SetInputSource sets the InputSource property of class WebXRController.
//
// https://doc.babylonjs.com/api/classes/babylon.webxrcontroller#inputsource
func (w *WebXRController) SetInputSource(inputSource js.Value) *WebXRController {
	w.p.Set("inputSource", inputSource)
	return w
}

// OnDisposeObservable returns the OnDisposeObservable property of class WebXRController.
//
// https://doc.babylonjs.com/api/classes/babylon.webxrcontroller#ondisposeobservable
func (w *WebXRController) OnDisposeObservable() *Observable {
	retVal := w.p.Get("onDisposeObservable")
	return ObservableFromJSObject(retVal, w.ctx)
}

// SetOnDisposeObservable sets the OnDisposeObservable property of class WebXRController.
//
// https://doc.babylonjs.com/api/classes/babylon.webxrcontroller#ondisposeobservable
func (w *WebXRController) SetOnDisposeObservable(onDisposeObservable *Observable) *WebXRController {
	w.p.Set("onDisposeObservable", onDisposeObservable.JSObject())
	return w
}

// Pointer returns the Pointer property of class WebXRController.
//
// https://doc.babylonjs.com/api/classes/babylon.webxrcontroller#pointer
func (w *WebXRController) Pointer() *AbstractMesh {
	retVal := w.p.Get("pointer")
	return AbstractMeshFromJSObject(retVal, w.ctx)
}

// SetPointer sets the Pointer property of class WebXRController.
//
// https://doc.babylonjs.com/api/classes/babylon.webxrcontroller#pointer
func (w *WebXRController) SetPointer(pointer *AbstractMesh) *WebXRController {
	w.p.Set("pointer", pointer.JSObject())
	return w
}
