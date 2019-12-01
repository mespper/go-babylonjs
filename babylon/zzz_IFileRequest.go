// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// IFileRequest represents a babylon.js IFileRequest.
// File request interface
type IFileRequest struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (i *IFileRequest) JSObject() js.Value { return i.p }

// IFileRequest returns a IFileRequest JavaScript class.
func (ba *Babylon) IFileRequest() *IFileRequest {
	p := ba.ctx.Get("IFileRequest")
	return IFileRequestFromJSObject(p, ba.ctx)
}

// IFileRequestFromJSObject returns a wrapped IFileRequest JavaScript class.
func IFileRequestFromJSObject(p js.Value, ctx js.Value) *IFileRequest {
	return &IFileRequest{p: p, ctx: ctx}
}

// IFileRequestArrayToJSArray returns a JavaScript Array for the wrapped array.
func IFileRequestArrayToJSArray(array []*IFileRequest) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

/*

// Abort returns the Abort property of class IFileRequest.
//
// https://doc.babylonjs.com/api/classes/babylon.ifilerequest#abort
func (i *IFileRequest) Abort(abort func()) *IFileRequest {
	p := ba.ctx.Get("IFileRequest").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {abort(); return nil}))
	return IFileRequestFromJSObject(p, ba.ctx)
}

// SetAbort sets the Abort property of class IFileRequest.
//
// https://doc.babylonjs.com/api/classes/babylon.ifilerequest#abort
func (i *IFileRequest) SetAbort(abort func()) *IFileRequest {
	p := ba.ctx.Get("IFileRequest").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {abort(); return nil}))
	return IFileRequestFromJSObject(p, ba.ctx)
}

// OnCompleteObservable returns the OnCompleteObservable property of class IFileRequest.
//
// https://doc.babylonjs.com/api/classes/babylon.ifilerequest#oncompleteobservable
func (i *IFileRequest) OnCompleteObservable(onCompleteObservable *Observable) *IFileRequest {
	p := ba.ctx.Get("IFileRequest").New(onCompleteObservable.JSObject())
	return IFileRequestFromJSObject(p, ba.ctx)
}

// SetOnCompleteObservable sets the OnCompleteObservable property of class IFileRequest.
//
// https://doc.babylonjs.com/api/classes/babylon.ifilerequest#oncompleteobservable
func (i *IFileRequest) SetOnCompleteObservable(onCompleteObservable *Observable) *IFileRequest {
	p := ba.ctx.Get("IFileRequest").New(onCompleteObservable.JSObject())
	return IFileRequestFromJSObject(p, ba.ctx)
}

*/