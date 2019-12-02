// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ISpriteManager represents a babylon.js ISpriteManager.
// Defines the minimum interface to fullfil in order to be a sprite manager.
type ISpriteManager struct {
	*IDisposable
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (i *ISpriteManager) JSObject() js.Value { return i.p }

// ISpriteManager returns a ISpriteManager JavaScript class.
func (ba *Babylon) ISpriteManager() *ISpriteManager {
	p := ba.ctx.Get("ISpriteManager")
	return ISpriteManagerFromJSObject(p, ba.ctx)
}

// ISpriteManagerFromJSObject returns a wrapped ISpriteManager JavaScript class.
func ISpriteManagerFromJSObject(p js.Value, ctx js.Value) *ISpriteManager {
	return &ISpriteManager{IDisposable: IDisposableFromJSObject(p, ctx), ctx: ctx}
}

// ISpriteManagerArrayToJSArray returns a JavaScript Array for the wrapped array.
func ISpriteManagerArrayToJSArray(array []*ISpriteManager) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// ISpriteManagerIntersectsOpts contains optional parameters for ISpriteManager.Intersects.
type ISpriteManagerIntersectsOpts struct {
	Predicate func()
	FastCheck *bool
}

// Intersects calls the Intersects method on the ISpriteManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.ispritemanager#intersects
func (i *ISpriteManager) Intersects(ray *Ray, camera *Camera, opts *ISpriteManagerIntersectsOpts) *PickingInfo {
	if opts == nil {
		opts = &ISpriteManagerIntersectsOpts{}
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

	retVal := i.p.Call("intersects", args...)
	return PickingInfoFromJSObject(retVal, i.ctx)
}

// ISpriteManagerMultiIntersectsOpts contains optional parameters for ISpriteManager.MultiIntersects.
type ISpriteManagerMultiIntersectsOpts struct {
	Predicate func()
}

// MultiIntersects calls the MultiIntersects method on the ISpriteManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.ispritemanager#multiintersects
func (i *ISpriteManager) MultiIntersects(ray *Ray, camera *Camera, opts *ISpriteManagerMultiIntersectsOpts) *PickingInfo {
	if opts == nil {
		opts = &ISpriteManagerMultiIntersectsOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, ray.JSObject())
	args = append(args, camera.JSObject())

	if opts.Predicate == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Predicate)
	}

	retVal := i.p.Call("multiIntersects", args...)
	return PickingInfoFromJSObject(retVal, i.ctx)
}

// Render calls the Render method on the ISpriteManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.ispritemanager#render
func (i *ISpriteManager) Render() {

	i.p.Call("render")
}

// IsPickable returns the IsPickable property of class ISpriteManager.
//
// https://doc.babylonjs.com/api/classes/babylon.ispritemanager#ispickable
func (i *ISpriteManager) IsPickable() bool {
	retVal := i.p.Get("isPickable")
	return retVal.Bool()
}

// SetIsPickable sets the IsPickable property of class ISpriteManager.
//
// https://doc.babylonjs.com/api/classes/babylon.ispritemanager#ispickable
func (i *ISpriteManager) SetIsPickable(isPickable bool) *ISpriteManager {
	i.p.Set("isPickable", isPickable)
	return i
}

// LayerMask returns the LayerMask property of class ISpriteManager.
//
// https://doc.babylonjs.com/api/classes/babylon.ispritemanager#layermask
func (i *ISpriteManager) LayerMask() float64 {
	retVal := i.p.Get("layerMask")
	return retVal.Float()
}

// SetLayerMask sets the LayerMask property of class ISpriteManager.
//
// https://doc.babylonjs.com/api/classes/babylon.ispritemanager#layermask
func (i *ISpriteManager) SetLayerMask(layerMask float64) *ISpriteManager {
	i.p.Set("layerMask", layerMask)
	return i
}

// RenderingGroupId returns the RenderingGroupId property of class ISpriteManager.
//
// https://doc.babylonjs.com/api/classes/babylon.ispritemanager#renderinggroupid
func (i *ISpriteManager) RenderingGroupId() float64 {
	retVal := i.p.Get("renderingGroupId")
	return retVal.Float()
}

// SetRenderingGroupId sets the RenderingGroupId property of class ISpriteManager.
//
// https://doc.babylonjs.com/api/classes/babylon.ispritemanager#renderinggroupid
func (i *ISpriteManager) SetRenderingGroupId(renderingGroupId float64) *ISpriteManager {
	i.p.Set("renderingGroupId", renderingGroupId)
	return i
}

// Sprites returns the Sprites property of class ISpriteManager.
//
// https://doc.babylonjs.com/api/classes/babylon.ispritemanager#sprites
func (i *ISpriteManager) Sprites() []*Sprite {
	retVal := i.p.Get("sprites")
	result := []*Sprite{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, SpriteFromJSObject(retVal.Index(ri), i.ctx))
	}
	return result
}

// SetSprites sets the Sprites property of class ISpriteManager.
//
// https://doc.babylonjs.com/api/classes/babylon.ispritemanager#sprites
func (i *ISpriteManager) SetSprites(sprites []*Sprite) *ISpriteManager {
	i.p.Set("sprites", sprites)
	return i
}
