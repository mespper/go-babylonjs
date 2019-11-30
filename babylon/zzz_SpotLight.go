// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// SpotLight represents a babylon.js SpotLight.
// A spot light is defined by a position, a direction, an angle, and an exponent.
// These values define a cone of light starting from the position, emitting toward the direction.
// The angle, in radians, defines the size (field of illumination) of the spotlight&amp;#39;s conical beam,
// and the exponent defines the speed of the decay of the light with distance (reach).
// Documentation: &lt;a href=&#34;https://doc.babylonjs.com/babylon101/lights&#34;&gt;https://doc.babylonjs.com/babylon101/lights&lt;/a&gt;
type SpotLight struct{ *ShadowLight }

// JSObject returns the underlying js.Value.
func (s *SpotLight) JSObject() js.Value { return s.p }

// SpotLight returns a SpotLight JavaScript class.
func (ba *Babylon) SpotLight() *SpotLight {
	p := ba.ctx.Get("SpotLight")
	return SpotLightFromJSObject(p)
}

// SpotLightFromJSObject returns a wrapped SpotLight JavaScript class.
func SpotLightFromJSObject(p js.Value) *SpotLight {
	return &SpotLight{ShadowLightFromJSObject(p)}
}

// NewSpotLight returns a new SpotLight object.
//
// https://doc.babylonjs.com/api/classes/babylon.spotlight
func (ba *Babylon) NewSpotLight(name string, position *Vector3, direction *Vector3, angle float64, exponent float64, scene *Scene) *SpotLight {
	p := ba.ctx.Get("SpotLight").New(name, position.JSObject(), direction.JSObject(), angle, exponent, scene.JSObject())
	return SpotLightFromJSObject(p)
}

// TODO: methods
