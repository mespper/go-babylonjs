// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// NodeMaterial represents a babylon.js NodeMaterial.
// Class used to create a node based material built by assembling shader blocks
type NodeMaterial struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (n *NodeMaterial) JSObject() js.Value { return n.p }

// NodeMaterial returns a NodeMaterial JavaScript class.
func (ba *Babylon) NodeMaterial() *NodeMaterial {
	p := ba.ctx.Get("NodeMaterial")
	return NodeMaterialFromJSObject(p, ba.ctx)
}

// NodeMaterialFromJSObject returns a wrapped NodeMaterial JavaScript class.
func NodeMaterialFromJSObject(p js.Value, ctx js.Value) *NodeMaterial {
	return &NodeMaterial{p: p, ctx: ctx}
}

// NodeMaterialArrayToJSArray returns a JavaScript Array for the wrapped array.
func NodeMaterialArrayToJSArray(array []*NodeMaterial) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewNodeMaterialOpts contains optional parameters for NewNodeMaterial.
type NewNodeMaterialOpts struct {
	Scene   *Scene
	Options *INodeMaterialOptions
}

// NewNodeMaterial returns a new NodeMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.nodematerial
func (ba *Babylon) NewNodeMaterial(name string, opts *NewNodeMaterialOpts) *NodeMaterial {
	if opts == nil {
		opts = &NewNodeMaterialOpts{}
	}

	args := make([]interface{}, 0, 1+2)

	args = append(args, name)

	if opts.Scene == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Scene.JSObject())
	}
	if opts.Options == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Options.JSObject())
	}

	p := ba.ctx.Get("NodeMaterial").New(args...)
	return NodeMaterialFromJSObject(p, ba.ctx)
}

// AddOutputNode calls the AddOutputNode method on the NodeMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.nodematerial#addoutputnode
func (n *NodeMaterial) AddOutputNode(node *NodeMaterialBlock) *NodeMaterial {

	args := make([]interface{}, 0, 1+0)

	args = append(args, node.JSObject())

	retVal := n.p.Call("addOutputNode", args...)
	return NodeMaterialFromJSObject(retVal, n.ctx)
}

// BindForSubMesh calls the BindForSubMesh method on the NodeMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.nodematerial#bindforsubmesh
func (n *NodeMaterial) BindForSubMesh(world *Matrix, mesh *Mesh, subMesh *SubMesh) {

	args := make([]interface{}, 0, 3+0)

	args = append(args, world.JSObject())
	args = append(args, mesh.JSObject())
	args = append(args, subMesh.JSObject())

	n.p.Call("bindForSubMesh", args...)
}

// BindOnlyWorldMatrix calls the BindOnlyWorldMatrix method on the NodeMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.nodematerial#bindonlyworldmatrix
func (n *NodeMaterial) BindOnlyWorldMatrix(world *Matrix) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, world.JSObject())

	n.p.Call("bindOnlyWorldMatrix", args...)
}

// NodeMaterialBuildOpts contains optional parameters for NodeMaterial.Build.
type NodeMaterialBuildOpts struct {
	Verbose *bool
}

// Build calls the Build method on the NodeMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.nodematerial#build
func (n *NodeMaterial) Build(opts *NodeMaterialBuildOpts) {
	if opts == nil {
		opts = &NodeMaterialBuildOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.Verbose == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Verbose)
	}

	n.p.Call("build", args...)
}

// Clear calls the Clear method on the NodeMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.nodematerial#clear
func (n *NodeMaterial) Clear() {

	n.p.Call("clear")
}

// NodeMaterialCreateDefaultOpts contains optional parameters for NodeMaterial.CreateDefault.
type NodeMaterialCreateDefaultOpts struct {
	Scene *Scene
}

// CreateDefault calls the CreateDefault method on the NodeMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.nodematerial#createdefault
func (n *NodeMaterial) CreateDefault(name string, opts *NodeMaterialCreateDefaultOpts) *NodeMaterial {
	if opts == nil {
		opts = &NodeMaterialCreateDefaultOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, name)

	if opts.Scene == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Scene.JSObject())
	}

	retVal := n.p.Call("CreateDefault", args...)
	return NodeMaterialFromJSObject(retVal, n.ctx)
}

// NodeMaterialDisposeOpts contains optional parameters for NodeMaterial.Dispose.
type NodeMaterialDisposeOpts struct {
	ForceDisposeEffect   *bool
	ForceDisposeTextures *bool
	NotBoundToMesh       *bool
}

// Dispose calls the Dispose method on the NodeMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.nodematerial#dispose
func (n *NodeMaterial) Dispose(opts *NodeMaterialDisposeOpts) {
	if opts == nil {
		opts = &NodeMaterialDisposeOpts{}
	}

	args := make([]interface{}, 0, 0+3)

	if opts.ForceDisposeEffect == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.ForceDisposeEffect)
	}
	if opts.ForceDisposeTextures == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.ForceDisposeTextures)
	}
	if opts.NotBoundToMesh == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.NotBoundToMesh)
	}

	n.p.Call("dispose", args...)
}

// NodeMaterialEditOpts contains optional parameters for NodeMaterial.Edit.
type NodeMaterialEditOpts struct {
	Config *INodeMaterialEditorOptions
}

// Edit calls the Edit method on the NodeMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.nodematerial#edit
func (n *NodeMaterial) Edit(opts *NodeMaterialEditOpts) *Promise {
	if opts == nil {
		opts = &NodeMaterialEditOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.Config == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Config.JSObject())
	}

	retVal := n.p.Call("edit", args...)
	return PromiseFromJSObject(retVal, n.ctx)
}

// GenerateCode calls the GenerateCode method on the NodeMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.nodematerial#generatecode
func (n *NodeMaterial) GenerateCode() string {

	retVal := n.p.Call("generateCode")
	return retVal.String()
}

// GetActiveTextures calls the GetActiveTextures method on the NodeMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.nodematerial#getactivetextures
func (n *NodeMaterial) GetActiveTextures() []*BaseTexture {

	retVal := n.p.Call("getActiveTextures")
	result := []*BaseTexture{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, BaseTextureFromJSObject(retVal.Index(ri), n.ctx))
	}
	return result
}

// GetBlockByName calls the GetBlockByName method on the NodeMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.nodematerial#getblockbyname
func (n *NodeMaterial) GetBlockByName(name string) *NodeMaterialBlock {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	retVal := n.p.Call("getBlockByName", args...)
	return NodeMaterialBlockFromJSObject(retVal, n.ctx)
}

// GetBlockByPredicate calls the GetBlockByPredicate method on the NodeMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.nodematerial#getblockbypredicate
func (n *NodeMaterial) GetBlockByPredicate(predicate JSFunc) *NodeMaterialBlock {

	args := make([]interface{}, 0, 1+0)

	args = append(args, js.FuncOf(predicate))

	retVal := n.p.Call("getBlockByPredicate", args...)
	return NodeMaterialBlockFromJSObject(retVal, n.ctx)
}

// GetClassName calls the GetClassName method on the NodeMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.nodematerial#getclassname
func (n *NodeMaterial) GetClassName() string {

	retVal := n.p.Call("getClassName")
	return retVal.String()
}

// GetInputBlockByPredicate calls the GetInputBlockByPredicate method on the NodeMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.nodematerial#getinputblockbypredicate
func (n *NodeMaterial) GetInputBlockByPredicate(predicate JSFunc) *InputBlock {

	args := make([]interface{}, 0, 1+0)

	args = append(args, js.FuncOf(predicate))

	retVal := n.p.Call("getInputBlockByPredicate", args...)
	return InputBlockFromJSObject(retVal, n.ctx)
}

// GetInputBlocks calls the GetInputBlocks method on the NodeMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.nodematerial#getinputblocks
func (n *NodeMaterial) GetInputBlocks() []*InputBlock {

	retVal := n.p.Call("getInputBlocks")
	result := []*InputBlock{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, InputBlockFromJSObject(retVal.Index(ri), n.ctx))
	}
	return result
}

// GetTextureBlocks calls the GetTextureBlocks method on the NodeMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.nodematerial#gettextureblocks
func (n *NodeMaterial) GetTextureBlocks() []*ReflectionTextureBlock {

	retVal := n.p.Call("getTextureBlocks")
	result := []*ReflectionTextureBlock{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, ReflectionTextureBlockFromJSObject(retVal.Index(ri), n.ctx))
	}
	return result
}

// HasTexture calls the HasTexture method on the NodeMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.nodematerial#hastexture
func (n *NodeMaterial) HasTexture(texture *BaseTexture) bool {

	args := make([]interface{}, 0, 1+0)

	args = append(args, texture.JSObject())

	retVal := n.p.Call("hasTexture", args...)
	return retVal.Bool()
}

// NodeMaterialIsReadyForSubMeshOpts contains optional parameters for NodeMaterial.IsReadyForSubMesh.
type NodeMaterialIsReadyForSubMeshOpts struct {
	UseInstances *bool
}

// IsReadyForSubMesh calls the IsReadyForSubMesh method on the NodeMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.nodematerial#isreadyforsubmesh
func (n *NodeMaterial) IsReadyForSubMesh(mesh *AbstractMesh, subMesh *SubMesh, opts *NodeMaterialIsReadyForSubMeshOpts) bool {
	if opts == nil {
		opts = &NodeMaterialIsReadyForSubMeshOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, mesh.JSObject())
	args = append(args, subMesh.JSObject())

	if opts.UseInstances == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.UseInstances)
	}

	retVal := n.p.Call("isReadyForSubMesh", args...)
	return retVal.Bool()
}

// LoadAsync calls the LoadAsync method on the NodeMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.nodematerial#loadasync
func (n *NodeMaterial) LoadAsync(url string) *Promise {

	args := make([]interface{}, 0, 1+0)

	args = append(args, url)

	retVal := n.p.Call("loadAsync", args...)
	return PromiseFromJSObject(retVal, n.ctx)
}

// NodeMaterialLoadFromSerializationOpts contains optional parameters for NodeMaterial.LoadFromSerialization.
type NodeMaterialLoadFromSerializationOpts struct {
	RootUrl *string
}

// LoadFromSerialization calls the LoadFromSerialization method on the NodeMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.nodematerial#loadfromserialization
func (n *NodeMaterial) LoadFromSerialization(source interface{}, opts *NodeMaterialLoadFromSerializationOpts) {
	if opts == nil {
		opts = &NodeMaterialLoadFromSerializationOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, source)

	if opts.RootUrl == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.RootUrl)
	}

	n.p.Call("loadFromSerialization", args...)
}

// NeedAlphaBlending calls the NeedAlphaBlending method on the NodeMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.nodematerial#needalphablending
func (n *NodeMaterial) NeedAlphaBlending() bool {

	retVal := n.p.Call("needAlphaBlending")
	return retVal.Bool()
}

// NeedAlphaTesting calls the NeedAlphaTesting method on the NodeMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.nodematerial#needalphatesting
func (n *NodeMaterial) NeedAlphaTesting() bool {

	retVal := n.p.Call("needAlphaTesting")
	return retVal.Bool()
}

// Optimize calls the Optimize method on the NodeMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.nodematerial#optimize
func (n *NodeMaterial) Optimize() {

	n.p.Call("optimize")
}

// NodeMaterialParseOpts contains optional parameters for NodeMaterial.Parse.
type NodeMaterialParseOpts struct {
	RootUrl *string
}

// Parse calls the Parse method on the NodeMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.nodematerial#parse
func (n *NodeMaterial) Parse(source interface{}, scene *Scene, opts *NodeMaterialParseOpts) *NodeMaterial {
	if opts == nil {
		opts = &NodeMaterialParseOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, source)
	args = append(args, scene.JSObject())

	if opts.RootUrl == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.RootUrl)
	}

	retVal := n.p.Call("Parse", args...)
	return NodeMaterialFromJSObject(retVal, n.ctx)
}

// RegisterOptimizer calls the RegisterOptimizer method on the NodeMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.nodematerial#registeroptimizer
func (n *NodeMaterial) RegisterOptimizer(optimizer *NodeMaterialOptimizer) *NodeMaterial {

	args := make([]interface{}, 0, 1+0)

	args = append(args, optimizer.JSObject())

	retVal := n.p.Call("registerOptimizer", args...)
	return NodeMaterialFromJSObject(retVal, n.ctx)
}

// RemoveBlock calls the RemoveBlock method on the NodeMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.nodematerial#removeblock
func (n *NodeMaterial) RemoveBlock(block *NodeMaterialBlock) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, block.JSObject())

	n.p.Call("removeBlock", args...)
}

// RemoveOutputNode calls the RemoveOutputNode method on the NodeMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.nodematerial#removeoutputnode
func (n *NodeMaterial) RemoveOutputNode(node *NodeMaterialBlock) *NodeMaterial {

	args := make([]interface{}, 0, 1+0)

	args = append(args, node.JSObject())

	retVal := n.p.Call("removeOutputNode", args...)
	return NodeMaterialFromJSObject(retVal, n.ctx)
}

// Serialize calls the Serialize method on the NodeMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.nodematerial#serialize
func (n *NodeMaterial) Serialize() interface{} {

	retVal := n.p.Call("serialize")
	return retVal
}

// SetToDefault calls the SetToDefault method on the NodeMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.nodematerial#settodefault
func (n *NodeMaterial) SetToDefault() {

	n.p.Call("setToDefault")
}

// UnregisterOptimizer calls the UnregisterOptimizer method on the NodeMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.nodematerial#unregisteroptimizer
func (n *NodeMaterial) UnregisterOptimizer(optimizer *NodeMaterialOptimizer) *NodeMaterial {

	args := make([]interface{}, 0, 1+0)

	args = append(args, optimizer.JSObject())

	retVal := n.p.Call("unregisterOptimizer", args...)
	return NodeMaterialFromJSObject(retVal, n.ctx)
}

// AttachedBlocks returns the AttachedBlocks property of class NodeMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.nodematerial#attachedblocks
func (n *NodeMaterial) AttachedBlocks() []*NodeMaterialBlock {
	retVal := n.p.Get("attachedBlocks")
	result := []*NodeMaterialBlock{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, NodeMaterialBlockFromJSObject(retVal.Index(ri), n.ctx))
	}
	return result
}

// SetAttachedBlocks sets the AttachedBlocks property of class NodeMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.nodematerial#attachedblocks
func (n *NodeMaterial) SetAttachedBlocks(attachedBlocks []*NodeMaterialBlock) *NodeMaterial {
	n.p.Set("attachedBlocks", attachedBlocks)
	return n
}

// CompiledShaders returns the CompiledShaders property of class NodeMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.nodematerial#compiledshaders
func (n *NodeMaterial) CompiledShaders() string {
	retVal := n.p.Get("compiledShaders")
	return retVal.String()
}

// SetCompiledShaders sets the CompiledShaders property of class NodeMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.nodematerial#compiledshaders
func (n *NodeMaterial) SetCompiledShaders(compiledShaders string) *NodeMaterial {
	n.p.Set("compiledShaders", compiledShaders)
	return n
}

// EditorData returns the EditorData property of class NodeMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.nodematerial#editordata
func (n *NodeMaterial) EditorData() interface{} {
	retVal := n.p.Get("editorData")
	return retVal
}

// SetEditorData sets the EditorData property of class NodeMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.nodematerial#editordata
func (n *NodeMaterial) SetEditorData(editorData interface{}) *NodeMaterial {
	n.p.Set("editorData", editorData)
	return n
}

// EditorURL returns the EditorURL property of class NodeMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.nodematerial#editorurl
func (n *NodeMaterial) EditorURL() string {
	retVal := n.p.Get("EditorURL")
	return retVal.String()
}

// SetEditorURL sets the EditorURL property of class NodeMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.nodematerial#editorurl
func (n *NodeMaterial) SetEditorURL(EditorURL string) *NodeMaterial {
	n.p.Set("EditorURL", EditorURL)
	return n
}

// IgnoreAlpha returns the IgnoreAlpha property of class NodeMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.nodematerial#ignorealpha
func (n *NodeMaterial) IgnoreAlpha() bool {
	retVal := n.p.Get("ignoreAlpha")
	return retVal.Bool()
}

// SetIgnoreAlpha sets the IgnoreAlpha property of class NodeMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.nodematerial#ignorealpha
func (n *NodeMaterial) SetIgnoreAlpha(ignoreAlpha bool) *NodeMaterial {
	n.p.Set("ignoreAlpha", ignoreAlpha)
	return n
}

// ImageProcessingConfiguration returns the ImageProcessingConfiguration property of class NodeMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.nodematerial#imageprocessingconfiguration
func (n *NodeMaterial) ImageProcessingConfiguration() *ImageProcessingConfiguration {
	retVal := n.p.Get("imageProcessingConfiguration")
	return ImageProcessingConfigurationFromJSObject(retVal, n.ctx)
}

// SetImageProcessingConfiguration sets the ImageProcessingConfiguration property of class NodeMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.nodematerial#imageprocessingconfiguration
func (n *NodeMaterial) SetImageProcessingConfiguration(imageProcessingConfiguration *ImageProcessingConfiguration) *NodeMaterial {
	n.p.Set("imageProcessingConfiguration", imageProcessingConfiguration.JSObject())
	return n
}

// MaxSimultaneousLights returns the MaxSimultaneousLights property of class NodeMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.nodematerial#maxsimultaneouslights
func (n *NodeMaterial) MaxSimultaneousLights() float64 {
	retVal := n.p.Get("maxSimultaneousLights")
	return retVal.Float()
}

// SetMaxSimultaneousLights sets the MaxSimultaneousLights property of class NodeMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.nodematerial#maxsimultaneouslights
func (n *NodeMaterial) SetMaxSimultaneousLights(maxSimultaneousLights float64) *NodeMaterial {
	n.p.Set("maxSimultaneousLights", maxSimultaneousLights)
	return n
}

// OnBuildObservable returns the OnBuildObservable property of class NodeMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.nodematerial#onbuildobservable
func (n *NodeMaterial) OnBuildObservable() *Observable {
	retVal := n.p.Get("onBuildObservable")
	return ObservableFromJSObject(retVal, n.ctx)
}

// SetOnBuildObservable sets the OnBuildObservable property of class NodeMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.nodematerial#onbuildobservable
func (n *NodeMaterial) SetOnBuildObservable(onBuildObservable *Observable) *NodeMaterial {
	n.p.Set("onBuildObservable", onBuildObservable.JSObject())
	return n
}

// Options returns the Options property of class NodeMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.nodematerial#options
func (n *NodeMaterial) Options() *INodeMaterialOptions {
	retVal := n.p.Get("options")
	return INodeMaterialOptionsFromJSObject(retVal, n.ctx)
}

// SetOptions sets the Options property of class NodeMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.nodematerial#options
func (n *NodeMaterial) SetOptions(options *INodeMaterialOptions) *NodeMaterial {
	n.p.Set("options", options.JSObject())
	return n
}

// _fragmentOutputNodes returns the _fragmentOutputNodes property of class NodeMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.nodematerial#_fragmentoutputnodes
func (n *NodeMaterial) _fragmentOutputNodes() []*NodeMaterialBlock {
	retVal := n.p.Get("_fragmentOutputNodes")
	result := []*NodeMaterialBlock{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, NodeMaterialBlockFromJSObject(retVal.Index(ri), n.ctx))
	}
	return result
}

// Set_fragmentOutputNodes sets the _fragmentOutputNodes property of class NodeMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.nodematerial#_fragmentoutputnodes
func (n *NodeMaterial) Set_fragmentOutputNodes(_fragmentOutputNodes []*NodeMaterialBlock) *NodeMaterial {
	n.p.Set("_fragmentOutputNodes", _fragmentOutputNodes)
	return n
}

// _vertexOutputNodes returns the _vertexOutputNodes property of class NodeMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.nodematerial#_vertexoutputnodes
func (n *NodeMaterial) _vertexOutputNodes() []*NodeMaterialBlock {
	retVal := n.p.Get("_vertexOutputNodes")
	result := []*NodeMaterialBlock{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, NodeMaterialBlockFromJSObject(retVal.Index(ri), n.ctx))
	}
	return result
}

// Set_vertexOutputNodes sets the _vertexOutputNodes property of class NodeMaterial.
//
// https://doc.babylonjs.com/api/classes/babylon.nodematerial#_vertexoutputnodes
func (n *NodeMaterial) Set_vertexOutputNodes(_vertexOutputNodes []*NodeMaterialBlock) *NodeMaterial {
	n.p.Set("_vertexOutputNodes", _vertexOutputNodes)
	return n
}
