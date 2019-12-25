// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// MultiLine represents a babylon.js MultiLine.
// Class used to create multi line control
type MultiLine struct {
	*Control
	ctx js.Value
}

// JSObject returns the underlying js.Value.
func (m *MultiLine) JSObject() js.Value { return m.p }

// MultiLine returns a MultiLine JavaScript class.
func (gui *GUI) MultiLine() *MultiLine {
	p := gui.ctx.Get("MultiLine")
	return MultiLineFromJSObject(p, gui.ctx)
}

// MultiLineFromJSObject returns a wrapped MultiLine JavaScript class.
func MultiLineFromJSObject(p js.Value, ctx js.Value) *MultiLine {
	return &MultiLine{Control: ControlFromJSObject(p, ctx), ctx: ctx}
}

// MultiLineArrayToJSArray returns a JavaScript Array for the wrapped array.
func MultiLineArrayToJSArray(array []*MultiLine) []interface{} {
	var result []interface{}
	for _, v := range array {
		result = append(result, v.JSObject())
	}
	return result
}

// NewMultiLineOpts contains optional parameters for NewMultiLine.
type NewMultiLineOpts struct {
	Name *string
}

// NewMultiLine returns a new MultiLine object.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.multiline
func (gui *GUI) NewMultiLine(opts *NewMultiLineOpts) *MultiLine {
	if opts == nil {
		opts = &NewMultiLineOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.Name == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, *opts.Name)
	}

	p := gui.ctx.Get("MultiLine").New(args...)
	return MultiLineFromJSObject(p, gui.ctx)
}

// Add calls the Add method on the MultiLine object.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.multiline#add
func (m *MultiLine) Add(items []*AbstractMesh) []*MultiLinePoint {

	args := make([]interface{}, 0, 1+0)

	args = append(args, AbstractMeshArrayToJSArray(items))

	retVal := m.p.Call("add", args...)
	result := []*MultiLinePoint{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, MultiLinePointFromJSObject(retVal.Index(ri), m.ctx))
	}
	return result
}

// Dispose calls the Dispose method on the MultiLine object.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.multiline#dispose
func (m *MultiLine) Dispose() {

	m.p.Call("dispose")
}

// GetAt calls the GetAt method on the MultiLine object.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.multiline#getat
func (m *MultiLine) GetAt(index float64) *MultiLinePoint {

	args := make([]interface{}, 0, 1+0)

	args = append(args, index)

	retVal := m.p.Call("getAt", args...)
	return MultiLinePointFromJSObject(retVal, m.ctx)
}

// MultiLinePushOpts contains optional parameters for MultiLine.Push.
type MultiLinePushOpts struct {
	Item *AbstractMesh
}

// Push calls the Push method on the MultiLine object.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.multiline#push
func (m *MultiLine) Push(opts *MultiLinePushOpts) *MultiLinePoint {
	if opts == nil {
		opts = &MultiLinePushOpts{}
	}

	args := make([]interface{}, 0, 0+1)

	if opts.Item == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.Item.JSObject())
	}

	retVal := m.p.Call("push", args...)
	return MultiLinePointFromJSObject(retVal, m.ctx)
}

// Remove calls the Remove method on the MultiLine object.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.multiline#remove
func (m *MultiLine) Remove(value float64) {

	args := make([]interface{}, 0, 1+0)

	args = append(args, value)

	m.p.Call("remove", args...)
}

// Reset calls the Reset method on the MultiLine object.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.multiline#reset
func (m *MultiLine) Reset() {

	m.p.Call("reset")
}

// ResetLinks calls the ResetLinks method on the MultiLine object.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.multiline#resetlinks
func (m *MultiLine) ResetLinks() {

	m.p.Call("resetLinks")
}

// MultiLine_drawOpts contains optional parameters for MultiLine._draw.
type MultiLine_drawOpts struct {
	InvalidatedRectangle *Measure
}

// _draw calls the _draw method on the MultiLine object.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.multiline#_draw
func (m *MultiLine) _draw(context js.Value, opts *MultiLine_drawOpts) {
	if opts == nil {
		opts = &MultiLine_drawOpts{}
	}

	args := make([]interface{}, 0, 1+1)

	args = append(args, context)

	if opts.InvalidatedRectangle == nil {
		args = append(args, js.Undefined())
	} else {
		args = append(args, opts.InvalidatedRectangle.JSObject())
	}

	m.p.Call("_draw", args...)
}

// _measure calls the _measure method on the MultiLine object.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.multiline#_measure
func (m *MultiLine) _measure() {

	m.p.Call("_measure")
}

// Dash returns the Dash property of class MultiLine.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.multiline#dash
func (m *MultiLine) Dash() []float64 {
	retVal := m.p.Get("dash")
	result := []float64{}
	for ri := 0; ri < retVal.Length(); ri++ {
		result = append(result, retVal.Index(ri).Float())
	}
	return result
}

// SetDash sets the Dash property of class MultiLine.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.multiline#dash
func (m *MultiLine) SetDash(dash []float64) *MultiLine {
	m.p.Set("dash", dash)
	return m
}

// HorizontalAlignment returns the HorizontalAlignment property of class MultiLine.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.multiline#horizontalalignment
func (m *MultiLine) HorizontalAlignment() float64 {
	retVal := m.p.Get("horizontalAlignment")
	return retVal.Float()
}

// SetHorizontalAlignment sets the HorizontalAlignment property of class MultiLine.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.multiline#horizontalalignment
func (m *MultiLine) SetHorizontalAlignment(horizontalAlignment float64) *MultiLine {
	m.p.Set("horizontalAlignment", horizontalAlignment)
	return m
}

// LineWidth returns the LineWidth property of class MultiLine.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.multiline#linewidth
func (m *MultiLine) LineWidth() float64 {
	retVal := m.p.Get("lineWidth")
	return retVal.Float()
}

// SetLineWidth sets the LineWidth property of class MultiLine.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.multiline#linewidth
func (m *MultiLine) SetLineWidth(lineWidth float64) *MultiLine {
	m.p.Set("lineWidth", lineWidth)
	return m
}

// Name returns the Name property of class MultiLine.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.multiline#name
func (m *MultiLine) Name() string {
	retVal := m.p.Get("name")
	return retVal.String()
}

// SetName sets the Name property of class MultiLine.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.multiline#name
func (m *MultiLine) SetName(name string) *MultiLine {
	m.p.Set("name", name)
	return m
}

// OnPointUpdate returns the OnPointUpdate property of class MultiLine.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.multiline#onpointupdate
func (m *MultiLine) OnPointUpdate() js.Value {
	retVal := m.p.Get("onPointUpdate")
	return retVal
}

// SetOnPointUpdate sets the OnPointUpdate property of class MultiLine.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.multiline#onpointupdate
func (m *MultiLine) SetOnPointUpdate(onPointUpdate JSFunc) *MultiLine {
	m.p.Set("onPointUpdate", js.FuncOf(onPointUpdate))
	return m
}

// VerticalAlignment returns the VerticalAlignment property of class MultiLine.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.multiline#verticalalignment
func (m *MultiLine) VerticalAlignment() float64 {
	retVal := m.p.Get("verticalAlignment")
	return retVal.Float()
}

// SetVerticalAlignment sets the VerticalAlignment property of class MultiLine.
//
// https://doc.babylonjs.com/api/classes/babylon.gui.multiline#verticalalignment
func (m *MultiLine) SetVerticalAlignment(verticalAlignment float64) *MultiLine {
	m.p.Set("verticalAlignment", verticalAlignment)
	return m
}
