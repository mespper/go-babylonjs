// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// AndOrNotEvaluator represents a babylon.js AndOrNotEvaluator.
// Class used to evalaute queries containing <code>and</code> and <code>or</code> operators
type AndOrNotEvaluator struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (a *AndOrNotEvaluator) JSObject() js.Value { return a.p }

// AndOrNotEvaluator returns a AndOrNotEvaluator JavaScript class.
func (ba *Babylon) AndOrNotEvaluator() *AndOrNotEvaluator {
	p := ba.ctx.Get("AndOrNotEvaluator")
	return AndOrNotEvaluatorFromJSObject(p, ba.ctx)
}

// AndOrNotEvaluatorFromJSObject returns a wrapped AndOrNotEvaluator JavaScript class.
func AndOrNotEvaluatorFromJSObject(p js.Value, ctx js.Value) *AndOrNotEvaluator {
	return &AndOrNotEvaluator{p: p, ctx: ctx}
}

// AndOrNotEvaluatorArrayToJSArray returns a JavaScript Array for the wrapped array.
func AndOrNotEvaluatorArrayToJSArray(array []*AndOrNotEvaluator) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// Eval calls the Eval method on the AndOrNotEvaluator object.
//
// https://doc.babylonjs.com/api/classes/babylon.andornotevaluator#eval
func (a *AndOrNotEvaluator) Eval(query string, evaluateCallback func()) bool {

	args := make([]interface{}, 0, 2+0)

	args = append(args, query)
	args = append(args, js.FuncOf(func(this js.Value, args []js.Value) interface{} { evaluateCallback(); return nil }))

	retVal := a.p.Call("Eval", args...)
	return retVal.Bool()
}

/*

 */
