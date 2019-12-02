// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// IExplorerExtensibilityOption represents a babylon.js IExplorerExtensibilityOption.
// Interface used to define scene explorer extensibility option
type IExplorerExtensibilityOption struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (i *IExplorerExtensibilityOption) JSObject() js.Value { return i.p }

// IExplorerExtensibilityOption returns a IExplorerExtensibilityOption JavaScript class.
func (ba *Babylon) IExplorerExtensibilityOption() *IExplorerExtensibilityOption {
	p := ba.ctx.Get("IExplorerExtensibilityOption")
	return IExplorerExtensibilityOptionFromJSObject(p, ba.ctx)
}

// IExplorerExtensibilityOptionFromJSObject returns a wrapped IExplorerExtensibilityOption JavaScript class.
func IExplorerExtensibilityOptionFromJSObject(p js.Value, ctx js.Value) *IExplorerExtensibilityOption {
	return &IExplorerExtensibilityOption{p: p, ctx: ctx}
}

// IExplorerExtensibilityOptionArrayToJSArray returns a JavaScript Array for the wrapped array.
func IExplorerExtensibilityOptionArrayToJSArray(array []*IExplorerExtensibilityOption) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// Action returns the Action property of class IExplorerExtensibilityOption.
//
// https://doc.babylonjs.com/api/classes/babylon.iexplorerextensibilityoption#action
func (i *IExplorerExtensibilityOption) Action() js.Value {
	retVal := i.p.Get("action")
	return retVal
}

// SetAction sets the Action property of class IExplorerExtensibilityOption.
//
// https://doc.babylonjs.com/api/classes/babylon.iexplorerextensibilityoption#action
func (i *IExplorerExtensibilityOption) SetAction(action func()) *IExplorerExtensibilityOption {
	i.p.Set("action", js.FuncOf(func(this js.Value, args []js.Value) interface{} { action(); return nil }))
	return i
}

// Label returns the Label property of class IExplorerExtensibilityOption.
//
// https://doc.babylonjs.com/api/classes/babylon.iexplorerextensibilityoption#label
func (i *IExplorerExtensibilityOption) Label() string {
	retVal := i.p.Get("label")
	return retVal.String()
}

// SetLabel sets the Label property of class IExplorerExtensibilityOption.
//
// https://doc.babylonjs.com/api/classes/babylon.iexplorerextensibilityoption#label
func (i *IExplorerExtensibilityOption) SetLabel(label string) *IExplorerExtensibilityOption {
	i.p.Set("label", label)
	return i
}
