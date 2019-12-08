// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ISimplifier represents a babylon.js ISimplifier.
// A simplifier interface for future simplification implementations
//
// See: http://doc.babylonjs.com/how_to/in-browser_mesh_simplification
type ISimplifier struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (i *ISimplifier) JSObject() js.Value { return i.p }

// ISimplifier returns a ISimplifier JavaScript class.
func (ba *Babylon) ISimplifier() *ISimplifier {
	p := ba.ctx.Get("ISimplifier")
	return ISimplifierFromJSObject(p, ba.ctx)
}

// ISimplifierFromJSObject returns a wrapped ISimplifier JavaScript class.
func ISimplifierFromJSObject(p js.Value, ctx js.Value) *ISimplifier {
	return &ISimplifier{p: p, ctx: ctx}
}

// ISimplifierArrayToJSArray returns a JavaScript Array for the wrapped array.
func ISimplifierArrayToJSArray(array []*ISimplifier) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// ISimplifierSimplifyOpts contains optional parameters for ISimplifier.Simplify.
type ISimplifierSimplifyOpts struct {
	ErrorCallback JSFunc
}

// Simplify calls the Simplify method on the ISimplifier object.
//
// https://doc.babylonjs.com/api/classes/babylon.isimplifier#simplify
func (i *ISimplifier) Simplify(settings *ISimplificationSettings, successCallback JSFunc, opts *ISimplifierSimplifyOpts) {
	if opts == nil {
		opts = &ISimplifierSimplifyOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	if settings == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, settings.JSObject())
	}

	args = append(args, js.FuncOf(successCallback))

	if opts.ErrorCallback == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, js.FuncOf(opts.ErrorCallback) /* never freed! */)
	}

	i.p.Call("simplify", args...)
}
