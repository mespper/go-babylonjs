// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// TextureOptimization represents a babylon.js TextureOptimization.
// Defines an optimization used to reduce the size of render target textures
//
// See: http://doc.babylonjs.com/how_to/how_to_use_sceneoptimizer
type TextureOptimization struct{ *SceneOptimization }

// JSObject returns the underlying js.Value.
func (t *TextureOptimization) JSObject() js.Value { return t.p }

// TextureOptimization returns a TextureOptimization JavaScript class.
func (ba *Babylon) TextureOptimization() *TextureOptimization {
	p := ba.ctx.Get("TextureOptimization")
	return TextureOptimizationFromJSObject(p)
}

// TextureOptimizationFromJSObject returns a wrapped TextureOptimization JavaScript class.
func TextureOptimizationFromJSObject(p js.Value) *TextureOptimization {
	return &TextureOptimization{SceneOptimizationFromJSObject(p)}
}

// NewTextureOptimizationOpts contains optional parameters for NewTextureOptimization.
type NewTextureOptimizationOpts struct {
	Priority *JSFloat64

	MaximumSize *JSFloat64

	Step *JSFloat64
}

// NewTextureOptimization returns a new TextureOptimization object.
//
// https://doc.babylonjs.com/api/classes/babylon.textureoptimization
func (ba *Babylon) NewTextureOptimization(opts *NewTextureOptimizationOpts) *TextureOptimization {
	if opts == nil {
		opts = &NewTextureOptimizationOpts{}
	}

	p := ba.ctx.Get("TextureOptimization").New(opts.Priority, opts.MaximumSize, opts.Step)
	return TextureOptimizationFromJSObject(p)
}

// TODO: methods
