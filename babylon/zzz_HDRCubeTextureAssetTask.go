// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// HDRCubeTextureAssetTask represents a babylon.js HDRCubeTextureAssetTask.
// Define a task used by AssetsManager to load HDR cube textures
type HDRCubeTextureAssetTask struct{ *AbstractAssetTask }

// JSObject returns the underlying js.Value.
func (h *HDRCubeTextureAssetTask) JSObject() js.Value { return h.p }

// HDRCubeTextureAssetTask returns a HDRCubeTextureAssetTask JavaScript class.
func (ba *Babylon) HDRCubeTextureAssetTask() *HDRCubeTextureAssetTask {
	p := ba.ctx.Get("HDRCubeTextureAssetTask")
	return HDRCubeTextureAssetTaskFromJSObject(p)
}

// HDRCubeTextureAssetTaskFromJSObject returns a wrapped HDRCubeTextureAssetTask JavaScript class.
func HDRCubeTextureAssetTaskFromJSObject(p js.Value) *HDRCubeTextureAssetTask {
	return &HDRCubeTextureAssetTask{AbstractAssetTaskFromJSObject(p)}
}

// NewHDRCubeTextureAssetTaskOpts contains optional parameters for NewHDRCubeTextureAssetTask.
type NewHDRCubeTextureAssetTaskOpts struct {
	NoMipmap *JSBool

	GenerateHarmonics *JSBool

	GammaSpace *JSBool

	Reserved *JSBool
}

// NewHDRCubeTextureAssetTask returns a new HDRCubeTextureAssetTask object.
//
// https://doc.babylonjs.com/api/classes/babylon.hdrcubetextureassettask
func (ba *Babylon) NewHDRCubeTextureAssetTask(name string, url string, size float64, opts *NewHDRCubeTextureAssetTaskOpts) *HDRCubeTextureAssetTask {
	if opts == nil {
		opts = &NewHDRCubeTextureAssetTaskOpts{}
	}

	p := ba.ctx.Get("HDRCubeTextureAssetTask").New(name, url, size, opts.NoMipmap.JSObject(), opts.GenerateHarmonics.JSObject(), opts.GammaSpace.JSObject(), opts.Reserved.JSObject())
	return HDRCubeTextureAssetTaskFromJSObject(p)
}

// TODO: methods
