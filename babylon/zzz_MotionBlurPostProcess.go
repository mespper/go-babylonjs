// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// MotionBlurPostProcess represents a babylon.js MotionBlurPostProcess.
// The Motion Blur Post Process which blurs an image based on the objects velocity in scene.
// Velocity can be affected by each object&#39;s rotation, position and scale depending on the transformation speed.
// As an example, all you have to do is to create the post-process:
// var mb = new BABYLON.MotionBlurPostProcess(
// &#39;mb&#39;, // The name of the effect.
// scene, // The scene containing the objects to blur according to their velocity.
// 1.0, // The required width/height ratio to downsize to before computing the render pass.
// camera // The camera to apply the render pass to.
// );
// Then, all objects moving, rotating and/or scaling will be blurred depending on the transformation speed.
type MotionBlurPostProcess struct {
	*PostProcess
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (m *MotionBlurPostProcess) JSObject() js.Value { return m.p }

// MotionBlurPostProcess returns a MotionBlurPostProcess JavaScript class.
func (ba *Babylon) MotionBlurPostProcess() *MotionBlurPostProcess {
	p := ba.ctx.Get("MotionBlurPostProcess")
	return MotionBlurPostProcessFromJSObject(p, ba.ctx)
}

// MotionBlurPostProcessFromJSObject returns a wrapped MotionBlurPostProcess JavaScript class.
func MotionBlurPostProcessFromJSObject(p js.Value, ctx js.Value) *MotionBlurPostProcess {
	return &MotionBlurPostProcess{PostProcess: PostProcessFromJSObject(p, ctx), ctx: ctx}
}

// MotionBlurPostProcessArrayToJSArray returns a JavaScript Array for the wrapped array.
func MotionBlurPostProcessArrayToJSArray(array []*MotionBlurPostProcess) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewMotionBlurPostProcessOpts contains optional parameters for NewMotionBlurPostProcess.
type NewMotionBlurPostProcessOpts struct {
	SamplingMode     *float64
	Engine           *Engine
	Reusable         *bool
	TextureType      *float64
	BlockCompilation *bool
}

// NewMotionBlurPostProcess returns a new MotionBlurPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess
func (ba *Babylon) NewMotionBlurPostProcess(name string, scene *Scene, options float64, camera *Camera, opts *NewMotionBlurPostProcessOpts) *MotionBlurPostProcess {
	if opts == nil {
		opts = &NewMotionBlurPostProcessOpts{}
	}

	args := make([]interface{}, 0, 4+5)

	args = append(args, name)
	args = append(args, scene.JSObject())
	args = append(args, options)
	args = append(args, camera.JSObject())

	if opts.SamplingMode == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.SamplingMode)
	}
	if opts.Engine == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Engine.JSObject())
	}
	if opts.Reusable == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Reusable)
	}
	if opts.TextureType == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.TextureType)
	}
	if opts.BlockCompilation == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.BlockCompilation)
	}

	p := ba.ctx.Get("MotionBlurPostProcess").New(args...)
	return MotionBlurPostProcessFromJSObject(p, ba.ctx)
}

// MotionBlurPostProcessActivateOpts contains optional parameters for MotionBlurPostProcess.Activate.
type MotionBlurPostProcessActivateOpts struct {
	SourceTexture     *InternalTexture
	ForceDepthStencil *bool
}

// Activate calls the Activate method on the MotionBlurPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess#activate
func (m *MotionBlurPostProcess) Activate(camera *Camera, opts *MotionBlurPostProcessActivateOpts) *InternalTexture {
	if opts == nil {
		opts = &MotionBlurPostProcessActivateOpts{}
	}

	args := make([]interface{}, 0, 1+2)

	args = append(args, camera.JSObject())

	if opts.SourceTexture == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.SourceTexture.JSObject())
	}
	if opts.ForceDepthStencil == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.ForceDepthStencil)
	}

	retVal := m.p.Call("activate", args...)
	return InternalTextureFromJSObject(retVal, m.ctx)
}

// Apply calls the Apply method on the MotionBlurPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess#apply
func (m *MotionBlurPostProcess) Apply() *Effect {

	retVal := m.p.Call("apply")
	return EffectFromJSObject(retVal, m.ctx)
}

// MotionBlurPostProcessDisposeOpts contains optional parameters for MotionBlurPostProcess.Dispose.
type MotionBlurPostProcessDisposeOpts struct {
	Camera *Camera
}

// Dispose calls the Dispose method on the MotionBlurPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess#dispose
func (m *MotionBlurPostProcess) Dispose(opts *MotionBlurPostProcessDisposeOpts) {
	if opts == nil {
		opts = &MotionBlurPostProcessDisposeOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.Camera == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Camera.JSObject())
	}

	m.p.Call("dispose", args...)
}

// ExcludeSkinnedMesh calls the ExcludeSkinnedMesh method on the MotionBlurPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess#excludeskinnedmesh
func (m *MotionBlurPostProcess) ExcludeSkinnedMesh(skinnedMesh *AbstractMesh) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, skinnedMesh.JSObject())

	m.p.Call("excludeSkinnedMesh", args...)
}

// GetCamera calls the GetCamera method on the MotionBlurPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess#getcamera
func (m *MotionBlurPostProcess) GetCamera() *Camera {

	retVal := m.p.Call("getCamera")
	return CameraFromJSObject(retVal, m.ctx)
}

// GetClassName calls the GetClassName method on the MotionBlurPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess#getclassname
func (m *MotionBlurPostProcess) GetClassName() string {

	retVal := m.p.Call("getClassName")
	return retVal.String()
}

// GetEffect calls the GetEffect method on the MotionBlurPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess#geteffect
func (m *MotionBlurPostProcess) GetEffect() *Effect {

	retVal := m.p.Call("getEffect")
	return EffectFromJSObject(retVal, m.ctx)
}

// GetEffectName calls the GetEffectName method on the MotionBlurPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess#geteffectname
func (m *MotionBlurPostProcess) GetEffectName() string {

	retVal := m.p.Call("getEffectName")
	return retVal.String()
}

// GetEngine calls the GetEngine method on the MotionBlurPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess#getengine
func (m *MotionBlurPostProcess) GetEngine() *Engine {

	retVal := m.p.Call("getEngine")
	return EngineFromJSObject(retVal, m.ctx)
}

// IsReady calls the IsReady method on the MotionBlurPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess#isready
func (m *MotionBlurPostProcess) IsReady() bool {

	retVal := m.p.Call("isReady")
	return retVal.Bool()
}

// IsReusable calls the IsReusable method on the MotionBlurPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess#isreusable
func (m *MotionBlurPostProcess) IsReusable() bool {

	retVal := m.p.Call("isReusable")
	return retVal.Bool()
}

// MarkTextureDirty calls the MarkTextureDirty method on the MotionBlurPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess#marktexturedirty
func (m *MotionBlurPostProcess) MarkTextureDirty() {

	m.p.Call("markTextureDirty")
}

// RemoveExcludedSkinnedMesh calls the RemoveExcludedSkinnedMesh method on the MotionBlurPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess#removeexcludedskinnedmesh
func (m *MotionBlurPostProcess) RemoveExcludedSkinnedMesh(skinnedMesh *AbstractMesh) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, skinnedMesh.JSObject())

	m.p.Call("removeExcludedSkinnedMesh", args...)
}

// ShareOutputWith calls the ShareOutputWith method on the MotionBlurPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess#shareoutputwith
func (m *MotionBlurPostProcess) ShareOutputWith(postProcess *PostProcess) *PostProcess {

	args := make([]interface{}, 0, 1+0)

	args = append(args, postProcess.JSObject())

	retVal := m.p.Call("shareOutputWith", args...)
	return PostProcessFromJSObject(retVal, m.ctx)
}

// MotionBlurPostProcessUpdateEffectOpts contains optional parameters for MotionBlurPostProcess.UpdateEffect.
type MotionBlurPostProcessUpdateEffectOpts struct {
	Defines         *string
	Uniforms        *string
	Samplers        *string
	IndexParameters *interface{}
	OnCompiled      func()
	OnError         func()
}

// UpdateEffect calls the UpdateEffect method on the MotionBlurPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess#updateeffect
func (m *MotionBlurPostProcess) UpdateEffect(opts *MotionBlurPostProcessUpdateEffectOpts) {
	if opts == nil {
		opts = &MotionBlurPostProcessUpdateEffectOpts{}
	}

	args := make([]interface{}, 0, 0+6)

	if opts.Defines == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Defines)
	}
	if opts.Uniforms == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Uniforms)
	}
	if opts.Samplers == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Samplers)
	}
	if opts.IndexParameters == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.IndexParameters)
	}
	if opts.OnCompiled == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.OnCompiled)
	}
	if opts.OnError == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.OnError)
	}

	m.p.Call("updateEffect", args...)
}

// UseOwnOutput calls the UseOwnOutput method on the MotionBlurPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess#useownoutput
func (m *MotionBlurPostProcess) UseOwnOutput() {

	m.p.Call("useOwnOutput")
}

/*

// AdaptScaleToCurrentViewport returns the AdaptScaleToCurrentViewport property of class MotionBlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess#adaptscaletocurrentviewport
func (m *MotionBlurPostProcess) AdaptScaleToCurrentViewport(adaptScaleToCurrentViewport bool) *MotionBlurPostProcess {
	p := ba.ctx.Get("MotionBlurPostProcess").New(adaptScaleToCurrentViewport)
	return MotionBlurPostProcessFromJSObject(p, ba.ctx)
}

// SetAdaptScaleToCurrentViewport sets the AdaptScaleToCurrentViewport property of class MotionBlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess#adaptscaletocurrentviewport
func (m *MotionBlurPostProcess) SetAdaptScaleToCurrentViewport(adaptScaleToCurrentViewport bool) *MotionBlurPostProcess {
	p := ba.ctx.Get("MotionBlurPostProcess").New(adaptScaleToCurrentViewport)
	return MotionBlurPostProcessFromJSObject(p, ba.ctx)
}

// AlphaConstants returns the AlphaConstants property of class MotionBlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess#alphaconstants
func (m *MotionBlurPostProcess) AlphaConstants(alphaConstants *Color4) *MotionBlurPostProcess {
	p := ba.ctx.Get("MotionBlurPostProcess").New(alphaConstants.JSObject())
	return MotionBlurPostProcessFromJSObject(p, ba.ctx)
}

// SetAlphaConstants sets the AlphaConstants property of class MotionBlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess#alphaconstants
func (m *MotionBlurPostProcess) SetAlphaConstants(alphaConstants *Color4) *MotionBlurPostProcess {
	p := ba.ctx.Get("MotionBlurPostProcess").New(alphaConstants.JSObject())
	return MotionBlurPostProcessFromJSObject(p, ba.ctx)
}

// AlphaMode returns the AlphaMode property of class MotionBlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess#alphamode
func (m *MotionBlurPostProcess) AlphaMode(alphaMode float64) *MotionBlurPostProcess {
	p := ba.ctx.Get("MotionBlurPostProcess").New(alphaMode)
	return MotionBlurPostProcessFromJSObject(p, ba.ctx)
}

// SetAlphaMode sets the AlphaMode property of class MotionBlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess#alphamode
func (m *MotionBlurPostProcess) SetAlphaMode(alphaMode float64) *MotionBlurPostProcess {
	p := ba.ctx.Get("MotionBlurPostProcess").New(alphaMode)
	return MotionBlurPostProcessFromJSObject(p, ba.ctx)
}

// AlwaysForcePOT returns the AlwaysForcePOT property of class MotionBlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess#alwaysforcepot
func (m *MotionBlurPostProcess) AlwaysForcePOT(alwaysForcePOT bool) *MotionBlurPostProcess {
	p := ba.ctx.Get("MotionBlurPostProcess").New(alwaysForcePOT)
	return MotionBlurPostProcessFromJSObject(p, ba.ctx)
}

// SetAlwaysForcePOT sets the AlwaysForcePOT property of class MotionBlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess#alwaysforcepot
func (m *MotionBlurPostProcess) SetAlwaysForcePOT(alwaysForcePOT bool) *MotionBlurPostProcess {
	p := ba.ctx.Get("MotionBlurPostProcess").New(alwaysForcePOT)
	return MotionBlurPostProcessFromJSObject(p, ba.ctx)
}

// Animations returns the Animations property of class MotionBlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess#animations
func (m *MotionBlurPostProcess) Animations(animations *Animation) *MotionBlurPostProcess {
	p := ba.ctx.Get("MotionBlurPostProcess").New(animations.JSObject())
	return MotionBlurPostProcessFromJSObject(p, ba.ctx)
}

// SetAnimations sets the Animations property of class MotionBlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess#animations
func (m *MotionBlurPostProcess) SetAnimations(animations *Animation) *MotionBlurPostProcess {
	p := ba.ctx.Get("MotionBlurPostProcess").New(animations.JSObject())
	return MotionBlurPostProcessFromJSObject(p, ba.ctx)
}

// AspectRatio returns the AspectRatio property of class MotionBlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess#aspectratio
func (m *MotionBlurPostProcess) AspectRatio(aspectRatio float64) *MotionBlurPostProcess {
	p := ba.ctx.Get("MotionBlurPostProcess").New(aspectRatio)
	return MotionBlurPostProcessFromJSObject(p, ba.ctx)
}

// SetAspectRatio sets the AspectRatio property of class MotionBlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess#aspectratio
func (m *MotionBlurPostProcess) SetAspectRatio(aspectRatio float64) *MotionBlurPostProcess {
	p := ba.ctx.Get("MotionBlurPostProcess").New(aspectRatio)
	return MotionBlurPostProcessFromJSObject(p, ba.ctx)
}

// AutoClear returns the AutoClear property of class MotionBlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess#autoclear
func (m *MotionBlurPostProcess) AutoClear(autoClear bool) *MotionBlurPostProcess {
	p := ba.ctx.Get("MotionBlurPostProcess").New(autoClear)
	return MotionBlurPostProcessFromJSObject(p, ba.ctx)
}

// SetAutoClear sets the AutoClear property of class MotionBlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess#autoclear
func (m *MotionBlurPostProcess) SetAutoClear(autoClear bool) *MotionBlurPostProcess {
	p := ba.ctx.Get("MotionBlurPostProcess").New(autoClear)
	return MotionBlurPostProcessFromJSObject(p, ba.ctx)
}

// ClearColor returns the ClearColor property of class MotionBlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess#clearcolor
func (m *MotionBlurPostProcess) ClearColor(clearColor *Color4) *MotionBlurPostProcess {
	p := ba.ctx.Get("MotionBlurPostProcess").New(clearColor.JSObject())
	return MotionBlurPostProcessFromJSObject(p, ba.ctx)
}

// SetClearColor sets the ClearColor property of class MotionBlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess#clearcolor
func (m *MotionBlurPostProcess) SetClearColor(clearColor *Color4) *MotionBlurPostProcess {
	p := ba.ctx.Get("MotionBlurPostProcess").New(clearColor.JSObject())
	return MotionBlurPostProcessFromJSObject(p, ba.ctx)
}

// EnablePixelPerfectMode returns the EnablePixelPerfectMode property of class MotionBlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess#enablepixelperfectmode
func (m *MotionBlurPostProcess) EnablePixelPerfectMode(enablePixelPerfectMode bool) *MotionBlurPostProcess {
	p := ba.ctx.Get("MotionBlurPostProcess").New(enablePixelPerfectMode)
	return MotionBlurPostProcessFromJSObject(p, ba.ctx)
}

// SetEnablePixelPerfectMode sets the EnablePixelPerfectMode property of class MotionBlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess#enablepixelperfectmode
func (m *MotionBlurPostProcess) SetEnablePixelPerfectMode(enablePixelPerfectMode bool) *MotionBlurPostProcess {
	p := ba.ctx.Get("MotionBlurPostProcess").New(enablePixelPerfectMode)
	return MotionBlurPostProcessFromJSObject(p, ba.ctx)
}

// ForceFullscreenViewport returns the ForceFullscreenViewport property of class MotionBlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess#forcefullscreenviewport
func (m *MotionBlurPostProcess) ForceFullscreenViewport(forceFullscreenViewport bool) *MotionBlurPostProcess {
	p := ba.ctx.Get("MotionBlurPostProcess").New(forceFullscreenViewport)
	return MotionBlurPostProcessFromJSObject(p, ba.ctx)
}

// SetForceFullscreenViewport sets the ForceFullscreenViewport property of class MotionBlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess#forcefullscreenviewport
func (m *MotionBlurPostProcess) SetForceFullscreenViewport(forceFullscreenViewport bool) *MotionBlurPostProcess {
	p := ba.ctx.Get("MotionBlurPostProcess").New(forceFullscreenViewport)
	return MotionBlurPostProcessFromJSObject(p, ba.ctx)
}

// Height returns the Height property of class MotionBlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess#height
func (m *MotionBlurPostProcess) Height(height float64) *MotionBlurPostProcess {
	p := ba.ctx.Get("MotionBlurPostProcess").New(height)
	return MotionBlurPostProcessFromJSObject(p, ba.ctx)
}

// SetHeight sets the Height property of class MotionBlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess#height
func (m *MotionBlurPostProcess) SetHeight(height float64) *MotionBlurPostProcess {
	p := ba.ctx.Get("MotionBlurPostProcess").New(height)
	return MotionBlurPostProcessFromJSObject(p, ba.ctx)
}

// InputTexture returns the InputTexture property of class MotionBlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess#inputtexture
func (m *MotionBlurPostProcess) InputTexture(inputTexture *InternalTexture) *MotionBlurPostProcess {
	p := ba.ctx.Get("MotionBlurPostProcess").New(inputTexture.JSObject())
	return MotionBlurPostProcessFromJSObject(p, ba.ctx)
}

// SetInputTexture sets the InputTexture property of class MotionBlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess#inputtexture
func (m *MotionBlurPostProcess) SetInputTexture(inputTexture *InternalTexture) *MotionBlurPostProcess {
	p := ba.ctx.Get("MotionBlurPostProcess").New(inputTexture.JSObject())
	return MotionBlurPostProcessFromJSObject(p, ba.ctx)
}

// InspectableCustomProperties returns the InspectableCustomProperties property of class MotionBlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess#inspectablecustomproperties
func (m *MotionBlurPostProcess) InspectableCustomProperties(inspectableCustomProperties *IInspectable) *MotionBlurPostProcess {
	p := ba.ctx.Get("MotionBlurPostProcess").New(inspectableCustomProperties.JSObject())
	return MotionBlurPostProcessFromJSObject(p, ba.ctx)
}

// SetInspectableCustomProperties sets the InspectableCustomProperties property of class MotionBlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess#inspectablecustomproperties
func (m *MotionBlurPostProcess) SetInspectableCustomProperties(inspectableCustomProperties *IInspectable) *MotionBlurPostProcess {
	p := ba.ctx.Get("MotionBlurPostProcess").New(inspectableCustomProperties.JSObject())
	return MotionBlurPostProcessFromJSObject(p, ba.ctx)
}

// IsSupported returns the IsSupported property of class MotionBlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess#issupported
func (m *MotionBlurPostProcess) IsSupported(isSupported bool) *MotionBlurPostProcess {
	p := ba.ctx.Get("MotionBlurPostProcess").New(isSupported)
	return MotionBlurPostProcessFromJSObject(p, ba.ctx)
}

// SetIsSupported sets the IsSupported property of class MotionBlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess#issupported
func (m *MotionBlurPostProcess) SetIsSupported(isSupported bool) *MotionBlurPostProcess {
	p := ba.ctx.Get("MotionBlurPostProcess").New(isSupported)
	return MotionBlurPostProcessFromJSObject(p, ba.ctx)
}

// MotionBlurSamples returns the MotionBlurSamples property of class MotionBlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess#motionblursamples
func (m *MotionBlurPostProcess) MotionBlurSamples(motionBlurSamples float64) *MotionBlurPostProcess {
	p := ba.ctx.Get("MotionBlurPostProcess").New(motionBlurSamples)
	return MotionBlurPostProcessFromJSObject(p, ba.ctx)
}

// SetMotionBlurSamples sets the MotionBlurSamples property of class MotionBlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess#motionblursamples
func (m *MotionBlurPostProcess) SetMotionBlurSamples(motionBlurSamples float64) *MotionBlurPostProcess {
	p := ba.ctx.Get("MotionBlurPostProcess").New(motionBlurSamples)
	return MotionBlurPostProcessFromJSObject(p, ba.ctx)
}

// MotionStrength returns the MotionStrength property of class MotionBlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess#motionstrength
func (m *MotionBlurPostProcess) MotionStrength(motionStrength float64) *MotionBlurPostProcess {
	p := ba.ctx.Get("MotionBlurPostProcess").New(motionStrength)
	return MotionBlurPostProcessFromJSObject(p, ba.ctx)
}

// SetMotionStrength sets the MotionStrength property of class MotionBlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess#motionstrength
func (m *MotionBlurPostProcess) SetMotionStrength(motionStrength float64) *MotionBlurPostProcess {
	p := ba.ctx.Get("MotionBlurPostProcess").New(motionStrength)
	return MotionBlurPostProcessFromJSObject(p, ba.ctx)
}

// Name returns the Name property of class MotionBlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess#name
func (m *MotionBlurPostProcess) Name(name string) *MotionBlurPostProcess {
	p := ba.ctx.Get("MotionBlurPostProcess").New(name)
	return MotionBlurPostProcessFromJSObject(p, ba.ctx)
}

// SetName sets the Name property of class MotionBlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess#name
func (m *MotionBlurPostProcess) SetName(name string) *MotionBlurPostProcess {
	p := ba.ctx.Get("MotionBlurPostProcess").New(name)
	return MotionBlurPostProcessFromJSObject(p, ba.ctx)
}

// OnActivate returns the OnActivate property of class MotionBlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess#onactivate
func (m *MotionBlurPostProcess) OnActivate(onActivate func()) *MotionBlurPostProcess {
	p := ba.ctx.Get("MotionBlurPostProcess").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {onActivate(); return nil}))
	return MotionBlurPostProcessFromJSObject(p, ba.ctx)
}

// SetOnActivate sets the OnActivate property of class MotionBlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess#onactivate
func (m *MotionBlurPostProcess) SetOnActivate(onActivate func()) *MotionBlurPostProcess {
	p := ba.ctx.Get("MotionBlurPostProcess").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {onActivate(); return nil}))
	return MotionBlurPostProcessFromJSObject(p, ba.ctx)
}

// OnActivateObservable returns the OnActivateObservable property of class MotionBlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess#onactivateobservable
func (m *MotionBlurPostProcess) OnActivateObservable(onActivateObservable *Observable) *MotionBlurPostProcess {
	p := ba.ctx.Get("MotionBlurPostProcess").New(onActivateObservable.JSObject())
	return MotionBlurPostProcessFromJSObject(p, ba.ctx)
}

// SetOnActivateObservable sets the OnActivateObservable property of class MotionBlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess#onactivateobservable
func (m *MotionBlurPostProcess) SetOnActivateObservable(onActivateObservable *Observable) *MotionBlurPostProcess {
	p := ba.ctx.Get("MotionBlurPostProcess").New(onActivateObservable.JSObject())
	return MotionBlurPostProcessFromJSObject(p, ba.ctx)
}

// OnAfterRender returns the OnAfterRender property of class MotionBlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess#onafterrender
func (m *MotionBlurPostProcess) OnAfterRender(onAfterRender func()) *MotionBlurPostProcess {
	p := ba.ctx.Get("MotionBlurPostProcess").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {onAfterRender(); return nil}))
	return MotionBlurPostProcessFromJSObject(p, ba.ctx)
}

// SetOnAfterRender sets the OnAfterRender property of class MotionBlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess#onafterrender
func (m *MotionBlurPostProcess) SetOnAfterRender(onAfterRender func()) *MotionBlurPostProcess {
	p := ba.ctx.Get("MotionBlurPostProcess").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {onAfterRender(); return nil}))
	return MotionBlurPostProcessFromJSObject(p, ba.ctx)
}

// OnAfterRenderObservable returns the OnAfterRenderObservable property of class MotionBlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess#onafterrenderobservable
func (m *MotionBlurPostProcess) OnAfterRenderObservable(onAfterRenderObservable *Observable) *MotionBlurPostProcess {
	p := ba.ctx.Get("MotionBlurPostProcess").New(onAfterRenderObservable.JSObject())
	return MotionBlurPostProcessFromJSObject(p, ba.ctx)
}

// SetOnAfterRenderObservable sets the OnAfterRenderObservable property of class MotionBlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess#onafterrenderobservable
func (m *MotionBlurPostProcess) SetOnAfterRenderObservable(onAfterRenderObservable *Observable) *MotionBlurPostProcess {
	p := ba.ctx.Get("MotionBlurPostProcess").New(onAfterRenderObservable.JSObject())
	return MotionBlurPostProcessFromJSObject(p, ba.ctx)
}

// OnApply returns the OnApply property of class MotionBlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess#onapply
func (m *MotionBlurPostProcess) OnApply(onApply func()) *MotionBlurPostProcess {
	p := ba.ctx.Get("MotionBlurPostProcess").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {onApply(); return nil}))
	return MotionBlurPostProcessFromJSObject(p, ba.ctx)
}

// SetOnApply sets the OnApply property of class MotionBlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess#onapply
func (m *MotionBlurPostProcess) SetOnApply(onApply func()) *MotionBlurPostProcess {
	p := ba.ctx.Get("MotionBlurPostProcess").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {onApply(); return nil}))
	return MotionBlurPostProcessFromJSObject(p, ba.ctx)
}

// OnApplyObservable returns the OnApplyObservable property of class MotionBlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess#onapplyobservable
func (m *MotionBlurPostProcess) OnApplyObservable(onApplyObservable *Observable) *MotionBlurPostProcess {
	p := ba.ctx.Get("MotionBlurPostProcess").New(onApplyObservable.JSObject())
	return MotionBlurPostProcessFromJSObject(p, ba.ctx)
}

// SetOnApplyObservable sets the OnApplyObservable property of class MotionBlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess#onapplyobservable
func (m *MotionBlurPostProcess) SetOnApplyObservable(onApplyObservable *Observable) *MotionBlurPostProcess {
	p := ba.ctx.Get("MotionBlurPostProcess").New(onApplyObservable.JSObject())
	return MotionBlurPostProcessFromJSObject(p, ba.ctx)
}

// OnBeforeRender returns the OnBeforeRender property of class MotionBlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess#onbeforerender
func (m *MotionBlurPostProcess) OnBeforeRender(onBeforeRender func()) *MotionBlurPostProcess {
	p := ba.ctx.Get("MotionBlurPostProcess").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {onBeforeRender(); return nil}))
	return MotionBlurPostProcessFromJSObject(p, ba.ctx)
}

// SetOnBeforeRender sets the OnBeforeRender property of class MotionBlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess#onbeforerender
func (m *MotionBlurPostProcess) SetOnBeforeRender(onBeforeRender func()) *MotionBlurPostProcess {
	p := ba.ctx.Get("MotionBlurPostProcess").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {onBeforeRender(); return nil}))
	return MotionBlurPostProcessFromJSObject(p, ba.ctx)
}

// OnBeforeRenderObservable returns the OnBeforeRenderObservable property of class MotionBlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess#onbeforerenderobservable
func (m *MotionBlurPostProcess) OnBeforeRenderObservable(onBeforeRenderObservable *Observable) *MotionBlurPostProcess {
	p := ba.ctx.Get("MotionBlurPostProcess").New(onBeforeRenderObservable.JSObject())
	return MotionBlurPostProcessFromJSObject(p, ba.ctx)
}

// SetOnBeforeRenderObservable sets the OnBeforeRenderObservable property of class MotionBlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess#onbeforerenderobservable
func (m *MotionBlurPostProcess) SetOnBeforeRenderObservable(onBeforeRenderObservable *Observable) *MotionBlurPostProcess {
	p := ba.ctx.Get("MotionBlurPostProcess").New(onBeforeRenderObservable.JSObject())
	return MotionBlurPostProcessFromJSObject(p, ba.ctx)
}

// OnSizeChanged returns the OnSizeChanged property of class MotionBlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess#onsizechanged
func (m *MotionBlurPostProcess) OnSizeChanged(onSizeChanged func()) *MotionBlurPostProcess {
	p := ba.ctx.Get("MotionBlurPostProcess").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {onSizeChanged(); return nil}))
	return MotionBlurPostProcessFromJSObject(p, ba.ctx)
}

// SetOnSizeChanged sets the OnSizeChanged property of class MotionBlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess#onsizechanged
func (m *MotionBlurPostProcess) SetOnSizeChanged(onSizeChanged func()) *MotionBlurPostProcess {
	p := ba.ctx.Get("MotionBlurPostProcess").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {onSizeChanged(); return nil}))
	return MotionBlurPostProcessFromJSObject(p, ba.ctx)
}

// OnSizeChangedObservable returns the OnSizeChangedObservable property of class MotionBlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess#onsizechangedobservable
func (m *MotionBlurPostProcess) OnSizeChangedObservable(onSizeChangedObservable *Observable) *MotionBlurPostProcess {
	p := ba.ctx.Get("MotionBlurPostProcess").New(onSizeChangedObservable.JSObject())
	return MotionBlurPostProcessFromJSObject(p, ba.ctx)
}

// SetOnSizeChangedObservable sets the OnSizeChangedObservable property of class MotionBlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess#onsizechangedobservable
func (m *MotionBlurPostProcess) SetOnSizeChangedObservable(onSizeChangedObservable *Observable) *MotionBlurPostProcess {
	p := ba.ctx.Get("MotionBlurPostProcess").New(onSizeChangedObservable.JSObject())
	return MotionBlurPostProcessFromJSObject(p, ba.ctx)
}

// RenderTargetSamplingMode returns the RenderTargetSamplingMode property of class MotionBlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess#rendertargetsamplingmode
func (m *MotionBlurPostProcess) RenderTargetSamplingMode(renderTargetSamplingMode float64) *MotionBlurPostProcess {
	p := ba.ctx.Get("MotionBlurPostProcess").New(renderTargetSamplingMode)
	return MotionBlurPostProcessFromJSObject(p, ba.ctx)
}

// SetRenderTargetSamplingMode sets the RenderTargetSamplingMode property of class MotionBlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess#rendertargetsamplingmode
func (m *MotionBlurPostProcess) SetRenderTargetSamplingMode(renderTargetSamplingMode float64) *MotionBlurPostProcess {
	p := ba.ctx.Get("MotionBlurPostProcess").New(renderTargetSamplingMode)
	return MotionBlurPostProcessFromJSObject(p, ba.ctx)
}

// Samples returns the Samples property of class MotionBlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess#samples
func (m *MotionBlurPostProcess) Samples(samples float64) *MotionBlurPostProcess {
	p := ba.ctx.Get("MotionBlurPostProcess").New(samples)
	return MotionBlurPostProcessFromJSObject(p, ba.ctx)
}

// SetSamples sets the Samples property of class MotionBlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess#samples
func (m *MotionBlurPostProcess) SetSamples(samples float64) *MotionBlurPostProcess {
	p := ba.ctx.Get("MotionBlurPostProcess").New(samples)
	return MotionBlurPostProcessFromJSObject(p, ba.ctx)
}

// ScaleMode returns the ScaleMode property of class MotionBlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess#scalemode
func (m *MotionBlurPostProcess) ScaleMode(scaleMode float64) *MotionBlurPostProcess {
	p := ba.ctx.Get("MotionBlurPostProcess").New(scaleMode)
	return MotionBlurPostProcessFromJSObject(p, ba.ctx)
}

// SetScaleMode sets the ScaleMode property of class MotionBlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess#scalemode
func (m *MotionBlurPostProcess) SetScaleMode(scaleMode float64) *MotionBlurPostProcess {
	p := ba.ctx.Get("MotionBlurPostProcess").New(scaleMode)
	return MotionBlurPostProcessFromJSObject(p, ba.ctx)
}

// TexelSize returns the TexelSize property of class MotionBlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess#texelsize
func (m *MotionBlurPostProcess) TexelSize(texelSize *Vector2) *MotionBlurPostProcess {
	p := ba.ctx.Get("MotionBlurPostProcess").New(texelSize.JSObject())
	return MotionBlurPostProcessFromJSObject(p, ba.ctx)
}

// SetTexelSize sets the TexelSize property of class MotionBlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess#texelsize
func (m *MotionBlurPostProcess) SetTexelSize(texelSize *Vector2) *MotionBlurPostProcess {
	p := ba.ctx.Get("MotionBlurPostProcess").New(texelSize.JSObject())
	return MotionBlurPostProcessFromJSObject(p, ba.ctx)
}

// UniqueId returns the UniqueId property of class MotionBlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess#uniqueid
func (m *MotionBlurPostProcess) UniqueId(uniqueId float64) *MotionBlurPostProcess {
	p := ba.ctx.Get("MotionBlurPostProcess").New(uniqueId)
	return MotionBlurPostProcessFromJSObject(p, ba.ctx)
}

// SetUniqueId sets the UniqueId property of class MotionBlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess#uniqueid
func (m *MotionBlurPostProcess) SetUniqueId(uniqueId float64) *MotionBlurPostProcess {
	p := ba.ctx.Get("MotionBlurPostProcess").New(uniqueId)
	return MotionBlurPostProcessFromJSObject(p, ba.ctx)
}

// Width returns the Width property of class MotionBlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess#width
func (m *MotionBlurPostProcess) Width(width float64) *MotionBlurPostProcess {
	p := ba.ctx.Get("MotionBlurPostProcess").New(width)
	return MotionBlurPostProcessFromJSObject(p, ba.ctx)
}

// SetWidth sets the Width property of class MotionBlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.motionblurpostprocess#width
func (m *MotionBlurPostProcess) SetWidth(width float64) *MotionBlurPostProcess {
	p := ba.ctx.Get("MotionBlurPostProcess").New(width)
	return MotionBlurPostProcessFromJSObject(p, ba.ctx)
}

*/
