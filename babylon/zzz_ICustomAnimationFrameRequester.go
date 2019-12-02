// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ICustomAnimationFrameRequester represents a babylon.js ICustomAnimationFrameRequester.
// Interface for any object that can request an animation frame
type ICustomAnimationFrameRequester struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (i *ICustomAnimationFrameRequester) JSObject() js.Value { return i.p }

// ICustomAnimationFrameRequester returns a ICustomAnimationFrameRequester JavaScript class.
func (ba *Babylon) ICustomAnimationFrameRequester() *ICustomAnimationFrameRequester {
	p := ba.ctx.Get("ICustomAnimationFrameRequester")
	return ICustomAnimationFrameRequesterFromJSObject(p, ba.ctx)
}

// ICustomAnimationFrameRequesterFromJSObject returns a wrapped ICustomAnimationFrameRequester JavaScript class.
func ICustomAnimationFrameRequesterFromJSObject(p js.Value, ctx js.Value) *ICustomAnimationFrameRequester {
	return &ICustomAnimationFrameRequester{p: p, ctx: ctx}
}

// ICustomAnimationFrameRequesterArrayToJSArray returns a JavaScript Array for the wrapped array.
func ICustomAnimationFrameRequesterArrayToJSArray(array []*ICustomAnimationFrameRequester) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// RenderFunction returns the RenderFunction property of class ICustomAnimationFrameRequester.
//
// https://doc.babylonjs.com/api/classes/babylon.icustomanimationframerequester#renderfunction
func (i *ICustomAnimationFrameRequester) RenderFunction() js.Value {
	retVal := i.p.Get("renderFunction")
	return retVal
}

// SetRenderFunction sets the RenderFunction property of class ICustomAnimationFrameRequester.
//
// https://doc.babylonjs.com/api/classes/babylon.icustomanimationframerequester#renderfunction
func (i *ICustomAnimationFrameRequester) SetRenderFunction(renderFunction func()) *ICustomAnimationFrameRequester {
	i.p.Set("renderFunction", js.FuncOf(func(this js.Value, args []js.Value) interface{} { renderFunction(); return nil }))
	return i
}

// RequestAnimationFrame returns the RequestAnimationFrame property of class ICustomAnimationFrameRequester.
//
// https://doc.babylonjs.com/api/classes/babylon.icustomanimationframerequester#requestanimationframe
func (i *ICustomAnimationFrameRequester) RequestAnimationFrame() js.Value {
	retVal := i.p.Get("requestAnimationFrame")
	return retVal
}

// SetRequestAnimationFrame sets the RequestAnimationFrame property of class ICustomAnimationFrameRequester.
//
// https://doc.babylonjs.com/api/classes/babylon.icustomanimationframerequester#requestanimationframe
func (i *ICustomAnimationFrameRequester) SetRequestAnimationFrame(requestAnimationFrame func()) *ICustomAnimationFrameRequester {
	i.p.Set("requestAnimationFrame", js.FuncOf(func(this js.Value, args []js.Value) interface{} { requestAnimationFrame(); return nil }))
	return i
}

// RequestID returns the RequestID property of class ICustomAnimationFrameRequester.
//
// https://doc.babylonjs.com/api/classes/babylon.icustomanimationframerequester#requestid
func (i *ICustomAnimationFrameRequester) RequestID() float64 {
	retVal := i.p.Get("requestID")
	return retVal.Float()
}

// SetRequestID sets the RequestID property of class ICustomAnimationFrameRequester.
//
// https://doc.babylonjs.com/api/classes/babylon.icustomanimationframerequester#requestid
func (i *ICustomAnimationFrameRequester) SetRequestID(requestID float64) *ICustomAnimationFrameRequester {
	i.p.Set("requestID", requestID)
	return i
}
