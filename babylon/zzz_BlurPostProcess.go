// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// BlurPostProcess represents a babylon.js BlurPostProcess.
// The Blur Post Process which blurs an image based on a kernel and direction.
// Can be used twice in x and y directions to perform a guassian blur in two passes.
type BlurPostProcess struct {
	*PostProcess
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (b *BlurPostProcess) JSObject() js.Value { return b.p }

// BlurPostProcess returns a BlurPostProcess JavaScript class.
func (ba *Babylon) BlurPostProcess() *BlurPostProcess {
	p := ba.ctx.Get("BlurPostProcess")
	return BlurPostProcessFromJSObject(p, ba.ctx)
}

// BlurPostProcessFromJSObject returns a wrapped BlurPostProcess JavaScript class.
func BlurPostProcessFromJSObject(p js.Value, ctx js.Value) *BlurPostProcess {
	return &BlurPostProcess{PostProcess: PostProcessFromJSObject(p, ctx), ctx: ctx}
}

// BlurPostProcessArrayToJSArray returns a JavaScript Array for the wrapped array.
func BlurPostProcessArrayToJSArray(array []*BlurPostProcess) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewBlurPostProcessOpts contains optional parameters for NewBlurPostProcess.
type NewBlurPostProcessOpts struct {
	SamplingMode     *float64
	Engine           *Engine
	Reusable         *bool
	TextureType      *float64
	Defines          *string
	BlockCompilation *bool
}

// NewBlurPostProcess returns a new BlurPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess
func (ba *Babylon) NewBlurPostProcess(name string, direction *Vector2, kernel float64, options float64, camera *Camera, opts *NewBlurPostProcessOpts) *BlurPostProcess {
	if opts == nil {
		opts = &NewBlurPostProcessOpts{}
	}

	args := make([]interface{}, 0, 5+6)

	args = append(args, name)
	args = append(args, direction.JSObject())
	args = append(args, kernel)
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
	if opts.Defines == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Defines)
	}
	if opts.BlockCompilation == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.BlockCompilation)
	}

	p := ba.ctx.Get("BlurPostProcess").New(args...)
	return BlurPostProcessFromJSObject(p, ba.ctx)
}

// BlurPostProcessActivateOpts contains optional parameters for BlurPostProcess.Activate.
type BlurPostProcessActivateOpts struct {
	SourceTexture     *InternalTexture
	ForceDepthStencil *bool
}

// Activate calls the Activate method on the BlurPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess#activate
func (b *BlurPostProcess) Activate(camera *Camera, opts *BlurPostProcessActivateOpts) *InternalTexture {
	if opts == nil {
		opts = &BlurPostProcessActivateOpts{}
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

	retVal := b.p.Call("activate", args...)
	return InternalTextureFromJSObject(retVal, b.ctx)
}

// Apply calls the Apply method on the BlurPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess#apply
func (b *BlurPostProcess) Apply() *Effect {

	retVal := b.p.Call("apply")
	return EffectFromJSObject(retVal, b.ctx)
}

// BlurPostProcessDisposeOpts contains optional parameters for BlurPostProcess.Dispose.
type BlurPostProcessDisposeOpts struct {
	Camera *Camera
}

// Dispose calls the Dispose method on the BlurPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess#dispose
func (b *BlurPostProcess) Dispose(opts *BlurPostProcessDisposeOpts) {
	if opts == nil {
		opts = &BlurPostProcessDisposeOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.Camera == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Camera.JSObject())
	}

	b.p.Call("dispose", args...)
}

// GetCamera calls the GetCamera method on the BlurPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess#getcamera
func (b *BlurPostProcess) GetCamera() *Camera {

	retVal := b.p.Call("getCamera")
	return CameraFromJSObject(retVal, b.ctx)
}

// GetClassName calls the GetClassName method on the BlurPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess#getclassname
func (b *BlurPostProcess) GetClassName() string {

	retVal := b.p.Call("getClassName")
	return retVal.String()
}

// GetEffect calls the GetEffect method on the BlurPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess#geteffect
func (b *BlurPostProcess) GetEffect() *Effect {

	retVal := b.p.Call("getEffect")
	return EffectFromJSObject(retVal, b.ctx)
}

// GetEffectName calls the GetEffectName method on the BlurPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess#geteffectname
func (b *BlurPostProcess) GetEffectName() string {

	retVal := b.p.Call("getEffectName")
	return retVal.String()
}

// GetEngine calls the GetEngine method on the BlurPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess#getengine
func (b *BlurPostProcess) GetEngine() *Engine {

	retVal := b.p.Call("getEngine")
	return EngineFromJSObject(retVal, b.ctx)
}

// IsReady calls the IsReady method on the BlurPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess#isready
func (b *BlurPostProcess) IsReady() bool {

	retVal := b.p.Call("isReady")
	return retVal.Bool()
}

// IsReusable calls the IsReusable method on the BlurPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess#isreusable
func (b *BlurPostProcess) IsReusable() bool {

	retVal := b.p.Call("isReusable")
	return retVal.Bool()
}

// MarkTextureDirty calls the MarkTextureDirty method on the BlurPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess#marktexturedirty
func (b *BlurPostProcess) MarkTextureDirty() {

	b.p.Call("markTextureDirty")
}

// ShareOutputWith calls the ShareOutputWith method on the BlurPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess#shareoutputwith
func (b *BlurPostProcess) ShareOutputWith(postProcess *PostProcess) *PostProcess {

	args := make([]interface{}, 0, 1+0)

	args = append(args, postProcess.JSObject())

	retVal := b.p.Call("shareOutputWith", args...)
	return PostProcessFromJSObject(retVal, b.ctx)
}

// BlurPostProcessUpdateEffectOpts contains optional parameters for BlurPostProcess.UpdateEffect.
type BlurPostProcessUpdateEffectOpts struct {
	Defines         *string
	Uniforms        *string
	Samplers        *string
	IndexParameters *interface{}
	OnCompiled      func()
	OnError         func()
}

// UpdateEffect calls the UpdateEffect method on the BlurPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess#updateeffect
func (b *BlurPostProcess) UpdateEffect(opts *BlurPostProcessUpdateEffectOpts) {
	if opts == nil {
		opts = &BlurPostProcessUpdateEffectOpts{}
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

	b.p.Call("updateEffect", args...)
}

// UseOwnOutput calls the UseOwnOutput method on the BlurPostProcess object.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess#useownoutput
func (b *BlurPostProcess) UseOwnOutput() {

	b.p.Call("useOwnOutput")
}

/*

// AdaptScaleToCurrentViewport returns the AdaptScaleToCurrentViewport property of class BlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess#adaptscaletocurrentviewport
func (b *BlurPostProcess) AdaptScaleToCurrentViewport(adaptScaleToCurrentViewport bool) *BlurPostProcess {
	p := ba.ctx.Get("BlurPostProcess").New(adaptScaleToCurrentViewport)
	return BlurPostProcessFromJSObject(p, ba.ctx)
}

// SetAdaptScaleToCurrentViewport sets the AdaptScaleToCurrentViewport property of class BlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess#adaptscaletocurrentviewport
func (b *BlurPostProcess) SetAdaptScaleToCurrentViewport(adaptScaleToCurrentViewport bool) *BlurPostProcess {
	p := ba.ctx.Get("BlurPostProcess").New(adaptScaleToCurrentViewport)
	return BlurPostProcessFromJSObject(p, ba.ctx)
}

// AlphaConstants returns the AlphaConstants property of class BlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess#alphaconstants
func (b *BlurPostProcess) AlphaConstants(alphaConstants *Color4) *BlurPostProcess {
	p := ba.ctx.Get("BlurPostProcess").New(alphaConstants.JSObject())
	return BlurPostProcessFromJSObject(p, ba.ctx)
}

// SetAlphaConstants sets the AlphaConstants property of class BlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess#alphaconstants
func (b *BlurPostProcess) SetAlphaConstants(alphaConstants *Color4) *BlurPostProcess {
	p := ba.ctx.Get("BlurPostProcess").New(alphaConstants.JSObject())
	return BlurPostProcessFromJSObject(p, ba.ctx)
}

// AlphaMode returns the AlphaMode property of class BlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess#alphamode
func (b *BlurPostProcess) AlphaMode(alphaMode float64) *BlurPostProcess {
	p := ba.ctx.Get("BlurPostProcess").New(alphaMode)
	return BlurPostProcessFromJSObject(p, ba.ctx)
}

// SetAlphaMode sets the AlphaMode property of class BlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess#alphamode
func (b *BlurPostProcess) SetAlphaMode(alphaMode float64) *BlurPostProcess {
	p := ba.ctx.Get("BlurPostProcess").New(alphaMode)
	return BlurPostProcessFromJSObject(p, ba.ctx)
}

// AlwaysForcePOT returns the AlwaysForcePOT property of class BlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess#alwaysforcepot
func (b *BlurPostProcess) AlwaysForcePOT(alwaysForcePOT bool) *BlurPostProcess {
	p := ba.ctx.Get("BlurPostProcess").New(alwaysForcePOT)
	return BlurPostProcessFromJSObject(p, ba.ctx)
}

// SetAlwaysForcePOT sets the AlwaysForcePOT property of class BlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess#alwaysforcepot
func (b *BlurPostProcess) SetAlwaysForcePOT(alwaysForcePOT bool) *BlurPostProcess {
	p := ba.ctx.Get("BlurPostProcess").New(alwaysForcePOT)
	return BlurPostProcessFromJSObject(p, ba.ctx)
}

// Animations returns the Animations property of class BlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess#animations
func (b *BlurPostProcess) Animations(animations *Animation) *BlurPostProcess {
	p := ba.ctx.Get("BlurPostProcess").New(animations.JSObject())
	return BlurPostProcessFromJSObject(p, ba.ctx)
}

// SetAnimations sets the Animations property of class BlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess#animations
func (b *BlurPostProcess) SetAnimations(animations *Animation) *BlurPostProcess {
	p := ba.ctx.Get("BlurPostProcess").New(animations.JSObject())
	return BlurPostProcessFromJSObject(p, ba.ctx)
}

// AspectRatio returns the AspectRatio property of class BlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess#aspectratio
func (b *BlurPostProcess) AspectRatio(aspectRatio float64) *BlurPostProcess {
	p := ba.ctx.Get("BlurPostProcess").New(aspectRatio)
	return BlurPostProcessFromJSObject(p, ba.ctx)
}

// SetAspectRatio sets the AspectRatio property of class BlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess#aspectratio
func (b *BlurPostProcess) SetAspectRatio(aspectRatio float64) *BlurPostProcess {
	p := ba.ctx.Get("BlurPostProcess").New(aspectRatio)
	return BlurPostProcessFromJSObject(p, ba.ctx)
}

// AutoClear returns the AutoClear property of class BlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess#autoclear
func (b *BlurPostProcess) AutoClear(autoClear bool) *BlurPostProcess {
	p := ba.ctx.Get("BlurPostProcess").New(autoClear)
	return BlurPostProcessFromJSObject(p, ba.ctx)
}

// SetAutoClear sets the AutoClear property of class BlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess#autoclear
func (b *BlurPostProcess) SetAutoClear(autoClear bool) *BlurPostProcess {
	p := ba.ctx.Get("BlurPostProcess").New(autoClear)
	return BlurPostProcessFromJSObject(p, ba.ctx)
}

// ClearColor returns the ClearColor property of class BlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess#clearcolor
func (b *BlurPostProcess) ClearColor(clearColor *Color4) *BlurPostProcess {
	p := ba.ctx.Get("BlurPostProcess").New(clearColor.JSObject())
	return BlurPostProcessFromJSObject(p, ba.ctx)
}

// SetClearColor sets the ClearColor property of class BlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess#clearcolor
func (b *BlurPostProcess) SetClearColor(clearColor *Color4) *BlurPostProcess {
	p := ba.ctx.Get("BlurPostProcess").New(clearColor.JSObject())
	return BlurPostProcessFromJSObject(p, ba.ctx)
}

// Direction returns the Direction property of class BlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess#direction
func (b *BlurPostProcess) Direction(direction *Vector2) *BlurPostProcess {
	p := ba.ctx.Get("BlurPostProcess").New(direction.JSObject())
	return BlurPostProcessFromJSObject(p, ba.ctx)
}

// SetDirection sets the Direction property of class BlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess#direction
func (b *BlurPostProcess) SetDirection(direction *Vector2) *BlurPostProcess {
	p := ba.ctx.Get("BlurPostProcess").New(direction.JSObject())
	return BlurPostProcessFromJSObject(p, ba.ctx)
}

// EnablePixelPerfectMode returns the EnablePixelPerfectMode property of class BlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess#enablepixelperfectmode
func (b *BlurPostProcess) EnablePixelPerfectMode(enablePixelPerfectMode bool) *BlurPostProcess {
	p := ba.ctx.Get("BlurPostProcess").New(enablePixelPerfectMode)
	return BlurPostProcessFromJSObject(p, ba.ctx)
}

// SetEnablePixelPerfectMode sets the EnablePixelPerfectMode property of class BlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess#enablepixelperfectmode
func (b *BlurPostProcess) SetEnablePixelPerfectMode(enablePixelPerfectMode bool) *BlurPostProcess {
	p := ba.ctx.Get("BlurPostProcess").New(enablePixelPerfectMode)
	return BlurPostProcessFromJSObject(p, ba.ctx)
}

// ForceFullscreenViewport returns the ForceFullscreenViewport property of class BlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess#forcefullscreenviewport
func (b *BlurPostProcess) ForceFullscreenViewport(forceFullscreenViewport bool) *BlurPostProcess {
	p := ba.ctx.Get("BlurPostProcess").New(forceFullscreenViewport)
	return BlurPostProcessFromJSObject(p, ba.ctx)
}

// SetForceFullscreenViewport sets the ForceFullscreenViewport property of class BlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess#forcefullscreenviewport
func (b *BlurPostProcess) SetForceFullscreenViewport(forceFullscreenViewport bool) *BlurPostProcess {
	p := ba.ctx.Get("BlurPostProcess").New(forceFullscreenViewport)
	return BlurPostProcessFromJSObject(p, ba.ctx)
}

// Height returns the Height property of class BlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess#height
func (b *BlurPostProcess) Height(height float64) *BlurPostProcess {
	p := ba.ctx.Get("BlurPostProcess").New(height)
	return BlurPostProcessFromJSObject(p, ba.ctx)
}

// SetHeight sets the Height property of class BlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess#height
func (b *BlurPostProcess) SetHeight(height float64) *BlurPostProcess {
	p := ba.ctx.Get("BlurPostProcess").New(height)
	return BlurPostProcessFromJSObject(p, ba.ctx)
}

// InputTexture returns the InputTexture property of class BlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess#inputtexture
func (b *BlurPostProcess) InputTexture(inputTexture *InternalTexture) *BlurPostProcess {
	p := ba.ctx.Get("BlurPostProcess").New(inputTexture.JSObject())
	return BlurPostProcessFromJSObject(p, ba.ctx)
}

// SetInputTexture sets the InputTexture property of class BlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess#inputtexture
func (b *BlurPostProcess) SetInputTexture(inputTexture *InternalTexture) *BlurPostProcess {
	p := ba.ctx.Get("BlurPostProcess").New(inputTexture.JSObject())
	return BlurPostProcessFromJSObject(p, ba.ctx)
}

// InspectableCustomProperties returns the InspectableCustomProperties property of class BlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess#inspectablecustomproperties
func (b *BlurPostProcess) InspectableCustomProperties(inspectableCustomProperties *IInspectable) *BlurPostProcess {
	p := ba.ctx.Get("BlurPostProcess").New(inspectableCustomProperties.JSObject())
	return BlurPostProcessFromJSObject(p, ba.ctx)
}

// SetInspectableCustomProperties sets the InspectableCustomProperties property of class BlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess#inspectablecustomproperties
func (b *BlurPostProcess) SetInspectableCustomProperties(inspectableCustomProperties *IInspectable) *BlurPostProcess {
	p := ba.ctx.Get("BlurPostProcess").New(inspectableCustomProperties.JSObject())
	return BlurPostProcessFromJSObject(p, ba.ctx)
}

// IsSupported returns the IsSupported property of class BlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess#issupported
func (b *BlurPostProcess) IsSupported(isSupported bool) *BlurPostProcess {
	p := ba.ctx.Get("BlurPostProcess").New(isSupported)
	return BlurPostProcessFromJSObject(p, ba.ctx)
}

// SetIsSupported sets the IsSupported property of class BlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess#issupported
func (b *BlurPostProcess) SetIsSupported(isSupported bool) *BlurPostProcess {
	p := ba.ctx.Get("BlurPostProcess").New(isSupported)
	return BlurPostProcessFromJSObject(p, ba.ctx)
}

// Kernel returns the Kernel property of class BlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess#kernel
func (b *BlurPostProcess) Kernel(kernel float64) *BlurPostProcess {
	p := ba.ctx.Get("BlurPostProcess").New(kernel)
	return BlurPostProcessFromJSObject(p, ba.ctx)
}

// SetKernel sets the Kernel property of class BlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess#kernel
func (b *BlurPostProcess) SetKernel(kernel float64) *BlurPostProcess {
	p := ba.ctx.Get("BlurPostProcess").New(kernel)
	return BlurPostProcessFromJSObject(p, ba.ctx)
}

// Name returns the Name property of class BlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess#name
func (b *BlurPostProcess) Name(name string) *BlurPostProcess {
	p := ba.ctx.Get("BlurPostProcess").New(name)
	return BlurPostProcessFromJSObject(p, ba.ctx)
}

// SetName sets the Name property of class BlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess#name
func (b *BlurPostProcess) SetName(name string) *BlurPostProcess {
	p := ba.ctx.Get("BlurPostProcess").New(name)
	return BlurPostProcessFromJSObject(p, ba.ctx)
}

// OnActivate returns the OnActivate property of class BlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess#onactivate
func (b *BlurPostProcess) OnActivate(onActivate func()) *BlurPostProcess {
	p := ba.ctx.Get("BlurPostProcess").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {onActivate(); return nil}))
	return BlurPostProcessFromJSObject(p, ba.ctx)
}

// SetOnActivate sets the OnActivate property of class BlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess#onactivate
func (b *BlurPostProcess) SetOnActivate(onActivate func()) *BlurPostProcess {
	p := ba.ctx.Get("BlurPostProcess").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {onActivate(); return nil}))
	return BlurPostProcessFromJSObject(p, ba.ctx)
}

// OnActivateObservable returns the OnActivateObservable property of class BlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess#onactivateobservable
func (b *BlurPostProcess) OnActivateObservable(onActivateObservable *Observable) *BlurPostProcess {
	p := ba.ctx.Get("BlurPostProcess").New(onActivateObservable.JSObject())
	return BlurPostProcessFromJSObject(p, ba.ctx)
}

// SetOnActivateObservable sets the OnActivateObservable property of class BlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess#onactivateobservable
func (b *BlurPostProcess) SetOnActivateObservable(onActivateObservable *Observable) *BlurPostProcess {
	p := ba.ctx.Get("BlurPostProcess").New(onActivateObservable.JSObject())
	return BlurPostProcessFromJSObject(p, ba.ctx)
}

// OnAfterRender returns the OnAfterRender property of class BlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess#onafterrender
func (b *BlurPostProcess) OnAfterRender(onAfterRender func()) *BlurPostProcess {
	p := ba.ctx.Get("BlurPostProcess").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {onAfterRender(); return nil}))
	return BlurPostProcessFromJSObject(p, ba.ctx)
}

// SetOnAfterRender sets the OnAfterRender property of class BlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess#onafterrender
func (b *BlurPostProcess) SetOnAfterRender(onAfterRender func()) *BlurPostProcess {
	p := ba.ctx.Get("BlurPostProcess").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {onAfterRender(); return nil}))
	return BlurPostProcessFromJSObject(p, ba.ctx)
}

// OnAfterRenderObservable returns the OnAfterRenderObservable property of class BlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess#onafterrenderobservable
func (b *BlurPostProcess) OnAfterRenderObservable(onAfterRenderObservable *Observable) *BlurPostProcess {
	p := ba.ctx.Get("BlurPostProcess").New(onAfterRenderObservable.JSObject())
	return BlurPostProcessFromJSObject(p, ba.ctx)
}

// SetOnAfterRenderObservable sets the OnAfterRenderObservable property of class BlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess#onafterrenderobservable
func (b *BlurPostProcess) SetOnAfterRenderObservable(onAfterRenderObservable *Observable) *BlurPostProcess {
	p := ba.ctx.Get("BlurPostProcess").New(onAfterRenderObservable.JSObject())
	return BlurPostProcessFromJSObject(p, ba.ctx)
}

// OnApply returns the OnApply property of class BlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess#onapply
func (b *BlurPostProcess) OnApply(onApply func()) *BlurPostProcess {
	p := ba.ctx.Get("BlurPostProcess").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {onApply(); return nil}))
	return BlurPostProcessFromJSObject(p, ba.ctx)
}

// SetOnApply sets the OnApply property of class BlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess#onapply
func (b *BlurPostProcess) SetOnApply(onApply func()) *BlurPostProcess {
	p := ba.ctx.Get("BlurPostProcess").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {onApply(); return nil}))
	return BlurPostProcessFromJSObject(p, ba.ctx)
}

// OnApplyObservable returns the OnApplyObservable property of class BlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess#onapplyobservable
func (b *BlurPostProcess) OnApplyObservable(onApplyObservable *Observable) *BlurPostProcess {
	p := ba.ctx.Get("BlurPostProcess").New(onApplyObservable.JSObject())
	return BlurPostProcessFromJSObject(p, ba.ctx)
}

// SetOnApplyObservable sets the OnApplyObservable property of class BlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess#onapplyobservable
func (b *BlurPostProcess) SetOnApplyObservable(onApplyObservable *Observable) *BlurPostProcess {
	p := ba.ctx.Get("BlurPostProcess").New(onApplyObservable.JSObject())
	return BlurPostProcessFromJSObject(p, ba.ctx)
}

// OnBeforeRender returns the OnBeforeRender property of class BlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess#onbeforerender
func (b *BlurPostProcess) OnBeforeRender(onBeforeRender func()) *BlurPostProcess {
	p := ba.ctx.Get("BlurPostProcess").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {onBeforeRender(); return nil}))
	return BlurPostProcessFromJSObject(p, ba.ctx)
}

// SetOnBeforeRender sets the OnBeforeRender property of class BlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess#onbeforerender
func (b *BlurPostProcess) SetOnBeforeRender(onBeforeRender func()) *BlurPostProcess {
	p := ba.ctx.Get("BlurPostProcess").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {onBeforeRender(); return nil}))
	return BlurPostProcessFromJSObject(p, ba.ctx)
}

// OnBeforeRenderObservable returns the OnBeforeRenderObservable property of class BlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess#onbeforerenderobservable
func (b *BlurPostProcess) OnBeforeRenderObservable(onBeforeRenderObservable *Observable) *BlurPostProcess {
	p := ba.ctx.Get("BlurPostProcess").New(onBeforeRenderObservable.JSObject())
	return BlurPostProcessFromJSObject(p, ba.ctx)
}

// SetOnBeforeRenderObservable sets the OnBeforeRenderObservable property of class BlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess#onbeforerenderobservable
func (b *BlurPostProcess) SetOnBeforeRenderObservable(onBeforeRenderObservable *Observable) *BlurPostProcess {
	p := ba.ctx.Get("BlurPostProcess").New(onBeforeRenderObservable.JSObject())
	return BlurPostProcessFromJSObject(p, ba.ctx)
}

// OnSizeChanged returns the OnSizeChanged property of class BlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess#onsizechanged
func (b *BlurPostProcess) OnSizeChanged(onSizeChanged func()) *BlurPostProcess {
	p := ba.ctx.Get("BlurPostProcess").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {onSizeChanged(); return nil}))
	return BlurPostProcessFromJSObject(p, ba.ctx)
}

// SetOnSizeChanged sets the OnSizeChanged property of class BlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess#onsizechanged
func (b *BlurPostProcess) SetOnSizeChanged(onSizeChanged func()) *BlurPostProcess {
	p := ba.ctx.Get("BlurPostProcess").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {onSizeChanged(); return nil}))
	return BlurPostProcessFromJSObject(p, ba.ctx)
}

// OnSizeChangedObservable returns the OnSizeChangedObservable property of class BlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess#onsizechangedobservable
func (b *BlurPostProcess) OnSizeChangedObservable(onSizeChangedObservable *Observable) *BlurPostProcess {
	p := ba.ctx.Get("BlurPostProcess").New(onSizeChangedObservable.JSObject())
	return BlurPostProcessFromJSObject(p, ba.ctx)
}

// SetOnSizeChangedObservable sets the OnSizeChangedObservable property of class BlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess#onsizechangedobservable
func (b *BlurPostProcess) SetOnSizeChangedObservable(onSizeChangedObservable *Observable) *BlurPostProcess {
	p := ba.ctx.Get("BlurPostProcess").New(onSizeChangedObservable.JSObject())
	return BlurPostProcessFromJSObject(p, ba.ctx)
}

// PackedFloat returns the PackedFloat property of class BlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess#packedfloat
func (b *BlurPostProcess) PackedFloat(packedFloat bool) *BlurPostProcess {
	p := ba.ctx.Get("BlurPostProcess").New(packedFloat)
	return BlurPostProcessFromJSObject(p, ba.ctx)
}

// SetPackedFloat sets the PackedFloat property of class BlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess#packedfloat
func (b *BlurPostProcess) SetPackedFloat(packedFloat bool) *BlurPostProcess {
	p := ba.ctx.Get("BlurPostProcess").New(packedFloat)
	return BlurPostProcessFromJSObject(p, ba.ctx)
}

// RenderTargetSamplingMode returns the RenderTargetSamplingMode property of class BlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess#rendertargetsamplingmode
func (b *BlurPostProcess) RenderTargetSamplingMode(renderTargetSamplingMode float64) *BlurPostProcess {
	p := ba.ctx.Get("BlurPostProcess").New(renderTargetSamplingMode)
	return BlurPostProcessFromJSObject(p, ba.ctx)
}

// SetRenderTargetSamplingMode sets the RenderTargetSamplingMode property of class BlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess#rendertargetsamplingmode
func (b *BlurPostProcess) SetRenderTargetSamplingMode(renderTargetSamplingMode float64) *BlurPostProcess {
	p := ba.ctx.Get("BlurPostProcess").New(renderTargetSamplingMode)
	return BlurPostProcessFromJSObject(p, ba.ctx)
}

// Samples returns the Samples property of class BlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess#samples
func (b *BlurPostProcess) Samples(samples float64) *BlurPostProcess {
	p := ba.ctx.Get("BlurPostProcess").New(samples)
	return BlurPostProcessFromJSObject(p, ba.ctx)
}

// SetSamples sets the Samples property of class BlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess#samples
func (b *BlurPostProcess) SetSamples(samples float64) *BlurPostProcess {
	p := ba.ctx.Get("BlurPostProcess").New(samples)
	return BlurPostProcessFromJSObject(p, ba.ctx)
}

// ScaleMode returns the ScaleMode property of class BlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess#scalemode
func (b *BlurPostProcess) ScaleMode(scaleMode float64) *BlurPostProcess {
	p := ba.ctx.Get("BlurPostProcess").New(scaleMode)
	return BlurPostProcessFromJSObject(p, ba.ctx)
}

// SetScaleMode sets the ScaleMode property of class BlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess#scalemode
func (b *BlurPostProcess) SetScaleMode(scaleMode float64) *BlurPostProcess {
	p := ba.ctx.Get("BlurPostProcess").New(scaleMode)
	return BlurPostProcessFromJSObject(p, ba.ctx)
}

// TexelSize returns the TexelSize property of class BlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess#texelsize
func (b *BlurPostProcess) TexelSize(texelSize *Vector2) *BlurPostProcess {
	p := ba.ctx.Get("BlurPostProcess").New(texelSize.JSObject())
	return BlurPostProcessFromJSObject(p, ba.ctx)
}

// SetTexelSize sets the TexelSize property of class BlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess#texelsize
func (b *BlurPostProcess) SetTexelSize(texelSize *Vector2) *BlurPostProcess {
	p := ba.ctx.Get("BlurPostProcess").New(texelSize.JSObject())
	return BlurPostProcessFromJSObject(p, ba.ctx)
}

// UniqueId returns the UniqueId property of class BlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess#uniqueid
func (b *BlurPostProcess) UniqueId(uniqueId float64) *BlurPostProcess {
	p := ba.ctx.Get("BlurPostProcess").New(uniqueId)
	return BlurPostProcessFromJSObject(p, ba.ctx)
}

// SetUniqueId sets the UniqueId property of class BlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess#uniqueid
func (b *BlurPostProcess) SetUniqueId(uniqueId float64) *BlurPostProcess {
	p := ba.ctx.Get("BlurPostProcess").New(uniqueId)
	return BlurPostProcessFromJSObject(p, ba.ctx)
}

// Width returns the Width property of class BlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess#width
func (b *BlurPostProcess) Width(width float64) *BlurPostProcess {
	p := ba.ctx.Get("BlurPostProcess").New(width)
	return BlurPostProcessFromJSObject(p, ba.ctx)
}

// SetWidth sets the Width property of class BlurPostProcess.
//
// https://doc.babylonjs.com/api/classes/babylon.blurpostprocess#width
func (b *BlurPostProcess) SetWidth(width float64) *BlurPostProcess {
	p := ba.ctx.Get("BlurPostProcess").New(width)
	return BlurPostProcessFromJSObject(p, ba.ctx)
}

*/
