// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// BezierCurveEase represents a babylon.js BezierCurveEase.
// Easing function with a bezier shape (see link below).
//
// See: http://doc.babylonjs.com/babylon101/animations#easing-functions
type BezierCurveEase struct{ *EasingFunction }

// JSObject returns the underlying js.Value.
func (b *BezierCurveEase) JSObject() js.Value { return b.p }

// BezierCurveEase returns a BezierCurveEase JavaScript class.
func (ba *Babylon) BezierCurveEase() *BezierCurveEase {
	p := ba.ctx.Get("BezierCurveEase")
	return BezierCurveEaseFromJSObject(p)
}

// BezierCurveEaseFromJSObject returns a wrapped BezierCurveEase JavaScript class.
func BezierCurveEaseFromJSObject(p js.Value) *BezierCurveEase {
	return &BezierCurveEase{EasingFunctionFromJSObject(p)}
}

// NewBezierCurveEaseOpts contains optional parameters for NewBezierCurveEase.
type NewBezierCurveEaseOpts struct {
	X1 *JSFloat64

	Y1 *JSFloat64

	X2 *JSFloat64

	Y2 *JSFloat64
}

// NewBezierCurveEase returns a new BezierCurveEase object.
//
// https://doc.babylonjs.com/api/classes/babylon.beziercurveease
func (ba *Babylon) NewBezierCurveEase(opts *NewBezierCurveEaseOpts) *BezierCurveEase {
	if opts == nil {
		opts = &NewBezierCurveEaseOpts{}
	}

	p := ba.ctx.Get("BezierCurveEase").New(opts.X1.JSObject(), opts.Y1.JSObject(), opts.X2.JSObject(), opts.Y2.JSObject())
	return BezierCurveEaseFromJSObject(p)
}

// TODO: methods
