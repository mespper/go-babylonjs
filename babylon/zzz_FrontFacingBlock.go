// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// FrontFacingBlock represents a babylon.js FrontFacingBlock.
// Block used to test if the fragment shader is front facing
type FrontFacingBlock struct {
	*NodeMaterialBlock
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (f *FrontFacingBlock) JSObject() js.Value { return f.p }

// FrontFacingBlock returns a FrontFacingBlock JavaScript class.
func (ba *Babylon) FrontFacingBlock() *FrontFacingBlock {
	p := ba.ctx.Get("FrontFacingBlock")
	return FrontFacingBlockFromJSObject(p, ba.ctx)
}

// FrontFacingBlockFromJSObject returns a wrapped FrontFacingBlock JavaScript class.
func FrontFacingBlockFromJSObject(p js.Value, ctx js.Value) *FrontFacingBlock {
	return &FrontFacingBlock{NodeMaterialBlock: NodeMaterialBlockFromJSObject(p, ctx), ctx: ctx}
}

// FrontFacingBlockArrayToJSArray returns a JavaScript Array for the wrapped array.
func FrontFacingBlockArrayToJSArray(array []*FrontFacingBlock) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewFrontFacingBlock returns a new FrontFacingBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.frontfacingblock
func (ba *Babylon) NewFrontFacingBlock(name string) *FrontFacingBlock {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	p := ba.ctx.Get("FrontFacingBlock").New(args...)
	return FrontFacingBlockFromJSObject(p, ba.ctx)
}

// AutoConfigure calls the AutoConfigure method on the FrontFacingBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.frontfacingblock#autoconfigure
func (f *FrontFacingBlock) AutoConfigure(material *NodeMaterial) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, material.JSObject())

	f.p.Call("autoConfigure", args...)
}

// FrontFacingBlockBindOpts contains optional parameters for FrontFacingBlock.Bind.
type FrontFacingBlockBindOpts struct {
	Mesh *Mesh
}

// Bind calls the Bind method on the FrontFacingBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.frontfacingblock#bind
func (f *FrontFacingBlock) Bind(effect *Effect, nodeMaterial *NodeMaterial, opts *FrontFacingBlockBindOpts) {
	if opts == nil {
		opts = &FrontFacingBlockBindOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, effect.JSObject())
	args = append(args, nodeMaterial.JSObject())

	if opts.Mesh == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Mesh.JSObject())
	}

	f.p.Call("bind", args...)
}

// Build calls the Build method on the FrontFacingBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.frontfacingblock#build
func (f *FrontFacingBlock) Build(state *NodeMaterialBuildState, activeBlocks *NodeMaterialBlock) bool {

	args := make([]interface{}, 0, 2+0)

	args = append(args, state.JSObject())
	args = append(args, activeBlocks.JSObject())

	retVal := f.p.Call("build", args...)
	return retVal.Bool()
}

// FrontFacingBlockCloneOpts contains optional parameters for FrontFacingBlock.Clone.
type FrontFacingBlockCloneOpts struct {
	RootUrl *string
}

// Clone calls the Clone method on the FrontFacingBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.frontfacingblock#clone
func (f *FrontFacingBlock) Clone(scene *Scene, opts *FrontFacingBlockCloneOpts) *NodeMaterialBlock {
	if opts == nil {
		opts = &FrontFacingBlockCloneOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, scene.JSObject())

	if opts.RootUrl == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.RootUrl)
	}

	retVal := f.p.Call("clone", args...)
	return NodeMaterialBlockFromJSObject(retVal, f.ctx)
}

// FrontFacingBlockConnectToOpts contains optional parameters for FrontFacingBlock.ConnectTo.
type FrontFacingBlockConnectToOpts struct {
	Options map[string]interface{}
}

// ConnectTo calls the ConnectTo method on the FrontFacingBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.frontfacingblock#connectto
func (f *FrontFacingBlock) ConnectTo(other *NodeMaterialBlock, opts *FrontFacingBlockConnectToOpts) *FrontFacingBlock {
	if opts == nil {
		opts = &FrontFacingBlockConnectToOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, other.JSObject())

	if opts.Options == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Options)
	}

	retVal := f.p.Call("connectTo", args...)
	return FrontFacingBlockFromJSObject(retVal, f.ctx)
}

// Dispose calls the Dispose method on the FrontFacingBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.frontfacingblock#dispose
func (f *FrontFacingBlock) Dispose() {

	f.p.Call("dispose")
}

// GetClassName calls the GetClassName method on the FrontFacingBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.frontfacingblock#getclassname
func (f *FrontFacingBlock) GetClassName() string {

	retVal := f.p.Call("getClassName")
	return retVal.String()
}

// FrontFacingBlockGetFirstAvailableInputOpts contains optional parameters for FrontFacingBlock.GetFirstAvailableInput.
type FrontFacingBlockGetFirstAvailableInputOpts struct {
	ForOutput *NodeMaterialConnectionPoint
}

// GetFirstAvailableInput calls the GetFirstAvailableInput method on the FrontFacingBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.frontfacingblock#getfirstavailableinput
func (f *FrontFacingBlock) GetFirstAvailableInput(opts *FrontFacingBlockGetFirstAvailableInputOpts) *NodeMaterialConnectionPoint {
	if opts == nil {
		opts = &FrontFacingBlockGetFirstAvailableInputOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.ForOutput == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.ForOutput.JSObject())
	}

	retVal := f.p.Call("getFirstAvailableInput", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, f.ctx)
}

// FrontFacingBlockGetFirstAvailableOutputOpts contains optional parameters for FrontFacingBlock.GetFirstAvailableOutput.
type FrontFacingBlockGetFirstAvailableOutputOpts struct {
	ForBlock *NodeMaterialBlock
}

// GetFirstAvailableOutput calls the GetFirstAvailableOutput method on the FrontFacingBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.frontfacingblock#getfirstavailableoutput
func (f *FrontFacingBlock) GetFirstAvailableOutput(opts *FrontFacingBlockGetFirstAvailableOutputOpts) *NodeMaterialConnectionPoint {
	if opts == nil {
		opts = &FrontFacingBlockGetFirstAvailableOutputOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.ForBlock == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.ForBlock.JSObject())
	}

	retVal := f.p.Call("getFirstAvailableOutput", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, f.ctx)
}

// GetInputByName calls the GetInputByName method on the FrontFacingBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.frontfacingblock#getinputbyname
func (f *FrontFacingBlock) GetInputByName(name string) *NodeMaterialConnectionPoint {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	retVal := f.p.Call("getInputByName", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, f.ctx)
}

// GetOutputByName calls the GetOutputByName method on the FrontFacingBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.frontfacingblock#getoutputbyname
func (f *FrontFacingBlock) GetOutputByName(name string) *NodeMaterialConnectionPoint {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	retVal := f.p.Call("getOutputByName", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, f.ctx)
}

// GetSiblingOutput calls the GetSiblingOutput method on the FrontFacingBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.frontfacingblock#getsiblingoutput
func (f *FrontFacingBlock) GetSiblingOutput(current *NodeMaterialConnectionPoint) *NodeMaterialConnectionPoint {

	args := make([]interface{}, 0, 1+0)

	args = append(args, current.JSObject())

	retVal := f.p.Call("getSiblingOutput", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, f.ctx)
}

// Initialize calls the Initialize method on the FrontFacingBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.frontfacingblock#initialize
func (f *FrontFacingBlock) Initialize(state *NodeMaterialBuildState) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, state.JSObject())

	f.p.Call("initialize", args...)
}

// FrontFacingBlockInitializeDefinesOpts contains optional parameters for FrontFacingBlock.InitializeDefines.
type FrontFacingBlockInitializeDefinesOpts struct {
	UseInstances *bool
}

// InitializeDefines calls the InitializeDefines method on the FrontFacingBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.frontfacingblock#initializedefines
func (f *FrontFacingBlock) InitializeDefines(mesh *AbstractMesh, nodeMaterial *NodeMaterial, defines js.Value, opts *FrontFacingBlockInitializeDefinesOpts) {
	if opts == nil {
		opts = &FrontFacingBlockInitializeDefinesOpts{}
	}

	args := make([]interface{}, 0, 3+1)

	args = append(args, mesh.JSObject())
	args = append(args, nodeMaterial.JSObject())
	args = append(args, defines)

	if opts.UseInstances == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.UseInstances)
	}

	f.p.Call("initializeDefines", args...)
}

// FrontFacingBlockIsReadyOpts contains optional parameters for FrontFacingBlock.IsReady.
type FrontFacingBlockIsReadyOpts struct {
	UseInstances *bool
}

// IsReady calls the IsReady method on the FrontFacingBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.frontfacingblock#isready
func (f *FrontFacingBlock) IsReady(mesh *AbstractMesh, nodeMaterial *NodeMaterial, defines js.Value, opts *FrontFacingBlockIsReadyOpts) bool {
	if opts == nil {
		opts = &FrontFacingBlockIsReadyOpts{}
	}

	args := make([]interface{}, 0, 3+1)

	args = append(args, mesh.JSObject())
	args = append(args, nodeMaterial.JSObject())
	args = append(args, defines)

	if opts.UseInstances == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.UseInstances)
	}

	retVal := f.p.Call("isReady", args...)
	return retVal.Bool()
}

// FrontFacingBlockPrepareDefinesOpts contains optional parameters for FrontFacingBlock.PrepareDefines.
type FrontFacingBlockPrepareDefinesOpts struct {
	UseInstances *bool
}

// PrepareDefines calls the PrepareDefines method on the FrontFacingBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.frontfacingblock#preparedefines
func (f *FrontFacingBlock) PrepareDefines(mesh *AbstractMesh, nodeMaterial *NodeMaterial, defines js.Value, opts *FrontFacingBlockPrepareDefinesOpts) {
	if opts == nil {
		opts = &FrontFacingBlockPrepareDefinesOpts{}
	}

	args := make([]interface{}, 0, 3+1)

	args = append(args, mesh.JSObject())
	args = append(args, nodeMaterial.JSObject())
	args = append(args, defines)

	if opts.UseInstances == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.UseInstances)
	}

	f.p.Call("prepareDefines", args...)
}

// ProvideFallbacks calls the ProvideFallbacks method on the FrontFacingBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.frontfacingblock#providefallbacks
func (f *FrontFacingBlock) ProvideFallbacks(mesh *AbstractMesh, fallbacks *EffectFallbacks) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, mesh.JSObject())
	args = append(args, fallbacks.JSObject())

	f.p.Call("provideFallbacks", args...)
}

// FrontFacingBlockRegisterInputOpts contains optional parameters for FrontFacingBlock.RegisterInput.
type FrontFacingBlockRegisterInputOpts struct {
	IsOptional *bool
	Target     js.Value
}

// RegisterInput calls the RegisterInput method on the FrontFacingBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.frontfacingblock#registerinput
func (f *FrontFacingBlock) RegisterInput(name string, jsType js.Value, opts *FrontFacingBlockRegisterInputOpts) *FrontFacingBlock {
	if opts == nil {
		opts = &FrontFacingBlockRegisterInputOpts{}
	}

	args := make([]interface{}, 0, 2+2)

	args = append(args, name)
	args = append(args, jsType)

	if opts.IsOptional == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.IsOptional)
	}
	args = append(args, opts.Target)

	retVal := f.p.Call("registerInput", args...)
	return FrontFacingBlockFromJSObject(retVal, f.ctx)
}

// FrontFacingBlockRegisterOutputOpts contains optional parameters for FrontFacingBlock.RegisterOutput.
type FrontFacingBlockRegisterOutputOpts struct {
	Target js.Value
}

// RegisterOutput calls the RegisterOutput method on the FrontFacingBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.frontfacingblock#registeroutput
func (f *FrontFacingBlock) RegisterOutput(name string, jsType js.Value, opts *FrontFacingBlockRegisterOutputOpts) *FrontFacingBlock {
	if opts == nil {
		opts = &FrontFacingBlockRegisterOutputOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, name)
	args = append(args, jsType)

	args = append(args, opts.Target)

	retVal := f.p.Call("registerOutput", args...)
	return FrontFacingBlockFromJSObject(retVal, f.ctx)
}

// ReplaceRepeatableContent calls the ReplaceRepeatableContent method on the FrontFacingBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.frontfacingblock#replacerepeatablecontent
func (f *FrontFacingBlock) ReplaceRepeatableContent(vertexShaderState *NodeMaterialBuildState, fragmentShaderState *NodeMaterialBuildState, mesh *AbstractMesh, defines js.Value) {

	args := make([]interface{}, 0, 4+0)

	args = append(args, vertexShaderState.JSObject())
	args = append(args, fragmentShaderState.JSObject())
	args = append(args, mesh.JSObject())
	args = append(args, defines)

	f.p.Call("replaceRepeatableContent", args...)
}

// Serialize calls the Serialize method on the FrontFacingBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.frontfacingblock#serialize
func (f *FrontFacingBlock) Serialize() interface{} {

	retVal := f.p.Call("serialize")
	return retVal
}

// UpdateUniformsAndSamples calls the UpdateUniformsAndSamples method on the FrontFacingBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.frontfacingblock#updateuniformsandsamples
func (f *FrontFacingBlock) UpdateUniformsAndSamples(state *NodeMaterialBuildState, nodeMaterial *NodeMaterial, defines js.Value, uniformBuffers string) {

	args := make([]interface{}, 0, 4+0)

	args = append(args, state.JSObject())
	args = append(args, nodeMaterial.JSObject())
	args = append(args, defines)
	args = append(args, uniformBuffers)

	f.p.Call("updateUniformsAndSamples", args...)
}

/*

// BuildId returns the BuildId property of class FrontFacingBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.frontfacingblock#buildid
func (f *FrontFacingBlock) BuildId(buildId float64) *FrontFacingBlock {
	p := ba.ctx.Get("FrontFacingBlock").New(buildId)
	return FrontFacingBlockFromJSObject(p, ba.ctx)
}

// SetBuildId sets the BuildId property of class FrontFacingBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.frontfacingblock#buildid
func (f *FrontFacingBlock) SetBuildId(buildId float64) *FrontFacingBlock {
	p := ba.ctx.Get("FrontFacingBlock").New(buildId)
	return FrontFacingBlockFromJSObject(p, ba.ctx)
}

// Comments returns the Comments property of class FrontFacingBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.frontfacingblock#comments
func (f *FrontFacingBlock) Comments(comments string) *FrontFacingBlock {
	p := ba.ctx.Get("FrontFacingBlock").New(comments)
	return FrontFacingBlockFromJSObject(p, ba.ctx)
}

// SetComments sets the Comments property of class FrontFacingBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.frontfacingblock#comments
func (f *FrontFacingBlock) SetComments(comments string) *FrontFacingBlock {
	p := ba.ctx.Get("FrontFacingBlock").New(comments)
	return FrontFacingBlockFromJSObject(p, ba.ctx)
}

// Inputs returns the Inputs property of class FrontFacingBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.frontfacingblock#inputs
func (f *FrontFacingBlock) Inputs(inputs *NodeMaterialConnectionPoint) *FrontFacingBlock {
	p := ba.ctx.Get("FrontFacingBlock").New(inputs.JSObject())
	return FrontFacingBlockFromJSObject(p, ba.ctx)
}

// SetInputs sets the Inputs property of class FrontFacingBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.frontfacingblock#inputs
func (f *FrontFacingBlock) SetInputs(inputs *NodeMaterialConnectionPoint) *FrontFacingBlock {
	p := ba.ctx.Get("FrontFacingBlock").New(inputs.JSObject())
	return FrontFacingBlockFromJSObject(p, ba.ctx)
}

// IsFinalMerger returns the IsFinalMerger property of class FrontFacingBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.frontfacingblock#isfinalmerger
func (f *FrontFacingBlock) IsFinalMerger(isFinalMerger bool) *FrontFacingBlock {
	p := ba.ctx.Get("FrontFacingBlock").New(isFinalMerger)
	return FrontFacingBlockFromJSObject(p, ba.ctx)
}

// SetIsFinalMerger sets the IsFinalMerger property of class FrontFacingBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.frontfacingblock#isfinalmerger
func (f *FrontFacingBlock) SetIsFinalMerger(isFinalMerger bool) *FrontFacingBlock {
	p := ba.ctx.Get("FrontFacingBlock").New(isFinalMerger)
	return FrontFacingBlockFromJSObject(p, ba.ctx)
}

// IsInput returns the IsInput property of class FrontFacingBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.frontfacingblock#isinput
func (f *FrontFacingBlock) IsInput(isInput bool) *FrontFacingBlock {
	p := ba.ctx.Get("FrontFacingBlock").New(isInput)
	return FrontFacingBlockFromJSObject(p, ba.ctx)
}

// SetIsInput sets the IsInput property of class FrontFacingBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.frontfacingblock#isinput
func (f *FrontFacingBlock) SetIsInput(isInput bool) *FrontFacingBlock {
	p := ba.ctx.Get("FrontFacingBlock").New(isInput)
	return FrontFacingBlockFromJSObject(p, ba.ctx)
}

// IsUnique returns the IsUnique property of class FrontFacingBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.frontfacingblock#isunique
func (f *FrontFacingBlock) IsUnique(isUnique bool) *FrontFacingBlock {
	p := ba.ctx.Get("FrontFacingBlock").New(isUnique)
	return FrontFacingBlockFromJSObject(p, ba.ctx)
}

// SetIsUnique sets the IsUnique property of class FrontFacingBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.frontfacingblock#isunique
func (f *FrontFacingBlock) SetIsUnique(isUnique bool) *FrontFacingBlock {
	p := ba.ctx.Get("FrontFacingBlock").New(isUnique)
	return FrontFacingBlockFromJSObject(p, ba.ctx)
}

// Name returns the Name property of class FrontFacingBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.frontfacingblock#name
func (f *FrontFacingBlock) Name(name string) *FrontFacingBlock {
	p := ba.ctx.Get("FrontFacingBlock").New(name)
	return FrontFacingBlockFromJSObject(p, ba.ctx)
}

// SetName sets the Name property of class FrontFacingBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.frontfacingblock#name
func (f *FrontFacingBlock) SetName(name string) *FrontFacingBlock {
	p := ba.ctx.Get("FrontFacingBlock").New(name)
	return FrontFacingBlockFromJSObject(p, ba.ctx)
}

// Output returns the Output property of class FrontFacingBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.frontfacingblock#output
func (f *FrontFacingBlock) Output(output *NodeMaterialConnectionPoint) *FrontFacingBlock {
	p := ba.ctx.Get("FrontFacingBlock").New(output.JSObject())
	return FrontFacingBlockFromJSObject(p, ba.ctx)
}

// SetOutput sets the Output property of class FrontFacingBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.frontfacingblock#output
func (f *FrontFacingBlock) SetOutput(output *NodeMaterialConnectionPoint) *FrontFacingBlock {
	p := ba.ctx.Get("FrontFacingBlock").New(output.JSObject())
	return FrontFacingBlockFromJSObject(p, ba.ctx)
}

// Outputs returns the Outputs property of class FrontFacingBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.frontfacingblock#outputs
func (f *FrontFacingBlock) Outputs(outputs *NodeMaterialConnectionPoint) *FrontFacingBlock {
	p := ba.ctx.Get("FrontFacingBlock").New(outputs.JSObject())
	return FrontFacingBlockFromJSObject(p, ba.ctx)
}

// SetOutputs sets the Outputs property of class FrontFacingBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.frontfacingblock#outputs
func (f *FrontFacingBlock) SetOutputs(outputs *NodeMaterialConnectionPoint) *FrontFacingBlock {
	p := ba.ctx.Get("FrontFacingBlock").New(outputs.JSObject())
	return FrontFacingBlockFromJSObject(p, ba.ctx)
}

// Target returns the Target property of class FrontFacingBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.frontfacingblock#target
func (f *FrontFacingBlock) Target(target js.Value) *FrontFacingBlock {
	p := ba.ctx.Get("FrontFacingBlock").New(target)
	return FrontFacingBlockFromJSObject(p, ba.ctx)
}

// SetTarget sets the Target property of class FrontFacingBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.frontfacingblock#target
func (f *FrontFacingBlock) SetTarget(target js.Value) *FrontFacingBlock {
	p := ba.ctx.Get("FrontFacingBlock").New(target)
	return FrontFacingBlockFromJSObject(p, ba.ctx)
}

// UniqueId returns the UniqueId property of class FrontFacingBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.frontfacingblock#uniqueid
func (f *FrontFacingBlock) UniqueId(uniqueId float64) *FrontFacingBlock {
	p := ba.ctx.Get("FrontFacingBlock").New(uniqueId)
	return FrontFacingBlockFromJSObject(p, ba.ctx)
}

// SetUniqueId sets the UniqueId property of class FrontFacingBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.frontfacingblock#uniqueid
func (f *FrontFacingBlock) SetUniqueId(uniqueId float64) *FrontFacingBlock {
	p := ba.ctx.Get("FrontFacingBlock").New(uniqueId)
	return FrontFacingBlockFromJSObject(p, ba.ctx)
}

*/
