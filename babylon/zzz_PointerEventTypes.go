// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// PointerEventTypes represents a babylon.js PointerEventTypes.
// Gather the list of pointer event types as constants.
type PointerEventTypes struct{ p js.Value }

// JSObject returns the underlying js.Value.
func (p *PointerEventTypes) JSObject() js.Value { return p.p }

// PointerEventTypes returns a PointerEventTypes JavaScript class.
func (ba *Babylon) PointerEventTypes() *PointerEventTypes {
	p := ba.ctx.Get("PointerEventTypes")
	return PointerEventTypesFromJSObject(p)
}

// PointerEventTypesFromJSObject returns a wrapped PointerEventTypes JavaScript class.
func PointerEventTypesFromJSObject(p js.Value) *PointerEventTypes {
	return &PointerEventTypes{p: p}
}

// TODO: methods
