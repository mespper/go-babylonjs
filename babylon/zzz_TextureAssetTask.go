// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// TextureAssetTask represents a babylon.js TextureAssetTask.
// Define a task used by AssetsManager to load 2D textures
type TextureAssetTask struct{ *AbstractAssetTask }

// JSObject returns the underlying js.Value.
func (t *TextureAssetTask) JSObject() js.Value { return t.p }

// TextureAssetTask returns a TextureAssetTask JavaScript class.
func (b *Babylon) TextureAssetTask() *TextureAssetTask {
	p := b.ctx.Get("TextureAssetTask")
	return TextureAssetTaskFromJSObject(p)
}

// TextureAssetTaskFromJSObject returns a wrapped TextureAssetTask JavaScript class.
func TextureAssetTaskFromJSObject(p js.Value) *TextureAssetTask {
	return &TextureAssetTask{AbstractAssetTaskFromJSObject(p)}
}

// NewTextureAssetTask returns a new TextureAssetTask object.
//
// https://doc.babylonjs.com/api/classes/babylon.textureassettask
func (b *Babylon) NewTextureAssetTask(todo parameters) *TextureAssetTask {
	p := b.ctx.Get("TextureAssetTask").New(todo)
	return TextureAssetTaskFromJSObject(p)
}

// TODO: methods