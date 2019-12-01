// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// CameraInputsMap represents a babylon.js CameraInputsMap.
// Represents a map of input types to input instance or input index to input instance.
type CameraInputsMap struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (c *CameraInputsMap) JSObject() js.Value { return c.p }

// CameraInputsMap returns a CameraInputsMap JavaScript class.
func (ba *Babylon) CameraInputsMap() *CameraInputsMap {
	p := ba.ctx.Get("CameraInputsMap")
	return CameraInputsMapFromJSObject(p, ba.ctx)
}

// CameraInputsMapFromJSObject returns a wrapped CameraInputsMap JavaScript class.
func CameraInputsMapFromJSObject(p js.Value, ctx js.Value) *CameraInputsMap {
	return &CameraInputsMap{p: p, ctx: ctx}
}

// CameraInputsMapArrayToJSArray returns a JavaScript Array for the wrapped array.
func CameraInputsMapArrayToJSArray(array []*CameraInputsMap) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

/*

 */
