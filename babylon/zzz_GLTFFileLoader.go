// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// GLTFFileLoader represents a babylon.js GLTFFileLoader.
// File loader for loading glTF files into a scene.
type GLTFFileLoader struct{}

// JSObject returns the underlying js.Value.
func (g *GLTFFileLoader) JSObject() js.Value { return g.p }

// GLTFFileLoader returns a GLTFFileLoader JavaScript class.
func (b *Babylon) GLTFFileLoader() *GLTFFileLoader {
	p := b.ctx.Get("GLTFFileLoader")
	return GLTFFileLoaderFromJSObject(p)
}

// GLTFFileLoaderFromJSObject returns a wrapped GLTFFileLoader JavaScript class.
func GLTFFileLoaderFromJSObject(p js.Value) *GLTFFileLoader {
	return &GLTFFileLoader{p: p}
}

// NewGLTFFileLoader returns a new GLTFFileLoader object.
//
// https://doc.babylonjs.com/api/classes/babylon.gltffileloader
func (b *Babylon) NewGLTFFileLoader(todo parameters) *GLTFFileLoader {
	p := b.ctx.Get("GLTFFileLoader").New(todo)
	return GLTFFileLoaderFromJSObject(p)
}

// TODO: methods