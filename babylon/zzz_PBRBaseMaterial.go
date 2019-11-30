// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// PBRBaseMaterial represents a babylon.js PBRBaseMaterial.
// The Physically based material base class of BJS.
//
// This offers the main features of a standard PBR material.
// For more information, please refer to the documentation :
// &lt;a href=&#34;https://doc.babylonjs.com/how_to/physically_based_rendering&#34;&gt;https://doc.babylonjs.com/how_to/physically_based_rendering&lt;/a&gt;
type PBRBaseMaterial struct{ p js.Value }

// JSObject returns the underlying js.Value.
func (p *PBRBaseMaterial) JSObject() js.Value { return p.p }

// PBRBaseMaterial returns a PBRBaseMaterial JavaScript class.
func (ba *Babylon) PBRBaseMaterial() *PBRBaseMaterial {
	p := ba.ctx.Get("PBRBaseMaterial")
	return PBRBaseMaterialFromJSObject(p)
}

// PBRBaseMaterialFromJSObject returns a wrapped PBRBaseMaterial JavaScript class.
func PBRBaseMaterialFromJSObject(p js.Value) *PBRBaseMaterial {
	return &PBRBaseMaterial{p: p}
}

// NewPBRBaseMaterial returns a new PBRBaseMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.pbrbasematerial
func (ba *Babylon) NewPBRBaseMaterial(name string, scene *Scene) *PBRBaseMaterial {
	p := ba.ctx.Get("PBRBaseMaterial").New(name, scene.JSObject())
	return PBRBaseMaterialFromJSObject(p)
}

// TODO: methods
