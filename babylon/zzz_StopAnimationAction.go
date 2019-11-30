// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// StopAnimationAction represents a babylon.js StopAnimationAction.
// This defines an action responsible to stop an animation once triggered.
//
// See: http://doc.babylonjs.com/how_to/how_to_use_actions
type StopAnimationAction struct{ *Action }

// JSObject returns the underlying js.Value.
func (s *StopAnimationAction) JSObject() js.Value { return s.p }

// StopAnimationAction returns a StopAnimationAction JavaScript class.
func (ba *Babylon) StopAnimationAction() *StopAnimationAction {
	p := ba.ctx.Get("StopAnimationAction")
	return StopAnimationActionFromJSObject(p)
}

// StopAnimationActionFromJSObject returns a wrapped StopAnimationAction JavaScript class.
func StopAnimationActionFromJSObject(p js.Value) *StopAnimationAction {
	return &StopAnimationAction{ActionFromJSObject(p)}
}

// NewStopAnimationActionOpts contains optional parameters for NewStopAnimationAction.
type NewStopAnimationActionOpts struct {
	Condition *Condition
}

// NewStopAnimationAction returns a new StopAnimationAction object.
//
// https://doc.babylonjs.com/api/classes/babylon.stopanimationaction
func (ba *Babylon) NewStopAnimationAction(triggerOptions interface{}, target interface{}, opts *NewStopAnimationActionOpts) *StopAnimationAction {
	if opts == nil {
		opts = &NewStopAnimationActionOpts{}
	}

	p := ba.ctx.Get("StopAnimationAction").New(triggerOptions, target, opts.Condition.JSObject())
	return StopAnimationActionFromJSObject(p)
}

// TODO: methods
