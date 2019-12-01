// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// RenderingManager represents a babylon.js RenderingManager.
// This is the manager responsible of all the rendering for meshes sprites and particles.
// It is enable to manage the different groups as well as the different necessary sort functions.
// This should not be used directly aside of the few static configurations
type RenderingManager struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (r *RenderingManager) JSObject() js.Value { return r.p }

// RenderingManager returns a RenderingManager JavaScript class.
func (ba *Babylon) RenderingManager() *RenderingManager {
	p := ba.ctx.Get("RenderingManager")
	return RenderingManagerFromJSObject(p, ba.ctx)
}

// RenderingManagerFromJSObject returns a wrapped RenderingManager JavaScript class.
func RenderingManagerFromJSObject(p js.Value, ctx js.Value) *RenderingManager {
	return &RenderingManager{p: p, ctx: ctx}
}

// RenderingManagerArrayToJSArray returns a JavaScript Array for the wrapped array.
func RenderingManagerArrayToJSArray(array []*RenderingManager) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewRenderingManager returns a new RenderingManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.renderingmanager
func (ba *Babylon) NewRenderingManager(scene *Scene) *RenderingManager {

	args := make([]interface{}, 0, 1+0)

	args = append(args, scene.JSObject())

	p := ba.ctx.Get("RenderingManager").New(args...)
	return RenderingManagerFromJSObject(p, ba.ctx)
}

// RenderingManagerDispatchOpts contains optional parameters for RenderingManager.Dispatch.
type RenderingManagerDispatchOpts struct {
	Mesh     *AbstractMesh
	Material *Material
}

// Dispatch calls the Dispatch method on the RenderingManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.renderingmanager#dispatch
func (r *RenderingManager) Dispatch(subMesh *SubMesh, opts *RenderingManagerDispatchOpts) {
	if opts == nil {
		opts = &RenderingManagerDispatchOpts{}
	}

	args := make([]interface{}, 0, 1+2)

	args = append(args, subMesh.JSObject())

	if opts.Mesh == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Mesh.JSObject())
	}
	if opts.Material == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Material.JSObject())
	}

	r.p.Call("dispatch", args...)
}

// DispatchParticles calls the DispatchParticles method on the RenderingManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.renderingmanager#dispatchparticles
func (r *RenderingManager) DispatchParticles(particleSystem *IParticleSystem) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, particleSystem.JSObject())

	r.p.Call("dispatchParticles", args...)
}

// DispatchSprites calls the DispatchSprites method on the RenderingManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.renderingmanager#dispatchsprites
func (r *RenderingManager) DispatchSprites(spriteManager *ISpriteManager) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, spriteManager.JSObject())

	r.p.Call("dispatchSprites", args...)
}

// FreeRenderingGroups calls the FreeRenderingGroups method on the RenderingManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.renderingmanager#freerenderinggroups
func (r *RenderingManager) FreeRenderingGroups() {

	r.p.Call("freeRenderingGroups")
}

// GetAutoClearDepthStencilSetup calls the GetAutoClearDepthStencilSetup method on the RenderingManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.renderingmanager#getautocleardepthstencilsetup
func (r *RenderingManager) GetAutoClearDepthStencilSetup(index float64) *IRenderingManagerAutoClearSetup {

	args := make([]interface{}, 0, 1+0)

	args = append(args, index)

	retVal := r.p.Call("getAutoClearDepthStencilSetup", args...)
	return IRenderingManagerAutoClearSetupFromJSObject(retVal, r.ctx)
}

// RenderingManagerSetRenderingAutoClearDepthStencilOpts contains optional parameters for RenderingManager.SetRenderingAutoClearDepthStencil.
type RenderingManagerSetRenderingAutoClearDepthStencilOpts struct {
	Depth   *bool
	Stencil *bool
}

// SetRenderingAutoClearDepthStencil calls the SetRenderingAutoClearDepthStencil method on the RenderingManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.renderingmanager#setrenderingautocleardepthstencil
func (r *RenderingManager) SetRenderingAutoClearDepthStencil(renderingGroupId float64, autoClearDepthStencil bool, opts *RenderingManagerSetRenderingAutoClearDepthStencilOpts) {
	if opts == nil {
		opts = &RenderingManagerSetRenderingAutoClearDepthStencilOpts{}
	}

	args := make([]interface{}, 0, 2+2)

	args = append(args, renderingGroupId)
	args = append(args, autoClearDepthStencil)

	if opts.Depth == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Depth)
	}
	if opts.Stencil == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Stencil)
	}

	r.p.Call("setRenderingAutoClearDepthStencil", args...)
}

// RenderingManagerSetRenderingOrderOpts contains optional parameters for RenderingManager.SetRenderingOrder.
type RenderingManagerSetRenderingOrderOpts struct {
	OpaqueSortCompareFn      func()
	AlphaTestSortCompareFn   func()
	TransparentSortCompareFn func()
}

// SetRenderingOrder calls the SetRenderingOrder method on the RenderingManager object.
//
// https://doc.babylonjs.com/api/classes/babylon.renderingmanager#setrenderingorder
func (r *RenderingManager) SetRenderingOrder(renderingGroupId float64, opts *RenderingManagerSetRenderingOrderOpts) {
	if opts == nil {
		opts = &RenderingManagerSetRenderingOrderOpts{}
	}

	args := make([]interface{}, 0, 1+3)

	args = append(args, renderingGroupId)

	if opts.OpaqueSortCompareFn == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.OpaqueSortCompareFn)
	}
	if opts.AlphaTestSortCompareFn == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.AlphaTestSortCompareFn)
	}
	if opts.TransparentSortCompareFn == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.TransparentSortCompareFn)
	}

	r.p.Call("setRenderingOrder", args...)
}

/*

// AUTOCLEAR returns the AUTOCLEAR property of class RenderingManager.
//
// https://doc.babylonjs.com/api/classes/babylon.renderingmanager#autoclear
func (r *RenderingManager) AUTOCLEAR(AUTOCLEAR bool) *RenderingManager {
	p := ba.ctx.Get("RenderingManager").New(AUTOCLEAR)
	return RenderingManagerFromJSObject(p, ba.ctx)
}

// SetAUTOCLEAR sets the AUTOCLEAR property of class RenderingManager.
//
// https://doc.babylonjs.com/api/classes/babylon.renderingmanager#autoclear
func (r *RenderingManager) SetAUTOCLEAR(AUTOCLEAR bool) *RenderingManager {
	p := ba.ctx.Get("RenderingManager").New(AUTOCLEAR)
	return RenderingManagerFromJSObject(p, ba.ctx)
}

// MAX_RENDERINGGROUPS returns the MAX_RENDERINGGROUPS property of class RenderingManager.
//
// https://doc.babylonjs.com/api/classes/babylon.renderingmanager#max_renderinggroups
func (r *RenderingManager) MAX_RENDERINGGROUPS(MAX_RENDERINGGROUPS float64) *RenderingManager {
	p := ba.ctx.Get("RenderingManager").New(MAX_RENDERINGGROUPS)
	return RenderingManagerFromJSObject(p, ba.ctx)
}

// SetMAX_RENDERINGGROUPS sets the MAX_RENDERINGGROUPS property of class RenderingManager.
//
// https://doc.babylonjs.com/api/classes/babylon.renderingmanager#max_renderinggroups
func (r *RenderingManager) SetMAX_RENDERINGGROUPS(MAX_RENDERINGGROUPS float64) *RenderingManager {
	p := ba.ctx.Get("RenderingManager").New(MAX_RENDERINGGROUPS)
	return RenderingManagerFromJSObject(p, ba.ctx)
}

// MIN_RENDERINGGROUPS returns the MIN_RENDERINGGROUPS property of class RenderingManager.
//
// https://doc.babylonjs.com/api/classes/babylon.renderingmanager#min_renderinggroups
func (r *RenderingManager) MIN_RENDERINGGROUPS(MIN_RENDERINGGROUPS float64) *RenderingManager {
	p := ba.ctx.Get("RenderingManager").New(MIN_RENDERINGGROUPS)
	return RenderingManagerFromJSObject(p, ba.ctx)
}

// SetMIN_RENDERINGGROUPS sets the MIN_RENDERINGGROUPS property of class RenderingManager.
//
// https://doc.babylonjs.com/api/classes/babylon.renderingmanager#min_renderinggroups
func (r *RenderingManager) SetMIN_RENDERINGGROUPS(MIN_RENDERINGGROUPS float64) *RenderingManager {
	p := ba.ctx.Get("RenderingManager").New(MIN_RENDERINGGROUPS)
	return RenderingManagerFromJSObject(p, ba.ctx)
}

*/
