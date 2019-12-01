// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ISceneLoaderPlugin represents a babylon.js ISceneLoaderPlugin.
// Interface used to define a SceneLoader plugin
type ISceneLoaderPlugin struct {
	*ISceneLoaderPluginBase
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (i *ISceneLoaderPlugin) JSObject() js.Value { return i.p }

// ISceneLoaderPlugin returns a ISceneLoaderPlugin JavaScript class.
func (ba *Babylon) ISceneLoaderPlugin() *ISceneLoaderPlugin {
	p := ba.ctx.Get("ISceneLoaderPlugin")
	return ISceneLoaderPluginFromJSObject(p, ba.ctx)
}

// ISceneLoaderPluginFromJSObject returns a wrapped ISceneLoaderPlugin JavaScript class.
func ISceneLoaderPluginFromJSObject(p js.Value, ctx js.Value) *ISceneLoaderPlugin {
	return &ISceneLoaderPlugin{ISceneLoaderPluginBase: ISceneLoaderPluginBaseFromJSObject(p, ctx), ctx: ctx}
}

// ISceneLoaderPluginArrayToJSArray returns a JavaScript Array for the wrapped array.
func ISceneLoaderPluginArrayToJSArray(array []*ISceneLoaderPlugin) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// CanDirectLoad calls the CanDirectLoad method on the ISceneLoaderPlugin object.
//
// https://doc.babylonjs.com/api/classes/babylon.isceneloaderplugin#candirectload
func (i *ISceneLoaderPlugin) CanDirectLoad(data string) bool {

	args := make([]interface{}, 0, 1+0)

	args = append(args, data)

	retVal := i.p.Call("canDirectLoad", args...)
	return retVal.Bool()
}

// DirectLoad calls the DirectLoad method on the ISceneLoaderPlugin object.
//
// https://doc.babylonjs.com/api/classes/babylon.isceneloaderplugin#directload
func (i *ISceneLoaderPlugin) DirectLoad(scene *Scene, data string) interface{} {

	args := make([]interface{}, 0, 2+0)

	args = append(args, scene.JSObject())
	args = append(args, data)

	retVal := i.p.Call("directLoad", args...)
	return retVal
}

// ISceneLoaderPluginImportMeshOpts contains optional parameters for ISceneLoaderPlugin.ImportMesh.
type ISceneLoaderPluginImportMeshOpts struct {
	OnError func()
}

// ImportMesh calls the ImportMesh method on the ISceneLoaderPlugin object.
//
// https://doc.babylonjs.com/api/classes/babylon.isceneloaderplugin#importmesh
func (i *ISceneLoaderPlugin) ImportMesh(meshesNames interface{}, scene *Scene, data interface{}, rootUrl string, meshes *AbstractMesh, particleSystems *IParticleSystem, skeletons *Skeleton, opts *ISceneLoaderPluginImportMeshOpts) bool {
	if opts == nil {
		opts = &ISceneLoaderPluginImportMeshOpts{}
	}

	args := make([]interface{}, 0, 7+1)

	args = append(args, meshesNames)
	args = append(args, scene.JSObject())
	args = append(args, data)
	args = append(args, rootUrl)
	args = append(args, meshes.JSObject())
	args = append(args, particleSystems.JSObject())
	args = append(args, skeletons.JSObject())

	if opts.OnError == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.OnError)
	}

	retVal := i.p.Call("importMesh", args...)
	return retVal.Bool()
}

// ISceneLoaderPluginLoadOpts contains optional parameters for ISceneLoaderPlugin.Load.
type ISceneLoaderPluginLoadOpts struct {
	OnError func()
}

// Load calls the Load method on the ISceneLoaderPlugin object.
//
// https://doc.babylonjs.com/api/classes/babylon.isceneloaderplugin#load
func (i *ISceneLoaderPlugin) Load(scene *Scene, data interface{}, rootUrl string, opts *ISceneLoaderPluginLoadOpts) bool {
	if opts == nil {
		opts = &ISceneLoaderPluginLoadOpts{}
	}

	args := make([]interface{}, 0, 3+1)

	args = append(args, scene.JSObject())
	args = append(args, data)
	args = append(args, rootUrl)

	if opts.OnError == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.OnError)
	}

	retVal := i.p.Call("load", args...)
	return retVal.Bool()
}

// ISceneLoaderPluginLoadAssetContainerOpts contains optional parameters for ISceneLoaderPlugin.LoadAssetContainer.
type ISceneLoaderPluginLoadAssetContainerOpts struct {
	OnError func()
}

// LoadAssetContainer calls the LoadAssetContainer method on the ISceneLoaderPlugin object.
//
// https://doc.babylonjs.com/api/classes/babylon.isceneloaderplugin#loadassetcontainer
func (i *ISceneLoaderPlugin) LoadAssetContainer(scene *Scene, data interface{}, rootUrl string, opts *ISceneLoaderPluginLoadAssetContainerOpts) *AssetContainer {
	if opts == nil {
		opts = &ISceneLoaderPluginLoadAssetContainerOpts{}
	}

	args := make([]interface{}, 0, 3+1)

	args = append(args, scene.JSObject())
	args = append(args, data)
	args = append(args, rootUrl)

	if opts.OnError == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.OnError)
	}

	retVal := i.p.Call("loadAssetContainer", args...)
	return AssetContainerFromJSObject(retVal, i.ctx)
}

// ISceneLoaderPluginReadFileOpts contains optional parameters for ISceneLoaderPlugin.ReadFile.
type ISceneLoaderPluginReadFileOpts struct {
	OnProgress     func()
	UseArrayBuffer *bool
	OnError        func()
}

// ReadFile calls the ReadFile method on the ISceneLoaderPlugin object.
//
// https://doc.babylonjs.com/api/classes/babylon.isceneloaderplugin#readfile
func (i *ISceneLoaderPlugin) ReadFile(scene *Scene, file js.Value, onSuccess func(), opts *ISceneLoaderPluginReadFileOpts) *IFileRequest {
	if opts == nil {
		opts = &ISceneLoaderPluginReadFileOpts{}
	}

	args := make([]interface{}, 0, 3+3)

	args = append(args, scene.JSObject())
	args = append(args, file)
	args = append(args, js.FuncOf(func(this js.Value, args []js.Value) interface{} { onSuccess(); return nil }))

	if opts.OnProgress == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.OnProgress)
	}
	if opts.UseArrayBuffer == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.UseArrayBuffer)
	}
	if opts.OnError == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.OnError)
	}

	retVal := i.p.Call("readFile", args...)
	return IFileRequestFromJSObject(retVal, i.ctx)
}

// ISceneLoaderPluginRequestFileOpts contains optional parameters for ISceneLoaderPlugin.RequestFile.
type ISceneLoaderPluginRequestFileOpts struct {
	OnProgress     func()
	UseArrayBuffer *bool
	OnError        func()
}

// RequestFile calls the RequestFile method on the ISceneLoaderPlugin object.
//
// https://doc.babylonjs.com/api/classes/babylon.isceneloaderplugin#requestfile
func (i *ISceneLoaderPlugin) RequestFile(scene *Scene, url string, onSuccess func(), opts *ISceneLoaderPluginRequestFileOpts) *IFileRequest {
	if opts == nil {
		opts = &ISceneLoaderPluginRequestFileOpts{}
	}

	args := make([]interface{}, 0, 3+3)

	args = append(args, scene.JSObject())
	args = append(args, url)
	args = append(args, js.FuncOf(func(this js.Value, args []js.Value) interface{} { onSuccess(); return nil }))

	if opts.OnProgress == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.OnProgress)
	}
	if opts.UseArrayBuffer == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.UseArrayBuffer)
	}
	if opts.OnError == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.OnError)
	}

	retVal := i.p.Call("requestFile", args...)
	return IFileRequestFromJSObject(retVal, i.ctx)
}

// ISceneLoaderPluginRewriteRootURLOpts contains optional parameters for ISceneLoaderPlugin.RewriteRootURL.
type ISceneLoaderPluginRewriteRootURLOpts struct {
	ResponseURL *string
}

// RewriteRootURL calls the RewriteRootURL method on the ISceneLoaderPlugin object.
//
// https://doc.babylonjs.com/api/classes/babylon.isceneloaderplugin#rewriterooturl
func (i *ISceneLoaderPlugin) RewriteRootURL(rootUrl string, opts *ISceneLoaderPluginRewriteRootURLOpts) string {
	if opts == nil {
		opts = &ISceneLoaderPluginRewriteRootURLOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, rootUrl)

	if opts.ResponseURL == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.ResponseURL)
	}

	retVal := i.p.Call("rewriteRootURL", args...)
	return retVal.String()
}

/*

// Extensions returns the Extensions property of class ISceneLoaderPlugin.
//
// https://doc.babylonjs.com/api/classes/babylon.isceneloaderplugin#extensions
func (i *ISceneLoaderPlugin) Extensions(extensions string) *ISceneLoaderPlugin {
	p := ba.ctx.Get("ISceneLoaderPlugin").New(extensions)
	return ISceneLoaderPluginFromJSObject(p, ba.ctx)
}

// SetExtensions sets the Extensions property of class ISceneLoaderPlugin.
//
// https://doc.babylonjs.com/api/classes/babylon.isceneloaderplugin#extensions
func (i *ISceneLoaderPlugin) SetExtensions(extensions string) *ISceneLoaderPlugin {
	p := ba.ctx.Get("ISceneLoaderPlugin").New(extensions)
	return ISceneLoaderPluginFromJSObject(p, ba.ctx)
}

// Name returns the Name property of class ISceneLoaderPlugin.
//
// https://doc.babylonjs.com/api/classes/babylon.isceneloaderplugin#name
func (i *ISceneLoaderPlugin) Name(name string) *ISceneLoaderPlugin {
	p := ba.ctx.Get("ISceneLoaderPlugin").New(name)
	return ISceneLoaderPluginFromJSObject(p, ba.ctx)
}

// SetName sets the Name property of class ISceneLoaderPlugin.
//
// https://doc.babylonjs.com/api/classes/babylon.isceneloaderplugin#name
func (i *ISceneLoaderPlugin) SetName(name string) *ISceneLoaderPlugin {
	p := ba.ctx.Get("ISceneLoaderPlugin").New(name)
	return ISceneLoaderPluginFromJSObject(p, ba.ctx)
}

*/
