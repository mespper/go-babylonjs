// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// TrigonometryBlock represents a babylon.js TrigonometryBlock.
// Block used to apply trigonometry operation to floats
type TrigonometryBlock struct {
	*NodeMaterialBlock
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (t *TrigonometryBlock) JSObject() js.Value { return t.p }

// TrigonometryBlock returns a TrigonometryBlock JavaScript class.
func (ba *Babylon) TrigonometryBlock() *TrigonometryBlock {
	p := ba.ctx.Get("TrigonometryBlock")
	return TrigonometryBlockFromJSObject(p, ba.ctx)
}

// TrigonometryBlockFromJSObject returns a wrapped TrigonometryBlock JavaScript class.
func TrigonometryBlockFromJSObject(p js.Value, ctx js.Value) *TrigonometryBlock {
	return &TrigonometryBlock{NodeMaterialBlock: NodeMaterialBlockFromJSObject(p, ctx), ctx: ctx}
}

// TrigonometryBlockArrayToJSArray returns a JavaScript Array for the wrapped array.
func TrigonometryBlockArrayToJSArray(array []*TrigonometryBlock) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewTrigonometryBlock returns a new TrigonometryBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.trigonometryblock
func (ba *Babylon) NewTrigonometryBlock(name string) *TrigonometryBlock {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	p := ba.ctx.Get("TrigonometryBlock").New(args...)
	return TrigonometryBlockFromJSObject(p, ba.ctx)
}

// AutoConfigure calls the AutoConfigure method on the TrigonometryBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.trigonometryblock#autoconfigure
func (t *TrigonometryBlock) AutoConfigure(material *NodeMaterial) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, material.JSObject())

	t.p.Call("autoConfigure", args...)
}

// TrigonometryBlockBindOpts contains optional parameters for TrigonometryBlock.Bind.
type TrigonometryBlockBindOpts struct {
	Mesh *Mesh
}

// Bind calls the Bind method on the TrigonometryBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.trigonometryblock#bind
func (t *TrigonometryBlock) Bind(effect *Effect, nodeMaterial *NodeMaterial, opts *TrigonometryBlockBindOpts) {
	if opts == nil {
		opts = &TrigonometryBlockBindOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, effect.JSObject())
	args = append(args, nodeMaterial.JSObject())

	if opts.Mesh == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Mesh.JSObject())
	}

	t.p.Call("bind", args...)
}

// Build calls the Build method on the TrigonometryBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.trigonometryblock#build
func (t *TrigonometryBlock) Build(state *NodeMaterialBuildState, activeBlocks *NodeMaterialBlock) bool {

	args := make([]interface{}, 0, 2+0)

	args = append(args, state.JSObject())
	args = append(args, activeBlocks.JSObject())

	retVal := t.p.Call("build", args...)
	return retVal.Bool()
}

// TrigonometryBlockCloneOpts contains optional parameters for TrigonometryBlock.Clone.
type TrigonometryBlockCloneOpts struct {
	RootUrl *string
}

// Clone calls the Clone method on the TrigonometryBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.trigonometryblock#clone
func (t *TrigonometryBlock) Clone(scene *Scene, opts *TrigonometryBlockCloneOpts) *NodeMaterialBlock {
	if opts == nil {
		opts = &TrigonometryBlockCloneOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, scene.JSObject())

	if opts.RootUrl == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.RootUrl)
	}

	retVal := t.p.Call("clone", args...)
	return NodeMaterialBlockFromJSObject(retVal, t.ctx)
}

// TrigonometryBlockConnectToOpts contains optional parameters for TrigonometryBlock.ConnectTo.
type TrigonometryBlockConnectToOpts struct {
	Options map[string]interface{}
}

// ConnectTo calls the ConnectTo method on the TrigonometryBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.trigonometryblock#connectto
func (t *TrigonometryBlock) ConnectTo(other *NodeMaterialBlock, opts *TrigonometryBlockConnectToOpts) *TrigonometryBlock {
	if opts == nil {
		opts = &TrigonometryBlockConnectToOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, other.JSObject())

	if opts.Options == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Options)
	}

	retVal := t.p.Call("connectTo", args...)
	return TrigonometryBlockFromJSObject(retVal, t.ctx)
}

// Dispose calls the Dispose method on the TrigonometryBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.trigonometryblock#dispose
func (t *TrigonometryBlock) Dispose() {

	t.p.Call("dispose")
}

// GetClassName calls the GetClassName method on the TrigonometryBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.trigonometryblock#getclassname
func (t *TrigonometryBlock) GetClassName() string {

	retVal := t.p.Call("getClassName")
	return retVal.String()
}

// TrigonometryBlockGetFirstAvailableInputOpts contains optional parameters for TrigonometryBlock.GetFirstAvailableInput.
type TrigonometryBlockGetFirstAvailableInputOpts struct {
	ForOutput *NodeMaterialConnectionPoint
}

// GetFirstAvailableInput calls the GetFirstAvailableInput method on the TrigonometryBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.trigonometryblock#getfirstavailableinput
func (t *TrigonometryBlock) GetFirstAvailableInput(opts *TrigonometryBlockGetFirstAvailableInputOpts) *NodeMaterialConnectionPoint {
	if opts == nil {
		opts = &TrigonometryBlockGetFirstAvailableInputOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.ForOutput == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.ForOutput.JSObject())
	}

	retVal := t.p.Call("getFirstAvailableInput", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, t.ctx)
}

// TrigonometryBlockGetFirstAvailableOutputOpts contains optional parameters for TrigonometryBlock.GetFirstAvailableOutput.
type TrigonometryBlockGetFirstAvailableOutputOpts struct {
	ForBlock *NodeMaterialBlock
}

// GetFirstAvailableOutput calls the GetFirstAvailableOutput method on the TrigonometryBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.trigonometryblock#getfirstavailableoutput
func (t *TrigonometryBlock) GetFirstAvailableOutput(opts *TrigonometryBlockGetFirstAvailableOutputOpts) *NodeMaterialConnectionPoint {
	if opts == nil {
		opts = &TrigonometryBlockGetFirstAvailableOutputOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.ForBlock == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.ForBlock.JSObject())
	}

	retVal := t.p.Call("getFirstAvailableOutput", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, t.ctx)
}

// GetInputByName calls the GetInputByName method on the TrigonometryBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.trigonometryblock#getinputbyname
func (t *TrigonometryBlock) GetInputByName(name string) *NodeMaterialConnectionPoint {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	retVal := t.p.Call("getInputByName", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, t.ctx)
}

// GetOutputByName calls the GetOutputByName method on the TrigonometryBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.trigonometryblock#getoutputbyname
func (t *TrigonometryBlock) GetOutputByName(name string) *NodeMaterialConnectionPoint {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	retVal := t.p.Call("getOutputByName", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, t.ctx)
}

// GetSiblingOutput calls the GetSiblingOutput method on the TrigonometryBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.trigonometryblock#getsiblingoutput
func (t *TrigonometryBlock) GetSiblingOutput(current *NodeMaterialConnectionPoint) *NodeMaterialConnectionPoint {

	args := make([]interface{}, 0, 1+0)

	args = append(args, current.JSObject())

	retVal := t.p.Call("getSiblingOutput", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, t.ctx)
}

// Initialize calls the Initialize method on the TrigonometryBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.trigonometryblock#initialize
func (t *TrigonometryBlock) Initialize(state *NodeMaterialBuildState) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, state.JSObject())

	t.p.Call("initialize", args...)
}

// TrigonometryBlockInitializeDefinesOpts contains optional parameters for TrigonometryBlock.InitializeDefines.
type TrigonometryBlockInitializeDefinesOpts struct {
	UseInstances *bool
}

// InitializeDefines calls the InitializeDefines method on the TrigonometryBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.trigonometryblock#initializedefines
func (t *TrigonometryBlock) InitializeDefines(mesh *AbstractMesh, nodeMaterial *NodeMaterial, defines js.Value, opts *TrigonometryBlockInitializeDefinesOpts) {
	if opts == nil {
		opts = &TrigonometryBlockInitializeDefinesOpts{}
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

	t.p.Call("initializeDefines", args...)
}

// TrigonometryBlockIsReadyOpts contains optional parameters for TrigonometryBlock.IsReady.
type TrigonometryBlockIsReadyOpts struct {
	UseInstances *bool
}

// IsReady calls the IsReady method on the TrigonometryBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.trigonometryblock#isready
func (t *TrigonometryBlock) IsReady(mesh *AbstractMesh, nodeMaterial *NodeMaterial, defines js.Value, opts *TrigonometryBlockIsReadyOpts) bool {
	if opts == nil {
		opts = &TrigonometryBlockIsReadyOpts{}
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

	retVal := t.p.Call("isReady", args...)
	return retVal.Bool()
}

// TrigonometryBlockPrepareDefinesOpts contains optional parameters for TrigonometryBlock.PrepareDefines.
type TrigonometryBlockPrepareDefinesOpts struct {
	UseInstances *bool
}

// PrepareDefines calls the PrepareDefines method on the TrigonometryBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.trigonometryblock#preparedefines
func (t *TrigonometryBlock) PrepareDefines(mesh *AbstractMesh, nodeMaterial *NodeMaterial, defines js.Value, opts *TrigonometryBlockPrepareDefinesOpts) {
	if opts == nil {
		opts = &TrigonometryBlockPrepareDefinesOpts{}
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

	t.p.Call("prepareDefines", args...)
}

// ProvideFallbacks calls the ProvideFallbacks method on the TrigonometryBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.trigonometryblock#providefallbacks
func (t *TrigonometryBlock) ProvideFallbacks(mesh *AbstractMesh, fallbacks *EffectFallbacks) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, mesh.JSObject())
	args = append(args, fallbacks.JSObject())

	t.p.Call("provideFallbacks", args...)
}

// TrigonometryBlockRegisterInputOpts contains optional parameters for TrigonometryBlock.RegisterInput.
type TrigonometryBlockRegisterInputOpts struct {
	IsOptional *bool
	Target     js.Value
}

// RegisterInput calls the RegisterInput method on the TrigonometryBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.trigonometryblock#registerinput
func (t *TrigonometryBlock) RegisterInput(name string, jsType js.Value, opts *TrigonometryBlockRegisterInputOpts) *TrigonometryBlock {
	if opts == nil {
		opts = &TrigonometryBlockRegisterInputOpts{}
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

	retVal := t.p.Call("registerInput", args...)
	return TrigonometryBlockFromJSObject(retVal, t.ctx)
}

// TrigonometryBlockRegisterOutputOpts contains optional parameters for TrigonometryBlock.RegisterOutput.
type TrigonometryBlockRegisterOutputOpts struct {
	Target js.Value
}

// RegisterOutput calls the RegisterOutput method on the TrigonometryBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.trigonometryblock#registeroutput
func (t *TrigonometryBlock) RegisterOutput(name string, jsType js.Value, opts *TrigonometryBlockRegisterOutputOpts) *TrigonometryBlock {
	if opts == nil {
		opts = &TrigonometryBlockRegisterOutputOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, name)
	args = append(args, jsType)

	args = append(args, opts.Target)

	retVal := t.p.Call("registerOutput", args...)
	return TrigonometryBlockFromJSObject(retVal, t.ctx)
}

// ReplaceRepeatableContent calls the ReplaceRepeatableContent method on the TrigonometryBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.trigonometryblock#replacerepeatablecontent
func (t *TrigonometryBlock) ReplaceRepeatableContent(vertexShaderState *NodeMaterialBuildState, fragmentShaderState *NodeMaterialBuildState, mesh *AbstractMesh, defines js.Value) {

	args := make([]interface{}, 0, 4+0)

	args = append(args, vertexShaderState.JSObject())
	args = append(args, fragmentShaderState.JSObject())
	args = append(args, mesh.JSObject())
	args = append(args, defines)

	t.p.Call("replaceRepeatableContent", args...)
}

// Serialize calls the Serialize method on the TrigonometryBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.trigonometryblock#serialize
func (t *TrigonometryBlock) Serialize() interface{} {

	retVal := t.p.Call("serialize")
	return retVal
}

// UpdateUniformsAndSamples calls the UpdateUniformsAndSamples method on the TrigonometryBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.trigonometryblock#updateuniformsandsamples
func (t *TrigonometryBlock) UpdateUniformsAndSamples(state *NodeMaterialBuildState, nodeMaterial *NodeMaterial, defines js.Value, uniformBuffers string) {

	args := make([]interface{}, 0, 4+0)

	args = append(args, state.JSObject())
	args = append(args, nodeMaterial.JSObject())
	args = append(args, defines)
	args = append(args, uniformBuffers)

	t.p.Call("updateUniformsAndSamples", args...)
}

// _deserialize calls the _deserialize method on the TrigonometryBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.trigonometryblock#_deserialize
func (t *TrigonometryBlock) _deserialize(serializationObject interface{}, scene *Scene, rootUrl string) {

	args := make([]interface{}, 0, 3+0)

	args = append(args, serializationObject)
	args = append(args, scene.JSObject())
	args = append(args, rootUrl)

	t.p.Call("_deserialize", args...)
}

/*

// BuildId returns the BuildId property of class TrigonometryBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.trigonometryblock#buildid
func (t *TrigonometryBlock) BuildId(buildId float64) *TrigonometryBlock {
	p := ba.ctx.Get("TrigonometryBlock").New(buildId)
	return TrigonometryBlockFromJSObject(p, ba.ctx)
}

// SetBuildId sets the BuildId property of class TrigonometryBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.trigonometryblock#buildid
func (t *TrigonometryBlock) SetBuildId(buildId float64) *TrigonometryBlock {
	p := ba.ctx.Get("TrigonometryBlock").New(buildId)
	return TrigonometryBlockFromJSObject(p, ba.ctx)
}

// Comments returns the Comments property of class TrigonometryBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.trigonometryblock#comments
func (t *TrigonometryBlock) Comments(comments string) *TrigonometryBlock {
	p := ba.ctx.Get("TrigonometryBlock").New(comments)
	return TrigonometryBlockFromJSObject(p, ba.ctx)
}

// SetComments sets the Comments property of class TrigonometryBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.trigonometryblock#comments
func (t *TrigonometryBlock) SetComments(comments string) *TrigonometryBlock {
	p := ba.ctx.Get("TrigonometryBlock").New(comments)
	return TrigonometryBlockFromJSObject(p, ba.ctx)
}

// Input returns the Input property of class TrigonometryBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.trigonometryblock#input
func (t *TrigonometryBlock) Input(input *NodeMaterialConnectionPoint) *TrigonometryBlock {
	p := ba.ctx.Get("TrigonometryBlock").New(input.JSObject())
	return TrigonometryBlockFromJSObject(p, ba.ctx)
}

// SetInput sets the Input property of class TrigonometryBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.trigonometryblock#input
func (t *TrigonometryBlock) SetInput(input *NodeMaterialConnectionPoint) *TrigonometryBlock {
	p := ba.ctx.Get("TrigonometryBlock").New(input.JSObject())
	return TrigonometryBlockFromJSObject(p, ba.ctx)
}

// Inputs returns the Inputs property of class TrigonometryBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.trigonometryblock#inputs
func (t *TrigonometryBlock) Inputs(inputs *NodeMaterialConnectionPoint) *TrigonometryBlock {
	p := ba.ctx.Get("TrigonometryBlock").New(inputs.JSObject())
	return TrigonometryBlockFromJSObject(p, ba.ctx)
}

// SetInputs sets the Inputs property of class TrigonometryBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.trigonometryblock#inputs
func (t *TrigonometryBlock) SetInputs(inputs *NodeMaterialConnectionPoint) *TrigonometryBlock {
	p := ba.ctx.Get("TrigonometryBlock").New(inputs.JSObject())
	return TrigonometryBlockFromJSObject(p, ba.ctx)
}

// IsFinalMerger returns the IsFinalMerger property of class TrigonometryBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.trigonometryblock#isfinalmerger
func (t *TrigonometryBlock) IsFinalMerger(isFinalMerger bool) *TrigonometryBlock {
	p := ba.ctx.Get("TrigonometryBlock").New(isFinalMerger)
	return TrigonometryBlockFromJSObject(p, ba.ctx)
}

// SetIsFinalMerger sets the IsFinalMerger property of class TrigonometryBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.trigonometryblock#isfinalmerger
func (t *TrigonometryBlock) SetIsFinalMerger(isFinalMerger bool) *TrigonometryBlock {
	p := ba.ctx.Get("TrigonometryBlock").New(isFinalMerger)
	return TrigonometryBlockFromJSObject(p, ba.ctx)
}

// IsInput returns the IsInput property of class TrigonometryBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.trigonometryblock#isinput
func (t *TrigonometryBlock) IsInput(isInput bool) *TrigonometryBlock {
	p := ba.ctx.Get("TrigonometryBlock").New(isInput)
	return TrigonometryBlockFromJSObject(p, ba.ctx)
}

// SetIsInput sets the IsInput property of class TrigonometryBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.trigonometryblock#isinput
func (t *TrigonometryBlock) SetIsInput(isInput bool) *TrigonometryBlock {
	p := ba.ctx.Get("TrigonometryBlock").New(isInput)
	return TrigonometryBlockFromJSObject(p, ba.ctx)
}

// IsUnique returns the IsUnique property of class TrigonometryBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.trigonometryblock#isunique
func (t *TrigonometryBlock) IsUnique(isUnique bool) *TrigonometryBlock {
	p := ba.ctx.Get("TrigonometryBlock").New(isUnique)
	return TrigonometryBlockFromJSObject(p, ba.ctx)
}

// SetIsUnique sets the IsUnique property of class TrigonometryBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.trigonometryblock#isunique
func (t *TrigonometryBlock) SetIsUnique(isUnique bool) *TrigonometryBlock {
	p := ba.ctx.Get("TrigonometryBlock").New(isUnique)
	return TrigonometryBlockFromJSObject(p, ba.ctx)
}

// Name returns the Name property of class TrigonometryBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.trigonometryblock#name
func (t *TrigonometryBlock) Name(name string) *TrigonometryBlock {
	p := ba.ctx.Get("TrigonometryBlock").New(name)
	return TrigonometryBlockFromJSObject(p, ba.ctx)
}

// SetName sets the Name property of class TrigonometryBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.trigonometryblock#name
func (t *TrigonometryBlock) SetName(name string) *TrigonometryBlock {
	p := ba.ctx.Get("TrigonometryBlock").New(name)
	return TrigonometryBlockFromJSObject(p, ba.ctx)
}

// Operation returns the Operation property of class TrigonometryBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.trigonometryblock#operation
func (t *TrigonometryBlock) Operation(operation *TrigonometryBlockOperations) *TrigonometryBlock {
	p := ba.ctx.Get("TrigonometryBlock").New(operation.JSObject())
	return TrigonometryBlockFromJSObject(p, ba.ctx)
}

// SetOperation sets the Operation property of class TrigonometryBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.trigonometryblock#operation
func (t *TrigonometryBlock) SetOperation(operation *TrigonometryBlockOperations) *TrigonometryBlock {
	p := ba.ctx.Get("TrigonometryBlock").New(operation.JSObject())
	return TrigonometryBlockFromJSObject(p, ba.ctx)
}

// Output returns the Output property of class TrigonometryBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.trigonometryblock#output
func (t *TrigonometryBlock) Output(output *NodeMaterialConnectionPoint) *TrigonometryBlock {
	p := ba.ctx.Get("TrigonometryBlock").New(output.JSObject())
	return TrigonometryBlockFromJSObject(p, ba.ctx)
}

// SetOutput sets the Output property of class TrigonometryBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.trigonometryblock#output
func (t *TrigonometryBlock) SetOutput(output *NodeMaterialConnectionPoint) *TrigonometryBlock {
	p := ba.ctx.Get("TrigonometryBlock").New(output.JSObject())
	return TrigonometryBlockFromJSObject(p, ba.ctx)
}

// Outputs returns the Outputs property of class TrigonometryBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.trigonometryblock#outputs
func (t *TrigonometryBlock) Outputs(outputs *NodeMaterialConnectionPoint) *TrigonometryBlock {
	p := ba.ctx.Get("TrigonometryBlock").New(outputs.JSObject())
	return TrigonometryBlockFromJSObject(p, ba.ctx)
}

// SetOutputs sets the Outputs property of class TrigonometryBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.trigonometryblock#outputs
func (t *TrigonometryBlock) SetOutputs(outputs *NodeMaterialConnectionPoint) *TrigonometryBlock {
	p := ba.ctx.Get("TrigonometryBlock").New(outputs.JSObject())
	return TrigonometryBlockFromJSObject(p, ba.ctx)
}

// Target returns the Target property of class TrigonometryBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.trigonometryblock#target
func (t *TrigonometryBlock) Target(target js.Value) *TrigonometryBlock {
	p := ba.ctx.Get("TrigonometryBlock").New(target)
	return TrigonometryBlockFromJSObject(p, ba.ctx)
}

// SetTarget sets the Target property of class TrigonometryBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.trigonometryblock#target
func (t *TrigonometryBlock) SetTarget(target js.Value) *TrigonometryBlock {
	p := ba.ctx.Get("TrigonometryBlock").New(target)
	return TrigonometryBlockFromJSObject(p, ba.ctx)
}

// UniqueId returns the UniqueId property of class TrigonometryBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.trigonometryblock#uniqueid
func (t *TrigonometryBlock) UniqueId(uniqueId float64) *TrigonometryBlock {
	p := ba.ctx.Get("TrigonometryBlock").New(uniqueId)
	return TrigonometryBlockFromJSObject(p, ba.ctx)
}

// SetUniqueId sets the UniqueId property of class TrigonometryBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.trigonometryblock#uniqueid
func (t *TrigonometryBlock) SetUniqueId(uniqueId float64) *TrigonometryBlock {
	p := ba.ctx.Get("TrigonometryBlock").New(uniqueId)
	return TrigonometryBlockFromJSObject(p, ba.ctx)
}

*/
