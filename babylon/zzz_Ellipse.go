// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// Ellipse represents a babylon.js Ellipse.
// Class used to create 2D ellipse containers
type Ellipse struct{ *Container }

// JSObject returns the underlying js.Value.
func (e *Ellipse) JSObject() js.Value { return e.p }

// Ellipse returns a Ellipse JavaScript class.
func (ba *Babylon) Ellipse() *Ellipse {
	p := ba.ctx.Get("Ellipse")
	return EllipseFromJSObject(p)
}

// EllipseFromJSObject returns a wrapped Ellipse JavaScript class.
func EllipseFromJSObject(p js.Value) *Ellipse {
	return &Ellipse{ContainerFromJSObject(p)}
}

// NewEllipseOpts contains optional parameters for NewEllipse.
type NewEllipseOpts struct {
	Name *JSString
}

// NewEllipse returns a new Ellipse object.
//
// https://doc.babylonjs.com/api/classes/babylon.ellipse
func (ba *Babylon) NewEllipse(opts *NewEllipseOpts) *Ellipse {
	if opts == nil {
		opts = &NewEllipseOpts{}
	}

	p := ba.ctx.Get("Ellipse").New(opts.Name)
	return EllipseFromJSObject(p)
}

// TODO: methods
