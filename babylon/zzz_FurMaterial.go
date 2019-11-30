// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// FurMaterial represents a babylon.js FurMaterial.
//
type FurMaterial struct{ p js.Value }

// JSObject returns the underlying js.Value.
func (f *FurMaterial) JSObject() js.Value { return f.p }

// FurMaterial returns a FurMaterial JavaScript class.
func (ba *Babylon) FurMaterial() *FurMaterial {
	p := ba.ctx.Get("FurMaterial")
	return FurMaterialFromJSObject(p)
}

// FurMaterialFromJSObject returns a wrapped FurMaterial JavaScript class.
func FurMaterialFromJSObject(p js.Value) *FurMaterial {
	return &FurMaterial{p: p}
}

// NewFurMaterial returns a new FurMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.furmaterial
func (ba *Babylon) NewFurMaterial(name string, scene *Scene) *FurMaterial {
	p := ba.ctx.Get("FurMaterial").New(name, scene.JSObject())
	return FurMaterialFromJSObject(p)
}

// TODO: methods
