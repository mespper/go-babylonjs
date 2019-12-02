// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// MultiLinePoint represents a babylon.js MultiLinePoint.
// Class used to store a point for a MultiLine object.
// The point can be pure 2D coordinates, a mesh or a control
type MultiLinePoint struct {
	p   js.Value
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (m *MultiLinePoint) JSObject() js.Value { return m.p }

// MultiLinePoint returns a MultiLinePoint JavaScript class.
func (ba *Babylon) MultiLinePoint() *MultiLinePoint {
	p := ba.ctx.Get("MultiLinePoint")
	return MultiLinePointFromJSObject(p, ba.ctx)
}

// MultiLinePointFromJSObject returns a wrapped MultiLinePoint JavaScript class.
func MultiLinePointFromJSObject(p js.Value, ctx js.Value) *MultiLinePoint {
	return &MultiLinePoint{p: p, ctx: ctx}
}

// MultiLinePointArrayToJSArray returns a JavaScript Array for the wrapped array.
func MultiLinePointArrayToJSArray(array []*MultiLinePoint) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewMultiLinePoint returns a new MultiLinePoint object.
//
// https://doc.babylonjs.com/api/classes/babylon.multilinepoint
func (ba *Babylon) NewMultiLinePoint(multiLine *MultiLine) *MultiLinePoint {

	args := make([]interface{}, 0, 1+0)

	args = append(args, multiLine.JSObject())

	p := ba.ctx.Get("MultiLinePoint").New(args...)
	return MultiLinePointFromJSObject(p, ba.ctx)
}

// Dispose calls the Dispose method on the MultiLinePoint object.
//
// https://doc.babylonjs.com/api/classes/babylon.multilinepoint#dispose
func (m *MultiLinePoint) Dispose() {

	m.p.Call("dispose")
}

// ResetLinks calls the ResetLinks method on the MultiLinePoint object.
//
// https://doc.babylonjs.com/api/classes/babylon.multilinepoint#resetlinks
func (m *MultiLinePoint) ResetLinks() {

	m.p.Call("resetLinks")
}

// Translate calls the Translate method on the MultiLinePoint object.
//
// https://doc.babylonjs.com/api/classes/babylon.multilinepoint#translate
func (m *MultiLinePoint) Translate() *Vector2 {

	retVal := m.p.Call("translate")
	return Vector2FromJSObject(retVal, m.ctx)
}

// Control returns the Control property of class MultiLinePoint.
//
// https://doc.babylonjs.com/api/classes/babylon.multilinepoint#control
func (m *MultiLinePoint) Control() *Control {
	retVal := m.p.Get("control")
	return ControlFromJSObject(retVal, m.ctx)
}

// SetControl sets the Control property of class MultiLinePoint.
//
// https://doc.babylonjs.com/api/classes/babylon.multilinepoint#control
func (m *MultiLinePoint) SetControl(control *Control) *MultiLinePoint {
	m.p.Set("control", control.JSObject())
	return m
}

// Mesh returns the Mesh property of class MultiLinePoint.
//
// https://doc.babylonjs.com/api/classes/babylon.multilinepoint#mesh
func (m *MultiLinePoint) Mesh() *AbstractMesh {
	retVal := m.p.Get("mesh")
	return AbstractMeshFromJSObject(retVal, m.ctx)
}

// SetMesh sets the Mesh property of class MultiLinePoint.
//
// https://doc.babylonjs.com/api/classes/babylon.multilinepoint#mesh
func (m *MultiLinePoint) SetMesh(mesh *AbstractMesh) *MultiLinePoint {
	m.p.Set("mesh", mesh.JSObject())
	return m
}

// X returns the X property of class MultiLinePoint.
//
// https://doc.babylonjs.com/api/classes/babylon.multilinepoint#x
func (m *MultiLinePoint) X() string {
	retVal := m.p.Get("x")
	return retVal.String()
}

// SetX sets the X property of class MultiLinePoint.
//
// https://doc.babylonjs.com/api/classes/babylon.multilinepoint#x
func (m *MultiLinePoint) SetX(x string) *MultiLinePoint {
	m.p.Set("x", x)
	return m
}

// Y returns the Y property of class MultiLinePoint.
//
// https://doc.babylonjs.com/api/classes/babylon.multilinepoint#y
func (m *MultiLinePoint) Y() string {
	retVal := m.p.Get("y")
	return retVal.String()
}

// SetY sets the Y property of class MultiLinePoint.
//
// https://doc.babylonjs.com/api/classes/babylon.multilinepoint#y
func (m *MultiLinePoint) SetY(y string) *MultiLinePoint {
	m.p.Set("y", y)
	return m
}
