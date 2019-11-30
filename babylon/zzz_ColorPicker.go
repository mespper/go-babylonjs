// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ColorPicker represents a babylon.js ColorPicker.
// Class used to create color pickers
type ColorPicker struct{ *Control }

// JSObject returns the underlying js.Value.
func (c *ColorPicker) JSObject() js.Value { return c.p }

// ColorPicker returns a ColorPicker JavaScript class.
func (ba *Babylon) ColorPicker() *ColorPicker {
	p := ba.ctx.Get("ColorPicker")
	return ColorPickerFromJSObject(p)
}

// ColorPickerFromJSObject returns a wrapped ColorPicker JavaScript class.
func ColorPickerFromJSObject(p js.Value) *ColorPicker {
	return &ColorPicker{ControlFromJSObject(p)}
}

// NewColorPickerOpts contains optional parameters for NewColorPicker.
type NewColorPickerOpts struct {
	Name *JSString
}

// NewColorPicker returns a new ColorPicker object.
//
// https://doc.babylonjs.com/api/classes/babylon.colorpicker
func (ba *Babylon) NewColorPicker(opts *NewColorPickerOpts) *ColorPicker {
	if opts == nil {
		opts = &NewColorPickerOpts{}
	}

	p := ba.ctx.Get("ColorPicker").New(opts.Name)
	return ColorPickerFromJSObject(p)
}

// TODO: methods
