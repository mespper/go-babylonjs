// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// WebVROptions represents a babylon.js WebVROptions.
// Set of options to customize the webVRCamera
type WebVROptions struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (w *WebVROptions) JSObject() js.Value { return w.p }

// WebVROptions returns a WebVROptions JavaScript class.
func (ba *Babylon) WebVROptions() *WebVROptions {
	p := ba.ctx.Get("WebVROptions")
	return WebVROptionsFromJSObject(p, ba.ctx)
}

// WebVROptionsFromJSObject returns a wrapped WebVROptions JavaScript class.
func WebVROptionsFromJSObject(p js.Value, ctx js.Value) *WebVROptions {
	return &WebVROptions{p: p, ctx: ctx}
}

// WebVROptionsArrayToJSArray returns a JavaScript Array for the wrapped array.
func WebVROptionsArrayToJSArray(array []*WebVROptions) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

/*

// ControllerMeshes returns the ControllerMeshes property of class WebVROptions.
//
// https://doc.babylonjs.com/api/classes/babylon.webvroptions#controllermeshes
func (w *WebVROptions) ControllerMeshes(controllerMeshes bool) *WebVROptions {
	p := ba.ctx.Get("WebVROptions").New(controllerMeshes)
	return WebVROptionsFromJSObject(p, ba.ctx)
}

// SetControllerMeshes sets the ControllerMeshes property of class WebVROptions.
//
// https://doc.babylonjs.com/api/classes/babylon.webvroptions#controllermeshes
func (w *WebVROptions) SetControllerMeshes(controllerMeshes bool) *WebVROptions {
	p := ba.ctx.Get("WebVROptions").New(controllerMeshes)
	return WebVROptionsFromJSObject(p, ba.ctx)
}

// CustomVRButton returns the CustomVRButton property of class WebVROptions.
//
// https://doc.babylonjs.com/api/classes/babylon.webvroptions#customvrbutton
func (w *WebVROptions) CustomVRButton(customVRButton *HTMLButtonElement) *WebVROptions {
	p := ba.ctx.Get("WebVROptions").New(customVRButton.JSObject())
	return WebVROptionsFromJSObject(p, ba.ctx)
}

// SetCustomVRButton sets the CustomVRButton property of class WebVROptions.
//
// https://doc.babylonjs.com/api/classes/babylon.webvroptions#customvrbutton
func (w *WebVROptions) SetCustomVRButton(customVRButton *HTMLButtonElement) *WebVROptions {
	p := ba.ctx.Get("WebVROptions").New(customVRButton.JSObject())
	return WebVROptionsFromJSObject(p, ba.ctx)
}

// DefaultHeight returns the DefaultHeight property of class WebVROptions.
//
// https://doc.babylonjs.com/api/classes/babylon.webvroptions#defaultheight
func (w *WebVROptions) DefaultHeight(defaultHeight float64) *WebVROptions {
	p := ba.ctx.Get("WebVROptions").New(defaultHeight)
	return WebVROptionsFromJSObject(p, ba.ctx)
}

// SetDefaultHeight sets the DefaultHeight property of class WebVROptions.
//
// https://doc.babylonjs.com/api/classes/babylon.webvroptions#defaultheight
func (w *WebVROptions) SetDefaultHeight(defaultHeight float64) *WebVROptions {
	p := ba.ctx.Get("WebVROptions").New(defaultHeight)
	return WebVROptionsFromJSObject(p, ba.ctx)
}

// DefaultLightingOnControllers returns the DefaultLightingOnControllers property of class WebVROptions.
//
// https://doc.babylonjs.com/api/classes/babylon.webvroptions#defaultlightingoncontrollers
func (w *WebVROptions) DefaultLightingOnControllers(defaultLightingOnControllers bool) *WebVROptions {
	p := ba.ctx.Get("WebVROptions").New(defaultLightingOnControllers)
	return WebVROptionsFromJSObject(p, ba.ctx)
}

// SetDefaultLightingOnControllers sets the DefaultLightingOnControllers property of class WebVROptions.
//
// https://doc.babylonjs.com/api/classes/babylon.webvroptions#defaultlightingoncontrollers
func (w *WebVROptions) SetDefaultLightingOnControllers(defaultLightingOnControllers bool) *WebVROptions {
	p := ba.ctx.Get("WebVROptions").New(defaultLightingOnControllers)
	return WebVROptionsFromJSObject(p, ba.ctx)
}

// DisplayName returns the DisplayName property of class WebVROptions.
//
// https://doc.babylonjs.com/api/classes/babylon.webvroptions#displayname
func (w *WebVROptions) DisplayName(displayName string) *WebVROptions {
	p := ba.ctx.Get("WebVROptions").New(displayName)
	return WebVROptionsFromJSObject(p, ba.ctx)
}

// SetDisplayName sets the DisplayName property of class WebVROptions.
//
// https://doc.babylonjs.com/api/classes/babylon.webvroptions#displayname
func (w *WebVROptions) SetDisplayName(displayName string) *WebVROptions {
	p := ba.ctx.Get("WebVROptions").New(displayName)
	return WebVROptionsFromJSObject(p, ba.ctx)
}

// PositionScale returns the PositionScale property of class WebVROptions.
//
// https://doc.babylonjs.com/api/classes/babylon.webvroptions#positionscale
func (w *WebVROptions) PositionScale(positionScale float64) *WebVROptions {
	p := ba.ctx.Get("WebVROptions").New(positionScale)
	return WebVROptionsFromJSObject(p, ba.ctx)
}

// SetPositionScale sets the PositionScale property of class WebVROptions.
//
// https://doc.babylonjs.com/api/classes/babylon.webvroptions#positionscale
func (w *WebVROptions) SetPositionScale(positionScale float64) *WebVROptions {
	p := ba.ctx.Get("WebVROptions").New(positionScale)
	return WebVROptionsFromJSObject(p, ba.ctx)
}

// RayLength returns the RayLength property of class WebVROptions.
//
// https://doc.babylonjs.com/api/classes/babylon.webvroptions#raylength
func (w *WebVROptions) RayLength(rayLength float64) *WebVROptions {
	p := ba.ctx.Get("WebVROptions").New(rayLength)
	return WebVROptionsFromJSObject(p, ba.ctx)
}

// SetRayLength sets the RayLength property of class WebVROptions.
//
// https://doc.babylonjs.com/api/classes/babylon.webvroptions#raylength
func (w *WebVROptions) SetRayLength(rayLength float64) *WebVROptions {
	p := ba.ctx.Get("WebVROptions").New(rayLength)
	return WebVROptionsFromJSObject(p, ba.ctx)
}

// TrackPosition returns the TrackPosition property of class WebVROptions.
//
// https://doc.babylonjs.com/api/classes/babylon.webvroptions#trackposition
func (w *WebVROptions) TrackPosition(trackPosition bool) *WebVROptions {
	p := ba.ctx.Get("WebVROptions").New(trackPosition)
	return WebVROptionsFromJSObject(p, ba.ctx)
}

// SetTrackPosition sets the TrackPosition property of class WebVROptions.
//
// https://doc.babylonjs.com/api/classes/babylon.webvroptions#trackposition
func (w *WebVROptions) SetTrackPosition(trackPosition bool) *WebVROptions {
	p := ba.ctx.Get("WebVROptions").New(trackPosition)
	return WebVROptionsFromJSObject(p, ba.ctx)
}

// UseCustomVRButton returns the UseCustomVRButton property of class WebVROptions.
//
// https://doc.babylonjs.com/api/classes/babylon.webvroptions#usecustomvrbutton
func (w *WebVROptions) UseCustomVRButton(useCustomVRButton bool) *WebVROptions {
	p := ba.ctx.Get("WebVROptions").New(useCustomVRButton)
	return WebVROptionsFromJSObject(p, ba.ctx)
}

// SetUseCustomVRButton sets the UseCustomVRButton property of class WebVROptions.
//
// https://doc.babylonjs.com/api/classes/babylon.webvroptions#usecustomvrbutton
func (w *WebVROptions) SetUseCustomVRButton(useCustomVRButton bool) *WebVROptions {
	p := ba.ctx.Get("WebVROptions").New(useCustomVRButton)
	return WebVROptionsFromJSObject(p, ba.ctx)
}

// UseMultiview returns the UseMultiview property of class WebVROptions.
//
// https://doc.babylonjs.com/api/classes/babylon.webvroptions#usemultiview
func (w *WebVROptions) UseMultiview(useMultiview bool) *WebVROptions {
	p := ba.ctx.Get("WebVROptions").New(useMultiview)
	return WebVROptionsFromJSObject(p, ba.ctx)
}

// SetUseMultiview sets the UseMultiview property of class WebVROptions.
//
// https://doc.babylonjs.com/api/classes/babylon.webvroptions#usemultiview
func (w *WebVROptions) SetUseMultiview(useMultiview bool) *WebVROptions {
	p := ba.ctx.Get("WebVROptions").New(useMultiview)
	return WebVROptionsFromJSObject(p, ba.ctx)
}

*/