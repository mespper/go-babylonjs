// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ReplaceColorBlock represents a babylon.js ReplaceColorBlock.
// Block used to replace a color by another one
type ReplaceColorBlock struct {
	*NodeMaterialBlock
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (r *ReplaceColorBlock) JSObject() js.Value { return r.p }

// ReplaceColorBlock returns a ReplaceColorBlock JavaScript class.
func (ba *Babylon) ReplaceColorBlock() *ReplaceColorBlock {
	p := ba.ctx.Get("ReplaceColorBlock")
	return ReplaceColorBlockFromJSObject(p, ba.ctx)
}

// ReplaceColorBlockFromJSObject returns a wrapped ReplaceColorBlock JavaScript class.
func ReplaceColorBlockFromJSObject(p js.Value, ctx js.Value) *ReplaceColorBlock {
	return &ReplaceColorBlock{NodeMaterialBlock: NodeMaterialBlockFromJSObject(p, ctx), ctx: ctx}
}

// ReplaceColorBlockArrayToJSArray returns a JavaScript Array for the wrapped array.
func ReplaceColorBlockArrayToJSArray(array []*ReplaceColorBlock) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewReplaceColorBlock returns a new ReplaceColorBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.replacecolorblock
func (ba *Babylon) NewReplaceColorBlock(name string) *ReplaceColorBlock {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	p := ba.ctx.Get("ReplaceColorBlock").New(args...)
	return ReplaceColorBlockFromJSObject(p, ba.ctx)
}

// GetClassName calls the GetClassName method on the ReplaceColorBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.replacecolorblock#getclassname
func (r *ReplaceColorBlock) GetClassName() string {

	retVal := r.p.Call("getClassName")
	return retVal.String()
}

// Distance returns the Distance property of class ReplaceColorBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.replacecolorblock#distance
func (r *ReplaceColorBlock) Distance() *NodeMaterialConnectionPoint {
	retVal := r.p.Get("distance")
	return NodeMaterialConnectionPointFromJSObject(retVal, r.ctx)
}

// SetDistance sets the Distance property of class ReplaceColorBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.replacecolorblock#distance
func (r *ReplaceColorBlock) SetDistance(distance *NodeMaterialConnectionPoint) *ReplaceColorBlock {
	r.p.Set("distance", distance.JSObject())
	return r
}

// Output returns the Output property of class ReplaceColorBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.replacecolorblock#output
func (r *ReplaceColorBlock) Output() *NodeMaterialConnectionPoint {
	retVal := r.p.Get("output")
	return NodeMaterialConnectionPointFromJSObject(retVal, r.ctx)
}

// SetOutput sets the Output property of class ReplaceColorBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.replacecolorblock#output
func (r *ReplaceColorBlock) SetOutput(output *NodeMaterialConnectionPoint) *ReplaceColorBlock {
	r.p.Set("output", output.JSObject())
	return r
}

// Reference returns the Reference property of class ReplaceColorBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.replacecolorblock#reference
func (r *ReplaceColorBlock) Reference() *NodeMaterialConnectionPoint {
	retVal := r.p.Get("reference")
	return NodeMaterialConnectionPointFromJSObject(retVal, r.ctx)
}

// SetReference sets the Reference property of class ReplaceColorBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.replacecolorblock#reference
func (r *ReplaceColorBlock) SetReference(reference *NodeMaterialConnectionPoint) *ReplaceColorBlock {
	r.p.Set("reference", reference.JSObject())
	return r
}

// Replacement returns the Replacement property of class ReplaceColorBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.replacecolorblock#replacement
func (r *ReplaceColorBlock) Replacement() *NodeMaterialConnectionPoint {
	retVal := r.p.Get("replacement")
	return NodeMaterialConnectionPointFromJSObject(retVal, r.ctx)
}

// SetReplacement sets the Replacement property of class ReplaceColorBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.replacecolorblock#replacement
func (r *ReplaceColorBlock) SetReplacement(replacement *NodeMaterialConnectionPoint) *ReplaceColorBlock {
	r.p.Set("replacement", replacement.JSObject())
	return r
}

// Value returns the Value property of class ReplaceColorBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.replacecolorblock#value
func (r *ReplaceColorBlock) Value() *NodeMaterialConnectionPoint {
	retVal := r.p.Get("value")
	return NodeMaterialConnectionPointFromJSObject(retVal, r.ctx)
}

// SetValue sets the Value property of class ReplaceColorBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.replacecolorblock#value
func (r *ReplaceColorBlock) SetValue(value *NodeMaterialConnectionPoint) *ReplaceColorBlock {
	r.p.Set("value", value.JSObject())
	return r
}
