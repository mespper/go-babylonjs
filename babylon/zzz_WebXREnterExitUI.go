// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// WebXREnterExitUI represents a babylon.js WebXREnterExitUI.
// UI to allow the user to enter/exit XR mode
type WebXREnterExitUI struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (w *WebXREnterExitUI) JSObject() js.Value { return w.p }

// WebXREnterExitUI returns a WebXREnterExitUI JavaScript class.
func (ba *Babylon) WebXREnterExitUI() *WebXREnterExitUI {
	p := ba.ctx.Get("WebXREnterExitUI")
	return WebXREnterExitUIFromJSObject(p, ba.ctx)
}

// WebXREnterExitUIFromJSObject returns a wrapped WebXREnterExitUI JavaScript class.
func WebXREnterExitUIFromJSObject(p js.Value, ctx js.Value) *WebXREnterExitUI {
	return &WebXREnterExitUI{p: p, ctx: ctx}
}

// WebXREnterExitUIArrayToJSArray returns a JavaScript Array for the wrapped array.
func WebXREnterExitUIArrayToJSArray(array []*WebXREnterExitUI) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// CreateAsync calls the CreateAsync method on the WebXREnterExitUI object.
//
// https://doc.babylonjs.com/api/classes/babylon.webxrenterexitui#createasync
func (w *WebXREnterExitUI) CreateAsync(scene *Scene, helper *WebXRExperienceHelper, options *WebXREnterExitUIOptions) *Promise {

	args := make([]interface{}, 0, 3+0)

	if scene == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, scene.JSObject())
	}

	if helper == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, helper.JSObject())
	}

	if options == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, options.JSObject())
	}

	retVal := w.p.Call("CreateAsync", args...)
	return PromiseFromJSObject(retVal, w.ctx)
}

// Dispose calls the Dispose method on the WebXREnterExitUI object.
//
// https://doc.babylonjs.com/api/classes/babylon.webxrenterexitui#dispose
func (w *WebXREnterExitUI) Dispose() {

	w.p.Call("dispose")
}

// ActiveButtonChangedObservable returns the ActiveButtonChangedObservable property of class WebXREnterExitUI.
//
// https://doc.babylonjs.com/api/classes/babylon.webxrenterexitui#activebuttonchangedobservable
func (w *WebXREnterExitUI) ActiveButtonChangedObservable() *Observable {
	retVal := w.p.Get("activeButtonChangedObservable")
	return ObservableFromJSObject(retVal, w.ctx)
}

// SetActiveButtonChangedObservable sets the ActiveButtonChangedObservable property of class WebXREnterExitUI.
//
// https://doc.babylonjs.com/api/classes/babylon.webxrenterexitui#activebuttonchangedobservable
func (w *WebXREnterExitUI) SetActiveButtonChangedObservable(activeButtonChangedObservable *Observable) *WebXREnterExitUI {
	w.p.Set("activeButtonChangedObservable", activeButtonChangedObservable.JSObject())
	return w
}
