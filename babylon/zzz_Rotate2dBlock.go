// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// Rotate2dBlock represents a babylon.js Rotate2dBlock.
// Block used to rotate a 2d vector by a given angle
type Rotate2dBlock struct {
	*NodeMaterialBlock
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (r *Rotate2dBlock) JSObject() js.Value { return r.p }

// Rotate2dBlock returns a Rotate2dBlock JavaScript class.
func (ba *Babylon) Rotate2dBlock() *Rotate2dBlock {
	p := ba.ctx.Get("Rotate2dBlock")
	return Rotate2dBlockFromJSObject(p, ba.ctx)
}

// Rotate2dBlockFromJSObject returns a wrapped Rotate2dBlock JavaScript class.
func Rotate2dBlockFromJSObject(p js.Value, ctx js.Value) *Rotate2dBlock {
	return &Rotate2dBlock{NodeMaterialBlock: NodeMaterialBlockFromJSObject(p, ctx), ctx: ctx}
}

// Rotate2dBlockArrayToJSArray returns a JavaScript Array for the wrapped array.
func Rotate2dBlockArrayToJSArray(array []*Rotate2dBlock) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewRotate2dBlock returns a new Rotate2dBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.rotate2dblock
func (ba *Babylon) NewRotate2dBlock(name string) *Rotate2dBlock {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	p := ba.ctx.Get("Rotate2dBlock").New(args...)
	return Rotate2dBlockFromJSObject(p, ba.ctx)
}

// AutoConfigure calls the AutoConfigure method on the Rotate2dBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.rotate2dblock#autoconfigure
func (r *Rotate2dBlock) AutoConfigure(material *NodeMaterial) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, material.JSObject())

	r.p.Call("autoConfigure", args...)
}

// Rotate2dBlockBindOpts contains optional parameters for Rotate2dBlock.Bind.
type Rotate2dBlockBindOpts struct {
	Mesh *Mesh
}

// Bind calls the Bind method on the Rotate2dBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.rotate2dblock#bind
func (r *Rotate2dBlock) Bind(effect *Effect, nodeMaterial *NodeMaterial, opts *Rotate2dBlockBindOpts) {
	if opts == nil {
		opts = &Rotate2dBlockBindOpts{}
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

// Build calls the Build method on the Rotate2dBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.rotate2dblock#build
func (r *Rotate2dBlock) Build(state *NodeMaterialBuildState, activeBlocks *NodeMaterialBlock) bool {

	args := make([]interface{}, 0, 2+0)

	args = append(args, state.JSObject())
	args = append(args, activeBlocks.JSObject())

	retVal := r.p.Call("build", args...)
	return retVal.Bool()
}

// Rotate2dBlockCloneOpts contains optional parameters for Rotate2dBlock.Clone.
type Rotate2dBlockCloneOpts struct {
	RootUrl *string
}

// Clone calls the Clone method on the Rotate2dBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.rotate2dblock#clone
func (r *Rotate2dBlock) Clone(scene *Scene, opts *Rotate2dBlockCloneOpts) *NodeMaterialBlock {
	if opts == nil {
		opts = &Rotate2dBlockCloneOpts{}
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

// Rotate2dBlockConnectToOpts contains optional parameters for Rotate2dBlock.ConnectTo.
type Rotate2dBlockConnectToOpts struct {
	Options map[string]interface{}
}

// ConnectTo calls the ConnectTo method on the Rotate2dBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.rotate2dblock#connectto
func (r *Rotate2dBlock) ConnectTo(other *NodeMaterialBlock, opts *Rotate2dBlockConnectToOpts) *Rotate2dBlock {
	if opts == nil {
		opts = &Rotate2dBlockConnectToOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, other.JSObject())

	if opts.Options == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Options)
	}

	retVal := r.p.Call("connectTo", args...)
	return Rotate2dBlockFromJSObject(retVal, r.ctx)
}

// Dispose calls the Dispose method on the Rotate2dBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.rotate2dblock#dispose
func (r *Rotate2dBlock) Dispose() {

	r.p.Call("dispose")
}

// GetClassName calls the GetClassName method on the Rotate2dBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.rotate2dblock#getclassname
func (r *Rotate2dBlock) GetClassName() string {

	retVal := r.p.Call("getClassName")
	return retVal.String()
}

// Rotate2dBlockGetFirstAvailableInputOpts contains optional parameters for Rotate2dBlock.GetFirstAvailableInput.
type Rotate2dBlockGetFirstAvailableInputOpts struct {
	ForOutput *NodeMaterialConnectionPoint
}

// GetFirstAvailableInput calls the GetFirstAvailableInput method on the Rotate2dBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.rotate2dblock#getfirstavailableinput
func (r *Rotate2dBlock) GetFirstAvailableInput(opts *Rotate2dBlockGetFirstAvailableInputOpts) *NodeMaterialConnectionPoint {
	if opts == nil {
		opts = &Rotate2dBlockGetFirstAvailableInputOpts{}
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

// Rotate2dBlockGetFirstAvailableOutputOpts contains optional parameters for Rotate2dBlock.GetFirstAvailableOutput.
type Rotate2dBlockGetFirstAvailableOutputOpts struct {
	ForBlock *NodeMaterialBlock
}

// GetFirstAvailableOutput calls the GetFirstAvailableOutput method on the Rotate2dBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.rotate2dblock#getfirstavailableoutput
func (r *Rotate2dBlock) GetFirstAvailableOutput(opts *Rotate2dBlockGetFirstAvailableOutputOpts) *NodeMaterialConnectionPoint {
	if opts == nil {
		opts = &Rotate2dBlockGetFirstAvailableOutputOpts{}
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

// GetInputByName calls the GetInputByName method on the Rotate2dBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.rotate2dblock#getinputbyname
func (r *Rotate2dBlock) GetInputByName(name string) *NodeMaterialConnectionPoint {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	retVal := r.p.Call("getInputByName", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, r.ctx)
}

// GetOutputByName calls the GetOutputByName method on the Rotate2dBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.rotate2dblock#getoutputbyname
func (r *Rotate2dBlock) GetOutputByName(name string) *NodeMaterialConnectionPoint {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	retVal := r.p.Call("getOutputByName", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, r.ctx)
}

// GetSiblingOutput calls the GetSiblingOutput method on the Rotate2dBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.rotate2dblock#getsiblingoutput
func (r *Rotate2dBlock) GetSiblingOutput(current *NodeMaterialConnectionPoint) *NodeMaterialConnectionPoint {

	args := make([]interface{}, 0, 1+0)

	args = append(args, current.JSObject())

	retVal := r.p.Call("getSiblingOutput", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, r.ctx)
}

// Initialize calls the Initialize method on the Rotate2dBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.rotate2dblock#initialize
func (r *Rotate2dBlock) Initialize(state *NodeMaterialBuildState) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, state.JSObject())

	r.p.Call("initialize", args...)
}

// Rotate2dBlockInitializeDefinesOpts contains optional parameters for Rotate2dBlock.InitializeDefines.
type Rotate2dBlockInitializeDefinesOpts struct {
	UseInstances *bool
}

// InitializeDefines calls the InitializeDefines method on the Rotate2dBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.rotate2dblock#initializedefines
func (r *Rotate2dBlock) InitializeDefines(mesh *AbstractMesh, nodeMaterial *NodeMaterial, defines js.Value, opts *Rotate2dBlockInitializeDefinesOpts) {
	if opts == nil {
		opts = &Rotate2dBlockInitializeDefinesOpts{}
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

// Rotate2dBlockIsReadyOpts contains optional parameters for Rotate2dBlock.IsReady.
type Rotate2dBlockIsReadyOpts struct {
	UseInstances *bool
}

// IsReady calls the IsReady method on the Rotate2dBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.rotate2dblock#isready
func (r *Rotate2dBlock) IsReady(mesh *AbstractMesh, nodeMaterial *NodeMaterial, defines js.Value, opts *Rotate2dBlockIsReadyOpts) bool {
	if opts == nil {
		opts = &Rotate2dBlockIsReadyOpts{}
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

// Rotate2dBlockPrepareDefinesOpts contains optional parameters for Rotate2dBlock.PrepareDefines.
type Rotate2dBlockPrepareDefinesOpts struct {
	UseInstances *bool
}

// PrepareDefines calls the PrepareDefines method on the Rotate2dBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.rotate2dblock#preparedefines
func (r *Rotate2dBlock) PrepareDefines(mesh *AbstractMesh, nodeMaterial *NodeMaterial, defines js.Value, opts *Rotate2dBlockPrepareDefinesOpts) {
	if opts == nil {
		opts = &Rotate2dBlockPrepareDefinesOpts{}
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

// ProvideFallbacks calls the ProvideFallbacks method on the Rotate2dBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.rotate2dblock#providefallbacks
func (r *Rotate2dBlock) ProvideFallbacks(mesh *AbstractMesh, fallbacks *EffectFallbacks) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, mesh.JSObject())
	args = append(args, fallbacks.JSObject())

	r.p.Call("provideFallbacks", args...)
}

// Rotate2dBlockRegisterInputOpts contains optional parameters for Rotate2dBlock.RegisterInput.
type Rotate2dBlockRegisterInputOpts struct {
	IsOptional *bool
	Target     js.Value
}

// RegisterInput calls the RegisterInput method on the Rotate2dBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.rotate2dblock#registerinput
func (r *Rotate2dBlock) RegisterInput(name string, jsType js.Value, opts *Rotate2dBlockRegisterInputOpts) *Rotate2dBlock {
	if opts == nil {
		opts = &Rotate2dBlockRegisterInputOpts{}
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
	return Rotate2dBlockFromJSObject(retVal, r.ctx)
}

// Rotate2dBlockRegisterOutputOpts contains optional parameters for Rotate2dBlock.RegisterOutput.
type Rotate2dBlockRegisterOutputOpts struct {
	Target js.Value
}

// RegisterOutput calls the RegisterOutput method on the Rotate2dBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.rotate2dblock#registeroutput
func (r *Rotate2dBlock) RegisterOutput(name string, jsType js.Value, opts *Rotate2dBlockRegisterOutputOpts) *Rotate2dBlock {
	if opts == nil {
		opts = &Rotate2dBlockRegisterOutputOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, name)
	args = append(args, jsType)

	args = append(args, opts.Target)

	retVal := r.p.Call("registerOutput", args...)
	return Rotate2dBlockFromJSObject(retVal, r.ctx)
}

// ReplaceRepeatableContent calls the ReplaceRepeatableContent method on the Rotate2dBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.rotate2dblock#replacerepeatablecontent
func (r *Rotate2dBlock) ReplaceRepeatableContent(vertexShaderState *NodeMaterialBuildState, fragmentShaderState *NodeMaterialBuildState, mesh *AbstractMesh, defines js.Value) {

	args := make([]interface{}, 0, 4+0)

	args = append(args, vertexShaderState.JSObject())
	args = append(args, fragmentShaderState.JSObject())
	args = append(args, mesh.JSObject())
	args = append(args, defines)

	r.p.Call("replaceRepeatableContent", args...)
}

// Serialize calls the Serialize method on the Rotate2dBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.rotate2dblock#serialize
func (r *Rotate2dBlock) Serialize() interface{} {

	retVal := r.p.Call("serialize")
	return retVal
}

// UpdateUniformsAndSamples calls the UpdateUniformsAndSamples method on the Rotate2dBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.rotate2dblock#updateuniformsandsamples
func (r *Rotate2dBlock) UpdateUniformsAndSamples(state *NodeMaterialBuildState, nodeMaterial *NodeMaterial, defines js.Value, uniformBuffers string) {

	args := make([]interface{}, 0, 4+0)

	args = append(args, state.JSObject())
	args = append(args, nodeMaterial.JSObject())
	args = append(args, defines)
	args = append(args, uniformBuffers)

	r.p.Call("updateUniformsAndSamples", args...)
}

/*

// Angle returns the Angle property of class Rotate2dBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.rotate2dblock#angle
func (r *Rotate2dBlock) Angle(angle *NodeMaterialConnectionPoint) *Rotate2dBlock {
	p := ba.ctx.Get("Rotate2dBlock").New(angle.JSObject())
	return Rotate2dBlockFromJSObject(p, ba.ctx)
}

// SetAngle sets the Angle property of class Rotate2dBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.rotate2dblock#angle
func (r *Rotate2dBlock) SetAngle(angle *NodeMaterialConnectionPoint) *Rotate2dBlock {
	p := ba.ctx.Get("Rotate2dBlock").New(angle.JSObject())
	return Rotate2dBlockFromJSObject(p, ba.ctx)
}

// BuildId returns the BuildId property of class Rotate2dBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.rotate2dblock#buildid
func (r *Rotate2dBlock) BuildId(buildId float64) *Rotate2dBlock {
	p := ba.ctx.Get("Rotate2dBlock").New(buildId)
	return Rotate2dBlockFromJSObject(p, ba.ctx)
}

// SetBuildId sets the BuildId property of class Rotate2dBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.rotate2dblock#buildid
func (r *Rotate2dBlock) SetBuildId(buildId float64) *Rotate2dBlock {
	p := ba.ctx.Get("Rotate2dBlock").New(buildId)
	return Rotate2dBlockFromJSObject(p, ba.ctx)
}

// Comments returns the Comments property of class Rotate2dBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.rotate2dblock#comments
func (r *Rotate2dBlock) Comments(comments string) *Rotate2dBlock {
	p := ba.ctx.Get("Rotate2dBlock").New(comments)
	return Rotate2dBlockFromJSObject(p, ba.ctx)
}

// SetComments sets the Comments property of class Rotate2dBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.rotate2dblock#comments
func (r *Rotate2dBlock) SetComments(comments string) *Rotate2dBlock {
	p := ba.ctx.Get("Rotate2dBlock").New(comments)
	return Rotate2dBlockFromJSObject(p, ba.ctx)
}

// Input returns the Input property of class Rotate2dBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.rotate2dblock#input
func (r *Rotate2dBlock) Input(input *NodeMaterialConnectionPoint) *Rotate2dBlock {
	p := ba.ctx.Get("Rotate2dBlock").New(input.JSObject())
	return Rotate2dBlockFromJSObject(p, ba.ctx)
}

// SetInput sets the Input property of class Rotate2dBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.rotate2dblock#input
func (r *Rotate2dBlock) SetInput(input *NodeMaterialConnectionPoint) *Rotate2dBlock {
	p := ba.ctx.Get("Rotate2dBlock").New(input.JSObject())
	return Rotate2dBlockFromJSObject(p, ba.ctx)
}

// Inputs returns the Inputs property of class Rotate2dBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.rotate2dblock#inputs
func (r *Rotate2dBlock) Inputs(inputs *NodeMaterialConnectionPoint) *Rotate2dBlock {
	p := ba.ctx.Get("Rotate2dBlock").New(inputs.JSObject())
	return Rotate2dBlockFromJSObject(p, ba.ctx)
}

// SetInputs sets the Inputs property of class Rotate2dBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.rotate2dblock#inputs
func (r *Rotate2dBlock) SetInputs(inputs *NodeMaterialConnectionPoint) *Rotate2dBlock {
	p := ba.ctx.Get("Rotate2dBlock").New(inputs.JSObject())
	return Rotate2dBlockFromJSObject(p, ba.ctx)
}

// IsFinalMerger returns the IsFinalMerger property of class Rotate2dBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.rotate2dblock#isfinalmerger
func (r *Rotate2dBlock) IsFinalMerger(isFinalMerger bool) *Rotate2dBlock {
	p := ba.ctx.Get("Rotate2dBlock").New(isFinalMerger)
	return Rotate2dBlockFromJSObject(p, ba.ctx)
}

// SetIsFinalMerger sets the IsFinalMerger property of class Rotate2dBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.rotate2dblock#isfinalmerger
func (r *Rotate2dBlock) SetIsFinalMerger(isFinalMerger bool) *Rotate2dBlock {
	p := ba.ctx.Get("Rotate2dBlock").New(isFinalMerger)
	return Rotate2dBlockFromJSObject(p, ba.ctx)
}

// IsInput returns the IsInput property of class Rotate2dBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.rotate2dblock#isinput
func (r *Rotate2dBlock) IsInput(isInput bool) *Rotate2dBlock {
	p := ba.ctx.Get("Rotate2dBlock").New(isInput)
	return Rotate2dBlockFromJSObject(p, ba.ctx)
}

// SetIsInput sets the IsInput property of class Rotate2dBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.rotate2dblock#isinput
func (r *Rotate2dBlock) SetIsInput(isInput bool) *Rotate2dBlock {
	p := ba.ctx.Get("Rotate2dBlock").New(isInput)
	return Rotate2dBlockFromJSObject(p, ba.ctx)
}

// IsUnique returns the IsUnique property of class Rotate2dBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.rotate2dblock#isunique
func (r *Rotate2dBlock) IsUnique(isUnique bool) *Rotate2dBlock {
	p := ba.ctx.Get("Rotate2dBlock").New(isUnique)
	return Rotate2dBlockFromJSObject(p, ba.ctx)
}

// SetIsUnique sets the IsUnique property of class Rotate2dBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.rotate2dblock#isunique
func (r *Rotate2dBlock) SetIsUnique(isUnique bool) *Rotate2dBlock {
	p := ba.ctx.Get("Rotate2dBlock").New(isUnique)
	return Rotate2dBlockFromJSObject(p, ba.ctx)
}

// Name returns the Name property of class Rotate2dBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.rotate2dblock#name
func (r *Rotate2dBlock) Name(name string) *Rotate2dBlock {
	p := ba.ctx.Get("Rotate2dBlock").New(name)
	return Rotate2dBlockFromJSObject(p, ba.ctx)
}

// SetName sets the Name property of class Rotate2dBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.rotate2dblock#name
func (r *Rotate2dBlock) SetName(name string) *Rotate2dBlock {
	p := ba.ctx.Get("Rotate2dBlock").New(name)
	return Rotate2dBlockFromJSObject(p, ba.ctx)
}

// Output returns the Output property of class Rotate2dBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.rotate2dblock#output
func (r *Rotate2dBlock) Output(output *NodeMaterialConnectionPoint) *Rotate2dBlock {
	p := ba.ctx.Get("Rotate2dBlock").New(output.JSObject())
	return Rotate2dBlockFromJSObject(p, ba.ctx)
}

// SetOutput sets the Output property of class Rotate2dBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.rotate2dblock#output
func (r *Rotate2dBlock) SetOutput(output *NodeMaterialConnectionPoint) *Rotate2dBlock {
	p := ba.ctx.Get("Rotate2dBlock").New(output.JSObject())
	return Rotate2dBlockFromJSObject(p, ba.ctx)
}

// Outputs returns the Outputs property of class Rotate2dBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.rotate2dblock#outputs
func (r *Rotate2dBlock) Outputs(outputs *NodeMaterialConnectionPoint) *Rotate2dBlock {
	p := ba.ctx.Get("Rotate2dBlock").New(outputs.JSObject())
	return Rotate2dBlockFromJSObject(p, ba.ctx)
}

// SetOutputs sets the Outputs property of class Rotate2dBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.rotate2dblock#outputs
func (r *Rotate2dBlock) SetOutputs(outputs *NodeMaterialConnectionPoint) *Rotate2dBlock {
	p := ba.ctx.Get("Rotate2dBlock").New(outputs.JSObject())
	return Rotate2dBlockFromJSObject(p, ba.ctx)
}

// Target returns the Target property of class Rotate2dBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.rotate2dblock#target
func (r *Rotate2dBlock) Target(target js.Value) *Rotate2dBlock {
	p := ba.ctx.Get("Rotate2dBlock").New(target)
	return Rotate2dBlockFromJSObject(p, ba.ctx)
}

// SetTarget sets the Target property of class Rotate2dBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.rotate2dblock#target
func (r *Rotate2dBlock) SetTarget(target js.Value) *Rotate2dBlock {
	p := ba.ctx.Get("Rotate2dBlock").New(target)
	return Rotate2dBlockFromJSObject(p, ba.ctx)
}

// UniqueId returns the UniqueId property of class Rotate2dBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.rotate2dblock#uniqueid
func (r *Rotate2dBlock) UniqueId(uniqueId float64) *Rotate2dBlock {
	p := ba.ctx.Get("Rotate2dBlock").New(uniqueId)
	return Rotate2dBlockFromJSObject(p, ba.ctx)
}

// SetUniqueId sets the UniqueId property of class Rotate2dBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.rotate2dblock#uniqueid
func (r *Rotate2dBlock) SetUniqueId(uniqueId float64) *Rotate2dBlock {
	p := ba.ctx.Get("Rotate2dBlock").New(uniqueId)
	return Rotate2dBlockFromJSObject(p, ba.ctx)
}

*/
