// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// VRDistortionCorrectionPostProcess represents a babylon.js VRDistortionCorrectionPostProcess.
// VRDistortionCorrectionPostProcess used for mobile VR
type VRDistortionCorrectionPostProcess struct {
	*PostProcess
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (v *VRDistortionCorrectionPostProcess) JSObject() js.Value { return v.p }

// VRDistortionCorrectionPostProcess returns a VRDistortionCorrectionPostProcess JavaScript class.
func (ba *Babylon) VRDistortionCorrectionPostProcess() *VRDistortionCorrectionPostProcess {
	p := ba.ctx.Get("VRDistortionCorrectionPostProcess")
	return VRDistortionCorrectionPostProcessFromJSObject(p, ba.ctx)
}

// VRDistortionCorrectionPostProcessFromJSObject returns a wrapped VRDistortionCorrectionPostProcess JavaScript class.
func VRDistortionCorrectionPostProcessFromJSObject(p js.Value, ctx js.Value) *VRDistortionCorrectionPostProcess {
	return &VRDistortionCorrectionPostProcess{PostProcess: PostProcessFromJSObject(p, ctx), ctx: ctx}
}

// VRDistortionCorrectionPostProcessArrayToJSArray returns a JavaScript Array for the wrapped array.
func VRDistortionCorrectionPostProcessArrayToJSArray(array []*VRDistortionCorrectionPostProcess) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewVRDistortionCorrectionPostProcess returns a new VRDistortionCorrectionPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdistortioncorrectionpostprocess
func (ba *Babylon) NewVRDistortionCorrectionPostProcess(name string, camera *Camera, isRightEye bool, vrMetrics *VRCameraMetrics) *VRDistortionCorrectionPostProcess {

	args := make([]interface{}, 0, 4+0)

	args = append(args, name)
	args = append(args, camera.JSObject())
	args = append(args, isRightEye)
	args = append(args, vrMetrics.JSObject())

	p := ba.ctx.Get("VRDistortionCorrectionPostProcess").New(args...)
	return VRDistortionCorrectionPostProcessFromJSObject(p, ba.ctx)
}

// VRDistortionCorrectionPostProcessActivateOpts contains optional parameters for VRDistortionCorrectionPostProcess.Activate.
type VRDistortionCorrectionPostProcessActivateOpts struct {
	SourceTexture     *InternalTexture
	ForceDepthStencil *bool
}

// Activate calls the Activate method on the VRDistortionCorrectionPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdistortioncorrectionpostprocess#activate
func (v *VRDistortionCorrectionPostProcess) Activate(camera *Camera, opts *VRDistortionCorrectionPostProcessActivateOpts) *InternalTexture {
	if opts == nil {
		opts = &VRDistortionCorrectionPostProcessActivateOpts{}
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

	retVal := v.p.Call("activate", args...)
	return InternalTextureFromJSObject(retVal, v.ctx)
}

// Apply calls the Apply method on the VRDistortionCorrectionPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdistortioncorrectionpostprocess#apply
func (v *VRDistortionCorrectionPostProcess) Apply() *Effect {

	retVal := v.p.Call("apply")
	return EffectFromJSObject(retVal, v.ctx)
}

// VRDistortionCorrectionPostProcessDisposeOpts contains optional parameters for VRDistortionCorrectionPostProcess.Dispose.
type VRDistortionCorrectionPostProcessDisposeOpts struct {
	Camera *Camera
}

// Dispose calls the Dispose method on the VRDistortionCorrectionPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdistortioncorrectionpostprocess#dispose
func (v *VRDistortionCorrectionPostProcess) Dispose(opts *VRDistortionCorrectionPostProcessDisposeOpts) {
	if opts == nil {
		opts = &VRDistortionCorrectionPostProcessDisposeOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.Camera == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Camera.JSObject())
	}

	v.p.Call("dispose", args...)
}

// GetCamera calls the GetCamera method on the VRDistortionCorrectionPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdistortioncorrectionpostprocess#getcamera
func (v *VRDistortionCorrectionPostProcess) GetCamera() *Camera {

	retVal := v.p.Call("getCamera")
	return CameraFromJSObject(retVal, v.ctx)
}

// GetClassName calls the GetClassName method on the VRDistortionCorrectionPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdistortioncorrectionpostprocess#getclassname
func (v *VRDistortionCorrectionPostProcess) GetClassName() string {

	retVal := v.p.Call("getClassName")
	return retVal.String()
}

// GetEffect calls the GetEffect method on the VRDistortionCorrectionPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdistortioncorrectionpostprocess#geteffect
func (v *VRDistortionCorrectionPostProcess) GetEffect() *Effect {

	retVal := v.p.Call("getEffect")
	return EffectFromJSObject(retVal, v.ctx)
}

// GetEffectName calls the GetEffectName method on the VRDistortionCorrectionPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdistortioncorrectionpostprocess#geteffectname
func (v *VRDistortionCorrectionPostProcess) GetEffectName() string {

	retVal := v.p.Call("getEffectName")
	return retVal.String()
}

// GetEngine calls the GetEngine method on the VRDistortionCorrectionPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdistortioncorrectionpostprocess#getengine
func (v *VRDistortionCorrectionPostProcess) GetEngine() *Engine {

	retVal := v.p.Call("getEngine")
	return EngineFromJSObject(retVal, v.ctx)
}

// IsReady calls the IsReady method on the VRDistortionCorrectionPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdistortioncorrectionpostprocess#isready
func (v *VRDistortionCorrectionPostProcess) IsReady() bool {

	retVal := v.p.Call("isReady")
	return retVal.Bool()
}

// IsReusable calls the IsReusable method on the VRDistortionCorrectionPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdistortioncorrectionpostprocess#isreusable
func (v *VRDistortionCorrectionPostProcess) IsReusable() bool {

	retVal := v.p.Call("isReusable")
	return retVal.Bool()
}

// MarkTextureDirty calls the MarkTextureDirty method on the VRDistortionCorrectionPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdistortioncorrectionpostprocess#marktexturedirty
func (v *VRDistortionCorrectionPostProcess) MarkTextureDirty() {

	v.p.Call("markTextureDirty")
}

// ShareOutputWith calls the ShareOutputWith method on the VRDistortionCorrectionPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdistortioncorrectionpostprocess#shareoutputwith
func (v *VRDistortionCorrectionPostProcess) ShareOutputWith(postProcess *PostProcess) *PostProcess {

	args := make([]interface{}, 0, 1+0)

	args = append(args, postProcess.JSObject())

	retVal := v.p.Call("shareOutputWith", args...)
	return PostProcessFromJSObject(retVal, v.ctx)
}

// VRDistortionCorrectionPostProcessUpdateEffectOpts contains optional parameters for VRDistortionCorrectionPostProcess.UpdateEffect.
type VRDistortionCorrectionPostProcessUpdateEffectOpts struct {
	Defines         *string
	Uniforms        *string
	Samplers        *string
	IndexParameters *interface{}
	OnCompiled      func()
	OnError         func()
}

// UpdateEffect calls the UpdateEffect method on the VRDistortionCorrectionPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdistortioncorrectionpostprocess#updateeffect
func (v *VRDistortionCorrectionPostProcess) UpdateEffect(opts *VRDistortionCorrectionPostProcessUpdateEffectOpts) {
	if opts == nil {
		opts = &VRDistortionCorrectionPostProcessUpdateEffectOpts{}
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

	v.p.Call("updateEffect", args...)
}

// UseOwnOutput calls the UseOwnOutput method on the VRDistortionCorrectionPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdistortioncorrectionpostprocess#useownoutput
func (v *VRDistortionCorrectionPostProcess) UseOwnOutput() {

	v.p.Call("useOwnOutput")
}

/*

// AdaptScaleToCurrentViewport returns the AdaptScaleToCurrentViewport property of class VRDistortionCorrectionPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdistortioncorrectionpostprocess#adaptscaletocurrentviewport
func (v *VRDistortionCorrectionPostProcess) AdaptScaleToCurrentViewport(adaptScaleToCurrentViewport bool) *VRDistortionCorrectionPostProcess {
	p := ba.ctx.Get("VRDistortionCorrectionPostProcess").New(adaptScaleToCurrentViewport)
	return VRDistortionCorrectionPostProcessFromJSObject(p, ba.ctx)
}

// SetAdaptScaleToCurrentViewport sets the AdaptScaleToCurrentViewport property of class VRDistortionCorrectionPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdistortioncorrectionpostprocess#adaptscaletocurrentviewport
func (v *VRDistortionCorrectionPostProcess) SetAdaptScaleToCurrentViewport(adaptScaleToCurrentViewport bool) *VRDistortionCorrectionPostProcess {
	p := ba.ctx.Get("VRDistortionCorrectionPostProcess").New(adaptScaleToCurrentViewport)
	return VRDistortionCorrectionPostProcessFromJSObject(p, ba.ctx)
}

// AlphaConstants returns the AlphaConstants property of class VRDistortionCorrectionPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdistortioncorrectionpostprocess#alphaconstants
func (v *VRDistortionCorrectionPostProcess) AlphaConstants(alphaConstants *Color4) *VRDistortionCorrectionPostProcess {
	p := ba.ctx.Get("VRDistortionCorrectionPostProcess").New(alphaConstants.JSObject())
	return VRDistortionCorrectionPostProcessFromJSObject(p, ba.ctx)
}

// SetAlphaConstants sets the AlphaConstants property of class VRDistortionCorrectionPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdistortioncorrectionpostprocess#alphaconstants
func (v *VRDistortionCorrectionPostProcess) SetAlphaConstants(alphaConstants *Color4) *VRDistortionCorrectionPostProcess {
	p := ba.ctx.Get("VRDistortionCorrectionPostProcess").New(alphaConstants.JSObject())
	return VRDistortionCorrectionPostProcessFromJSObject(p, ba.ctx)
}

// AlphaMode returns the AlphaMode property of class VRDistortionCorrectionPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdistortioncorrectionpostprocess#alphamode
func (v *VRDistortionCorrectionPostProcess) AlphaMode(alphaMode float64) *VRDistortionCorrectionPostProcess {
	p := ba.ctx.Get("VRDistortionCorrectionPostProcess").New(alphaMode)
	return VRDistortionCorrectionPostProcessFromJSObject(p, ba.ctx)
}

// SetAlphaMode sets the AlphaMode property of class VRDistortionCorrectionPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdistortioncorrectionpostprocess#alphamode
func (v *VRDistortionCorrectionPostProcess) SetAlphaMode(alphaMode float64) *VRDistortionCorrectionPostProcess {
	p := ba.ctx.Get("VRDistortionCorrectionPostProcess").New(alphaMode)
	return VRDistortionCorrectionPostProcessFromJSObject(p, ba.ctx)
}

// AlwaysForcePOT returns the AlwaysForcePOT property of class VRDistortionCorrectionPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdistortioncorrectionpostprocess#alwaysforcepot
func (v *VRDistortionCorrectionPostProcess) AlwaysForcePOT(alwaysForcePOT bool) *VRDistortionCorrectionPostProcess {
	p := ba.ctx.Get("VRDistortionCorrectionPostProcess").New(alwaysForcePOT)
	return VRDistortionCorrectionPostProcessFromJSObject(p, ba.ctx)
}

// SetAlwaysForcePOT sets the AlwaysForcePOT property of class VRDistortionCorrectionPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdistortioncorrectionpostprocess#alwaysforcepot
func (v *VRDistortionCorrectionPostProcess) SetAlwaysForcePOT(alwaysForcePOT bool) *VRDistortionCorrectionPostProcess {
	p := ba.ctx.Get("VRDistortionCorrectionPostProcess").New(alwaysForcePOT)
	return VRDistortionCorrectionPostProcessFromJSObject(p, ba.ctx)
}

// Animations returns the Animations property of class VRDistortionCorrectionPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdistortioncorrectionpostprocess#animations
func (v *VRDistortionCorrectionPostProcess) Animations(animations *Animation) *VRDistortionCorrectionPostProcess {
	p := ba.ctx.Get("VRDistortionCorrectionPostProcess").New(animations.JSObject())
	return VRDistortionCorrectionPostProcessFromJSObject(p, ba.ctx)
}

// SetAnimations sets the Animations property of class VRDistortionCorrectionPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdistortioncorrectionpostprocess#animations
func (v *VRDistortionCorrectionPostProcess) SetAnimations(animations *Animation) *VRDistortionCorrectionPostProcess {
	p := ba.ctx.Get("VRDistortionCorrectionPostProcess").New(animations.JSObject())
	return VRDistortionCorrectionPostProcessFromJSObject(p, ba.ctx)
}

// AspectRatio returns the AspectRatio property of class VRDistortionCorrectionPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdistortioncorrectionpostprocess#aspectratio
func (v *VRDistortionCorrectionPostProcess) AspectRatio(aspectRatio float64) *VRDistortionCorrectionPostProcess {
	p := ba.ctx.Get("VRDistortionCorrectionPostProcess").New(aspectRatio)
	return VRDistortionCorrectionPostProcessFromJSObject(p, ba.ctx)
}

// SetAspectRatio sets the AspectRatio property of class VRDistortionCorrectionPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdistortioncorrectionpostprocess#aspectratio
func (v *VRDistortionCorrectionPostProcess) SetAspectRatio(aspectRatio float64) *VRDistortionCorrectionPostProcess {
	p := ba.ctx.Get("VRDistortionCorrectionPostProcess").New(aspectRatio)
	return VRDistortionCorrectionPostProcessFromJSObject(p, ba.ctx)
}

// AutoClear returns the AutoClear property of class VRDistortionCorrectionPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdistortioncorrectionpostprocess#autoclear
func (v *VRDistortionCorrectionPostProcess) AutoClear(autoClear bool) *VRDistortionCorrectionPostProcess {
	p := ba.ctx.Get("VRDistortionCorrectionPostProcess").New(autoClear)
	return VRDistortionCorrectionPostProcessFromJSObject(p, ba.ctx)
}

// SetAutoClear sets the AutoClear property of class VRDistortionCorrectionPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdistortioncorrectionpostprocess#autoclear
func (v *VRDistortionCorrectionPostProcess) SetAutoClear(autoClear bool) *VRDistortionCorrectionPostProcess {
	p := ba.ctx.Get("VRDistortionCorrectionPostProcess").New(autoClear)
	return VRDistortionCorrectionPostProcessFromJSObject(p, ba.ctx)
}

// ClearColor returns the ClearColor property of class VRDistortionCorrectionPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdistortioncorrectionpostprocess#clearcolor
func (v *VRDistortionCorrectionPostProcess) ClearColor(clearColor *Color4) *VRDistortionCorrectionPostProcess {
	p := ba.ctx.Get("VRDistortionCorrectionPostProcess").New(clearColor.JSObject())
	return VRDistortionCorrectionPostProcessFromJSObject(p, ba.ctx)
}

// SetClearColor sets the ClearColor property of class VRDistortionCorrectionPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdistortioncorrectionpostprocess#clearcolor
func (v *VRDistortionCorrectionPostProcess) SetClearColor(clearColor *Color4) *VRDistortionCorrectionPostProcess {
	p := ba.ctx.Get("VRDistortionCorrectionPostProcess").New(clearColor.JSObject())
	return VRDistortionCorrectionPostProcessFromJSObject(p, ba.ctx)
}

// EnablePixelPerfectMode returns the EnablePixelPerfectMode property of class VRDistortionCorrectionPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdistortioncorrectionpostprocess#enablepixelperfectmode
func (v *VRDistortionCorrectionPostProcess) EnablePixelPerfectMode(enablePixelPerfectMode bool) *VRDistortionCorrectionPostProcess {
	p := ba.ctx.Get("VRDistortionCorrectionPostProcess").New(enablePixelPerfectMode)
	return VRDistortionCorrectionPostProcessFromJSObject(p, ba.ctx)
}

// SetEnablePixelPerfectMode sets the EnablePixelPerfectMode property of class VRDistortionCorrectionPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdistortioncorrectionpostprocess#enablepixelperfectmode
func (v *VRDistortionCorrectionPostProcess) SetEnablePixelPerfectMode(enablePixelPerfectMode bool) *VRDistortionCorrectionPostProcess {
	p := ba.ctx.Get("VRDistortionCorrectionPostProcess").New(enablePixelPerfectMode)
	return VRDistortionCorrectionPostProcessFromJSObject(p, ba.ctx)
}

// ForceFullscreenViewport returns the ForceFullscreenViewport property of class VRDistortionCorrectionPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdistortioncorrectionpostprocess#forcefullscreenviewport
func (v *VRDistortionCorrectionPostProcess) ForceFullscreenViewport(forceFullscreenViewport bool) *VRDistortionCorrectionPostProcess {
	p := ba.ctx.Get("VRDistortionCorrectionPostProcess").New(forceFullscreenViewport)
	return VRDistortionCorrectionPostProcessFromJSObject(p, ba.ctx)
}

// SetForceFullscreenViewport sets the ForceFullscreenViewport property of class VRDistortionCorrectionPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdistortioncorrectionpostprocess#forcefullscreenviewport
func (v *VRDistortionCorrectionPostProcess) SetForceFullscreenViewport(forceFullscreenViewport bool) *VRDistortionCorrectionPostProcess {
	p := ba.ctx.Get("VRDistortionCorrectionPostProcess").New(forceFullscreenViewport)
	return VRDistortionCorrectionPostProcessFromJSObject(p, ba.ctx)
}

// Height returns the Height property of class VRDistortionCorrectionPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdistortioncorrectionpostprocess#height
func (v *VRDistortionCorrectionPostProcess) Height(height float64) *VRDistortionCorrectionPostProcess {
	p := ba.ctx.Get("VRDistortionCorrectionPostProcess").New(height)
	return VRDistortionCorrectionPostProcessFromJSObject(p, ba.ctx)
}

// SetHeight sets the Height property of class VRDistortionCorrectionPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdistortioncorrectionpostprocess#height
func (v *VRDistortionCorrectionPostProcess) SetHeight(height float64) *VRDistortionCorrectionPostProcess {
	p := ba.ctx.Get("VRDistortionCorrectionPostProcess").New(height)
	return VRDistortionCorrectionPostProcessFromJSObject(p, ba.ctx)
}

// InputTexture returns the InputTexture property of class VRDistortionCorrectionPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdistortioncorrectionpostprocess#inputtexture
func (v *VRDistortionCorrectionPostProcess) InputTexture(inputTexture *InternalTexture) *VRDistortionCorrectionPostProcess {
	p := ba.ctx.Get("VRDistortionCorrectionPostProcess").New(inputTexture.JSObject())
	return VRDistortionCorrectionPostProcessFromJSObject(p, ba.ctx)
}

// SetInputTexture sets the InputTexture property of class VRDistortionCorrectionPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdistortioncorrectionpostprocess#inputtexture
func (v *VRDistortionCorrectionPostProcess) SetInputTexture(inputTexture *InternalTexture) *VRDistortionCorrectionPostProcess {
	p := ba.ctx.Get("VRDistortionCorrectionPostProcess").New(inputTexture.JSObject())
	return VRDistortionCorrectionPostProcessFromJSObject(p, ba.ctx)
}

// InspectableCustomProperties returns the InspectableCustomProperties property of class VRDistortionCorrectionPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdistortioncorrectionpostprocess#inspectablecustomproperties
func (v *VRDistortionCorrectionPostProcess) InspectableCustomProperties(inspectableCustomProperties *IInspectable) *VRDistortionCorrectionPostProcess {
	p := ba.ctx.Get("VRDistortionCorrectionPostProcess").New(inspectableCustomProperties.JSObject())
	return VRDistortionCorrectionPostProcessFromJSObject(p, ba.ctx)
}

// SetInspectableCustomProperties sets the InspectableCustomProperties property of class VRDistortionCorrectionPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdistortioncorrectionpostprocess#inspectablecustomproperties
func (v *VRDistortionCorrectionPostProcess) SetInspectableCustomProperties(inspectableCustomProperties *IInspectable) *VRDistortionCorrectionPostProcess {
	p := ba.ctx.Get("VRDistortionCorrectionPostProcess").New(inspectableCustomProperties.JSObject())
	return VRDistortionCorrectionPostProcessFromJSObject(p, ba.ctx)
}

// IsSupported returns the IsSupported property of class VRDistortionCorrectionPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdistortioncorrectionpostprocess#issupported
func (v *VRDistortionCorrectionPostProcess) IsSupported(isSupported bool) *VRDistortionCorrectionPostProcess {
	p := ba.ctx.Get("VRDistortionCorrectionPostProcess").New(isSupported)
	return VRDistortionCorrectionPostProcessFromJSObject(p, ba.ctx)
}

// SetIsSupported sets the IsSupported property of class VRDistortionCorrectionPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdistortioncorrectionpostprocess#issupported
func (v *VRDistortionCorrectionPostProcess) SetIsSupported(isSupported bool) *VRDistortionCorrectionPostProcess {
	p := ba.ctx.Get("VRDistortionCorrectionPostProcess").New(isSupported)
	return VRDistortionCorrectionPostProcessFromJSObject(p, ba.ctx)
}

// Name returns the Name property of class VRDistortionCorrectionPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdistortioncorrectionpostprocess#name
func (v *VRDistortionCorrectionPostProcess) Name(name string) *VRDistortionCorrectionPostProcess {
	p := ba.ctx.Get("VRDistortionCorrectionPostProcess").New(name)
	return VRDistortionCorrectionPostProcessFromJSObject(p, ba.ctx)
}

// SetName sets the Name property of class VRDistortionCorrectionPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdistortioncorrectionpostprocess#name
func (v *VRDistortionCorrectionPostProcess) SetName(name string) *VRDistortionCorrectionPostProcess {
	p := ba.ctx.Get("VRDistortionCorrectionPostProcess").New(name)
	return VRDistortionCorrectionPostProcessFromJSObject(p, ba.ctx)
}

// OnActivate returns the OnActivate property of class VRDistortionCorrectionPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdistortioncorrectionpostprocess#onactivate
func (v *VRDistortionCorrectionPostProcess) OnActivate(onActivate func()) *VRDistortionCorrectionPostProcess {
	p := ba.ctx.Get("VRDistortionCorrectionPostProcess").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {onActivate(); return nil}))
	return VRDistortionCorrectionPostProcessFromJSObject(p, ba.ctx)
}

// SetOnActivate sets the OnActivate property of class VRDistortionCorrectionPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdistortioncorrectionpostprocess#onactivate
func (v *VRDistortionCorrectionPostProcess) SetOnActivate(onActivate func()) *VRDistortionCorrectionPostProcess {
	p := ba.ctx.Get("VRDistortionCorrectionPostProcess").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {onActivate(); return nil}))
	return VRDistortionCorrectionPostProcessFromJSObject(p, ba.ctx)
}

// OnActivateObservable returns the OnActivateObservable property of class VRDistortionCorrectionPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdistortioncorrectionpostprocess#onactivateobservable
func (v *VRDistortionCorrectionPostProcess) OnActivateObservable(onActivateObservable *Observable) *VRDistortionCorrectionPostProcess {
	p := ba.ctx.Get("VRDistortionCorrectionPostProcess").New(onActivateObservable.JSObject())
	return VRDistortionCorrectionPostProcessFromJSObject(p, ba.ctx)
}

// SetOnActivateObservable sets the OnActivateObservable property of class VRDistortionCorrectionPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdistortioncorrectionpostprocess#onactivateobservable
func (v *VRDistortionCorrectionPostProcess) SetOnActivateObservable(onActivateObservable *Observable) *VRDistortionCorrectionPostProcess {
	p := ba.ctx.Get("VRDistortionCorrectionPostProcess").New(onActivateObservable.JSObject())
	return VRDistortionCorrectionPostProcessFromJSObject(p, ba.ctx)
}

// OnAfterRender returns the OnAfterRender property of class VRDistortionCorrectionPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdistortioncorrectionpostprocess#onafterrender
func (v *VRDistortionCorrectionPostProcess) OnAfterRender(onAfterRender func()) *VRDistortionCorrectionPostProcess {
	p := ba.ctx.Get("VRDistortionCorrectionPostProcess").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {onAfterRender(); return nil}))
	return VRDistortionCorrectionPostProcessFromJSObject(p, ba.ctx)
}

// SetOnAfterRender sets the OnAfterRender property of class VRDistortionCorrectionPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdistortioncorrectionpostprocess#onafterrender
func (v *VRDistortionCorrectionPostProcess) SetOnAfterRender(onAfterRender func()) *VRDistortionCorrectionPostProcess {
	p := ba.ctx.Get("VRDistortionCorrectionPostProcess").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {onAfterRender(); return nil}))
	return VRDistortionCorrectionPostProcessFromJSObject(p, ba.ctx)
}

// OnAfterRenderObservable returns the OnAfterRenderObservable property of class VRDistortionCorrectionPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdistortioncorrectionpostprocess#onafterrenderobservable
func (v *VRDistortionCorrectionPostProcess) OnAfterRenderObservable(onAfterRenderObservable *Observable) *VRDistortionCorrectionPostProcess {
	p := ba.ctx.Get("VRDistortionCorrectionPostProcess").New(onAfterRenderObservable.JSObject())
	return VRDistortionCorrectionPostProcessFromJSObject(p, ba.ctx)
}

// SetOnAfterRenderObservable sets the OnAfterRenderObservable property of class VRDistortionCorrectionPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdistortioncorrectionpostprocess#onafterrenderobservable
func (v *VRDistortionCorrectionPostProcess) SetOnAfterRenderObservable(onAfterRenderObservable *Observable) *VRDistortionCorrectionPostProcess {
	p := ba.ctx.Get("VRDistortionCorrectionPostProcess").New(onAfterRenderObservable.JSObject())
	return VRDistortionCorrectionPostProcessFromJSObject(p, ba.ctx)
}

// OnApply returns the OnApply property of class VRDistortionCorrectionPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdistortioncorrectionpostprocess#onapply
func (v *VRDistortionCorrectionPostProcess) OnApply(onApply func()) *VRDistortionCorrectionPostProcess {
	p := ba.ctx.Get("VRDistortionCorrectionPostProcess").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {onApply(); return nil}))
	return VRDistortionCorrectionPostProcessFromJSObject(p, ba.ctx)
}

// SetOnApply sets the OnApply property of class VRDistortionCorrectionPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdistortioncorrectionpostprocess#onapply
func (v *VRDistortionCorrectionPostProcess) SetOnApply(onApply func()) *VRDistortionCorrectionPostProcess {
	p := ba.ctx.Get("VRDistortionCorrectionPostProcess").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {onApply(); return nil}))
	return VRDistortionCorrectionPostProcessFromJSObject(p, ba.ctx)
}

// OnApplyObservable returns the OnApplyObservable property of class VRDistortionCorrectionPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdistortioncorrectionpostprocess#onapplyobservable
func (v *VRDistortionCorrectionPostProcess) OnApplyObservable(onApplyObservable *Observable) *VRDistortionCorrectionPostProcess {
	p := ba.ctx.Get("VRDistortionCorrectionPostProcess").New(onApplyObservable.JSObject())
	return VRDistortionCorrectionPostProcessFromJSObject(p, ba.ctx)
}

// SetOnApplyObservable sets the OnApplyObservable property of class VRDistortionCorrectionPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdistortioncorrectionpostprocess#onapplyobservable
func (v *VRDistortionCorrectionPostProcess) SetOnApplyObservable(onApplyObservable *Observable) *VRDistortionCorrectionPostProcess {
	p := ba.ctx.Get("VRDistortionCorrectionPostProcess").New(onApplyObservable.JSObject())
	return VRDistortionCorrectionPostProcessFromJSObject(p, ba.ctx)
}

// OnBeforeRender returns the OnBeforeRender property of class VRDistortionCorrectionPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdistortioncorrectionpostprocess#onbeforerender
func (v *VRDistortionCorrectionPostProcess) OnBeforeRender(onBeforeRender func()) *VRDistortionCorrectionPostProcess {
	p := ba.ctx.Get("VRDistortionCorrectionPostProcess").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {onBeforeRender(); return nil}))
	return VRDistortionCorrectionPostProcessFromJSObject(p, ba.ctx)
}

// SetOnBeforeRender sets the OnBeforeRender property of class VRDistortionCorrectionPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdistortioncorrectionpostprocess#onbeforerender
func (v *VRDistortionCorrectionPostProcess) SetOnBeforeRender(onBeforeRender func()) *VRDistortionCorrectionPostProcess {
	p := ba.ctx.Get("VRDistortionCorrectionPostProcess").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {onBeforeRender(); return nil}))
	return VRDistortionCorrectionPostProcessFromJSObject(p, ba.ctx)
}

// OnBeforeRenderObservable returns the OnBeforeRenderObservable property of class VRDistortionCorrectionPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdistortioncorrectionpostprocess#onbeforerenderobservable
func (v *VRDistortionCorrectionPostProcess) OnBeforeRenderObservable(onBeforeRenderObservable *Observable) *VRDistortionCorrectionPostProcess {
	p := ba.ctx.Get("VRDistortionCorrectionPostProcess").New(onBeforeRenderObservable.JSObject())
	return VRDistortionCorrectionPostProcessFromJSObject(p, ba.ctx)
}

// SetOnBeforeRenderObservable sets the OnBeforeRenderObservable property of class VRDistortionCorrectionPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdistortioncorrectionpostprocess#onbeforerenderobservable
func (v *VRDistortionCorrectionPostProcess) SetOnBeforeRenderObservable(onBeforeRenderObservable *Observable) *VRDistortionCorrectionPostProcess {
	p := ba.ctx.Get("VRDistortionCorrectionPostProcess").New(onBeforeRenderObservable.JSObject())
	return VRDistortionCorrectionPostProcessFromJSObject(p, ba.ctx)
}

// OnSizeChanged returns the OnSizeChanged property of class VRDistortionCorrectionPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdistortioncorrectionpostprocess#onsizechanged
func (v *VRDistortionCorrectionPostProcess) OnSizeChanged(onSizeChanged func()) *VRDistortionCorrectionPostProcess {
	p := ba.ctx.Get("VRDistortionCorrectionPostProcess").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {onSizeChanged(); return nil}))
	return VRDistortionCorrectionPostProcessFromJSObject(p, ba.ctx)
}

// SetOnSizeChanged sets the OnSizeChanged property of class VRDistortionCorrectionPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdistortioncorrectionpostprocess#onsizechanged
func (v *VRDistortionCorrectionPostProcess) SetOnSizeChanged(onSizeChanged func()) *VRDistortionCorrectionPostProcess {
	p := ba.ctx.Get("VRDistortionCorrectionPostProcess").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {onSizeChanged(); return nil}))
	return VRDistortionCorrectionPostProcessFromJSObject(p, ba.ctx)
}

// OnSizeChangedObservable returns the OnSizeChangedObservable property of class VRDistortionCorrectionPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdistortioncorrectionpostprocess#onsizechangedobservable
func (v *VRDistortionCorrectionPostProcess) OnSizeChangedObservable(onSizeChangedObservable *Observable) *VRDistortionCorrectionPostProcess {
	p := ba.ctx.Get("VRDistortionCorrectionPostProcess").New(onSizeChangedObservable.JSObject())
	return VRDistortionCorrectionPostProcessFromJSObject(p, ba.ctx)
}

// SetOnSizeChangedObservable sets the OnSizeChangedObservable property of class VRDistortionCorrectionPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdistortioncorrectionpostprocess#onsizechangedobservable
func (v *VRDistortionCorrectionPostProcess) SetOnSizeChangedObservable(onSizeChangedObservable *Observable) *VRDistortionCorrectionPostProcess {
	p := ba.ctx.Get("VRDistortionCorrectionPostProcess").New(onSizeChangedObservable.JSObject())
	return VRDistortionCorrectionPostProcessFromJSObject(p, ba.ctx)
}

// RenderTargetSamplingMode returns the RenderTargetSamplingMode property of class VRDistortionCorrectionPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdistortioncorrectionpostprocess#rendertargetsamplingmode
func (v *VRDistortionCorrectionPostProcess) RenderTargetSamplingMode(renderTargetSamplingMode float64) *VRDistortionCorrectionPostProcess {
	p := ba.ctx.Get("VRDistortionCorrectionPostProcess").New(renderTargetSamplingMode)
	return VRDistortionCorrectionPostProcessFromJSObject(p, ba.ctx)
}

// SetRenderTargetSamplingMode sets the RenderTargetSamplingMode property of class VRDistortionCorrectionPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdistortioncorrectionpostprocess#rendertargetsamplingmode
func (v *VRDistortionCorrectionPostProcess) SetRenderTargetSamplingMode(renderTargetSamplingMode float64) *VRDistortionCorrectionPostProcess {
	p := ba.ctx.Get("VRDistortionCorrectionPostProcess").New(renderTargetSamplingMode)
	return VRDistortionCorrectionPostProcessFromJSObject(p, ba.ctx)
}

// Samples returns the Samples property of class VRDistortionCorrectionPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdistortioncorrectionpostprocess#samples
func (v *VRDistortionCorrectionPostProcess) Samples(samples float64) *VRDistortionCorrectionPostProcess {
	p := ba.ctx.Get("VRDistortionCorrectionPostProcess").New(samples)
	return VRDistortionCorrectionPostProcessFromJSObject(p, ba.ctx)
}

// SetSamples sets the Samples property of class VRDistortionCorrectionPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdistortioncorrectionpostprocess#samples
func (v *VRDistortionCorrectionPostProcess) SetSamples(samples float64) *VRDistortionCorrectionPostProcess {
	p := ba.ctx.Get("VRDistortionCorrectionPostProcess").New(samples)
	return VRDistortionCorrectionPostProcessFromJSObject(p, ba.ctx)
}

// ScaleMode returns the ScaleMode property of class VRDistortionCorrectionPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdistortioncorrectionpostprocess#scalemode
func (v *VRDistortionCorrectionPostProcess) ScaleMode(scaleMode float64) *VRDistortionCorrectionPostProcess {
	p := ba.ctx.Get("VRDistortionCorrectionPostProcess").New(scaleMode)
	return VRDistortionCorrectionPostProcessFromJSObject(p, ba.ctx)
}

// SetScaleMode sets the ScaleMode property of class VRDistortionCorrectionPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdistortioncorrectionpostprocess#scalemode
func (v *VRDistortionCorrectionPostProcess) SetScaleMode(scaleMode float64) *VRDistortionCorrectionPostProcess {
	p := ba.ctx.Get("VRDistortionCorrectionPostProcess").New(scaleMode)
	return VRDistortionCorrectionPostProcessFromJSObject(p, ba.ctx)
}

// TexelSize returns the TexelSize property of class VRDistortionCorrectionPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdistortioncorrectionpostprocess#texelsize
func (v *VRDistortionCorrectionPostProcess) TexelSize(texelSize *Vector2) *VRDistortionCorrectionPostProcess {
	p := ba.ctx.Get("VRDistortionCorrectionPostProcess").New(texelSize.JSObject())
	return VRDistortionCorrectionPostProcessFromJSObject(p, ba.ctx)
}

// SetTexelSize sets the TexelSize property of class VRDistortionCorrectionPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdistortioncorrectionpostprocess#texelsize
func (v *VRDistortionCorrectionPostProcess) SetTexelSize(texelSize *Vector2) *VRDistortionCorrectionPostProcess {
	p := ba.ctx.Get("VRDistortionCorrectionPostProcess").New(texelSize.JSObject())
	return VRDistortionCorrectionPostProcessFromJSObject(p, ba.ctx)
}

// UniqueId returns the UniqueId property of class VRDistortionCorrectionPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdistortioncorrectionpostprocess#uniqueid
func (v *VRDistortionCorrectionPostProcess) UniqueId(uniqueId float64) *VRDistortionCorrectionPostProcess {
	p := ba.ctx.Get("VRDistortionCorrectionPostProcess").New(uniqueId)
	return VRDistortionCorrectionPostProcessFromJSObject(p, ba.ctx)
}

// SetUniqueId sets the UniqueId property of class VRDistortionCorrectionPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdistortioncorrectionpostprocess#uniqueid
func (v *VRDistortionCorrectionPostProcess) SetUniqueId(uniqueId float64) *VRDistortionCorrectionPostProcess {
	p := ba.ctx.Get("VRDistortionCorrectionPostProcess").New(uniqueId)
	return VRDistortionCorrectionPostProcessFromJSObject(p, ba.ctx)
}

// Width returns the Width property of class VRDistortionCorrectionPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdistortioncorrectionpostprocess#width
func (v *VRDistortionCorrectionPostProcess) Width(width float64) *VRDistortionCorrectionPostProcess {
	p := ba.ctx.Get("VRDistortionCorrectionPostProcess").New(width)
	return VRDistortionCorrectionPostProcessFromJSObject(p, ba.ctx)
}

// SetWidth sets the Width property of class VRDistortionCorrectionPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.vrdistortioncorrectionpostprocess#width
func (v *VRDistortionCorrectionPostProcess) SetWidth(width float64) *VRDistortionCorrectionPostProcess {
	p := ba.ctx.Get("VRDistortionCorrectionPostProcess").New(width)
	return VRDistortionCorrectionPostProcessFromJSObject(p, ba.ctx)
}

*/
