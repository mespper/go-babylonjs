// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// IViewportOwnerLike represents a babylon.js IViewportOwnerLike.
// Defines the interface used by objects containing a viewport (like a camera)
type IViewportOwnerLike struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (i *IViewportOwnerLike) JSObject() js.Value { return i.p }

// IViewportOwnerLike returns a IViewportOwnerLike JavaScript class.
func (ba *Babylon) IViewportOwnerLike() *IViewportOwnerLike {
	p := ba.ctx.Get("IViewportOwnerLike")
	return IViewportOwnerLikeFromJSObject(p, ba.ctx)
}

// IViewportOwnerLikeFromJSObject returns a wrapped IViewportOwnerLike JavaScript class.
func IViewportOwnerLikeFromJSObject(p js.Value, ctx js.Value) *IViewportOwnerLike {
	return &IViewportOwnerLike{p: p, ctx: ctx}
}

// IViewportOwnerLikeArrayToJSArray returns a JavaScript Array for the wrapped array.
func IViewportOwnerLikeArrayToJSArray(array []*IViewportOwnerLike) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// Viewport returns the Viewport property of class IViewportOwnerLike.
//
// https://doc.babylonjs.com/api/classes/babylon.iviewportownerlike#viewport
func (i *IViewportOwnerLike) Viewport() js.Value {
	retVal := i.p.Get("viewport")
	return retVal
}

// SetViewport sets the Viewport property of class IViewportOwnerLike.
//
// https://doc.babylonjs.com/api/classes/babylon.iviewportownerlike#viewport
func (i *IViewportOwnerLike) SetViewport(viewport js.Value) *IViewportOwnerLike {
	i.p.Set("viewport", viewport)
	return i
}
