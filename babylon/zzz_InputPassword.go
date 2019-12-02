// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// InputPassword represents a babylon.js InputPassword.
// Class used to create a password control
type InputPassword struct {
	*InputText
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (i *InputPassword) JSObject() js.Value { return i.p }

// InputPassword returns a InputPassword JavaScript class.
func (ba *Babylon) InputPassword() *InputPassword {
	p := ba.ctx.Get("InputPassword")
	return InputPasswordFromJSObject(p, ba.ctx)
}

// InputPasswordFromJSObject returns a wrapped InputPassword JavaScript class.
func InputPasswordFromJSObject(p js.Value, ctx js.Value) *InputPassword {
	return &InputPassword{InputText: InputTextFromJSObject(p, ctx), ctx: ctx}
}

// InputPasswordArrayToJSArray returns a JavaScript Array for the wrapped array.
func InputPasswordArrayToJSArray(array []*InputPassword) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewInputPasswordOpts contains optional parameters for NewInputPassword.
type NewInputPasswordOpts struct {
	Name *string
	Text *string
}

// NewInputPassword returns a new InputPassword object.
//
// https://doc.babylonjs.com/api/classes/babylon.inputpassword
func (ba *Babylon) NewInputPassword(opts *NewInputPasswordOpts) *InputPassword {
	if opts == nil {
		opts = &NewInputPasswordOpts{}
	}

	args := make([]interface{}, 0, 0+2)

	if opts.Name == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Name)
	}
	if opts.Text == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Text)
	}

	p := ba.ctx.Get("InputPassword").New(args...)
	return InputPasswordFromJSObject(p, ba.ctx)
}
