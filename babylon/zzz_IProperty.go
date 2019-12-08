// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// IProperty represents a babylon.js IProperty.
// glTF Property
type IProperty struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (i *IProperty) JSObject() js.Value { return i.p }

// IProperty returns a IProperty JavaScript class.
func (ba *Babylon) IProperty() *IProperty {
	p := ba.ctx.Get("IProperty")
	return IPropertyFromJSObject(p, ba.ctx)
}

// IPropertyFromJSObject returns a wrapped IProperty JavaScript class.
func IPropertyFromJSObject(p js.Value, ctx js.Value) *IProperty {
	return &IProperty{p: p, ctx: ctx}
}

// IPropertyArrayToJSArray returns a JavaScript Array for the wrapped array.
func IPropertyArrayToJSArray(array []*IProperty) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// Extensions returns the Extensions property of class IProperty.
//
// https://doc.babylonjs.com/api/classes/babylon.iproperty#extensions
func (i *IProperty) Extensions() js.Value {
	retVal := i.p.Get("extensions")
	return retVal
}

// SetExtensions sets the Extensions property of class IProperty.
//
// https://doc.babylonjs.com/api/classes/babylon.iproperty#extensions
func (i *IProperty) SetExtensions(extensions js.Value) *IProperty {
	i.p.Set("extensions", extensions)
	return i
}

// Extras returns the Extras property of class IProperty.
//
// https://doc.babylonjs.com/api/classes/babylon.iproperty#extras
func (i *IProperty) Extras() js.Value {
	retVal := i.p.Get("extras")
	return retVal
}

// SetExtras sets the Extras property of class IProperty.
//
// https://doc.babylonjs.com/api/classes/babylon.iproperty#extras
func (i *IProperty) SetExtras(extras JSObject) *IProperty {
	i.p.Set("extras", extras.JSObject())
	return i
}
