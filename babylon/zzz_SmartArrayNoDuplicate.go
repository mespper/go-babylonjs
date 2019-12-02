// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// SmartArrayNoDuplicate represents a babylon.js SmartArrayNoDuplicate.
// Defines an GC Friendly array where the backfield array do not shrink to prevent over allocations.
// The data in this array can only be present once
type SmartArrayNoDuplicate struct {
	*SmartArray
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (s *SmartArrayNoDuplicate) JSObject() js.Value { return s.p }

// SmartArrayNoDuplicate returns a SmartArrayNoDuplicate JavaScript class.
func (ba *Babylon) SmartArrayNoDuplicate() *SmartArrayNoDuplicate {
	p := ba.ctx.Get("SmartArrayNoDuplicate")
	return SmartArrayNoDuplicateFromJSObject(p, ba.ctx)
}

// SmartArrayNoDuplicateFromJSObject returns a wrapped SmartArrayNoDuplicate JavaScript class.
func SmartArrayNoDuplicateFromJSObject(p js.Value, ctx js.Value) *SmartArrayNoDuplicate {
	return &SmartArrayNoDuplicate{SmartArray: SmartArrayFromJSObject(p, ctx), ctx: ctx}
}

// SmartArrayNoDuplicateArrayToJSArray returns a JavaScript Array for the wrapped array.
func SmartArrayNoDuplicateArrayToJSArray(array []*SmartArrayNoDuplicate) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewSmartArrayNoDuplicate returns a new SmartArrayNoDuplicate object.
//
// https://doc.babylonjs.com/api/classes/babylon.smartarraynoduplicate
func (ba *Babylon) NewSmartArrayNoDuplicate(capacity float64) *SmartArrayNoDuplicate {

	args := make([]interface{}, 0, 1+0)

	args = append(args, capacity)

	p := ba.ctx.Get("SmartArrayNoDuplicate").New(args...)
	return SmartArrayNoDuplicateFromJSObject(p, ba.ctx)
}

// ConcatWithNoDuplicate calls the ConcatWithNoDuplicate method on the SmartArrayNoDuplicate object.
//
// https://doc.babylonjs.com/api/classes/babylon.smartarraynoduplicate#concatwithnoduplicate
func (s *SmartArrayNoDuplicate) ConcatWithNoDuplicate(array interface{}) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, array)

	s.p.Call("concatWithNoDuplicate", args...)
}

// Push calls the Push method on the SmartArrayNoDuplicate object.
//
// https://doc.babylonjs.com/api/classes/babylon.smartarraynoduplicate#push
func (s *SmartArrayNoDuplicate) Push(value *T) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, value.JSObject())

	s.p.Call("push", args...)
}

// PushNoDuplicate calls the PushNoDuplicate method on the SmartArrayNoDuplicate object.
//
// https://doc.babylonjs.com/api/classes/babylon.smartarraynoduplicate#pushnoduplicate
func (s *SmartArrayNoDuplicate) PushNoDuplicate(value *T) bool {

	args := make([]interface{}, 0, 1+0)

	args = append(args, value.JSObject())

	retVal := s.p.Call("pushNoDuplicate", args...)
	return retVal.Bool()
}

// Reset calls the Reset method on the SmartArrayNoDuplicate object.
//
// https://doc.babylonjs.com/api/classes/babylon.smartarraynoduplicate#reset
func (s *SmartArrayNoDuplicate) Reset() {

	s.p.Call("reset")
}
