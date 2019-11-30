// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// BloomMergePostProcess represents a babylon.js BloomMergePostProcess.
// The BloomMergePostProcess merges blurred images with the original based on the values of the circle of confusion.
type BloomMergePostProcess struct{ *PostProcess }

// JSObject returns the underlying js.Value.
func (b *BloomMergePostProcess) JSObject() js.Value { return b.p }

// BloomMergePostProcess returns a BloomMergePostProcess JavaScript class.
func (ba *Babylon) BloomMergePostProcess() *BloomMergePostProcess {
	p := ba.ctx.Get("BloomMergePostProcess")
	return BloomMergePostProcessFromJSObject(p)
}

// BloomMergePostProcessFromJSObject returns a wrapped BloomMergePostProcess JavaScript class.
func BloomMergePostProcessFromJSObject(p js.Value) *BloomMergePostProcess {
	return &BloomMergePostProcess{PostProcessFromJSObject(p)}
}

// NewBloomMergePostProcessOpts contains optional parameters for NewBloomMergePostProcess.
type NewBloomMergePostProcessOpts struct {
	SamplingMode *JSFloat64

	Engine *Engine

	Reusable *JSBool

	TextureType *JSFloat64

	BlockCompilation *JSBool
}

// NewBloomMergePostProcess returns a new BloomMergePostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.bloommergepostprocess
func (ba *Babylon) NewBloomMergePostProcess(name string, originalFromInput *PostProcess, blurred *PostProcess, weight float64, options float64, camera *Camera, opts *NewBloomMergePostProcessOpts) *BloomMergePostProcess {
	if opts == nil {
		opts = &NewBloomMergePostProcessOpts{}
	}

	p := ba.ctx.Get("BloomMergePostProcess").New(name, originalFromInput.JSObject(), blurred.JSObject(), weight, options, camera.JSObject(), opts.SamplingMode, opts.Engine.JSObject(), opts.Reusable, opts.TextureType, opts.BlockCompilation)
	return BloomMergePostProcessFromJSObject(p)
}

// TODO: methods
