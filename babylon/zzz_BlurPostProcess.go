// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// BlurPostProcess represents a babylon.js BlurPostProcess.
// The Blur Post Process which blurs an image based on a kernel and direction.
// Can be used twice in x and y directions to perform a guassian blur in two passes.
type BlurPostProcess struct{ *PostProcess }

// JSObject returns the underlying js.Value.
func (b *BlurPostProcess) JSObject() js.Value { return b.p }

// BlurPostProcess returns a BlurPostProcess JavaScript class.
func (ba *Babylon) BlurPostProcess() *BlurPostProcess {
	p := ba.ctx.Get("BlurPostProcess")
	return BlurPostProcessFromJSObject(p)
}

// BlurPostProcessFromJSObject returns a wrapped BlurPostProcess JavaScript class.
func BlurPostProcessFromJSObject(p js.Value) *BlurPostProcess {
	return &BlurPostProcess{PostProcessFromJSObject(p)}
}

// NewBlurPostProcessOpts contains optional parameters for NewBlurPostProcess.
type NewBlurPostProcessOpts struct {
	SamplingMode *JSFloat64

	Engine *Engine

	Reusable *JSBool

	TextureType *JSFloat64

	Defines *JSString

	BlockCompilation *JSBool
}

// NewBlurPostProcess returns a new BlurPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess
func (ba *Babylon) NewBlurPostProcess(name string, direction *Vector2, kernel float64, options float64, camera *Camera, opts *NewBlurPostProcessOpts) *BlurPostProcess {
	if opts == nil {
		opts = &NewBlurPostProcessOpts{}
	}

	p := ba.ctx.Get("BlurPostProcess").New(name, direction.JSObject(), kernel, options, camera.JSObject(), opts.SamplingMode, opts.Engine.JSObject(), opts.Reusable, opts.TextureType, opts.Defines, opts.BlockCompilation)
	return BlurPostProcessFromJSObject(p)
}

// TODO: methods
