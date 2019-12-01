// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ScaleBlock represents a babylon.js ScaleBlock.
// Block used to scale a vector by a float
type ScaleBlock struct {
	*NodeMaterialBlock
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (s *ScaleBlock) JSObject() js.Value { return s.p }

// ScaleBlock returns a ScaleBlock JavaScript class.
func (ba *Babylon) ScaleBlock() *ScaleBlock {
	p := ba.ctx.Get("ScaleBlock")
	return ScaleBlockFromJSObject(p, ba.ctx)
}

// ScaleBlockFromJSObject returns a wrapped ScaleBlock JavaScript class.
func ScaleBlockFromJSObject(p js.Value, ctx js.Value) *ScaleBlock {
	return &ScaleBlock{NodeMaterialBlock: NodeMaterialBlockFromJSObject(p, ctx), ctx: ctx}
}

// ScaleBlockArrayToJSArray returns a JavaScript Array for the wrapped array.
func ScaleBlockArrayToJSArray(array []*ScaleBlock) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewScaleBlock returns a new ScaleBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.scaleblock
func (ba *Babylon) NewScaleBlock(name string) *ScaleBlock {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	p := ba.ctx.Get("ScaleBlock").New(args...)
	return ScaleBlockFromJSObject(p, ba.ctx)
}

// AutoConfigure calls the AutoConfigure method on the ScaleBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.scaleblock#autoconfigure
func (s *ScaleBlock) AutoConfigure(material *NodeMaterial) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, material.JSObject())

	s.p.Call("autoConfigure", args...)
}

// ScaleBlockBindOpts contains optional parameters for ScaleBlock.Bind.
type ScaleBlockBindOpts struct {
	Mesh *Mesh
}

// Bind calls the Bind method on the ScaleBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.scaleblock#bind
func (s *ScaleBlock) Bind(effect *Effect, nodeMaterial *NodeMaterial, opts *ScaleBlockBindOpts) {
	if opts == nil {
		opts = &ScaleBlockBindOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, effect.JSObject())
	args = append(args, nodeMaterial.JSObject())

	if opts.Mesh == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Mesh.JSObject())
	}

	s.p.Call("bind", args...)
}

// Build calls the Build method on the ScaleBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.scaleblock#build
func (s *ScaleBlock) Build(state *NodeMaterialBuildState, activeBlocks *NodeMaterialBlock) bool {

	args := make([]interface{}, 0, 2+0)

	args = append(args, state.JSObject())
	args = append(args, activeBlocks.JSObject())

	retVal := s.p.Call("build", args...)
	return retVal.Bool()
}

// ScaleBlockCloneOpts contains optional parameters for ScaleBlock.Clone.
type ScaleBlockCloneOpts struct {
	RootUrl *string
}

// Clone calls the Clone method on the ScaleBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.scaleblock#clone
func (s *ScaleBlock) Clone(scene *Scene, opts *ScaleBlockCloneOpts) *NodeMaterialBlock {
	if opts == nil {
		opts = &ScaleBlockCloneOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, scene.JSObject())

	if opts.RootUrl == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.RootUrl)
	}

	retVal := s.p.Call("clone", args...)
	return NodeMaterialBlockFromJSObject(retVal, s.ctx)
}

// ScaleBlockConnectToOpts contains optional parameters for ScaleBlock.ConnectTo.
type ScaleBlockConnectToOpts struct {
	Options map[string]interface{}
}

// ConnectTo calls the ConnectTo method on the ScaleBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.scaleblock#connectto
func (s *ScaleBlock) ConnectTo(other *NodeMaterialBlock, opts *ScaleBlockConnectToOpts) *ScaleBlock {
	if opts == nil {
		opts = &ScaleBlockConnectToOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, other.JSObject())

	if opts.Options == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Options)
	}

	retVal := s.p.Call("connectTo", args...)
	return ScaleBlockFromJSObject(retVal, s.ctx)
}

// Dispose calls the Dispose method on the ScaleBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.scaleblock#dispose
func (s *ScaleBlock) Dispose() {

	s.p.Call("dispose")
}

// GetClassName calls the GetClassName method on the ScaleBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.scaleblock#getclassname
func (s *ScaleBlock) GetClassName() string {

	retVal := s.p.Call("getClassName")
	return retVal.String()
}

// ScaleBlockGetFirstAvailableInputOpts contains optional parameters for ScaleBlock.GetFirstAvailableInput.
type ScaleBlockGetFirstAvailableInputOpts struct {
	ForOutput *NodeMaterialConnectionPoint
}

// GetFirstAvailableInput calls the GetFirstAvailableInput method on the ScaleBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.scaleblock#getfirstavailableinput
func (s *ScaleBlock) GetFirstAvailableInput(opts *ScaleBlockGetFirstAvailableInputOpts) *NodeMaterialConnectionPoint {
	if opts == nil {
		opts = &ScaleBlockGetFirstAvailableInputOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.ForOutput == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.ForOutput.JSObject())
	}

	retVal := s.p.Call("getFirstAvailableInput", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, s.ctx)
}

// ScaleBlockGetFirstAvailableOutputOpts contains optional parameters for ScaleBlock.GetFirstAvailableOutput.
type ScaleBlockGetFirstAvailableOutputOpts struct {
	ForBlock *NodeMaterialBlock
}

// GetFirstAvailableOutput calls the GetFirstAvailableOutput method on the ScaleBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.scaleblock#getfirstavailableoutput
func (s *ScaleBlock) GetFirstAvailableOutput(opts *ScaleBlockGetFirstAvailableOutputOpts) *NodeMaterialConnectionPoint {
	if opts == nil {
		opts = &ScaleBlockGetFirstAvailableOutputOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.ForBlock == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.ForBlock.JSObject())
	}

	retVal := s.p.Call("getFirstAvailableOutput", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, s.ctx)
}

// GetInputByName calls the GetInputByName method on the ScaleBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.scaleblock#getinputbyname
func (s *ScaleBlock) GetInputByName(name string) *NodeMaterialConnectionPoint {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	retVal := s.p.Call("getInputByName", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, s.ctx)
}

// GetOutputByName calls the GetOutputByName method on the ScaleBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.scaleblock#getoutputbyname
func (s *ScaleBlock) GetOutputByName(name string) *NodeMaterialConnectionPoint {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	retVal := s.p.Call("getOutputByName", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, s.ctx)
}

// GetSiblingOutput calls the GetSiblingOutput method on the ScaleBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.scaleblock#getsiblingoutput
func (s *ScaleBlock) GetSiblingOutput(current *NodeMaterialConnectionPoint) *NodeMaterialConnectionPoint {

	args := make([]interface{}, 0, 1+0)

	args = append(args, current.JSObject())

	retVal := s.p.Call("getSiblingOutput", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, s.ctx)
}

// Initialize calls the Initialize method on the ScaleBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.scaleblock#initialize
func (s *ScaleBlock) Initialize(state *NodeMaterialBuildState) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, state.JSObject())

	s.p.Call("initialize", args...)
}

// ScaleBlockInitializeDefinesOpts contains optional parameters for ScaleBlock.InitializeDefines.
type ScaleBlockInitializeDefinesOpts struct {
	UseInstances *bool
}

// InitializeDefines calls the InitializeDefines method on the ScaleBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.scaleblock#initializedefines
func (s *ScaleBlock) InitializeDefines(mesh *AbstractMesh, nodeMaterial *NodeMaterial, defines js.Value, opts *ScaleBlockInitializeDefinesOpts) {
	if opts == nil {
		opts = &ScaleBlockInitializeDefinesOpts{}
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

	s.p.Call("initializeDefines", args...)
}

// ScaleBlockIsReadyOpts contains optional parameters for ScaleBlock.IsReady.
type ScaleBlockIsReadyOpts struct {
	UseInstances *bool
}

// IsReady calls the IsReady method on the ScaleBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.scaleblock#isready
func (s *ScaleBlock) IsReady(mesh *AbstractMesh, nodeMaterial *NodeMaterial, defines js.Value, opts *ScaleBlockIsReadyOpts) bool {
	if opts == nil {
		opts = &ScaleBlockIsReadyOpts{}
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

	retVal := s.p.Call("isReady", args...)
	return retVal.Bool()
}

// ScaleBlockPrepareDefinesOpts contains optional parameters for ScaleBlock.PrepareDefines.
type ScaleBlockPrepareDefinesOpts struct {
	UseInstances *bool
}

// PrepareDefines calls the PrepareDefines method on the ScaleBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.scaleblock#preparedefines
func (s *ScaleBlock) PrepareDefines(mesh *AbstractMesh, nodeMaterial *NodeMaterial, defines js.Value, opts *ScaleBlockPrepareDefinesOpts) {
	if opts == nil {
		opts = &ScaleBlockPrepareDefinesOpts{}
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

	s.p.Call("prepareDefines", args...)
}

// ProvideFallbacks calls the ProvideFallbacks method on the ScaleBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.scaleblock#providefallbacks
func (s *ScaleBlock) ProvideFallbacks(mesh *AbstractMesh, fallbacks *EffectFallbacks) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, mesh.JSObject())
	args = append(args, fallbacks.JSObject())

	s.p.Call("provideFallbacks", args...)
}

// ScaleBlockRegisterInputOpts contains optional parameters for ScaleBlock.RegisterInput.
type ScaleBlockRegisterInputOpts struct {
	IsOptional *bool
	Target     js.Value
}

// RegisterInput calls the RegisterInput method on the ScaleBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.scaleblock#registerinput
func (s *ScaleBlock) RegisterInput(name string, jsType js.Value, opts *ScaleBlockRegisterInputOpts) *ScaleBlock {
	if opts == nil {
		opts = &ScaleBlockRegisterInputOpts{}
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

	retVal := s.p.Call("registerInput", args...)
	return ScaleBlockFromJSObject(retVal, s.ctx)
}

// ScaleBlockRegisterOutputOpts contains optional parameters for ScaleBlock.RegisterOutput.
type ScaleBlockRegisterOutputOpts struct {
	Target js.Value
}

// RegisterOutput calls the RegisterOutput method on the ScaleBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.scaleblock#registeroutput
func (s *ScaleBlock) RegisterOutput(name string, jsType js.Value, opts *ScaleBlockRegisterOutputOpts) *ScaleBlock {
	if opts == nil {
		opts = &ScaleBlockRegisterOutputOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, name)
	args = append(args, jsType)

	args = append(args, opts.Target)

	retVal := s.p.Call("registerOutput", args...)
	return ScaleBlockFromJSObject(retVal, s.ctx)
}

// ReplaceRepeatableContent calls the ReplaceRepeatableContent method on the ScaleBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.scaleblock#replacerepeatablecontent
func (s *ScaleBlock) ReplaceRepeatableContent(vertexShaderState *NodeMaterialBuildState, fragmentShaderState *NodeMaterialBuildState, mesh *AbstractMesh, defines js.Value) {

	args := make([]interface{}, 0, 4+0)

	args = append(args, vertexShaderState.JSObject())
	args = append(args, fragmentShaderState.JSObject())
	args = append(args, mesh.JSObject())
	args = append(args, defines)

	s.p.Call("replaceRepeatableContent", args...)
}

// Serialize calls the Serialize method on the ScaleBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.scaleblock#serialize
func (s *ScaleBlock) Serialize() interface{} {

	retVal := s.p.Call("serialize")
	return retVal
}

// UpdateUniformsAndSamples calls the UpdateUniformsAndSamples method on the ScaleBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.scaleblock#updateuniformsandsamples
func (s *ScaleBlock) UpdateUniformsAndSamples(state *NodeMaterialBuildState, nodeMaterial *NodeMaterial, defines js.Value, uniformBuffers string) {

	args := make([]interface{}, 0, 4+0)

	args = append(args, state.JSObject())
	args = append(args, nodeMaterial.JSObject())
	args = append(args, defines)
	args = append(args, uniformBuffers)

	s.p.Call("updateUniformsAndSamples", args...)
}

/*

// BuildId returns the BuildId property of class ScaleBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.scaleblock#buildid
func (s *ScaleBlock) BuildId(buildId float64) *ScaleBlock {
	p := ba.ctx.Get("ScaleBlock").New(buildId)
	return ScaleBlockFromJSObject(p, ba.ctx)
}

// SetBuildId sets the BuildId property of class ScaleBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.scaleblock#buildid
func (s *ScaleBlock) SetBuildId(buildId float64) *ScaleBlock {
	p := ba.ctx.Get("ScaleBlock").New(buildId)
	return ScaleBlockFromJSObject(p, ba.ctx)
}

// Comments returns the Comments property of class ScaleBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.scaleblock#comments
func (s *ScaleBlock) Comments(comments string) *ScaleBlock {
	p := ba.ctx.Get("ScaleBlock").New(comments)
	return ScaleBlockFromJSObject(p, ba.ctx)
}

// SetComments sets the Comments property of class ScaleBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.scaleblock#comments
func (s *ScaleBlock) SetComments(comments string) *ScaleBlock {
	p := ba.ctx.Get("ScaleBlock").New(comments)
	return ScaleBlockFromJSObject(p, ba.ctx)
}

// Factor returns the Factor property of class ScaleBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.scaleblock#factor
func (s *ScaleBlock) Factor(factor *NodeMaterialConnectionPoint) *ScaleBlock {
	p := ba.ctx.Get("ScaleBlock").New(factor.JSObject())
	return ScaleBlockFromJSObject(p, ba.ctx)
}

// SetFactor sets the Factor property of class ScaleBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.scaleblock#factor
func (s *ScaleBlock) SetFactor(factor *NodeMaterialConnectionPoint) *ScaleBlock {
	p := ba.ctx.Get("ScaleBlock").New(factor.JSObject())
	return ScaleBlockFromJSObject(p, ba.ctx)
}

// Input returns the Input property of class ScaleBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.scaleblock#input
func (s *ScaleBlock) Input(input *NodeMaterialConnectionPoint) *ScaleBlock {
	p := ba.ctx.Get("ScaleBlock").New(input.JSObject())
	return ScaleBlockFromJSObject(p, ba.ctx)
}

// SetInput sets the Input property of class ScaleBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.scaleblock#input
func (s *ScaleBlock) SetInput(input *NodeMaterialConnectionPoint) *ScaleBlock {
	p := ba.ctx.Get("ScaleBlock").New(input.JSObject())
	return ScaleBlockFromJSObject(p, ba.ctx)
}

// Inputs returns the Inputs property of class ScaleBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.scaleblock#inputs
func (s *ScaleBlock) Inputs(inputs *NodeMaterialConnectionPoint) *ScaleBlock {
	p := ba.ctx.Get("ScaleBlock").New(inputs.JSObject())
	return ScaleBlockFromJSObject(p, ba.ctx)
}

// SetInputs sets the Inputs property of class ScaleBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.scaleblock#inputs
func (s *ScaleBlock) SetInputs(inputs *NodeMaterialConnectionPoint) *ScaleBlock {
	p := ba.ctx.Get("ScaleBlock").New(inputs.JSObject())
	return ScaleBlockFromJSObject(p, ba.ctx)
}

// IsFinalMerger returns the IsFinalMerger property of class ScaleBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.scaleblock#isfinalmerger
func (s *ScaleBlock) IsFinalMerger(isFinalMerger bool) *ScaleBlock {
	p := ba.ctx.Get("ScaleBlock").New(isFinalMerger)
	return ScaleBlockFromJSObject(p, ba.ctx)
}

// SetIsFinalMerger sets the IsFinalMerger property of class ScaleBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.scaleblock#isfinalmerger
func (s *ScaleBlock) SetIsFinalMerger(isFinalMerger bool) *ScaleBlock {
	p := ba.ctx.Get("ScaleBlock").New(isFinalMerger)
	return ScaleBlockFromJSObject(p, ba.ctx)
}

// IsInput returns the IsInput property of class ScaleBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.scaleblock#isinput
func (s *ScaleBlock) IsInput(isInput bool) *ScaleBlock {
	p := ba.ctx.Get("ScaleBlock").New(isInput)
	return ScaleBlockFromJSObject(p, ba.ctx)
}

// SetIsInput sets the IsInput property of class ScaleBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.scaleblock#isinput
func (s *ScaleBlock) SetIsInput(isInput bool) *ScaleBlock {
	p := ba.ctx.Get("ScaleBlock").New(isInput)
	return ScaleBlockFromJSObject(p, ba.ctx)
}

// IsUnique returns the IsUnique property of class ScaleBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.scaleblock#isunique
func (s *ScaleBlock) IsUnique(isUnique bool) *ScaleBlock {
	p := ba.ctx.Get("ScaleBlock").New(isUnique)
	return ScaleBlockFromJSObject(p, ba.ctx)
}

// SetIsUnique sets the IsUnique property of class ScaleBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.scaleblock#isunique
func (s *ScaleBlock) SetIsUnique(isUnique bool) *ScaleBlock {
	p := ba.ctx.Get("ScaleBlock").New(isUnique)
	return ScaleBlockFromJSObject(p, ba.ctx)
}

// Name returns the Name property of class ScaleBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.scaleblock#name
func (s *ScaleBlock) Name(name string) *ScaleBlock {
	p := ba.ctx.Get("ScaleBlock").New(name)
	return ScaleBlockFromJSObject(p, ba.ctx)
}

// SetName sets the Name property of class ScaleBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.scaleblock#name
func (s *ScaleBlock) SetName(name string) *ScaleBlock {
	p := ba.ctx.Get("ScaleBlock").New(name)
	return ScaleBlockFromJSObject(p, ba.ctx)
}

// Output returns the Output property of class ScaleBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.scaleblock#output
func (s *ScaleBlock) Output(output *NodeMaterialConnectionPoint) *ScaleBlock {
	p := ba.ctx.Get("ScaleBlock").New(output.JSObject())
	return ScaleBlockFromJSObject(p, ba.ctx)
}

// SetOutput sets the Output property of class ScaleBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.scaleblock#output
func (s *ScaleBlock) SetOutput(output *NodeMaterialConnectionPoint) *ScaleBlock {
	p := ba.ctx.Get("ScaleBlock").New(output.JSObject())
	return ScaleBlockFromJSObject(p, ba.ctx)
}

// Outputs returns the Outputs property of class ScaleBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.scaleblock#outputs
func (s *ScaleBlock) Outputs(outputs *NodeMaterialConnectionPoint) *ScaleBlock {
	p := ba.ctx.Get("ScaleBlock").New(outputs.JSObject())
	return ScaleBlockFromJSObject(p, ba.ctx)
}

// SetOutputs sets the Outputs property of class ScaleBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.scaleblock#outputs
func (s *ScaleBlock) SetOutputs(outputs *NodeMaterialConnectionPoint) *ScaleBlock {
	p := ba.ctx.Get("ScaleBlock").New(outputs.JSObject())
	return ScaleBlockFromJSObject(p, ba.ctx)
}

// Target returns the Target property of class ScaleBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.scaleblock#target
func (s *ScaleBlock) Target(target js.Value) *ScaleBlock {
	p := ba.ctx.Get("ScaleBlock").New(target)
	return ScaleBlockFromJSObject(p, ba.ctx)
}

// SetTarget sets the Target property of class ScaleBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.scaleblock#target
func (s *ScaleBlock) SetTarget(target js.Value) *ScaleBlock {
	p := ba.ctx.Get("ScaleBlock").New(target)
	return ScaleBlockFromJSObject(p, ba.ctx)
}

// UniqueId returns the UniqueId property of class ScaleBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.scaleblock#uniqueid
func (s *ScaleBlock) UniqueId(uniqueId float64) *ScaleBlock {
	p := ba.ctx.Get("ScaleBlock").New(uniqueId)
	return ScaleBlockFromJSObject(p, ba.ctx)
}

// SetUniqueId sets the UniqueId property of class ScaleBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.scaleblock#uniqueid
func (s *ScaleBlock) SetUniqueId(uniqueId float64) *ScaleBlock {
	p := ba.ctx.Get("ScaleBlock").New(uniqueId)
	return ScaleBlockFromJSObject(p, ba.ctx)
}

*/
