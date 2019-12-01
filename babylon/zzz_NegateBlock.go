// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// NegateBlock represents a babylon.js NegateBlock.
// Block used to get negative version of a value (i.e. x * -1)
type NegateBlock struct {
	*NodeMaterialBlock
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (n *NegateBlock) JSObject() js.Value { return n.p }

// NegateBlock returns a NegateBlock JavaScript class.
func (ba *Babylon) NegateBlock() *NegateBlock {
	p := ba.ctx.Get("NegateBlock")
	return NegateBlockFromJSObject(p, ba.ctx)
}

// NegateBlockFromJSObject returns a wrapped NegateBlock JavaScript class.
func NegateBlockFromJSObject(p js.Value, ctx js.Value) *NegateBlock {
	return &NegateBlock{NodeMaterialBlock: NodeMaterialBlockFromJSObject(p, ctx), ctx: ctx}
}

// NegateBlockArrayToJSArray returns a JavaScript Array for the wrapped array.
func NegateBlockArrayToJSArray(array []*NegateBlock) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewNegateBlock returns a new NegateBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.negateblock
func (ba *Babylon) NewNegateBlock(name string) *NegateBlock {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	p := ba.ctx.Get("NegateBlock").New(args...)
	return NegateBlockFromJSObject(p, ba.ctx)
}

// AutoConfigure calls the AutoConfigure method on the NegateBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.negateblock#autoconfigure
func (n *NegateBlock) AutoConfigure(material *NodeMaterial) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, material.JSObject())

	n.p.Call("autoConfigure", args...)
}

// NegateBlockBindOpts contains optional parameters for NegateBlock.Bind.
type NegateBlockBindOpts struct {
	Mesh *Mesh
}

// Bind calls the Bind method on the NegateBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.negateblock#bind
func (n *NegateBlock) Bind(effect *Effect, nodeMaterial *NodeMaterial, opts *NegateBlockBindOpts) {
	if opts == nil {
		opts = &NegateBlockBindOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, effect.JSObject())
	args = append(args, nodeMaterial.JSObject())

	if opts.Mesh == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Mesh.JSObject())
	}

	n.p.Call("bind", args...)
}

// Build calls the Build method on the NegateBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.negateblock#build
func (n *NegateBlock) Build(state *NodeMaterialBuildState, activeBlocks *NodeMaterialBlock) bool {

	args := make([]interface{}, 0, 2+0)

	args = append(args, state.JSObject())
	args = append(args, activeBlocks.JSObject())

	retVal := n.p.Call("build", args...)
	return retVal.Bool()
}

// NegateBlockCloneOpts contains optional parameters for NegateBlock.Clone.
type NegateBlockCloneOpts struct {
	RootUrl *string
}

// Clone calls the Clone method on the NegateBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.negateblock#clone
func (n *NegateBlock) Clone(scene *Scene, opts *NegateBlockCloneOpts) *NodeMaterialBlock {
	if opts == nil {
		opts = &NegateBlockCloneOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, scene.JSObject())

	if opts.RootUrl == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.RootUrl)
	}

	retVal := n.p.Call("clone", args...)
	return NodeMaterialBlockFromJSObject(retVal, n.ctx)
}

// NegateBlockConnectToOpts contains optional parameters for NegateBlock.ConnectTo.
type NegateBlockConnectToOpts struct {
	Options map[string]interface{}
}

// ConnectTo calls the ConnectTo method on the NegateBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.negateblock#connectto
func (n *NegateBlock) ConnectTo(other *NodeMaterialBlock, opts *NegateBlockConnectToOpts) *NegateBlock {
	if opts == nil {
		opts = &NegateBlockConnectToOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, other.JSObject())

	if opts.Options == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Options)
	}

	retVal := n.p.Call("connectTo", args...)
	return NegateBlockFromJSObject(retVal, n.ctx)
}

// Dispose calls the Dispose method on the NegateBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.negateblock#dispose
func (n *NegateBlock) Dispose() {

	n.p.Call("dispose")
}

// GetClassName calls the GetClassName method on the NegateBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.negateblock#getclassname
func (n *NegateBlock) GetClassName() string {

	retVal := n.p.Call("getClassName")
	return retVal.String()
}

// NegateBlockGetFirstAvailableInputOpts contains optional parameters for NegateBlock.GetFirstAvailableInput.
type NegateBlockGetFirstAvailableInputOpts struct {
	ForOutput *NodeMaterialConnectionPoint
}

// GetFirstAvailableInput calls the GetFirstAvailableInput method on the NegateBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.negateblock#getfirstavailableinput
func (n *NegateBlock) GetFirstAvailableInput(opts *NegateBlockGetFirstAvailableInputOpts) *NodeMaterialConnectionPoint {
	if opts == nil {
		opts = &NegateBlockGetFirstAvailableInputOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.ForOutput == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.ForOutput.JSObject())
	}

	retVal := n.p.Call("getFirstAvailableInput", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, n.ctx)
}

// NegateBlockGetFirstAvailableOutputOpts contains optional parameters for NegateBlock.GetFirstAvailableOutput.
type NegateBlockGetFirstAvailableOutputOpts struct {
	ForBlock *NodeMaterialBlock
}

// GetFirstAvailableOutput calls the GetFirstAvailableOutput method on the NegateBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.negateblock#getfirstavailableoutput
func (n *NegateBlock) GetFirstAvailableOutput(opts *NegateBlockGetFirstAvailableOutputOpts) *NodeMaterialConnectionPoint {
	if opts == nil {
		opts = &NegateBlockGetFirstAvailableOutputOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.ForBlock == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.ForBlock.JSObject())
	}

	retVal := n.p.Call("getFirstAvailableOutput", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, n.ctx)
}

// GetInputByName calls the GetInputByName method on the NegateBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.negateblock#getinputbyname
func (n *NegateBlock) GetInputByName(name string) *NodeMaterialConnectionPoint {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	retVal := n.p.Call("getInputByName", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, n.ctx)
}

// GetOutputByName calls the GetOutputByName method on the NegateBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.negateblock#getoutputbyname
func (n *NegateBlock) GetOutputByName(name string) *NodeMaterialConnectionPoint {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	retVal := n.p.Call("getOutputByName", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, n.ctx)
}

// GetSiblingOutput calls the GetSiblingOutput method on the NegateBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.negateblock#getsiblingoutput
func (n *NegateBlock) GetSiblingOutput(current *NodeMaterialConnectionPoint) *NodeMaterialConnectionPoint {

	args := make([]interface{}, 0, 1+0)

	args = append(args, current.JSObject())

	retVal := n.p.Call("getSiblingOutput", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, n.ctx)
}

// Initialize calls the Initialize method on the NegateBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.negateblock#initialize
func (n *NegateBlock) Initialize(state *NodeMaterialBuildState) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, state.JSObject())

	n.p.Call("initialize", args...)
}

// NegateBlockInitializeDefinesOpts contains optional parameters for NegateBlock.InitializeDefines.
type NegateBlockInitializeDefinesOpts struct {
	UseInstances *bool
}

// InitializeDefines calls the InitializeDefines method on the NegateBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.negateblock#initializedefines
func (n *NegateBlock) InitializeDefines(mesh *AbstractMesh, nodeMaterial *NodeMaterial, defines js.Value, opts *NegateBlockInitializeDefinesOpts) {
	if opts == nil {
		opts = &NegateBlockInitializeDefinesOpts{}
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

	n.p.Call("initializeDefines", args...)
}

// NegateBlockIsReadyOpts contains optional parameters for NegateBlock.IsReady.
type NegateBlockIsReadyOpts struct {
	UseInstances *bool
}

// IsReady calls the IsReady method on the NegateBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.negateblock#isready
func (n *NegateBlock) IsReady(mesh *AbstractMesh, nodeMaterial *NodeMaterial, defines js.Value, opts *NegateBlockIsReadyOpts) bool {
	if opts == nil {
		opts = &NegateBlockIsReadyOpts{}
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

	retVal := n.p.Call("isReady", args...)
	return retVal.Bool()
}

// NegateBlockPrepareDefinesOpts contains optional parameters for NegateBlock.PrepareDefines.
type NegateBlockPrepareDefinesOpts struct {
	UseInstances *bool
}

// PrepareDefines calls the PrepareDefines method on the NegateBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.negateblock#preparedefines
func (n *NegateBlock) PrepareDefines(mesh *AbstractMesh, nodeMaterial *NodeMaterial, defines js.Value, opts *NegateBlockPrepareDefinesOpts) {
	if opts == nil {
		opts = &NegateBlockPrepareDefinesOpts{}
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

	n.p.Call("prepareDefines", args...)
}

// ProvideFallbacks calls the ProvideFallbacks method on the NegateBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.negateblock#providefallbacks
func (n *NegateBlock) ProvideFallbacks(mesh *AbstractMesh, fallbacks *EffectFallbacks) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, mesh.JSObject())
	args = append(args, fallbacks.JSObject())

	n.p.Call("provideFallbacks", args...)
}

// NegateBlockRegisterInputOpts contains optional parameters for NegateBlock.RegisterInput.
type NegateBlockRegisterInputOpts struct {
	IsOptional *bool
	Target     js.Value
}

// RegisterInput calls the RegisterInput method on the NegateBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.negateblock#registerinput
func (n *NegateBlock) RegisterInput(name string, jsType js.Value, opts *NegateBlockRegisterInputOpts) *NegateBlock {
	if opts == nil {
		opts = &NegateBlockRegisterInputOpts{}
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

	retVal := n.p.Call("registerInput", args...)
	return NegateBlockFromJSObject(retVal, n.ctx)
}

// NegateBlockRegisterOutputOpts contains optional parameters for NegateBlock.RegisterOutput.
type NegateBlockRegisterOutputOpts struct {
	Target js.Value
}

// RegisterOutput calls the RegisterOutput method on the NegateBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.negateblock#registeroutput
func (n *NegateBlock) RegisterOutput(name string, jsType js.Value, opts *NegateBlockRegisterOutputOpts) *NegateBlock {
	if opts == nil {
		opts = &NegateBlockRegisterOutputOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, name)
	args = append(args, jsType)

	args = append(args, opts.Target)

	retVal := n.p.Call("registerOutput", args...)
	return NegateBlockFromJSObject(retVal, n.ctx)
}

// ReplaceRepeatableContent calls the ReplaceRepeatableContent method on the NegateBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.negateblock#replacerepeatablecontent
func (n *NegateBlock) ReplaceRepeatableContent(vertexShaderState *NodeMaterialBuildState, fragmentShaderState *NodeMaterialBuildState, mesh *AbstractMesh, defines js.Value) {

	args := make([]interface{}, 0, 4+0)

	args = append(args, vertexShaderState.JSObject())
	args = append(args, fragmentShaderState.JSObject())
	args = append(args, mesh.JSObject())
	args = append(args, defines)

	n.p.Call("replaceRepeatableContent", args...)
}

// Serialize calls the Serialize method on the NegateBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.negateblock#serialize
func (n *NegateBlock) Serialize() interface{} {

	retVal := n.p.Call("serialize")
	return retVal
}

// UpdateUniformsAndSamples calls the UpdateUniformsAndSamples method on the NegateBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.negateblock#updateuniformsandsamples
func (n *NegateBlock) UpdateUniformsAndSamples(state *NodeMaterialBuildState, nodeMaterial *NodeMaterial, defines js.Value, uniformBuffers string) {

	args := make([]interface{}, 0, 4+0)

	args = append(args, state.JSObject())
	args = append(args, nodeMaterial.JSObject())
	args = append(args, defines)
	args = append(args, uniformBuffers)

	n.p.Call("updateUniformsAndSamples", args...)
}

/*

// BuildId returns the BuildId property of class NegateBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.negateblock#buildid
func (n *NegateBlock) BuildId(buildId float64) *NegateBlock {
	p := ba.ctx.Get("NegateBlock").New(buildId)
	return NegateBlockFromJSObject(p, ba.ctx)
}

// SetBuildId sets the BuildId property of class NegateBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.negateblock#buildid
func (n *NegateBlock) SetBuildId(buildId float64) *NegateBlock {
	p := ba.ctx.Get("NegateBlock").New(buildId)
	return NegateBlockFromJSObject(p, ba.ctx)
}

// Comments returns the Comments property of class NegateBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.negateblock#comments
func (n *NegateBlock) Comments(comments string) *NegateBlock {
	p := ba.ctx.Get("NegateBlock").New(comments)
	return NegateBlockFromJSObject(p, ba.ctx)
}

// SetComments sets the Comments property of class NegateBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.negateblock#comments
func (n *NegateBlock) SetComments(comments string) *NegateBlock {
	p := ba.ctx.Get("NegateBlock").New(comments)
	return NegateBlockFromJSObject(p, ba.ctx)
}

// Inputs returns the Inputs property of class NegateBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.negateblock#inputs
func (n *NegateBlock) Inputs(inputs *NodeMaterialConnectionPoint) *NegateBlock {
	p := ba.ctx.Get("NegateBlock").New(inputs.JSObject())
	return NegateBlockFromJSObject(p, ba.ctx)
}

// SetInputs sets the Inputs property of class NegateBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.negateblock#inputs
func (n *NegateBlock) SetInputs(inputs *NodeMaterialConnectionPoint) *NegateBlock {
	p := ba.ctx.Get("NegateBlock").New(inputs.JSObject())
	return NegateBlockFromJSObject(p, ba.ctx)
}

// IsFinalMerger returns the IsFinalMerger property of class NegateBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.negateblock#isfinalmerger
func (n *NegateBlock) IsFinalMerger(isFinalMerger bool) *NegateBlock {
	p := ba.ctx.Get("NegateBlock").New(isFinalMerger)
	return NegateBlockFromJSObject(p, ba.ctx)
}

// SetIsFinalMerger sets the IsFinalMerger property of class NegateBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.negateblock#isfinalmerger
func (n *NegateBlock) SetIsFinalMerger(isFinalMerger bool) *NegateBlock {
	p := ba.ctx.Get("NegateBlock").New(isFinalMerger)
	return NegateBlockFromJSObject(p, ba.ctx)
}

// IsInput returns the IsInput property of class NegateBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.negateblock#isinput
func (n *NegateBlock) IsInput(isInput bool) *NegateBlock {
	p := ba.ctx.Get("NegateBlock").New(isInput)
	return NegateBlockFromJSObject(p, ba.ctx)
}

// SetIsInput sets the IsInput property of class NegateBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.negateblock#isinput
func (n *NegateBlock) SetIsInput(isInput bool) *NegateBlock {
	p := ba.ctx.Get("NegateBlock").New(isInput)
	return NegateBlockFromJSObject(p, ba.ctx)
}

// IsUnique returns the IsUnique property of class NegateBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.negateblock#isunique
func (n *NegateBlock) IsUnique(isUnique bool) *NegateBlock {
	p := ba.ctx.Get("NegateBlock").New(isUnique)
	return NegateBlockFromJSObject(p, ba.ctx)
}

// SetIsUnique sets the IsUnique property of class NegateBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.negateblock#isunique
func (n *NegateBlock) SetIsUnique(isUnique bool) *NegateBlock {
	p := ba.ctx.Get("NegateBlock").New(isUnique)
	return NegateBlockFromJSObject(p, ba.ctx)
}

// Name returns the Name property of class NegateBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.negateblock#name
func (n *NegateBlock) Name(name string) *NegateBlock {
	p := ba.ctx.Get("NegateBlock").New(name)
	return NegateBlockFromJSObject(p, ba.ctx)
}

// SetName sets the Name property of class NegateBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.negateblock#name
func (n *NegateBlock) SetName(name string) *NegateBlock {
	p := ba.ctx.Get("NegateBlock").New(name)
	return NegateBlockFromJSObject(p, ba.ctx)
}

// Output returns the Output property of class NegateBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.negateblock#output
func (n *NegateBlock) Output(output *NodeMaterialConnectionPoint) *NegateBlock {
	p := ba.ctx.Get("NegateBlock").New(output.JSObject())
	return NegateBlockFromJSObject(p, ba.ctx)
}

// SetOutput sets the Output property of class NegateBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.negateblock#output
func (n *NegateBlock) SetOutput(output *NodeMaterialConnectionPoint) *NegateBlock {
	p := ba.ctx.Get("NegateBlock").New(output.JSObject())
	return NegateBlockFromJSObject(p, ba.ctx)
}

// Outputs returns the Outputs property of class NegateBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.negateblock#outputs
func (n *NegateBlock) Outputs(outputs *NodeMaterialConnectionPoint) *NegateBlock {
	p := ba.ctx.Get("NegateBlock").New(outputs.JSObject())
	return NegateBlockFromJSObject(p, ba.ctx)
}

// SetOutputs sets the Outputs property of class NegateBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.negateblock#outputs
func (n *NegateBlock) SetOutputs(outputs *NodeMaterialConnectionPoint) *NegateBlock {
	p := ba.ctx.Get("NegateBlock").New(outputs.JSObject())
	return NegateBlockFromJSObject(p, ba.ctx)
}

// Target returns the Target property of class NegateBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.negateblock#target
func (n *NegateBlock) Target(target js.Value) *NegateBlock {
	p := ba.ctx.Get("NegateBlock").New(target)
	return NegateBlockFromJSObject(p, ba.ctx)
}

// SetTarget sets the Target property of class NegateBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.negateblock#target
func (n *NegateBlock) SetTarget(target js.Value) *NegateBlock {
	p := ba.ctx.Get("NegateBlock").New(target)
	return NegateBlockFromJSObject(p, ba.ctx)
}

// UniqueId returns the UniqueId property of class NegateBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.negateblock#uniqueid
func (n *NegateBlock) UniqueId(uniqueId float64) *NegateBlock {
	p := ba.ctx.Get("NegateBlock").New(uniqueId)
	return NegateBlockFromJSObject(p, ba.ctx)
}

// SetUniqueId sets the UniqueId property of class NegateBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.negateblock#uniqueid
func (n *NegateBlock) SetUniqueId(uniqueId float64) *NegateBlock {
	p := ba.ctx.Get("NegateBlock").New(uniqueId)
	return NegateBlockFromJSObject(p, ba.ctx)
}

// Value returns the Value property of class NegateBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.negateblock#value
func (n *NegateBlock) Value(value *NodeMaterialConnectionPoint) *NegateBlock {
	p := ba.ctx.Get("NegateBlock").New(value.JSObject())
	return NegateBlockFromJSObject(p, ba.ctx)
}

// SetValue sets the Value property of class NegateBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.negateblock#value
func (n *NegateBlock) SetValue(value *NodeMaterialConnectionPoint) *NegateBlock {
	p := ba.ctx.Get("NegateBlock").New(value.JSObject())
	return NegateBlockFromJSObject(p, ba.ctx)
}

*/
