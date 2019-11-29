// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// TextFileAssetTask represents a babylon.js TextFileAssetTask.
// Define a task used by AssetsManager to load text content
type TextFileAssetTask struct{ *AbstractAssetTask }

// JSObject returns the underlying js.Value.
func (t *TextFileAssetTask) JSObject() js.Value { return t.p }

// TextFileAssetTask returns a TextFileAssetTask JavaScript class.
func (b *Babylon) TextFileAssetTask() *TextFileAssetTask {
	p := b.ctx.Get("TextFileAssetTask")
	return TextFileAssetTaskFromJSObject(p)
}

// TextFileAssetTaskFromJSObject returns a wrapped TextFileAssetTask JavaScript class.
func TextFileAssetTaskFromJSObject(p js.Value) *TextFileAssetTask {
	return &TextFileAssetTask{AbstractAssetTaskFromJSObject(p)}
}

// NewTextFileAssetTask returns a new TextFileAssetTask object.
//
// https://doc.babylonjs.com/api/classes/babylon.textfileassettask
func (b *Babylon) NewTextFileAssetTask(todo parameters) *TextFileAssetTask {
	p := b.ctx.Get("TextFileAssetTask").New(todo)
	return TextFileAssetTaskFromJSObject(p)
}

// TODO: methods