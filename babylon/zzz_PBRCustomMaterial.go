// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// PBRCustomMaterial represents a babylon.js PBRCustomMaterial.
//
type PBRCustomMaterial struct{ *PBRMaterial }

// JSObject returns the underlying js.Value.
func (p *PBRCustomMaterial) JSObject() js.Value { return p.p }

// PBRCustomMaterial returns a PBRCustomMaterial JavaScript class.
func (ba *Babylon) PBRCustomMaterial() *PBRCustomMaterial {
	p := ba.ctx.Get("PBRCustomMaterial")
	return PBRCustomMaterialFromJSObject(p)
}

// PBRCustomMaterialFromJSObject returns a wrapped PBRCustomMaterial JavaScript class.
func PBRCustomMaterialFromJSObject(p js.Value) *PBRCustomMaterial {
	return &PBRCustomMaterial{PBRMaterialFromJSObject(p)}
}

// NewPBRCustomMaterial returns a new PBRCustomMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrcustommaterial
func (ba *Babylon) NewPBRCustomMaterial(name string, scene *Scene) *PBRCustomMaterial {
	p := ba.ctx.Get("PBRCustomMaterial").New(name, scene.JSObject())
	return PBRCustomMaterialFromJSObject(p)
}

// TODO: methods
