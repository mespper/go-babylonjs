// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// WebXRControllerPointerSelection represents a babylon.js WebXRControllerPointerSelection.
// Handles pointer input automatically for the pointer of XR controllers
type WebXRControllerPointerSelection struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (w *WebXRControllerPointerSelection) JSObject() js.Value { return w.p }

// WebXRControllerPointerSelection returns a WebXRControllerPointerSelection JavaScript class.
func (ba *Babylon) WebXRControllerPointerSelection() *WebXRControllerPointerSelection {
	p := ba.ctx.Get("WebXRControllerPointerSelection")
	return WebXRControllerPointerSelectionFromJSObject(p, ba.ctx)
}

// WebXRControllerPointerSelectionFromJSObject returns a wrapped WebXRControllerPointerSelection JavaScript class.
func WebXRControllerPointerSelectionFromJSObject(p js.Value, ctx js.Value) *WebXRControllerPointerSelection {
	return &WebXRControllerPointerSelection{p: p, ctx: ctx}
}

// WebXRControllerPointerSelectionArrayToJSArray returns a JavaScript Array for the wrapped array.
func WebXRControllerPointerSelectionArrayToJSArray(array []*WebXRControllerPointerSelection) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewWebXRControllerPointerSelection returns a new WebXRControllerPointerSelection object.
//
// https://doc.babylonjs.com/api/classes/babylon.webxrcontrollerpointerselection
func (ba *Babylon) NewWebXRControllerPointerSelection(input *WebXRInput) *WebXRControllerPointerSelection {

	args := make([]interface{}, 0, 1+0)

	args = append(args, input.JSObject())

	p := ba.ctx.Get("WebXRControllerPointerSelection").New(args...)
	return WebXRControllerPointerSelectionFromJSObject(p, ba.ctx)
}
