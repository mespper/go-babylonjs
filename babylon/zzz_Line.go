// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// Line represents a babylon.js Line.
// Class used to render 2D lines
type Line struct{ *Control }

// JSObject returns the underlying js.Value.
func (l *Line) JSObject() js.Value { return l.p }

// Line returns a Line JavaScript class.
func (ba *Babylon) Line() *Line {
	p := ba.ctx.Get("Line")
	return LineFromJSObject(p)
}

// LineFromJSObject returns a wrapped Line JavaScript class.
func LineFromJSObject(p js.Value) *Line {
	return &Line{ControlFromJSObject(p)}
}

// NewLineOpts contains optional parameters for NewLine.
type NewLineOpts struct {
	Name *JSString
}

// NewLine returns a new Line object.
//
// https://doc.babylonjs.com/api/classes/babylon.line
func (ba *Babylon) NewLine(opts *NewLineOpts) *Line {
	if opts == nil {
		opts = &NewLineOpts{}
	}

	p := ba.ctx.Get("Line").New(opts.Name.JSObject())
	return LineFromJSObject(p)
}

// TODO: methods
