// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// PointerInfoBase represents a babylon.js PointerInfoBase.
// Base class of pointer info types.
type PointerInfoBase struct{ p js.Value }

// JSObject returns the underlying js.Value.
func (p *PointerInfoBase) JSObject() js.Value { return p.p }

// PointerInfoBase returns a PointerInfoBase JavaScript class.
func (ba *Babylon) PointerInfoBase() *PointerInfoBase {
	p := ba.ctx.Get("PointerInfoBase")
	return PointerInfoBaseFromJSObject(p)
}

// PointerInfoBaseFromJSObject returns a wrapped PointerInfoBase JavaScript class.
func PointerInfoBaseFromJSObject(p js.Value) *PointerInfoBase {
	return &PointerInfoBase{p: p}
}

// NewPointerInfoBase returns a new PointerInfoBase object.
//
// https://doc.babylonjs.com/api/classes/babylon.pointerinfobase
func (ba *Babylon) NewPointerInfoBase(jsType float64, event js.Value) *PointerInfoBase {
	p := ba.ctx.Get("PointerInfoBase").New(jsType, event)
	return PointerInfoBaseFromJSObject(p)
}

// TODO: methods
