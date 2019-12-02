// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// SmoothStepBlock represents a babylon.js SmoothStepBlock.
// Block used to smooth step a value
type SmoothStepBlock struct {
	*NodeMaterialBlock
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (s *SmoothStepBlock) JSObject() js.Value { return s.p }

// SmoothStepBlock returns a SmoothStepBlock JavaScript class.
func (ba *Babylon) SmoothStepBlock() *SmoothStepBlock {
	p := ba.ctx.Get("SmoothStepBlock")
	return SmoothStepBlockFromJSObject(p, ba.ctx)
}

// SmoothStepBlockFromJSObject returns a wrapped SmoothStepBlock JavaScript class.
func SmoothStepBlockFromJSObject(p js.Value, ctx js.Value) *SmoothStepBlock {
	return &SmoothStepBlock{NodeMaterialBlock: NodeMaterialBlockFromJSObject(p, ctx), ctx: ctx}
}

// SmoothStepBlockArrayToJSArray returns a JavaScript Array for the wrapped array.
func SmoothStepBlockArrayToJSArray(array []*SmoothStepBlock) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewSmoothStepBlock returns a new SmoothStepBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.smoothstepblock
func (ba *Babylon) NewSmoothStepBlock(name string) *SmoothStepBlock {

	args := make([]interface{}, 0, 1+0)

	args = append(args, name)

	p := ba.ctx.Get("SmoothStepBlock").New(args...)
	return SmoothStepBlockFromJSObject(p, ba.ctx)
}

// GetClassName calls the GetClassName method on the SmoothStepBlock object.
//
// https://doc.babylonjs.com/api/classes/babylon.smoothstepblock#getclassname
func (s *SmoothStepBlock) GetClassName() string {

	retVal := s.p.Call("getClassName")
	return retVal.String()
}

// Edge0 returns the Edge0 property of class SmoothStepBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.smoothstepblock#edge0
func (s *SmoothStepBlock) Edge0() *NodeMaterialConnectionPoint {
	retVal := s.p.Get("edge0")
	return NodeMaterialConnectionPointFromJSObject(retVal, s.ctx)
}

// SetEdge0 sets the Edge0 property of class SmoothStepBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.smoothstepblock#edge0
func (s *SmoothStepBlock) SetEdge0(edge0 *NodeMaterialConnectionPoint) *SmoothStepBlock {
	s.p.Set("edge0", edge0.JSObject())
	return s
}

// Edge1 returns the Edge1 property of class SmoothStepBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.smoothstepblock#edge1
func (s *SmoothStepBlock) Edge1() *NodeMaterialConnectionPoint {
	retVal := s.p.Get("edge1")
	return NodeMaterialConnectionPointFromJSObject(retVal, s.ctx)
}

// SetEdge1 sets the Edge1 property of class SmoothStepBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.smoothstepblock#edge1
func (s *SmoothStepBlock) SetEdge1(edge1 *NodeMaterialConnectionPoint) *SmoothStepBlock {
	s.p.Set("edge1", edge1.JSObject())
	return s
}

// Output returns the Output property of class SmoothStepBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.smoothstepblock#output
func (s *SmoothStepBlock) Output() *NodeMaterialConnectionPoint {
	retVal := s.p.Get("output")
	return NodeMaterialConnectionPointFromJSObject(retVal, s.ctx)
}

// SetOutput sets the Output property of class SmoothStepBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.smoothstepblock#output
func (s *SmoothStepBlock) SetOutput(output *NodeMaterialConnectionPoint) *SmoothStepBlock {
	s.p.Set("output", output.JSObject())
	return s
}

// Value returns the Value property of class SmoothStepBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.smoothstepblock#value
func (s *SmoothStepBlock) Value() *NodeMaterialConnectionPoint {
	retVal := s.p.Get("value")
	return NodeMaterialConnectionPointFromJSObject(retVal, s.ctx)
}

// SetValue sets the Value property of class SmoothStepBlock.
//
// https://doc.babylonjs.com/api/classes/babylon.smoothstepblock#value
func (s *SmoothStepBlock) SetValue(value *NodeMaterialConnectionPoint) *SmoothStepBlock {
	s.p.Set("value", value.JSObject())
	return s
}
