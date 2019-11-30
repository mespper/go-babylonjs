// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// HemisphericParticleEmitter represents a babylon.js HemisphericParticleEmitter.
// Particle emitter emitting particles from the inside of a hemisphere.
// It emits the particles alongside the hemisphere radius. The emission direction might be randomized.
type HemisphericParticleEmitter struct{ p js.Value }

// JSObject returns the underlying js.Value.
func (h *HemisphericParticleEmitter) JSObject() js.Value { return h.p }

// HemisphericParticleEmitter returns a HemisphericParticleEmitter JavaScript class.
func (ba *Babylon) HemisphericParticleEmitter() *HemisphericParticleEmitter {
	p := ba.ctx.Get("HemisphericParticleEmitter")
	return HemisphericParticleEmitterFromJSObject(p)
}

// HemisphericParticleEmitterFromJSObject returns a wrapped HemisphericParticleEmitter JavaScript class.
func HemisphericParticleEmitterFromJSObject(p js.Value) *HemisphericParticleEmitter {
	return &HemisphericParticleEmitter{p: p}
}

// NewHemisphericParticleEmitterOpts contains optional parameters for NewHemisphericParticleEmitter.
type NewHemisphericParticleEmitterOpts struct {
	Radius *JSFloat64

	RadiusRange *JSFloat64

	DirectionRandomizer *JSFloat64
}

// NewHemisphericParticleEmitter returns a new HemisphericParticleEmitter object.
//
// https://doc.babylonjs.com/api/classes/babylon.hemisphericparticleemitter
func (ba *Babylon) NewHemisphericParticleEmitter(opts *NewHemisphericParticleEmitterOpts) *HemisphericParticleEmitter {
	if opts == nil {
		opts = &NewHemisphericParticleEmitterOpts{}
	}

	p := ba.ctx.Get("HemisphericParticleEmitter").New(opts.Radius.JSObject(), opts.RadiusRange.JSObject(), opts.DirectionRandomizer.JSObject())
	return HemisphericParticleEmitterFromJSObject(p)
}

// TODO: methods
