// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// VolumetricLightScatteringPostProcess represents a babylon.js VolumetricLightScatteringPostProcess.
// Inspired by &lt;a href=&#34;http://http.developer.nvidia.com/GPUGems3/gpugems3_ch13.html&#34;&gt;http://http.developer.nvidia.com/GPUGems3/gpugems3_ch13.html&lt;/a&gt;
type VolumetricLightScatteringPostProcess struct{ *PostProcess }

// JSObject returns the underlying js.Value.
func (v *VolumetricLightScatteringPostProcess) JSObject() js.Value { return v.p }

// VolumetricLightScatteringPostProcess returns a VolumetricLightScatteringPostProcess JavaScript class.
func (ba *Babylon) VolumetricLightScatteringPostProcess() *VolumetricLightScatteringPostProcess {
	p := ba.ctx.Get("VolumetricLightScatteringPostProcess")
	return VolumetricLightScatteringPostProcessFromJSObject(p)
}

// VolumetricLightScatteringPostProcessFromJSObject returns a wrapped VolumetricLightScatteringPostProcess JavaScript class.
func VolumetricLightScatteringPostProcessFromJSObject(p js.Value) *VolumetricLightScatteringPostProcess {
	return &VolumetricLightScatteringPostProcess{PostProcessFromJSObject(p)}
}

// NewVolumetricLightScatteringPostProcessOpts contains optional parameters for NewVolumetricLightScatteringPostProcess.
type NewVolumetricLightScatteringPostProcessOpts struct {
	Mesh *Mesh

	Samples *JSFloat64

	SamplingMode *JSFloat64

	Engine *Engine

	Reusable *JSBool

	Scene *Scene
}

// NewVolumetricLightScatteringPostProcess returns a new VolumetricLightScatteringPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.volumetriclightscatteringpostprocess
func (ba *Babylon) NewVolumetricLightScatteringPostProcess(name string, ratio interface{}, camera *Camera, opts *NewVolumetricLightScatteringPostProcessOpts) *VolumetricLightScatteringPostProcess {
	if opts == nil {
		opts = &NewVolumetricLightScatteringPostProcessOpts{}
	}

	p := ba.ctx.Get("VolumetricLightScatteringPostProcess").New(name, ratio, camera.JSObject(), opts.Mesh.JSObject(), opts.Samples.JSObject(), opts.SamplingMode.JSObject(), opts.Engine.JSObject(), opts.Reusable.JSObject(), opts.Scene.JSObject())
	return VolumetricLightScatteringPostProcessFromJSObject(p)
}

// TODO: methods
