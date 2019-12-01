// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ReflectBlock represents a babylon.js ReflectBlock.
// Block used to get the reflected vector from a direction and a normal
type ReflectBlock struct {
	*NodeMaterialBlock
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (r *ReflectBlock) JSObject() js.Value { return r.p }

// ReflectBlock returns a ReflectBlock JavaScript class.
func (ba *Babylon) ReflectBlock() *ReflectBlock {
	p := ba.ctx.Get("ReflectBlock")
	return ReflectBlockFromJSObject(p, ba.ctx)
}

// ReflectBlockFromJSObject returns a wrapped ReflectBlock JavaScript class.
func ReflectBlockFromJSObject(p js.Value, ctx js.Value) *ReflectBlock {
	return &ReflectBlock{NodeMaterialBlock: NodeMaterialBlockFromJSObject(p, ctx), ctx: ctx}
}

// ReflectBlockArrayToJSArray returns a JavaScript Array for the wrapped array.
func ReflectBlockArrayToJSArray(array []*ReflectBlock) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewReflectBlock returns a new ReflectBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.reflectblock
func (ba *Babylon) NewReflectBlock(name string) *ReflectBlock {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	p := ba.ctx.Get("ReflectBlock").New(args...)
	return ReflectBlockFromJSObject(p, ba.ctx)
}

// AutoConfigure calls the AutoConfigure method on the ReflectBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.reflectblock#autoconfigure
func (r *ReflectBlock) AutoConfigure(material *NodeMaterial) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, material.JSObject())

	r.p.Call("autoConfigure", args...)
}

// ReflectBlockBindOpts contains optional parameters for ReflectBlock.Bind.
type ReflectBlockBindOpts struct {
	Mesh *Mesh
}

// Bind calls the Bind method on the ReflectBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.reflectblock#bind
func (r *ReflectBlock) Bind(effect *Effect, nodeMaterial *NodeMaterial, opts *ReflectBlockBindOpts) {
	if opts == nil {
		opts = &ReflectBlockBindOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, effect.JSObject())
	args = append(args, nodeMaterial.JSObject())

	if opts.Mesh == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Mesh.JSObject())
	}

	r.p.Call("bind", args...)
}

// Build calls the Build method on the ReflectBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.reflectblock#build
func (r *ReflectBlock) Build(state *NodeMaterialBuildState, activeBlocks *NodeMaterialBlock) bool {

	args := make([]interface{}, 0, 2+0)

	args = append(args, state.JSObject())
	args = append(args, activeBlocks.JSObject())

	retVal := r.p.Call("build", args...)
	return retVal.Bool()
}

// ReflectBlockCloneOpts contains optional parameters for ReflectBlock.Clone.
type ReflectBlockCloneOpts struct {
	RootUrl *string
}

// Clone calls the Clone method on the ReflectBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.reflectblock#clone
func (r *ReflectBlock) Clone(scene *Scene, opts *ReflectBlockCloneOpts) *NodeMaterialBlock {
	if opts == nil {
		opts = &ReflectBlockCloneOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, scene.JSObject())

	if opts.RootUrl == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.RootUrl)
	}

	retVal := r.p.Call("clone", args...)
	return NodeMaterialBlockFromJSObject(retVal, r.ctx)
}

// ReflectBlockConnectToOpts contains optional parameters for ReflectBlock.ConnectTo.
type ReflectBlockConnectToOpts struct {
	Options map[string]interface{}
}

// ConnectTo calls the ConnectTo method on the ReflectBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.reflectblock#connectto
func (r *ReflectBlock) ConnectTo(other *NodeMaterialBlock, opts *ReflectBlockConnectToOpts) *ReflectBlock {
	if opts == nil {
		opts = &ReflectBlockConnectToOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, other.JSObject())

	if opts.Options == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Options)
	}

	retVal := r.p.Call("connectTo", args...)
	return ReflectBlockFromJSObject(retVal, r.ctx)
}

// Dispose calls the Dispose method on the ReflectBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.reflectblock#dispose
func (r *ReflectBlock) Dispose() {

	r.p.Call("dispose")
}

// GetClassName calls the GetClassName method on the ReflectBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.reflectblock#getclassname
func (r *ReflectBlock) GetClassName() string {

	retVal := r.p.Call("getClassName")
	return retVal.String()
}

// ReflectBlockGetFirstAvailableInputOpts contains optional parameters for ReflectBlock.GetFirstAvailableInput.
type ReflectBlockGetFirstAvailableInputOpts struct {
	ForOutput *NodeMaterialConnectionPoint
}

// GetFirstAvailableInput calls the GetFirstAvailableInput method on the ReflectBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.reflectblock#getfirstavailableinput
func (r *ReflectBlock) GetFirstAvailableInput(opts *ReflectBlockGetFirstAvailableInputOpts) *NodeMaterialConnectionPoint {
	if opts == nil {
		opts = &ReflectBlockGetFirstAvailableInputOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.ForOutput == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.ForOutput.JSObject())
	}

	retVal := r.p.Call("getFirstAvailableInput", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, r.ctx)
}

// ReflectBlockGetFirstAvailableOutputOpts contains optional parameters for ReflectBlock.GetFirstAvailableOutput.
type ReflectBlockGetFirstAvailableOutputOpts struct {
	ForBlock *NodeMaterialBlock
}

// GetFirstAvailableOutput calls the GetFirstAvailableOutput method on the ReflectBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.reflectblock#getfirstavailableoutput
func (r *ReflectBlock) GetFirstAvailableOutput(opts *ReflectBlockGetFirstAvailableOutputOpts) *NodeMaterialConnectionPoint {
	if opts == nil {
		opts = &ReflectBlockGetFirstAvailableOutputOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.ForBlock == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.ForBlock.JSObject())
	}

	retVal := r.p.Call("getFirstAvailableOutput", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, r.ctx)
}

// GetInputByName calls the GetInputByName method on the ReflectBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.reflectblock#getinputbyname
func (r *ReflectBlock) GetInputByName(name string) *NodeMaterialConnectionPoint {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	retVal := r.p.Call("getInputByName", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, r.ctx)
}

// GetOutputByName calls the GetOutputByName method on the ReflectBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.reflectblock#getoutputbyname
func (r *ReflectBlock) GetOutputByName(name string) *NodeMaterialConnectionPoint {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	retVal := r.p.Call("getOutputByName", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, r.ctx)
}

// GetSiblingOutput calls the GetSiblingOutput method on the ReflectBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.reflectblock#getsiblingoutput
func (r *ReflectBlock) GetSiblingOutput(current *NodeMaterialConnectionPoint) *NodeMaterialConnectionPoint {

	args := make([]interface{}, 0, 1+0)

	args = append(args, current.JSObject())

	retVal := r.p.Call("getSiblingOutput", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, r.ctx)
}

// Initialize calls the Initialize method on the ReflectBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.reflectblock#initialize
func (r *ReflectBlock) Initialize(state *NodeMaterialBuildState) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, state.JSObject())

	r.p.Call("initialize", args...)
}

// ReflectBlockInitializeDefinesOpts contains optional parameters for ReflectBlock.InitializeDefines.
type ReflectBlockInitializeDefinesOpts struct {
	UseInstances *bool
}

// InitializeDefines calls the InitializeDefines method on the ReflectBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.reflectblock#initializedefines
func (r *ReflectBlock) InitializeDefines(mesh *AbstractMesh, nodeMaterial *NodeMaterial, defines js.Value, opts *ReflectBlockInitializeDefinesOpts) {
	if opts == nil {
		opts = &ReflectBlockInitializeDefinesOpts{}
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

	r.p.Call("initializeDefines", args...)
}

// ReflectBlockIsReadyOpts contains optional parameters for ReflectBlock.IsReady.
type ReflectBlockIsReadyOpts struct {
	UseInstances *bool
}

// IsReady calls the IsReady method on the ReflectBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.reflectblock#isready
func (r *ReflectBlock) IsReady(mesh *AbstractMesh, nodeMaterial *NodeMaterial, defines js.Value, opts *ReflectBlockIsReadyOpts) bool {
	if opts == nil {
		opts = &ReflectBlockIsReadyOpts{}
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

	retVal := r.p.Call("isReady", args...)
	return retVal.Bool()
}

// ReflectBlockPrepareDefinesOpts contains optional parameters for ReflectBlock.PrepareDefines.
type ReflectBlockPrepareDefinesOpts struct {
	UseInstances *bool
}

// PrepareDefines calls the PrepareDefines method on the ReflectBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.reflectblock#preparedefines
func (r *ReflectBlock) PrepareDefines(mesh *AbstractMesh, nodeMaterial *NodeMaterial, defines js.Value, opts *ReflectBlockPrepareDefinesOpts) {
	if opts == nil {
		opts = &ReflectBlockPrepareDefinesOpts{}
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

	r.p.Call("prepareDefines", args...)
}

// ProvideFallbacks calls the ProvideFallbacks method on the ReflectBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.reflectblock#providefallbacks
func (r *ReflectBlock) ProvideFallbacks(mesh *AbstractMesh, fallbacks *EffectFallbacks) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, mesh.JSObject())
	args = append(args, fallbacks.JSObject())

	r.p.Call("provideFallbacks", args...)
}

// ReflectBlockRegisterInputOpts contains optional parameters for ReflectBlock.RegisterInput.
type ReflectBlockRegisterInputOpts struct {
	IsOptional *bool
	Target     js.Value
}

// RegisterInput calls the RegisterInput method on the ReflectBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.reflectblock#registerinput
func (r *ReflectBlock) RegisterInput(name string, jsType js.Value, opts *ReflectBlockRegisterInputOpts) *ReflectBlock {
	if opts == nil {
		opts = &ReflectBlockRegisterInputOpts{}
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

	retVal := r.p.Call("registerInput", args...)
	return ReflectBlockFromJSObject(retVal, r.ctx)
}

// ReflectBlockRegisterOutputOpts contains optional parameters for ReflectBlock.RegisterOutput.
type ReflectBlockRegisterOutputOpts struct {
	Target js.Value
}

// RegisterOutput calls the RegisterOutput method on the ReflectBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.reflectblock#registeroutput
func (r *ReflectBlock) RegisterOutput(name string, jsType js.Value, opts *ReflectBlockRegisterOutputOpts) *ReflectBlock {
	if opts == nil {
		opts = &ReflectBlockRegisterOutputOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, name)
	args = append(args, jsType)

	args = append(args, opts.Target)

	retVal := r.p.Call("registerOutput", args...)
	return ReflectBlockFromJSObject(retVal, r.ctx)
}

// ReplaceRepeatableContent calls the ReplaceRepeatableContent method on the ReflectBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.reflectblock#replacerepeatablecontent
func (r *ReflectBlock) ReplaceRepeatableContent(vertexShaderState *NodeMaterialBuildState, fragmentShaderState *NodeMaterialBuildState, mesh *AbstractMesh, defines js.Value) {

	args := make([]interface{}, 0, 4+0)

	args = append(args, vertexShaderState.JSObject())
	args = append(args, fragmentShaderState.JSObject())
	args = append(args, mesh.JSObject())
	args = append(args, defines)

	r.p.Call("replaceRepeatableContent", args...)
}

// Serialize calls the Serialize method on the ReflectBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.reflectblock#serialize
func (r *ReflectBlock) Serialize() interface{} {

	retVal := r.p.Call("serialize")
	return retVal
}

// UpdateUniformsAndSamples calls the UpdateUniformsAndSamples method on the ReflectBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.reflectblock#updateuniformsandsamples
func (r *ReflectBlock) UpdateUniformsAndSamples(state *NodeMaterialBuildState, nodeMaterial *NodeMaterial, defines js.Value, uniformBuffers string) {

	args := make([]interface{}, 0, 4+0)

	args = append(args, state.JSObject())
	args = append(args, nodeMaterial.JSObject())
	args = append(args, defines)
	args = append(args, uniformBuffers)

	r.p.Call("updateUniformsAndSamples", args...)
}

/*

// BuildId returns the BuildId property of class ReflectBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.reflectblock#buildid
func (r *ReflectBlock) BuildId(buildId float64) *ReflectBlock {
	p := ba.ctx.Get("ReflectBlock").New(buildId)
	return ReflectBlockFromJSObject(p, ba.ctx)
}

// SetBuildId sets the BuildId property of class ReflectBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.reflectblock#buildid
func (r *ReflectBlock) SetBuildId(buildId float64) *ReflectBlock {
	p := ba.ctx.Get("ReflectBlock").New(buildId)
	return ReflectBlockFromJSObject(p, ba.ctx)
}

// Comments returns the Comments property of class ReflectBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.reflectblock#comments
func (r *ReflectBlock) Comments(comments string) *ReflectBlock {
	p := ba.ctx.Get("ReflectBlock").New(comments)
	return ReflectBlockFromJSObject(p, ba.ctx)
}

// SetComments sets the Comments property of class ReflectBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.reflectblock#comments
func (r *ReflectBlock) SetComments(comments string) *ReflectBlock {
	p := ba.ctx.Get("ReflectBlock").New(comments)
	return ReflectBlockFromJSObject(p, ba.ctx)
}

// Incident returns the Incident property of class ReflectBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.reflectblock#incident
func (r *ReflectBlock) Incident(incident *NodeMaterialConnectionPoint) *ReflectBlock {
	p := ba.ctx.Get("ReflectBlock").New(incident.JSObject())
	return ReflectBlockFromJSObject(p, ba.ctx)
}

// SetIncident sets the Incident property of class ReflectBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.reflectblock#incident
func (r *ReflectBlock) SetIncident(incident *NodeMaterialConnectionPoint) *ReflectBlock {
	p := ba.ctx.Get("ReflectBlock").New(incident.JSObject())
	return ReflectBlockFromJSObject(p, ba.ctx)
}

// Inputs returns the Inputs property of class ReflectBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.reflectblock#inputs
func (r *ReflectBlock) Inputs(inputs *NodeMaterialConnectionPoint) *ReflectBlock {
	p := ba.ctx.Get("ReflectBlock").New(inputs.JSObject())
	return ReflectBlockFromJSObject(p, ba.ctx)
}

// SetInputs sets the Inputs property of class ReflectBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.reflectblock#inputs
func (r *ReflectBlock) SetInputs(inputs *NodeMaterialConnectionPoint) *ReflectBlock {
	p := ba.ctx.Get("ReflectBlock").New(inputs.JSObject())
	return ReflectBlockFromJSObject(p, ba.ctx)
}

// IsFinalMerger returns the IsFinalMerger property of class ReflectBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.reflectblock#isfinalmerger
func (r *ReflectBlock) IsFinalMerger(isFinalMerger bool) *ReflectBlock {
	p := ba.ctx.Get("ReflectBlock").New(isFinalMerger)
	return ReflectBlockFromJSObject(p, ba.ctx)
}

// SetIsFinalMerger sets the IsFinalMerger property of class ReflectBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.reflectblock#isfinalmerger
func (r *ReflectBlock) SetIsFinalMerger(isFinalMerger bool) *ReflectBlock {
	p := ba.ctx.Get("ReflectBlock").New(isFinalMerger)
	return ReflectBlockFromJSObject(p, ba.ctx)
}

// IsInput returns the IsInput property of class ReflectBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.reflectblock#isinput
func (r *ReflectBlock) IsInput(isInput bool) *ReflectBlock {
	p := ba.ctx.Get("ReflectBlock").New(isInput)
	return ReflectBlockFromJSObject(p, ba.ctx)
}

// SetIsInput sets the IsInput property of class ReflectBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.reflectblock#isinput
func (r *ReflectBlock) SetIsInput(isInput bool) *ReflectBlock {
	p := ba.ctx.Get("ReflectBlock").New(isInput)
	return ReflectBlockFromJSObject(p, ba.ctx)
}

// IsUnique returns the IsUnique property of class ReflectBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.reflectblock#isunique
func (r *ReflectBlock) IsUnique(isUnique bool) *ReflectBlock {
	p := ba.ctx.Get("ReflectBlock").New(isUnique)
	return ReflectBlockFromJSObject(p, ba.ctx)
}

// SetIsUnique sets the IsUnique property of class ReflectBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.reflectblock#isunique
func (r *ReflectBlock) SetIsUnique(isUnique bool) *ReflectBlock {
	p := ba.ctx.Get("ReflectBlock").New(isUnique)
	return ReflectBlockFromJSObject(p, ba.ctx)
}

// Name returns the Name property of class ReflectBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.reflectblock#name
func (r *ReflectBlock) Name(name string) *ReflectBlock {
	p := ba.ctx.Get("ReflectBlock").New(name)
	return ReflectBlockFromJSObject(p, ba.ctx)
}

// SetName sets the Name property of class ReflectBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.reflectblock#name
func (r *ReflectBlock) SetName(name string) *ReflectBlock {
	p := ba.ctx.Get("ReflectBlock").New(name)
	return ReflectBlockFromJSObject(p, ba.ctx)
}

// Normal returns the Normal property of class ReflectBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.reflectblock#normal
func (r *ReflectBlock) Normal(normal *NodeMaterialConnectionPoint) *ReflectBlock {
	p := ba.ctx.Get("ReflectBlock").New(normal.JSObject())
	return ReflectBlockFromJSObject(p, ba.ctx)
}

// SetNormal sets the Normal property of class ReflectBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.reflectblock#normal
func (r *ReflectBlock) SetNormal(normal *NodeMaterialConnectionPoint) *ReflectBlock {
	p := ba.ctx.Get("ReflectBlock").New(normal.JSObject())
	return ReflectBlockFromJSObject(p, ba.ctx)
}

// Output returns the Output property of class ReflectBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.reflectblock#output
func (r *ReflectBlock) Output(output *NodeMaterialConnectionPoint) *ReflectBlock {
	p := ba.ctx.Get("ReflectBlock").New(output.JSObject())
	return ReflectBlockFromJSObject(p, ba.ctx)
}

// SetOutput sets the Output property of class ReflectBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.reflectblock#output
func (r *ReflectBlock) SetOutput(output *NodeMaterialConnectionPoint) *ReflectBlock {
	p := ba.ctx.Get("ReflectBlock").New(output.JSObject())
	return ReflectBlockFromJSObject(p, ba.ctx)
}

// Outputs returns the Outputs property of class ReflectBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.reflectblock#outputs
func (r *ReflectBlock) Outputs(outputs *NodeMaterialConnectionPoint) *ReflectBlock {
	p := ba.ctx.Get("ReflectBlock").New(outputs.JSObject())
	return ReflectBlockFromJSObject(p, ba.ctx)
}

// SetOutputs sets the Outputs property of class ReflectBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.reflectblock#outputs
func (r *ReflectBlock) SetOutputs(outputs *NodeMaterialConnectionPoint) *ReflectBlock {
	p := ba.ctx.Get("ReflectBlock").New(outputs.JSObject())
	return ReflectBlockFromJSObject(p, ba.ctx)
}

// Target returns the Target property of class ReflectBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.reflectblock#target
func (r *ReflectBlock) Target(target js.Value) *ReflectBlock {
	p := ba.ctx.Get("ReflectBlock").New(target)
	return ReflectBlockFromJSObject(p, ba.ctx)
}

// SetTarget sets the Target property of class ReflectBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.reflectblock#target
func (r *ReflectBlock) SetTarget(target js.Value) *ReflectBlock {
	p := ba.ctx.Get("ReflectBlock").New(target)
	return ReflectBlockFromJSObject(p, ba.ctx)
}

// UniqueId returns the UniqueId property of class ReflectBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.reflectblock#uniqueid
func (r *ReflectBlock) UniqueId(uniqueId float64) *ReflectBlock {
	p := ba.ctx.Get("ReflectBlock").New(uniqueId)
	return ReflectBlockFromJSObject(p, ba.ctx)
}

// SetUniqueId sets the UniqueId property of class ReflectBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.reflectblock#uniqueid
func (r *ReflectBlock) SetUniqueId(uniqueId float64) *ReflectBlock {
	p := ba.ctx.Get("ReflectBlock").New(uniqueId)
	return ReflectBlockFromJSObject(p, ba.ctx)
}

*/
