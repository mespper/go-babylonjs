// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// SpriteManager represents a babylon.js SpriteManager.
// Class used to manage multiple sprites on the same spritesheet
//
// See: http://doc.babylonjs.com/babylon101/sprites
type SpriteManager struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (s *SpriteManager) JSObject() js.Value { return s.p }

// SpriteManager returns a SpriteManager JavaScript class.
func (ba *Babylon) SpriteManager() *SpriteManager {
	p := ba.ctx.Get("SpriteManager")
	return SpriteManagerFromJSObject(p, ba.ctx)
}

// SpriteManagerFromJSObject returns a wrapped SpriteManager JavaScript class.
func SpriteManagerFromJSObject(p js.Value, ctx js.Value) *SpriteManager {
	return &SpriteManager{p: p, ctx: ctx}
}

// SpriteManagerArrayToJSArray returns a JavaScript Array for the wrapped array.
func SpriteManagerArrayToJSArray(array []*SpriteManager) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewSpriteManagerOpts contains optional parameters for NewSpriteManager.
type NewSpriteManagerOpts struct {
	Epsilon      *float64
	SamplingMode *float64
	FromPacked   *bool
	SpriteJSON   *string
}

// NewSpriteManager returns a new SpriteManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.spritemanager
func (ba *Babylon) NewSpriteManager(name string, imgUrl string, capacity float64, cellSize interface{}, scene *Scene, opts *NewSpriteManagerOpts) *SpriteManager {
	if opts == nil {
		opts = &NewSpriteManagerOpts{}
	}

	args := make([]interface{}, 0, 5+4)

	args = append(args, name)
	args = append(args, imgUrl)
	args = append(args, capacity)
	args = append(args, cellSize)
	args = append(args, scene.JSObject())

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
	if opts.FromPacked == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.FromPacked)
	}
	if opts.SpriteJSON == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.SpriteJSON)
	}

	p := ba.ctx.Get("SpriteManager").New(args...)
	return SpriteManagerFromJSObject(p, ba.ctx)
}

// Dispose calls the Dispose method on the SpriteManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.spritemanager#dispose
func (s *SpriteManager) Dispose() {

	s.p.Call("dispose")
}

// SpriteManagerIntersectsOpts contains optional parameters for SpriteManager.Intersects.
type SpriteManagerIntersectsOpts struct {
	Predicate func()
	FastCheck *bool
}

// Intersects calls the Intersects method on the SpriteManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.spritemanager#intersects
func (s *SpriteManager) Intersects(ray *Ray, camera *Camera, opts *SpriteManagerIntersectsOpts) *PickingInfo {
	if opts == nil {
		opts = &SpriteManagerIntersectsOpts{}
	}

	args := make([]interface{}, 0, 2+2)

	args = append(args, ray.JSObject())
	args = append(args, camera.JSObject())

	if opts.Predicate == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Predicate)
	}
	if opts.FastCheck == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.FastCheck)
	}

	retVal := s.p.Call("intersects", args...)
	return PickingInfoFromJSObject(retVal, s.ctx)
}

// SpriteManagerMultiIntersectsOpts contains optional parameters for SpriteManager.MultiIntersects.
type SpriteManagerMultiIntersectsOpts struct {
	Predicate func()
}

// MultiIntersects calls the MultiIntersects method on the SpriteManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.spritemanager#multiintersects
func (s *SpriteManager) MultiIntersects(ray *Ray, camera *Camera, opts *SpriteManagerMultiIntersectsOpts) *PickingInfo {
	if opts == nil {
		opts = &SpriteManagerMultiIntersectsOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, ray.JSObject())
	args = append(args, camera.JSObject())

	if opts.Predicate == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Predicate)
	}

	retVal := s.p.Call("multiIntersects", args...)
	return PickingInfoFromJSObject(retVal, s.ctx)
}

// Render calls the Render method on the SpriteManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.spritemanager#render
func (s *SpriteManager) Render() {

	s.p.Call("render")
}

/*

// CellHeight returns the CellHeight property of class SpriteManager.
//
// https://doc.babylonjs.com/api/classes/babylon.spritemanager#cellheight
func (s *SpriteManager) CellHeight(cellHeight float64) *SpriteManager {
	p := ba.ctx.Get("SpriteManager").New(cellHeight)
	return SpriteManagerFromJSObject(p, ba.ctx)
}

// SetCellHeight sets the CellHeight property of class SpriteManager.
//
// https://doc.babylonjs.com/api/classes/babylon.spritemanager#cellheight
func (s *SpriteManager) SetCellHeight(cellHeight float64) *SpriteManager {
	p := ba.ctx.Get("SpriteManager").New(cellHeight)
	return SpriteManagerFromJSObject(p, ba.ctx)
}

// CellWidth returns the CellWidth property of class SpriteManager.
//
// https://doc.babylonjs.com/api/classes/babylon.spritemanager#cellwidth
func (s *SpriteManager) CellWidth(cellWidth float64) *SpriteManager {
	p := ba.ctx.Get("SpriteManager").New(cellWidth)
	return SpriteManagerFromJSObject(p, ba.ctx)
}

// SetCellWidth sets the CellWidth property of class SpriteManager.
//
// https://doc.babylonjs.com/api/classes/babylon.spritemanager#cellwidth
func (s *SpriteManager) SetCellWidth(cellWidth float64) *SpriteManager {
	p := ba.ctx.Get("SpriteManager").New(cellWidth)
	return SpriteManagerFromJSObject(p, ba.ctx)
}

// FogEnabled returns the FogEnabled property of class SpriteManager.
//
// https://doc.babylonjs.com/api/classes/babylon.spritemanager#fogenabled
func (s *SpriteManager) FogEnabled(fogEnabled bool) *SpriteManager {
	p := ba.ctx.Get("SpriteManager").New(fogEnabled)
	return SpriteManagerFromJSObject(p, ba.ctx)
}

// SetFogEnabled sets the FogEnabled property of class SpriteManager.
//
// https://doc.babylonjs.com/api/classes/babylon.spritemanager#fogenabled
func (s *SpriteManager) SetFogEnabled(fogEnabled bool) *SpriteManager {
	p := ba.ctx.Get("SpriteManager").New(fogEnabled)
	return SpriteManagerFromJSObject(p, ba.ctx)
}

// IsPickable returns the IsPickable property of class SpriteManager.
//
// https://doc.babylonjs.com/api/classes/babylon.spritemanager#ispickable
func (s *SpriteManager) IsPickable(isPickable bool) *SpriteManager {
	p := ba.ctx.Get("SpriteManager").New(isPickable)
	return SpriteManagerFromJSObject(p, ba.ctx)
}

// SetIsPickable sets the IsPickable property of class SpriteManager.
//
// https://doc.babylonjs.com/api/classes/babylon.spritemanager#ispickable
func (s *SpriteManager) SetIsPickable(isPickable bool) *SpriteManager {
	p := ba.ctx.Get("SpriteManager").New(isPickable)
	return SpriteManagerFromJSObject(p, ba.ctx)
}

// LayerMask returns the LayerMask property of class SpriteManager.
//
// https://doc.babylonjs.com/api/classes/babylon.spritemanager#layermask
func (s *SpriteManager) LayerMask(layerMask float64) *SpriteManager {
	p := ba.ctx.Get("SpriteManager").New(layerMask)
	return SpriteManagerFromJSObject(p, ba.ctx)
}

// SetLayerMask sets the LayerMask property of class SpriteManager.
//
// https://doc.babylonjs.com/api/classes/babylon.spritemanager#layermask
func (s *SpriteManager) SetLayerMask(layerMask float64) *SpriteManager {
	p := ba.ctx.Get("SpriteManager").New(layerMask)
	return SpriteManagerFromJSObject(p, ba.ctx)
}

// Name returns the Name property of class SpriteManager.
//
// https://doc.babylonjs.com/api/classes/babylon.spritemanager#name
func (s *SpriteManager) Name(name string) *SpriteManager {
	p := ba.ctx.Get("SpriteManager").New(name)
	return SpriteManagerFromJSObject(p, ba.ctx)
}

// SetName sets the Name property of class SpriteManager.
//
// https://doc.babylonjs.com/api/classes/babylon.spritemanager#name
func (s *SpriteManager) SetName(name string) *SpriteManager {
	p := ba.ctx.Get("SpriteManager").New(name)
	return SpriteManagerFromJSObject(p, ba.ctx)
}

// OnDispose returns the OnDispose property of class SpriteManager.
//
// https://doc.babylonjs.com/api/classes/babylon.spritemanager#ondispose
func (s *SpriteManager) OnDispose(onDispose func()) *SpriteManager {
	p := ba.ctx.Get("SpriteManager").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {onDispose(); return nil}))
	return SpriteManagerFromJSObject(p, ba.ctx)
}

// SetOnDispose sets the OnDispose property of class SpriteManager.
//
// https://doc.babylonjs.com/api/classes/babylon.spritemanager#ondispose
func (s *SpriteManager) SetOnDispose(onDispose func()) *SpriteManager {
	p := ba.ctx.Get("SpriteManager").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {onDispose(); return nil}))
	return SpriteManagerFromJSObject(p, ba.ctx)
}

// OnDisposeObservable returns the OnDisposeObservable property of class SpriteManager.
//
// https://doc.babylonjs.com/api/classes/babylon.spritemanager#ondisposeobservable
func (s *SpriteManager) OnDisposeObservable(onDisposeObservable *Observable) *SpriteManager {
	p := ba.ctx.Get("SpriteManager").New(onDisposeObservable.JSObject())
	return SpriteManagerFromJSObject(p, ba.ctx)
}

// SetOnDisposeObservable sets the OnDisposeObservable property of class SpriteManager.
//
// https://doc.babylonjs.com/api/classes/babylon.spritemanager#ondisposeobservable
func (s *SpriteManager) SetOnDisposeObservable(onDisposeObservable *Observable) *SpriteManager {
	p := ba.ctx.Get("SpriteManager").New(onDisposeObservable.JSObject())
	return SpriteManagerFromJSObject(p, ba.ctx)
}

// RenderingGroupId returns the RenderingGroupId property of class SpriteManager.
//
// https://doc.babylonjs.com/api/classes/babylon.spritemanager#renderinggroupid
func (s *SpriteManager) RenderingGroupId(renderingGroupId float64) *SpriteManager {
	p := ba.ctx.Get("SpriteManager").New(renderingGroupId)
	return SpriteManagerFromJSObject(p, ba.ctx)
}

// SetRenderingGroupId sets the RenderingGroupId property of class SpriteManager.
//
// https://doc.babylonjs.com/api/classes/babylon.spritemanager#renderinggroupid
func (s *SpriteManager) SetRenderingGroupId(renderingGroupId float64) *SpriteManager {
	p := ba.ctx.Get("SpriteManager").New(renderingGroupId)
	return SpriteManagerFromJSObject(p, ba.ctx)
}

// Sprites returns the Sprites property of class SpriteManager.
//
// https://doc.babylonjs.com/api/classes/babylon.spritemanager#sprites
func (s *SpriteManager) Sprites(sprites *Sprite) *SpriteManager {
	p := ba.ctx.Get("SpriteManager").New(sprites.JSObject())
	return SpriteManagerFromJSObject(p, ba.ctx)
}

// SetSprites sets the Sprites property of class SpriteManager.
//
// https://doc.babylonjs.com/api/classes/babylon.spritemanager#sprites
func (s *SpriteManager) SetSprites(sprites *Sprite) *SpriteManager {
	p := ba.ctx.Get("SpriteManager").New(sprites.JSObject())
	return SpriteManagerFromJSObject(p, ba.ctx)
}

// Texture returns the Texture property of class SpriteManager.
//
// https://doc.babylonjs.com/api/classes/babylon.spritemanager#texture
func (s *SpriteManager) Texture(texture *Texture) *SpriteManager {
	p := ba.ctx.Get("SpriteManager").New(texture.JSObject())
	return SpriteManagerFromJSObject(p, ba.ctx)
}

// SetTexture sets the Texture property of class SpriteManager.
//
// https://doc.babylonjs.com/api/classes/babylon.spritemanager#texture
func (s *SpriteManager) SetTexture(texture *Texture) *SpriteManager {
	p := ba.ctx.Get("SpriteManager").New(texture.JSObject())
	return SpriteManagerFromJSObject(p, ba.ctx)
}

*/
