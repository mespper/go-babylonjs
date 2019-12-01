// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ParticleSystemSet represents a babylon.js ParticleSystemSet.
// Represents a set of particle systems working together to create a specific effect
type ParticleSystemSet struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (p *ParticleSystemSet) JSObject() js.Value { return p.p }

// ParticleSystemSet returns a ParticleSystemSet JavaScript class.
func (ba *Babylon) ParticleSystemSet() *ParticleSystemSet {
	p := ba.ctx.Get("ParticleSystemSet")
	return ParticleSystemSetFromJSObject(p, ba.ctx)
}

// ParticleSystemSetFromJSObject returns a wrapped ParticleSystemSet JavaScript class.
func ParticleSystemSetFromJSObject(p js.Value, ctx js.Value) *ParticleSystemSet {
	return &ParticleSystemSet{p: p, ctx: ctx}
}

// ParticleSystemSetArrayToJSArray returns a JavaScript Array for the wrapped array.
func ParticleSystemSetArrayToJSArray(array []*ParticleSystemSet) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// Dispose calls the Dispose method on the ParticleSystemSet object.
//
// https://doc.babylonjs.com/api/classes/babylon.particlesystemset#dispose
func (p *ParticleSystemSet) Dispose() {

	p.p.Call("dispose")
}

// ParticleSystemSetParseOpts contains optional parameters for ParticleSystemSet.Parse.
type ParticleSystemSetParseOpts struct {
	Gpu *bool
}

// Parse calls the Parse method on the ParticleSystemSet object.
//
// https://doc.babylonjs.com/api/classes/babylon.particlesystemset#parse
func (p *ParticleSystemSet) Parse(data interface{}, scene *Scene, opts *ParticleSystemSetParseOpts) *ParticleSystemSet {
	if opts == nil {
		opts = &ParticleSystemSetParseOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, data)
	args = append(args, scene.JSObject())

	if opts.Gpu == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Gpu)
	}

	retVal := p.p.Call("Parse", args...)
	return ParticleSystemSetFromJSObject(retVal, p.ctx)
}

// Serialize calls the Serialize method on the ParticleSystemSet object.
//
// https://doc.babylonjs.com/api/classes/babylon.particlesystemset#serialize
func (p *ParticleSystemSet) Serialize() interface{} {

	retVal := p.p.Call("serialize")
	return retVal
}

// SetEmitterAsSphere calls the SetEmitterAsSphere method on the ParticleSystemSet object.
//
// https://doc.babylonjs.com/api/classes/babylon.particlesystemset#setemitterassphere
func (p *ParticleSystemSet) SetEmitterAsSphere(options js.Value, renderingGroupId float64, scene *Scene) {

	args := make([]interface{}, 0, 3+0)

	args = append(args, options)
	args = append(args, renderingGroupId)
	args = append(args, scene.JSObject())

	p.p.Call("setEmitterAsSphere", args...)
}

// ParticleSystemSetStartOpts contains optional parameters for ParticleSystemSet.Start.
type ParticleSystemSetStartOpts struct {
	Emitter *AbstractMesh
}

// Start calls the Start method on the ParticleSystemSet object.
//
// https://doc.babylonjs.com/api/classes/babylon.particlesystemset#start
func (p *ParticleSystemSet) Start(opts *ParticleSystemSetStartOpts) {
	if opts == nil {
		opts = &ParticleSystemSetStartOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.Emitter == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Emitter.JSObject())
	}

	p.p.Call("start", args...)
}

/*

// BaseAssetsUrl returns the BaseAssetsUrl property of class ParticleSystemSet.
//
// https://doc.babylonjs.com/api/classes/babylon.particlesystemset#baseassetsurl
func (p *ParticleSystemSet) BaseAssetsUrl(BaseAssetsUrl string) *ParticleSystemSet {
	p := ba.ctx.Get("ParticleSystemSet").New(BaseAssetsUrl)
	return ParticleSystemSetFromJSObject(p, ba.ctx)
}

// SetBaseAssetsUrl sets the BaseAssetsUrl property of class ParticleSystemSet.
//
// https://doc.babylonjs.com/api/classes/babylon.particlesystemset#baseassetsurl
func (p *ParticleSystemSet) SetBaseAssetsUrl(BaseAssetsUrl string) *ParticleSystemSet {
	p := ba.ctx.Get("ParticleSystemSet").New(BaseAssetsUrl)
	return ParticleSystemSetFromJSObject(p, ba.ctx)
}

// EmitterNode returns the EmitterNode property of class ParticleSystemSet.
//
// https://doc.babylonjs.com/api/classes/babylon.particlesystemset#emitternode
func (p *ParticleSystemSet) EmitterNode(emitterNode *TransformNode) *ParticleSystemSet {
	p := ba.ctx.Get("ParticleSystemSet").New(emitterNode.JSObject())
	return ParticleSystemSetFromJSObject(p, ba.ctx)
}

// SetEmitterNode sets the EmitterNode property of class ParticleSystemSet.
//
// https://doc.babylonjs.com/api/classes/babylon.particlesystemset#emitternode
func (p *ParticleSystemSet) SetEmitterNode(emitterNode *TransformNode) *ParticleSystemSet {
	p := ba.ctx.Get("ParticleSystemSet").New(emitterNode.JSObject())
	return ParticleSystemSetFromJSObject(p, ba.ctx)
}

// Systems returns the Systems property of class ParticleSystemSet.
//
// https://doc.babylonjs.com/api/classes/babylon.particlesystemset#systems
func (p *ParticleSystemSet) Systems(systems *IParticleSystem) *ParticleSystemSet {
	p := ba.ctx.Get("ParticleSystemSet").New(systems.JSObject())
	return ParticleSystemSetFromJSObject(p, ba.ctx)
}

// SetSystems sets the Systems property of class ParticleSystemSet.
//
// https://doc.babylonjs.com/api/classes/babylon.particlesystemset#systems
func (p *ParticleSystemSet) SetSystems(systems *IParticleSystem) *ParticleSystemSet {
	p := ba.ctx.Get("ParticleSystemSet").New(systems.JSObject())
	return ParticleSystemSetFromJSObject(p, ba.ctx)
}

*/
