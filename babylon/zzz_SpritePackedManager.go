// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// SpritePackedManager represents a babylon.js SpritePackedManager.
// Class used to manage multiple sprites of different sizes on the same spritesheet
//
// See: http://doc.babylonjs.com/babylon101/sprites
type SpritePackedManager struct {
	*SpriteManager
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (s *SpritePackedManager) JSObject() js.Value { return s.p }

// SpritePackedManager returns a SpritePackedManager JavaScript class.
func (ba *Babylon) SpritePackedManager() *SpritePackedManager {
	p := ba.ctx.Get("SpritePackedManager")
	return SpritePackedManagerFromJSObject(p, ba.ctx)
}

// SpritePackedManagerFromJSObject returns a wrapped SpritePackedManager JavaScript class.
func SpritePackedManagerFromJSObject(p js.Value, ctx js.Value) *SpritePackedManager {
	return &SpritePackedManager{SpriteManager: SpriteManagerFromJSObject(p, ctx), ctx: ctx}
}

// SpritePackedManagerArrayToJSArray returns a JavaScript Array for the wrapped array.
func SpritePackedManagerArrayToJSArray(array []*SpritePackedManager) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewSpritePackedManagerOpts contains optional parameters for NewSpritePackedManager.
type NewSpritePackedManagerOpts struct {
	SpriteJSON   *string
	Epsilon      *float64
	SamplingMode *float64
}

// NewSpritePackedManager returns a new SpritePackedManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.spritepackedmanager
func (ba *Babylon) NewSpritePackedManager(name string, imgUrl string, capacity float64, scene *Scene, opts *NewSpritePackedManagerOpts) *SpritePackedManager {
	if opts == nil {
		opts = &NewSpritePackedManagerOpts{}
	}

	args := make([]interface{}, 0, 4+3)

	args = append(args, name)
	args = append(args, imgUrl)
	args = append(args, capacity)
	args = append(args, scene.JSObject())

	if opts.SpriteJSON == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.SpriteJSON)
	}
	if opts.Epsilon == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Epsilon)
	}
	if opts.SamplingMode == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.SamplingMode)
	}

	p := ba.ctx.Get("SpritePackedManager").New(args...)
	return SpritePackedManagerFromJSObject(p, ba.ctx)
}

// Name returns the Name property of class SpritePackedManager.
//
// https://doc.babylonjs.com/api/classes/babylon.spritepackedmanager#name
func (s *SpritePackedManager) Name() string {
	retVal := s.p.Get("name")
	return retVal.String()
}

// SetName sets the Name property of class SpritePackedManager.
//
// https://doc.babylonjs.com/api/classes/babylon.spritepackedmanager#name
func (s *SpritePackedManager) SetName(name string) *SpritePackedManager {
	s.p.Set("name", name)
	return s
}
