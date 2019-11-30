// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// AsyncLoop represents a babylon.js AsyncLoop.
// An implementation of a loop for asynchronous functions.
type AsyncLoop struct{ p js.Value }

// JSObject returns the underlying js.Value.
func (a *AsyncLoop) JSObject() js.Value { return a.p }

// AsyncLoop returns a AsyncLoop JavaScript class.
func (ba *Babylon) AsyncLoop() *AsyncLoop {
	p := ba.ctx.Get("AsyncLoop")
	return AsyncLoopFromJSObject(p)
}

// AsyncLoopFromJSObject returns a wrapped AsyncLoop JavaScript class.
func AsyncLoopFromJSObject(p js.Value) *AsyncLoop {
	return &AsyncLoop{p: p}
}

// NewAsyncLoopOpts contains optional parameters for NewAsyncLoop.
type NewAsyncLoopOpts struct {
	Offset *JSFloat64
}

// NewAsyncLoop returns a new AsyncLoop object.
//
// https://doc.babylonjs.com/api/classes/babylon.asyncloop
func (ba *Babylon) NewAsyncLoop(iterations float64, jsFunc func(), successCallback func(), opts *NewAsyncLoopOpts) *AsyncLoop {
	if opts == nil {
		opts = &NewAsyncLoopOpts{}
	}

	p := ba.ctx.Get("AsyncLoop").New(iterations, jsFunc, successCallback, opts.Offset)
	return AsyncLoopFromJSObject(p)
}

// TODO: methods
