// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// DefaultRenderingPipeline represents a babylon.js DefaultRenderingPipeline.
// The default rendering pipeline can be added to a scene to apply common post processing effects such as anti-aliasing or depth of field.
// See &lt;a href=&#34;https://doc.babylonjs.com/how_to/using_default_rendering_pipeline&#34;&gt;https://doc.babylonjs.com/how_to/using_default_rendering_pipeline&lt;/a&gt;
type DefaultRenderingPipeline struct{ *PostProcessRenderPipeline }

// JSObject returns the underlying js.Value.
func (d *DefaultRenderingPipeline) JSObject() js.Value { return d.p }

// DefaultRenderingPipeline returns a DefaultRenderingPipeline JavaScript class.
func (ba *Babylon) DefaultRenderingPipeline() *DefaultRenderingPipeline {
	p := ba.ctx.Get("DefaultRenderingPipeline")
	return DefaultRenderingPipelineFromJSObject(p)
}

// DefaultRenderingPipelineFromJSObject returns a wrapped DefaultRenderingPipeline JavaScript class.
func DefaultRenderingPipelineFromJSObject(p js.Value) *DefaultRenderingPipeline {
	return &DefaultRenderingPipeline{PostProcessRenderPipelineFromJSObject(p)}
}

// NewDefaultRenderingPipelineOpts contains optional parameters for NewDefaultRenderingPipeline.
type NewDefaultRenderingPipelineOpts struct {
	Name *JSString

	Hdr *JSBool

	Scene *Scene

	Cameras *Camera

	AutomaticBuild *JSBool
}

// NewDefaultRenderingPipeline returns a new DefaultRenderingPipeline object.
//
// https://doc.babylonjs.com/api/classes/babylon.defaultrenderingpipeline
func (ba *Babylon) NewDefaultRenderingPipeline(opts *NewDefaultRenderingPipelineOpts) *DefaultRenderingPipeline {
	if opts == nil {
		opts = &NewDefaultRenderingPipelineOpts{}
	}

	p := ba.ctx.Get("DefaultRenderingPipeline").New(opts.Name.JSObject(), opts.Hdr.JSObject(), opts.Scene.JSObject(), opts.Cameras.JSObject(), opts.AutomaticBuild.JSObject())
	return DefaultRenderingPipelineFromJSObject(p)
}

// TODO: methods
