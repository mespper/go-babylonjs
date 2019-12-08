// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// DepthRenderer represents a babylon.js DepthRenderer.
// This represents a depth renderer in Babylon.
// A depth renderer will render to it&#39;s depth map every frame which can be displayed or used in post processing
type DepthRenderer struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (d *DepthRenderer) JSObject() js.Value { return d.p }

// DepthRenderer returns a DepthRenderer JavaScript class.
func (ba *Babylon) DepthRenderer() *DepthRenderer {
	p := ba.ctx.Get("DepthRenderer")
	return DepthRendererFromJSObject(p, ba.ctx)
}

// DepthRendererFromJSObject returns a wrapped DepthRenderer JavaScript class.
func DepthRendererFromJSObject(p js.Value, ctx js.Value) *DepthRenderer {
	return &DepthRenderer{p: p, ctx: ctx}
}

// DepthRendererArrayToJSArray returns a JavaScript Array for the wrapped array.
func DepthRendererArrayToJSArray(array []*DepthRenderer) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewDepthRendererOpts contains optional parameters for NewDepthRenderer.
type NewDepthRendererOpts struct {
	Type                *float64
	Camera              *Camera
	StoreNonLinearDepth *bool
}

// NewDepthRenderer returns a new DepthRenderer object.
//
// https://doc.babylonjs.com/api/classes/babylon.depthrenderer
func (ba *Babylon) NewDepthRenderer(scene *Scene, opts *NewDepthRendererOpts) *DepthRenderer {
	if opts == nil {
		opts = &NewDepthRendererOpts{}
	}

	args := make([]interface{}, 0, 1+3)

	args = append(args, scene.JSObject())

	if opts.Type == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Type)
	}
	if opts.Camera == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Camera.JSObject())
	}
	if opts.StoreNonLinearDepth == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.StoreNonLinearDepth)
	}

	p := ba.ctx.Get("DepthRenderer").New(args...)
	return DepthRendererFromJSObject(p, ba.ctx)
}

// Dispose calls the Dispose method on the DepthRenderer object.
//
// https://doc.babylonjs.com/api/classes/babylon.depthrenderer#dispose
func (d *DepthRenderer) Dispose() {

	d.p.Call("dispose")
}

// GetDepthMap calls the GetDepthMap method on the DepthRenderer object.
//
// https://doc.babylonjs.com/api/classes/babylon.depthrenderer#getdepthmap
func (d *DepthRenderer) GetDepthMap() *RenderTargetTexture {

	retVal := d.p.Call("getDepthMap")
	return RenderTargetTextureFromJSObject(retVal, d.ctx)
}

// IsReady calls the IsReady method on the DepthRenderer object.
//
// https://doc.babylonjs.com/api/classes/babylon.depthrenderer#isready
func (d *DepthRenderer) IsReady(subMesh *SubMesh, useInstances bool) bool {

	args := make([]interface{}, 0, 2+0)

	if subMesh == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, subMesh.JSObject())
	}

	args = append(args, useInstances)

	retVal := d.p.Call("isReady", args...)
	return retVal.Bool()
}

// IsPacked returns the IsPacked property of class DepthRenderer.
//
// https://doc.babylonjs.com/api/classes/babylon.depthrenderer#ispacked
func (d *DepthRenderer) IsPacked() bool {
	retVal := d.p.Get("isPacked")
	return retVal.Bool()
}

// SetIsPacked sets the IsPacked property of class DepthRenderer.
//
// https://doc.babylonjs.com/api/classes/babylon.depthrenderer#ispacked
func (d *DepthRenderer) SetIsPacked(isPacked bool) *DepthRenderer {
	d.p.Set("isPacked", isPacked)
	return d
}

// UseOnlyInActiveCamera returns the UseOnlyInActiveCamera property of class DepthRenderer.
//
// https://doc.babylonjs.com/api/classes/babylon.depthrenderer#useonlyinactivecamera
func (d *DepthRenderer) UseOnlyInActiveCamera() bool {
	retVal := d.p.Get("useOnlyInActiveCamera")
	return retVal.Bool()
}

// SetUseOnlyInActiveCamera sets the UseOnlyInActiveCamera property of class DepthRenderer.
//
// https://doc.babylonjs.com/api/classes/babylon.depthrenderer#useonlyinactivecamera
func (d *DepthRenderer) SetUseOnlyInActiveCamera(useOnlyInActiveCamera bool) *DepthRenderer {
	d.p.Set("useOnlyInActiveCamera", useOnlyInActiveCamera)
	return d
}
