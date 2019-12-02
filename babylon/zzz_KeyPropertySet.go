// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// KeyPropertySet represents a babylon.js KeyPropertySet.
// Class used to store key control properties
type KeyPropertySet struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (k *KeyPropertySet) JSObject() js.Value { return k.p }

// KeyPropertySet returns a KeyPropertySet JavaScript class.
func (ba *Babylon) KeyPropertySet() *KeyPropertySet {
	p := ba.ctx.Get("KeyPropertySet")
	return KeyPropertySetFromJSObject(p, ba.ctx)
}

// KeyPropertySetFromJSObject returns a wrapped KeyPropertySet JavaScript class.
func KeyPropertySetFromJSObject(p js.Value, ctx js.Value) *KeyPropertySet {
	return &KeyPropertySet{p: p, ctx: ctx}
}

// KeyPropertySetArrayToJSArray returns a JavaScript Array for the wrapped array.
func KeyPropertySetArrayToJSArray(array []*KeyPropertySet) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// Background returns the Background property of class KeyPropertySet.
//
// https://doc.babylonjs.com/api/classes/babylon.keypropertyset#background
func (k *KeyPropertySet) Background() string {
	retVal := k.p.Get("background")
	return retVal.String()
}

// SetBackground sets the Background property of class KeyPropertySet.
//
// https://doc.babylonjs.com/api/classes/babylon.keypropertyset#background
func (k *KeyPropertySet) SetBackground(background string) *KeyPropertySet {
	k.p.Set("background", background)
	return k
}

// Color returns the Color property of class KeyPropertySet.
//
// https://doc.babylonjs.com/api/classes/babylon.keypropertyset#color
func (k *KeyPropertySet) Color() string {
	retVal := k.p.Get("color")
	return retVal.String()
}

// SetColor sets the Color property of class KeyPropertySet.
//
// https://doc.babylonjs.com/api/classes/babylon.keypropertyset#color
func (k *KeyPropertySet) SetColor(color string) *KeyPropertySet {
	k.p.Set("color", color)
	return k
}

// Height returns the Height property of class KeyPropertySet.
//
// https://doc.babylonjs.com/api/classes/babylon.keypropertyset#height
func (k *KeyPropertySet) Height() string {
	retVal := k.p.Get("height")
	return retVal.String()
}

// SetHeight sets the Height property of class KeyPropertySet.
//
// https://doc.babylonjs.com/api/classes/babylon.keypropertyset#height
func (k *KeyPropertySet) SetHeight(height string) *KeyPropertySet {
	k.p.Set("height", height)
	return k
}

// PaddingBottom returns the PaddingBottom property of class KeyPropertySet.
//
// https://doc.babylonjs.com/api/classes/babylon.keypropertyset#paddingbottom
func (k *KeyPropertySet) PaddingBottom() string {
	retVal := k.p.Get("paddingBottom")
	return retVal.String()
}

// SetPaddingBottom sets the PaddingBottom property of class KeyPropertySet.
//
// https://doc.babylonjs.com/api/classes/babylon.keypropertyset#paddingbottom
func (k *KeyPropertySet) SetPaddingBottom(paddingBottom string) *KeyPropertySet {
	k.p.Set("paddingBottom", paddingBottom)
	return k
}

// PaddingLeft returns the PaddingLeft property of class KeyPropertySet.
//
// https://doc.babylonjs.com/api/classes/babylon.keypropertyset#paddingleft
func (k *KeyPropertySet) PaddingLeft() string {
	retVal := k.p.Get("paddingLeft")
	return retVal.String()
}

// SetPaddingLeft sets the PaddingLeft property of class KeyPropertySet.
//
// https://doc.babylonjs.com/api/classes/babylon.keypropertyset#paddingleft
func (k *KeyPropertySet) SetPaddingLeft(paddingLeft string) *KeyPropertySet {
	k.p.Set("paddingLeft", paddingLeft)
	return k
}

// PaddingRight returns the PaddingRight property of class KeyPropertySet.
//
// https://doc.babylonjs.com/api/classes/babylon.keypropertyset#paddingright
func (k *KeyPropertySet) PaddingRight() string {
	retVal := k.p.Get("paddingRight")
	return retVal.String()
}

// SetPaddingRight sets the PaddingRight property of class KeyPropertySet.
//
// https://doc.babylonjs.com/api/classes/babylon.keypropertyset#paddingright
func (k *KeyPropertySet) SetPaddingRight(paddingRight string) *KeyPropertySet {
	k.p.Set("paddingRight", paddingRight)
	return k
}

// PaddingTop returns the PaddingTop property of class KeyPropertySet.
//
// https://doc.babylonjs.com/api/classes/babylon.keypropertyset#paddingtop
func (k *KeyPropertySet) PaddingTop() string {
	retVal := k.p.Get("paddingTop")
	return retVal.String()
}

// SetPaddingTop sets the PaddingTop property of class KeyPropertySet.
//
// https://doc.babylonjs.com/api/classes/babylon.keypropertyset#paddingtop
func (k *KeyPropertySet) SetPaddingTop(paddingTop string) *KeyPropertySet {
	k.p.Set("paddingTop", paddingTop)
	return k
}

// Width returns the Width property of class KeyPropertySet.
//
// https://doc.babylonjs.com/api/classes/babylon.keypropertyset#width
func (k *KeyPropertySet) Width() string {
	retVal := k.p.Get("width")
	return retVal.String()
}

// SetWidth sets the Width property of class KeyPropertySet.
//
// https://doc.babylonjs.com/api/classes/babylon.keypropertyset#width
func (k *KeyPropertySet) SetWidth(width string) *KeyPropertySet {
	k.p.Set("width", width)
	return k
}
