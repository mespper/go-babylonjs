// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// IArrayItem represents a babylon.js IArrayItem.
// Loader interface with an index field.
type IArrayItem struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (i *IArrayItem) JSObject() js.Value { return i.p }

// IArrayItem returns a IArrayItem JavaScript class.
func (ba *Babylon) IArrayItem() *IArrayItem {
	p := ba.ctx.Get("IArrayItem")
	return IArrayItemFromJSObject(p, ba.ctx)
}

// IArrayItemFromJSObject returns a wrapped IArrayItem JavaScript class.
func IArrayItemFromJSObject(p js.Value, ctx js.Value) *IArrayItem {
	return &IArrayItem{p: p, ctx: ctx}
}

// IArrayItemArrayToJSArray returns a JavaScript Array for the wrapped array.
func IArrayItemArrayToJSArray(array []*IArrayItem) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

/*

// Index returns the Index property of class IArrayItem.
//
// https://doc.babylonjs.com/api/classes/babylon.iarrayitem#index
func (i *IArrayItem) Index(index float64) *IArrayItem {
	p := ba.ctx.Get("IArrayItem").New(index)
	return IArrayItemFromJSObject(p, ba.ctx)
}

// SetIndex sets the Index property of class IArrayItem.
//
// https://doc.babylonjs.com/api/classes/babylon.iarrayitem#index
func (i *IArrayItem) SetIndex(index float64) *IArrayItem {
	p := ba.ctx.Get("IArrayItem").New(index)
	return IArrayItemFromJSObject(p, ba.ctx)
}

*/
