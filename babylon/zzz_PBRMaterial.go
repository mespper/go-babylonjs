// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// PBRMaterial represents a babylon.js PBRMaterial.
// The Physically based material of BJS.

//
// This offers the main features of a standard PBR material.
// For more information, please refer to the documentation :
// &lt;a href=&#34;https://doc.babylonjs.com/how_to/physically_based_rendering&#34;&gt;https://doc.babylonjs.com/how_to/physically_based_rendering&lt;/a&gt;
type PBRMaterial struct{ *PBRBaseMaterial }

// JSObject returns the underlying js.Value.
func (p *PBRMaterial) JSObject() js.Value { return p.p }

// PBRMaterial returns a PBRMaterial JavaScript class.
func (b *Babylon) PBRMaterial() *PBRMaterial {
	p := b.ctx.Get("PBRMaterial")
	return PBRMaterialFromJSObject(p)
}

// PBRMaterialFromJSObject returns a wrapped PBRMaterial JavaScript class.
func PBRMaterialFromJSObject(p js.Value) *PBRMaterial {
	return &PBRMaterial{PBRBaseMaterialFromJSObject(p)}
}

// NewPBRMaterial returns a new PBRMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrmaterial
func (b *Babylon) NewPBRMaterial(todo parameters) *PBRMaterial {
	p := b.ctx.Get("PBRMaterial").New(todo)
	return PBRMaterialFromJSObject(p)
}

// TODO: methods