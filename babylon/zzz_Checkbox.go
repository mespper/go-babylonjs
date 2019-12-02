// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// Checkbox represents a babylon.js Checkbox.
// Class used to represent a 2D checkbox
type Checkbox struct {
	*Control
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (c *Checkbox) JSObject() js.Value { return c.p }

// Checkbox returns a Checkbox JavaScript class.
func (ba *Babylon) Checkbox() *Checkbox {
	p := ba.ctx.Get("Checkbox")
	return CheckboxFromJSObject(p, ba.ctx)
}

// CheckboxFromJSObject returns a wrapped Checkbox JavaScript class.
func CheckboxFromJSObject(p js.Value, ctx js.Value) *Checkbox {
	return &Checkbox{Control: ControlFromJSObject(p, ctx), ctx: ctx}
}

// CheckboxArrayToJSArray returns a JavaScript Array for the wrapped array.
func CheckboxArrayToJSArray(array []*Checkbox) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewCheckboxOpts contains optional parameters for NewCheckbox.
type NewCheckboxOpts struct {
	Name *string
}

// NewCheckbox returns a new Checkbox object.
//
// https://doc.babylonjs.com/api/classes/babylon.checkbox
func (ba *Babylon) NewCheckbox(opts *NewCheckboxOpts) *Checkbox {
	if opts == nil {
		opts = &NewCheckboxOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.Name == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Name)
	}

	p := ba.ctx.Get("Checkbox").New(args...)
	return CheckboxFromJSObject(p, ba.ctx)
}

// AddCheckBoxWithHeader calls the AddCheckBoxWithHeader method on the Checkbox object.
//
// https://doc.babylonjs.com/api/classes/babylon.checkbox#addcheckboxwithheader
func (c *Checkbox) AddCheckBoxWithHeader(title string, onValueChanged func()) *StackPanel {

	args := make([]interface{}, 0, 2+0)

	args = append(args, title)
	args = append(args, js.FuncOf(func(this js.Value, args []js.Value) interface{} { onValueChanged(); return nil }))

	retVal := c.p.Call("AddCheckBoxWithHeader", args...)
	return StackPanelFromJSObject(retVal, c.ctx)
}

// Background returns the Background property of class Checkbox.
//
// https://doc.babylonjs.com/api/classes/babylon.checkbox#background
func (c *Checkbox) Background() string {
	retVal := c.p.Get("background")
	return retVal.String()
}

// SetBackground sets the Background property of class Checkbox.
//
// https://doc.babylonjs.com/api/classes/babylon.checkbox#background
func (c *Checkbox) SetBackground(background string) *Checkbox {
	c.p.Set("background", background)
	return c
}

// CheckSizeRatio returns the CheckSizeRatio property of class Checkbox.
//
// https://doc.babylonjs.com/api/classes/babylon.checkbox#checksizeratio
func (c *Checkbox) CheckSizeRatio() float64 {
	retVal := c.p.Get("checkSizeRatio")
	return retVal.Float()
}

// SetCheckSizeRatio sets the CheckSizeRatio property of class Checkbox.
//
// https://doc.babylonjs.com/api/classes/babylon.checkbox#checksizeratio
func (c *Checkbox) SetCheckSizeRatio(checkSizeRatio float64) *Checkbox {
	c.p.Set("checkSizeRatio", checkSizeRatio)
	return c
}

// IsChecked returns the IsChecked property of class Checkbox.
//
// https://doc.babylonjs.com/api/classes/babylon.checkbox#ischecked
func (c *Checkbox) IsChecked() bool {
	retVal := c.p.Get("isChecked")
	return retVal.Bool()
}

// SetIsChecked sets the IsChecked property of class Checkbox.
//
// https://doc.babylonjs.com/api/classes/babylon.checkbox#ischecked
func (c *Checkbox) SetIsChecked(isChecked bool) *Checkbox {
	c.p.Set("isChecked", isChecked)
	return c
}

// Name returns the Name property of class Checkbox.
//
// https://doc.babylonjs.com/api/classes/babylon.checkbox#name
func (c *Checkbox) Name() string {
	retVal := c.p.Get("name")
	return retVal.String()
}

// SetName sets the Name property of class Checkbox.
//
// https://doc.babylonjs.com/api/classes/babylon.checkbox#name
func (c *Checkbox) SetName(name string) *Checkbox {
	c.p.Set("name", name)
	return c
}

// OnIsCheckedChangedObservable returns the OnIsCheckedChangedObservable property of class Checkbox.
//
// https://doc.babylonjs.com/api/classes/babylon.checkbox#onischeckedchangedobservable
func (c *Checkbox) OnIsCheckedChangedObservable() *Observable {
	retVal := c.p.Get("onIsCheckedChangedObservable")
	return ObservableFromJSObject(retVal, c.ctx)
}

// SetOnIsCheckedChangedObservable sets the OnIsCheckedChangedObservable property of class Checkbox.
//
// https://doc.babylonjs.com/api/classes/babylon.checkbox#onischeckedchangedobservable
func (c *Checkbox) SetOnIsCheckedChangedObservable(onIsCheckedChangedObservable *Observable) *Checkbox {
	c.p.Set("onIsCheckedChangedObservable", onIsCheckedChangedObservable.JSObject())
	return c
}

// Thickness returns the Thickness property of class Checkbox.
//
// https://doc.babylonjs.com/api/classes/babylon.checkbox#thickness
func (c *Checkbox) Thickness() float64 {
	retVal := c.p.Get("thickness")
	return retVal.Float()
}

// SetThickness sets the Thickness property of class Checkbox.
//
// https://doc.babylonjs.com/api/classes/babylon.checkbox#thickness
func (c *Checkbox) SetThickness(thickness float64) *Checkbox {
	c.p.Set("thickness", thickness)
	return c
}
