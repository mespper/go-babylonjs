// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ISceneLoaderPluginAsync represents a babylon.js ISceneLoaderPluginAsync.
// Interface used to define an async SceneLoader plugin
type ISceneLoaderPluginAsync struct {
	*ISceneLoaderPluginBase
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (i *ISceneLoaderPluginAsync) JSObject() js.Value { return i.p }

// ISceneLoaderPluginAsync returns a ISceneLoaderPluginAsync JavaScript class.
func (ba *Babylon) ISceneLoaderPluginAsync() *ISceneLoaderPluginAsync {
	p := ba.ctx.Get("ISceneLoaderPluginAsync")
	return ISceneLoaderPluginAsyncFromJSObject(p, ba.ctx)
}

// ISceneLoaderPluginAsyncFromJSObject returns a wrapped ISceneLoaderPluginAsync JavaScript class.
func ISceneLoaderPluginAsyncFromJSObject(p js.Value, ctx js.Value) *ISceneLoaderPluginAsync {
	return &ISceneLoaderPluginAsync{ISceneLoaderPluginBase: ISceneLoaderPluginBaseFromJSObject(p, ctx), ctx: ctx}
}

// ISceneLoaderPluginAsyncArrayToJSArray returns a JavaScript Array for the wrapped array.
func ISceneLoaderPluginAsyncArrayToJSArray(array []*ISceneLoaderPluginAsync) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// CanDirectLoad calls the CanDirectLoad method on the ISceneLoaderPluginAsync object.
//
// https://doc.babylonjs.com/api/classes/babylon.isceneloaderpluginasync#candirectload
func (i *ISceneLoaderPluginAsync) CanDirectLoad(data string) bool {

	args := make([]interface{}, 0, 1+0)

	args = append(args, data)

	retVal := i.p.Call("canDirectLoad", args...)
	return retVal.Bool()
}

// DirectLoad calls the DirectLoad method on the ISceneLoaderPluginAsync object.
//
// https://doc.babylonjs.com/api/classes/babylon.isceneloaderpluginasync#directload
func (i *ISceneLoaderPluginAsync) DirectLoad(scene *Scene, data string) interface{} {

	args := make([]interface{}, 0, 2+0)

	args = append(args, scene.JSObject())
	args = append(args, data)

	retVal := i.p.Call("directLoad", args...)
	return retVal
}

// ISceneLoaderPluginAsyncImportMeshAsyncOpts contains optional parameters for ISceneLoaderPluginAsync.ImportMeshAsync.
type ISceneLoaderPluginAsyncImportMeshAsyncOpts struct {
	OnProgress *func()
	FileName   *string
}

// ImportMeshAsync calls the ImportMeshAsync method on the ISceneLoaderPluginAsync object.
//
// https://doc.babylonjs.com/api/classes/babylon.isceneloaderpluginasync#importmeshasync
func (i *ISceneLoaderPluginAsync) ImportMeshAsync(meshesNames interface{}, scene *Scene, data interface{}, rootUrl string, opts *ISceneLoaderPluginAsyncImportMeshAsyncOpts) *Promise {
	if opts == nil {
		opts = &ISceneLoaderPluginAsyncImportMeshAsyncOpts{}
	}

	args := make([]interface{}, 0, 4+2)

	args = append(args, meshesNames)
	args = append(args, scene.JSObject())
	args = append(args, data)
	args = append(args, rootUrl)

	if opts.OnProgress == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.OnProgress)
	}
	if opts.FileName == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.FileName)
	}

	retVal := i.p.Call("importMeshAsync", args...)
	return PromiseFromJSObject(retVal, i.ctx)
}

// ISceneLoaderPluginAsyncLoadAssetContainerAsyncOpts contains optional parameters for ISceneLoaderPluginAsync.LoadAssetContainerAsync.
type ISceneLoaderPluginAsyncLoadAssetContainerAsyncOpts struct {
	OnProgress *func()
	FileName   *string
}

// LoadAssetContainerAsync calls the LoadAssetContainerAsync method on the ISceneLoaderPluginAsync object.
//
// https://doc.babylonjs.com/api/classes/babylon.isceneloaderpluginasync#loadassetcontainerasync
func (i *ISceneLoaderPluginAsync) LoadAssetContainerAsync(scene *Scene, data interface{}, rootUrl string, opts *ISceneLoaderPluginAsyncLoadAssetContainerAsyncOpts) *Promise {
	if opts == nil {
		opts = &ISceneLoaderPluginAsyncLoadAssetContainerAsyncOpts{}
	}

	args := make([]interface{}, 0, 3+2)

	args = append(args, scene.JSObject())
	args = append(args, data)
	args = append(args, rootUrl)

	if opts.OnProgress == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.OnProgress)
	}
	if opts.FileName == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.FileName)
	}

	retVal := i.p.Call("loadAssetContainerAsync", args...)
	return PromiseFromJSObject(retVal, i.ctx)
}

// ISceneLoaderPluginAsyncLoadAsyncOpts contains optional parameters for ISceneLoaderPluginAsync.LoadAsync.
type ISceneLoaderPluginAsyncLoadAsyncOpts struct {
	OnProgress *func()
	FileName   *string
}

// LoadAsync calls the LoadAsync method on the ISceneLoaderPluginAsync object.
//
// https://doc.babylonjs.com/api/classes/babylon.isceneloaderpluginasync#loadasync
func (i *ISceneLoaderPluginAsync) LoadAsync(scene *Scene, data interface{}, rootUrl string, opts *ISceneLoaderPluginAsyncLoadAsyncOpts) *Promise {
	if opts == nil {
		opts = &ISceneLoaderPluginAsyncLoadAsyncOpts{}
	}

	args := make([]interface{}, 0, 3+2)

	args = append(args, scene.JSObject())
	args = append(args, data)
	args = append(args, rootUrl)

	if opts.OnProgress == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.OnProgress)
	}
	if opts.FileName == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.FileName)
	}

	retVal := i.p.Call("loadAsync", args...)
	return PromiseFromJSObject(retVal, i.ctx)
}

// ISceneLoaderPluginAsyncReadFileOpts contains optional parameters for ISceneLoaderPluginAsync.ReadFile.
type ISceneLoaderPluginAsyncReadFileOpts struct {
	OnProgress     *func()
	UseArrayBuffer *bool
	OnError        *func()
}

// ReadFile calls the ReadFile method on the ISceneLoaderPluginAsync object.
//
// https://doc.babylonjs.com/api/classes/babylon.isceneloaderpluginasync#readfile
func (i *ISceneLoaderPluginAsync) ReadFile(scene *Scene, file js.Value, onSuccess func(), opts *ISceneLoaderPluginAsyncReadFileOpts) *IFileRequest {
	if opts == nil {
		opts = &ISceneLoaderPluginAsyncReadFileOpts{}
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

// ISceneLoaderPluginAsyncRequestFileOpts contains optional parameters for ISceneLoaderPluginAsync.RequestFile.
type ISceneLoaderPluginAsyncRequestFileOpts struct {
	OnProgress     *func()
	UseArrayBuffer *bool
	OnError        *func()
}

// RequestFile calls the RequestFile method on the ISceneLoaderPluginAsync object.
//
// https://doc.babylonjs.com/api/classes/babylon.isceneloaderpluginasync#requestfile
func (i *ISceneLoaderPluginAsync) RequestFile(scene *Scene, url string, onSuccess func(), opts *ISceneLoaderPluginAsyncRequestFileOpts) *IFileRequest {
	if opts == nil {
		opts = &ISceneLoaderPluginAsyncRequestFileOpts{}
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

// ISceneLoaderPluginAsyncRewriteRootURLOpts contains optional parameters for ISceneLoaderPluginAsync.RewriteRootURL.
type ISceneLoaderPluginAsyncRewriteRootURLOpts struct {
	ResponseURL *string
}

// RewriteRootURL calls the RewriteRootURL method on the ISceneLoaderPluginAsync object.
//
// https://doc.babylonjs.com/api/classes/babylon.isceneloaderpluginasync#rewriterooturl
func (i *ISceneLoaderPluginAsync) RewriteRootURL(rootUrl string, opts *ISceneLoaderPluginAsyncRewriteRootURLOpts) string {
	if opts == nil {
		opts = &ISceneLoaderPluginAsyncRewriteRootURLOpts{}
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

// Extensions returns the Extensions property of class ISceneLoaderPluginAsync.
//
// https://doc.babylonjs.com/api/classes/babylon.isceneloaderpluginasync#extensions
func (i *ISceneLoaderPluginAsync) Extensions(extensions string) *ISceneLoaderPluginAsync {
	p := ba.ctx.Get("ISceneLoaderPluginAsync").New(extensions)
	return ISceneLoaderPluginAsyncFromJSObject(p, ba.ctx)
}

// SetExtensions sets the Extensions property of class ISceneLoaderPluginAsync.
//
// https://doc.babylonjs.com/api/classes/babylon.isceneloaderpluginasync#extensions
func (i *ISceneLoaderPluginAsync) SetExtensions(extensions string) *ISceneLoaderPluginAsync {
	p := ba.ctx.Get("ISceneLoaderPluginAsync").New(extensions)
	return ISceneLoaderPluginAsyncFromJSObject(p, ba.ctx)
}

// Name returns the Name property of class ISceneLoaderPluginAsync.
//
// https://doc.babylonjs.com/api/classes/babylon.isceneloaderpluginasync#name
func (i *ISceneLoaderPluginAsync) Name(name string) *ISceneLoaderPluginAsync {
	p := ba.ctx.Get("ISceneLoaderPluginAsync").New(name)
	return ISceneLoaderPluginAsyncFromJSObject(p, ba.ctx)
}

// SetName sets the Name property of class ISceneLoaderPluginAsync.
//
// https://doc.babylonjs.com/api/classes/babylon.isceneloaderpluginasync#name
func (i *ISceneLoaderPluginAsync) SetName(name string) *ISceneLoaderPluginAsync {
	p := ba.ctx.Get("ISceneLoaderPluginAsync").New(name)
	return ISceneLoaderPluginAsyncFromJSObject(p, ba.ctx)
}

*/
