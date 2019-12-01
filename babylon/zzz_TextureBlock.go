// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// TextureBlock represents a babylon.js TextureBlock.
// Block used to read a texture from a sampler
type TextureBlock struct {
	*NodeMaterialBlock
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (t *TextureBlock) JSObject() js.Value { return t.p }

// TextureBlock returns a TextureBlock JavaScript class.
func (ba *Babylon) TextureBlock() *TextureBlock {
	p := ba.ctx.Get("TextureBlock")
	return TextureBlockFromJSObject(p, ba.ctx)
}

// TextureBlockFromJSObject returns a wrapped TextureBlock JavaScript class.
func TextureBlockFromJSObject(p js.Value, ctx js.Value) *TextureBlock {
	return &TextureBlock{NodeMaterialBlock: NodeMaterialBlockFromJSObject(p, ctx), ctx: ctx}
}

// TextureBlockArrayToJSArray returns a JavaScript Array for the wrapped array.
func TextureBlockArrayToJSArray(array []*TextureBlock) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewTextureBlock returns a new TextureBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.textureblock
func (ba *Babylon) NewTextureBlock(name string) *TextureBlock {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	p := ba.ctx.Get("TextureBlock").New(args...)
	return TextureBlockFromJSObject(p, ba.ctx)
}

// AutoConfigure calls the AutoConfigure method on the TextureBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.textureblock#autoconfigure
func (t *TextureBlock) AutoConfigure(material *NodeMaterial) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, material.JSObject())

	t.p.Call("autoConfigure", args...)
}

// TextureBlockBindOpts contains optional parameters for TextureBlock.Bind.
type TextureBlockBindOpts struct {
	Mesh *Mesh
}

// Bind calls the Bind method on the TextureBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.textureblock#bind
func (t *TextureBlock) Bind(effect *Effect, nodeMaterial *NodeMaterial, opts *TextureBlockBindOpts) {
	if opts == nil {
		opts = &TextureBlockBindOpts{}
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

// Build calls the Build method on the TextureBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.textureblock#build
func (t *TextureBlock) Build(state *NodeMaterialBuildState, activeBlocks *NodeMaterialBlock) bool {

	args := make([]interface{}, 0, 2+0)

	args = append(args, state.JSObject())
	args = append(args, activeBlocks.JSObject())

	retVal := t.p.Call("build", args...)
	return retVal.Bool()
}

// TextureBlockCloneOpts contains optional parameters for TextureBlock.Clone.
type TextureBlockCloneOpts struct {
	RootUrl *string
}

// Clone calls the Clone method on the TextureBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.textureblock#clone
func (t *TextureBlock) Clone(scene *Scene, opts *TextureBlockCloneOpts) *NodeMaterialBlock {
	if opts == nil {
		opts = &TextureBlockCloneOpts{}
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

// TextureBlockConnectToOpts contains optional parameters for TextureBlock.ConnectTo.
type TextureBlockConnectToOpts struct {
	Options map[string]interface{}
}

// ConnectTo calls the ConnectTo method on the TextureBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.textureblock#connectto
func (t *TextureBlock) ConnectTo(other *NodeMaterialBlock, opts *TextureBlockConnectToOpts) *TextureBlock {
	if opts == nil {
		opts = &TextureBlockConnectToOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, other.JSObject())

	if opts.Options == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Options)
	}

	retVal := t.p.Call("connectTo", args...)
	return TextureBlockFromJSObject(retVal, t.ctx)
}

// Dispose calls the Dispose method on the TextureBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.textureblock#dispose
func (t *TextureBlock) Dispose() {

	t.p.Call("dispose")
}

// GetClassName calls the GetClassName method on the TextureBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.textureblock#getclassname
func (t *TextureBlock) GetClassName() string {

	retVal := t.p.Call("getClassName")
	return retVal.String()
}

// TextureBlockGetFirstAvailableInputOpts contains optional parameters for TextureBlock.GetFirstAvailableInput.
type TextureBlockGetFirstAvailableInputOpts struct {
	ForOutput *NodeMaterialConnectionPoint
}

// GetFirstAvailableInput calls the GetFirstAvailableInput method on the TextureBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.textureblock#getfirstavailableinput
func (t *TextureBlock) GetFirstAvailableInput(opts *TextureBlockGetFirstAvailableInputOpts) *NodeMaterialConnectionPoint {
	if opts == nil {
		opts = &TextureBlockGetFirstAvailableInputOpts{}
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

// TextureBlockGetFirstAvailableOutputOpts contains optional parameters for TextureBlock.GetFirstAvailableOutput.
type TextureBlockGetFirstAvailableOutputOpts struct {
	ForBlock *NodeMaterialBlock
}

// GetFirstAvailableOutput calls the GetFirstAvailableOutput method on the TextureBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.textureblock#getfirstavailableoutput
func (t *TextureBlock) GetFirstAvailableOutput(opts *TextureBlockGetFirstAvailableOutputOpts) *NodeMaterialConnectionPoint {
	if opts == nil {
		opts = &TextureBlockGetFirstAvailableOutputOpts{}
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

// GetInputByName calls the GetInputByName method on the TextureBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.textureblock#getinputbyname
func (t *TextureBlock) GetInputByName(name string) *NodeMaterialConnectionPoint {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	retVal := t.p.Call("getInputByName", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, t.ctx)
}

// GetOutputByName calls the GetOutputByName method on the TextureBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.textureblock#getoutputbyname
func (t *TextureBlock) GetOutputByName(name string) *NodeMaterialConnectionPoint {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	retVal := t.p.Call("getOutputByName", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, t.ctx)
}

// GetSiblingOutput calls the GetSiblingOutput method on the TextureBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.textureblock#getsiblingoutput
func (t *TextureBlock) GetSiblingOutput(current *NodeMaterialConnectionPoint) *NodeMaterialConnectionPoint {

	args := make([]interface{}, 0, 1+0)

	args = append(args, current.JSObject())

	retVal := t.p.Call("getSiblingOutput", args...)
	return NodeMaterialConnectionPointFromJSObject(retVal, t.ctx)
}

// Initialize calls the Initialize method on the TextureBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.textureblock#initialize
func (t *TextureBlock) Initialize(state *NodeMaterialBuildState) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, state.JSObject())

	t.p.Call("initialize", args...)
}

// TextureBlockInitializeDefinesOpts contains optional parameters for TextureBlock.InitializeDefines.
type TextureBlockInitializeDefinesOpts struct {
	UseInstances *bool
}

// InitializeDefines calls the InitializeDefines method on the TextureBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.textureblock#initializedefines
func (t *TextureBlock) InitializeDefines(mesh *AbstractMesh, nodeMaterial *NodeMaterial, defines js.Value, opts *TextureBlockInitializeDefinesOpts) {
	if opts == nil {
		opts = &TextureBlockInitializeDefinesOpts{}
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

// IsReady calls the IsReady method on the TextureBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.textureblock#isready
func (t *TextureBlock) IsReady() bool {

	retVal := t.p.Call("isReady")
	return retVal.Bool()
}

// PrepareDefines calls the PrepareDefines method on the TextureBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.textureblock#preparedefines
func (t *TextureBlock) PrepareDefines(mesh *AbstractMesh, nodeMaterial *NodeMaterial, defines js.Value) {

	args := make([]interface{}, 0, 3+0)

	args = append(args, mesh.JSObject())
	args = append(args, nodeMaterial.JSObject())
	args = append(args, defines)

	t.p.Call("prepareDefines", args...)
}

// ProvideFallbacks calls the ProvideFallbacks method on the TextureBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.textureblock#providefallbacks
func (t *TextureBlock) ProvideFallbacks(mesh *AbstractMesh, fallbacks *EffectFallbacks) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, mesh.JSObject())
	args = append(args, fallbacks.JSObject())

	t.p.Call("provideFallbacks", args...)
}

// TextureBlockRegisterInputOpts contains optional parameters for TextureBlock.RegisterInput.
type TextureBlockRegisterInputOpts struct {
	IsOptional *bool
	Target     js.Value
}

// RegisterInput calls the RegisterInput method on the TextureBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.textureblock#registerinput
func (t *TextureBlock) RegisterInput(name string, jsType js.Value, opts *TextureBlockRegisterInputOpts) *TextureBlock {
	if opts == nil {
		opts = &TextureBlockRegisterInputOpts{}
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
	return TextureBlockFromJSObject(retVal, t.ctx)
}

// TextureBlockRegisterOutputOpts contains optional parameters for TextureBlock.RegisterOutput.
type TextureBlockRegisterOutputOpts struct {
	Target js.Value
}

// RegisterOutput calls the RegisterOutput method on the TextureBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.textureblock#registeroutput
func (t *TextureBlock) RegisterOutput(name string, jsType js.Value, opts *TextureBlockRegisterOutputOpts) *TextureBlock {
	if opts == nil {
		opts = &TextureBlockRegisterOutputOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, name)
	args = append(args, jsType)

	args = append(args, opts.Target)

	retVal := t.p.Call("registerOutput", args...)
	return TextureBlockFromJSObject(retVal, t.ctx)
}

// ReplaceRepeatableContent calls the ReplaceRepeatableContent method on the TextureBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.textureblock#replacerepeatablecontent
func (t *TextureBlock) ReplaceRepeatableContent(vertexShaderState *NodeMaterialBuildState, fragmentShaderState *NodeMaterialBuildState, mesh *AbstractMesh, defines js.Value) {

	args := make([]interface{}, 0, 4+0)

	args = append(args, vertexShaderState.JSObject())
	args = append(args, fragmentShaderState.JSObject())
	args = append(args, mesh.JSObject())
	args = append(args, defines)

	t.p.Call("replaceRepeatableContent", args...)
}

// Serialize calls the Serialize method on the TextureBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.textureblock#serialize
func (t *TextureBlock) Serialize() interface{} {

	retVal := t.p.Call("serialize")
	return retVal
}

// UpdateUniformsAndSamples calls the UpdateUniformsAndSamples method on the TextureBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.textureblock#updateuniformsandsamples
func (t *TextureBlock) UpdateUniformsAndSamples(state *NodeMaterialBuildState, nodeMaterial *NodeMaterial, defines js.Value, uniformBuffers string) {

	args := make([]interface{}, 0, 4+0)

	args = append(args, state.JSObject())
	args = append(args, nodeMaterial.JSObject())
	args = append(args, defines)
	args = append(args, uniformBuffers)

	t.p.Call("updateUniformsAndSamples", args...)
}

// _deserialize calls the _deserialize method on the TextureBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.textureblock#_deserialize
func (t *TextureBlock) _deserialize(serializationObject interface{}, scene *Scene, rootUrl string) {

	args := make([]interface{}, 0, 3+0)

	args = append(args, serializationObject)
	args = append(args, scene.JSObject())
	args = append(args, rootUrl)

	t.p.Call("_deserialize", args...)
}

/*

// A returns the A property of class TextureBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.textureblock#a
func (t *TextureBlock) A(a *NodeMaterialConnectionPoint) *TextureBlock {
	p := ba.ctx.Get("TextureBlock").New(a.JSObject())
	return TextureBlockFromJSObject(p, ba.ctx)
}

// SetA sets the A property of class TextureBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.textureblock#a
func (t *TextureBlock) SetA(a *NodeMaterialConnectionPoint) *TextureBlock {
	p := ba.ctx.Get("TextureBlock").New(a.JSObject())
	return TextureBlockFromJSObject(p, ba.ctx)
}

// B returns the B property of class TextureBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.textureblock#b
func (t *TextureBlock) B(b *NodeMaterialConnectionPoint) *TextureBlock {
	p := ba.ctx.Get("TextureBlock").New(b.JSObject())
	return TextureBlockFromJSObject(p, ba.ctx)
}

// SetB sets the B property of class TextureBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.textureblock#b
func (t *TextureBlock) SetB(b *NodeMaterialConnectionPoint) *TextureBlock {
	p := ba.ctx.Get("TextureBlock").New(b.JSObject())
	return TextureBlockFromJSObject(p, ba.ctx)
}

// BuildId returns the BuildId property of class TextureBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.textureblock#buildid
func (t *TextureBlock) BuildId(buildId float64) *TextureBlock {
	p := ba.ctx.Get("TextureBlock").New(buildId)
	return TextureBlockFromJSObject(p, ba.ctx)
}

// SetBuildId sets the BuildId property of class TextureBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.textureblock#buildid
func (t *TextureBlock) SetBuildId(buildId float64) *TextureBlock {
	p := ba.ctx.Get("TextureBlock").New(buildId)
	return TextureBlockFromJSObject(p, ba.ctx)
}

// Comments returns the Comments property of class TextureBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.textureblock#comments
func (t *TextureBlock) Comments(comments string) *TextureBlock {
	p := ba.ctx.Get("TextureBlock").New(comments)
	return TextureBlockFromJSObject(p, ba.ctx)
}

// SetComments sets the Comments property of class TextureBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.textureblock#comments
func (t *TextureBlock) SetComments(comments string) *TextureBlock {
	p := ba.ctx.Get("TextureBlock").New(comments)
	return TextureBlockFromJSObject(p, ba.ctx)
}

// G returns the G property of class TextureBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.textureblock#g
func (t *TextureBlock) G(g *NodeMaterialConnectionPoint) *TextureBlock {
	p := ba.ctx.Get("TextureBlock").New(g.JSObject())
	return TextureBlockFromJSObject(p, ba.ctx)
}

// SetG sets the G property of class TextureBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.textureblock#g
func (t *TextureBlock) SetG(g *NodeMaterialConnectionPoint) *TextureBlock {
	p := ba.ctx.Get("TextureBlock").New(g.JSObject())
	return TextureBlockFromJSObject(p, ba.ctx)
}

// Inputs returns the Inputs property of class TextureBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.textureblock#inputs
func (t *TextureBlock) Inputs(inputs *NodeMaterialConnectionPoint) *TextureBlock {
	p := ba.ctx.Get("TextureBlock").New(inputs.JSObject())
	return TextureBlockFromJSObject(p, ba.ctx)
}

// SetInputs sets the Inputs property of class TextureBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.textureblock#inputs
func (t *TextureBlock) SetInputs(inputs *NodeMaterialConnectionPoint) *TextureBlock {
	p := ba.ctx.Get("TextureBlock").New(inputs.JSObject())
	return TextureBlockFromJSObject(p, ba.ctx)
}

// IsFinalMerger returns the IsFinalMerger property of class TextureBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.textureblock#isfinalmerger
func (t *TextureBlock) IsFinalMerger(isFinalMerger bool) *TextureBlock {
	p := ba.ctx.Get("TextureBlock").New(isFinalMerger)
	return TextureBlockFromJSObject(p, ba.ctx)
}

// SetIsFinalMerger sets the IsFinalMerger property of class TextureBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.textureblock#isfinalmerger
func (t *TextureBlock) SetIsFinalMerger(isFinalMerger bool) *TextureBlock {
	p := ba.ctx.Get("TextureBlock").New(isFinalMerger)
	return TextureBlockFromJSObject(p, ba.ctx)
}

// IsInput returns the IsInput property of class TextureBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.textureblock#isinput
func (t *TextureBlock) IsInput(isInput bool) *TextureBlock {
	p := ba.ctx.Get("TextureBlock").New(isInput)
	return TextureBlockFromJSObject(p, ba.ctx)
}

// SetIsInput sets the IsInput property of class TextureBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.textureblock#isinput
func (t *TextureBlock) SetIsInput(isInput bool) *TextureBlock {
	p := ba.ctx.Get("TextureBlock").New(isInput)
	return TextureBlockFromJSObject(p, ba.ctx)
}

// IsUnique returns the IsUnique property of class TextureBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.textureblock#isunique
func (t *TextureBlock) IsUnique(isUnique bool) *TextureBlock {
	p := ba.ctx.Get("TextureBlock").New(isUnique)
	return TextureBlockFromJSObject(p, ba.ctx)
}

// SetIsUnique sets the IsUnique property of class TextureBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.textureblock#isunique
func (t *TextureBlock) SetIsUnique(isUnique bool) *TextureBlock {
	p := ba.ctx.Get("TextureBlock").New(isUnique)
	return TextureBlockFromJSObject(p, ba.ctx)
}

// Name returns the Name property of class TextureBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.textureblock#name
func (t *TextureBlock) Name(name string) *TextureBlock {
	p := ba.ctx.Get("TextureBlock").New(name)
	return TextureBlockFromJSObject(p, ba.ctx)
}

// SetName sets the Name property of class TextureBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.textureblock#name
func (t *TextureBlock) SetName(name string) *TextureBlock {
	p := ba.ctx.Get("TextureBlock").New(name)
	return TextureBlockFromJSObject(p, ba.ctx)
}

// Outputs returns the Outputs property of class TextureBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.textureblock#outputs
func (t *TextureBlock) Outputs(outputs *NodeMaterialConnectionPoint) *TextureBlock {
	p := ba.ctx.Get("TextureBlock").New(outputs.JSObject())
	return TextureBlockFromJSObject(p, ba.ctx)
}

// SetOutputs sets the Outputs property of class TextureBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.textureblock#outputs
func (t *TextureBlock) SetOutputs(outputs *NodeMaterialConnectionPoint) *TextureBlock {
	p := ba.ctx.Get("TextureBlock").New(outputs.JSObject())
	return TextureBlockFromJSObject(p, ba.ctx)
}

// R returns the R property of class TextureBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.textureblock#r
func (t *TextureBlock) R(r *NodeMaterialConnectionPoint) *TextureBlock {
	p := ba.ctx.Get("TextureBlock").New(r.JSObject())
	return TextureBlockFromJSObject(p, ba.ctx)
}

// SetR sets the R property of class TextureBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.textureblock#r
func (t *TextureBlock) SetR(r *NodeMaterialConnectionPoint) *TextureBlock {
	p := ba.ctx.Get("TextureBlock").New(r.JSObject())
	return TextureBlockFromJSObject(p, ba.ctx)
}

// Rgb returns the Rgb property of class TextureBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.textureblock#rgb
func (t *TextureBlock) Rgb(rgb *NodeMaterialConnectionPoint) *TextureBlock {
	p := ba.ctx.Get("TextureBlock").New(rgb.JSObject())
	return TextureBlockFromJSObject(p, ba.ctx)
}

// SetRgb sets the Rgb property of class TextureBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.textureblock#rgb
func (t *TextureBlock) SetRgb(rgb *NodeMaterialConnectionPoint) *TextureBlock {
	p := ba.ctx.Get("TextureBlock").New(rgb.JSObject())
	return TextureBlockFromJSObject(p, ba.ctx)
}

// Rgba returns the Rgba property of class TextureBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.textureblock#rgba
func (t *TextureBlock) Rgba(rgba *NodeMaterialConnectionPoint) *TextureBlock {
	p := ba.ctx.Get("TextureBlock").New(rgba.JSObject())
	return TextureBlockFromJSObject(p, ba.ctx)
}

// SetRgba sets the Rgba property of class TextureBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.textureblock#rgba
func (t *TextureBlock) SetRgba(rgba *NodeMaterialConnectionPoint) *TextureBlock {
	p := ba.ctx.Get("TextureBlock").New(rgba.JSObject())
	return TextureBlockFromJSObject(p, ba.ctx)
}

// Target returns the Target property of class TextureBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.textureblock#target
func (t *TextureBlock) Target(target js.Value) *TextureBlock {
	p := ba.ctx.Get("TextureBlock").New(target)
	return TextureBlockFromJSObject(p, ba.ctx)
}

// SetTarget sets the Target property of class TextureBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.textureblock#target
func (t *TextureBlock) SetTarget(target js.Value) *TextureBlock {
	p := ba.ctx.Get("TextureBlock").New(target)
	return TextureBlockFromJSObject(p, ba.ctx)
}

// Texture returns the Texture property of class TextureBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.textureblock#texture
func (t *TextureBlock) Texture(texture *Texture) *TextureBlock {
	p := ba.ctx.Get("TextureBlock").New(texture.JSObject())
	return TextureBlockFromJSObject(p, ba.ctx)
}

// SetTexture sets the Texture property of class TextureBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.textureblock#texture
func (t *TextureBlock) SetTexture(texture *Texture) *TextureBlock {
	p := ba.ctx.Get("TextureBlock").New(texture.JSObject())
	return TextureBlockFromJSObject(p, ba.ctx)
}

// UniqueId returns the UniqueId property of class TextureBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.textureblock#uniqueid
func (t *TextureBlock) UniqueId(uniqueId float64) *TextureBlock {
	p := ba.ctx.Get("TextureBlock").New(uniqueId)
	return TextureBlockFromJSObject(p, ba.ctx)
}

// SetUniqueId sets the UniqueId property of class TextureBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.textureblock#uniqueid
func (t *TextureBlock) SetUniqueId(uniqueId float64) *TextureBlock {
	p := ba.ctx.Get("TextureBlock").New(uniqueId)
	return TextureBlockFromJSObject(p, ba.ctx)
}

// Uv returns the Uv property of class TextureBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.textureblock#uv
func (t *TextureBlock) Uv(uv *NodeMaterialConnectionPoint) *TextureBlock {
	p := ba.ctx.Get("TextureBlock").New(uv.JSObject())
	return TextureBlockFromJSObject(p, ba.ctx)
}

// SetUv sets the Uv property of class TextureBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.textureblock#uv
func (t *TextureBlock) SetUv(uv *NodeMaterialConnectionPoint) *TextureBlock {
	p := ba.ctx.Get("TextureBlock").New(uv.JSObject())
	return TextureBlockFromJSObject(p, ba.ctx)
}

*/
