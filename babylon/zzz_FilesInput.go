// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// FilesInput represents a babylon.js FilesInput.
// Class used to help managing file picking and drag&#39;n&#39;drop
type FilesInput struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (f *FilesInput) JSObject() js.Value { return f.p }

// FilesInput returns a FilesInput JavaScript class.
func (ba *Babylon) FilesInput() *FilesInput {
	p := ba.ctx.Get("FilesInput")
	return FilesInputFromJSObject(p, ba.ctx)
}

// FilesInputFromJSObject returns a wrapped FilesInput JavaScript class.
func FilesInputFromJSObject(p js.Value, ctx js.Value) *FilesInput {
	return &FilesInput{p: p, ctx: ctx}
}

// FilesInputArrayToJSArray returns a JavaScript Array for the wrapped array.
func FilesInputArrayToJSArray(array []*FilesInput) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewFilesInput returns a new FilesInput object.
//
// https://doc.babylonjs.com/api/classes/babylon.filesinput
func (ba *Babylon) NewFilesInput(engine *Engine, scene *Scene, sceneLoadedCallback func(), progressCallback func(), additionalRenderLoopLogicCallback func(), textureLoadingCallback func(), startingProcessingFilesCallback func(), onReloadCallback func(), errorCallback func()) *FilesInput {

	args := make([]interface{}, 0, 9+0)

	args = append(args, engine.JSObject())
	args = append(args, scene.JSObject())
	args = append(args, js.FuncOf(func(this js.Value, args []js.Value) interface{} { sceneLoadedCallback(); return nil }))
	args = append(args, js.FuncOf(func(this js.Value, args []js.Value) interface{} { progressCallback(); return nil }))
	args = append(args, js.FuncOf(func(this js.Value, args []js.Value) interface{} { additionalRenderLoopLogicCallback(); return nil }))
	args = append(args, js.FuncOf(func(this js.Value, args []js.Value) interface{} { textureLoadingCallback(); return nil }))
	args = append(args, js.FuncOf(func(this js.Value, args []js.Value) interface{} { startingProcessingFilesCallback(); return nil }))
	args = append(args, js.FuncOf(func(this js.Value, args []js.Value) interface{} { onReloadCallback(); return nil }))
	args = append(args, js.FuncOf(func(this js.Value, args []js.Value) interface{} { errorCallback(); return nil }))

	p := ba.ctx.Get("FilesInput").New(args...)
	return FilesInputFromJSObject(p, ba.ctx)
}

// Dispose calls the Dispose method on the FilesInput object.
//
// https://doc.babylonjs.com/api/classes/babylon.filesinput#dispose
func (f *FilesInput) Dispose() {

	f.p.Call("dispose")
}

// LoadFiles calls the LoadFiles method on the FilesInput object.
//
// https://doc.babylonjs.com/api/classes/babylon.filesinput#loadfiles
func (f *FilesInput) LoadFiles(event interface{}) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, event)

	f.p.Call("loadFiles", args...)
}

// MonitorElementForDragNDrop calls the MonitorElementForDragNDrop method on the FilesInput object.
//
// https://doc.babylonjs.com/api/classes/babylon.filesinput#monitorelementfordragndrop
func (f *FilesInput) MonitorElementForDragNDrop(elementToMonitor js.Value) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, elementToMonitor)

	f.p.Call("monitorElementForDragNDrop", args...)
}

// Reload calls the Reload method on the FilesInput object.
//
// https://doc.babylonjs.com/api/classes/babylon.filesinput#reload
func (f *FilesInput) Reload() {

	f.p.Call("reload")
}

/*

// FilesToLoad returns the FilesToLoad property of class FilesInput.
//
// https://doc.babylonjs.com/api/classes/babylon.filesinput#filestoload
func (f *FilesInput) FilesToLoad(FilesToLoad js.Value) *FilesInput {
	p := ba.ctx.Get("FilesInput").New(FilesToLoad)
	return FilesInputFromJSObject(p, ba.ctx)
}

// SetFilesToLoad sets the FilesToLoad property of class FilesInput.
//
// https://doc.babylonjs.com/api/classes/babylon.filesinput#filestoload
func (f *FilesInput) SetFilesToLoad(FilesToLoad js.Value) *FilesInput {
	p := ba.ctx.Get("FilesInput").New(FilesToLoad)
	return FilesInputFromJSObject(p, ba.ctx)
}

// OnProcessFileCallback returns the OnProcessFileCallback property of class FilesInput.
//
// https://doc.babylonjs.com/api/classes/babylon.filesinput#onprocessfilecallback
func (f *FilesInput) OnProcessFileCallback(onProcessFileCallback func()) *FilesInput {
	p := ba.ctx.Get("FilesInput").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {onProcessFileCallback(); return nil}))
	return FilesInputFromJSObject(p, ba.ctx)
}

// SetOnProcessFileCallback sets the OnProcessFileCallback property of class FilesInput.
//
// https://doc.babylonjs.com/api/classes/babylon.filesinput#onprocessfilecallback
func (f *FilesInput) SetOnProcessFileCallback(onProcessFileCallback func()) *FilesInput {
	p := ba.ctx.Get("FilesInput").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {onProcessFileCallback(); return nil}))
	return FilesInputFromJSObject(p, ba.ctx)
}

*/
