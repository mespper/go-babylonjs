// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// Viewport represents a babylon.js Viewport.
// Class used to represent a viewport on screen
type Viewport struct{ p js.Value }

// JSObject returns the underlying js.Value.
func (v *Viewport) JSObject() js.Value { return v.p }

// Viewport returns a Viewport JavaScript class.
func (ba *Babylon) Viewport() *Viewport {
	p := ba.ctx.Get("Viewport")
	return ViewportFromJSObject(p)
}

// ViewportFromJSObject returns a wrapped Viewport JavaScript class.
func ViewportFromJSObject(p js.Value) *Viewport {
	return &Viewport{p: p}
}

// NewViewport returns a new Viewport object.
//
// https://doc.babylonjs.com/api/classes/babylon.viewport
func (ba *Babylon) NewViewport(x float64, y float64, width float64, height float64) *Viewport {
	p := ba.ctx.Get("Viewport").New(x, y, width, height)
	return ViewportFromJSObject(p)
}

// TODO: methods
