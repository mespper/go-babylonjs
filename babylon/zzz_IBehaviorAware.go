// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// IBehaviorAware represents a babylon.js IBehaviorAware.
// Interface implemented by classes supporting behaviors
type IBehaviorAware struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (i *IBehaviorAware) JSObject() js.Value { return i.p }

// IBehaviorAware returns a IBehaviorAware JavaScript class.
func (ba *Babylon) IBehaviorAware() *IBehaviorAware {
	p := ba.ctx.Get("IBehaviorAware")
	return IBehaviorAwareFromJSObject(p, ba.ctx)
}

// IBehaviorAwareFromJSObject returns a wrapped IBehaviorAware JavaScript class.
func IBehaviorAwareFromJSObject(p js.Value, ctx js.Value) *IBehaviorAware {
	return &IBehaviorAware{p: p, ctx: ctx}
}

// IBehaviorAwareArrayToJSArray returns a JavaScript Array for the wrapped array.
func IBehaviorAwareArrayToJSArray(array []*IBehaviorAware) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// AddBehavior calls the AddBehavior method on the IBehaviorAware object.
//
// https://doc.babylonjs.com/api/classes/babylon.ibehavioraware#addbehavior
func (i *IBehaviorAware) AddBehavior(behavior js.Value) *T {

	args := make([]interface{}, 0, 1+0)

	args = append(args, behavior)

	retVal := i.p.Call("addBehavior", args...)
	return TFromJSObject(retVal, i.ctx)
}

// GetBehaviorByName calls the GetBehaviorByName method on the IBehaviorAware object.
//
// https://doc.babylonjs.com/api/classes/babylon.ibehavioraware#getbehaviorbyname
func (i *IBehaviorAware) GetBehaviorByName(name string) js.Value {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	retVal := i.p.Call("getBehaviorByName", args...)
	return retVal
}

// RemoveBehavior calls the RemoveBehavior method on the IBehaviorAware object.
//
// https://doc.babylonjs.com/api/classes/babylon.ibehavioraware#removebehavior
func (i *IBehaviorAware) RemoveBehavior(behavior js.Value) *T {

	args := make([]interface{}, 0, 1+0)

	args = append(args, behavior)

	retVal := i.p.Call("removeBehavior", args...)
	return TFromJSObject(retVal, i.ctx)
}

/*

 */