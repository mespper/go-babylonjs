// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// SpriteManager represents a babylon.js SpriteManager.
// Class used to manage multiple sprites on the same spritesheet
//
// See: http://doc.babylonjs.com/babylon101/sprites
type SpriteManager struct{ p js.Value }

// JSObject returns the underlying js.Value.
func (s *SpriteManager) JSObject() js.Value { return s.p }

// SpriteManager returns a SpriteManager JavaScript class.
func (ba *Babylon) SpriteManager() *SpriteManager {
	p := ba.ctx.Get("SpriteManager")
	return SpriteManagerFromJSObject(p)
}

// SpriteManagerFromJSObject returns a wrapped SpriteManager JavaScript class.
func SpriteManagerFromJSObject(p js.Value) *SpriteManager {
	return &SpriteManager{p: p}
}

// NewSpriteManagerOpts contains optional parameters for NewSpriteManager.
type NewSpriteManagerOpts struct {
	Epsilon *JSFloat64

	SamplingMode *JSFloat64

	FromPacked *JSBool

	SpriteJSON *JSString
}

// NewSpriteManager returns a new SpriteManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.spritemanager
func (ba *Babylon) NewSpriteManager(name string, imgUrl string, capacity float64, cellSize interface{}, scene *Scene, opts *NewSpriteManagerOpts) *SpriteManager {
	if opts == nil {
		opts = &NewSpriteManagerOpts{}
	}

	p := ba.ctx.Get("SpriteManager").New(name, imgUrl, capacity, cellSize, scene.JSObject(), opts.Epsilon, opts.SamplingMode, opts.FromPacked, opts.SpriteJSON)
	return SpriteManagerFromJSObject(p)
}

// TODO: methods
