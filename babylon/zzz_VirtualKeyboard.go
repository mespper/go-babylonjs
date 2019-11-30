// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// VirtualKeyboard represents a babylon.js VirtualKeyboard.
// Class used to create virtual keyboard
type VirtualKeyboard struct{ *StackPanel }

// JSObject returns the underlying js.Value.
func (v *VirtualKeyboard) JSObject() js.Value { return v.p }

// VirtualKeyboard returns a VirtualKeyboard JavaScript class.
func (ba *Babylon) VirtualKeyboard() *VirtualKeyboard {
	p := ba.ctx.Get("VirtualKeyboard")
	return VirtualKeyboardFromJSObject(p)
}

// VirtualKeyboardFromJSObject returns a wrapped VirtualKeyboard JavaScript class.
func VirtualKeyboardFromJSObject(p js.Value) *VirtualKeyboard {
	return &VirtualKeyboard{StackPanelFromJSObject(p)}
}

// NewVirtualKeyboardOpts contains optional parameters for NewVirtualKeyboard.
type NewVirtualKeyboardOpts struct {
	Name *JSString
}

// NewVirtualKeyboard returns a new VirtualKeyboard object.
//
// https://doc.babylonjs.com/api/classes/babylon.virtualkeyboard
func (ba *Babylon) NewVirtualKeyboard(opts *NewVirtualKeyboardOpts) *VirtualKeyboard {
	if opts == nil {
		opts = &NewVirtualKeyboardOpts{}
	}

	p := ba.ctx.Get("VirtualKeyboard").New(opts.Name.JSObject())
	return VirtualKeyboardFromJSObject(p)
}

// TODO: methods
