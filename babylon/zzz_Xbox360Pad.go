// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// Xbox360Pad represents a babylon.js Xbox360Pad.
// Defines a XBox360 gamepad
type Xbox360Pad struct{ *Gamepad }

// JSObject returns the underlying js.Value.
func (x *Xbox360Pad) JSObject() js.Value { return x.p }

// Xbox360Pad returns a Xbox360Pad JavaScript class.
func (ba *Babylon) Xbox360Pad() *Xbox360Pad {
	p := ba.ctx.Get("Xbox360Pad")
	return Xbox360PadFromJSObject(p)
}

// Xbox360PadFromJSObject returns a wrapped Xbox360Pad JavaScript class.
func Xbox360PadFromJSObject(p js.Value) *Xbox360Pad {
	return &Xbox360Pad{GamepadFromJSObject(p)}
}

// NewXbox360PadOpts contains optional parameters for NewXbox360Pad.
type NewXbox360PadOpts struct {
	XboxOne *JSBool
}

// NewXbox360Pad returns a new Xbox360Pad object.
//
// https://doc.babylonjs.com/api/classes/babylon.xbox360pad
func (ba *Babylon) NewXbox360Pad(id string, index float64, gamepad interface{}, opts *NewXbox360PadOpts) *Xbox360Pad {
	if opts == nil {
		opts = &NewXbox360PadOpts{}
	}

	p := ba.ctx.Get("Xbox360Pad").New(id, index, gamepad, opts.XboxOne)
	return Xbox360PadFromJSObject(p)
}

// TODO: methods
