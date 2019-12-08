// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// AnimationEvent represents a babylon.js AnimationEvent.
// Composed of a frame, and an action function
type AnimationEvent struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (a *AnimationEvent) JSObject() js.Value { return a.p }

// AnimationEvent returns a AnimationEvent JavaScript class.
func (ba *Babylon) AnimationEvent() *AnimationEvent {
	p := ba.ctx.Get("AnimationEvent")
	return AnimationEventFromJSObject(p, ba.ctx)
}

// AnimationEventFromJSObject returns a wrapped AnimationEvent JavaScript class.
func AnimationEventFromJSObject(p js.Value, ctx js.Value) *AnimationEvent {
	return &AnimationEvent{p: p, ctx: ctx}
}

// AnimationEventArrayToJSArray returns a JavaScript Array for the wrapped array.
func AnimationEventArrayToJSArray(array []*AnimationEvent) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewAnimationEventOpts contains optional parameters for NewAnimationEvent.
type NewAnimationEventOpts struct {
	OnlyOnce *bool
}

// NewAnimationEvent returns a new AnimationEvent object.
//
// https://doc.babylonjs.com/api/classes/babylon.animationevent
func (ba *Babylon) NewAnimationEvent(frame float64, action JSFunc, opts *NewAnimationEventOpts) *AnimationEvent {
	if opts == nil {
		opts = &NewAnimationEventOpts{}
	}

	args := make([]interface{}, 0, 2+1)

	args = append(args, frame)
	args = append(args, js.FuncOf(action))

	if opts.OnlyOnce == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.OnlyOnce)
	}

	p := ba.ctx.Get("AnimationEvent").New(args...)
	return AnimationEventFromJSObject(p, ba.ctx)
}

// Action returns the Action property of class AnimationEvent.
//
// https://doc.babylonjs.com/api/classes/babylon.animationevent#action
func (a *AnimationEvent) Action() js.Value {
	retVal := a.p.Get("action")
	return retVal
}

// SetAction sets the Action property of class AnimationEvent.
//
// https://doc.babylonjs.com/api/classes/babylon.animationevent#action
func (a *AnimationEvent) SetAction(action JSFunc) *AnimationEvent {
	a.p.Set("action", js.FuncOf(action))
	return a
}

// Frame returns the Frame property of class AnimationEvent.
//
// https://doc.babylonjs.com/api/classes/babylon.animationevent#frame
func (a *AnimationEvent) Frame() float64 {
	retVal := a.p.Get("frame")
	return retVal.Float()
}

// SetFrame sets the Frame property of class AnimationEvent.
//
// https://doc.babylonjs.com/api/classes/babylon.animationevent#frame
func (a *AnimationEvent) SetFrame(frame float64) *AnimationEvent {
	a.p.Set("frame", frame)
	return a
}

// IsDone returns the IsDone property of class AnimationEvent.
//
// https://doc.babylonjs.com/api/classes/babylon.animationevent#isdone
func (a *AnimationEvent) IsDone() bool {
	retVal := a.p.Get("isDone")
	return retVal.Bool()
}

// SetIsDone sets the IsDone property of class AnimationEvent.
//
// https://doc.babylonjs.com/api/classes/babylon.animationevent#isdone
func (a *AnimationEvent) SetIsDone(isDone bool) *AnimationEvent {
	a.p.Set("isDone", isDone)
	return a
}

// OnlyOnce returns the OnlyOnce property of class AnimationEvent.
//
// https://doc.babylonjs.com/api/classes/babylon.animationevent#onlyonce
func (a *AnimationEvent) OnlyOnce() bool {
	retVal := a.p.Get("onlyOnce")
	return retVal.Bool()
}

// SetOnlyOnce sets the OnlyOnce property of class AnimationEvent.
//
// https://doc.babylonjs.com/api/classes/babylon.animationevent#onlyonce
func (a *AnimationEvent) SetOnlyOnce(onlyOnce bool) *AnimationEvent {
	a.p.Set("onlyOnce", onlyOnce)
	return a
}
