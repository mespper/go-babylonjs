// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// RadioGroup represents a babylon.js RadioGroup.
// Class used to create a RadioGroup
// which contains groups of radio buttons
type RadioGroup struct {
	*SelectorGroup
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (r *RadioGroup) JSObject() js.Value { return r.p }

// RadioGroup returns a RadioGroup JavaScript class.
func (ba *Babylon) RadioGroup() *RadioGroup {
	p := ba.ctx.Get("RadioGroup")
	return RadioGroupFromJSObject(p, ba.ctx)
}

// RadioGroupFromJSObject returns a wrapped RadioGroup JavaScript class.
func RadioGroupFromJSObject(p js.Value, ctx js.Value) *RadioGroup {
	return &RadioGroup{SelectorGroup: SelectorGroupFromJSObject(p, ctx), ctx: ctx}
}

// RadioGroupArrayToJSArray returns a JavaScript Array for the wrapped array.
func RadioGroupArrayToJSArray(array []*RadioGroup) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewRadioGroup returns a new RadioGroup object.
//
// https://doc.babylonjs.com/api/classes/babylon.radiogroup
func (ba *Babylon) NewRadioGroup(name string) *RadioGroup {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	p := ba.ctx.Get("RadioGroup").New(args...)
	return RadioGroupFromJSObject(p, ba.ctx)
}

// RadioGroupAddRadioOpts contains optional parameters for RadioGroup.AddRadio.
type RadioGroupAddRadioOpts struct {
	Func    func()
	Checked *bool
}

// AddRadio calls the AddRadio method on the RadioGroup object.
//
// https://doc.babylonjs.com/api/classes/babylon.radiogroup#addradio
func (r *RadioGroup) AddRadio(label string, opts *RadioGroupAddRadioOpts) {
	if opts == nil {
		opts = &RadioGroupAddRadioOpts{}
	}

	args := make([]interface{}, 0, 1+2)

	args = append(args, label)

	if opts.Func == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Func)
	}
	if opts.Checked == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Checked)
	}

	r.p.Call("addRadio", args...)
}
