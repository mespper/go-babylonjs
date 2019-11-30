// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// SolidParticleSystem represents a babylon.js SolidParticleSystem.
// The SPS is a single updatable mesh. The solid particles are simply separate parts or faces fo this big mesh.
// As it is just a mesh, the SPS has all the same properties than any other BJS mesh : not more, not less. It can be scaled, rotated, translated, enlighted, textured, moved, etc.
// The SPS is also a particle system. It provides some methods to manage the particles.
// However it is behavior agnostic. This means it has no emitter, no particle physics, no particle recycler. You have to implement your own behavior.
//
// Full documentation here : &lt;a href=&#34;http://doc.babylonjs.com/how_to/Solid_Particle_System&#34;&gt;http://doc.babylonjs.com/how_to/Solid_Particle_System&lt;/a&gt;
type SolidParticleSystem struct{ p js.Value }

// JSObject returns the underlying js.Value.
func (s *SolidParticleSystem) JSObject() js.Value { return s.p }

// SolidParticleSystem returns a SolidParticleSystem JavaScript class.
func (ba *Babylon) SolidParticleSystem() *SolidParticleSystem {
	p := ba.ctx.Get("SolidParticleSystem")
	return SolidParticleSystemFromJSObject(p)
}

// SolidParticleSystemFromJSObject returns a wrapped SolidParticleSystem JavaScript class.
func SolidParticleSystemFromJSObject(p js.Value) *SolidParticleSystem {
	return &SolidParticleSystem{p: p}
}

// NewSolidParticleSystemOpts contains optional parameters for NewSolidParticleSystem.
type NewSolidParticleSystemOpts struct {
	Options *JSValue
}

// NewSolidParticleSystem returns a new SolidParticleSystem object.
//
// https://doc.babylonjs.com/api/classes/babylon.solidparticlesystem
func (ba *Babylon) NewSolidParticleSystem(name string, scene *Scene, opts *NewSolidParticleSystemOpts) *SolidParticleSystem {
	if opts == nil {
		opts = &NewSolidParticleSystemOpts{}
	}

	p := ba.ctx.Get("SolidParticleSystem").New(name, scene.JSObject(), opts.Options.JSObject())
	return SolidParticleSystemFromJSObject(p)
}

// TODO: methods
