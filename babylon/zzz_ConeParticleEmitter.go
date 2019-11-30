// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ConeParticleEmitter represents a babylon.js ConeParticleEmitter.
// Particle emitter emitting particles from the inside of a cone.
// It emits the particles alongside the cone volume from the base to the particle.
// The emission direction might be randomized.
type ConeParticleEmitter struct{ p js.Value }

// JSObject returns the underlying js.Value.
func (c *ConeParticleEmitter) JSObject() js.Value { return c.p }

// ConeParticleEmitter returns a ConeParticleEmitter JavaScript class.
func (ba *Babylon) ConeParticleEmitter() *ConeParticleEmitter {
	p := ba.ctx.Get("ConeParticleEmitter")
	return ConeParticleEmitterFromJSObject(p)
}

// ConeParticleEmitterFromJSObject returns a wrapped ConeParticleEmitter JavaScript class.
func ConeParticleEmitterFromJSObject(p js.Value) *ConeParticleEmitter {
	return &ConeParticleEmitter{p: p}
}

// NewConeParticleEmitterOpts contains optional parameters for NewConeParticleEmitter.
type NewConeParticleEmitterOpts struct {
	Radius *JSFloat64

	Angle *JSFloat64

	DirectionRandomizer *JSFloat64
}

// NewConeParticleEmitter returns a new ConeParticleEmitter object.
//
// https://doc.babylonjs.com/api/classes/babylon.coneparticleemitter
func (ba *Babylon) NewConeParticleEmitter(opts *NewConeParticleEmitterOpts) *ConeParticleEmitter {
	if opts == nil {
		opts = &NewConeParticleEmitterOpts{}
	}

	p := ba.ctx.Get("ConeParticleEmitter").New(opts.Radius, opts.Angle, opts.DirectionRandomizer)
	return ConeParticleEmitterFromJSObject(p)
}

// TODO: methods
