// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// SmartArray represents a babylon.js SmartArray.
// Defines an GC Friendly array where the backfield array do not shrink to prevent over allocations.
type SmartArray struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (s *SmartArray) JSObject() js.Value { return s.p }

// SmartArray returns a SmartArray JavaScript class.
func (ba *Babylon) SmartArray() *SmartArray {
	p := ba.ctx.Get("SmartArray")
	return SmartArrayFromJSObject(p, ba.ctx)
}

// SmartArrayFromJSObject returns a wrapped SmartArray JavaScript class.
func SmartArrayFromJSObject(p js.Value, ctx js.Value) *SmartArray {
	return &SmartArray{p: p, ctx: ctx}
}

// SmartArrayArrayToJSArray returns a JavaScript Array for the wrapped array.
func SmartArrayArrayToJSArray(array []*SmartArray) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewSmartArray returns a new SmartArray object.
//
// https://doc.babylonjs.com/api/classes/babylon.smartarray
func (ba *Babylon) NewSmartArray(capacity float64) *SmartArray {

	args := make([]interface{}, 0, 1+0)

	args = append(args, capacity)

	p := ba.ctx.Get("SmartArray").New(args...)
	return SmartArrayFromJSObject(p, ba.ctx)
}

// Concat calls the Concat method on the SmartArray object.
//
// https://doc.babylonjs.com/api/classes/babylon.smartarray#concat
func (s *SmartArray) Concat(array JSObject) {

	args := make([]interface{}, 0, 1+0)

	if array == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, array.JSObject())
	}

	s.p.Call("concat", args...)
}

// Contains calls the Contains method on the SmartArray object.
//
// https://doc.babylonjs.com/api/classes/babylon.smartarray#contains
func (s *SmartArray) Contains(value *T) bool {

	args := make([]interface{}, 0, 1+0)

	if value == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, value.JSObject())
	}

	retVal := s.p.Call("contains", args...)
	return retVal.Bool()
}

// Dispose calls the Dispose method on the SmartArray object.
//
// https://doc.babylonjs.com/api/classes/babylon.smartarray#dispose
func (s *SmartArray) Dispose() {

	s.p.Call("dispose")
}

// ForEach calls the ForEach method on the SmartArray object.
//
// https://doc.babylonjs.com/api/classes/babylon.smartarray#foreach
func (s *SmartArray) ForEach(jsFunc JSFunc) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, js.FuncOf(jsFunc))

	s.p.Call("forEach", args...)
}

// IndexOf calls the IndexOf method on the SmartArray object.
//
// https://doc.babylonjs.com/api/classes/babylon.smartarray#indexof
func (s *SmartArray) IndexOf(value *T) float64 {

	args := make([]interface{}, 0, 1+0)

	if value == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, value.JSObject())
	}

	retVal := s.p.Call("indexOf", args...)
	return retVal.Float()
}

// Push calls the Push method on the SmartArray object.
//
// https://doc.babylonjs.com/api/classes/babylon.smartarray#push
func (s *SmartArray) Push(value *T) {

	args := make([]interface{}, 0, 1+0)

	if value == nil {
		args = append(args, js.Null())
	} else {
		args = append(args, value.JSObject())
	}

	s.p.Call("push", args...)
}

// Reset calls the Reset method on the SmartArray object.
//
// https://doc.babylonjs.com/api/classes/babylon.smartarray#reset
func (s *SmartArray) Reset() {

	s.p.Call("reset")
}

// Sort calls the Sort method on the SmartArray object.
//
// https://doc.babylonjs.com/api/classes/babylon.smartarray#sort
func (s *SmartArray) Sort(compareFn JSFunc) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, js.FuncOf(compareFn))

	s.p.Call("sort", args...)
}

// Data returns the Data property of class SmartArray.
//
// https://doc.babylonjs.com/api/classes/babylon.smartarray#data
func (s *SmartArray) Data() []*T {
	retVal := s.p.Get("data")
	result := []*T{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, TFromJSObject(retVal.Index(ri), s.ctx))
	}
	return result
}

// SetData sets the Data property of class SmartArray.
//
// https://doc.babylonjs.com/api/classes/babylon.smartarray#data
func (s *SmartArray) SetData(data []*T) *SmartArray {
	s.p.Set("data", data)
	return s
}

// Length returns the Length property of class SmartArray.
//
// https://doc.babylonjs.com/api/classes/babylon.smartarray#length
func (s *SmartArray) Length() float64 {
	retVal := s.p.Get("length")
	return retVal.Float()
}

// SetLength sets the Length property of class SmartArray.
//
// https://doc.babylonjs.com/api/classes/babylon.smartarray#length
func (s *SmartArray) SetLength(length float64) *SmartArray {
	s.p.Set("length", length)
	return s
}
