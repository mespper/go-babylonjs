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
func (b *Babylon) HDRCubeTextureAssetTask() *HDRCubeTextureAssetTask {
	p := b.ctx.Get("HDRCubeTextureAssetTask")
	return HDRCubeTextureAssetTaskFromJSObject(p)
}

// HDRCubeTextureAssetTaskFromJSObject returns a wrapped HDRCubeTextureAssetTask JavaScript class.
func HDRCubeTextureAssetTaskFromJSObject(p js.Value) *HDRCubeTextureAssetTask {
	return &HDRCubeTextureAssetTask{AbstractAssetTaskFromJSObject(p)}
}

// NewHDRCubeTextureAssetTask returns a new HDRCubeTextureAssetTask object.
//
// https://doc.babylonjs.com/api/classes/babylon.hdrcubetextureassettask
func (b *Babylon) NewHDRCubeTextureAssetTask(todo parameters) *HDRCubeTextureAssetTask {
	p := b.ctx.Get("HDRCubeTextureAssetTask").New(todo)
	return HDRCubeTextureAssetTaskFromJSObject(p)
}

// TODO: methods