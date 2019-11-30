// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// Control3D represents a babylon.js Control3D.
// Class used as base class for controls
type Control3D struct{ p js.Value }

// JSObject returns the underlying js.Value.
func (c *Control3D) JSObject() js.Value { return c.p }

// Control3D returns a Control3D JavaScript class.
func (ba *Babylon) Control3D() *Control3D {
	p := ba.ctx.Get("Control3D")
	return Control3DFromJSObject(p)
}

// Control3DFromJSObject returns a wrapped Control3D JavaScript class.
func Control3DFromJSObject(p js.Value) *Control3D {
	return &Control3D{p: p}
}

// NewControl3DOpts contains optional parameters for NewControl3D.
type NewControl3DOpts struct {
	Name *JSString
}

// NewControl3D returns a new Control3D object.
//
// https://doc.babylonjs.com/api/classes/babylon.control3d
func (ba *Babylon) NewControl3D(opts *NewControl3DOpts) *Control3D {
	if opts == nil {
		opts = &NewControl3DOpts{}
	}

	p := ba.ctx.Get("Control3D").New(opts.Name.JSObject())
	return Control3DFromJSObject(p)
}

// TODO: methods
