// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// InterpolateValueAction represents a babylon.js InterpolateValueAction.
// This defines an action responsible to change the value of a property
// by interpolating between its current value and the newly set one once triggered.
//
// See: http://doc.babylonjs.com/how_to/how_to_use_actions
type InterpolateValueAction struct{ *Action }

// JSObject returns the underlying js.Value.
func (i *InterpolateValueAction) JSObject() js.Value { return i.p }

// InterpolateValueAction returns a InterpolateValueAction JavaScript class.
func (ba *Babylon) InterpolateValueAction() *InterpolateValueAction {
	p := ba.ctx.Get("InterpolateValueAction")
	return InterpolateValueActionFromJSObject(p)
}

// InterpolateValueActionFromJSObject returns a wrapped InterpolateValueAction JavaScript class.
func InterpolateValueActionFromJSObject(p js.Value) *InterpolateValueAction {
	return &InterpolateValueAction{ActionFromJSObject(p)}
}

// NewInterpolateValueActionOpts contains optional parameters for NewInterpolateValueAction.
type NewInterpolateValueActionOpts struct {
	Duration *JSFloat64

	Condition *Condition

	StopOtherAnimations *JSBool

	OnInterpolationDone *func()
}

// NewInterpolateValueAction returns a new InterpolateValueAction object.
//
// https://doc.babylonjs.com/api/classes/babylon.interpolatevalueaction
func (ba *Babylon) NewInterpolateValueAction(triggerOptions interface{}, target interface{}, propertyPath string, value interface{}, opts *NewInterpolateValueActionOpts) *InterpolateValueAction {
	if opts == nil {
		opts = &NewInterpolateValueActionOpts{}
	}

	p := ba.ctx.Get("InterpolateValueAction").New(triggerOptions, target, propertyPath, value, opts.Duration, opts.Condition.JSObject(), opts.StopOtherAnimations, opts.OnInterpolationDone)
	return InterpolateValueActionFromJSObject(p)
}

// TODO: methods
