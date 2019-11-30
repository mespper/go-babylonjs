// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// Arc2 represents a babylon.js Arc2.
// This represents an arc in a 2d space.
type Arc2 struct{ p js.Value }

// JSObject returns the underlying js.Value.
func (a *Arc2) JSObject() js.Value { return a.p }

// Arc2 returns a Arc2 JavaScript class.
func (ba *Babylon) Arc2() *Arc2 {
	p := ba.ctx.Get("Arc2")
	return Arc2FromJSObject(p)
}

// Arc2FromJSObject returns a wrapped Arc2 JavaScript class.
func Arc2FromJSObject(p js.Value) *Arc2 {
	return &Arc2{p: p}
}

// NewArc2 returns a new Arc2 object.
//
// https://doc.babylonjs.com/api/classes/babylon.arc2
func (ba *Babylon) NewArc2(startPoint *Vector2, midPoint *Vector2, endPoint *Vector2) *Arc2 {
	p := ba.ctx.Get("Arc2").New(startPoint.JSObject(), midPoint.JSObject(), endPoint.JSObject())
	return Arc2FromJSObject(p)
}

// TODO: methods
