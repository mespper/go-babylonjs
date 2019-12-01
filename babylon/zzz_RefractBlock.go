// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// RefractBlock represents a babylon.js RefractBlock.
// Block used to get the refracted vector from a direction and a normal
type RefractBlock struct {
	*NodeMaterialBlock
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (r *RefractBlock) JSObject() js.Value { return r.p }

// RefractBlock returns a RefractBlock JavaScript class.
func (ba *Babylon) RefractBlock() *RefractBlock {
	p := ba.ctx.Get("RefractBlock")
	return RefractBlockFromJSObject(p, ba.ctx)
}

// RefractBlockFromJSObject returns a wrapped RefractBlock JavaScript class.
func RefractBlockFromJSObject(p js.Value, ctx js.Value) *RefractBlock {
	return &RefractBlock{NodeMaterialBlock: NodeMaterialBlockFromJSObject(p, ctx), ctx: ctx}
}

// RefractBlockArrayToJSArray returns a JavaScript Array for the wrapped array.
func RefractBlockArrayToJSArray(array []*RefractBlock) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewRefractBlock returns a new RefractBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.refractblock
func (ba *Babylon) NewRefractBlock(name string) *RefractBlock {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	p := ba.ctx.Get("RefractBlock").New(args...)
	return RefractBlockFromJSObject(p, ba.ctx)
}

// AutoConfigure calls the AutoConfigure method on the RefractBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.refractblock#autoconfigure
func (r *RefractBlock) AutoConfigure(material *NodeMaterial) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, material.JSObject())

	r.p.Call("autoConfigure", args...)
}

// RefractBlockBindOpts contains optional parameters for RefractBlock.Bind.
type RefractBlockBindOpts struct {
	Mesh *Mesh
}

// Bind calls the Bind method on the RefractBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.refractblock#bind
func (r *RefractBlock) Bind(effect *Effect, nodeMaterial *NodeMaterial, opts *RefractBlockBindOpts) {
	if opts == nil {
		opts = &RefractBlockBindOpts{}
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

// Build calls the Build method on the RefractBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.refractblock#build
func (r *RefractBlock) Build(state *NodeMaterialBuildState, activeBlocks *NodeMaterialBlock) bool {

	args := make([]interface{}, 0, 2+0)

	args = append(args, state.JSObject())
	args = append(args, activeBlocks.JSObject())

	retVal := r.p.Call("build", args...)
	return retVal.Bool()
}

// RefractBlockCloneOpts contains optional parameters for RefractBlock.Clone.
type RefractBlockCloneOpts struct {
	RootUrl *string
}

// Clone calls the Clone method on the RefractBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.refractblock#clone
func (r *RefractBlock) Clone(scene *Scene, opts *RefractBlockCloneOpts) *NodeMaterialBlock {
	if opts == nil {
		opts = &RefractBlockCloneOpts{}
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

// RefractBlockConnectToOpts contains optional parameters for RefractBlock.ConnectTo.
type RefractBlockConnectToOpts struct {
	Options map[string]interface{}
}

// ConnectTo calls the ConnectTo method on the RefractBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.refractblock#connectto
func (r *RefractBlock) ConnectTo(other *NodeMaterialBlock, opts *RefractBlockConnectToOpts) *RefractBlock {
	if opts == nil {
		opts = &RefractBlockConnectToOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, other.JSObject())

	if opts.Options == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Options)
	}

	retVal := r.p.Call("connectTo", args...)
	return RefractBlockFromJSObject(retVal, r.ctx)
}

// Dispose calls the Dispose method on the RefractBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.refractblock#dispose
func (r *RefractBlock) Dispose() {

	r.p.Call("dispose")
}

// GetClassName calls the GetClassName method on the RefractBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.refractblock#getclassname
func (r *RefractBlock) GetClassName() string {

	retVal := r.p.Call("getClassName")
	return retVal.String()
}

// RefractBlockGetFirstAvailableInputOpts contains optional parameters for RefractBlock.GetFirstAvailableInput.
type RefractBlockGetFirstAvailableInputOpts struct {
	ForOutput *NodeMaterialConnectionPoint
}

// GetFirstAvailableInput calls the GetFirstAvailableInput method on the RefractBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.refractblock#getfirstavailableinput
func (r *RefractBlock) GetFirstAvailableInput(opts *RefractBlockGetFirstAvailableInputOpts) *NodeMaterialConnectionPoint {
	if opts == nil {
		opts = &RefractBlockGetFirstAvailableInputOpts{}
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

// RefractBlockGetFirstAvailableOutputOpts contains optional parameters for RefractBlock.GetFirstAvailableOutput.
type RefractBlockGetFirstAvailableOutputOpts struct {
	ForBlock *NodeMaterialBlock
}

// GetFirstAvailableOutput calls the GetFirstAvailableOutput method on the RefractBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.refractblock#getfirstavailableoutput
func (r *RefractBlock) GetFirstAvailableOutput(opts *RefractBlockGetFirstAvailableOutputOpts) *NodeMaterialConnectionPoint {
	if opts == nil {
		opts = &RefractBlockGetFirstAvailableOutputOpts{}
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

// GetInputByName calls the GetInputByName method on the RefractBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.refractblock#getinputbyname
func (r *RefractBlock) GetInputByName(name string) *NodeMaterialConnectionPoint {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	retVal := r.p.Call("getInputByName", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, r.ctx)
}

// GetOutputByName calls the GetOutputByName method on the RefractBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.refractblock#getoutputbyname
func (r *RefractBlock) GetOutputByName(name string) *NodeMaterialConnectionPoint {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	retVal := r.p.Call("getOutputByName", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, r.ctx)
}

// GetSiblingOutput calls the GetSiblingOutput method on the RefractBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.refractblock#getsiblingoutput
func (r *RefractBlock) GetSiblingOutput(current *NodeMaterialConnectionPoint) *NodeMaterialConnectionPoint {

	args := make([]interface{}, 0, 1+0)

	args = append(args, current.JSObject())

	retVal := r.p.Call("getSiblingOutput", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, r.ctx)
}

// Initialize calls the Initialize method on the RefractBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.refractblock#initialize
func (r *RefractBlock) Initialize(state *NodeMaterialBuildState) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, state.JSObject())

	r.p.Call("initialize", args...)
}

// RefractBlockInitializeDefinesOpts contains optional parameters for RefractBlock.InitializeDefines.
type RefractBlockInitializeDefinesOpts struct {
	UseInstances *bool
}

// InitializeDefines calls the InitializeDefines method on the RefractBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.refractblock#initializedefines
func (r *RefractBlock) InitializeDefines(mesh *AbstractMesh, nodeMaterial *NodeMaterial, defines js.Value, opts *RefractBlockInitializeDefinesOpts) {
	if opts == nil {
		opts = &RefractBlockInitializeDefinesOpts{}
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

// RefractBlockIsReadyOpts contains optional parameters for RefractBlock.IsReady.
type RefractBlockIsReadyOpts struct {
	UseInstances *bool
}

// IsReady calls the IsReady method on the RefractBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.refractblock#isready
func (r *RefractBlock) IsReady(mesh *AbstractMesh, nodeMaterial *NodeMaterial, defines js.Value, opts *RefractBlockIsReadyOpts) bool {
	if opts == nil {
		opts = &RefractBlockIsReadyOpts{}
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

// RefractBlockPrepareDefinesOpts contains optional parameters for RefractBlock.PrepareDefines.
type RefractBlockPrepareDefinesOpts struct {
	UseInstances *bool
}

// PrepareDefines calls the PrepareDefines method on the RefractBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.refractblock#preparedefines
func (r *RefractBlock) PrepareDefines(mesh *AbstractMesh, nodeMaterial *NodeMaterial, defines js.Value, opts *RefractBlockPrepareDefinesOpts) {
	if opts == nil {
		opts = &RefractBlockPrepareDefinesOpts{}
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

// ProvideFallbacks calls the ProvideFallbacks method on the RefractBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.refractblock#providefallbacks
func (r *RefractBlock) ProvideFallbacks(mesh *AbstractMesh, fallbacks *EffectFallbacks) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, mesh.JSObject())
	args = append(args, fallbacks.JSObject())

	r.p.Call("provideFallbacks", args...)
}

// RefractBlockRegisterInputOpts contains optional parameters for RefractBlock.RegisterInput.
type RefractBlockRegisterInputOpts struct {
	IsOptional *bool
	Target     js.Value
}

// RegisterInput calls the RegisterInput method on the RefractBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.refractblock#registerinput
func (r *RefractBlock) RegisterInput(name string, jsType js.Value, opts *RefractBlockRegisterInputOpts) *RefractBlock {
	if opts == nil {
		opts = &RefractBlockRegisterInputOpts{}
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
	return RefractBlockFromJSObject(retVal, r.ctx)
}

// RefractBlockRegisterOutputOpts contains optional parameters for RefractBlock.RegisterOutput.
type RefractBlockRegisterOutputOpts struct {
	Target js.Value
}

// RegisterOutput calls the RegisterOutput method on the RefractBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.refractblock#registeroutput
func (r *RefractBlock) RegisterOutput(name string, jsType js.Value, opts *RefractBlockRegisterOutputOpts) *RefractBlock {
	if opts == nil {
		opts = &RefractBlockRegisterOutputOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, name)
	args = append(args, jsType)

	args = append(args, opts.Target)

	retVal := r.p.Call("registerOutput", args...)
	return RefractBlockFromJSObject(retVal, r.ctx)
}

// ReplaceRepeatableContent calls the ReplaceRepeatableContent method on the RefractBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.refractblock#replacerepeatablecontent
func (r *RefractBlock) ReplaceRepeatableContent(vertexShaderState *NodeMaterialBuildState, fragmentShaderState *NodeMaterialBuildState, mesh *AbstractMesh, defines js.Value) {

	args := make([]interface{}, 0, 4+0)

	args = append(args, vertexShaderState.JSObject())
	args = append(args, fragmentShaderState.JSObject())
	args = append(args, mesh.JSObject())
	args = append(args, defines)

	r.p.Call("replaceRepeatableContent", args...)
}

// Serialize calls the Serialize method on the RefractBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.refractblock#serialize
func (r *RefractBlock) Serialize() interface{} {

	retVal := r.p.Call("serialize")
	return retVal
}

// UpdateUniformsAndSamples calls the UpdateUniformsAndSamples method on the RefractBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.refractblock#updateuniformsandsamples
func (r *RefractBlock) UpdateUniformsAndSamples(state *NodeMaterialBuildState, nodeMaterial *NodeMaterial, defines js.Value, uniformBuffers string) {

	args := make([]interface{}, 0, 4+0)

	args = append(args, state.JSObject())
	args = append(args, nodeMaterial.JSObject())
	args = append(args, defines)
	args = append(args, uniformBuffers)

	r.p.Call("updateUniformsAndSamples", args...)
}

/*

// BuildId returns the BuildId property of class RefractBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.refractblock#buildid
func (r *RefractBlock) BuildId(buildId float64) *RefractBlock {
	p := ba.ctx.Get("RefractBlock").New(buildId)
	return RefractBlockFromJSObject(p, ba.ctx)
}

// SetBuildId sets the BuildId property of class RefractBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.refractblock#buildid
func (r *RefractBlock) SetBuildId(buildId float64) *RefractBlock {
	p := ba.ctx.Get("RefractBlock").New(buildId)
	return RefractBlockFromJSObject(p, ba.ctx)
}

// Comments returns the Comments property of class RefractBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.refractblock#comments
func (r *RefractBlock) Comments(comments string) *RefractBlock {
	p := ba.ctx.Get("RefractBlock").New(comments)
	return RefractBlockFromJSObject(p, ba.ctx)
}

// SetComments sets the Comments property of class RefractBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.refractblock#comments
func (r *RefractBlock) SetComments(comments string) *RefractBlock {
	p := ba.ctx.Get("RefractBlock").New(comments)
	return RefractBlockFromJSObject(p, ba.ctx)
}

// Incident returns the Incident property of class RefractBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.refractblock#incident
func (r *RefractBlock) Incident(incident *NodeMaterialConnectionPoint) *RefractBlock {
	p := ba.ctx.Get("RefractBlock").New(incident.JSObject())
	return RefractBlockFromJSObject(p, ba.ctx)
}

// SetIncident sets the Incident property of class RefractBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.refractblock#incident
func (r *RefractBlock) SetIncident(incident *NodeMaterialConnectionPoint) *RefractBlock {
	p := ba.ctx.Get("RefractBlock").New(incident.JSObject())
	return RefractBlockFromJSObject(p, ba.ctx)
}

// Inputs returns the Inputs property of class RefractBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.refractblock#inputs
func (r *RefractBlock) Inputs(inputs *NodeMaterialConnectionPoint) *RefractBlock {
	p := ba.ctx.Get("RefractBlock").New(inputs.JSObject())
	return RefractBlockFromJSObject(p, ba.ctx)
}

// SetInputs sets the Inputs property of class RefractBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.refractblock#inputs
func (r *RefractBlock) SetInputs(inputs *NodeMaterialConnectionPoint) *RefractBlock {
	p := ba.ctx.Get("RefractBlock").New(inputs.JSObject())
	return RefractBlockFromJSObject(p, ba.ctx)
}

// Ior returns the Ior property of class RefractBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.refractblock#ior
func (r *RefractBlock) Ior(ior *NodeMaterialConnectionPoint) *RefractBlock {
	p := ba.ctx.Get("RefractBlock").New(ior.JSObject())
	return RefractBlockFromJSObject(p, ba.ctx)
}

// SetIor sets the Ior property of class RefractBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.refractblock#ior
func (r *RefractBlock) SetIor(ior *NodeMaterialConnectionPoint) *RefractBlock {
	p := ba.ctx.Get("RefractBlock").New(ior.JSObject())
	return RefractBlockFromJSObject(p, ba.ctx)
}

// IsFinalMerger returns the IsFinalMerger property of class RefractBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.refractblock#isfinalmerger
func (r *RefractBlock) IsFinalMerger(isFinalMerger bool) *RefractBlock {
	p := ba.ctx.Get("RefractBlock").New(isFinalMerger)
	return RefractBlockFromJSObject(p, ba.ctx)
}

// SetIsFinalMerger sets the IsFinalMerger property of class RefractBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.refractblock#isfinalmerger
func (r *RefractBlock) SetIsFinalMerger(isFinalMerger bool) *RefractBlock {
	p := ba.ctx.Get("RefractBlock").New(isFinalMerger)
	return RefractBlockFromJSObject(p, ba.ctx)
}

// IsInput returns the IsInput property of class RefractBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.refractblock#isinput
func (r *RefractBlock) IsInput(isInput bool) *RefractBlock {
	p := ba.ctx.Get("RefractBlock").New(isInput)
	return RefractBlockFromJSObject(p, ba.ctx)
}

// SetIsInput sets the IsInput property of class RefractBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.refractblock#isinput
func (r *RefractBlock) SetIsInput(isInput bool) *RefractBlock {
	p := ba.ctx.Get("RefractBlock").New(isInput)
	return RefractBlockFromJSObject(p, ba.ctx)
}

// IsUnique returns the IsUnique property of class RefractBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.refractblock#isunique
func (r *RefractBlock) IsUnique(isUnique bool) *RefractBlock {
	p := ba.ctx.Get("RefractBlock").New(isUnique)
	return RefractBlockFromJSObject(p, ba.ctx)
}

// SetIsUnique sets the IsUnique property of class RefractBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.refractblock#isunique
func (r *RefractBlock) SetIsUnique(isUnique bool) *RefractBlock {
	p := ba.ctx.Get("RefractBlock").New(isUnique)
	return RefractBlockFromJSObject(p, ba.ctx)
}

// Name returns the Name property of class RefractBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.refractblock#name
func (r *RefractBlock) Name(name string) *RefractBlock {
	p := ba.ctx.Get("RefractBlock").New(name)
	return RefractBlockFromJSObject(p, ba.ctx)
}

// SetName sets the Name property of class RefractBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.refractblock#name
func (r *RefractBlock) SetName(name string) *RefractBlock {
	p := ba.ctx.Get("RefractBlock").New(name)
	return RefractBlockFromJSObject(p, ba.ctx)
}

// Normal returns the Normal property of class RefractBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.refractblock#normal
func (r *RefractBlock) Normal(normal *NodeMaterialConnectionPoint) *RefractBlock {
	p := ba.ctx.Get("RefractBlock").New(normal.JSObject())
	return RefractBlockFromJSObject(p, ba.ctx)
}

// SetNormal sets the Normal property of class RefractBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.refractblock#normal
func (r *RefractBlock) SetNormal(normal *NodeMaterialConnectionPoint) *RefractBlock {
	p := ba.ctx.Get("RefractBlock").New(normal.JSObject())
	return RefractBlockFromJSObject(p, ba.ctx)
}

// Output returns the Output property of class RefractBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.refractblock#output
func (r *RefractBlock) Output(output *NodeMaterialConnectionPoint) *RefractBlock {
	p := ba.ctx.Get("RefractBlock").New(output.JSObject())
	return RefractBlockFromJSObject(p, ba.ctx)
}

// SetOutput sets the Output property of class RefractBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.refractblock#output
func (r *RefractBlock) SetOutput(output *NodeMaterialConnectionPoint) *RefractBlock {
	p := ba.ctx.Get("RefractBlock").New(output.JSObject())
	return RefractBlockFromJSObject(p, ba.ctx)
}

// Outputs returns the Outputs property of class RefractBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.refractblock#outputs
func (r *RefractBlock) Outputs(outputs *NodeMaterialConnectionPoint) *RefractBlock {
	p := ba.ctx.Get("RefractBlock").New(outputs.JSObject())
	return RefractBlockFromJSObject(p, ba.ctx)
}

// SetOutputs sets the Outputs property of class RefractBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.refractblock#outputs
func (r *RefractBlock) SetOutputs(outputs *NodeMaterialConnectionPoint) *RefractBlock {
	p := ba.ctx.Get("RefractBlock").New(outputs.JSObject())
	return RefractBlockFromJSObject(p, ba.ctx)
}

// Target returns the Target property of class RefractBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.refractblock#target
func (r *RefractBlock) Target(target js.Value) *RefractBlock {
	p := ba.ctx.Get("RefractBlock").New(target)
	return RefractBlockFromJSObject(p, ba.ctx)
}

// SetTarget sets the Target property of class RefractBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.refractblock#target
func (r *RefractBlock) SetTarget(target js.Value) *RefractBlock {
	p := ba.ctx.Get("RefractBlock").New(target)
	return RefractBlockFromJSObject(p, ba.ctx)
}

// UniqueId returns the UniqueId property of class RefractBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.refractblock#uniqueid
func (r *RefractBlock) UniqueId(uniqueId float64) *RefractBlock {
	p := ba.ctx.Get("RefractBlock").New(uniqueId)
	return RefractBlockFromJSObject(p, ba.ctx)
}

// SetUniqueId sets the UniqueId property of class RefractBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.refractblock#uniqueid
func (r *RefractBlock) SetUniqueId(uniqueId float64) *RefractBlock {
	p := ba.ctx.Get("RefractBlock").New(uniqueId)
	return RefractBlockFromJSObject(p, ba.ctx)
}

*/
