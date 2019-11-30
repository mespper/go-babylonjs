// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// DomManagement represents a babylon.js DomManagement.
// Sets of helpers dealing with the DOM and some of the recurrent functions needed in
// Babylon.js
type DomManagement struct{ p js.Value }

// JSObject returns the underlying js.Value.
func (d *DomManagement) JSObject() js.Value { return d.p }

// DomManagement returns a DomManagement JavaScript class.
func (ba *Babylon) DomManagement() *DomManagement {
	p := ba.ctx.Get("DomManagement")
	return DomManagementFromJSObject(p)
}

// DomManagementFromJSObject returns a wrapped DomManagement JavaScript class.
func DomManagementFromJSObject(p js.Value) *DomManagement {
	return &DomManagement{p: p}
}

// TODO: methods
