// Code generated by docs2go. DO NOT EDIT.

package babylon

import (
	"syscall/js"
)

// ScrollBar represents a babylon.js ScrollBar.
// Class used to create slider controls
type ScrollBar struct{ *BaseSlider }

// JSObject returns the underlying js.Value.
func (s *ScrollBar) JSObject() js.Value { return s.p }

// ScrollBar returns a ScrollBar JavaScript class.
func (ba *Babylon) ScrollBar() *ScrollBar {
	p := ba.ctx.Get("ScrollBar")
	return ScrollBarFromJSObject(p)
}

// ScrollBarFromJSObject returns a wrapped ScrollBar JavaScript class.
func ScrollBarFromJSObject(p js.Value) *ScrollBar {
	return &ScrollBar{BaseSliderFromJSObject(p)}
}

// NewScrollBarOpts contains optional parameters for NewScrollBar.
type NewScrollBarOpts struct {
	Name *JSString
}

// NewScrollBar returns a new ScrollBar object.
//
// https://doc.babylonjs.com/api/classes/babylon.scrollbar
func (ba *Babylon) NewScrollBar(opts *NewScrollBarOpts) *ScrollBar {
	if opts == nil {
		opts = &NewScrollBarOpts{}
	}

	p := ba.ctx.Get("ScrollBar").New(opts.Name.JSObject())
	return ScrollBarFromJSObject(p)
}

// TODO: methods
