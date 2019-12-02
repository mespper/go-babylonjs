// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// Particle represents a babylon.js Particle.
// A particle represents one of the element emitted by a particle system.
// This is mainly define by its coordinates, direction, velocity and age.
type Particle struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (p *Particle) JSObject() js.Value { return p.p }

// Particle returns a Particle JavaScript class.
func (ba *Babylon) Particle() *Particle {
	p := ba.ctx.Get("Particle")
	return ParticleFromJSObject(p, ba.ctx)
}

// ParticleFromJSObject returns a wrapped Particle JavaScript class.
func ParticleFromJSObject(p js.Value, ctx js.Value) *Particle {
	return &Particle{p: p, ctx: ctx}
}

// ParticleArrayToJSArray returns a JavaScript Array for the wrapped array.
func ParticleArrayToJSArray(array []*Particle) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewParticle returns a new Particle object.
//
// https://doc.babylonjs.com/api/classes/babylon.particle
func (ba *Babylon) NewParticle(particleSystem *ParticleSystem) *Particle {

	args := make([]interface{}, 0, 1+0)

	args = append(args, particleSystem.JSObject())

	p := ba.ctx.Get("Particle").New(args...)
	return ParticleFromJSObject(p, ba.ctx)
}

// CopyTo calls the CopyTo method on the Particle object.
//
// https://doc.babylonjs.com/api/classes/babylon.particle#copyto
func (p *Particle) CopyTo(other *Particle) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, other.JSObject())

	p.p.Call("copyTo", args...)
}

// UpdateCellIndex calls the UpdateCellIndex method on the Particle object.
//
// https://doc.babylonjs.com/api/classes/babylon.particle#updatecellindex
func (p *Particle) UpdateCellIndex() {

	p.p.Call("updateCellIndex")
}

// Age returns the Age property of class Particle.
//
// https://doc.babylonjs.com/api/classes/babylon.particle#age
func (p *Particle) Age() float64 {
	retVal := p.p.Get("age")
	return retVal.Float()
}

// SetAge sets the Age property of class Particle.
//
// https://doc.babylonjs.com/api/classes/babylon.particle#age
func (p *Particle) SetAge(age float64) *Particle {
	p.p.Set("age", age)
	return p
}

// Angle returns the Angle property of class Particle.
//
// https://doc.babylonjs.com/api/classes/babylon.particle#angle
func (p *Particle) Angle() float64 {
	retVal := p.p.Get("angle")
	return retVal.Float()
}

// SetAngle sets the Angle property of class Particle.
//
// https://doc.babylonjs.com/api/classes/babylon.particle#angle
func (p *Particle) SetAngle(angle float64) *Particle {
	p.p.Set("angle", angle)
	return p
}

// AngularSpeed returns the AngularSpeed property of class Particle.
//
// https://doc.babylonjs.com/api/classes/babylon.particle#angularspeed
func (p *Particle) AngularSpeed() float64 {
	retVal := p.p.Get("angularSpeed")
	return retVal.Float()
}

// SetAngularSpeed sets the AngularSpeed property of class Particle.
//
// https://doc.babylonjs.com/api/classes/babylon.particle#angularspeed
func (p *Particle) SetAngularSpeed(angularSpeed float64) *Particle {
	p.p.Set("angularSpeed", angularSpeed)
	return p
}

// CellIndex returns the CellIndex property of class Particle.
//
// https://doc.babylonjs.com/api/classes/babylon.particle#cellindex
func (p *Particle) CellIndex() float64 {
	retVal := p.p.Get("cellIndex")
	return retVal.Float()
}

// SetCellIndex sets the CellIndex property of class Particle.
//
// https://doc.babylonjs.com/api/classes/babylon.particle#cellindex
func (p *Particle) SetCellIndex(cellIndex float64) *Particle {
	p.p.Set("cellIndex", cellIndex)
	return p
}

// Color returns the Color property of class Particle.
//
// https://doc.babylonjs.com/api/classes/babylon.particle#color
func (p *Particle) Color() *Color4 {
	retVal := p.p.Get("color")
	return Color4FromJSObject(retVal, p.ctx)
}

// SetColor sets the Color property of class Particle.
//
// https://doc.babylonjs.com/api/classes/babylon.particle#color
func (p *Particle) SetColor(color *Color4) *Particle {
	p.p.Set("color", color.JSObject())
	return p
}

// ColorStep returns the ColorStep property of class Particle.
//
// https://doc.babylonjs.com/api/classes/babylon.particle#colorstep
func (p *Particle) ColorStep() *Color4 {
	retVal := p.p.Get("colorStep")
	return Color4FromJSObject(retVal, p.ctx)
}

// SetColorStep sets the ColorStep property of class Particle.
//
// https://doc.babylonjs.com/api/classes/babylon.particle#colorstep
func (p *Particle) SetColorStep(colorStep *Color4) *Particle {
	p.p.Set("colorStep", colorStep.JSObject())
	return p
}

// Direction returns the Direction property of class Particle.
//
// https://doc.babylonjs.com/api/classes/babylon.particle#direction
func (p *Particle) Direction() *Vector3 {
	retVal := p.p.Get("direction")
	return Vector3FromJSObject(retVal, p.ctx)
}

// SetDirection sets the Direction property of class Particle.
//
// https://doc.babylonjs.com/api/classes/babylon.particle#direction
func (p *Particle) SetDirection(direction *Vector3) *Particle {
	p.p.Set("direction", direction.JSObject())
	return p
}

// Id returns the Id property of class Particle.
//
// https://doc.babylonjs.com/api/classes/babylon.particle#id
func (p *Particle) Id() float64 {
	retVal := p.p.Get("id")
	return retVal.Float()
}

// SetId sets the Id property of class Particle.
//
// https://doc.babylonjs.com/api/classes/babylon.particle#id
func (p *Particle) SetId(id float64) *Particle {
	p.p.Set("id", id)
	return p
}

// LifeTime returns the LifeTime property of class Particle.
//
// https://doc.babylonjs.com/api/classes/babylon.particle#lifetime
func (p *Particle) LifeTime() float64 {
	retVal := p.p.Get("lifeTime")
	return retVal.Float()
}

// SetLifeTime sets the LifeTime property of class Particle.
//
// https://doc.babylonjs.com/api/classes/babylon.particle#lifetime
func (p *Particle) SetLifeTime(lifeTime float64) *Particle {
	p.p.Set("lifeTime", lifeTime)
	return p
}

// ParticleSystem returns the ParticleSystem property of class Particle.
//
// https://doc.babylonjs.com/api/classes/babylon.particle#particlesystem
func (p *Particle) ParticleSystem() *ParticleSystem {
	retVal := p.p.Get("particleSystem")
	return ParticleSystemFromJSObject(retVal, p.ctx)
}

// SetParticleSystem sets the ParticleSystem property of class Particle.
//
// https://doc.babylonjs.com/api/classes/babylon.particle#particlesystem
func (p *Particle) SetParticleSystem(particleSystem *ParticleSystem) *Particle {
	p.p.Set("particleSystem", particleSystem.JSObject())
	return p
}

// Position returns the Position property of class Particle.
//
// https://doc.babylonjs.com/api/classes/babylon.particle#position
func (p *Particle) Position() *Vector3 {
	retVal := p.p.Get("position")
	return Vector3FromJSObject(retVal, p.ctx)
}

// SetPosition sets the Position property of class Particle.
//
// https://doc.babylonjs.com/api/classes/babylon.particle#position
func (p *Particle) SetPosition(position *Vector3) *Particle {
	p.p.Set("position", position.JSObject())
	return p
}

// RemapData returns the RemapData property of class Particle.
//
// https://doc.babylonjs.com/api/classes/babylon.particle#remapdata
func (p *Particle) RemapData() *Vector4 {
	retVal := p.p.Get("remapData")
	return Vector4FromJSObject(retVal, p.ctx)
}

// SetRemapData sets the RemapData property of class Particle.
//
// https://doc.babylonjs.com/api/classes/babylon.particle#remapdata
func (p *Particle) SetRemapData(remapData *Vector4) *Particle {
	p.p.Set("remapData", remapData.JSObject())
	return p
}

// Scale returns the Scale property of class Particle.
//
// https://doc.babylonjs.com/api/classes/babylon.particle#scale
func (p *Particle) Scale() *Vector2 {
	retVal := p.p.Get("scale")
	return Vector2FromJSObject(retVal, p.ctx)
}

// SetScale sets the Scale property of class Particle.
//
// https://doc.babylonjs.com/api/classes/babylon.particle#scale
func (p *Particle) SetScale(scale *Vector2) *Particle {
	p.p.Set("scale", scale.JSObject())
	return p
}

// Size returns the Size property of class Particle.
//
// https://doc.babylonjs.com/api/classes/babylon.particle#size
func (p *Particle) Size() float64 {
	retVal := p.p.Get("size")
	return retVal.Float()
}

// SetSize sets the Size property of class Particle.
//
// https://doc.babylonjs.com/api/classes/babylon.particle#size
func (p *Particle) SetSize(size float64) *Particle {
	p.p.Set("size", size)
	return p
}
