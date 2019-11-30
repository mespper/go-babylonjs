// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// LavaMaterial represents a babylon.js LavaMaterial.
//
type LavaMaterial struct{ p js.Value }

// JSObject returns the underlying js.Value.
func (l *LavaMaterial) JSObject() js.Value { return l.p }

// LavaMaterial returns a LavaMaterial JavaScript class.
func (ba *Babylon) LavaMaterial() *LavaMaterial {
	p := ba.ctx.Get("LavaMaterial")
	return LavaMaterialFromJSObject(p)
}

// LavaMaterialFromJSObject returns a wrapped LavaMaterial JavaScript class.
func LavaMaterialFromJSObject(p js.Value) *LavaMaterial {
	return &LavaMaterial{p: p}
}

// NewLavaMaterial returns a new LavaMaterial object.
//
// https://doc.babylonjs.com/api/classes/babylon.lavamaterial
func (ba *Babylon) NewLavaMaterial(name string, scene *Scene) *LavaMaterial {
	p := ba.ctx.Get("LavaMaterial").New(name, scene.JSObject())
	return LavaMaterialFromJSObject(p)
}

// TODO: methods
