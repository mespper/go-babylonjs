// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ArcRotateCameraVRDeviceOrientationInput represents a babylon.js ArcRotateCameraVRDeviceOrientationInput.
// Manage the device orientation inputs (gyroscope) to control an arc rotate camera.
//
// See: http://doc.babylonjs.com/how_to/customizing_camera_inputs
type ArcRotateCameraVRDeviceOrientationInput struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (a *ArcRotateCameraVRDeviceOrientationInput) JSObject() js.Value { return a.p }

// ArcRotateCameraVRDeviceOrientationInput returns a ArcRotateCameraVRDeviceOrientationInput JavaScript class.
func (ba *Babylon) ArcRotateCameraVRDeviceOrientationInput() *ArcRotateCameraVRDeviceOrientationInput {
	p := ba.ctx.Get("ArcRotateCameraVRDeviceOrientationInput")
	return ArcRotateCameraVRDeviceOrientationInputFromJSObject(p, ba.ctx)
}

// ArcRotateCameraVRDeviceOrientationInputFromJSObject returns a wrapped ArcRotateCameraVRDeviceOrientationInput JavaScript class.
func ArcRotateCameraVRDeviceOrientationInputFromJSObject(p js.Value, ctx js.Value) *ArcRotateCameraVRDeviceOrientationInput {
	return &ArcRotateCameraVRDeviceOrientationInput{p: p, ctx: ctx}
}

// ArcRotateCameraVRDeviceOrientationInputArrayToJSArray returns a JavaScript Array for the wrapped array.
func ArcRotateCameraVRDeviceOrientationInputArrayToJSArray(array []*ArcRotateCameraVRDeviceOrientationInput) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewArcRotateCameraVRDeviceOrientationInput returns a new ArcRotateCameraVRDeviceOrientationInput object.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecameravrdeviceorientationinput
func (ba *Babylon) NewArcRotateCameraVRDeviceOrientationInput() *ArcRotateCameraVRDeviceOrientationInput {

	args := make([]interface{}, 0, 0+0)

	p := ba.ctx.Get("ArcRotateCameraVRDeviceOrientationInput").New(args...)
	return ArcRotateCameraVRDeviceOrientationInputFromJSObject(p, ba.ctx)
}

// ArcRotateCameraVRDeviceOrientationInputAttachControlOpts contains optional parameters for ArcRotateCameraVRDeviceOrientationInput.AttachControl.
type ArcRotateCameraVRDeviceOrientationInputAttachControlOpts struct {
	NoPreventDefault *bool
}

// AttachControl calls the AttachControl method on the ArcRotateCameraVRDeviceOrientationInput object.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecameravrdeviceorientationinput#attachcontrol
func (a *ArcRotateCameraVRDeviceOrientationInput) AttachControl(element js.Value, opts *ArcRotateCameraVRDeviceOrientationInputAttachControlOpts) {
	if opts == nil {
		opts = &ArcRotateCameraVRDeviceOrientationInputAttachControlOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, element)

	if opts.NoPreventDefault == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.NoPreventDefault)
	}

	a.p.Call("attachControl", args...)
}

// CheckInputs calls the CheckInputs method on the ArcRotateCameraVRDeviceOrientationInput object.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecameravrdeviceorientationinput#checkinputs
func (a *ArcRotateCameraVRDeviceOrientationInput) CheckInputs() {

	a.p.Call("checkInputs")
}

// DetachControl calls the DetachControl method on the ArcRotateCameraVRDeviceOrientationInput object.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecameravrdeviceorientationinput#detachcontrol
func (a *ArcRotateCameraVRDeviceOrientationInput) DetachControl(element js.Value) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, element)

	a.p.Call("detachControl", args...)
}

// GetClassName calls the GetClassName method on the ArcRotateCameraVRDeviceOrientationInput object.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecameravrdeviceorientationinput#getclassname
func (a *ArcRotateCameraVRDeviceOrientationInput) GetClassName() string {

	retVal := a.p.Call("getClassName")
	return retVal.String()
}

// GetSimpleName calls the GetSimpleName method on the ArcRotateCameraVRDeviceOrientationInput object.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecameravrdeviceorientationinput#getsimplename
func (a *ArcRotateCameraVRDeviceOrientationInput) GetSimpleName() string {

	retVal := a.p.Call("getSimpleName")
	return retVal.String()
}

// AlphaCorrection returns the AlphaCorrection property of class ArcRotateCameraVRDeviceOrientationInput.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecameravrdeviceorientationinput#alphacorrection
func (a *ArcRotateCameraVRDeviceOrientationInput) AlphaCorrection() float64 {
	retVal := a.p.Get("alphaCorrection")
	return retVal.Float()
}

// SetAlphaCorrection sets the AlphaCorrection property of class ArcRotateCameraVRDeviceOrientationInput.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecameravrdeviceorientationinput#alphacorrection
func (a *ArcRotateCameraVRDeviceOrientationInput) SetAlphaCorrection(alphaCorrection float64) *ArcRotateCameraVRDeviceOrientationInput {
	a.p.Set("alphaCorrection", alphaCorrection)
	return a
}

// Camera returns the Camera property of class ArcRotateCameraVRDeviceOrientationInput.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecameravrdeviceorientationinput#camera
func (a *ArcRotateCameraVRDeviceOrientationInput) Camera() *ArcRotateCamera {
	retVal := a.p.Get("camera")
	return ArcRotateCameraFromJSObject(retVal, a.ctx)
}

// SetCamera sets the Camera property of class ArcRotateCameraVRDeviceOrientationInput.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecameravrdeviceorientationinput#camera
func (a *ArcRotateCameraVRDeviceOrientationInput) SetCamera(camera *ArcRotateCamera) *ArcRotateCameraVRDeviceOrientationInput {
	a.p.Set("camera", camera.JSObject())
	return a
}

// GammaCorrection returns the GammaCorrection property of class ArcRotateCameraVRDeviceOrientationInput.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecameravrdeviceorientationinput#gammacorrection
func (a *ArcRotateCameraVRDeviceOrientationInput) GammaCorrection() float64 {
	retVal := a.p.Get("gammaCorrection")
	return retVal.Float()
}

// SetGammaCorrection sets the GammaCorrection property of class ArcRotateCameraVRDeviceOrientationInput.
//
// https://doc.babylonjs.com/api/classes/babylon.arcrotatecameravrdeviceorientationinput#gammacorrection
func (a *ArcRotateCameraVRDeviceOrientationInput) SetGammaCorrection(gammaCorrection float64) *ArcRotateCameraVRDeviceOrientationInput {
	a.p.Set("gammaCorrection", gammaCorrection)
	return a
}
