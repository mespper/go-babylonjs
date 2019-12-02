// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// GLTFFileLoader represents a babylon.js GLTFFileLoader.
// File loader for loading glTF files into a scene.
type GLTFFileLoader struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (g *GLTFFileLoader) JSObject() js.Value { return g.p }

// GLTFFileLoader returns a GLTFFileLoader JavaScript class.
func (ba *Babylon) GLTFFileLoader() *GLTFFileLoader {
	p := ba.ctx.Get("GLTFFileLoader")
	return GLTFFileLoaderFromJSObject(p, ba.ctx)
}

// GLTFFileLoaderFromJSObject returns a wrapped GLTFFileLoader JavaScript class.
func GLTFFileLoaderFromJSObject(p js.Value, ctx js.Value) *GLTFFileLoader {
	return &GLTFFileLoader{p: p, ctx: ctx}
}

// GLTFFileLoaderArrayToJSArray returns a JavaScript Array for the wrapped array.
func GLTFFileLoaderArrayToJSArray(array []*GLTFFileLoader) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// Dispose calls the Dispose method on the GLTFFileLoader object.
//
// https://doc.babylonjs.com/api/classes/babylon.gltffileloader#dispose
func (g *GLTFFileLoader) Dispose() {

	g.p.Call("dispose")
}

// GLTFFileLoaderRewriteRootURLOpts contains optional parameters for GLTFFileLoader.RewriteRootURL.
type GLTFFileLoaderRewriteRootURLOpts struct {
	ResponseURL *string
}

// RewriteRootURL calls the RewriteRootURL method on the GLTFFileLoader object.
//
// https://doc.babylonjs.com/api/classes/babylon.gltffileloader#rewriterooturl
func (g *GLTFFileLoader) RewriteRootURL(rootUrl string, opts *GLTFFileLoaderRewriteRootURLOpts) string {
	if opts == nil {
		opts = &GLTFFileLoaderRewriteRootURLOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, rootUrl)

	if opts.ResponseURL == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.ResponseURL)
	}

	retVal := g.p.Call("rewriteRootURL", args...)
	return retVal.String()
}

// WhenCompleteAsync calls the WhenCompleteAsync method on the GLTFFileLoader object.
//
// https://doc.babylonjs.com/api/classes/babylon.gltffileloader#whencompleteasync
func (g *GLTFFileLoader) WhenCompleteAsync() *Promise {

	retVal := g.p.Call("whenCompleteAsync")
	return PromiseFromJSObject(retVal, g.ctx)
}

// AnimationStartMode returns the AnimationStartMode property of class GLTFFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.gltffileloader#animationstartmode
func (g *GLTFFileLoader) AnimationStartMode() js.Value {
	retVal := g.p.Get("animationStartMode")
	return retVal
}

// SetAnimationStartMode sets the AnimationStartMode property of class GLTFFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.gltffileloader#animationstartmode
func (g *GLTFFileLoader) SetAnimationStartMode(animationStartMode js.Value) *GLTFFileLoader {
	g.p.Set("animationStartMode", animationStartMode)
	return g
}

// CapturePerformanceCounters returns the CapturePerformanceCounters property of class GLTFFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.gltffileloader#captureperformancecounters
func (g *GLTFFileLoader) CapturePerformanceCounters() bool {
	retVal := g.p.Get("capturePerformanceCounters")
	return retVal.Bool()
}

// SetCapturePerformanceCounters sets the CapturePerformanceCounters property of class GLTFFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.gltffileloader#captureperformancecounters
func (g *GLTFFileLoader) SetCapturePerformanceCounters(capturePerformanceCounters bool) *GLTFFileLoader {
	g.p.Set("capturePerformanceCounters", capturePerformanceCounters)
	return g
}

// CompileMaterials returns the CompileMaterials property of class GLTFFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.gltffileloader#compilematerials
func (g *GLTFFileLoader) CompileMaterials() bool {
	retVal := g.p.Get("compileMaterials")
	return retVal.Bool()
}

// SetCompileMaterials sets the CompileMaterials property of class GLTFFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.gltffileloader#compilematerials
func (g *GLTFFileLoader) SetCompileMaterials(compileMaterials bool) *GLTFFileLoader {
	g.p.Set("compileMaterials", compileMaterials)
	return g
}

// CompileShadowGenerators returns the CompileShadowGenerators property of class GLTFFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.gltffileloader#compileshadowgenerators
func (g *GLTFFileLoader) CompileShadowGenerators() bool {
	retVal := g.p.Get("compileShadowGenerators")
	return retVal.Bool()
}

// SetCompileShadowGenerators sets the CompileShadowGenerators property of class GLTFFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.gltffileloader#compileshadowgenerators
func (g *GLTFFileLoader) SetCompileShadowGenerators(compileShadowGenerators bool) *GLTFFileLoader {
	g.p.Set("compileShadowGenerators", compileShadowGenerators)
	return g
}

// CoordinateSystemMode returns the CoordinateSystemMode property of class GLTFFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.gltffileloader#coordinatesystemmode
func (g *GLTFFileLoader) CoordinateSystemMode() js.Value {
	retVal := g.p.Get("coordinateSystemMode")
	return retVal
}

// SetCoordinateSystemMode sets the CoordinateSystemMode property of class GLTFFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.gltffileloader#coordinatesystemmode
func (g *GLTFFileLoader) SetCoordinateSystemMode(coordinateSystemMode js.Value) *GLTFFileLoader {
	g.p.Set("coordinateSystemMode", coordinateSystemMode)
	return g
}

// CreateInstances returns the CreateInstances property of class GLTFFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.gltffileloader#createinstances
func (g *GLTFFileLoader) CreateInstances() bool {
	retVal := g.p.Get("createInstances")
	return retVal.Bool()
}

// SetCreateInstances sets the CreateInstances property of class GLTFFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.gltffileloader#createinstances
func (g *GLTFFileLoader) SetCreateInstances(createInstances bool) *GLTFFileLoader {
	g.p.Set("createInstances", createInstances)
	return g
}

// LoaderState returns the LoaderState property of class GLTFFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.gltffileloader#loaderstate
func (g *GLTFFileLoader) LoaderState() js.Value {
	retVal := g.p.Get("loaderState")
	return retVal
}

// SetLoaderState sets the LoaderState property of class GLTFFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.gltffileloader#loaderstate
func (g *GLTFFileLoader) SetLoaderState(loaderState js.Value) *GLTFFileLoader {
	g.p.Set("loaderState", loaderState)
	return g
}

// LoggingEnabled returns the LoggingEnabled property of class GLTFFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.gltffileloader#loggingenabled
func (g *GLTFFileLoader) LoggingEnabled() bool {
	retVal := g.p.Get("loggingEnabled")
	return retVal.Bool()
}

// SetLoggingEnabled sets the LoggingEnabled property of class GLTFFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.gltffileloader#loggingenabled
func (g *GLTFFileLoader) SetLoggingEnabled(loggingEnabled bool) *GLTFFileLoader {
	g.p.Set("loggingEnabled", loggingEnabled)
	return g
}

// Name returns the Name property of class GLTFFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.gltffileloader#name
func (g *GLTFFileLoader) Name() string {
	retVal := g.p.Get("name")
	return retVal.String()
}

// SetName sets the Name property of class GLTFFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.gltffileloader#name
func (g *GLTFFileLoader) SetName(name string) *GLTFFileLoader {
	g.p.Set("name", name)
	return g
}

// OnCameraLoaded returns the OnCameraLoaded property of class GLTFFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.gltffileloader#oncameraloaded
func (g *GLTFFileLoader) OnCameraLoaded() js.Value {
	retVal := g.p.Get("onCameraLoaded")
	return retVal
}

// SetOnCameraLoaded sets the OnCameraLoaded property of class GLTFFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.gltffileloader#oncameraloaded
func (g *GLTFFileLoader) SetOnCameraLoaded(onCameraLoaded func()) *GLTFFileLoader {
	g.p.Set("onCameraLoaded", js.FuncOf(func(this js.Value, args []js.Value) interface{} { onCameraLoaded(); return nil }))
	return g
}

// OnCameraLoadedObservable returns the OnCameraLoadedObservable property of class GLTFFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.gltffileloader#oncameraloadedobservable
func (g *GLTFFileLoader) OnCameraLoadedObservable() *Observable {
	retVal := g.p.Get("onCameraLoadedObservable")
	return ObservableFromJSObject(retVal, g.ctx)
}

// SetOnCameraLoadedObservable sets the OnCameraLoadedObservable property of class GLTFFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.gltffileloader#oncameraloadedobservable
func (g *GLTFFileLoader) SetOnCameraLoadedObservable(onCameraLoadedObservable *Observable) *GLTFFileLoader {
	g.p.Set("onCameraLoadedObservable", onCameraLoadedObservable.JSObject())
	return g
}

// OnComplete returns the OnComplete property of class GLTFFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.gltffileloader#oncomplete
func (g *GLTFFileLoader) OnComplete() js.Value {
	retVal := g.p.Get("onComplete")
	return retVal
}

// SetOnComplete sets the OnComplete property of class GLTFFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.gltffileloader#oncomplete
func (g *GLTFFileLoader) SetOnComplete(onComplete func()) *GLTFFileLoader {
	g.p.Set("onComplete", js.FuncOf(func(this js.Value, args []js.Value) interface{} { onComplete(); return nil }))
	return g
}

// OnCompleteObservable returns the OnCompleteObservable property of class GLTFFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.gltffileloader#oncompleteobservable
func (g *GLTFFileLoader) OnCompleteObservable() *Observable {
	retVal := g.p.Get("onCompleteObservable")
	return ObservableFromJSObject(retVal, g.ctx)
}

// SetOnCompleteObservable sets the OnCompleteObservable property of class GLTFFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.gltffileloader#oncompleteobservable
func (g *GLTFFileLoader) SetOnCompleteObservable(onCompleteObservable *Observable) *GLTFFileLoader {
	g.p.Set("onCompleteObservable", onCompleteObservable.JSObject())
	return g
}

// OnDispose returns the OnDispose property of class GLTFFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.gltffileloader#ondispose
func (g *GLTFFileLoader) OnDispose() js.Value {
	retVal := g.p.Get("onDispose")
	return retVal
}

// SetOnDispose sets the OnDispose property of class GLTFFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.gltffileloader#ondispose
func (g *GLTFFileLoader) SetOnDispose(onDispose func()) *GLTFFileLoader {
	g.p.Set("onDispose", js.FuncOf(func(this js.Value, args []js.Value) interface{} { onDispose(); return nil }))
	return g
}

// OnDisposeObservable returns the OnDisposeObservable property of class GLTFFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.gltffileloader#ondisposeobservable
func (g *GLTFFileLoader) OnDisposeObservable() *Observable {
	retVal := g.p.Get("onDisposeObservable")
	return ObservableFromJSObject(retVal, g.ctx)
}

// SetOnDisposeObservable sets the OnDisposeObservable property of class GLTFFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.gltffileloader#ondisposeobservable
func (g *GLTFFileLoader) SetOnDisposeObservable(onDisposeObservable *Observable) *GLTFFileLoader {
	g.p.Set("onDisposeObservable", onDisposeObservable.JSObject())
	return g
}

// OnError returns the OnError property of class GLTFFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.gltffileloader#onerror
func (g *GLTFFileLoader) OnError() js.Value {
	retVal := g.p.Get("onError")
	return retVal
}

// SetOnError sets the OnError property of class GLTFFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.gltffileloader#onerror
func (g *GLTFFileLoader) SetOnError(onError func()) *GLTFFileLoader {
	g.p.Set("onError", js.FuncOf(func(this js.Value, args []js.Value) interface{} { onError(); return nil }))
	return g
}

// OnErrorObservable returns the OnErrorObservable property of class GLTFFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.gltffileloader#onerrorobservable
func (g *GLTFFileLoader) OnErrorObservable() *Observable {
	retVal := g.p.Get("onErrorObservable")
	return ObservableFromJSObject(retVal, g.ctx)
}

// SetOnErrorObservable sets the OnErrorObservable property of class GLTFFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.gltffileloader#onerrorobservable
func (g *GLTFFileLoader) SetOnErrorObservable(onErrorObservable *Observable) *GLTFFileLoader {
	g.p.Set("onErrorObservable", onErrorObservable.JSObject())
	return g
}

// OnExtensionLoaded returns the OnExtensionLoaded property of class GLTFFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.gltffileloader#onextensionloaded
func (g *GLTFFileLoader) OnExtensionLoaded() js.Value {
	retVal := g.p.Get("onExtensionLoaded")
	return retVal
}

// SetOnExtensionLoaded sets the OnExtensionLoaded property of class GLTFFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.gltffileloader#onextensionloaded
func (g *GLTFFileLoader) SetOnExtensionLoaded(onExtensionLoaded func()) *GLTFFileLoader {
	g.p.Set("onExtensionLoaded", js.FuncOf(func(this js.Value, args []js.Value) interface{} { onExtensionLoaded(); return nil }))
	return g
}

// OnExtensionLoadedObservable returns the OnExtensionLoadedObservable property of class GLTFFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.gltffileloader#onextensionloadedobservable
func (g *GLTFFileLoader) OnExtensionLoadedObservable() *Observable {
	retVal := g.p.Get("onExtensionLoadedObservable")
	return ObservableFromJSObject(retVal, g.ctx)
}

// SetOnExtensionLoadedObservable sets the OnExtensionLoadedObservable property of class GLTFFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.gltffileloader#onextensionloadedobservable
func (g *GLTFFileLoader) SetOnExtensionLoadedObservable(onExtensionLoadedObservable *Observable) *GLTFFileLoader {
	g.p.Set("onExtensionLoadedObservable", onExtensionLoadedObservable.JSObject())
	return g
}

// OnMaterialLoaded returns the OnMaterialLoaded property of class GLTFFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.gltffileloader#onmaterialloaded
func (g *GLTFFileLoader) OnMaterialLoaded() js.Value {
	retVal := g.p.Get("onMaterialLoaded")
	return retVal
}

// SetOnMaterialLoaded sets the OnMaterialLoaded property of class GLTFFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.gltffileloader#onmaterialloaded
func (g *GLTFFileLoader) SetOnMaterialLoaded(onMaterialLoaded func()) *GLTFFileLoader {
	g.p.Set("onMaterialLoaded", js.FuncOf(func(this js.Value, args []js.Value) interface{} { onMaterialLoaded(); return nil }))
	return g
}

// OnMaterialLoadedObservable returns the OnMaterialLoadedObservable property of class GLTFFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.gltffileloader#onmaterialloadedobservable
func (g *GLTFFileLoader) OnMaterialLoadedObservable() *Observable {
	retVal := g.p.Get("onMaterialLoadedObservable")
	return ObservableFromJSObject(retVal, g.ctx)
}

// SetOnMaterialLoadedObservable sets the OnMaterialLoadedObservable property of class GLTFFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.gltffileloader#onmaterialloadedobservable
func (g *GLTFFileLoader) SetOnMaterialLoadedObservable(onMaterialLoadedObservable *Observable) *GLTFFileLoader {
	g.p.Set("onMaterialLoadedObservable", onMaterialLoadedObservable.JSObject())
	return g
}

// OnMeshLoaded returns the OnMeshLoaded property of class GLTFFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.gltffileloader#onmeshloaded
func (g *GLTFFileLoader) OnMeshLoaded() js.Value {
	retVal := g.p.Get("onMeshLoaded")
	return retVal
}

// SetOnMeshLoaded sets the OnMeshLoaded property of class GLTFFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.gltffileloader#onmeshloaded
func (g *GLTFFileLoader) SetOnMeshLoaded(onMeshLoaded func()) *GLTFFileLoader {
	g.p.Set("onMeshLoaded", js.FuncOf(func(this js.Value, args []js.Value) interface{} { onMeshLoaded(); return nil }))
	return g
}

// OnMeshLoadedObservable returns the OnMeshLoadedObservable property of class GLTFFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.gltffileloader#onmeshloadedobservable
func (g *GLTFFileLoader) OnMeshLoadedObservable() *Observable {
	retVal := g.p.Get("onMeshLoadedObservable")
	return ObservableFromJSObject(retVal, g.ctx)
}

// SetOnMeshLoadedObservable sets the OnMeshLoadedObservable property of class GLTFFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.gltffileloader#onmeshloadedobservable
func (g *GLTFFileLoader) SetOnMeshLoadedObservable(onMeshLoadedObservable *Observable) *GLTFFileLoader {
	g.p.Set("onMeshLoadedObservable", onMeshLoadedObservable.JSObject())
	return g
}

// OnParsed returns the OnParsed property of class GLTFFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.gltffileloader#onparsed
func (g *GLTFFileLoader) OnParsed() js.Value {
	retVal := g.p.Get("onParsed")
	return retVal
}

// SetOnParsed sets the OnParsed property of class GLTFFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.gltffileloader#onparsed
func (g *GLTFFileLoader) SetOnParsed(onParsed func()) *GLTFFileLoader {
	g.p.Set("onParsed", js.FuncOf(func(this js.Value, args []js.Value) interface{} { onParsed(); return nil }))
	return g
}

// OnParsedObservable returns the OnParsedObservable property of class GLTFFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.gltffileloader#onparsedobservable
func (g *GLTFFileLoader) OnParsedObservable() *Observable {
	retVal := g.p.Get("onParsedObservable")
	return ObservableFromJSObject(retVal, g.ctx)
}

// SetOnParsedObservable sets the OnParsedObservable property of class GLTFFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.gltffileloader#onparsedobservable
func (g *GLTFFileLoader) SetOnParsedObservable(onParsedObservable *Observable) *GLTFFileLoader {
	g.p.Set("onParsedObservable", onParsedObservable.JSObject())
	return g
}

// OnTextureLoaded returns the OnTextureLoaded property of class GLTFFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.gltffileloader#ontextureloaded
func (g *GLTFFileLoader) OnTextureLoaded() js.Value {
	retVal := g.p.Get("onTextureLoaded")
	return retVal
}

// SetOnTextureLoaded sets the OnTextureLoaded property of class GLTFFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.gltffileloader#ontextureloaded
func (g *GLTFFileLoader) SetOnTextureLoaded(onTextureLoaded func()) *GLTFFileLoader {
	g.p.Set("onTextureLoaded", js.FuncOf(func(this js.Value, args []js.Value) interface{} { onTextureLoaded(); return nil }))
	return g
}

// OnTextureLoadedObservable returns the OnTextureLoadedObservable property of class GLTFFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.gltffileloader#ontextureloadedobservable
func (g *GLTFFileLoader) OnTextureLoadedObservable() *Observable {
	retVal := g.p.Get("onTextureLoadedObservable")
	return ObservableFromJSObject(retVal, g.ctx)
}

// SetOnTextureLoadedObservable sets the OnTextureLoadedObservable property of class GLTFFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.gltffileloader#ontextureloadedobservable
func (g *GLTFFileLoader) SetOnTextureLoadedObservable(onTextureLoadedObservable *Observable) *GLTFFileLoader {
	g.p.Set("onTextureLoadedObservable", onTextureLoadedObservable.JSObject())
	return g
}

// OnValidated returns the OnValidated property of class GLTFFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.gltffileloader#onvalidated
func (g *GLTFFileLoader) OnValidated() js.Value {
	retVal := g.p.Get("onValidated")
	return retVal
}

// SetOnValidated sets the OnValidated property of class GLTFFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.gltffileloader#onvalidated
func (g *GLTFFileLoader) SetOnValidated(onValidated func()) *GLTFFileLoader {
	g.p.Set("onValidated", js.FuncOf(func(this js.Value, args []js.Value) interface{} { onValidated(); return nil }))
	return g
}

// OnValidatedObservable returns the OnValidatedObservable property of class GLTFFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.gltffileloader#onvalidatedobservable
func (g *GLTFFileLoader) OnValidatedObservable() *Observable {
	retVal := g.p.Get("onValidatedObservable")
	return ObservableFromJSObject(retVal, g.ctx)
}

// SetOnValidatedObservable sets the OnValidatedObservable property of class GLTFFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.gltffileloader#onvalidatedobservable
func (g *GLTFFileLoader) SetOnValidatedObservable(onValidatedObservable *Observable) *GLTFFileLoader {
	g.p.Set("onValidatedObservable", onValidatedObservable.JSObject())
	return g
}

// PreprocessUrlAsync returns the PreprocessUrlAsync property of class GLTFFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.gltffileloader#preprocessurlasync
func (g *GLTFFileLoader) PreprocessUrlAsync() js.Value {
	retVal := g.p.Get("preprocessUrlAsync")
	return retVal
}

// SetPreprocessUrlAsync sets the PreprocessUrlAsync property of class GLTFFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.gltffileloader#preprocessurlasync
func (g *GLTFFileLoader) SetPreprocessUrlAsync(preprocessUrlAsync func()) *GLTFFileLoader {
	g.p.Set("preprocessUrlAsync", js.FuncOf(func(this js.Value, args []js.Value) interface{} { preprocessUrlAsync(); return nil }))
	return g
}

// TransparencyAsCoverage returns the TransparencyAsCoverage property of class GLTFFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.gltffileloader#transparencyascoverage
func (g *GLTFFileLoader) TransparencyAsCoverage() bool {
	retVal := g.p.Get("transparencyAsCoverage")
	return retVal.Bool()
}

// SetTransparencyAsCoverage sets the TransparencyAsCoverage property of class GLTFFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.gltffileloader#transparencyascoverage
func (g *GLTFFileLoader) SetTransparencyAsCoverage(transparencyAsCoverage bool) *GLTFFileLoader {
	g.p.Set("transparencyAsCoverage", transparencyAsCoverage)
	return g
}

// UseClipPlane returns the UseClipPlane property of class GLTFFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.gltffileloader#useclipplane
func (g *GLTFFileLoader) UseClipPlane() bool {
	retVal := g.p.Get("useClipPlane")
	return retVal.Bool()
}

// SetUseClipPlane sets the UseClipPlane property of class GLTFFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.gltffileloader#useclipplane
func (g *GLTFFileLoader) SetUseClipPlane(useClipPlane bool) *GLTFFileLoader {
	g.p.Set("useClipPlane", useClipPlane)
	return g
}

// UseRangeRequests returns the UseRangeRequests property of class GLTFFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.gltffileloader#userangerequests
func (g *GLTFFileLoader) UseRangeRequests() bool {
	retVal := g.p.Get("useRangeRequests")
	return retVal.Bool()
}

// SetUseRangeRequests sets the UseRangeRequests property of class GLTFFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.gltffileloader#userangerequests
func (g *GLTFFileLoader) SetUseRangeRequests(useRangeRequests bool) *GLTFFileLoader {
	g.p.Set("useRangeRequests", useRangeRequests)
	return g
}

// Validate returns the Validate property of class GLTFFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.gltffileloader#validate
func (g *GLTFFileLoader) Validate() bool {
	retVal := g.p.Get("validate")
	return retVal.Bool()
}

// SetValidate sets the Validate property of class GLTFFileLoader.
//
// https://doc.babylonjs.com/api/classes/babylon.gltffileloader#validate
func (g *GLTFFileLoader) SetValidate(validate bool) *GLTFFileLoader {
	g.p.Set("validate", validate)
	return g
}
