// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// CubeTextureAssetTask represents a babylon.js CubeTextureAssetTask.
// Define a task used by AssetsManager to load cube textures
type CubeTextureAssetTask struct{ *AbstractAssetTask }

// JSObject returns the underlying js.Value.
func (c *CubeTextureAssetTask) JSObject() js.Value { return c.p }

// CubeTextureAssetTask returns a CubeTextureAssetTask JavaScript class.
func (b *Babylon) CubeTextureAssetTask() *CubeTextureAssetTask {
	p := b.ctx.Get("CubeTextureAssetTask")
	return CubeTextureAssetTaskFromJSObject(p)
}

// CubeTextureAssetTaskFromJSObject returns a wrapped CubeTextureAssetTask JavaScript class.
func CubeTextureAssetTaskFromJSObject(p js.Value) *CubeTextureAssetTask {
	return &CubeTextureAssetTask{AbstractAssetTaskFromJSObject(p)}
}

// NewCubeTextureAssetTask returns a new CubeTextureAssetTask object.
//
// https://doc.babylonjs.com/api/classes/babylon.cubetextureassettask
func (b *Babylon) NewCubeTextureAssetTask(todo parameters) *CubeTextureAssetTask {
	p := b.ctx.Get("CubeTextureAssetTask").New(todo)
	return CubeTextureAssetTaskFromJSObject(p)
}

// TODO: methods