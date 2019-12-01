// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// LerpBlock represents a babylon.js LerpBlock.
// Block used to lerp between 2 values
type LerpBlock struct {
	*NodeMaterialBlock
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (l *LerpBlock) JSObject() js.Value { return l.p }

// LerpBlock returns a LerpBlock JavaScript class.
func (ba *Babylon) LerpBlock() *LerpBlock {
	p := ba.ctx.Get("LerpBlock")
	return LerpBlockFromJSObject(p, ba.ctx)
}

// LerpBlockFromJSObject returns a wrapped LerpBlock JavaScript class.
func LerpBlockFromJSObject(p js.Value, ctx js.Value) *LerpBlock {
	return &LerpBlock{NodeMaterialBlock: NodeMaterialBlockFromJSObject(p, ctx), ctx: ctx}
}

// LerpBlockArrayToJSArray returns a JavaScript Array for the wrapped array.
func LerpBlockArrayToJSArray(array []*LerpBlock) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewLerpBlock returns a new LerpBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.lerpblock
func (ba *Babylon) NewLerpBlock(name string) *LerpBlock {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	p := ba.ctx.Get("LerpBlock").New(args...)
	return LerpBlockFromJSObject(p, ba.ctx)
}

// AutoConfigure calls the AutoConfigure method on the LerpBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.lerpblock#autoconfigure
func (l *LerpBlock) AutoConfigure(material *NodeMaterial) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, material.JSObject())

	l.p.Call("autoConfigure", args...)
}

// LerpBlockBindOpts contains optional parameters for LerpBlock.Bind.
type LerpBlockBindOpts struct {
	Mesh *Mesh
}

// Bind calls the Bind method on the LerpBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.lerpblock#bind
func (l *LerpBlock) Bind(effect *Effect, nodeMaterial *NodeMaterial, opts *LerpBlockBindOpts) {
	if opts == nil {
		opts = &LerpBlockBindOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, effect.JSObject())
	args = append(args, nodeMaterial.JSObject())

	if opts.Mesh == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Mesh.JSObject())
	}

	l.p.Call("bind", args...)
}

// Build calls the Build method on the LerpBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.lerpblock#build
func (l *LerpBlock) Build(state *NodeMaterialBuildState, activeBlocks *NodeMaterialBlock) bool {

	args := make([]interface{}, 0, 2+0)

	args = append(args, state.JSObject())
	args = append(args, activeBlocks.JSObject())

	retVal := l.p.Call("build", args...)
	return retVal.Bool()
}

// LerpBlockCloneOpts contains optional parameters for LerpBlock.Clone.
type LerpBlockCloneOpts struct {
	RootUrl *string
}

// Clone calls the Clone method on the LerpBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.lerpblock#clone
func (l *LerpBlock) Clone(scene *Scene, opts *LerpBlockCloneOpts) *NodeMaterialBlock {
	if opts == nil {
		opts = &LerpBlockCloneOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, scene.JSObject())

	if opts.RootUrl == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.RootUrl)
	}

	retVal := l.p.Call("clone", args...)
	return NodeMaterialBlockFromJSObject(retVal, l.ctx)
}

// LerpBlockConnectToOpts contains optional parameters for LerpBlock.ConnectTo.
type LerpBlockConnectToOpts struct {
	Options map[string]interface{}
}

// ConnectTo calls the ConnectTo method on the LerpBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.lerpblock#connectto
func (l *LerpBlock) ConnectTo(other *NodeMaterialBlock, opts *LerpBlockConnectToOpts) *LerpBlock {
	if opts == nil {
		opts = &LerpBlockConnectToOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, other.JSObject())

	if opts.Options == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Options)
	}

	retVal := l.p.Call("connectTo", args...)
	return LerpBlockFromJSObject(retVal, l.ctx)
}

// Dispose calls the Dispose method on the LerpBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.lerpblock#dispose
func (l *LerpBlock) Dispose() {

	l.p.Call("dispose")
}

// GetClassName calls the GetClassName method on the LerpBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.lerpblock#getclassname
func (l *LerpBlock) GetClassName() string {

	retVal := l.p.Call("getClassName")
	return retVal.String()
}

// LerpBlockGetFirstAvailableInputOpts contains optional parameters for LerpBlock.GetFirstAvailableInput.
type LerpBlockGetFirstAvailableInputOpts struct {
	ForOutput *NodeMaterialConnectionPoint
}

// GetFirstAvailableInput calls the GetFirstAvailableInput method on the LerpBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.lerpblock#getfirstavailableinput
func (l *LerpBlock) GetFirstAvailableInput(opts *LerpBlockGetFirstAvailableInputOpts) *NodeMaterialConnectionPoint {
	if opts == nil {
		opts = &LerpBlockGetFirstAvailableInputOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.ForOutput == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.ForOutput.JSObject())
	}

	retVal := l.p.Call("getFirstAvailableInput", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, l.ctx)
}

// LerpBlockGetFirstAvailableOutputOpts contains optional parameters for LerpBlock.GetFirstAvailableOutput.
type LerpBlockGetFirstAvailableOutputOpts struct {
	ForBlock *NodeMaterialBlock
}

// GetFirstAvailableOutput calls the GetFirstAvailableOutput method on the LerpBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.lerpblock#getfirstavailableoutput
func (l *LerpBlock) GetFirstAvailableOutput(opts *LerpBlockGetFirstAvailableOutputOpts) *NodeMaterialConnectionPoint {
	if opts == nil {
		opts = &LerpBlockGetFirstAvailableOutputOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.ForBlock == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.ForBlock.JSObject())
	}

	retVal := l.p.Call("getFirstAvailableOutput", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, l.ctx)
}

// GetInputByName calls the GetInputByName method on the LerpBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.lerpblock#getinputbyname
func (l *LerpBlock) GetInputByName(name string) *NodeMaterialConnectionPoint {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	retVal := l.p.Call("getInputByName", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, l.ctx)
}

// GetOutputByName calls the GetOutputByName method on the LerpBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.lerpblock#getoutputbyname
func (l *LerpBlock) GetOutputByName(name string) *NodeMaterialConnectionPoint {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	retVal := l.p.Call("getOutputByName", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, l.ctx)
}

// GetSiblingOutput calls the GetSiblingOutput method on the LerpBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.lerpblock#getsiblingoutput
func (l *LerpBlock) GetSiblingOutput(current *NodeMaterialConnectionPoint) *NodeMaterialConnectionPoint {

	args := make([]interface{}, 0, 1+0)

	args = append(args, current.JSObject())

	retVal := l.p.Call("getSiblingOutput", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, l.ctx)
}

// Initialize calls the Initialize method on the LerpBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.lerpblock#initialize
func (l *LerpBlock) Initialize(state *NodeMaterialBuildState) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, state.JSObject())

	l.p.Call("initialize", args...)
}

// LerpBlockInitializeDefinesOpts contains optional parameters for LerpBlock.InitializeDefines.
type LerpBlockInitializeDefinesOpts struct {
	UseInstances *bool
}

// InitializeDefines calls the InitializeDefines method on the LerpBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.lerpblock#initializedefines
func (l *LerpBlock) InitializeDefines(mesh *AbstractMesh, nodeMaterial *NodeMaterial, defines js.Value, opts *LerpBlockInitializeDefinesOpts) {
	if opts == nil {
		opts = &LerpBlockInitializeDefinesOpts{}
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

	l.p.Call("initializeDefines", args...)
}

// LerpBlockIsReadyOpts contains optional parameters for LerpBlock.IsReady.
type LerpBlockIsReadyOpts struct {
	UseInstances *bool
}

// IsReady calls the IsReady method on the LerpBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.lerpblock#isready
func (l *LerpBlock) IsReady(mesh *AbstractMesh, nodeMaterial *NodeMaterial, defines js.Value, opts *LerpBlockIsReadyOpts) bool {
	if opts == nil {
		opts = &LerpBlockIsReadyOpts{}
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

	retVal := l.p.Call("isReady", args...)
	return retVal.Bool()
}

// LerpBlockPrepareDefinesOpts contains optional parameters for LerpBlock.PrepareDefines.
type LerpBlockPrepareDefinesOpts struct {
	UseInstances *bool
}

// PrepareDefines calls the PrepareDefines method on the LerpBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.lerpblock#preparedefines
func (l *LerpBlock) PrepareDefines(mesh *AbstractMesh, nodeMaterial *NodeMaterial, defines js.Value, opts *LerpBlockPrepareDefinesOpts) {
	if opts == nil {
		opts = &LerpBlockPrepareDefinesOpts{}
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

	l.p.Call("prepareDefines", args...)
}

// ProvideFallbacks calls the ProvideFallbacks method on the LerpBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.lerpblock#providefallbacks
func (l *LerpBlock) ProvideFallbacks(mesh *AbstractMesh, fallbacks *EffectFallbacks) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, mesh.JSObject())
	args = append(args, fallbacks.JSObject())

	l.p.Call("provideFallbacks", args...)
}

// LerpBlockRegisterInputOpts contains optional parameters for LerpBlock.RegisterInput.
type LerpBlockRegisterInputOpts struct {
	IsOptional *bool
	Target     js.Value
}

// RegisterInput calls the RegisterInput method on the LerpBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.lerpblock#registerinput
func (l *LerpBlock) RegisterInput(name string, jsType js.Value, opts *LerpBlockRegisterInputOpts) *LerpBlock {
	if opts == nil {
		opts = &LerpBlockRegisterInputOpts{}
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

	retVal := l.p.Call("registerInput", args...)
	return LerpBlockFromJSObject(retVal, l.ctx)
}

// LerpBlockRegisterOutputOpts contains optional parameters for LerpBlock.RegisterOutput.
type LerpBlockRegisterOutputOpts struct {
	Target js.Value
}

// RegisterOutput calls the RegisterOutput method on the LerpBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.lerpblock#registeroutput
func (l *LerpBlock) RegisterOutput(name string, jsType js.Value, opts *LerpBlockRegisterOutputOpts) *LerpBlock {
	if opts == nil {
		opts = &LerpBlockRegisterOutputOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, name)
	args = append(args, jsType)

	args = append(args, opts.Target)

	retVal := l.p.Call("registerOutput", args...)
	return LerpBlockFromJSObject(retVal, l.ctx)
}

// ReplaceRepeatableContent calls the ReplaceRepeatableContent method on the LerpBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.lerpblock#replacerepeatablecontent
func (l *LerpBlock) ReplaceRepeatableContent(vertexShaderState *NodeMaterialBuildState, fragmentShaderState *NodeMaterialBuildState, mesh *AbstractMesh, defines js.Value) {

	args := make([]interface{}, 0, 4+0)

	args = append(args, vertexShaderState.JSObject())
	args = append(args, fragmentShaderState.JSObject())
	args = append(args, mesh.JSObject())
	args = append(args, defines)

	l.p.Call("replaceRepeatableContent", args...)
}

// Serialize calls the Serialize method on the LerpBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.lerpblock#serialize
func (l *LerpBlock) Serialize() interface{} {

	retVal := l.p.Call("serialize")
	return retVal
}

// UpdateUniformsAndSamples calls the UpdateUniformsAndSamples method on the LerpBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.lerpblock#updateuniformsandsamples
func (l *LerpBlock) UpdateUniformsAndSamples(state *NodeMaterialBuildState, nodeMaterial *NodeMaterial, defines js.Value, uniformBuffers string) {

	args := make([]interface{}, 0, 4+0)

	args = append(args, state.JSObject())
	args = append(args, nodeMaterial.JSObject())
	args = append(args, defines)
	args = append(args, uniformBuffers)

	l.p.Call("updateUniformsAndSamples", args...)
}

/*

// BuildId returns the BuildId property of class LerpBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.lerpblock#buildid
func (l *LerpBlock) BuildId(buildId float64) *LerpBlock {
	p := ba.ctx.Get("LerpBlock").New(buildId)
	return LerpBlockFromJSObject(p, ba.ctx)
}

// SetBuildId sets the BuildId property of class LerpBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.lerpblock#buildid
func (l *LerpBlock) SetBuildId(buildId float64) *LerpBlock {
	p := ba.ctx.Get("LerpBlock").New(buildId)
	return LerpBlockFromJSObject(p, ba.ctx)
}

// Comments returns the Comments property of class LerpBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.lerpblock#comments
func (l *LerpBlock) Comments(comments string) *LerpBlock {
	p := ba.ctx.Get("LerpBlock").New(comments)
	return LerpBlockFromJSObject(p, ba.ctx)
}

// SetComments sets the Comments property of class LerpBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.lerpblock#comments
func (l *LerpBlock) SetComments(comments string) *LerpBlock {
	p := ba.ctx.Get("LerpBlock").New(comments)
	return LerpBlockFromJSObject(p, ba.ctx)
}

// Gradient returns the Gradient property of class LerpBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.lerpblock#gradient
func (l *LerpBlock) Gradient(gradient *NodeMaterialConnectionPoint) *LerpBlock {
	p := ba.ctx.Get("LerpBlock").New(gradient.JSObject())
	return LerpBlockFromJSObject(p, ba.ctx)
}

// SetGradient sets the Gradient property of class LerpBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.lerpblock#gradient
func (l *LerpBlock) SetGradient(gradient *NodeMaterialConnectionPoint) *LerpBlock {
	p := ba.ctx.Get("LerpBlock").New(gradient.JSObject())
	return LerpBlockFromJSObject(p, ba.ctx)
}

// Inputs returns the Inputs property of class LerpBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.lerpblock#inputs
func (l *LerpBlock) Inputs(inputs *NodeMaterialConnectionPoint) *LerpBlock {
	p := ba.ctx.Get("LerpBlock").New(inputs.JSObject())
	return LerpBlockFromJSObject(p, ba.ctx)
}

// SetInputs sets the Inputs property of class LerpBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.lerpblock#inputs
func (l *LerpBlock) SetInputs(inputs *NodeMaterialConnectionPoint) *LerpBlock {
	p := ba.ctx.Get("LerpBlock").New(inputs.JSObject())
	return LerpBlockFromJSObject(p, ba.ctx)
}

// IsFinalMerger returns the IsFinalMerger property of class LerpBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.lerpblock#isfinalmerger
func (l *LerpBlock) IsFinalMerger(isFinalMerger bool) *LerpBlock {
	p := ba.ctx.Get("LerpBlock").New(isFinalMerger)
	return LerpBlockFromJSObject(p, ba.ctx)
}

// SetIsFinalMerger sets the IsFinalMerger property of class LerpBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.lerpblock#isfinalmerger
func (l *LerpBlock) SetIsFinalMerger(isFinalMerger bool) *LerpBlock {
	p := ba.ctx.Get("LerpBlock").New(isFinalMerger)
	return LerpBlockFromJSObject(p, ba.ctx)
}

// IsInput returns the IsInput property of class LerpBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.lerpblock#isinput
func (l *LerpBlock) IsInput(isInput bool) *LerpBlock {
	p := ba.ctx.Get("LerpBlock").New(isInput)
	return LerpBlockFromJSObject(p, ba.ctx)
}

// SetIsInput sets the IsInput property of class LerpBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.lerpblock#isinput
func (l *LerpBlock) SetIsInput(isInput bool) *LerpBlock {
	p := ba.ctx.Get("LerpBlock").New(isInput)
	return LerpBlockFromJSObject(p, ba.ctx)
}

// IsUnique returns the IsUnique property of class LerpBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.lerpblock#isunique
func (l *LerpBlock) IsUnique(isUnique bool) *LerpBlock {
	p := ba.ctx.Get("LerpBlock").New(isUnique)
	return LerpBlockFromJSObject(p, ba.ctx)
}

// SetIsUnique sets the IsUnique property of class LerpBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.lerpblock#isunique
func (l *LerpBlock) SetIsUnique(isUnique bool) *LerpBlock {
	p := ba.ctx.Get("LerpBlock").New(isUnique)
	return LerpBlockFromJSObject(p, ba.ctx)
}

// Left returns the Left property of class LerpBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.lerpblock#left
func (l *LerpBlock) Left(left *NodeMaterialConnectionPoint) *LerpBlock {
	p := ba.ctx.Get("LerpBlock").New(left.JSObject())
	return LerpBlockFromJSObject(p, ba.ctx)
}

// SetLeft sets the Left property of class LerpBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.lerpblock#left
func (l *LerpBlock) SetLeft(left *NodeMaterialConnectionPoint) *LerpBlock {
	p := ba.ctx.Get("LerpBlock").New(left.JSObject())
	return LerpBlockFromJSObject(p, ba.ctx)
}

// Name returns the Name property of class LerpBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.lerpblock#name
func (l *LerpBlock) Name(name string) *LerpBlock {
	p := ba.ctx.Get("LerpBlock").New(name)
	return LerpBlockFromJSObject(p, ba.ctx)
}

// SetName sets the Name property of class LerpBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.lerpblock#name
func (l *LerpBlock) SetName(name string) *LerpBlock {
	p := ba.ctx.Get("LerpBlock").New(name)
	return LerpBlockFromJSObject(p, ba.ctx)
}

// Output returns the Output property of class LerpBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.lerpblock#output
func (l *LerpBlock) Output(output *NodeMaterialConnectionPoint) *LerpBlock {
	p := ba.ctx.Get("LerpBlock").New(output.JSObject())
	return LerpBlockFromJSObject(p, ba.ctx)
}

// SetOutput sets the Output property of class LerpBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.lerpblock#output
func (l *LerpBlock) SetOutput(output *NodeMaterialConnectionPoint) *LerpBlock {
	p := ba.ctx.Get("LerpBlock").New(output.JSObject())
	return LerpBlockFromJSObject(p, ba.ctx)
}

// Outputs returns the Outputs property of class LerpBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.lerpblock#outputs
func (l *LerpBlock) Outputs(outputs *NodeMaterialConnectionPoint) *LerpBlock {
	p := ba.ctx.Get("LerpBlock").New(outputs.JSObject())
	return LerpBlockFromJSObject(p, ba.ctx)
}

// SetOutputs sets the Outputs property of class LerpBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.lerpblock#outputs
func (l *LerpBlock) SetOutputs(outputs *NodeMaterialConnectionPoint) *LerpBlock {
	p := ba.ctx.Get("LerpBlock").New(outputs.JSObject())
	return LerpBlockFromJSObject(p, ba.ctx)
}

// Right returns the Right property of class LerpBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.lerpblock#right
func (l *LerpBlock) Right(right *NodeMaterialConnectionPoint) *LerpBlock {
	p := ba.ctx.Get("LerpBlock").New(right.JSObject())
	return LerpBlockFromJSObject(p, ba.ctx)
}

// SetRight sets the Right property of class LerpBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.lerpblock#right
func (l *LerpBlock) SetRight(right *NodeMaterialConnectionPoint) *LerpBlock {
	p := ba.ctx.Get("LerpBlock").New(right.JSObject())
	return LerpBlockFromJSObject(p, ba.ctx)
}

// Target returns the Target property of class LerpBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.lerpblock#target
func (l *LerpBlock) Target(target js.Value) *LerpBlock {
	p := ba.ctx.Get("LerpBlock").New(target)
	return LerpBlockFromJSObject(p, ba.ctx)
}

// SetTarget sets the Target property of class LerpBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.lerpblock#target
func (l *LerpBlock) SetTarget(target js.Value) *LerpBlock {
	p := ba.ctx.Get("LerpBlock").New(target)
	return LerpBlockFromJSObject(p, ba.ctx)
}

// UniqueId returns the UniqueId property of class LerpBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.lerpblock#uniqueid
func (l *LerpBlock) UniqueId(uniqueId float64) *LerpBlock {
	p := ba.ctx.Get("LerpBlock").New(uniqueId)
	return LerpBlockFromJSObject(p, ba.ctx)
}

// SetUniqueId sets the UniqueId property of class LerpBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.lerpblock#uniqueid
func (l *LerpBlock) SetUniqueId(uniqueId float64) *LerpBlock {
	p := ba.ctx.Get("LerpBlock").New(uniqueId)
	return LerpBlockFromJSObject(p, ba.ctx)
}

*/
