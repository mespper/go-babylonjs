// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// VRDeviceOrientationGamepadCamera represents a babylon.js VRDeviceOrientationGamepadCamera.
// Camera used to simulate VR rendering (based on VRDeviceOrientationFreeCamera)
//
// See: http://doc.babylonjs.com/babylon101/cameras#vr-device-orientation-cameras
type VRDeviceOrientationGamepadCamera struct{ *VRDeviceOrientationFreeCamera }

// JSObject returns the underlying js.Value.
func (v *VRDeviceOrientationGamepadCamera) JSObject() js.Value { return v.p }

// VRDeviceOrientationGamepadCamera returns a VRDeviceOrientationGamepadCamera JavaScript class.
func (ba *Babylon) VRDeviceOrientationGamepadCamera() *VRDeviceOrientationGamepadCamera {
	p := ba.ctx.Get("VRDeviceOrientationGamepadCamera")
	return VRDeviceOrientationGamepadCameraFromJSObject(p)
}

// VRDeviceOrientationGamepadCameraFromJSObject returns a wrapped VRDeviceOrientationGamepadCamera JavaScript class.
func VRDeviceOrientationGamepadCameraFromJSObject(p js.Value) *VRDeviceOrientationGamepadCamera {
	return &VRDeviceOrientationGamepadCamera{VRDeviceOrientationFreeCameraFromJSObject(p)}
}

// NewVRDeviceOrientationGamepadCameraOpts contains optional parameters for NewVRDeviceOrientationGamepadCamera.
type NewVRDeviceOrientationGamepadCameraOpts struct {
	CompensateDistortion *JSBool

	VrCameraMetrics *VRCameraMetrics
}

// NewVRDeviceOrientationGamepadCamera returns a new VRDeviceOrientationGamepadCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdeviceorientationgamepadcamera
func (ba *Babylon) NewVRDeviceOrientationGamepadCamera(name string, position *Vector3, scene *Scene, opts *NewVRDeviceOrientationGamepadCameraOpts) *VRDeviceOrientationGamepadCamera {
	if opts == nil {
		opts = &NewVRDeviceOrientationGamepadCameraOpts{}
	}

	p := ba.ctx.Get("VRDeviceOrientationGamepadCamera").New(name, position.JSObject(), scene.JSObject(), opts.CompensateDistortion.JSObject(), opts.VrCameraMetrics.JSObject())
	return VRDeviceOrientationGamepadCameraFromJSObject(p)
}

// TODO: methods
