// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// PhysicsRadialExplosionEventOptions represents a babylon.js PhysicsRadialExplosionEventOptions.
// Options fot the radial explosion event
//
// See: https://doc.babylonjs.com/how_to/using_the_physics_engine#further-functionality-of-the-impostor-class
type PhysicsRadialExplosionEventOptions struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (p *PhysicsRadialExplosionEventOptions) JSObject() js.Value { return p.p }

// PhysicsRadialExplosionEventOptions returns a PhysicsRadialExplosionEventOptions JavaScript class.
func (ba *Babylon) PhysicsRadialExplosionEventOptions() *PhysicsRadialExplosionEventOptions {
	p := ba.ctx.Get("PhysicsRadialExplosionEventOptions")
	return PhysicsRadialExplosionEventOptionsFromJSObject(p, ba.ctx)
}

// PhysicsRadialExplosionEventOptionsFromJSObject returns a wrapped PhysicsRadialExplosionEventOptions JavaScript class.
func PhysicsRadialExplosionEventOptionsFromJSObject(p js.Value, ctx js.Value) *PhysicsRadialExplosionEventOptions {
	return &PhysicsRadialExplosionEventOptions{p: p, ctx: ctx}
}

// PhysicsRadialExplosionEventOptionsArrayToJSArray returns a JavaScript Array for the wrapped array.
func PhysicsRadialExplosionEventOptionsArrayToJSArray(array []*PhysicsRadialExplosionEventOptions) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// AffectedImpostorsCallback returns the AffectedImpostorsCallback property of class PhysicsRadialExplosionEventOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.physicsradialexplosioneventoptions#affectedimpostorscallback
func (p *PhysicsRadialExplosionEventOptions) AffectedImpostorsCallback() js.Value {
	retVal := p.p.Get("affectedImpostorsCallback")
	return retVal
}

// SetAffectedImpostorsCallback sets the AffectedImpostorsCallback property of class PhysicsRadialExplosionEventOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.physicsradialexplosioneventoptions#affectedimpostorscallback
func (p *PhysicsRadialExplosionEventOptions) SetAffectedImpostorsCallback(affectedImpostorsCallback JSFunc) *PhysicsRadialExplosionEventOptions {
	p.p.Set("affectedImpostorsCallback", js.FuncOf(affectedImpostorsCallback))
	return p
}

// Falloff returns the Falloff property of class PhysicsRadialExplosionEventOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.physicsradialexplosioneventoptions#falloff
func (p *PhysicsRadialExplosionEventOptions) Falloff() js.Value {
	retVal := p.p.Get("falloff")
	return retVal
}

// SetFalloff sets the Falloff property of class PhysicsRadialExplosionEventOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.physicsradialexplosioneventoptions#falloff
func (p *PhysicsRadialExplosionEventOptions) SetFalloff(falloff js.Value) *PhysicsRadialExplosionEventOptions {
	p.p.Set("falloff", falloff)
	return p
}

// Radius returns the Radius property of class PhysicsRadialExplosionEventOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.physicsradialexplosioneventoptions#radius
func (p *PhysicsRadialExplosionEventOptions) Radius() float64 {
	retVal := p.p.Get("radius")
	return retVal.Float()
}

// SetRadius sets the Radius property of class PhysicsRadialExplosionEventOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.physicsradialexplosioneventoptions#radius
func (p *PhysicsRadialExplosionEventOptions) SetRadius(radius float64) *PhysicsRadialExplosionEventOptions {
	p.p.Set("radius", radius)
	return p
}

// Sphere returns the Sphere property of class PhysicsRadialExplosionEventOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.physicsradialexplosioneventoptions#sphere
func (p *PhysicsRadialExplosionEventOptions) Sphere() js.Value {
	retVal := p.p.Get("sphere")
	return retVal
}

// SetSphere sets the Sphere property of class PhysicsRadialExplosionEventOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.physicsradialexplosioneventoptions#sphere
func (p *PhysicsRadialExplosionEventOptions) SetSphere(sphere js.Value) *PhysicsRadialExplosionEventOptions {
	p.p.Set("sphere", sphere)
	return p
}

// Strength returns the Strength property of class PhysicsRadialExplosionEventOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.physicsradialexplosioneventoptions#strength
func (p *PhysicsRadialExplosionEventOptions) Strength() float64 {
	retVal := p.p.Get("strength")
	return retVal.Float()
}

// SetStrength sets the Strength property of class PhysicsRadialExplosionEventOptions.
//
// https://doc.babylonjs.com/api/classes/babylon.physicsradialexplosioneventoptions#strength
func (p *PhysicsRadialExplosionEventOptions) SetStrength(strength float64) *PhysicsRadialExplosionEventOptions {
	p.p.Set("strength", strength)
	return p
}
