// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// Checkbox represents a babylon.js Checkbox.
// Class used to represent a 2D checkbox
type Checkbox struct{ *Control }

// JSObject returns the underlying js.Value.
func (c *Checkbox) JSObject() js.Value { return c.p }

// Checkbox returns a Checkbox JavaScript class.
func (ba *Babylon) Checkbox() *Checkbox {
	p := ba.ctx.Get("Checkbox")
	return CheckboxFromJSObject(p)
}

// CheckboxFromJSObject returns a wrapped Checkbox JavaScript class.
func CheckboxFromJSObject(p js.Value) *Checkbox {
	return &Checkbox{ControlFromJSObject(p)}
}

// NewCheckboxOpts contains optional parameters for NewCheckbox.
type NewCheckboxOpts struct {
	Name *JSString
}

// NewCheckbox returns a new Checkbox object.
//
// https://doc.babylonjs.com/api/classes/babylon.checkbox
func (ba *Babylon) NewCheckbox(opts *NewCheckboxOpts) *Checkbox {
	if opts == nil {
		opts = &NewCheckboxOpts{}
	}

	p := ba.ctx.Get("Checkbox").New(opts.Name.JSObject())
	return CheckboxFromJSObject(p)
}

// TODO: methods
