// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// AssetsManager represents a babylon.js AssetsManager.
// This class can be used to easily import assets into a scene
//
// See: http://doc.babylonjs.com/how_to/how_to_use_assetsmanager
type AssetsManager struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (a *AssetsManager) JSObject() js.Value { return a.p }

// AssetsManager returns a AssetsManager JavaScript class.
func (ba *Babylon) AssetsManager() *AssetsManager {
	p := ba.ctx.Get("AssetsManager")
	return AssetsManagerFromJSObject(p, ba.ctx)
}

// AssetsManagerFromJSObject returns a wrapped AssetsManager JavaScript class.
func AssetsManagerFromJSObject(p js.Value, ctx js.Value) *AssetsManager {
	return &AssetsManager{p: p, ctx: ctx}
}

// AssetsManagerArrayToJSArray returns a JavaScript Array for the wrapped array.
func AssetsManagerArrayToJSArray(array []*AssetsManager) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewAssetsManager returns a new AssetsManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.assetsmanager
func (ba *Babylon) NewAssetsManager(scene *Scene) *AssetsManager {

	args := make([]interface{}, 0, 1+0)

	args = append(args, scene.JSObject())

	p := ba.ctx.Get("AssetsManager").New(args...)
	return AssetsManagerFromJSObject(p, ba.ctx)
}

// AddBinaryFileTask calls the AddBinaryFileTask method on the AssetsManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.assetsmanager#addbinaryfiletask
func (a *AssetsManager) AddBinaryFileTask(taskName string, url string) *BinaryFileAssetTask {

	args := make([]interface{}, 0, 2+0)

	args = append(args, taskName)

	args = append(args, url)

	retVal := a.p.Call("addBinaryFileTask", args...)
	return BinaryFileAssetTaskFromJSObject(retVal, a.ctx)
}

// AssetsManagerAddCubeTextureTaskOpts contains optional parameters for AssetsManager.AddCubeTextureTask.
type AssetsManagerAddCubeTextureTaskOpts struct {
	Extensions []string
	NoMipmap   *bool
	Files      []string
}

// AddCubeTextureTask calls the AddCubeTextureTask method on the AssetsManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.assetsmanager#addcubetexturetask
func (a *AssetsManager) AddCubeTextureTask(taskName string, url string, opts *AssetsManagerAddCubeTextureTaskOpts) *CubeTextureAssetTask {
	if opts == nil {
		opts = &AssetsManagerAddCubeTextureTaskOpts{}
	}

	args := make([]interface{}, 0, 2+3)

	args = append(args, taskName)

	args = append(args, url)

	if opts.Extensions == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Extensions)
	}
	if opts.NoMipmap == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.NoMipmap)
	}
	if opts.Files == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Files)
	}

	retVal := a.p.Call("addCubeTextureTask", args...)
	return CubeTextureAssetTaskFromJSObject(retVal, a.ctx)
}

// AssetsManagerAddEquiRectangularCubeTextureAssetTaskOpts contains optional parameters for AssetsManager.AddEquiRectangularCubeTextureAssetTask.
type AssetsManagerAddEquiRectangularCubeTextureAssetTaskOpts struct {
	NoMipmap   *bool
	GammaSpace *bool
}

// AddEquiRectangularCubeTextureAssetTask calls the AddEquiRectangularCubeTextureAssetTask method on the AssetsManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.assetsmanager#addequirectangularcubetextureassettask
func (a *AssetsManager) AddEquiRectangularCubeTextureAssetTask(taskName string, url string, size float64, opts *AssetsManagerAddEquiRectangularCubeTextureAssetTaskOpts) *EquiRectangularCubeTextureAssetTask {
	if opts == nil {
		opts = &AssetsManagerAddEquiRectangularCubeTextureAssetTaskOpts{}
	}

	args := make([]interface{}, 0, 3+2)

	args = append(args, taskName)

	args = append(args, url)

	args = append(args, size)

	if opts.NoMipmap == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.NoMipmap)
	}
	if opts.GammaSpace == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.GammaSpace)
	}

	retVal := a.p.Call("addEquiRectangularCubeTextureAssetTask", args...)
	return EquiRectangularCubeTextureAssetTaskFromJSObject(retVal, a.ctx)
}

// AssetsManagerAddHDRCubeTextureTaskOpts contains optional parameters for AssetsManager.AddHDRCubeTextureTask.
type AssetsManagerAddHDRCubeTextureTaskOpts struct {
	NoMipmap          *bool
	GenerateHarmonics *bool
	GammaSpace        *bool
	Reserved          *bool
}

// AddHDRCubeTextureTask calls the AddHDRCubeTextureTask method on the AssetsManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.assetsmanager#addhdrcubetexturetask
func (a *AssetsManager) AddHDRCubeTextureTask(taskName string, url string, size float64, opts *AssetsManagerAddHDRCubeTextureTaskOpts) *HDRCubeTextureAssetTask {
	if opts == nil {
		opts = &AssetsManagerAddHDRCubeTextureTaskOpts{}
	}

	args := make([]interface{}, 0, 3+4)

	args = append(args, taskName)

	args = append(args, url)

	args = append(args, size)

	if opts.NoMipmap == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.NoMipmap)
	}
	if opts.GenerateHarmonics == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.GenerateHarmonics)
	}
	if opts.GammaSpace == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.GammaSpace)
	}
	if opts.Reserved == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Reserved)
	}

	retVal := a.p.Call("addHDRCubeTextureTask", args...)
	return HDRCubeTextureAssetTaskFromJSObject(retVal, a.ctx)
}

// AddImageTask calls the AddImageTask method on the AssetsManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.assetsmanager#addimagetask
func (a *AssetsManager) AddImageTask(taskName string, url string) *ImageAssetTask {

	args := make([]interface{}, 0, 2+0)

	args = append(args, taskName)

	args = append(args, url)

	retVal := a.p.Call("addImageTask", args...)
	return ImageAssetTaskFromJSObject(retVal, a.ctx)
}

// AddMeshTask calls the AddMeshTask method on the AssetsManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.assetsmanager#addmeshtask
func (a *AssetsManager) AddMeshTask(taskName string, meshesNames JSObject, rootUrl string, sceneFilename string) *MeshAssetTask {

	args := make([]interface{}, 0, 4+0)

	args = append(args, taskName)

	if meshesNames == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, meshesNames.JSObject())
	}

	args = append(args, rootUrl)

	args = append(args, sceneFilename)

	retVal := a.p.Call("addMeshTask", args...)
	return MeshAssetTaskFromJSObject(retVal, a.ctx)
}

// AddTextFileTask calls the AddTextFileTask method on the AssetsManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.assetsmanager#addtextfiletask
func (a *AssetsManager) AddTextFileTask(taskName string, url string) *TextFileAssetTask {

	args := make([]interface{}, 0, 2+0)

	args = append(args, taskName)

	args = append(args, url)

	retVal := a.p.Call("addTextFileTask", args...)
	return TextFileAssetTaskFromJSObject(retVal, a.ctx)
}

// AssetsManagerAddTextureTaskOpts contains optional parameters for AssetsManager.AddTextureTask.
type AssetsManagerAddTextureTaskOpts struct {
	NoMipmap     *bool
	InvertY      *bool
	SamplingMode *float64
}

// AddTextureTask calls the AddTextureTask method on the AssetsManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.assetsmanager#addtexturetask
func (a *AssetsManager) AddTextureTask(taskName string, url string, opts *AssetsManagerAddTextureTaskOpts) *TextureAssetTask {
	if opts == nil {
		opts = &AssetsManagerAddTextureTaskOpts{}
	}

	args := make([]interface{}, 0, 2+3)

	args = append(args, taskName)

	args = append(args, url)

	if opts.NoMipmap == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.NoMipmap)
	}
	if opts.InvertY == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.InvertY)
	}
	if opts.SamplingMode == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.SamplingMode)
	}

	retVal := a.p.Call("addTextureTask", args...)
	return TextureAssetTaskFromJSObject(retVal, a.ctx)
}

// Load calls the Load method on the AssetsManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.assetsmanager#load
func (a *AssetsManager) Load() *AssetsManager {

	retVal := a.p.Call("load")
	return AssetsManagerFromJSObject(retVal, a.ctx)
}

// LoadAsync calls the LoadAsync method on the AssetsManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.assetsmanager#loadasync
func (a *AssetsManager) LoadAsync() *Promise {

	retVal := a.p.Call("loadAsync")
	return PromiseFromJSObject(retVal, a.ctx)
}

// RemoveTask calls the RemoveTask method on the AssetsManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.assetsmanager#removetask
func (a *AssetsManager) RemoveTask(task *AbstractAssetTask) {

	args := make([]interface{}, 0, 1+0)

	if task == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, task.JSObject())
	}

	a.p.Call("removeTask", args...)
}

// Reset calls the Reset method on the AssetsManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.assetsmanager#reset
func (a *AssetsManager) Reset() *AssetsManager {

	retVal := a.p.Call("reset")
	return AssetsManagerFromJSObject(retVal, a.ctx)
}

// AutoHideLoadingUI returns the AutoHideLoadingUI property of class AssetsManager.
//
// https://doc.babylonjs.com/api/classes/babylon.assetsmanager#autohideloadingui
func (a *AssetsManager) AutoHideLoadingUI() bool {
	retVal := a.p.Get("autoHideLoadingUI")
	return retVal.Bool()
}

// SetAutoHideLoadingUI sets the AutoHideLoadingUI property of class AssetsManager.
//
// https://doc.babylonjs.com/api/classes/babylon.assetsmanager#autohideloadingui
func (a *AssetsManager) SetAutoHideLoadingUI(autoHideLoadingUI bool) *AssetsManager {
	a.p.Set("autoHideLoadingUI", autoHideLoadingUI)
	return a
}

// OnFinish returns the OnFinish property of class AssetsManager.
//
// https://doc.babylonjs.com/api/classes/babylon.assetsmanager#onfinish
func (a *AssetsManager) OnFinish() js.Value {
	retVal := a.p.Get("onFinish")
	return retVal
}

// SetOnFinish sets the OnFinish property of class AssetsManager.
//
// https://doc.babylonjs.com/api/classes/babylon.assetsmanager#onfinish
func (a *AssetsManager) SetOnFinish(onFinish JSFunc) *AssetsManager {
	a.p.Set("onFinish", js.FuncOf(onFinish))
	return a
}

// OnProgress returns the OnProgress property of class AssetsManager.
//
// https://doc.babylonjs.com/api/classes/babylon.assetsmanager#onprogress
func (a *AssetsManager) OnProgress() js.Value {
	retVal := a.p.Get("onProgress")
	return retVal
}

// SetOnProgress sets the OnProgress property of class AssetsManager.
//
// https://doc.babylonjs.com/api/classes/babylon.assetsmanager#onprogress
func (a *AssetsManager) SetOnProgress(onProgress JSFunc) *AssetsManager {
	a.p.Set("onProgress", js.FuncOf(onProgress))
	return a
}

// OnProgressObservable returns the OnProgressObservable property of class AssetsManager.
//
// https://doc.babylonjs.com/api/classes/babylon.assetsmanager#onprogressobservable
func (a *AssetsManager) OnProgressObservable() *Observable {
	retVal := a.p.Get("onProgressObservable")
	return ObservableFromJSObject(retVal, a.ctx)
}

// SetOnProgressObservable sets the OnProgressObservable property of class AssetsManager.
//
// https://doc.babylonjs.com/api/classes/babylon.assetsmanager#onprogressobservable
func (a *AssetsManager) SetOnProgressObservable(onProgressObservable *Observable) *AssetsManager {
	a.p.Set("onProgressObservable", onProgressObservable.JSObject())
	return a
}

// OnTaskError returns the OnTaskError property of class AssetsManager.
//
// https://doc.babylonjs.com/api/classes/babylon.assetsmanager#ontaskerror
func (a *AssetsManager) OnTaskError() js.Value {
	retVal := a.p.Get("onTaskError")
	return retVal
}

// SetOnTaskError sets the OnTaskError property of class AssetsManager.
//
// https://doc.babylonjs.com/api/classes/babylon.assetsmanager#ontaskerror
func (a *AssetsManager) SetOnTaskError(onTaskError JSFunc) *AssetsManager {
	a.p.Set("onTaskError", js.FuncOf(onTaskError))
	return a
}

// OnTaskErrorObservable returns the OnTaskErrorObservable property of class AssetsManager.
//
// https://doc.babylonjs.com/api/classes/babylon.assetsmanager#ontaskerrorobservable
func (a *AssetsManager) OnTaskErrorObservable() *Observable {
	retVal := a.p.Get("onTaskErrorObservable")
	return ObservableFromJSObject(retVal, a.ctx)
}

// SetOnTaskErrorObservable sets the OnTaskErrorObservable property of class AssetsManager.
//
// https://doc.babylonjs.com/api/classes/babylon.assetsmanager#ontaskerrorobservable
func (a *AssetsManager) SetOnTaskErrorObservable(onTaskErrorObservable *Observable) *AssetsManager {
	a.p.Set("onTaskErrorObservable", onTaskErrorObservable.JSObject())
	return a
}

// OnTaskSuccess returns the OnTaskSuccess property of class AssetsManager.
//
// https://doc.babylonjs.com/api/classes/babylon.assetsmanager#ontasksuccess
func (a *AssetsManager) OnTaskSuccess() js.Value {
	retVal := a.p.Get("onTaskSuccess")
	return retVal
}

// SetOnTaskSuccess sets the OnTaskSuccess property of class AssetsManager.
//
// https://doc.babylonjs.com/api/classes/babylon.assetsmanager#ontasksuccess
func (a *AssetsManager) SetOnTaskSuccess(onTaskSuccess JSFunc) *AssetsManager {
	a.p.Set("onTaskSuccess", js.FuncOf(onTaskSuccess))
	return a
}

// OnTaskSuccessObservable returns the OnTaskSuccessObservable property of class AssetsManager.
//
// https://doc.babylonjs.com/api/classes/babylon.assetsmanager#ontasksuccessobservable
func (a *AssetsManager) OnTaskSuccessObservable() *Observable {
	retVal := a.p.Get("onTaskSuccessObservable")
	return ObservableFromJSObject(retVal, a.ctx)
}

// SetOnTaskSuccessObservable sets the OnTaskSuccessObservable property of class AssetsManager.
//
// https://doc.babylonjs.com/api/classes/babylon.assetsmanager#ontasksuccessobservable
func (a *AssetsManager) SetOnTaskSuccessObservable(onTaskSuccessObservable *Observable) *AssetsManager {
	a.p.Set("onTaskSuccessObservable", onTaskSuccessObservable.JSObject())
	return a
}

// OnTasksDoneObservable returns the OnTasksDoneObservable property of class AssetsManager.
//
// https://doc.babylonjs.com/api/classes/babylon.assetsmanager#ontasksdoneobservable
func (a *AssetsManager) OnTasksDoneObservable() []*Observable {
	retVal := a.p.Get("onTasksDoneObservable")
	result := []*Observable{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, ObservableFromJSObject(retVal.Index(ri), a.ctx))
	}
	return result
}

// SetOnTasksDoneObservable sets the OnTasksDoneObservable property of class AssetsManager.
//
// https://doc.babylonjs.com/api/classes/babylon.assetsmanager#ontasksdoneobservable
func (a *AssetsManager) SetOnTasksDoneObservable(onTasksDoneObservable []*Observable) *AssetsManager {
	a.p.Set("onTasksDoneObservable", onTasksDoneObservable)
	return a
}

// UseDefaultLoadingScreen returns the UseDefaultLoadingScreen property of class AssetsManager.
//
// https://doc.babylonjs.com/api/classes/babylon.assetsmanager#usedefaultloadingscreen
func (a *AssetsManager) UseDefaultLoadingScreen() bool {
	retVal := a.p.Get("useDefaultLoadingScreen")
	return retVal.Bool()
}

// SetUseDefaultLoadingScreen sets the UseDefaultLoadingScreen property of class AssetsManager.
//
// https://doc.babylonjs.com/api/classes/babylon.assetsmanager#usedefaultloadingscreen
func (a *AssetsManager) SetUseDefaultLoadingScreen(useDefaultLoadingScreen bool) *AssetsManager {
	a.p.Set("useDefaultLoadingScreen", useDefaultLoadingScreen)
	return a
}
