// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// PhysicsHelper represents a babylon.js PhysicsHelper.
// A helper for physics simulations
//
// See: https://doc.babylonjs.com/how_to/using_the_physics_engine#further-functionality-of-the-impostor-class
type PhysicsHelper struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (p *PhysicsHelper) JSObject() js.Value { return p.p }

// PhysicsHelper returns a PhysicsHelper JavaScript class.
func (ba *Babylon) PhysicsHelper() *PhysicsHelper {
	p := ba.ctx.Get("PhysicsHelper")
	return PhysicsHelperFromJSObject(p, ba.ctx)
}

// PhysicsHelperFromJSObject returns a wrapped PhysicsHelper JavaScript class.
func PhysicsHelperFromJSObject(p js.Value, ctx js.Value) *PhysicsHelper {
	return &PhysicsHelper{p: p, ctx: ctx}
}

// PhysicsHelperArrayToJSArray returns a JavaScript Array for the wrapped array.
func PhysicsHelperArrayToJSArray(array []*PhysicsHelper) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewPhysicsHelper returns a new PhysicsHelper object.
//
// https://doc.babylonjs.com/api/classes/babylon.physicshelper
func (ba *Babylon) NewPhysicsHelper(scene *Scene) *PhysicsHelper {

	args := make([]interface{}, 0, 1+0)

	args = append(args, scene.JSObject())

	p := ba.ctx.Get("PhysicsHelper").New(args...)
	return PhysicsHelperFromJSObject(p, ba.ctx)
}

// PhysicsHelperApplyRadialExplosionForceOpts contains optional parameters for PhysicsHelper.ApplyRadialExplosionForce.
type PhysicsHelperApplyRadialExplosionForceOpts struct {
	Strength *float64
	Falloff  js.Value
}

// ApplyRadialExplosionForce calls the ApplyRadialExplosionForce method on the PhysicsHelper object.
//
// https://doc.babylonjs.com/api/classes/babylon.physicshelper#applyradialexplosionforce
func (p *PhysicsHelper) ApplyRadialExplosionForce(origin *Vector3, radiusOrEventOptions float64, opts *PhysicsHelperApplyRadialExplosionForceOpts) *PhysicsRadialExplosionEvent {
	if opts == nil {
		opts = &PhysicsHelperApplyRadialExplosionForceOpts{}
	}

	args := make([]interface{}, 0, 2+2)

	args = append(args, origin.JSObject())
	args = append(args, radiusOrEventOptions)

	if opts.Strength == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Strength)
	}
	args = append(args, opts.Falloff)

	retVal := p.p.Call("applyRadialExplosionForce", args...)
	return PhysicsRadialExplosionEventFromJSObject(retVal, p.ctx)
}

// PhysicsHelperApplyRadialExplosionImpulseOpts contains optional parameters for PhysicsHelper.ApplyRadialExplosionImpulse.
type PhysicsHelperApplyRadialExplosionImpulseOpts struct {
	Strength *float64
	Falloff  js.Value
}

// ApplyRadialExplosionImpulse calls the ApplyRadialExplosionImpulse method on the PhysicsHelper object.
//
// https://doc.babylonjs.com/api/classes/babylon.physicshelper#applyradialexplosionimpulse
func (p *PhysicsHelper) ApplyRadialExplosionImpulse(origin *Vector3, radiusOrEventOptions float64, opts *PhysicsHelperApplyRadialExplosionImpulseOpts) *PhysicsRadialExplosionEvent {
	if opts == nil {
		opts = &PhysicsHelperApplyRadialExplosionImpulseOpts{}
	}

	args := make([]interface{}, 0, 2+2)

	args = append(args, origin.JSObject())
	args = append(args, radiusOrEventOptions)

	if opts.Strength == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Strength)
	}
	args = append(args, opts.Falloff)

	retVal := p.p.Call("applyRadialExplosionImpulse", args...)
	return PhysicsRadialExplosionEventFromJSObject(retVal, p.ctx)
}

// PhysicsHelperGravitationalFieldOpts contains optional parameters for PhysicsHelper.GravitationalField.
type PhysicsHelperGravitationalFieldOpts struct {
	Strength *float64
	Falloff  js.Value
}

// GravitationalField calls the GravitationalField method on the PhysicsHelper object.
//
// https://doc.babylonjs.com/api/classes/babylon.physicshelper#gravitationalfield
func (p *PhysicsHelper) GravitationalField(origin *Vector3, radiusOrEventOptions float64, opts *PhysicsHelperGravitationalFieldOpts) *PhysicsGravitationalFieldEvent {
	if opts == nil {
		opts = &PhysicsHelperGravitationalFieldOpts{}
	}

	args := make([]interface{}, 0, 2+2)

	args = append(args, origin.JSObject())
	args = append(args, radiusOrEventOptions)

	if opts.Strength == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Strength)
	}
	args = append(args, opts.Falloff)

	retVal := p.p.Call("gravitationalField", args...)
	return PhysicsGravitationalFieldEventFromJSObject(retVal, p.ctx)
}

// PhysicsHelperUpdraftOpts contains optional parameters for PhysicsHelper.Updraft.
type PhysicsHelperUpdraftOpts struct {
	Strength    *float64
	Height      *float64
	UpdraftMode js.Value
}

// Updraft calls the Updraft method on the PhysicsHelper object.
//
// https://doc.babylonjs.com/api/classes/babylon.physicshelper#updraft
func (p *PhysicsHelper) Updraft(origin *Vector3, radiusOrEventOptions float64, opts *PhysicsHelperUpdraftOpts) *PhysicsUpdraftEvent {
	if opts == nil {
		opts = &PhysicsHelperUpdraftOpts{}
	}

	args := make([]interface{}, 0, 2+3)

	args = append(args, origin.JSObject())
	args = append(args, radiusOrEventOptions)

	if opts.Strength == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Strength)
	}
	if opts.Height == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Height)
	}
	args = append(args, opts.UpdraftMode)

	retVal := p.p.Call("updraft", args...)
	return PhysicsUpdraftEventFromJSObject(retVal, p.ctx)
}

// PhysicsHelperVortexOpts contains optional parameters for PhysicsHelper.Vortex.
type PhysicsHelperVortexOpts struct {
	Strength *float64
	Height   *float64
}

// Vortex calls the Vortex method on the PhysicsHelper object.
//
// https://doc.babylonjs.com/api/classes/babylon.physicshelper#vortex
func (p *PhysicsHelper) Vortex(origin *Vector3, radiusOrEventOptions float64, opts *PhysicsHelperVortexOpts) *PhysicsVortexEvent {
	if opts == nil {
		opts = &PhysicsHelperVortexOpts{}
	}

	args := make([]interface{}, 0, 2+2)

	args = append(args, origin.JSObject())
	args = append(args, radiusOrEventOptions)

	if opts.Strength == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Strength)
	}
	if opts.Height == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Height)
	}

	retVal := p.p.Call("vortex", args...)
	return PhysicsVortexEventFromJSObject(retVal, p.ctx)
}
