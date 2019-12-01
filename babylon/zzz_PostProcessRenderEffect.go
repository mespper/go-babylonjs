// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// PostProcessRenderEffect represents a babylon.js PostProcessRenderEffect.
// This represents a set of one or more post processes in Babylon.
// A post process can be used to apply a shader to a texture after it is rendered.
//
// See: https://doc.babylonjs.com/how_to/how_to_use_postprocessrenderpipeline
type PostProcessRenderEffect struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (p *PostProcessRenderEffect) JSObject() js.Value { return p.p }

// PostProcessRenderEffect returns a PostProcessRenderEffect JavaScript class.
func (ba *Babylon) PostProcessRenderEffect() *PostProcessRenderEffect {
	p := ba.ctx.Get("PostProcessRenderEffect")
	return PostProcessRenderEffectFromJSObject(p, ba.ctx)
}

// PostProcessRenderEffectFromJSObject returns a wrapped PostProcessRenderEffect JavaScript class.
func PostProcessRenderEffectFromJSObject(p js.Value, ctx js.Value) *PostProcessRenderEffect {
	return &PostProcessRenderEffect{p: p, ctx: ctx}
}

// PostProcessRenderEffectArrayToJSArray returns a JavaScript Array for the wrapped array.
func PostProcessRenderEffectArrayToJSArray(array []*PostProcessRenderEffect) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewPostProcessRenderEffectOpts contains optional parameters for NewPostProcessRenderEffect.
type NewPostProcessRenderEffectOpts struct {
	SingleInstance *bool
}

// NewPostProcessRenderEffect returns a new PostProcessRenderEffect object.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocessrendereffect
func (ba *Babylon) NewPostProcessRenderEffect(engine *Engine, name string, getPostProcesses func(), opts *NewPostProcessRenderEffectOpts) *PostProcessRenderEffect {
	if opts == nil {
		opts = &NewPostProcessRenderEffectOpts{}
	}

	args := make([]interface{}, 0, 3+1)

	args = append(args, engine.JSObject())
	args = append(args, name)
	args = append(args, js.FuncOf(func(this js.Value, args []js.Value) interface{} { getPostProcesses(); return nil }))

	if opts.SingleInstance == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.SingleInstance)
	}

	p := ba.ctx.Get("PostProcessRenderEffect").New(args...)
	return PostProcessRenderEffectFromJSObject(p, ba.ctx)
}

// PostProcessRenderEffectGetPostProcessesOpts contains optional parameters for PostProcessRenderEffect.GetPostProcesses.
type PostProcessRenderEffectGetPostProcessesOpts struct {
	Camera *Camera
}

// GetPostProcesses calls the GetPostProcesses method on the PostProcessRenderEffect object.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocessrendereffect#getpostprocesses
func (p *PostProcessRenderEffect) GetPostProcesses(opts *PostProcessRenderEffectGetPostProcessesOpts) []*PostProcess {
	if opts == nil {
		opts = &PostProcessRenderEffectGetPostProcessesOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.Camera == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Camera.JSObject())
	}

	retVal := p.p.Call("getPostProcesses", args...)
	result := []*PostProcess{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, PostProcessFromJSObject(retVal.Index(ri), p.ctx))
	}
	return result
}

/*

// IsSupported returns the IsSupported property of class PostProcessRenderEffect.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocessrendereffect#issupported
func (p *PostProcessRenderEffect) IsSupported(isSupported bool) *PostProcessRenderEffect {
	p := ba.ctx.Get("PostProcessRenderEffect").New(isSupported)
	return PostProcessRenderEffectFromJSObject(p, ba.ctx)
}

// SetIsSupported sets the IsSupported property of class PostProcessRenderEffect.
//
// https://doc.babylonjs.com/api/classes/babylon.postprocessrendereffect#issupported
func (p *PostProcessRenderEffect) SetIsSupported(isSupported bool) *PostProcessRenderEffect {
	p := ba.ctx.Get("PostProcessRenderEffect").New(isSupported)
	return PostProcessRenderEffectFromJSObject(p, ba.ctx)
}

*/
