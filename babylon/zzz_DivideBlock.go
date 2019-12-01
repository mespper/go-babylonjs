// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// DivideBlock represents a babylon.js DivideBlock.
// Block used to divide 2 vectors
type DivideBlock struct {
	*NodeMaterialBlock
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (d *DivideBlock) JSObject() js.Value { return d.p }

// DivideBlock returns a DivideBlock JavaScript class.
func (ba *Babylon) DivideBlock() *DivideBlock {
	p := ba.ctx.Get("DivideBlock")
	return DivideBlockFromJSObject(p, ba.ctx)
}

// DivideBlockFromJSObject returns a wrapped DivideBlock JavaScript class.
func DivideBlockFromJSObject(p js.Value, ctx js.Value) *DivideBlock {
	return &DivideBlock{NodeMaterialBlock: NodeMaterialBlockFromJSObject(p, ctx), ctx: ctx}
}

// DivideBlockArrayToJSArray returns a JavaScript Array for the wrapped array.
func DivideBlockArrayToJSArray(array []*DivideBlock) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewDivideBlock returns a new DivideBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.divideblock
func (ba *Babylon) NewDivideBlock(name string) *DivideBlock {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	p := ba.ctx.Get("DivideBlock").New(args...)
	return DivideBlockFromJSObject(p, ba.ctx)
}

// AutoConfigure calls the AutoConfigure method on the DivideBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.divideblock#autoconfigure
func (d *DivideBlock) AutoConfigure(material *NodeMaterial) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, material.JSObject())

	d.p.Call("autoConfigure", args...)
}

// DivideBlockBindOpts contains optional parameters for DivideBlock.Bind.
type DivideBlockBindOpts struct {
	Mesh *Mesh
}

// Bind calls the Bind method on the DivideBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.divideblock#bind
func (d *DivideBlock) Bind(effect *Effect, nodeMaterial *NodeMaterial, opts *DivideBlockBindOpts) {
	if opts == nil {
		opts = &DivideBlockBindOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, effect.JSObject())
	args = append(args, nodeMaterial.JSObject())

	if opts.Mesh == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Mesh.JSObject())
	}

	d.p.Call("bind", args...)
}

// Build calls the Build method on the DivideBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.divideblock#build
func (d *DivideBlock) Build(state *NodeMaterialBuildState, activeBlocks *NodeMaterialBlock) bool {

	args := make([]interface{}, 0, 2+0)

	args = append(args, state.JSObject())
	args = append(args, activeBlocks.JSObject())

	retVal := d.p.Call("build", args...)
	return retVal.Bool()
}

// DivideBlockCloneOpts contains optional parameters for DivideBlock.Clone.
type DivideBlockCloneOpts struct {
	RootUrl *string
}

// Clone calls the Clone method on the DivideBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.divideblock#clone
func (d *DivideBlock) Clone(scene *Scene, opts *DivideBlockCloneOpts) *NodeMaterialBlock {
	if opts == nil {
		opts = &DivideBlockCloneOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, scene.JSObject())

	if opts.RootUrl == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.RootUrl)
	}

	retVal := d.p.Call("clone", args...)
	return NodeMaterialBlockFromJSObject(retVal, d.ctx)
}

// DivideBlockConnectToOpts contains optional parameters for DivideBlock.ConnectTo.
type DivideBlockConnectToOpts struct {
	Options map[string]interface{}
}

// ConnectTo calls the ConnectTo method on the DivideBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.divideblock#connectto
func (d *DivideBlock) ConnectTo(other *NodeMaterialBlock, opts *DivideBlockConnectToOpts) *DivideBlock {
	if opts == nil {
		opts = &DivideBlockConnectToOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, other.JSObject())

	if opts.Options == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Options)
	}

	retVal := d.p.Call("connectTo", args...)
	return DivideBlockFromJSObject(retVal, d.ctx)
}

// Dispose calls the Dispose method on the DivideBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.divideblock#dispose
func (d *DivideBlock) Dispose() {

	d.p.Call("dispose")
}

// GetClassName calls the GetClassName method on the DivideBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.divideblock#getclassname
func (d *DivideBlock) GetClassName() string {

	retVal := d.p.Call("getClassName")
	return retVal.String()
}

// DivideBlockGetFirstAvailableInputOpts contains optional parameters for DivideBlock.GetFirstAvailableInput.
type DivideBlockGetFirstAvailableInputOpts struct {
	ForOutput *NodeMaterialConnectionPoint
}

// GetFirstAvailableInput calls the GetFirstAvailableInput method on the DivideBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.divideblock#getfirstavailableinput
func (d *DivideBlock) GetFirstAvailableInput(opts *DivideBlockGetFirstAvailableInputOpts) *NodeMaterialConnectionPoint {
	if opts == nil {
		opts = &DivideBlockGetFirstAvailableInputOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.ForOutput == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.ForOutput.JSObject())
	}

	retVal := d.p.Call("getFirstAvailableInput", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, d.ctx)
}

// DivideBlockGetFirstAvailableOutputOpts contains optional parameters for DivideBlock.GetFirstAvailableOutput.
type DivideBlockGetFirstAvailableOutputOpts struct {
	ForBlock *NodeMaterialBlock
}

// GetFirstAvailableOutput calls the GetFirstAvailableOutput method on the DivideBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.divideblock#getfirstavailableoutput
func (d *DivideBlock) GetFirstAvailableOutput(opts *DivideBlockGetFirstAvailableOutputOpts) *NodeMaterialConnectionPoint {
	if opts == nil {
		opts = &DivideBlockGetFirstAvailableOutputOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.ForBlock == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.ForBlock.JSObject())
	}

	retVal := d.p.Call("getFirstAvailableOutput", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, d.ctx)
}

// GetInputByName calls the GetInputByName method on the DivideBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.divideblock#getinputbyname
func (d *DivideBlock) GetInputByName(name string) *NodeMaterialConnectionPoint {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	retVal := d.p.Call("getInputByName", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, d.ctx)
}

// GetOutputByName calls the GetOutputByName method on the DivideBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.divideblock#getoutputbyname
func (d *DivideBlock) GetOutputByName(name string) *NodeMaterialConnectionPoint {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	retVal := d.p.Call("getOutputByName", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, d.ctx)
}

// GetSiblingOutput calls the GetSiblingOutput method on the DivideBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.divideblock#getsiblingoutput
func (d *DivideBlock) GetSiblingOutput(current *NodeMaterialConnectionPoint) *NodeMaterialConnectionPoint {

	args := make([]interface{}, 0, 1+0)

	args = append(args, current.JSObject())

	retVal := d.p.Call("getSiblingOutput", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, d.ctx)
}

// Initialize calls the Initialize method on the DivideBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.divideblock#initialize
func (d *DivideBlock) Initialize(state *NodeMaterialBuildState) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, state.JSObject())

	d.p.Call("initialize", args...)
}

// DivideBlockInitializeDefinesOpts contains optional parameters for DivideBlock.InitializeDefines.
type DivideBlockInitializeDefinesOpts struct {
	UseInstances *bool
}

// InitializeDefines calls the InitializeDefines method on the DivideBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.divideblock#initializedefines
func (d *DivideBlock) InitializeDefines(mesh *AbstractMesh, nodeMaterial *NodeMaterial, defines js.Value, opts *DivideBlockInitializeDefinesOpts) {
	if opts == nil {
		opts = &DivideBlockInitializeDefinesOpts{}
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

	d.p.Call("initializeDefines", args...)
}

// DivideBlockIsReadyOpts contains optional parameters for DivideBlock.IsReady.
type DivideBlockIsReadyOpts struct {
	UseInstances *bool
}

// IsReady calls the IsReady method on the DivideBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.divideblock#isready
func (d *DivideBlock) IsReady(mesh *AbstractMesh, nodeMaterial *NodeMaterial, defines js.Value, opts *DivideBlockIsReadyOpts) bool {
	if opts == nil {
		opts = &DivideBlockIsReadyOpts{}
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

	retVal := d.p.Call("isReady", args...)
	return retVal.Bool()
}

// DivideBlockPrepareDefinesOpts contains optional parameters for DivideBlock.PrepareDefines.
type DivideBlockPrepareDefinesOpts struct {
	UseInstances *bool
}

// PrepareDefines calls the PrepareDefines method on the DivideBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.divideblock#preparedefines
func (d *DivideBlock) PrepareDefines(mesh *AbstractMesh, nodeMaterial *NodeMaterial, defines js.Value, opts *DivideBlockPrepareDefinesOpts) {
	if opts == nil {
		opts = &DivideBlockPrepareDefinesOpts{}
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

	d.p.Call("prepareDefines", args...)
}

// ProvideFallbacks calls the ProvideFallbacks method on the DivideBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.divideblock#providefallbacks
func (d *DivideBlock) ProvideFallbacks(mesh *AbstractMesh, fallbacks *EffectFallbacks) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, mesh.JSObject())
	args = append(args, fallbacks.JSObject())

	d.p.Call("provideFallbacks", args...)
}

// DivideBlockRegisterInputOpts contains optional parameters for DivideBlock.RegisterInput.
type DivideBlockRegisterInputOpts struct {
	IsOptional *bool
	Target     js.Value
}

// RegisterInput calls the RegisterInput method on the DivideBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.divideblock#registerinput
func (d *DivideBlock) RegisterInput(name string, jsType js.Value, opts *DivideBlockRegisterInputOpts) *DivideBlock {
	if opts == nil {
		opts = &DivideBlockRegisterInputOpts{}
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

	retVal := d.p.Call("registerInput", args...)
	return DivideBlockFromJSObject(retVal, d.ctx)
}

// DivideBlockRegisterOutputOpts contains optional parameters for DivideBlock.RegisterOutput.
type DivideBlockRegisterOutputOpts struct {
	Target js.Value
}

// RegisterOutput calls the RegisterOutput method on the DivideBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.divideblock#registeroutput
func (d *DivideBlock) RegisterOutput(name string, jsType js.Value, opts *DivideBlockRegisterOutputOpts) *DivideBlock {
	if opts == nil {
		opts = &DivideBlockRegisterOutputOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, name)
	args = append(args, jsType)

	args = append(args, opts.Target)

	retVal := d.p.Call("registerOutput", args...)
	return DivideBlockFromJSObject(retVal, d.ctx)
}

// ReplaceRepeatableContent calls the ReplaceRepeatableContent method on the DivideBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.divideblock#replacerepeatablecontent
func (d *DivideBlock) ReplaceRepeatableContent(vertexShaderState *NodeMaterialBuildState, fragmentShaderState *NodeMaterialBuildState, mesh *AbstractMesh, defines js.Value) {

	args := make([]interface{}, 0, 4+0)

	args = append(args, vertexShaderState.JSObject())
	args = append(args, fragmentShaderState.JSObject())
	args = append(args, mesh.JSObject())
	args = append(args, defines)

	d.p.Call("replaceRepeatableContent", args...)
}

// Serialize calls the Serialize method on the DivideBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.divideblock#serialize
func (d *DivideBlock) Serialize() interface{} {

	retVal := d.p.Call("serialize")
	return retVal
}

// UpdateUniformsAndSamples calls the UpdateUniformsAndSamples method on the DivideBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.divideblock#updateuniformsandsamples
func (d *DivideBlock) UpdateUniformsAndSamples(state *NodeMaterialBuildState, nodeMaterial *NodeMaterial, defines js.Value, uniformBuffers string) {

	args := make([]interface{}, 0, 4+0)

	args = append(args, state.JSObject())
	args = append(args, nodeMaterial.JSObject())
	args = append(args, defines)
	args = append(args, uniformBuffers)

	d.p.Call("updateUniformsAndSamples", args...)
}

/*

// BuildId returns the BuildId property of class DivideBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.divideblock#buildid
func (d *DivideBlock) BuildId(buildId float64) *DivideBlock {
	p := ba.ctx.Get("DivideBlock").New(buildId)
	return DivideBlockFromJSObject(p, ba.ctx)
}

// SetBuildId sets the BuildId property of class DivideBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.divideblock#buildid
func (d *DivideBlock) SetBuildId(buildId float64) *DivideBlock {
	p := ba.ctx.Get("DivideBlock").New(buildId)
	return DivideBlockFromJSObject(p, ba.ctx)
}

// Comments returns the Comments property of class DivideBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.divideblock#comments
func (d *DivideBlock) Comments(comments string) *DivideBlock {
	p := ba.ctx.Get("DivideBlock").New(comments)
	return DivideBlockFromJSObject(p, ba.ctx)
}

// SetComments sets the Comments property of class DivideBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.divideblock#comments
func (d *DivideBlock) SetComments(comments string) *DivideBlock {
	p := ba.ctx.Get("DivideBlock").New(comments)
	return DivideBlockFromJSObject(p, ba.ctx)
}

// Inputs returns the Inputs property of class DivideBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.divideblock#inputs
func (d *DivideBlock) Inputs(inputs *NodeMaterialConnectionPoint) *DivideBlock {
	p := ba.ctx.Get("DivideBlock").New(inputs.JSObject())
	return DivideBlockFromJSObject(p, ba.ctx)
}

// SetInputs sets the Inputs property of class DivideBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.divideblock#inputs
func (d *DivideBlock) SetInputs(inputs *NodeMaterialConnectionPoint) *DivideBlock {
	p := ba.ctx.Get("DivideBlock").New(inputs.JSObject())
	return DivideBlockFromJSObject(p, ba.ctx)
}

// IsFinalMerger returns the IsFinalMerger property of class DivideBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.divideblock#isfinalmerger
func (d *DivideBlock) IsFinalMerger(isFinalMerger bool) *DivideBlock {
	p := ba.ctx.Get("DivideBlock").New(isFinalMerger)
	return DivideBlockFromJSObject(p, ba.ctx)
}

// SetIsFinalMerger sets the IsFinalMerger property of class DivideBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.divideblock#isfinalmerger
func (d *DivideBlock) SetIsFinalMerger(isFinalMerger bool) *DivideBlock {
	p := ba.ctx.Get("DivideBlock").New(isFinalMerger)
	return DivideBlockFromJSObject(p, ba.ctx)
}

// IsInput returns the IsInput property of class DivideBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.divideblock#isinput
func (d *DivideBlock) IsInput(isInput bool) *DivideBlock {
	p := ba.ctx.Get("DivideBlock").New(isInput)
	return DivideBlockFromJSObject(p, ba.ctx)
}

// SetIsInput sets the IsInput property of class DivideBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.divideblock#isinput
func (d *DivideBlock) SetIsInput(isInput bool) *DivideBlock {
	p := ba.ctx.Get("DivideBlock").New(isInput)
	return DivideBlockFromJSObject(p, ba.ctx)
}

// IsUnique returns the IsUnique property of class DivideBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.divideblock#isunique
func (d *DivideBlock) IsUnique(isUnique bool) *DivideBlock {
	p := ba.ctx.Get("DivideBlock").New(isUnique)
	return DivideBlockFromJSObject(p, ba.ctx)
}

// SetIsUnique sets the IsUnique property of class DivideBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.divideblock#isunique
func (d *DivideBlock) SetIsUnique(isUnique bool) *DivideBlock {
	p := ba.ctx.Get("DivideBlock").New(isUnique)
	return DivideBlockFromJSObject(p, ba.ctx)
}

// Left returns the Left property of class DivideBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.divideblock#left
func (d *DivideBlock) Left(left *NodeMaterialConnectionPoint) *DivideBlock {
	p := ba.ctx.Get("DivideBlock").New(left.JSObject())
	return DivideBlockFromJSObject(p, ba.ctx)
}

// SetLeft sets the Left property of class DivideBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.divideblock#left
func (d *DivideBlock) SetLeft(left *NodeMaterialConnectionPoint) *DivideBlock {
	p := ba.ctx.Get("DivideBlock").New(left.JSObject())
	return DivideBlockFromJSObject(p, ba.ctx)
}

// Name returns the Name property of class DivideBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.divideblock#name
func (d *DivideBlock) Name(name string) *DivideBlock {
	p := ba.ctx.Get("DivideBlock").New(name)
	return DivideBlockFromJSObject(p, ba.ctx)
}

// SetName sets the Name property of class DivideBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.divideblock#name
func (d *DivideBlock) SetName(name string) *DivideBlock {
	p := ba.ctx.Get("DivideBlock").New(name)
	return DivideBlockFromJSObject(p, ba.ctx)
}

// Output returns the Output property of class DivideBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.divideblock#output
func (d *DivideBlock) Output(output *NodeMaterialConnectionPoint) *DivideBlock {
	p := ba.ctx.Get("DivideBlock").New(output.JSObject())
	return DivideBlockFromJSObject(p, ba.ctx)
}

// SetOutput sets the Output property of class DivideBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.divideblock#output
func (d *DivideBlock) SetOutput(output *NodeMaterialConnectionPoint) *DivideBlock {
	p := ba.ctx.Get("DivideBlock").New(output.JSObject())
	return DivideBlockFromJSObject(p, ba.ctx)
}

// Outputs returns the Outputs property of class DivideBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.divideblock#outputs
func (d *DivideBlock) Outputs(outputs *NodeMaterialConnectionPoint) *DivideBlock {
	p := ba.ctx.Get("DivideBlock").New(outputs.JSObject())
	return DivideBlockFromJSObject(p, ba.ctx)
}

// SetOutputs sets the Outputs property of class DivideBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.divideblock#outputs
func (d *DivideBlock) SetOutputs(outputs *NodeMaterialConnectionPoint) *DivideBlock {
	p := ba.ctx.Get("DivideBlock").New(outputs.JSObject())
	return DivideBlockFromJSObject(p, ba.ctx)
}

// Right returns the Right property of class DivideBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.divideblock#right
func (d *DivideBlock) Right(right *NodeMaterialConnectionPoint) *DivideBlock {
	p := ba.ctx.Get("DivideBlock").New(right.JSObject())
	return DivideBlockFromJSObject(p, ba.ctx)
}

// SetRight sets the Right property of class DivideBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.divideblock#right
func (d *DivideBlock) SetRight(right *NodeMaterialConnectionPoint) *DivideBlock {
	p := ba.ctx.Get("DivideBlock").New(right.JSObject())
	return DivideBlockFromJSObject(p, ba.ctx)
}

// Target returns the Target property of class DivideBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.divideblock#target
func (d *DivideBlock) Target(target js.Value) *DivideBlock {
	p := ba.ctx.Get("DivideBlock").New(target)
	return DivideBlockFromJSObject(p, ba.ctx)
}

// SetTarget sets the Target property of class DivideBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.divideblock#target
func (d *DivideBlock) SetTarget(target js.Value) *DivideBlock {
	p := ba.ctx.Get("DivideBlock").New(target)
	return DivideBlockFromJSObject(p, ba.ctx)
}

// UniqueId returns the UniqueId property of class DivideBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.divideblock#uniqueid
func (d *DivideBlock) UniqueId(uniqueId float64) *DivideBlock {
	p := ba.ctx.Get("DivideBlock").New(uniqueId)
	return DivideBlockFromJSObject(p, ba.ctx)
}

// SetUniqueId sets the UniqueId property of class DivideBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.divideblock#uniqueid
func (d *DivideBlock) SetUniqueId(uniqueId float64) *DivideBlock {
	p := ba.ctx.Get("DivideBlock").New(uniqueId)
	return DivideBlockFromJSObject(p, ba.ctx)
}

*/
