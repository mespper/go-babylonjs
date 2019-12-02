// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// KeyboardInfo represents a babylon.js KeyboardInfo.
// This class is used to store keyboard related info for the onKeyboardObservable event.
type KeyboardInfo struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (k *KeyboardInfo) JSObject() js.Value { return k.p }

// KeyboardInfo returns a KeyboardInfo JavaScript class.
func (ba *Babylon) KeyboardInfo() *KeyboardInfo {
	p := ba.ctx.Get("KeyboardInfo")
	return KeyboardInfoFromJSObject(p, ba.ctx)
}

// KeyboardInfoFromJSObject returns a wrapped KeyboardInfo JavaScript class.
func KeyboardInfoFromJSObject(p js.Value, ctx js.Value) *KeyboardInfo {
	return &KeyboardInfo{p: p, ctx: ctx}
}

// KeyboardInfoArrayToJSArray returns a JavaScript Array for the wrapped array.
func KeyboardInfoArrayToJSArray(array []*KeyboardInfo) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewKeyboardInfo returns a new KeyboardInfo object.
//
// https://doc.babylonjs.com/api/classes/babylon.keyboardinfo
func (ba *Babylon) NewKeyboardInfo(jsType float64, event js.Value) *KeyboardInfo {

	args := make([]interface{}, 0, 2+0)

	args = append(args, jsType)
	args = append(args, event)

	p := ba.ctx.Get("KeyboardInfo").New(args...)
	return KeyboardInfoFromJSObject(p, ba.ctx)
}

// Event returns the Event property of class KeyboardInfo.
//
// https://doc.babylonjs.com/api/classes/babylon.keyboardinfo#event
func (k *KeyboardInfo) Event() js.Value {
	retVal := k.p.Get("event")
	return retVal
}

// SetEvent sets the Event property of class KeyboardInfo.
//
// https://doc.babylonjs.com/api/classes/babylon.keyboardinfo#event
func (k *KeyboardInfo) SetEvent(event js.Value) *KeyboardInfo {
	k.p.Set("event", event)
	return k
}

// Type returns the Type property of class KeyboardInfo.
//
// https://doc.babylonjs.com/api/classes/babylon.keyboardinfo#type
func (k *KeyboardInfo) Type() float64 {
	retVal := k.p.Get("type")
	return retVal.Float()
}

// SetType sets the Type property of class KeyboardInfo.
//
// https://doc.babylonjs.com/api/classes/babylon.keyboardinfo#type
func (k *KeyboardInfo) SetType(jsType float64) *KeyboardInfo {
	k.p.Set("type", jsType)
	return k
}
