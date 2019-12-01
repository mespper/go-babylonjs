// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// IOfflineProvider represents a babylon.js IOfflineProvider.
// Class used to enable access to offline support
//
// See: http://doc.babylonjs.com/how_to/caching_resources_in_indexeddb
type IOfflineProvider struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (i *IOfflineProvider) JSObject() js.Value { return i.p }

// IOfflineProvider returns a IOfflineProvider JavaScript class.
func (ba *Babylon) IOfflineProvider() *IOfflineProvider {
	p := ba.ctx.Get("IOfflineProvider")
	return IOfflineProviderFromJSObject(p, ba.ctx)
}

// IOfflineProviderFromJSObject returns a wrapped IOfflineProvider JavaScript class.
func IOfflineProviderFromJSObject(p js.Value, ctx js.Value) *IOfflineProvider {
	return &IOfflineProvider{p: p, ctx: ctx}
}

// IOfflineProviderArrayToJSArray returns a JavaScript Array for the wrapped array.
func IOfflineProviderArrayToJSArray(array []*IOfflineProvider) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// IOfflineProviderLoadFileOpts contains optional parameters for IOfflineProvider.LoadFile.
type IOfflineProviderLoadFileOpts struct {
	ProgressCallBack func()
	ErrorCallback    func()
	UseArrayBuffer   *bool
}

// LoadFile calls the LoadFile method on the IOfflineProvider object.
//
// https://doc.babylonjs.com/api/classes/babylon.iofflineprovider#loadfile
func (i *IOfflineProvider) LoadFile(url string, sceneLoaded func(), opts *IOfflineProviderLoadFileOpts) {
	if opts == nil {
		opts = &IOfflineProviderLoadFileOpts{}
	}

	args := make([]interface{}, 0, 2+3)

	args = append(args, url)
	args = append(args, js.FuncOf(func(this js.Value, args []js.Value) interface{} { sceneLoaded(); return nil }))

	if opts.ProgressCallBack == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.ProgressCallBack)
	}
	if opts.ErrorCallback == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.ErrorCallback)
	}
	if opts.UseArrayBuffer == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.UseArrayBuffer)
	}

	i.p.Call("loadFile", args...)
}

// LoadImage calls the LoadImage method on the IOfflineProvider object.
//
// https://doc.babylonjs.com/api/classes/babylon.iofflineprovider#loadimage
func (i *IOfflineProvider) LoadImage(url string, image js.Value) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, url)
	args = append(args, image)

	i.p.Call("loadImage", args...)
}

// Open calls the Open method on the IOfflineProvider object.
//
// https://doc.babylonjs.com/api/classes/babylon.iofflineprovider#open
func (i *IOfflineProvider) Open(successCallback func(), errorCallback func()) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, js.FuncOf(func(this js.Value, args []js.Value) interface{} { successCallback(); return nil }))
	args = append(args, js.FuncOf(func(this js.Value, args []js.Value) interface{} { errorCallback(); return nil }))

	i.p.Call("open", args...)
}

/*

// EnableSceneOffline returns the EnableSceneOffline property of class IOfflineProvider.
//
// https://doc.babylonjs.com/api/classes/babylon.iofflineprovider#enablesceneoffline
func (i *IOfflineProvider) EnableSceneOffline(enableSceneOffline bool) *IOfflineProvider {
	p := ba.ctx.Get("IOfflineProvider").New(enableSceneOffline)
	return IOfflineProviderFromJSObject(p, ba.ctx)
}

// SetEnableSceneOffline sets the EnableSceneOffline property of class IOfflineProvider.
//
// https://doc.babylonjs.com/api/classes/babylon.iofflineprovider#enablesceneoffline
func (i *IOfflineProvider) SetEnableSceneOffline(enableSceneOffline bool) *IOfflineProvider {
	p := ba.ctx.Get("IOfflineProvider").New(enableSceneOffline)
	return IOfflineProviderFromJSObject(p, ba.ctx)
}

// EnableTexturesOffline returns the EnableTexturesOffline property of class IOfflineProvider.
//
// https://doc.babylonjs.com/api/classes/babylon.iofflineprovider#enabletexturesoffline
func (i *IOfflineProvider) EnableTexturesOffline(enableTexturesOffline bool) *IOfflineProvider {
	p := ba.ctx.Get("IOfflineProvider").New(enableTexturesOffline)
	return IOfflineProviderFromJSObject(p, ba.ctx)
}

// SetEnableTexturesOffline sets the EnableTexturesOffline property of class IOfflineProvider.
//
// https://doc.babylonjs.com/api/classes/babylon.iofflineprovider#enabletexturesoffline
func (i *IOfflineProvider) SetEnableTexturesOffline(enableTexturesOffline bool) *IOfflineProvider {
	p := ba.ctx.Get("IOfflineProvider").New(enableTexturesOffline)
	return IOfflineProviderFromJSObject(p, ba.ctx)
}

*/
