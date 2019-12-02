// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ReciprocalBlock represents a babylon.js ReciprocalBlock.
// Block used to get the reciprocal (1 / x) of a value
type ReciprocalBlock struct {
	*NodeMaterialBlock
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (r *ReciprocalBlock) JSObject() js.Value { return r.p }

// ReciprocalBlock returns a ReciprocalBlock JavaScript class.
func (ba *Babylon) ReciprocalBlock() *ReciprocalBlock {
	p := ba.ctx.Get("ReciprocalBlock")
	return ReciprocalBlockFromJSObject(p, ba.ctx)
}

// ReciprocalBlockFromJSObject returns a wrapped ReciprocalBlock JavaScript class.
func ReciprocalBlockFromJSObject(p js.Value, ctx js.Value) *ReciprocalBlock {
	return &ReciprocalBlock{NodeMaterialBlock: NodeMaterialBlockFromJSObject(p, ctx), ctx: ctx}
}

// ReciprocalBlockArrayToJSArray returns a JavaScript Array for the wrapped array.
func ReciprocalBlockArrayToJSArray(array []*ReciprocalBlock) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewReciprocalBlock returns a new ReciprocalBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.reciprocalblock
func (ba *Babylon) NewReciprocalBlock(name string) *ReciprocalBlock {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	p := ba.ctx.Get("ReciprocalBlock").New(args...)
	return ReciprocalBlockFromJSObject(p, ba.ctx)
}

// GetClassName calls the GetClassName method on the ReciprocalBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.reciprocalblock#getclassname
func (r *ReciprocalBlock) GetClassName() string {

	retVal := r.p.Call("getClassName")
	return retVal.String()
}

// Input returns the Input property of class ReciprocalBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.reciprocalblock#input
func (r *ReciprocalBlock) Input() *NodeMaterialConnectionPoint {
	retVal := r.p.Get("input")
	return NodeMaterialConnectionPointFromJSObject(retVal, r.ctx)
}

// SetInput sets the Input property of class ReciprocalBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.reciprocalblock#input
func (r *ReciprocalBlock) SetInput(input *NodeMaterialConnectionPoint) *ReciprocalBlock {
	r.p.Set("input", input.JSObject())
	return r
}

// Output returns the Output property of class ReciprocalBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.reciprocalblock#output
func (r *ReciprocalBlock) Output() *NodeMaterialConnectionPoint {
	retVal := r.p.Get("output")
	return NodeMaterialConnectionPointFromJSObject(retVal, r.ctx)
}

// SetOutput sets the Output property of class ReciprocalBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.reciprocalblock#output
func (r *ReciprocalBlock) SetOutput(output *NodeMaterialConnectionPoint) *ReciprocalBlock {
	r.p.Set("output", output.JSObject())
	return r
}
