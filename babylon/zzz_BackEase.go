// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// BackEase represents a babylon.js BackEase.
// Easing function with a ease back shape (see link below).
//
// See: https://easings.net/#easeInBack
// See: http://doc.babylonjs.com/babylon101/animations#easing-functions
type BackEase struct {
	*EasingFunction
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (b *BackEase) JSObject() js.Value { return b.p }

// BackEase returns a BackEase JavaScript class.
func (ba *Babylon) BackEase() *BackEase {
	p := ba.ctx.Get("BackEase")
	return BackEaseFromJSObject(p, ba.ctx)
}

// BackEaseFromJSObject returns a wrapped BackEase JavaScript class.
func BackEaseFromJSObject(p js.Value, ctx js.Value) *BackEase {
	return &BackEase{EasingFunction: EasingFunctionFromJSObject(p, ctx), ctx: ctx}
}

// BackEaseArrayToJSArray returns a JavaScript Array for the wrapped array.
func BackEaseArrayToJSArray(array []*BackEase) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewBackEaseOpts contains optional parameters for NewBackEase.
type NewBackEaseOpts struct {
	Amplitude *float64
}

// NewBackEase returns a new BackEase object.
//
// https://doc.babylonjs.com/api/classes/babylon.backease
func (ba *Babylon) NewBackEase(opts *NewBackEaseOpts) *BackEase {
	if opts == nil {
		opts = &NewBackEaseOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.Amplitude == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Amplitude)
	}

	p := ba.ctx.Get("BackEase").New(args...)
	return BackEaseFromJSObject(p, ba.ctx)
}

// Amplitude returns the Amplitude property of class BackEase.
//
// https://doc.babylonjs.com/api/classes/babylon.backease#amplitude
func (b *BackEase) Amplitude() float64 {
	retVal := b.p.Get("amplitude")
	return retVal.Float()
}

// SetAmplitude sets the Amplitude property of class BackEase.
//
// https://doc.babylonjs.com/api/classes/babylon.backease#amplitude
func (b *BackEase) SetAmplitude(amplitude float64) *BackEase {
	b.p.Set("amplitude", amplitude)
	return b
}
