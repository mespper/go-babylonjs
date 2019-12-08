// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ParticleSystem represents a babylon.js ParticleSystem.
// This represents a particle system in Babylon.
// Particles are often small sprites used to simulate hard-to-reproduce phenomena like fire, smoke, water, or abstract visual effects like magic glitter and faery dust.
// Particles can take different shapes while emitted like box, sphere, cone or you can write your custom function.
//
// See: https://doc.babylonjs.com/babylon101/particles
type ParticleSystem struct {
	*BaseParticleSystem
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (p *ParticleSystem) JSObject() js.Value { return p.p }

// ParticleSystem returns a ParticleSystem JavaScript class.
func (ba *Babylon) ParticleSystem() *ParticleSystem {
	p := ba.ctx.Get("ParticleSystem")
	return ParticleSystemFromJSObject(p, ba.ctx)
}

// ParticleSystemFromJSObject returns a wrapped ParticleSystem JavaScript class.
func ParticleSystemFromJSObject(p js.Value, ctx js.Value) *ParticleSystem {
	return &ParticleSystem{BaseParticleSystem: BaseParticleSystemFromJSObject(p, ctx), ctx: ctx}
}

// ParticleSystemArrayToJSArray returns a JavaScript Array for the wrapped array.
func ParticleSystemArrayToJSArray(array []*ParticleSystem) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewParticleSystemOpts contains optional parameters for NewParticleSystem.
type NewParticleSystemOpts struct {
	CustomEffect            *Effect
	IsAnimationSheetEnabled *bool
	Epsilon                 *float64
}

// NewParticleSystem returns a new ParticleSystem object.
//
// https://doc.babylonjs.com/api/classes/babylon.particlesystem
func (ba *Babylon) NewParticleSystem(name string, capacity float64, scene *Scene, opts *NewParticleSystemOpts) *ParticleSystem {
	if opts == nil {
		opts = &NewParticleSystemOpts{}
	}

	args := make([]interface{}, 0, 3+3)

	args = append(args, name)
	args = append(args, capacity)
	args = append(args, scene.JSObject())

	if opts.CustomEffect == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.CustomEffect.JSObject())
	}
	if opts.IsAnimationSheetEnabled == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.IsAnimationSheetEnabled)
	}
	if opts.Epsilon == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Epsilon)
	}

	p := ba.ctx.Get("ParticleSystem").New(args...)
	return ParticleSystemFromJSObject(p, ba.ctx)
}

// AddAlphaRemapGradient calls the AddAlphaRemapGradient method on the ParticleSystem object.
//
// https://doc.babylonjs.com/api/classes/babylon.particlesystem#addalpharemapgradient
func (p *ParticleSystem) AddAlphaRemapGradient(gradient float64, min float64, max float64) *IParticleSystem {

	args := make([]interface{}, 0, 3+0)

	args = append(args, gradient)

	args = append(args, min)

	args = append(args, max)

	retVal := p.p.Call("addAlphaRemapGradient", args...)
	return IParticleSystemFromJSObject(retVal, p.ctx)
}

// ParticleSystemAddAngularSpeedGradientOpts contains optional parameters for ParticleSystem.AddAngularSpeedGradient.
type ParticleSystemAddAngularSpeedGradientOpts struct {
	Factor2 *float64
}

// AddAngularSpeedGradient calls the AddAngularSpeedGradient method on the ParticleSystem object.
//
// https://doc.babylonjs.com/api/classes/babylon.particlesystem#addangularspeedgradient
func (p *ParticleSystem) AddAngularSpeedGradient(gradient float64, factor float64, opts *ParticleSystemAddAngularSpeedGradientOpts) *IParticleSystem {
	if opts == nil {
		opts = &ParticleSystemAddAngularSpeedGradientOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, gradient)

	args = append(args, factor)

	if opts.Factor2 == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Factor2)
	}

	retVal := p.p.Call("addAngularSpeedGradient", args...)
	return IParticleSystemFromJSObject(retVal, p.ctx)
}

// ParticleSystemAddColorGradientOpts contains optional parameters for ParticleSystem.AddColorGradient.
type ParticleSystemAddColorGradientOpts struct {
	Color2 *Color4
}

// AddColorGradient calls the AddColorGradient method on the ParticleSystem object.
//
// https://doc.babylonjs.com/api/classes/babylon.particlesystem#addcolorgradient
func (p *ParticleSystem) AddColorGradient(gradient float64, color1 *Color4, opts *ParticleSystemAddColorGradientOpts) *IParticleSystem {
	if opts == nil {
		opts = &ParticleSystemAddColorGradientOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, gradient)

	if color1 == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, color1.JSObject())
	}

	if opts.Color2 == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Color2.JSObject())
	}

	retVal := p.p.Call("addColorGradient", args...)
	return IParticleSystemFromJSObject(retVal, p.ctx)
}

// AddColorRemapGradient calls the AddColorRemapGradient method on the ParticleSystem object.
//
// https://doc.babylonjs.com/api/classes/babylon.particlesystem#addcolorremapgradient
func (p *ParticleSystem) AddColorRemapGradient(gradient float64, min float64, max float64) *IParticleSystem {

	args := make([]interface{}, 0, 3+0)

	args = append(args, gradient)

	args = append(args, min)

	args = append(args, max)

	retVal := p.p.Call("addColorRemapGradient", args...)
	return IParticleSystemFromJSObject(retVal, p.ctx)
}

// ParticleSystemAddDragGradientOpts contains optional parameters for ParticleSystem.AddDragGradient.
type ParticleSystemAddDragGradientOpts struct {
	Factor2 *float64
}

// AddDragGradient calls the AddDragGradient method on the ParticleSystem object.
//
// https://doc.babylonjs.com/api/classes/babylon.particlesystem#adddraggradient
func (p *ParticleSystem) AddDragGradient(gradient float64, factor float64, opts *ParticleSystemAddDragGradientOpts) *IParticleSystem {
	if opts == nil {
		opts = &ParticleSystemAddDragGradientOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, gradient)

	args = append(args, factor)

	if opts.Factor2 == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Factor2)
	}

	retVal := p.p.Call("addDragGradient", args...)
	return IParticleSystemFromJSObject(retVal, p.ctx)
}

// ParticleSystemAddEmitRateGradientOpts contains optional parameters for ParticleSystem.AddEmitRateGradient.
type ParticleSystemAddEmitRateGradientOpts struct {
	Factor2 *float64
}

// AddEmitRateGradient calls the AddEmitRateGradient method on the ParticleSystem object.
//
// https://doc.babylonjs.com/api/classes/babylon.particlesystem#addemitrategradient
func (p *ParticleSystem) AddEmitRateGradient(gradient float64, factor float64, opts *ParticleSystemAddEmitRateGradientOpts) *IParticleSystem {
	if opts == nil {
		opts = &ParticleSystemAddEmitRateGradientOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, gradient)

	args = append(args, factor)

	if opts.Factor2 == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Factor2)
	}

	retVal := p.p.Call("addEmitRateGradient", args...)
	return IParticleSystemFromJSObject(retVal, p.ctx)
}

// ParticleSystemAddLifeTimeGradientOpts contains optional parameters for ParticleSystem.AddLifeTimeGradient.
type ParticleSystemAddLifeTimeGradientOpts struct {
	Factor2 *float64
}

// AddLifeTimeGradient calls the AddLifeTimeGradient method on the ParticleSystem object.
//
// https://doc.babylonjs.com/api/classes/babylon.particlesystem#addlifetimegradient
func (p *ParticleSystem) AddLifeTimeGradient(gradient float64, factor float64, opts *ParticleSystemAddLifeTimeGradientOpts) *IParticleSystem {
	if opts == nil {
		opts = &ParticleSystemAddLifeTimeGradientOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, gradient)

	args = append(args, factor)

	if opts.Factor2 == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Factor2)
	}

	retVal := p.p.Call("addLifeTimeGradient", args...)
	return IParticleSystemFromJSObject(retVal, p.ctx)
}

// ParticleSystemAddLimitVelocityGradientOpts contains optional parameters for ParticleSystem.AddLimitVelocityGradient.
type ParticleSystemAddLimitVelocityGradientOpts struct {
	Factor2 *float64
}

// AddLimitVelocityGradient calls the AddLimitVelocityGradient method on the ParticleSystem object.
//
// https://doc.babylonjs.com/api/classes/babylon.particlesystem#addlimitvelocitygradient
func (p *ParticleSystem) AddLimitVelocityGradient(gradient float64, factor float64, opts *ParticleSystemAddLimitVelocityGradientOpts) *IParticleSystem {
	if opts == nil {
		opts = &ParticleSystemAddLimitVelocityGradientOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, gradient)

	args = append(args, factor)

	if opts.Factor2 == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Factor2)
	}

	retVal := p.p.Call("addLimitVelocityGradient", args...)
	return IParticleSystemFromJSObject(retVal, p.ctx)
}

// AddRampGradient calls the AddRampGradient method on the ParticleSystem object.
//
// https://doc.babylonjs.com/api/classes/babylon.particlesystem#addrampgradient
func (p *ParticleSystem) AddRampGradient(gradient float64, color *Color3) *ParticleSystem {

	args := make([]interface{}, 0, 2+0)

	args = append(args, gradient)

	if color == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, color.JSObject())
	}

	retVal := p.p.Call("addRampGradient", args...)
	return ParticleSystemFromJSObject(retVal, p.ctx)
}

// ParticleSystemAddSizeGradientOpts contains optional parameters for ParticleSystem.AddSizeGradient.
type ParticleSystemAddSizeGradientOpts struct {
	Factor2 *float64
}

// AddSizeGradient calls the AddSizeGradient method on the ParticleSystem object.
//
// https://doc.babylonjs.com/api/classes/babylon.particlesystem#addsizegradient
func (p *ParticleSystem) AddSizeGradient(gradient float64, factor float64, opts *ParticleSystemAddSizeGradientOpts) *IParticleSystem {
	if opts == nil {
		opts = &ParticleSystemAddSizeGradientOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, gradient)

	args = append(args, factor)

	if opts.Factor2 == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Factor2)
	}

	retVal := p.p.Call("addSizeGradient", args...)
	return IParticleSystemFromJSObject(retVal, p.ctx)
}

// ParticleSystemAddStartSizeGradientOpts contains optional parameters for ParticleSystem.AddStartSizeGradient.
type ParticleSystemAddStartSizeGradientOpts struct {
	Factor2 *float64
}

// AddStartSizeGradient calls the AddStartSizeGradient method on the ParticleSystem object.
//
// https://doc.babylonjs.com/api/classes/babylon.particlesystem#addstartsizegradient
func (p *ParticleSystem) AddStartSizeGradient(gradient float64, factor float64, opts *ParticleSystemAddStartSizeGradientOpts) *IParticleSystem {
	if opts == nil {
		opts = &ParticleSystemAddStartSizeGradientOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, gradient)

	args = append(args, factor)

	if opts.Factor2 == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Factor2)
	}

	retVal := p.p.Call("addStartSizeGradient", args...)
	return IParticleSystemFromJSObject(retVal, p.ctx)
}

// ParticleSystemAddVelocityGradientOpts contains optional parameters for ParticleSystem.AddVelocityGradient.
type ParticleSystemAddVelocityGradientOpts struct {
	Factor2 *float64
}

// AddVelocityGradient calls the AddVelocityGradient method on the ParticleSystem object.
//
// https://doc.babylonjs.com/api/classes/babylon.particlesystem#addvelocitygradient
func (p *ParticleSystem) AddVelocityGradient(gradient float64, factor float64, opts *ParticleSystemAddVelocityGradientOpts) *IParticleSystem {
	if opts == nil {
		opts = &ParticleSystemAddVelocityGradientOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, gradient)

	args = append(args, factor)

	if opts.Factor2 == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Factor2)
	}

	retVal := p.p.Call("addVelocityGradient", args...)
	return IParticleSystemFromJSObject(retVal, p.ctx)
}

// ParticleSystemAnimateOpts contains optional parameters for ParticleSystem.Animate.
type ParticleSystemAnimateOpts struct {
	PreWarmOnly *bool
}

// Animate calls the Animate method on the ParticleSystem object.
//
// https://doc.babylonjs.com/api/classes/babylon.particlesystem#animate
func (p *ParticleSystem) Animate(opts *ParticleSystemAnimateOpts) {
	if opts == nil {
		opts = &ParticleSystemAnimateOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.PreWarmOnly == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.PreWarmOnly)
	}

	p.p.Call("animate", args...)
}

// Clone calls the Clone method on the ParticleSystem object.
//
// https://doc.babylonjs.com/api/classes/babylon.particlesystem#clone
func (p *ParticleSystem) Clone(name string, newEmitter JSObject) *ParticleSystem {

	args := make([]interface{}, 0, 2+0)

	args = append(args, name)

	if newEmitter == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, newEmitter.JSObject())
	}

	retVal := p.p.Call("clone", args...)
	return ParticleSystemFromJSObject(retVal, p.ctx)
}

// ParticleSystemDisposeOpts contains optional parameters for ParticleSystem.Dispose.
type ParticleSystemDisposeOpts struct {
	DisposeTexture *bool
}

// Dispose calls the Dispose method on the ParticleSystem object.
//
// https://doc.babylonjs.com/api/classes/babylon.particlesystem#dispose
func (p *ParticleSystem) Dispose(opts *ParticleSystemDisposeOpts) {
	if opts == nil {
		opts = &ParticleSystemDisposeOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.DisposeTexture == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.DisposeTexture)
	}

	p.p.Call("dispose", args...)
}

// GetCapacity calls the GetCapacity method on the ParticleSystem object.
//
// https://doc.babylonjs.com/api/classes/babylon.particlesystem#getcapacity
func (p *ParticleSystem) GetCapacity() float64 {

	retVal := p.p.Call("getCapacity")
	return retVal.Float()
}

// GetClassName calls the GetClassName method on the ParticleSystem object.
//
// https://doc.babylonjs.com/api/classes/babylon.particlesystem#getclassname
func (p *ParticleSystem) GetClassName() string {

	retVal := p.p.Call("getClassName")
	return retVal.String()
}

// GetRampGradients calls the GetRampGradients method on the ParticleSystem object.
//
// https://doc.babylonjs.com/api/classes/babylon.particlesystem#getrampgradients
func (p *ParticleSystem) GetRampGradients() []*Color3Gradient {

	retVal := p.p.Call("getRampGradients")
	result := []*Color3Gradient{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, Color3GradientFromJSObject(retVal.Index(ri), p.ctx))
	}
	return result
}

// IsAlive calls the IsAlive method on the ParticleSystem object.
//
// https://doc.babylonjs.com/api/classes/babylon.particlesystem#isalive
func (p *ParticleSystem) IsAlive() bool {

	retVal := p.p.Call("isAlive")
	return retVal.Bool()
}

// IsReady calls the IsReady method on the ParticleSystem object.
//
// https://doc.babylonjs.com/api/classes/babylon.particlesystem#isready
func (p *ParticleSystem) IsReady() bool {

	retVal := p.p.Call("isReady")
	return retVal.Bool()
}

// IsStarted calls the IsStarted method on the ParticleSystem object.
//
// https://doc.babylonjs.com/api/classes/babylon.particlesystem#isstarted
func (p *ParticleSystem) IsStarted() bool {

	retVal := p.p.Call("isStarted")
	return retVal.Bool()
}

// ParticleSystemParseOpts contains optional parameters for ParticleSystem.Parse.
type ParticleSystemParseOpts struct {
	DoNotStart *bool
}

// Parse calls the Parse method on the ParticleSystem object.
//
// https://doc.babylonjs.com/api/classes/babylon.particlesystem#parse
func (p *ParticleSystem) Parse(parsedParticleSystem JSObject, scene *Scene, rootUrl string, opts *ParticleSystemParseOpts) *ParticleSystem {
	if opts == nil {
		opts = &ParticleSystemParseOpts{}
	}

	args := make([]interface{}, 0, 3+1)

	if parsedParticleSystem == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, parsedParticleSystem.JSObject())
	}

	if scene == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, scene.JSObject())
	}

	args = append(args, rootUrl)

	if opts.DoNotStart == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.DoNotStart)
	}

	retVal := p.p.Call("Parse", args...)
	return ParticleSystemFromJSObject(retVal, p.ctx)
}

// Rebuild calls the Rebuild method on the ParticleSystem object.
//
// https://doc.babylonjs.com/api/classes/babylon.particlesystem#rebuild
func (p *ParticleSystem) Rebuild() {

	p.p.Call("rebuild")
}

// RemoveAlphaRemapGradient calls the RemoveAlphaRemapGradient method on the ParticleSystem object.
//
// https://doc.babylonjs.com/api/classes/babylon.particlesystem#removealpharemapgradient
func (p *ParticleSystem) RemoveAlphaRemapGradient(gradient float64) *IParticleSystem {

	args := make([]interface{}, 0, 1+0)

	args = append(args, gradient)

	retVal := p.p.Call("removeAlphaRemapGradient", args...)
	return IParticleSystemFromJSObject(retVal, p.ctx)
}

// RemoveAngularSpeedGradient calls the RemoveAngularSpeedGradient method on the ParticleSystem object.
//
// https://doc.babylonjs.com/api/classes/babylon.particlesystem#removeangularspeedgradient
func (p *ParticleSystem) RemoveAngularSpeedGradient(gradient float64) *IParticleSystem {

	args := make([]interface{}, 0, 1+0)

	args = append(args, gradient)

	retVal := p.p.Call("removeAngularSpeedGradient", args...)
	return IParticleSystemFromJSObject(retVal, p.ctx)
}

// RemoveColorGradient calls the RemoveColorGradient method on the ParticleSystem object.
//
// https://doc.babylonjs.com/api/classes/babylon.particlesystem#removecolorgradient
func (p *ParticleSystem) RemoveColorGradient(gradient float64) *IParticleSystem {

	args := make([]interface{}, 0, 1+0)

	args = append(args, gradient)

	retVal := p.p.Call("removeColorGradient", args...)
	return IParticleSystemFromJSObject(retVal, p.ctx)
}

// RemoveColorRemapGradient calls the RemoveColorRemapGradient method on the ParticleSystem object.
//
// https://doc.babylonjs.com/api/classes/babylon.particlesystem#removecolorremapgradient
func (p *ParticleSystem) RemoveColorRemapGradient(gradient float64) *IParticleSystem {

	args := make([]interface{}, 0, 1+0)

	args = append(args, gradient)

	retVal := p.p.Call("removeColorRemapGradient", args...)
	return IParticleSystemFromJSObject(retVal, p.ctx)
}

// RemoveDragGradient calls the RemoveDragGradient method on the ParticleSystem object.
//
// https://doc.babylonjs.com/api/classes/babylon.particlesystem#removedraggradient
func (p *ParticleSystem) RemoveDragGradient(gradient float64) *IParticleSystem {

	args := make([]interface{}, 0, 1+0)

	args = append(args, gradient)

	retVal := p.p.Call("removeDragGradient", args...)
	return IParticleSystemFromJSObject(retVal, p.ctx)
}

// RemoveEmitRateGradient calls the RemoveEmitRateGradient method on the ParticleSystem object.
//
// https://doc.babylonjs.com/api/classes/babylon.particlesystem#removeemitrategradient
func (p *ParticleSystem) RemoveEmitRateGradient(gradient float64) *IParticleSystem {

	args := make([]interface{}, 0, 1+0)

	args = append(args, gradient)

	retVal := p.p.Call("removeEmitRateGradient", args...)
	return IParticleSystemFromJSObject(retVal, p.ctx)
}

// RemoveLifeTimeGradient calls the RemoveLifeTimeGradient method on the ParticleSystem object.
//
// https://doc.babylonjs.com/api/classes/babylon.particlesystem#removelifetimegradient
func (p *ParticleSystem) RemoveLifeTimeGradient(gradient float64) *IParticleSystem {

	args := make([]interface{}, 0, 1+0)

	args = append(args, gradient)

	retVal := p.p.Call("removeLifeTimeGradient", args...)
	return IParticleSystemFromJSObject(retVal, p.ctx)
}

// RemoveLimitVelocityGradient calls the RemoveLimitVelocityGradient method on the ParticleSystem object.
//
// https://doc.babylonjs.com/api/classes/babylon.particlesystem#removelimitvelocitygradient
func (p *ParticleSystem) RemoveLimitVelocityGradient(gradient float64) *IParticleSystem {

	args := make([]interface{}, 0, 1+0)

	args = append(args, gradient)

	retVal := p.p.Call("removeLimitVelocityGradient", args...)
	return IParticleSystemFromJSObject(retVal, p.ctx)
}

// RemoveRampGradient calls the RemoveRampGradient method on the ParticleSystem object.
//
// https://doc.babylonjs.com/api/classes/babylon.particlesystem#removerampgradient
func (p *ParticleSystem) RemoveRampGradient(gradient float64) *ParticleSystem {

	args := make([]interface{}, 0, 1+0)

	args = append(args, gradient)

	retVal := p.p.Call("removeRampGradient", args...)
	return ParticleSystemFromJSObject(retVal, p.ctx)
}

// RemoveSizeGradient calls the RemoveSizeGradient method on the ParticleSystem object.
//
// https://doc.babylonjs.com/api/classes/babylon.particlesystem#removesizegradient
func (p *ParticleSystem) RemoveSizeGradient(gradient float64) *IParticleSystem {

	args := make([]interface{}, 0, 1+0)

	args = append(args, gradient)

	retVal := p.p.Call("removeSizeGradient", args...)
	return IParticleSystemFromJSObject(retVal, p.ctx)
}

// RemoveStartSizeGradient calls the RemoveStartSizeGradient method on the ParticleSystem object.
//
// https://doc.babylonjs.com/api/classes/babylon.particlesystem#removestartsizegradient
func (p *ParticleSystem) RemoveStartSizeGradient(gradient float64) *IParticleSystem {

	args := make([]interface{}, 0, 1+0)

	args = append(args, gradient)

	retVal := p.p.Call("removeStartSizeGradient", args...)
	return IParticleSystemFromJSObject(retVal, p.ctx)
}

// RemoveVelocityGradient calls the RemoveVelocityGradient method on the ParticleSystem object.
//
// https://doc.babylonjs.com/api/classes/babylon.particlesystem#removevelocitygradient
func (p *ParticleSystem) RemoveVelocityGradient(gradient float64) *IParticleSystem {

	args := make([]interface{}, 0, 1+0)

	args = append(args, gradient)

	retVal := p.p.Call("removeVelocityGradient", args...)
	return IParticleSystemFromJSObject(retVal, p.ctx)
}

// Render calls the Render method on the ParticleSystem object.
//
// https://doc.babylonjs.com/api/classes/babylon.particlesystem#render
func (p *ParticleSystem) Render() float64 {

	retVal := p.p.Call("render")
	return retVal.Float()
}

// Reset calls the Reset method on the ParticleSystem object.
//
// https://doc.babylonjs.com/api/classes/babylon.particlesystem#reset
func (p *ParticleSystem) Reset() {

	p.p.Call("reset")
}

// Serialize calls the Serialize method on the ParticleSystem object.
//
// https://doc.babylonjs.com/api/classes/babylon.particlesystem#serialize
func (p *ParticleSystem) Serialize() js.Value {

	retVal := p.p.Call("serialize")
	return retVal
}

// ParticleSystemStartOpts contains optional parameters for ParticleSystem.Start.
type ParticleSystemStartOpts struct {
	Delay *float64
}

// Start calls the Start method on the ParticleSystem object.
//
// https://doc.babylonjs.com/api/classes/babylon.particlesystem#start
func (p *ParticleSystem) Start(opts *ParticleSystemStartOpts) {
	if opts == nil {
		opts = &ParticleSystemStartOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.Delay == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Delay)
	}

	p.p.Call("start", args...)
}

// ParticleSystemStopOpts contains optional parameters for ParticleSystem.Stop.
type ParticleSystemStopOpts struct {
	StopSubEmitters *bool
}

// Stop calls the Stop method on the ParticleSystem object.
//
// https://doc.babylonjs.com/api/classes/babylon.particlesystem#stop
func (p *ParticleSystem) Stop(opts *ParticleSystemStopOpts) {
	if opts == nil {
		opts = &ParticleSystemStopOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.StopSubEmitters == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.StopSubEmitters)
	}

	p.p.Call("stop", args...)
}

// ActiveSubSystems returns the ActiveSubSystems property of class ParticleSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.particlesystem#activesubsystems
func (p *ParticleSystem) ActiveSubSystems() []*ParticleSystem {
	retVal := p.p.Get("activeSubSystems")
	result := []*ParticleSystem{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, ParticleSystemFromJSObject(retVal.Index(ri), p.ctx))
	}
	return result
}

// SetActiveSubSystems sets the ActiveSubSystems property of class ParticleSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.particlesystem#activesubsystems
func (p *ParticleSystem) SetActiveSubSystems(activeSubSystems []*ParticleSystem) *ParticleSystem {
	p.p.Set("activeSubSystems", activeSubSystems)
	return p
}

// BILLBOARDMODE_ALL returns the BILLBOARDMODE_ALL property of class ParticleSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.particlesystem#billboardmode_all
func (p *ParticleSystem) BILLBOARDMODE_ALL() float64 {
	retVal := p.p.Get("BILLBOARDMODE_ALL")
	return retVal.Float()
}

// SetBILLBOARDMODE_ALL sets the BILLBOARDMODE_ALL property of class ParticleSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.particlesystem#billboardmode_all
func (p *ParticleSystem) SetBILLBOARDMODE_ALL(BILLBOARDMODE_ALL float64) *ParticleSystem {
	p.p.Set("BILLBOARDMODE_ALL", BILLBOARDMODE_ALL)
	return p
}

// BILLBOARDMODE_STRETCHED returns the BILLBOARDMODE_STRETCHED property of class ParticleSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.particlesystem#billboardmode_stretched
func (p *ParticleSystem) BILLBOARDMODE_STRETCHED() float64 {
	retVal := p.p.Get("BILLBOARDMODE_STRETCHED")
	return retVal.Float()
}

// SetBILLBOARDMODE_STRETCHED sets the BILLBOARDMODE_STRETCHED property of class ParticleSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.particlesystem#billboardmode_stretched
func (p *ParticleSystem) SetBILLBOARDMODE_STRETCHED(BILLBOARDMODE_STRETCHED float64) *ParticleSystem {
	p.p.Set("BILLBOARDMODE_STRETCHED", BILLBOARDMODE_STRETCHED)
	return p
}

// BILLBOARDMODE_Y returns the BILLBOARDMODE_Y property of class ParticleSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.particlesystem#billboardmode_y
func (p *ParticleSystem) BILLBOARDMODE_Y() float64 {
	retVal := p.p.Get("BILLBOARDMODE_Y")
	return retVal.Float()
}

// SetBILLBOARDMODE_Y sets the BILLBOARDMODE_Y property of class ParticleSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.particlesystem#billboardmode_y
func (p *ParticleSystem) SetBILLBOARDMODE_Y(BILLBOARDMODE_Y float64) *ParticleSystem {
	p.p.Set("BILLBOARDMODE_Y", BILLBOARDMODE_Y)
	return p
}

// OnDispose returns the OnDispose property of class ParticleSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.particlesystem#ondispose
func (p *ParticleSystem) OnDispose() js.Value {
	retVal := p.p.Get("onDispose")
	return retVal
}

// SetOnDispose sets the OnDispose property of class ParticleSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.particlesystem#ondispose
func (p *ParticleSystem) SetOnDispose(onDispose JSFunc) *ParticleSystem {
	p.p.Set("onDispose", js.FuncOf(onDispose))
	return p
}

// OnDisposeObservable returns the OnDisposeObservable property of class ParticleSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.particlesystem#ondisposeobservable
func (p *ParticleSystem) OnDisposeObservable() *Observable {
	retVal := p.p.Get("onDisposeObservable")
	return ObservableFromJSObject(retVal, p.ctx)
}

// SetOnDisposeObservable sets the OnDisposeObservable property of class ParticleSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.particlesystem#ondisposeobservable
func (p *ParticleSystem) SetOnDisposeObservable(onDisposeObservable *Observable) *ParticleSystem {
	p.p.Set("onDisposeObservable", onDisposeObservable.JSObject())
	return p
}

// Particles returns the Particles property of class ParticleSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.particlesystem#particles
func (p *ParticleSystem) Particles() []*Particle {
	retVal := p.p.Get("particles")
	result := []*Particle{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, ParticleFromJSObject(retVal.Index(ri), p.ctx))
	}
	return result
}

// SetParticles sets the Particles property of class ParticleSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.particlesystem#particles
func (p *ParticleSystem) SetParticles(particles []*Particle) *ParticleSystem {
	p.p.Set("particles", particles)
	return p
}

// RecycleParticle returns the RecycleParticle property of class ParticleSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.particlesystem#recycleparticle
func (p *ParticleSystem) RecycleParticle() js.Value {
	retVal := p.p.Get("recycleParticle")
	return retVal
}

// SetRecycleParticle sets the RecycleParticle property of class ParticleSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.particlesystem#recycleparticle
func (p *ParticleSystem) SetRecycleParticle(recycleParticle JSFunc) *ParticleSystem {
	p.p.Set("recycleParticle", js.FuncOf(recycleParticle))
	return p
}

// StartDirectionFunction returns the StartDirectionFunction property of class ParticleSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.particlesystem#startdirectionfunction
func (p *ParticleSystem) StartDirectionFunction() js.Value {
	retVal := p.p.Get("startDirectionFunction")
	return retVal
}

// SetStartDirectionFunction sets the StartDirectionFunction property of class ParticleSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.particlesystem#startdirectionfunction
func (p *ParticleSystem) SetStartDirectionFunction(startDirectionFunction JSFunc) *ParticleSystem {
	p.p.Set("startDirectionFunction", js.FuncOf(startDirectionFunction))
	return p
}

// StartPositionFunction returns the StartPositionFunction property of class ParticleSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.particlesystem#startpositionfunction
func (p *ParticleSystem) StartPositionFunction() js.Value {
	retVal := p.p.Get("startPositionFunction")
	return retVal
}

// SetStartPositionFunction sets the StartPositionFunction property of class ParticleSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.particlesystem#startpositionfunction
func (p *ParticleSystem) SetStartPositionFunction(startPositionFunction JSFunc) *ParticleSystem {
	p.p.Set("startPositionFunction", js.FuncOf(startPositionFunction))
	return p
}

// SubEmitters returns the SubEmitters property of class ParticleSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.particlesystem#subemitters
func (p *ParticleSystem) SubEmitters() []*ParticleSystem {
	retVal := p.p.Get("subEmitters")
	result := []*ParticleSystem{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, ParticleSystemFromJSObject(retVal.Index(ri), p.ctx))
	}
	return result
}

// SetSubEmitters sets the SubEmitters property of class ParticleSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.particlesystem#subemitters
func (p *ParticleSystem) SetSubEmitters(subEmitters []*ParticleSystem) *ParticleSystem {
	p.p.Set("subEmitters", subEmitters)
	return p
}

// UpdateFunction returns the UpdateFunction property of class ParticleSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.particlesystem#updatefunction
func (p *ParticleSystem) UpdateFunction() js.Value {
	retVal := p.p.Get("updateFunction")
	return retVal
}

// SetUpdateFunction sets the UpdateFunction property of class ParticleSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.particlesystem#updatefunction
func (p *ParticleSystem) SetUpdateFunction(updateFunction JSFunc) *ParticleSystem {
	p.p.Set("updateFunction", js.FuncOf(updateFunction))
	return p
}

// UseRampGradients returns the UseRampGradients property of class ParticleSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.particlesystem#userampgradients
func (p *ParticleSystem) UseRampGradients() bool {
	retVal := p.p.Get("useRampGradients")
	return retVal.Bool()
}

// SetUseRampGradients sets the UseRampGradients property of class ParticleSystem.
//
// https://doc.babylonjs.com/api/classes/babylon.particlesystem#userampgradients
func (p *ParticleSystem) SetUseRampGradients(useRampGradients bool) *ParticleSystem {
	p.p.Set("useRampGradients", useRampGradients)
	return p
}
