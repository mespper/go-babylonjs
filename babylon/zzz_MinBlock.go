// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// MinBlock represents a babylon.js MinBlock.
// Block used to get the min of 2 values
type MinBlock struct {
	*NodeMaterialBlock
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (m *MinBlock) JSObject() js.Value { return m.p }

// MinBlock returns a MinBlock JavaScript class.
func (ba *Babylon) MinBlock() *MinBlock {
	p := ba.ctx.Get("MinBlock")
	return MinBlockFromJSObject(p, ba.ctx)
}

// MinBlockFromJSObject returns a wrapped MinBlock JavaScript class.
func MinBlockFromJSObject(p js.Value, ctx js.Value) *MinBlock {
	return &MinBlock{NodeMaterialBlock: NodeMaterialBlockFromJSObject(p, ctx), ctx: ctx}
}

// MinBlockArrayToJSArray returns a JavaScript Array for the wrapped array.
func MinBlockArrayToJSArray(array []*MinBlock) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewMinBlock returns a new MinBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.minblock
func (ba *Babylon) NewMinBlock(name string) *MinBlock {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	p := ba.ctx.Get("MinBlock").New(args...)
	return MinBlockFromJSObject(p, ba.ctx)
}

// AutoConfigure calls the AutoConfigure method on the MinBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.minblock#autoconfigure
func (m *MinBlock) AutoConfigure(material *NodeMaterial) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, material.JSObject())

	m.p.Call("autoConfigure", args...)
}

// MinBlockBindOpts contains optional parameters for MinBlock.Bind.
type MinBlockBindOpts struct {
	Mesh *Mesh
}

// Bind calls the Bind method on the MinBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.minblock#bind
func (m *MinBlock) Bind(effect *Effect, nodeMaterial *NodeMaterial, opts *MinBlockBindOpts) {
	if opts == nil {
		opts = &MinBlockBindOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, effect.JSObject())
	args = append(args, nodeMaterial.JSObject())

	if opts.Mesh == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Mesh.JSObject())
	}

	m.p.Call("bind", args...)
}

// Build calls the Build method on the MinBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.minblock#build
func (m *MinBlock) Build(state *NodeMaterialBuildState, activeBlocks *NodeMaterialBlock) bool {

	args := make([]interface{}, 0, 2+0)

	args = append(args, state.JSObject())
	args = append(args, activeBlocks.JSObject())

	retVal := m.p.Call("build", args...)
	return retVal.Bool()
}

// MinBlockCloneOpts contains optional parameters for MinBlock.Clone.
type MinBlockCloneOpts struct {
	RootUrl *string
}

// Clone calls the Clone method on the MinBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.minblock#clone
func (m *MinBlock) Clone(scene *Scene, opts *MinBlockCloneOpts) *NodeMaterialBlock {
	if opts == nil {
		opts = &MinBlockCloneOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, scene.JSObject())

	if opts.RootUrl == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.RootUrl)
	}

	retVal := m.p.Call("clone", args...)
	return NodeMaterialBlockFromJSObject(retVal, m.ctx)
}

// MinBlockConnectToOpts contains optional parameters for MinBlock.ConnectTo.
type MinBlockConnectToOpts struct {
	Options map[string]interface{}
}

// ConnectTo calls the ConnectTo method on the MinBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.minblock#connectto
func (m *MinBlock) ConnectTo(other *NodeMaterialBlock, opts *MinBlockConnectToOpts) *MinBlock {
	if opts == nil {
		opts = &MinBlockConnectToOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, other.JSObject())

	if opts.Options == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Options)
	}

	retVal := m.p.Call("connectTo", args...)
	return MinBlockFromJSObject(retVal, m.ctx)
}

// Dispose calls the Dispose method on the MinBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.minblock#dispose
func (m *MinBlock) Dispose() {

	m.p.Call("dispose")
}

// GetClassName calls the GetClassName method on the MinBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.minblock#getclassname
func (m *MinBlock) GetClassName() string {

	retVal := m.p.Call("getClassName")
	return retVal.String()
}

// MinBlockGetFirstAvailableInputOpts contains optional parameters for MinBlock.GetFirstAvailableInput.
type MinBlockGetFirstAvailableInputOpts struct {
	ForOutput *NodeMaterialConnectionPoint
}

// GetFirstAvailableInput calls the GetFirstAvailableInput method on the MinBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.minblock#getfirstavailableinput
func (m *MinBlock) GetFirstAvailableInput(opts *MinBlockGetFirstAvailableInputOpts) *NodeMaterialConnectionPoint {
	if opts == nil {
		opts = &MinBlockGetFirstAvailableInputOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.ForOutput == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.ForOutput.JSObject())
	}

	retVal := m.p.Call("getFirstAvailableInput", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, m.ctx)
}

// MinBlockGetFirstAvailableOutputOpts contains optional parameters for MinBlock.GetFirstAvailableOutput.
type MinBlockGetFirstAvailableOutputOpts struct {
	ForBlock *NodeMaterialBlock
}

// GetFirstAvailableOutput calls the GetFirstAvailableOutput method on the MinBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.minblock#getfirstavailableoutput
func (m *MinBlock) GetFirstAvailableOutput(opts *MinBlockGetFirstAvailableOutputOpts) *NodeMaterialConnectionPoint {
	if opts == nil {
		opts = &MinBlockGetFirstAvailableOutputOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.ForBlock == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.ForBlock.JSObject())
	}

	retVal := m.p.Call("getFirstAvailableOutput", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, m.ctx)
}

// GetInputByName calls the GetInputByName method on the MinBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.minblock#getinputbyname
func (m *MinBlock) GetInputByName(name string) *NodeMaterialConnectionPoint {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	retVal := m.p.Call("getInputByName", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, m.ctx)
}

// GetOutputByName calls the GetOutputByName method on the MinBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.minblock#getoutputbyname
func (m *MinBlock) GetOutputByName(name string) *NodeMaterialConnectionPoint {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	retVal := m.p.Call("getOutputByName", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, m.ctx)
}

// GetSiblingOutput calls the GetSiblingOutput method on the MinBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.minblock#getsiblingoutput
func (m *MinBlock) GetSiblingOutput(current *NodeMaterialConnectionPoint) *NodeMaterialConnectionPoint {

	args := make([]interface{}, 0, 1+0)

	args = append(args, current.JSObject())

	retVal := m.p.Call("getSiblingOutput", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, m.ctx)
}

// Initialize calls the Initialize method on the MinBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.minblock#initialize
func (m *MinBlock) Initialize(state *NodeMaterialBuildState) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, state.JSObject())

	m.p.Call("initialize", args...)
}

// MinBlockInitializeDefinesOpts contains optional parameters for MinBlock.InitializeDefines.
type MinBlockInitializeDefinesOpts struct {
	UseInstances *bool
}

// InitializeDefines calls the InitializeDefines method on the MinBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.minblock#initializedefines
func (m *MinBlock) InitializeDefines(mesh *AbstractMesh, nodeMaterial *NodeMaterial, defines js.Value, opts *MinBlockInitializeDefinesOpts) {
	if opts == nil {
		opts = &MinBlockInitializeDefinesOpts{}
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

	m.p.Call("initializeDefines", args...)
}

// MinBlockIsReadyOpts contains optional parameters for MinBlock.IsReady.
type MinBlockIsReadyOpts struct {
	UseInstances *bool
}

// IsReady calls the IsReady method on the MinBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.minblock#isready
func (m *MinBlock) IsReady(mesh *AbstractMesh, nodeMaterial *NodeMaterial, defines js.Value, opts *MinBlockIsReadyOpts) bool {
	if opts == nil {
		opts = &MinBlockIsReadyOpts{}
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

	retVal := m.p.Call("isReady", args...)
	return retVal.Bool()
}

// MinBlockPrepareDefinesOpts contains optional parameters for MinBlock.PrepareDefines.
type MinBlockPrepareDefinesOpts struct {
	UseInstances *bool
}

// PrepareDefines calls the PrepareDefines method on the MinBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.minblock#preparedefines
func (m *MinBlock) PrepareDefines(mesh *AbstractMesh, nodeMaterial *NodeMaterial, defines js.Value, opts *MinBlockPrepareDefinesOpts) {
	if opts == nil {
		opts = &MinBlockPrepareDefinesOpts{}
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

	m.p.Call("prepareDefines", args...)
}

// ProvideFallbacks calls the ProvideFallbacks method on the MinBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.minblock#providefallbacks
func (m *MinBlock) ProvideFallbacks(mesh *AbstractMesh, fallbacks *EffectFallbacks) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, mesh.JSObject())
	args = append(args, fallbacks.JSObject())

	m.p.Call("provideFallbacks", args...)
}

// MinBlockRegisterInputOpts contains optional parameters for MinBlock.RegisterInput.
type MinBlockRegisterInputOpts struct {
	IsOptional *bool
	Target     js.Value
}

// RegisterInput calls the RegisterInput method on the MinBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.minblock#registerinput
func (m *MinBlock) RegisterInput(name string, jsType js.Value, opts *MinBlockRegisterInputOpts) *MinBlock {
	if opts == nil {
		opts = &MinBlockRegisterInputOpts{}
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

	retVal := m.p.Call("registerInput", args...)
	return MinBlockFromJSObject(retVal, m.ctx)
}

// MinBlockRegisterOutputOpts contains optional parameters for MinBlock.RegisterOutput.
type MinBlockRegisterOutputOpts struct {
	Target js.Value
}

// RegisterOutput calls the RegisterOutput method on the MinBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.minblock#registeroutput
func (m *MinBlock) RegisterOutput(name string, jsType js.Value, opts *MinBlockRegisterOutputOpts) *MinBlock {
	if opts == nil {
		opts = &MinBlockRegisterOutputOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, name)
	args = append(args, jsType)

	args = append(args, opts.Target)

	retVal := m.p.Call("registerOutput", args...)
	return MinBlockFromJSObject(retVal, m.ctx)
}

// ReplaceRepeatableContent calls the ReplaceRepeatableContent method on the MinBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.minblock#replacerepeatablecontent
func (m *MinBlock) ReplaceRepeatableContent(vertexShaderState *NodeMaterialBuildState, fragmentShaderState *NodeMaterialBuildState, mesh *AbstractMesh, defines js.Value) {

	args := make([]interface{}, 0, 4+0)

	args = append(args, vertexShaderState.JSObject())
	args = append(args, fragmentShaderState.JSObject())
	args = append(args, mesh.JSObject())
	args = append(args, defines)

	m.p.Call("replaceRepeatableContent", args...)
}

// Serialize calls the Serialize method on the MinBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.minblock#serialize
func (m *MinBlock) Serialize() interface{} {

	retVal := m.p.Call("serialize")
	return retVal
}

// UpdateUniformsAndSamples calls the UpdateUniformsAndSamples method on the MinBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.minblock#updateuniformsandsamples
func (m *MinBlock) UpdateUniformsAndSamples(state *NodeMaterialBuildState, nodeMaterial *NodeMaterial, defines js.Value, uniformBuffers string) {

	args := make([]interface{}, 0, 4+0)

	args = append(args, state.JSObject())
	args = append(args, nodeMaterial.JSObject())
	args = append(args, defines)
	args = append(args, uniformBuffers)

	m.p.Call("updateUniformsAndSamples", args...)
}

/*

// BuildId returns the BuildId property of class MinBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.minblock#buildid
func (m *MinBlock) BuildId(buildId float64) *MinBlock {
	p := ba.ctx.Get("MinBlock").New(buildId)
	return MinBlockFromJSObject(p, ba.ctx)
}

// SetBuildId sets the BuildId property of class MinBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.minblock#buildid
func (m *MinBlock) SetBuildId(buildId float64) *MinBlock {
	p := ba.ctx.Get("MinBlock").New(buildId)
	return MinBlockFromJSObject(p, ba.ctx)
}

// Comments returns the Comments property of class MinBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.minblock#comments
func (m *MinBlock) Comments(comments string) *MinBlock {
	p := ba.ctx.Get("MinBlock").New(comments)
	return MinBlockFromJSObject(p, ba.ctx)
}

// SetComments sets the Comments property of class MinBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.minblock#comments
func (m *MinBlock) SetComments(comments string) *MinBlock {
	p := ba.ctx.Get("MinBlock").New(comments)
	return MinBlockFromJSObject(p, ba.ctx)
}

// Inputs returns the Inputs property of class MinBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.minblock#inputs
func (m *MinBlock) Inputs(inputs *NodeMaterialConnectionPoint) *MinBlock {
	p := ba.ctx.Get("MinBlock").New(inputs.JSObject())
	return MinBlockFromJSObject(p, ba.ctx)
}

// SetInputs sets the Inputs property of class MinBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.minblock#inputs
func (m *MinBlock) SetInputs(inputs *NodeMaterialConnectionPoint) *MinBlock {
	p := ba.ctx.Get("MinBlock").New(inputs.JSObject())
	return MinBlockFromJSObject(p, ba.ctx)
}

// IsFinalMerger returns the IsFinalMerger property of class MinBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.minblock#isfinalmerger
func (m *MinBlock) IsFinalMerger(isFinalMerger bool) *MinBlock {
	p := ba.ctx.Get("MinBlock").New(isFinalMerger)
	return MinBlockFromJSObject(p, ba.ctx)
}

// SetIsFinalMerger sets the IsFinalMerger property of class MinBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.minblock#isfinalmerger
func (m *MinBlock) SetIsFinalMerger(isFinalMerger bool) *MinBlock {
	p := ba.ctx.Get("MinBlock").New(isFinalMerger)
	return MinBlockFromJSObject(p, ba.ctx)
}

// IsInput returns the IsInput property of class MinBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.minblock#isinput
func (m *MinBlock) IsInput(isInput bool) *MinBlock {
	p := ba.ctx.Get("MinBlock").New(isInput)
	return MinBlockFromJSObject(p, ba.ctx)
}

// SetIsInput sets the IsInput property of class MinBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.minblock#isinput
func (m *MinBlock) SetIsInput(isInput bool) *MinBlock {
	p := ba.ctx.Get("MinBlock").New(isInput)
	return MinBlockFromJSObject(p, ba.ctx)
}

// IsUnique returns the IsUnique property of class MinBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.minblock#isunique
func (m *MinBlock) IsUnique(isUnique bool) *MinBlock {
	p := ba.ctx.Get("MinBlock").New(isUnique)
	return MinBlockFromJSObject(p, ba.ctx)
}

// SetIsUnique sets the IsUnique property of class MinBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.minblock#isunique
func (m *MinBlock) SetIsUnique(isUnique bool) *MinBlock {
	p := ba.ctx.Get("MinBlock").New(isUnique)
	return MinBlockFromJSObject(p, ba.ctx)
}

// Left returns the Left property of class MinBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.minblock#left
func (m *MinBlock) Left(left *NodeMaterialConnectionPoint) *MinBlock {
	p := ba.ctx.Get("MinBlock").New(left.JSObject())
	return MinBlockFromJSObject(p, ba.ctx)
}

// SetLeft sets the Left property of class MinBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.minblock#left
func (m *MinBlock) SetLeft(left *NodeMaterialConnectionPoint) *MinBlock {
	p := ba.ctx.Get("MinBlock").New(left.JSObject())
	return MinBlockFromJSObject(p, ba.ctx)
}

// Name returns the Name property of class MinBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.minblock#name
func (m *MinBlock) Name(name string) *MinBlock {
	p := ba.ctx.Get("MinBlock").New(name)
	return MinBlockFromJSObject(p, ba.ctx)
}

// SetName sets the Name property of class MinBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.minblock#name
func (m *MinBlock) SetName(name string) *MinBlock {
	p := ba.ctx.Get("MinBlock").New(name)
	return MinBlockFromJSObject(p, ba.ctx)
}

// Output returns the Output property of class MinBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.minblock#output
func (m *MinBlock) Output(output *NodeMaterialConnectionPoint) *MinBlock {
	p := ba.ctx.Get("MinBlock").New(output.JSObject())
	return MinBlockFromJSObject(p, ba.ctx)
}

// SetOutput sets the Output property of class MinBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.minblock#output
func (m *MinBlock) SetOutput(output *NodeMaterialConnectionPoint) *MinBlock {
	p := ba.ctx.Get("MinBlock").New(output.JSObject())
	return MinBlockFromJSObject(p, ba.ctx)
}

// Outputs returns the Outputs property of class MinBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.minblock#outputs
func (m *MinBlock) Outputs(outputs *NodeMaterialConnectionPoint) *MinBlock {
	p := ba.ctx.Get("MinBlock").New(outputs.JSObject())
	return MinBlockFromJSObject(p, ba.ctx)
}

// SetOutputs sets the Outputs property of class MinBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.minblock#outputs
func (m *MinBlock) SetOutputs(outputs *NodeMaterialConnectionPoint) *MinBlock {
	p := ba.ctx.Get("MinBlock").New(outputs.JSObject())
	return MinBlockFromJSObject(p, ba.ctx)
}

// Right returns the Right property of class MinBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.minblock#right
func (m *MinBlock) Right(right *NodeMaterialConnectionPoint) *MinBlock {
	p := ba.ctx.Get("MinBlock").New(right.JSObject())
	return MinBlockFromJSObject(p, ba.ctx)
}

// SetRight sets the Right property of class MinBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.minblock#right
func (m *MinBlock) SetRight(right *NodeMaterialConnectionPoint) *MinBlock {
	p := ba.ctx.Get("MinBlock").New(right.JSObject())
	return MinBlockFromJSObject(p, ba.ctx)
}

// Target returns the Target property of class MinBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.minblock#target
func (m *MinBlock) Target(target js.Value) *MinBlock {
	p := ba.ctx.Get("MinBlock").New(target)
	return MinBlockFromJSObject(p, ba.ctx)
}

// SetTarget sets the Target property of class MinBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.minblock#target
func (m *MinBlock) SetTarget(target js.Value) *MinBlock {
	p := ba.ctx.Get("MinBlock").New(target)
	return MinBlockFromJSObject(p, ba.ctx)
}

// UniqueId returns the UniqueId property of class MinBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.minblock#uniqueid
func (m *MinBlock) UniqueId(uniqueId float64) *MinBlock {
	p := ba.ctx.Get("MinBlock").New(uniqueId)
	return MinBlockFromJSObject(p, ba.ctx)
}

// SetUniqueId sets the UniqueId property of class MinBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.minblock#uniqueid
func (m *MinBlock) SetUniqueId(uniqueId float64) *MinBlock {
	p := ba.ctx.Get("MinBlock").New(uniqueId)
	return MinBlockFromJSObject(p, ba.ctx)
}

*/
