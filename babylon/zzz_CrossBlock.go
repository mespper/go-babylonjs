// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// CrossBlock represents a babylon.js CrossBlock.
// Block used to apply a cross product between 2 vectors
type CrossBlock struct {
	*NodeMaterialBlock
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (c *CrossBlock) JSObject() js.Value { return c.p }

// CrossBlock returns a CrossBlock JavaScript class.
func (ba *Babylon) CrossBlock() *CrossBlock {
	p := ba.ctx.Get("CrossBlock")
	return CrossBlockFromJSObject(p, ba.ctx)
}

// CrossBlockFromJSObject returns a wrapped CrossBlock JavaScript class.
func CrossBlockFromJSObject(p js.Value, ctx js.Value) *CrossBlock {
	return &CrossBlock{NodeMaterialBlock: NodeMaterialBlockFromJSObject(p, ctx), ctx: ctx}
}

// CrossBlockArrayToJSArray returns a JavaScript Array for the wrapped array.
func CrossBlockArrayToJSArray(array []*CrossBlock) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewCrossBlock returns a new CrossBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.crossblock
func (ba *Babylon) NewCrossBlock(name string) *CrossBlock {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	p := ba.ctx.Get("CrossBlock").New(args...)
	return CrossBlockFromJSObject(p, ba.ctx)
}

// GetClassName calls the GetClassName method on the CrossBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.crossblock#getclassname
func (c *CrossBlock) GetClassName() string {

	retVal := c.p.Call("getClassName")
	return retVal.String()
}

// Left returns the Left property of class CrossBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.crossblock#left
func (c *CrossBlock) Left() *NodeMaterialConnectionPoint {
	retVal := c.p.Get("left")
	return NodeMaterialConnectionPointFromJSObject(retVal, c.ctx)
}

// SetLeft sets the Left property of class CrossBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.crossblock#left
func (c *CrossBlock) SetLeft(left *NodeMaterialConnectionPoint) *CrossBlock {
	c.p.Set("left", left.JSObject())
	return c
}

// Output returns the Output property of class CrossBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.crossblock#output
func (c *CrossBlock) Output() *NodeMaterialConnectionPoint {
	retVal := c.p.Get("output")
	return NodeMaterialConnectionPointFromJSObject(retVal, c.ctx)
}

// SetOutput sets the Output property of class CrossBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.crossblock#output
func (c *CrossBlock) SetOutput(output *NodeMaterialConnectionPoint) *CrossBlock {
	c.p.Set("output", output.JSObject())
	return c
}

// Right returns the Right property of class CrossBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.crossblock#right
func (c *CrossBlock) Right() *NodeMaterialConnectionPoint {
	retVal := c.p.Get("right")
	return NodeMaterialConnectionPointFromJSObject(retVal, c.ctx)
}

// SetRight sets the Right property of class CrossBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.crossblock#right
func (c *CrossBlock) SetRight(right *NodeMaterialConnectionPoint) *CrossBlock {
	c.p.Set("right", right.JSObject())
	return c
}
