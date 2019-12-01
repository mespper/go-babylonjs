// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// GLTFLoader represents a babylon.js GLTFLoader.
// The glTF 2.0 loader
type GLTFLoader struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (g *GLTFLoader) JSObject() js.Value { return g.p }

// GLTFLoader returns a GLTFLoader JavaScript class.
func (ba *Babylon) GLTFLoader() *GLTFLoader {
	p := ba.ctx.Get("GLTFLoader")
	return GLTFLoaderFromJSObject(p, ba.ctx)
}

// GLTFLoaderFromJSObject returns a wrapped GLTFLoader JavaScript class.
func GLTFLoaderFromJSObject(p js.Value, ctx js.Value) *GLTFLoader {
	return &GLTFLoader{p: p, ctx: ctx}
}

// GLTFLoaderArrayToJSArray returns a JavaScript Array for the wrapped array.
func GLTFLoaderArrayToJSArray(array []*GLTFLoader) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// AddPointerMetadata calls the AddPointerMetadata method on the GLTFLoader object.
//
// https://doc.babylonjs.com/api/classes/babylon.gltfloader#addpointermetadata
func (g *GLTFLoader) AddPointerMetadata(babylonObject js.Value, pointer string) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, babylonObject)
	args = append(args, pointer)

	g.p.Call("AddPointerMetadata", args...)
}

// CreateMaterial calls the CreateMaterial method on the GLTFLoader object.
//
// https://doc.babylonjs.com/api/classes/babylon.gltfloader#creatematerial
func (g *GLTFLoader) CreateMaterial(context string, material *IMaterial, babylonDrawMode float64) *Material {

	args := make([]interface{}, 0, 3+0)

	args = append(args, context)
	args = append(args, material.JSObject())
	args = append(args, babylonDrawMode)

	retVal := g.p.Call("createMaterial", args...)
	return MaterialFromJSObject(retVal, g.ctx)
}

// EndPerformanceCounter calls the EndPerformanceCounter method on the GLTFLoader object.
//
// https://doc.babylonjs.com/api/classes/babylon.gltfloader#endperformancecounter
func (g *GLTFLoader) EndPerformanceCounter(counterName string) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, counterName)

	g.p.Call("endPerformanceCounter", args...)
}

// IsExtensionUsed calls the IsExtensionUsed method on the GLTFLoader object.
//
// https://doc.babylonjs.com/api/classes/babylon.gltfloader#isextensionused
func (g *GLTFLoader) IsExtensionUsed(name string) bool {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	retVal := g.p.Call("isExtensionUsed", args...)
	return retVal.Bool()
}

// LoadAnimationAsync calls the LoadAnimationAsync method on the GLTFLoader object.
//
// https://doc.babylonjs.com/api/classes/babylon.gltfloader#loadanimationasync
func (g *GLTFLoader) LoadAnimationAsync(context string, animation *IAnimation) *Promise {

	args := make([]interface{}, 0, 2+0)

	args = append(args, context)
	args = append(args, animation.JSObject())

	retVal := g.p.Call("loadAnimationAsync", args...)
	return PromiseFromJSObject(retVal, g.ctx)
}

// LoadBufferViewAsync calls the LoadBufferViewAsync method on the GLTFLoader object.
//
// https://doc.babylonjs.com/api/classes/babylon.gltfloader#loadbufferviewasync
func (g *GLTFLoader) LoadBufferViewAsync(context string, bufferView *IBufferView) *Promise {

	args := make([]interface{}, 0, 2+0)

	args = append(args, context)
	args = append(args, bufferView.JSObject())

	retVal := g.p.Call("loadBufferViewAsync", args...)
	return PromiseFromJSObject(retVal, g.ctx)
}

// GLTFLoaderLoadCameraAsyncOpts contains optional parameters for GLTFLoader.LoadCameraAsync.
type GLTFLoaderLoadCameraAsyncOpts struct {
	Assign func()
}

// LoadCameraAsync calls the LoadCameraAsync method on the GLTFLoader object.
//
// https://doc.babylonjs.com/api/classes/babylon.gltfloader#loadcameraasync
func (g *GLTFLoader) LoadCameraAsync(context string, camera *ICamera, opts *GLTFLoaderLoadCameraAsyncOpts) *Promise {
	if opts == nil {
		opts = &GLTFLoaderLoadCameraAsyncOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, context)
	args = append(args, camera.JSObject())

	if opts.Assign == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Assign)
	}

	retVal := g.p.Call("loadCameraAsync", args...)
	return PromiseFromJSObject(retVal, g.ctx)
}

// LoadImageAsync calls the LoadImageAsync method on the GLTFLoader object.
//
// https://doc.babylonjs.com/api/classes/babylon.gltfloader#loadimageasync
func (g *GLTFLoader) LoadImageAsync(context string, image *IImage) *Promise {

	args := make([]interface{}, 0, 2+0)

	args = append(args, context)
	args = append(args, image.JSObject())

	retVal := g.p.Call("loadImageAsync", args...)
	return PromiseFromJSObject(retVal, g.ctx)
}

// LoadMaterialAlphaProperties calls the LoadMaterialAlphaProperties method on the GLTFLoader object.
//
// https://doc.babylonjs.com/api/classes/babylon.gltfloader#loadmaterialalphaproperties
func (g *GLTFLoader) LoadMaterialAlphaProperties(context string, material *IMaterial, babylonMaterial *Material) {

	args := make([]interface{}, 0, 3+0)

	args = append(args, context)
	args = append(args, material.JSObject())
	args = append(args, babylonMaterial.JSObject())

	g.p.Call("loadMaterialAlphaProperties", args...)
}

// LoadMaterialBasePropertiesAsync calls the LoadMaterialBasePropertiesAsync method on the GLTFLoader object.
//
// https://doc.babylonjs.com/api/classes/babylon.gltfloader#loadmaterialbasepropertiesasync
func (g *GLTFLoader) LoadMaterialBasePropertiesAsync(context string, material *IMaterial, babylonMaterial *Material) *Promise {

	args := make([]interface{}, 0, 3+0)

	args = append(args, context)
	args = append(args, material.JSObject())
	args = append(args, babylonMaterial.JSObject())

	retVal := g.p.Call("loadMaterialBasePropertiesAsync", args...)
	return PromiseFromJSObject(retVal, g.ctx)
}

// LoadMaterialPropertiesAsync calls the LoadMaterialPropertiesAsync method on the GLTFLoader object.
//
// https://doc.babylonjs.com/api/classes/babylon.gltfloader#loadmaterialpropertiesasync
func (g *GLTFLoader) LoadMaterialPropertiesAsync(context string, material *IMaterial, babylonMaterial *Material) *Promise {

	args := make([]interface{}, 0, 3+0)

	args = append(args, context)
	args = append(args, material.JSObject())
	args = append(args, babylonMaterial.JSObject())

	retVal := g.p.Call("loadMaterialPropertiesAsync", args...)
	return PromiseFromJSObject(retVal, g.ctx)
}

// GLTFLoaderLoadNodeAsyncOpts contains optional parameters for GLTFLoader.LoadNodeAsync.
type GLTFLoaderLoadNodeAsyncOpts struct {
	Assign func()
}

// LoadNodeAsync calls the LoadNodeAsync method on the GLTFLoader object.
//
// https://doc.babylonjs.com/api/classes/babylon.gltfloader#loadnodeasync
func (g *GLTFLoader) LoadNodeAsync(context string, node *INode, opts *GLTFLoaderLoadNodeAsyncOpts) *Promise {
	if opts == nil {
		opts = &GLTFLoaderLoadNodeAsyncOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, context)
	args = append(args, node.JSObject())

	if opts.Assign == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Assign)
	}

	retVal := g.p.Call("loadNodeAsync", args...)
	return PromiseFromJSObject(retVal, g.ctx)
}

// LoadSceneAsync calls the LoadSceneAsync method on the GLTFLoader object.
//
// https://doc.babylonjs.com/api/classes/babylon.gltfloader#loadsceneasync
func (g *GLTFLoader) LoadSceneAsync(context string, scene *IScene) *Promise {

	args := make([]interface{}, 0, 2+0)

	args = append(args, context)
	args = append(args, scene.JSObject())

	retVal := g.p.Call("loadSceneAsync", args...)
	return PromiseFromJSObject(retVal, g.ctx)
}

// GLTFLoaderLoadTextureInfoAsyncOpts contains optional parameters for GLTFLoader.LoadTextureInfoAsync.
type GLTFLoaderLoadTextureInfoAsyncOpts struct {
	Assign func()
}

// LoadTextureInfoAsync calls the LoadTextureInfoAsync method on the GLTFLoader object.
//
// https://doc.babylonjs.com/api/classes/babylon.gltfloader#loadtextureinfoasync
func (g *GLTFLoader) LoadTextureInfoAsync(context string, textureInfo *ITextureInfo, opts *GLTFLoaderLoadTextureInfoAsyncOpts) *Promise {
	if opts == nil {
		opts = &GLTFLoaderLoadTextureInfoAsyncOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, context)
	args = append(args, textureInfo.JSObject())

	if opts.Assign == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Assign)
	}

	retVal := g.p.Call("loadTextureInfoAsync", args...)
	return PromiseFromJSObject(retVal, g.ctx)
}

// LoadUriAsync calls the LoadUriAsync method on the GLTFLoader object.
//
// https://doc.babylonjs.com/api/classes/babylon.gltfloader#loaduriasync
func (g *GLTFLoader) LoadUriAsync(context string, property *IProperty, uri string) *Promise {

	args := make([]interface{}, 0, 3+0)

	args = append(args, context)
	args = append(args, property.JSObject())
	args = append(args, uri)

	retVal := g.p.Call("loadUriAsync", args...)
	return PromiseFromJSObject(retVal, g.ctx)
}

// Log calls the Log method on the GLTFLoader object.
//
// https://doc.babylonjs.com/api/classes/babylon.gltfloader#log
func (g *GLTFLoader) Log(message string) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, message)

	g.p.Call("log", args...)
}

// LogClose calls the LogClose method on the GLTFLoader object.
//
// https://doc.babylonjs.com/api/classes/babylon.gltfloader#logclose
func (g *GLTFLoader) LogClose() {

	g.p.Call("logClose")
}

// LogOpen calls the LogOpen method on the GLTFLoader object.
//
// https://doc.babylonjs.com/api/classes/babylon.gltfloader#logopen
func (g *GLTFLoader) LogOpen(message string) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, message)

	g.p.Call("logOpen", args...)
}

// RegisterExtension calls the RegisterExtension method on the GLTFLoader object.
//
// https://doc.babylonjs.com/api/classes/babylon.gltfloader#registerextension
func (g *GLTFLoader) RegisterExtension(name string, factory func()) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, name)
	args = append(args, js.FuncOf(func(this js.Value, args []js.Value) interface{} { factory(); return nil }))

	g.p.Call("RegisterExtension", args...)
}

// StartPerformanceCounter calls the StartPerformanceCounter method on the GLTFLoader object.
//
// https://doc.babylonjs.com/api/classes/babylon.gltfloader#startperformancecounter
func (g *GLTFLoader) StartPerformanceCounter(counterName string) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, counterName)

	g.p.Call("startPerformanceCounter", args...)
}

// UnregisterExtension calls the UnregisterExtension method on the GLTFLoader object.
//
// https://doc.babylonjs.com/api/classes/babylon.gltfloader#unregisterextension
func (g *GLTFLoader) UnregisterExtension(name string) bool {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	retVal := g.p.Call("UnregisterExtension", args...)
	return retVal.Bool()
}

/*

// BabylonScene returns the BabylonScene property of class GLTFLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.gltfloader#babylonscene
func (g *GLTFLoader) BabylonScene(babylonScene *Scene) *GLTFLoader {
	p := ba.ctx.Get("GLTFLoader").New(babylonScene.JSObject())
	return GLTFLoaderFromJSObject(p, ba.ctx)
}

// SetBabylonScene sets the BabylonScene property of class GLTFLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.gltfloader#babylonscene
func (g *GLTFLoader) SetBabylonScene(babylonScene *Scene) *GLTFLoader {
	p := ba.ctx.Get("GLTFLoader").New(babylonScene.JSObject())
	return GLTFLoaderFromJSObject(p, ba.ctx)
}

// Bin returns the Bin property of class GLTFLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.gltfloader#bin
func (g *GLTFLoader) Bin(bin *IDataBuffer) *GLTFLoader {
	p := ba.ctx.Get("GLTFLoader").New(bin.JSObject())
	return GLTFLoaderFromJSObject(p, ba.ctx)
}

// SetBin sets the Bin property of class GLTFLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.gltfloader#bin
func (g *GLTFLoader) SetBin(bin *IDataBuffer) *GLTFLoader {
	p := ba.ctx.Get("GLTFLoader").New(bin.JSObject())
	return GLTFLoaderFromJSObject(p, ba.ctx)
}

// Gltf returns the Gltf property of class GLTFLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.gltfloader#gltf
func (g *GLTFLoader) Gltf(gltf *IGLTF) *GLTFLoader {
	p := ba.ctx.Get("GLTFLoader").New(gltf.JSObject())
	return GLTFLoaderFromJSObject(p, ba.ctx)
}

// SetGltf sets the Gltf property of class GLTFLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.gltfloader#gltf
func (g *GLTFLoader) SetGltf(gltf *IGLTF) *GLTFLoader {
	p := ba.ctx.Get("GLTFLoader").New(gltf.JSObject())
	return GLTFLoaderFromJSObject(p, ba.ctx)
}

// Parent returns the Parent property of class GLTFLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.gltfloader#parent
func (g *GLTFLoader) Parent(parent *GLTFFileLoader) *GLTFLoader {
	p := ba.ctx.Get("GLTFLoader").New(parent.JSObject())
	return GLTFLoaderFromJSObject(p, ba.ctx)
}

// SetParent sets the Parent property of class GLTFLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.gltfloader#parent
func (g *GLTFLoader) SetParent(parent *GLTFFileLoader) *GLTFLoader {
	p := ba.ctx.Get("GLTFLoader").New(parent.JSObject())
	return GLTFLoaderFromJSObject(p, ba.ctx)
}

// RootBabylonMesh returns the RootBabylonMesh property of class GLTFLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.gltfloader#rootbabylonmesh
func (g *GLTFLoader) RootBabylonMesh(rootBabylonMesh *Mesh) *GLTFLoader {
	p := ba.ctx.Get("GLTFLoader").New(rootBabylonMesh.JSObject())
	return GLTFLoaderFromJSObject(p, ba.ctx)
}

// SetRootBabylonMesh sets the RootBabylonMesh property of class GLTFLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.gltfloader#rootbabylonmesh
func (g *GLTFLoader) SetRootBabylonMesh(rootBabylonMesh *Mesh) *GLTFLoader {
	p := ba.ctx.Get("GLTFLoader").New(rootBabylonMesh.JSObject())
	return GLTFLoaderFromJSObject(p, ba.ctx)
}

// State returns the State property of class GLTFLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.gltfloader#state
func (g *GLTFLoader) State(state *GLTFLoaderState) *GLTFLoader {
	p := ba.ctx.Get("GLTFLoader").New(state.JSObject())
	return GLTFLoaderFromJSObject(p, ba.ctx)
}

// SetState sets the State property of class GLTFLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.gltfloader#state
func (g *GLTFLoader) SetState(state *GLTFLoaderState) *GLTFLoader {
	p := ba.ctx.Get("GLTFLoader").New(state.JSObject())
	return GLTFLoaderFromJSObject(p, ba.ctx)
}

*/
