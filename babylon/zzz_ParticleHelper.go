// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ParticleHelper represents a babylon.js ParticleHelper.
// This class is made for on one-liner static method to help creating particle system set.
type ParticleHelper struct{ p js.Value }

// JSObject returns the underlying js.Value.
func (p *ParticleHelper) JSObject() js.Value { return p.p }

// ParticleHelper returns a ParticleHelper JavaScript class.
func (ba *Babylon) ParticleHelper() *ParticleHelper {
	p := ba.ctx.Get("ParticleHelper")
	return ParticleHelperFromJSObject(p)
}

// ParticleHelperFromJSObject returns a wrapped ParticleHelper JavaScript class.
func ParticleHelperFromJSObject(p js.Value) *ParticleHelper {
	return &ParticleHelper{p: p}
}

// TODO: methods
