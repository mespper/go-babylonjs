// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// NullEngine represents a babylon.js NullEngine.
// The null engine class provides support for headless version of babylon.js.
// This can be used in server side scenario or for testing purposes
type NullEngine struct {
	*Engine
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (n *NullEngine) JSObject() js.Value { return n.p }

// NullEngine returns a NullEngine JavaScript class.
func (ba *Babylon) NullEngine() *NullEngine {
	p := ba.ctx.Get("NullEngine")
	return NullEngineFromJSObject(p, ba.ctx)
}

// NullEngineFromJSObject returns a wrapped NullEngine JavaScript class.
func NullEngineFromJSObject(p js.Value, ctx js.Value) *NullEngine {
	return &NullEngine{Engine: EngineFromJSObject(p, ctx), ctx: ctx}
}

// NullEngineArrayToJSArray returns a JavaScript Array for the wrapped array.
func NullEngineArrayToJSArray(array []*NullEngine) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewNullEngineOpts contains optional parameters for NewNullEngine.
type NewNullEngineOpts struct {
	Options *NullEngineOptions
}

// NewNullEngine returns a new NullEngine object.
//
// https://doc.babylonjs.com/api/classes/babylon.nullengine
func (ba *Babylon) NewNullEngine(opts *NewNullEngineOpts) *NullEngine {
	if opts == nil {
		opts = &NewNullEngineOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.Options == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Options.JSObject())
	}

	p := ba.ctx.Get("NullEngine").New(args...)
	return NullEngineFromJSObject(p, ba.ctx)
}

// AreAllEffectsReady calls the AreAllEffectsReady method on the NullEngine object.
//
// https://doc.babylonjs.com/api/classes/babylon.nullengine#arealleffectsready
func (n *NullEngine) AreAllEffectsReady() bool {

	retVal := n.p.Call("areAllEffectsReady")
	return retVal.Bool()
}

// BindBuffers calls the BindBuffers method on the NullEngine object.
//
// https://doc.babylonjs.com/api/classes/babylon.nullengine#bindbuffers
func (n *NullEngine) BindBuffers(vertexBuffers js.Value, indexBuffer *DataBuffer, effect *Effect) {

	args := make([]interface{}, 0, 3+0)

	args = append(args, vertexBuffers)
	args = append(args, indexBuffer.JSObject())
	args = append(args, effect.JSObject())

	n.p.Call("bindBuffers", args...)
}

// NullEngineBindFramebufferOpts contains optional parameters for NullEngine.BindFramebuffer.
type NullEngineBindFramebufferOpts struct {
	FaceIndex               *float64
	RequiredWidth           *float64
	RequiredHeight          *float64
	ForceFullscreenViewport *bool
}

// BindFramebuffer calls the BindFramebuffer method on the NullEngine object.
//
// https://doc.babylonjs.com/api/classes/babylon.nullengine#bindframebuffer
func (n *NullEngine) BindFramebuffer(texture *InternalTexture, opts *NullEngineBindFramebufferOpts) {
	if opts == nil {
		opts = &NullEngineBindFramebufferOpts{}
	}

	args := make([]interface{}, 0, 1+4)

	args = append(args, texture.JSObject())

	if opts.FaceIndex == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.FaceIndex)
	}
	if opts.RequiredWidth == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.RequiredWidth)
	}
	if opts.RequiredHeight == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.RequiredHeight)
	}
	if opts.ForceFullscreenViewport == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.ForceFullscreenViewport)
	}

	n.p.Call("bindFramebuffer", args...)
}

// BindSamplers calls the BindSamplers method on the NullEngine object.
//
// https://doc.babylonjs.com/api/classes/babylon.nullengine#bindsamplers
func (n *NullEngine) BindSamplers(effect *Effect) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, effect.JSObject())

	n.p.Call("bindSamplers", args...)
}

// NullEngineClearOpts contains optional parameters for NullEngine.Clear.
type NullEngineClearOpts struct {
	Stencil *bool
}

// Clear calls the Clear method on the NullEngine object.
//
// https://doc.babylonjs.com/api/classes/babylon.nullengine#clear
func (n *NullEngine) Clear(color js.Value, backBuffer bool, depth bool, opts *NullEngineClearOpts) {
	if opts == nil {
		opts = &NullEngineClearOpts{}
	}

	args := make([]interface{}, 0, 3+1)

	args = append(args, color)
	args = append(args, backBuffer)
	args = append(args, depth)

	if opts.Stencil == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Stencil)
	}

	n.p.Call("clear", args...)
}

// CreateDynamicVertexBuffer calls the CreateDynamicVertexBuffer method on the NullEngine object.
//
// https://doc.babylonjs.com/api/classes/babylon.nullengine#createdynamicvertexbuffer
func (n *NullEngine) CreateDynamicVertexBuffer(vertices js.Value) *DataBuffer {

	args := make([]interface{}, 0, 1+0)

	args = append(args, vertices)

	retVal := n.p.Call("createDynamicVertexBuffer", args...)
	return DataBufferFromJSObject(retVal, n.ctx)
}

// CreateIndexBuffer calls the CreateIndexBuffer method on the NullEngine object.
//
// https://doc.babylonjs.com/api/classes/babylon.nullengine#createindexbuffer
func (n *NullEngine) CreateIndexBuffer(indices js.Value) *DataBuffer {

	args := make([]interface{}, 0, 1+0)

	args = append(args, indices)

	retVal := n.p.Call("createIndexBuffer", args...)
	return DataBufferFromJSObject(retVal, n.ctx)
}

// CreateRenderTargetTexture calls the CreateRenderTargetTexture method on the NullEngine object.
//
// https://doc.babylonjs.com/api/classes/babylon.nullengine#createrendertargettexture
func (n *NullEngine) CreateRenderTargetTexture(size interface{}, options bool) *InternalTexture {

	args := make([]interface{}, 0, 2+0)

	args = append(args, size)
	args = append(args, options)

	retVal := n.p.Call("createRenderTargetTexture", args...)
	return InternalTextureFromJSObject(retVal, n.ctx)
}

// NullEngineCreateShaderProgramOpts contains optional parameters for NullEngine.CreateShaderProgram.
type NullEngineCreateShaderProgramOpts struct {
	Context js.Value
}

// CreateShaderProgram calls the CreateShaderProgram method on the NullEngine object.
//
// https://doc.babylonjs.com/api/classes/babylon.nullengine#createshaderprogram
func (n *NullEngine) CreateShaderProgram(pipelineContext *IPipelineContext, vertexCode string, fragmentCode string, defines string, opts *NullEngineCreateShaderProgramOpts) js.Value {
	if opts == nil {
		opts = &NullEngineCreateShaderProgramOpts{}
	}

	args := make([]interface{}, 0, 4+1)

	args = append(args, pipelineContext.JSObject())
	args = append(args, vertexCode)
	args = append(args, fragmentCode)
	args = append(args, defines)

	args = append(args, opts.Context)

	retVal := n.p.Call("createShaderProgram", args...)
	return retVal
}

// NullEngineCreateTextureOpts contains optional parameters for NullEngine.CreateTexture.
type NullEngineCreateTextureOpts struct {
	SamplingMode *float64
	OnLoad       JSFunc
	OnError      JSFunc
	Buffer       js.Value
	FallBack     *InternalTexture
	Format       *float64
}

// CreateTexture calls the CreateTexture method on the NullEngine object.
//
// https://doc.babylonjs.com/api/classes/babylon.nullengine#createtexture
func (n *NullEngine) CreateTexture(urlArg string, noMipmap bool, invertY bool, scene *Scene, opts *NullEngineCreateTextureOpts) *InternalTexture {
	if opts == nil {
		opts = &NullEngineCreateTextureOpts{}
	}

	args := make([]interface{}, 0, 4+6)

	args = append(args, urlArg)
	args = append(args, noMipmap)
	args = append(args, invertY)
	args = append(args, scene.JSObject())

	if opts.SamplingMode == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.SamplingMode)
	}
	if opts.OnLoad == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, js.FuncOf(opts.OnLoad) /* never freed! */)
	}
	if opts.OnError == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, js.FuncOf(opts.OnError) /* never freed! */)
	}
	args = append(args, opts.Buffer)
	if opts.FallBack == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.FallBack.JSObject())
	}
	if opts.Format == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Format)
	}

	retVal := n.p.Call("createTexture", args...)
	return InternalTextureFromJSObject(retVal, n.ctx)
}

// CreateVertexBuffer calls the CreateVertexBuffer method on the NullEngine object.
//
// https://doc.babylonjs.com/api/classes/babylon.nullengine#createvertexbuffer
func (n *NullEngine) CreateVertexBuffer(vertices js.Value) *DataBuffer {

	args := make([]interface{}, 0, 1+0)

	args = append(args, vertices)

	retVal := n.p.Call("createVertexBuffer", args...)
	return DataBufferFromJSObject(retVal, n.ctx)
}

// DisplayLoadingUI calls the DisplayLoadingUI method on the NullEngine object.
//
// https://doc.babylonjs.com/api/classes/babylon.nullengine#displayloadingui
func (n *NullEngine) DisplayLoadingUI() {

	n.p.Call("displayLoadingUI")
}

// NullEngineDrawOpts contains optional parameters for NullEngine.Draw.
type NullEngineDrawOpts struct {
	InstancesCount *float64
}

// Draw calls the Draw method on the NullEngine object.
//
// https://doc.babylonjs.com/api/classes/babylon.nullengine#draw
func (n *NullEngine) Draw(useTriangles bool, indexStart float64, indexCount float64, opts *NullEngineDrawOpts) {
	if opts == nil {
		opts = &NullEngineDrawOpts{}
	}

	args := make([]interface{}, 0, 3+1)

	args = append(args, useTriangles)
	args = append(args, indexStart)
	args = append(args, indexCount)

	if opts.InstancesCount == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.InstancesCount)
	}

	n.p.Call("draw", args...)
}

// NullEngineDrawArraysTypeOpts contains optional parameters for NullEngine.DrawArraysType.
type NullEngineDrawArraysTypeOpts struct {
	InstancesCount *float64
}

// DrawArraysType calls the DrawArraysType method on the NullEngine object.
//
// https://doc.babylonjs.com/api/classes/babylon.nullengine#drawarraystype
func (n *NullEngine) DrawArraysType(fillMode float64, verticesStart float64, verticesCount float64, opts *NullEngineDrawArraysTypeOpts) {
	if opts == nil {
		opts = &NullEngineDrawArraysTypeOpts{}
	}

	args := make([]interface{}, 0, 3+1)

	args = append(args, fillMode)
	args = append(args, verticesStart)
	args = append(args, verticesCount)

	if opts.InstancesCount == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.InstancesCount)
	}

	n.p.Call("drawArraysType", args...)
}

// NullEngineDrawElementsTypeOpts contains optional parameters for NullEngine.DrawElementsType.
type NullEngineDrawElementsTypeOpts struct {
	InstancesCount *float64
}

// DrawElementsType calls the DrawElementsType method on the NullEngine object.
//
// https://doc.babylonjs.com/api/classes/babylon.nullengine#drawelementstype
func (n *NullEngine) DrawElementsType(fillMode float64, indexStart float64, indexCount float64, opts *NullEngineDrawElementsTypeOpts) {
	if opts == nil {
		opts = &NullEngineDrawElementsTypeOpts{}
	}

	args := make([]interface{}, 0, 3+1)

	args = append(args, fillMode)
	args = append(args, indexStart)
	args = append(args, indexCount)

	if opts.InstancesCount == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.InstancesCount)
	}

	n.p.Call("drawElementsType", args...)
}

// EnableEffect calls the EnableEffect method on the NullEngine object.
//
// https://doc.babylonjs.com/api/classes/babylon.nullengine#enableeffect
func (n *NullEngine) EnableEffect(effect *Effect) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, effect.JSObject())

	n.p.Call("enableEffect", args...)
}

// GetAttributes calls the GetAttributes method on the NullEngine object.
//
// https://doc.babylonjs.com/api/classes/babylon.nullengine#getattributes
func (n *NullEngine) GetAttributes(pipelineContext *IPipelineContext, attributesNames []string) []float64 {

	args := make([]interface{}, 0, 2+0)

	args = append(args, pipelineContext.JSObject())
	args = append(args, attributesNames)

	retVal := n.p.Call("getAttributes", args...)
	result := []float64{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, retVal.Index(ri).Float())
	}
	return result
}

// GetHardwareScalingLevel calls the GetHardwareScalingLevel method on the NullEngine object.
//
// https://doc.babylonjs.com/api/classes/babylon.nullengine#gethardwarescalinglevel
func (n *NullEngine) GetHardwareScalingLevel() float64 {

	retVal := n.p.Call("getHardwareScalingLevel")
	return retVal.Float()
}

// GetLockstepMaxSteps calls the GetLockstepMaxSteps method on the NullEngine object.
//
// https://doc.babylonjs.com/api/classes/babylon.nullengine#getlockstepmaxsteps
func (n *NullEngine) GetLockstepMaxSteps() float64 {

	retVal := n.p.Call("getLockstepMaxSteps")
	return retVal.Float()
}

// NullEngineGetRenderHeightOpts contains optional parameters for NullEngine.GetRenderHeight.
type NullEngineGetRenderHeightOpts struct {
	UseScreen *bool
}

// GetRenderHeight calls the GetRenderHeight method on the NullEngine object.
//
// https://doc.babylonjs.com/api/classes/babylon.nullengine#getrenderheight
func (n *NullEngine) GetRenderHeight(opts *NullEngineGetRenderHeightOpts) float64 {
	if opts == nil {
		opts = &NullEngineGetRenderHeightOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.UseScreen == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.UseScreen)
	}

	retVal := n.p.Call("getRenderHeight", args...)
	return retVal.Float()
}

// NullEngineGetRenderWidthOpts contains optional parameters for NullEngine.GetRenderWidth.
type NullEngineGetRenderWidthOpts struct {
	UseScreen *bool
}

// GetRenderWidth calls the GetRenderWidth method on the NullEngine object.
//
// https://doc.babylonjs.com/api/classes/babylon.nullengine#getrenderwidth
func (n *NullEngine) GetRenderWidth(opts *NullEngineGetRenderWidthOpts) float64 {
	if opts == nil {
		opts = &NullEngineGetRenderWidthOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.UseScreen == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.UseScreen)
	}

	retVal := n.p.Call("getRenderWidth", args...)
	return retVal.Float()
}

// GetUniforms calls the GetUniforms method on the NullEngine object.
//
// https://doc.babylonjs.com/api/classes/babylon.nullengine#getuniforms
func (n *NullEngine) GetUniforms(pipelineContext *IPipelineContext, uniformsNames []string) js.Value {

	args := make([]interface{}, 0, 2+0)

	args = append(args, pipelineContext.JSObject())
	args = append(args, uniformsNames)

	retVal := n.p.Call("getUniforms", args...)
	return retVal
}

// HideLoadingUI calls the HideLoadingUI method on the NullEngine object.
//
// https://doc.babylonjs.com/api/classes/babylon.nullengine#hideloadingui
func (n *NullEngine) HideLoadingUI() {

	n.p.Call("hideLoadingUI")
}

// IsDeterministicLockStep calls the IsDeterministicLockStep method on the NullEngine object.
//
// https://doc.babylonjs.com/api/classes/babylon.nullengine#isdeterministiclockstep
func (n *NullEngine) IsDeterministicLockStep() bool {

	retVal := n.p.Call("isDeterministicLockStep")
	return retVal.Bool()
}

// ReleaseEffects calls the ReleaseEffects method on the NullEngine object.
//
// https://doc.babylonjs.com/api/classes/babylon.nullengine#releaseeffects
func (n *NullEngine) ReleaseEffects() {

	n.p.Call("releaseEffects")
}

// NullEngineSetAlphaModeOpts contains optional parameters for NullEngine.SetAlphaMode.
type NullEngineSetAlphaModeOpts struct {
	NoDepthWriteChange *bool
}

// SetAlphaMode calls the SetAlphaMode method on the NullEngine object.
//
// https://doc.babylonjs.com/api/classes/babylon.nullengine#setalphamode
func (n *NullEngine) SetAlphaMode(mode float64, opts *NullEngineSetAlphaModeOpts) {
	if opts == nil {
		opts = &NullEngineSetAlphaModeOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, mode)

	if opts.NoDepthWriteChange == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.NoDepthWriteChange)
	}

	n.p.Call("setAlphaMode", args...)
}

// SetArray calls the SetArray method on the NullEngine object.
//
// https://doc.babylonjs.com/api/classes/babylon.nullengine#setarray
func (n *NullEngine) SetArray(uniform js.Value, array []float64) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, uniform)
	args = append(args, array)

	n.p.Call("setArray", args...)
}

// SetArray2 calls the SetArray2 method on the NullEngine object.
//
// https://doc.babylonjs.com/api/classes/babylon.nullengine#setarray2
func (n *NullEngine) SetArray2(uniform js.Value, array []float64) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, uniform)
	args = append(args, array)

	n.p.Call("setArray2", args...)
}

// SetArray3 calls the SetArray3 method on the NullEngine object.
//
// https://doc.babylonjs.com/api/classes/babylon.nullengine#setarray3
func (n *NullEngine) SetArray3(uniform js.Value, array []float64) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, uniform)
	args = append(args, array)

	n.p.Call("setArray3", args...)
}

// SetArray4 calls the SetArray4 method on the NullEngine object.
//
// https://doc.babylonjs.com/api/classes/babylon.nullengine#setarray4
func (n *NullEngine) SetArray4(uniform js.Value, array []float64) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, uniform)
	args = append(args, array)

	n.p.Call("setArray4", args...)
}

// SetBool calls the SetBool method on the NullEngine object.
//
// https://doc.babylonjs.com/api/classes/babylon.nullengine#setbool
func (n *NullEngine) SetBool(uniform js.Value, bool float64) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, uniform)
	args = append(args, bool)

	n.p.Call("setBool", args...)
}

// SetFloat calls the SetFloat method on the NullEngine object.
//
// https://doc.babylonjs.com/api/classes/babylon.nullengine#setfloat
func (n *NullEngine) SetFloat(uniform js.Value, value float64) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, uniform)
	args = append(args, value)

	n.p.Call("setFloat", args...)
}

// SetFloat2 calls the SetFloat2 method on the NullEngine object.
//
// https://doc.babylonjs.com/api/classes/babylon.nullengine#setfloat2
func (n *NullEngine) SetFloat2(uniform js.Value, x float64, y float64) {

	args := make([]interface{}, 0, 3+0)

	args = append(args, uniform)
	args = append(args, x)
	args = append(args, y)

	n.p.Call("setFloat2", args...)
}

// SetFloat3 calls the SetFloat3 method on the NullEngine object.
//
// https://doc.babylonjs.com/api/classes/babylon.nullengine#setfloat3
func (n *NullEngine) SetFloat3(uniform js.Value, x float64, y float64, z float64) {

	args := make([]interface{}, 0, 4+0)

	args = append(args, uniform)
	args = append(args, x)
	args = append(args, y)
	args = append(args, z)

	n.p.Call("setFloat3", args...)
}

// SetFloat4 calls the SetFloat4 method on the NullEngine object.
//
// https://doc.babylonjs.com/api/classes/babylon.nullengine#setfloat4
func (n *NullEngine) SetFloat4(uniform js.Value, x float64, y float64, z float64, w float64) {

	args := make([]interface{}, 0, 5+0)

	args = append(args, uniform)
	args = append(args, x)
	args = append(args, y)
	args = append(args, z)
	args = append(args, w)

	n.p.Call("setFloat4", args...)
}

// SetFloatArray calls the SetFloatArray method on the NullEngine object.
//
// https://doc.babylonjs.com/api/classes/babylon.nullengine#setfloatarray
func (n *NullEngine) SetFloatArray(uniform js.Value, array js.Value) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, uniform)
	args = append(args, array)

	n.p.Call("setFloatArray", args...)
}

// SetFloatArray2 calls the SetFloatArray2 method on the NullEngine object.
//
// https://doc.babylonjs.com/api/classes/babylon.nullengine#setfloatarray2
func (n *NullEngine) SetFloatArray2(uniform js.Value, array js.Value) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, uniform)
	args = append(args, array)

	n.p.Call("setFloatArray2", args...)
}

// SetFloatArray3 calls the SetFloatArray3 method on the NullEngine object.
//
// https://doc.babylonjs.com/api/classes/babylon.nullengine#setfloatarray3
func (n *NullEngine) SetFloatArray3(uniform js.Value, array js.Value) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, uniform)
	args = append(args, array)

	n.p.Call("setFloatArray3", args...)
}

// SetFloatArray4 calls the SetFloatArray4 method on the NullEngine object.
//
// https://doc.babylonjs.com/api/classes/babylon.nullengine#setfloatarray4
func (n *NullEngine) SetFloatArray4(uniform js.Value, array js.Value) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, uniform)
	args = append(args, array)

	n.p.Call("setFloatArray4", args...)
}

// SetIntArray calls the SetIntArray method on the NullEngine object.
//
// https://doc.babylonjs.com/api/classes/babylon.nullengine#setintarray
func (n *NullEngine) SetIntArray(uniform js.Value, array js.Value) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, uniform)
	args = append(args, array)

	n.p.Call("setIntArray", args...)
}

// SetIntArray2 calls the SetIntArray2 method on the NullEngine object.
//
// https://doc.babylonjs.com/api/classes/babylon.nullengine#setintarray2
func (n *NullEngine) SetIntArray2(uniform js.Value, array js.Value) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, uniform)
	args = append(args, array)

	n.p.Call("setIntArray2", args...)
}

// SetIntArray3 calls the SetIntArray3 method on the NullEngine object.
//
// https://doc.babylonjs.com/api/classes/babylon.nullengine#setintarray3
func (n *NullEngine) SetIntArray3(uniform js.Value, array js.Value) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, uniform)
	args = append(args, array)

	n.p.Call("setIntArray3", args...)
}

// SetIntArray4 calls the SetIntArray4 method on the NullEngine object.
//
// https://doc.babylonjs.com/api/classes/babylon.nullengine#setintarray4
func (n *NullEngine) SetIntArray4(uniform js.Value, array js.Value) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, uniform)
	args = append(args, array)

	n.p.Call("setIntArray4", args...)
}

// SetMatrices calls the SetMatrices method on the NullEngine object.
//
// https://doc.babylonjs.com/api/classes/babylon.nullengine#setmatrices
func (n *NullEngine) SetMatrices(uniform js.Value, matrices js.Value) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, uniform)
	args = append(args, matrices)

	n.p.Call("setMatrices", args...)
}

// SetMatrix2x2 calls the SetMatrix2x2 method on the NullEngine object.
//
// https://doc.babylonjs.com/api/classes/babylon.nullengine#setmatrix2x2
func (n *NullEngine) SetMatrix2x2(uniform js.Value, matrix js.Value) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, uniform)
	args = append(args, matrix)

	n.p.Call("setMatrix2x2", args...)
}

// SetMatrix3x3 calls the SetMatrix3x3 method on the NullEngine object.
//
// https://doc.babylonjs.com/api/classes/babylon.nullengine#setmatrix3x3
func (n *NullEngine) SetMatrix3x3(uniform js.Value, matrix js.Value) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, uniform)
	args = append(args, matrix)

	n.p.Call("setMatrix3x3", args...)
}

// NullEngineSetStateOpts contains optional parameters for NullEngine.SetState.
type NullEngineSetStateOpts struct {
	ZOffset     *float64
	Force       *bool
	ReverseSide *bool
}

// SetState calls the SetState method on the NullEngine object.
//
// https://doc.babylonjs.com/api/classes/babylon.nullengine#setstate
func (n *NullEngine) SetState(culling bool, opts *NullEngineSetStateOpts) {
	if opts == nil {
		opts = &NullEngineSetStateOpts{}
	}

	args := make([]interface{}, 0, 1+3)

	args = append(args, culling)

	if opts.ZOffset == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.ZOffset)
	}
	if opts.Force == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Force)
	}
	if opts.ReverseSide == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.ReverseSide)
	}

	n.p.Call("setState", args...)
}

// NullEngineSetViewportOpts contains optional parameters for NullEngine.SetViewport.
type NullEngineSetViewportOpts struct {
	RequiredWidth  *float64
	RequiredHeight *float64
}

// SetViewport calls the SetViewport method on the NullEngine object.
//
// https://doc.babylonjs.com/api/classes/babylon.nullengine#setviewport
func (n *NullEngine) SetViewport(viewport js.Value, opts *NullEngineSetViewportOpts) {
	if opts == nil {
		opts = &NullEngineSetViewportOpts{}
	}

	args := make([]interface{}, 0, 1+2)

	args = append(args, viewport)

	if opts.RequiredWidth == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.RequiredWidth)
	}
	if opts.RequiredHeight == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.RequiredHeight)
	}

	n.p.Call("setViewport", args...)
}

// NullEngineUnBindFramebufferOpts contains optional parameters for NullEngine.UnBindFramebuffer.
type NullEngineUnBindFramebufferOpts struct {
	DisableGenerateMipMaps *bool
	OnBeforeUnbind         JSFunc
}

// UnBindFramebuffer calls the UnBindFramebuffer method on the NullEngine object.
//
// https://doc.babylonjs.com/api/classes/babylon.nullengine#unbindframebuffer
func (n *NullEngine) UnBindFramebuffer(texture *InternalTexture, opts *NullEngineUnBindFramebufferOpts) {
	if opts == nil {
		opts = &NullEngineUnBindFramebufferOpts{}
	}

	args := make([]interface{}, 0, 1+2)

	args = append(args, texture.JSObject())

	if opts.DisableGenerateMipMaps == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.DisableGenerateMipMaps)
	}
	if opts.OnBeforeUnbind == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, js.FuncOf(opts.OnBeforeUnbind) /* never freed! */)
	}

	n.p.Call("unBindFramebuffer", args...)
}

// NullEngineUpdateDynamicIndexBufferOpts contains optional parameters for NullEngine.UpdateDynamicIndexBuffer.
type NullEngineUpdateDynamicIndexBufferOpts struct {
	Offset *float64
}

// UpdateDynamicIndexBuffer calls the UpdateDynamicIndexBuffer method on the NullEngine object.
//
// https://doc.babylonjs.com/api/classes/babylon.nullengine#updatedynamicindexbuffer
func (n *NullEngine) UpdateDynamicIndexBuffer(indexBuffer js.Value, indices js.Value, opts *NullEngineUpdateDynamicIndexBufferOpts) {
	if opts == nil {
		opts = &NullEngineUpdateDynamicIndexBufferOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, indexBuffer)
	args = append(args, indices)

	if opts.Offset == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Offset)
	}

	n.p.Call("updateDynamicIndexBuffer", args...)
}

// NullEngineUpdateDynamicTextureOpts contains optional parameters for NullEngine.UpdateDynamicTexture.
type NullEngineUpdateDynamicTextureOpts struct {
	PremulAlpha *bool
	Format      *float64
}

// UpdateDynamicTexture calls the UpdateDynamicTexture method on the NullEngine object.
//
// https://doc.babylonjs.com/api/classes/babylon.nullengine#updatedynamictexture
func (n *NullEngine) UpdateDynamicTexture(texture *InternalTexture, canvas js.Value, invertY bool, opts *NullEngineUpdateDynamicTextureOpts) {
	if opts == nil {
		opts = &NullEngineUpdateDynamicTextureOpts{}
	}

	args := make([]interface{}, 0, 3+2)

	args = append(args, texture.JSObject())
	args = append(args, canvas)
	args = append(args, invertY)

	if opts.PremulAlpha == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.PremulAlpha)
	}
	if opts.Format == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Format)
	}

	n.p.Call("updateDynamicTexture", args...)
}

// NullEngineUpdateDynamicVertexBufferOpts contains optional parameters for NullEngine.UpdateDynamicVertexBuffer.
type NullEngineUpdateDynamicVertexBufferOpts struct {
	ByteOffset *float64
	ByteLength *float64
}

// UpdateDynamicVertexBuffer calls the UpdateDynamicVertexBuffer method on the NullEngine object.
//
// https://doc.babylonjs.com/api/classes/babylon.nullengine#updatedynamicvertexbuffer
func (n *NullEngine) UpdateDynamicVertexBuffer(vertexBuffer js.Value, vertices js.Value, opts *NullEngineUpdateDynamicVertexBufferOpts) {
	if opts == nil {
		opts = &NullEngineUpdateDynamicVertexBufferOpts{}
	}

	args := make([]interface{}, 0, 2+2)

	args = append(args, vertexBuffer)
	args = append(args, vertices)

	if opts.ByteOffset == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.ByteOffset)
	}
	if opts.ByteLength == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.ByteLength)
	}

	n.p.Call("updateDynamicVertexBuffer", args...)
}

// UpdateTextureSamplingMode calls the UpdateTextureSamplingMode method on the NullEngine object.
//
// https://doc.babylonjs.com/api/classes/babylon.nullengine#updatetexturesamplingmode
func (n *NullEngine) UpdateTextureSamplingMode(samplingMode float64, texture *InternalTexture) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, samplingMode)
	args = append(args, texture.JSObject())

	n.p.Call("updateTextureSamplingMode", args...)
}

// NullEngineWipeCachesOpts contains optional parameters for NullEngine.WipeCaches.
type NullEngineWipeCachesOpts struct {
	BruteForce *bool
}

// WipeCaches calls the WipeCaches method on the NullEngine object.
//
// https://doc.babylonjs.com/api/classes/babylon.nullengine#wipecaches
func (n *NullEngine) WipeCaches(opts *NullEngineWipeCachesOpts) {
	if opts == nil {
		opts = &NullEngineWipeCachesOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.BruteForce == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.BruteForce)
	}

	n.p.Call("wipeCaches", args...)
}
