// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ColorPicker represents a babylon.js ColorPicker.
// Class used to create color pickers
type ColorPicker struct {
	*Control
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (c *ColorPicker) JSObject() js.Value { return c.p }

// ColorPicker returns a ColorPicker JavaScript class.
func (ba *Babylon) ColorPicker() *ColorPicker {
	p := ba.ctx.Get("ColorPicker")
	return ColorPickerFromJSObject(p, ba.ctx)
}

// ColorPickerFromJSObject returns a wrapped ColorPicker JavaScript class.
func ColorPickerFromJSObject(p js.Value, ctx js.Value) *ColorPicker {
	return &ColorPicker{Control: ControlFromJSObject(p, ctx), ctx: ctx}
}

// ColorPickerArrayToJSArray returns a JavaScript Array for the wrapped array.
func ColorPickerArrayToJSArray(array []*ColorPicker) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewColorPickerOpts contains optional parameters for NewColorPicker.
type NewColorPickerOpts struct {
	Name *string
}

// NewColorPicker returns a new ColorPicker object.
//
// https://doc.babylonjs.com/api/classes/babylon.colorpicker
func (ba *Babylon) NewColorPicker(opts *NewColorPickerOpts) *ColorPicker {
	if opts == nil {
		opts = &NewColorPickerOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.Name == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Name)
	}

	p := ba.ctx.Get("ColorPicker").New(args...)
	return ColorPickerFromJSObject(p, ba.ctx)
}

// ShowPickerDialogAsync calls the ShowPickerDialogAsync method on the ColorPicker object.
//
// https://doc.babylonjs.com/api/classes/babylon.colorpicker#showpickerdialogasync
func (c *ColorPicker) ShowPickerDialogAsync(advancedTexture *AdvancedDynamicTexture, options js.Value) *Promise {

	args := make([]interface{}, 0, 2+0)

	args = append(args, advancedTexture.JSObject())
	args = append(args, options)

	retVal := c.p.Call("ShowPickerDialogAsync", args...)
	return PromiseFromJSObject(retVal, c.ctx)
}

// _onPointerDown calls the _onPointerDown method on the ColorPicker object.
//
// https://doc.babylonjs.com/api/classes/babylon.colorpicker#_onpointerdown
func (c *ColorPicker) _onPointerDown(target *Control, coordinates *Vector2, pointerId float64, buttonIndex float64) bool {

	args := make([]interface{}, 0, 4+0)

	args = append(args, target.JSObject())
	args = append(args, coordinates.JSObject())
	args = append(args, pointerId)
	args = append(args, buttonIndex)

	retVal := c.p.Call("_onPointerDown", args...)
	return retVal.Bool()
}

// _onPointerMove calls the _onPointerMove method on the ColorPicker object.
//
// https://doc.babylonjs.com/api/classes/babylon.colorpicker#_onpointermove
func (c *ColorPicker) _onPointerMove(target *Control, coordinates *Vector2, pointerId float64) {

	args := make([]interface{}, 0, 3+0)

	args = append(args, target.JSObject())
	args = append(args, coordinates.JSObject())
	args = append(args, pointerId)

	c.p.Call("_onPointerMove", args...)
}

// _onPointerUp calls the _onPointerUp method on the ColorPicker object.
//
// https://doc.babylonjs.com/api/classes/babylon.colorpicker#_onpointerup
func (c *ColorPicker) _onPointerUp(target *Control, coordinates *Vector2, pointerId float64, buttonIndex float64, notifyClick bool) {

	args := make([]interface{}, 0, 5+0)

	args = append(args, target.JSObject())
	args = append(args, coordinates.JSObject())
	args = append(args, pointerId)
	args = append(args, buttonIndex)
	args = append(args, notifyClick)

	c.p.Call("_onPointerUp", args...)
}

// Height returns the Height property of class ColorPicker.
//
// https://doc.babylonjs.com/api/classes/babylon.colorpicker#height
func (c *ColorPicker) Height() string {
	retVal := c.p.Get("height")
	return retVal.String()
}

// SetHeight sets the Height property of class ColorPicker.
//
// https://doc.babylonjs.com/api/classes/babylon.colorpicker#height
func (c *ColorPicker) SetHeight(height string) *ColorPicker {
	c.p.Set("height", height)
	return c
}

// Name returns the Name property of class ColorPicker.
//
// https://doc.babylonjs.com/api/classes/babylon.colorpicker#name
func (c *ColorPicker) Name() string {
	retVal := c.p.Get("name")
	return retVal.String()
}

// SetName sets the Name property of class ColorPicker.
//
// https://doc.babylonjs.com/api/classes/babylon.colorpicker#name
func (c *ColorPicker) SetName(name string) *ColorPicker {
	c.p.Set("name", name)
	return c
}

// OnValueChangedObservable returns the OnValueChangedObservable property of class ColorPicker.
//
// https://doc.babylonjs.com/api/classes/babylon.colorpicker#onvaluechangedobservable
func (c *ColorPicker) OnValueChangedObservable() *Observable {
	retVal := c.p.Get("onValueChangedObservable")
	return ObservableFromJSObject(retVal, c.ctx)
}

// SetOnValueChangedObservable sets the OnValueChangedObservable property of class ColorPicker.
//
// https://doc.babylonjs.com/api/classes/babylon.colorpicker#onvaluechangedobservable
func (c *ColorPicker) SetOnValueChangedObservable(onValueChangedObservable *Observable) *ColorPicker {
	c.p.Set("onValueChangedObservable", onValueChangedObservable.JSObject())
	return c
}

// Size returns the Size property of class ColorPicker.
//
// https://doc.babylonjs.com/api/classes/babylon.colorpicker#size
func (c *ColorPicker) Size() string {
	retVal := c.p.Get("size")
	return retVal.String()
}

// SetSize sets the Size property of class ColorPicker.
//
// https://doc.babylonjs.com/api/classes/babylon.colorpicker#size
func (c *ColorPicker) SetSize(size string) *ColorPicker {
	c.p.Set("size", size)
	return c
}

// Value returns the Value property of class ColorPicker.
//
// https://doc.babylonjs.com/api/classes/babylon.colorpicker#value
func (c *ColorPicker) Value() *Color3 {
	retVal := c.p.Get("value")
	return Color3FromJSObject(retVal, c.ctx)
}

// SetValue sets the Value property of class ColorPicker.
//
// https://doc.babylonjs.com/api/classes/babylon.colorpicker#value
func (c *ColorPicker) SetValue(value *Color3) *ColorPicker {
	c.p.Set("value", value.JSObject())
	return c
}

// Width returns the Width property of class ColorPicker.
//
// https://doc.babylonjs.com/api/classes/babylon.colorpicker#width
func (c *ColorPicker) Width() string {
	retVal := c.p.Get("width")
	return retVal.String()
}

// SetWidth sets the Width property of class ColorPicker.
//
// https://doc.babylonjs.com/api/classes/babylon.colorpicker#width
func (c *ColorPicker) SetWidth(width string) *ColorPicker {
	c.p.Set("width", width)
	return c
}
