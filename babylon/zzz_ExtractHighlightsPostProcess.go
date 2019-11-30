// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ExtractHighlightsPostProcess represents a babylon.js ExtractHighlightsPostProcess.
// The extract highlights post process sets all pixels to black except pixels above the specified luminance threshold. Used as the first step for a bloom effect.
type ExtractHighlightsPostProcess struct{ *PostProcess }

// JSObject returns the underlying js.Value.
func (e *ExtractHighlightsPostProcess) JSObject() js.Value { return e.p }

// ExtractHighlightsPostProcess returns a ExtractHighlightsPostProcess JavaScript class.
func (ba *Babylon) ExtractHighlightsPostProcess() *ExtractHighlightsPostProcess {
	p := ba.ctx.Get("ExtractHighlightsPostProcess")
	return ExtractHighlightsPostProcessFromJSObject(p)
}

// ExtractHighlightsPostProcessFromJSObject returns a wrapped ExtractHighlightsPostProcess JavaScript class.
func ExtractHighlightsPostProcessFromJSObject(p js.Value) *ExtractHighlightsPostProcess {
	return &ExtractHighlightsPostProcess{PostProcessFromJSObject(p)}
}

// NewExtractHighlightsPostProcessOpts contains optional parameters for NewExtractHighlightsPostProcess.
type NewExtractHighlightsPostProcessOpts struct {
	SamplingMode *JSFloat64

	Engine *Engine

	Reusable *JSBool

	TextureType *JSFloat64

	BlockCompilation *JSBool
}

// NewExtractHighlightsPostProcess returns a new ExtractHighlightsPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.extracthighlightspostprocess
func (ba *Babylon) NewExtractHighlightsPostProcess(name string, options float64, camera *Camera, opts *NewExtractHighlightsPostProcessOpts) *ExtractHighlightsPostProcess {
	if opts == nil {
		opts = &NewExtractHighlightsPostProcessOpts{}
	}

	p := ba.ctx.Get("ExtractHighlightsPostProcess").New(name, options, camera.JSObject(), opts.SamplingMode.JSObject(), opts.Engine.JSObject(), opts.Reusable.JSObject(), opts.TextureType.JSObject(), opts.BlockCompilation.JSObject())
	return ExtractHighlightsPostProcessFromJSObject(p)
}

// TODO: methods
