// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// Measure represents a babylon.js Measure.
// Class used to store 2D control sizes
type Measure struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (m *Measure) JSObject() js.Value { return m.p }

// Measure returns a Measure JavaScript class.
func (ba *Babylon) Measure() *Measure {
	p := ba.ctx.Get("Measure")
	return MeasureFromJSObject(p, ba.ctx)
}

// MeasureFromJSObject returns a wrapped Measure JavaScript class.
func MeasureFromJSObject(p js.Value, ctx js.Value) *Measure {
	return &Measure{p: p, ctx: ctx}
}

// MeasureArrayToJSArray returns a JavaScript Array for the wrapped array.
func MeasureArrayToJSArray(array []*Measure) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewMeasure returns a new Measure object.
//
// https://doc.babylonjs.com/api/classes/babylon.measure
func (ba *Babylon) NewMeasure(left float64, top float64, width float64, height float64) *Measure {

	args := make([]interface{}, 0, 4+0)

	args = append(args, left)
	args = append(args, top)
	args = append(args, width)
	args = append(args, height)

	p := ba.ctx.Get("Measure").New(args...)
	return MeasureFromJSObject(p, ba.ctx)
}

// CombineToRef calls the CombineToRef method on the Measure object.
//
// https://doc.babylonjs.com/api/classes/babylon.measure#combinetoref
func (m *Measure) CombineToRef(a *Measure, b *Measure, result *Measure) {

	args := make([]interface{}, 0, 3+0)

	args = append(args, a.JSObject())
	args = append(args, b.JSObject())
	args = append(args, result.JSObject())

	m.p.Call("CombineToRef", args...)
}

// CopyFrom calls the CopyFrom method on the Measure object.
//
// https://doc.babylonjs.com/api/classes/babylon.measure#copyfrom
func (m *Measure) CopyFrom(other *Measure) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, other.JSObject())

	m.p.Call("copyFrom", args...)
}

// CopyFromFloats calls the CopyFromFloats method on the Measure object.
//
// https://doc.babylonjs.com/api/classes/babylon.measure#copyfromfloats
func (m *Measure) CopyFromFloats(left float64, top float64, width float64, height float64) {

	args := make([]interface{}, 0, 4+0)

	args = append(args, left)
	args = append(args, top)
	args = append(args, width)
	args = append(args, height)

	m.p.Call("copyFromFloats", args...)
}

// Empty calls the Empty method on the Measure object.
//
// https://doc.babylonjs.com/api/classes/babylon.measure#empty
func (m *Measure) Empty() *Measure {

	retVal := m.p.Call("Empty")
	return MeasureFromJSObject(retVal, m.ctx)
}

// IsEqualsTo calls the IsEqualsTo method on the Measure object.
//
// https://doc.babylonjs.com/api/classes/babylon.measure#isequalsto
func (m *Measure) IsEqualsTo(other *Measure) bool {

	args := make([]interface{}, 0, 1+0)

	args = append(args, other.JSObject())

	retVal := m.p.Call("isEqualsTo", args...)
	return retVal.Bool()
}

// TransformToRef calls the TransformToRef method on the Measure object.
//
// https://doc.babylonjs.com/api/classes/babylon.measure#transformtoref
func (m *Measure) TransformToRef(transform *Matrix2D, result *Measure) {

	args := make([]interface{}, 0, 2+0)

	args = append(args, transform.JSObject())
	args = append(args, result.JSObject())

	m.p.Call("transformToRef", args...)
}

// Height returns the Height property of class Measure.
//
// https://doc.babylonjs.com/api/classes/babylon.measure#height
func (m *Measure) Height() float64 {
	retVal := m.p.Get("height")
	return retVal.Float()
}

// SetHeight sets the Height property of class Measure.
//
// https://doc.babylonjs.com/api/classes/babylon.measure#height
func (m *Measure) SetHeight(height float64) *Measure {
	m.p.Set("height", height)
	return m
}

// Left returns the Left property of class Measure.
//
// https://doc.babylonjs.com/api/classes/babylon.measure#left
func (m *Measure) Left() float64 {
	retVal := m.p.Get("left")
	return retVal.Float()
}

// SetLeft sets the Left property of class Measure.
//
// https://doc.babylonjs.com/api/classes/babylon.measure#left
func (m *Measure) SetLeft(left float64) *Measure {
	m.p.Set("left", left)
	return m
}

// Top returns the Top property of class Measure.
//
// https://doc.babylonjs.com/api/classes/babylon.measure#top
func (m *Measure) Top() float64 {
	retVal := m.p.Get("top")
	return retVal.Float()
}

// SetTop sets the Top property of class Measure.
//
// https://doc.babylonjs.com/api/classes/babylon.measure#top
func (m *Measure) SetTop(top float64) *Measure {
	m.p.Set("top", top)
	return m
}

// Width returns the Width property of class Measure.
//
// https://doc.babylonjs.com/api/classes/babylon.measure#width
func (m *Measure) Width() float64 {
	retVal := m.p.Get("width")
	return retVal.Float()
}

// SetWidth sets the Width property of class Measure.
//
// https://doc.babylonjs.com/api/classes/babylon.measure#width
func (m *Measure) SetWidth(width float64) *Measure {
	m.p.Set("width", width)
	return m
}
