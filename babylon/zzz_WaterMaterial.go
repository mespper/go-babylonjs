// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// WaterMaterial represents a babylon.js WaterMaterial.
//
type WaterMaterial struct{ p js.Value }

// JSObject returns the underlying js.Value.
func (w *WaterMaterial) JSObject() js.Value { return w.p }

// WaterMaterial returns a WaterMaterial JavaScript class.
func (ba *Babylon) WaterMaterial() *WaterMaterial {
	p := ba.ctx.Get("WaterMaterial")
	return WaterMaterialFromJSObject(p)
}

// WaterMaterialFromJSObject returns a wrapped WaterMaterial JavaScript class.
func WaterMaterialFromJSObject(p js.Value) *WaterMaterial {
	return &WaterMaterial{p: p}
}

// NewWaterMaterialOpts contains optional parameters for NewWaterMaterial.
type NewWaterMaterialOpts struct {
	RenderTargetSize *Vector2
}

// NewWaterMaterial returns a new WaterMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.watermaterial
func (ba *Babylon) NewWaterMaterial(name string, scene *Scene, opts *NewWaterMaterialOpts) *WaterMaterial {
	if opts == nil {
		opts = &NewWaterMaterialOpts{}
	}

	p := ba.ctx.Get("WaterMaterial").New(name, scene.JSObject(), opts.RenderTargetSize.JSObject())
	return WaterMaterialFromJSObject(p)
}

// TODO: methods
