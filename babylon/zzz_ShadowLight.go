// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ShadowLight represents a babylon.js ShadowLight.
// Base implementation IShadowLight
// It groups all the common behaviour in order to reduce dupplication and better follow the DRY pattern.
type ShadowLight struct{ *Light }

// JSObject returns the underlying js.Value.
func (s *ShadowLight) JSObject() js.Value { return s.p }

// ShadowLight returns a ShadowLight JavaScript class.
func (ba *Babylon) ShadowLight() *ShadowLight {
	p := ba.ctx.Get("ShadowLight")
	return ShadowLightFromJSObject(p)
}

// ShadowLightFromJSObject returns a wrapped ShadowLight JavaScript class.
func ShadowLightFromJSObject(p js.Value) *ShadowLight {
	return &ShadowLight{LightFromJSObject(p)}
}

// NewShadowLight returns a new ShadowLight object.
//
// https://doc.babylonjs.com/api/classes/babylon.shadowlight
func (ba *Babylon) NewShadowLight(name string, scene *Scene) *ShadowLight {
	p := ba.ctx.Get("ShadowLight").New(name, scene.JSObject())
	return ShadowLightFromJSObject(p)
}

// TODO: methods
