// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// IBufferView represents a babylon.js IBufferView.
// Loader interface with additional members.
type IBufferView struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (i *IBufferView) JSObject() js.Value { return i.p }

// IBufferView returns a IBufferView JavaScript class.
func (ba *Babylon) IBufferView() *IBufferView {
	p := ba.ctx.Get("IBufferView")
	return IBufferViewFromJSObject(p, ba.ctx)
}

// IBufferViewFromJSObject returns a wrapped IBufferView JavaScript class.
func IBufferViewFromJSObject(p js.Value, ctx js.Value) *IBufferView {
	return &IBufferView{p: p, ctx: ctx}
}

// IBufferViewArrayToJSArray returns a JavaScript Array for the wrapped array.
func IBufferViewArrayToJSArray(array []*IBufferView) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}
