// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// AnaglyphUniversalCamera represents a babylon.js AnaglyphUniversalCamera.
// Camera used to simulate anaglyphic rendering (based on UniversalCamera)
//
// See: http://doc.babylonjs.com/features/cameras#anaglyph-cameras
type AnaglyphUniversalCamera struct{ *UniversalCamera }

// JSObject returns the underlying js.Value.
func (a *AnaglyphUniversalCamera) JSObject() js.Value { return a.p }

// AnaglyphUniversalCamera returns a AnaglyphUniversalCamera JavaScript class.
func (ba *Babylon) AnaglyphUniversalCamera() *AnaglyphUniversalCamera {
	p := ba.ctx.Get("AnaglyphUniversalCamera")
	return AnaglyphUniversalCameraFromJSObject(p)
}

// AnaglyphUniversalCameraFromJSObject returns a wrapped AnaglyphUniversalCamera JavaScript class.
func AnaglyphUniversalCameraFromJSObject(p js.Value) *AnaglyphUniversalCamera {
	return &AnaglyphUniversalCamera{UniversalCameraFromJSObject(p)}
}

// NewAnaglyphUniversalCamera returns a new AnaglyphUniversalCamera object.
//
// https://doc.babylonjs.com/api/classes/babylon.anaglyphuniversalcamera
func (ba *Babylon) NewAnaglyphUniversalCamera(name string, position *Vector3, interaxialDistance float64, scene *Scene) *AnaglyphUniversalCamera {
	p := ba.ctx.Get("AnaglyphUniversalCamera").New(name, position.JSObject(), interaxialDistance, scene.JSObject())
	return AnaglyphUniversalCameraFromJSObject(p)
}

// TODO: methods
